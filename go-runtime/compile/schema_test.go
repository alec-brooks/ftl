package compile

import (
	"context"
	"go/token"
	"go/types"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/alecthomas/assert/v2"
	"github.com/alecthomas/participle/v2/lexer"

	"github.com/TBD54566975/ftl/backend/schema"
	"github.com/TBD54566975/ftl/internal/errors"
	"github.com/TBD54566975/ftl/internal/exec"
	"github.com/TBD54566975/ftl/internal/log"
	"github.com/TBD54566975/ftl/internal/slices"
)

// this is helpful when a test requires another module to be built before running
// eg: when module A depends on module B, we need to build module B before building module A
func prebuildTestModule(t *testing.T, args ...string) {
	t.Helper()

	ctx := log.ContextWithNewDefaultLogger(context.Background())

	dir, err := os.Getwd()
	assert.NoError(t, err, "Failed to get current directory: %v", err)

	ftlArgs := []string{"build"}
	ftlArgs = append(ftlArgs, args...)

	cmd := exec.Command(ctx, log.Debug, dir, "ftl", ftlArgs...)
	err = cmd.RunBuffered(ctx)
	assert.NoError(t, err, "ftl build failed with %s\n", err)
}

func TestExtractModuleSchema(t *testing.T) {
	if testing.Short() {
		t.SkipNow()
	}
	prebuildTestModule(t, "testdata/one", "testdata/two")

	_, actual, _, err := ExtractModuleSchema("testdata/one")
	assert.NoError(t, err)
	actual = schema.Normalise(actual)
	expected := `module one {
  config configValue one.Config
  secret secretValue String

  database testDb

  enum Color(String) {
    Red("Red")
    Blue("Blue")
    Green("Green")
    Yellow("Yellow")
  }

  // Comments about ColorInt.
  enum ColorInt(Int) {
    // RedInt is a color.
    RedInt(0)
    BlueInt(1)
    // GreenInt is also a color.
    GreenInt(2)
    YellowInt(3)
  }

  enum IotaExpr(Int) {
    First(1)
    Second(3)
    Third(5)
  }

  enum SimpleIota(Int) {
    Zero(0)
    One(1)
    Two(2)
  }

  data Config {
    field String
  }

  data Nested {
  }

  data Req {
    int Int
    int64 Int
    float Float
    string String
    slice [String]
    map {String: String}
    nested one.Nested
    optional one.Nested?
    time Time
    user two.User +alias json "u"
    bytes Bytes
    localEnumRef one.Color
    externalEnumRef two.TwoEnum
  }

  data Resp {
  }

  data SinkReq {
  }

  data SourceResp {
  }

  verb nothing(Unit) Unit

  verb sink(one.SinkReq) Unit

  verb source(Unit) one.SourceResp

  verb verb(one.Req) one.Resp
}
`
	assert.Equal(t, expected, actual.String())
}

func TestExtractModuleSchemaTwo(t *testing.T) {
	if testing.Short() {
		t.SkipNow()
	}
	_, actual, _, err := ExtractModuleSchema("testdata/two")
	assert.NoError(t, err)
	actual = schema.Normalise(actual)
	expected := `module two {
		enum TwoEnum(String) {
		  Red("Red")
		  Blue("Blue")
		  Green("Green")
        }

		data Exported {
		}

		data Payload<T> {
		  body T
		}

		data User {
		  name String
		}

		data UserResponse {
		  user two.User
		}

		verb callsTwo(two.Payload<String>) two.Payload<String>
			+calls two.two

		verb returnsUser(Unit) two.UserResponse

		verb two(two.Payload<String>) two.Payload<String>
	  }
`
	assert.Equal(t, normaliseString(expected), normaliseString(actual.String()))
}

func TestParseDirectives(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected directive
	}{
		{name: "Export", input: "ftl:export", expected: &directiveExport{Export: true}},
		{name: "Ingress", input: `ftl:ingress GET /foo`, expected: &directiveIngress{
			Method: "GET",
			Path: []schema.IngressPathComponent{
				&schema.IngressPathLiteral{
					Text: "foo",
				},
			},
		}},
		{name: "Ingress", input: `ftl:ingress GET /test_path/{something}/987-Your_File.txt%7E%21Misc%2A%28path%29info%40abc%3Fxyz`, expected: &directiveIngress{
			Method: "GET",
			Path: []schema.IngressPathComponent{
				&schema.IngressPathLiteral{
					Text: "test_path",
				},
				&schema.IngressPathParameter{
					Name: "something",
				},
				&schema.IngressPathLiteral{
					Text: "987-Your_File.txt%7E%21Misc%2A%28path%29info%40abc%3Fxyz",
				},
			},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := directiveParser.ParseString("", tt.input)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, got.Directive, assert.Exclude[lexer.Position](), assert.Exclude[schema.Position]())
		})
	}
}

