//go:build integration

package simple_test

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"connectrpc.com/connect"
	"github.com/alecthomas/assert/v2"
	"github.com/alecthomas/repr"
	"golang.org/x/sync/errgroup"

	ftlv1 "github.com/TBD54566975/ftl/backend/protos/xyz/block/ftl/v1"
	schemapb "github.com/TBD54566975/ftl/backend/protos/xyz/block/ftl/v1/schema"
)

func TestCron(t *testing.T) {
	dir := t.TempDir()
	// Due to some MacOS magic, /tmp differs between this test code and the
	// executing module, so we need to pass the file path as an environment
	// variable.
	tmpFile := filepath.Join(dir, "cron.txt")
	t.Setenv("DEST_FILE", tmpFile)

	t.Cleanup(func() { _ = os.Remove(tmpFile) })

	run(t, "",
		copyModule("cron"),
		deploy("cron"),
		func(t testing.TB, ic testContext) error {
			_, err := os.Stat(tmpFile)
			return err
		},
	)
}

func TestLifecycle(t *testing.T) {
	run(t, "",
		exec("ftl", "init", "go", ".", "echo"),
		deploy("echo"),
		call("echo", "echo", obj{"name": "Bob"}, func(response obj) error {
			if response["message"] != "Hello, Bob!" {
				return fmt.Errorf("unexpected response: %v", response)
			}
			return nil
		}),
	)
}

func TestInterModuleCall(t *testing.T) {
	run(t, "",
		copyModule("echo"),
		copyModule("time"),
		deploy("time"),
		deploy("echo"),
		call("echo", "echo", obj{"name": "Bob"}, func(response obj) error {
			message, ok := response["message"].(string)
			if !ok {
				return fmt.Errorf("unexpected response: %v", response)
			}
			if !strings.HasPrefix(message, "Hello, Bob!!! It is ") {
				return fmt.Errorf("unexpected response: %v", response)
			}
			return nil
		}),
	)
}

func TestNonExportedDecls(t *testing.T) {
	run(t, "",
		copyModule("time"),
		deploy("time"),
		copyModule("echo"),
		deploy("echo"),
		copyModule("notexportedverb"),
		expectError(execWithOutput("ftl", "deploy", "notexportedverb"), "call first argument must be a function but is an unresolved reference to echo.Echo, does it need to be exported?"),
	)
}

func TestUndefinedExportedDecls(t *testing.T) {
	run(t, "",
		copyModule("time"),
		deploy("time"),
		copyModule("echo"),
		deploy("echo"),
		copyModule("undefinedverb"),
		expectError(execWithOutput("ftl", "deploy", "undefinedverb"), "call first argument must be a function but is an unresolved reference to echo.Undefined"),
	)
}

func TestDatabase(t *testing.T) {
	run(t, "database/ftl-project.toml",
		// deploy real module against "testdb"
		copyModule("database"),
		createDBAction("database", "testdb", false),
		deploy("database"),
		call("database", "insert", obj{"data": "hello"}, func(response obj) error { return nil }),
		queryRow("testdb", "SELECT data FROM requests", "hello"),

		// run tests which should only affect "testdb_test"
		createDBAction("database", "testdb", true),
		testModule("database"),
		queryRow("testdb", "SELECT data FROM requests", "hello"),
	)
}

func TestSchemaGenerate(t *testing.T) {
	run(t, "",
		copyDir("../schema-generate", "schema-generate"),
		mkdir("build/schema-generate"),
		exec("ftl", "schema", "generate", "schema-generate", "build/schema-generate"),
		fileContains("build/schema-generate/test.txt", "olleh"),
	)
}

func TestHttpEncodeOmitempty(t *testing.T) {
	run(t, "",
		copyModule("omitempty"),
		deploy("omitempty"),
		httpCall(http.MethodGet, "/get", jsonData(t, obj{}), func(resp *httpResponse) error {
			assert.Equal(t, 200, resp.status)
			_, ok := resp.jsonBody["mustset"]
			assert.True(t, ok)
			_, ok = resp.jsonBody["error"]
			assert.False(t, ok)
			return nil
		}),
	)
}

