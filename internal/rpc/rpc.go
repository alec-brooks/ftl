package rpc

import (
	"context"
	"crypto/tls"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/alecthomas/errors"
	"github.com/bufbuild/connect-go"
	"github.com/jpillora/backoff"
	"golang.org/x/net/http2"

	"github.com/TBD54566975/ftl/internal/log"
	ftlv1 "github.com/TBD54566975/ftl/protos/xyz/block/ftl/v1"
)

var (
	dialer = &net.Dialer{
		Timeout: time.Second * 10,
	}
	h2cClient = func() *http.Client {
		var netTransport = &http2.Transport{
			AllowHTTP: true,
			DialTLS: func(network, addr string, _ *tls.Config) (net.Conn, error) {
				conn, err := dialer.Dial(network, addr)
				return conn, errors.WithStack(err)
			},
		}
		return &http.Client{
			// We can't have a client-wide timeout because it also applies to
			// streaming RPCs, timing them out.
			// Timeout:   time.Second * 10,
			Transport: netTransport,
		}
	}()
	tlsClient = func() *http.Client {
		netTransport := &http2.Transport{
			DialTLSContext: func(ctx context.Context, network, addr string, config *tls.Config) (net.Conn, error) {
				tlsDialer := tls.Dialer{Config: config, NetDialer: dialer}
				conn, err := tlsDialer.DialContext(ctx, network, addr)
				return conn, errors.WithStack(err)
			},
		}
		return &http.Client{
			// We can't have a client-wide timeout because it also applies to
			// streaming RPCs, timing them out.
			// Timeout:   time.Second * 10,
			Transport: netTransport,
		}
	}()
)

type Pingable interface {
	Ping(context.Context, *connect.Request[ftlv1.PingRequest]) (*connect.Response[ftlv1.PingResponse], error)
}

// ClientFactory is a function that creates a new client and is typically one of
// the New*Client functions generated by protoc-gen-connect-go.
type ClientFactory[Client Pingable] func(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) Client

func Dial[Client Pingable](factory ClientFactory[Client], baseURL string, errorLevel log.Level, opts ...connect.ClientOption) Client {
	client := tlsClient
	if strings.HasPrefix(baseURL, "http://") {
		client = h2cClient
	}
	opts = append(opts, DefaultClientOptions(errorLevel)...)
	return factory(client, baseURL, opts...)
}

// ContextValuesMiddleware injects values from a Context into the request Context.
func ContextValuesMiddleware(ctx context.Context, handler http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r = r.WithContext(mergedContext{values: ctx, Context: r.Context()})
		handler.ServeHTTP(w, r)
	}
}

var _ context.Context = (*mergedContext)(nil)

type mergedContext struct {
	values context.Context
	context.Context
}

func (m mergedContext) Value(key any) any {
	if value := m.Context.Value(key); value != nil {
		return value
	}
	return m.values.Value(key)
}

// Wait for a client to become available.
//
// This will repeatedly call Ping() every 100ms until the service becomes
// ready. TODO: This will probably need to be smarter at some point.
//
// If "ctx" is cancelled this will return ctx.Err()
func Wait(ctx context.Context, retry backoff.Backoff, client Pingable) error {
	logger := log.FromContext(ctx)
	for {
		select {
		case <-ctx.Done():
			return errors.WithStack(ctx.Err())
		default:
		}
		resp, err := client.Ping(ctx, connect.NewRequest(&ftlv1.PingRequest{}))
		if err == nil {
			if resp.Msg.NotReady == nil {
				return nil
			}
			err = errors.Errorf("service is not ready: %s", *resp.Msg.NotReady)
		}
		delay := retry.Duration()
		logger.Tracef("Ping failed waiting %s for client: %+v", delay, err)
		time.Sleep(delay)
	}
}

// RetryStreamingClientStream will repeatedly call handler with the stream
// returned by rpc until handler returns nil or the context is cancelled.
//
// If the stream errors, it will be closed and a new call will be issued.
func RetryStreamingClientStream[Req, Resp any](
	ctx context.Context,
	retry backoff.Backoff,
	rpc func(context.Context) *connect.ClientStreamForClient[Req, Resp],
	handler func(ctx context.Context, send func(*Req) error) error,
) {
	logger := log.FromContext(ctx)
	for {
		stream := rpc(ctx)
		var err error
		for {
			err = handler(ctx, stream.Send)
			if err != nil {
				break
			}
			select {
			case <-ctx.Done():
				return
			default:
			}
			retry.Reset()
		}
		_, _ = stream.CloseAndReceive()

		delay := retry.Duration()
		logger.Warnf("Stream handler failed, retrying in %s: %s", delay, err)
		select {
		case <-ctx.Done():
			return

		case <-time.After(delay):
		}

	}
}