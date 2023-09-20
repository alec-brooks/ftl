package model

import (
	"crypto/rand"
	"database/sql"
	"database/sql/driver"
	"encoding"
	"encoding/hex"
	"fmt"
	"regexp"
	"strings"

	"github.com/alecthomas/errors"
	"github.com/alecthomas/types"
)

// A RequestName represents an inbound request into the cluster.
type RequestName string

type MaybeRequestName types.Option[RequestName]

var _ interface {
	sql.Scanner
	driver.Valuer
	encoding.TextUnmarshaler
	encoding.TextMarshaler
} = (*RequestName)(nil)

type Origin string

const (
	OriginIngress Origin = "ingress"
	OriginCron    Origin = "cron"
	OriginPubsub  Origin = "pubsub"
)

func ParseOrigin(origin string) (Origin, error) {
	switch origin {
	case "ingress":
		return OriginIngress, nil
	case "cron":
		return OriginCron, nil
	case "pubsub":
		return OriginPubsub, nil
	default:
		return "", errors.Errorf("unknown origin %q", origin)
	}
}

var requestNameNormaliserRe = regexp.MustCompile("[^a-zA-Z0-9]+")

func NewRequestName(origin Origin, key string) RequestName {
	hash := make([]byte, 5)
	_, err := rand.Read(hash)
	if err != nil {
		panic(err)
	}
	key = requestNameNormaliserRe.ReplaceAllString(key, "-")
	key = strings.ToLower(key)
	return RequestName(fmt.Sprintf("%s-%s-%010x", origin, key, hash))
}

func ParseRequestName(name string) (Origin, RequestName, error) {
	parts := strings.Split(name, "-")
	if len(parts) < 3 {
		return "", "", errors.Errorf("should be <origin>-<key>-<hash>: invalid request name %q", name)
	}
	origin, err := ParseOrigin(parts[0])
	if err != nil {
		return "", "", errors.Wrapf(err, "invalid request name %q", name)
	}
	hash, err := hex.DecodeString(parts[len(parts)-1])
	if err != nil {
		return "", "", errors.Wrapf(err, "invalid request name %q", name)
	}
	if len(hash) != 5 {
		return "", "", errors.Errorf("hash should be 5 bytes: invalid request name %q", name)
	}
	return origin, RequestName(fmt.Sprintf("%s-%010x", strings.Join(parts[0:len(parts)-1], "-"), hash)), nil
}

func (d *RequestName) String() string {
	return string(*d)
}

func (d *RequestName) UnmarshalText(bytes []byte) error {
	_, name, err := ParseRequestName(string(bytes))
	if err != nil {
		return err
	}
	*d = name
	return nil
}

func (d *RequestName) MarshalText() ([]byte, error) {
	return []byte(*d), nil
}

func (d *RequestName) Scan(value any) error {
	if value == nil {
		return nil
	}
	str, ok := value.(string)
	if !ok {
		return errors.Errorf("expected string, got %T", value)
	}
	_, name, err := ParseRequestName(str)
	if err != nil {
		return err
	}
	*d = name
	return nil
}

func (d *RequestName) Value() (driver.Value, error) {
	return d.String(), nil
}