func TestHttpIngress(t *testing.T) {
	run(t, "",
		copyModule("httpingress"),
		deploy("httpingress"),
		httpCall(http.MethodGet, "/users/123/posts/456", jsonData(t, obj{}), func(resp *httpResponse) error {
			assert.Equal(t, 200, resp.status)
			assert.Equal(t, []string{"Header from FTL"}, resp.headers["Get"])
			assert.Equal(t, []string{"application/json; charset=utf-8"}, resp.headers["Content-Type"])

			message, ok := resp.jsonBody["msg"].(string)
			assert.True(t, ok, "msg is not a string: %s", repr.String(resp.jsonBody))
			assert.Equal(t, "UserID: 123, PostID: 456", message)

			nested, ok := resp.jsonBody["nested"].(map[string]any)
			assert.True(t, ok, "nested is not a map: %s", repr.String(resp.jsonBody))
			goodStuff, ok := nested["good_stuff"].(string)
			assert.True(t, ok, "good_stuff is not a string: %s", repr.String(resp.jsonBody))
			assert.Equal(t, "This is good stuff", goodStuff)
			return nil
		}),
		httpCall(http.MethodPost, "/users", jsonData(t, obj{"userId": 123, "postId": 345}), func(resp *httpResponse) error {
			assert.Equal(t, 201, resp.status)
			assert.Equal(t, []string{"Header from FTL"}, resp.headers["Post"])
			success, ok := resp.jsonBody["success"].(bool)
			assert.True(t, ok, "success is not a bool: %s", repr.String(resp.jsonBody))
			assert.True(t, success)
			return nil
		}),
		// contains aliased field
		httpCall(http.MethodPost, "/users", jsonData(t, obj{"user_id": 123, "postId": 345}), func(resp *httpResponse) error {
			assert.Equal(t, 201, resp.status)
			return nil
		}),
		httpCall(http.MethodPut, "/users/123", jsonData(t, obj{"postId": "346"}), func(resp *httpResponse) error {
			assert.Equal(t, 200, resp.status)
			assert.Equal(t, []string{"Header from FTL"}, resp.headers["Put"])
			assert.Equal(t, map[string]any{}, resp.jsonBody)
			return nil
		}),
		httpCall(http.MethodDelete, "/users/123", jsonData(t, obj{}), func(resp *httpResponse) error {
			assert.Equal(t, 200, resp.status)
			assert.Equal(t, []string{"Header from FTL"}, resp.headers["Delete"])
			assert.Equal(t, map[string]any{}, resp.jsonBody)
			return nil
		}),

		httpCall(http.MethodGet, "/html", jsonData(t, obj{}), func(resp *httpResponse) error {
			assert.Equal(t, 200, resp.status)
			assert.Equal(t, []string{"text/html; charset=utf-8"}, resp.headers["Content-Type"])
			assert.Equal(t, "<html><body><h1>HTML Page From FTL 🚀!</h1></body></html>", string(resp.bodyBytes))
			return nil
		}),

		httpCall(http.MethodPost, "/bytes", []byte("Hello, World!"), func(resp *httpResponse) error {
			assert.Equal(t, 200, resp.status)
			assert.Equal(t, []string{"application/octet-stream"}, resp.headers["Content-Type"])
			assert.Equal(t, []byte("Hello, World!"), resp.bodyBytes)
			return nil
		}),

		httpCall(http.MethodGet, "/empty", nil, func(resp *httpResponse) error {
			assert.Equal(t, 200, resp.status)
			assert.Equal(t, nil, resp.headers["Content-Type"])
			assert.Equal(t, nil, resp.bodyBytes)
			return nil
		}),

		httpCall(http.MethodGet, "/string", []byte("Hello, World!"), func(resp *httpResponse) error {
			assert.Equal(t, 200, resp.status)
			assert.Equal(t, []string{"text/plain; charset=utf-8"}, resp.headers["Content-Type"])
			assert.Equal(t, []byte("Hello, World!"), resp.bodyBytes)
			return nil
		}),

		httpCall(http.MethodGet, "/int", []byte("1234"), func(resp *httpResponse) error {
			assert.Equal(t, 200, resp.status)
			assert.Equal(t, []string{"text/plain; charset=utf-8"}, resp.headers["Content-Type"])
			assert.Equal(t, []byte("1234"), resp.bodyBytes)
			return nil
		}),
		httpCall(http.MethodGet, "/float", []byte("1234.56789"), func(resp *httpResponse) error {
			assert.Equal(t, 200, resp.status)
			assert.Equal(t, []string{"text/plain; charset=utf-8"}, resp.headers["Content-Type"])
			assert.Equal(t, []byte("1234.56789"), resp.bodyBytes)
			return nil
		}),
		httpCall(http.MethodGet, "/bool", []byte("true"), func(resp *httpResponse) error {
			assert.Equal(t, 200, resp.status)
			assert.Equal(t, []string{"text/plain; charset=utf-8"}, resp.headers["Content-Type"])
			assert.Equal(t, []byte("true"), resp.bodyBytes)
			return nil
		}),
		httpCall(http.MethodGet, "/error", nil, func(resp *httpResponse) error {
			assert.Equal(t, 500, resp.status)
			assert.Equal(t, []string{"text/plain; charset=utf-8"}, resp.headers["Content-Type"])
			assert.Equal(t, []byte("Error from FTL"), resp.bodyBytes)
			return nil
		}),
		httpCall(http.MethodGet, "/array/string", jsonData(t, []string{"hello", "world"}), func(resp *httpResponse) error {
			assert.Equal(t, 200, resp.status)
			assert.Equal(t, []string{"application/json; charset=utf-8"}, resp.headers["Content-Type"])
			assert.Equal(t, jsonData(t, []string{"hello", "world"}), resp.bodyBytes)
			return nil
		}),
		httpCall(http.MethodPost, "/array/data", jsonData(t, []obj{{"item": "a"}, {"item": "b"}}), func(resp *httpResponse) error {
			assert.Equal(t, 200, resp.status)
			assert.Equal(t, []string{"application/json; charset=utf-8"}, resp.headers["Content-Type"])
			assert.Equal(t, jsonData(t, []obj{{"item": "a"}, {"item": "b"}}), resp.bodyBytes)
			return nil
		}),
		httpCall(http.MethodGet, "/typeenum", jsonData(t, obj{"name": "A", "value": "hello"}), func(resp *httpResponse) error {
			assert.Equal(t, 200, resp.status)
			assert.Equal(t, []string{"application/json; charset=utf-8"}, resp.headers["Content-Type"])
			assert.Equal(t, jsonData(t, obj{"name": "A", "value": "hello"}), resp.bodyBytes)
			return nil
		}),
	)
}