func TestParseTypesTime(t *testing.T) {
	timeRef := mustLoadRef("time", "Time").Type()
	parsed, ok := visitType(nil, token.NoPos, timeRef).Get()
	assert.True(t, ok)
	_, ok = parsed.(*schema.Time)
	assert.True(t, ok)
}

func TestParseBasicTypes(t *testing.T) {
	tests := []struct {
		name     string
		input    types.Type
		expected schema.Type
	}{
		{name: "String", input: types.Typ[types.String], expected: &schema.String{}},
		{name: "Int", input: types.Typ[types.Int], expected: &schema.Int{}},
		{name: "Bool", input: types.Typ[types.Bool], expected: &schema.Bool{}},
		{name: "Float64", input: types.Typ[types.Float64], expected: &schema.Float{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parsed, ok := visitType(nil, token.NoPos, tt.input).Get()
			assert.True(t, ok)
			assert.Equal(t, tt.expected, parsed)
		})
	}
}

func normaliseString(s string) string {
	return strings.TrimSpace(strings.Join(slices.Map(strings.Split(s, "\n"), strings.TrimSpace), "\n"))
}

func TestErrorReporting(t *testing.T) {
	if testing.Short() {
		t.SkipNow()
	}
	pwd, _ := os.Getwd()
	_, _, schemaErrs, err := ExtractModuleSchema("testdata/failing")
	assert.NoError(t, err)

	filename := filepath.Join(pwd, `testdata/failing/failing.go`)
	assert.EqualError(t, errors.Join(genericizeErrors(schemaErrs)...),
		filename+":10:13-35: config and secret declarations must have a single string literal argument\n"+
			filename+":13:18-52: duplicate config declaration at 12:18-52\n"+
			filename+":16:18-49: duplicate secret declaration at 15:18-49\n"+
			filename+":19:2-10: unsupported type \"error\" for field \"BadParam\"\n"+
			filename+":22:2-17: unsupported type \"uint64\" for field \"AnotherBadParam\"\n"+
			filename+":25:3-3: unexpected token \"verb\" (expected Directive)\n"+
			filename+":31:36-39: unsupported request type \"ftl/failing.Request\"\n"+
			filename+":31:50-50: unsupported response type \"ftl/failing.Response\"\n"+
			filename+":32:16-29: call first argument must be a function in an ftl module\n"+
			filename+":33:2-46: call must have exactly three arguments\n"+
			filename+":34:16-25: call first argument must be a function\n"+
			filename+":39:1-2: must have at most two parameters (context.Context, struct)\n"+
			filename+":39:69-69: unsupported response type \"ftl/failing.Response\"\n"+
			filename+":44:22-27: first parameter must be of type context.Context but is ftl/failing.Request\n"+
			filename+":44:37-43: second parameter must be a struct but is string\n"+
			filename+":44:53-53: unsupported response type \"ftl/failing.Response\"\n"+
			filename+":49:43-47: second parameter must not be ftl.Unit\n"+
			filename+":49:59-59: unsupported response type \"ftl/failing.Response\"\n"+
			filename+":54:1-2: first parameter must be context.Context\n"+
			filename+":54:18-18: unsupported response type \"ftl/failing.Response\"\n"+
			filename+":59:1-2: must have at most two results (struct, error)\n"+
			filename+":59:41-44: unsupported request type \"ftl/failing.Request\"\n"+
			filename+":64:1-2: must at least return an error\n"+
			filename+":64:36-39: unsupported request type \"ftl/failing.Request\"\n"+
			filename+":68:35-38: unsupported request type \"ftl/failing.Request\"\n"+
			filename+":68:48-48: must return an error but is ftl/failing.Response\n"+
			filename+":73:41-44: unsupported request type \"ftl/failing.Request\"\n"+
			filename+":73:55-55: first result must be a struct but is string\n"+
			filename+":73:63-63: must return an error but is string\n"+
			filename+":73:63-63: second result must not be ftl.Unit\n"+
			filename+":80:1-1: verb \"WrongResponse\" already exported\n"+
			filename+":86:2-12: struct field unexported must be exported by starting with an uppercase letter",
	)
}

func TestDuplicateVerbNames(t *testing.T) {
	if testing.Short() {
		t.SkipNow()
	}
	pwd, _ := os.Getwd()
	_, _, schemaErrs, err := ExtractModuleSchema("testdata/duplicateverbs")
	assert.NoError(t, err)
	assert.EqualError(t, errors.Join(genericizeErrors(schemaErrs)...), filepath.Join(pwd, `testdata/duplicateverbs/duplicateverbs.go`)+`:23:1-1: verb "Time" already exported`)
}

func genericizeErrors(schemaErrs []*schema.Error) []error {
	errs := make([]error, len(schemaErrs))
	for i, schemaErr := range schemaErrs {
		errs[i] = schemaErr
	}
	return errs
}