func TestRuntimeReflection(t *testing.T) {
	run(t, "",
		copyModule("runtimereflection"),
		testModule("runtimereflection"),
	)
}

func TestModuleUnitTests(t *testing.T) {
	run(t, "",
		copyModule("time"),
		copyModule("wrapped"),
		copyModule("verbtypes"),
		build("time", "wrapped", "verbtypes"),
		testModule("wrapped"),
		testModule("verbtypes"),
	)
}

func TestLease(t *testing.T) {
	run(t, "",
		copyModule("leases"),
		build("leases"),
		// checks if leases work in a unit test environment
		testModule("leases"),
		deploy("leases"),
		// checks if it leases work with a real controller
		func(t testing.TB, ic testContext) error {
			// Start a lease.
			wg := errgroup.Group{}
			wg.Go(func() error {
				infof("Acquiring lease")
				resp, err := ic.verbs.Call(ic, connect.NewRequest(&ftlv1.CallRequest{
					Verb: &schemapb.Ref{Module: "leases", Name: "acquire"},
					Body: []byte("{}"),
				}))
				if respErr := resp.Msg.GetError(); respErr != nil {
					return fmt.Errorf("received error on first call: %v", respErr)
				}
				return err
			})

			time.Sleep(time.Second)

			infof("Trying to acquire lease again")
			// Trying to obtain the lease again should fail.
			resp, err := ic.verbs.Call(ic, connect.NewRequest(&ftlv1.CallRequest{
				Verb: &schemapb.Ref{Module: "leases", Name: "acquire"},
				Body: []byte("{}"),
			}))
			if err != nil {
				return err
			}
			if resp.Msg.GetError() == nil || !strings.Contains(resp.Msg.GetError().Message, "could not acquire lease") {
				return fmt.Errorf("expected error but got: %#v", resp.Msg.GetError())
			}
			return wg.Wait()
		},
	)
}
