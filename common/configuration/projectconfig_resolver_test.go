package configuration

import (
	"context"
	"net/url"
	"os"
	"path/filepath"
	"testing"

	"github.com/alecthomas/assert/v2"
	"github.com/alecthomas/types/optional"

	"github.com/TBD54566975/ftl/common/projectconfig"
	"github.com/TBD54566975/ftl/internal/log"
)

func TestSet(t *testing.T) {
	defaultPath := projectconfig.GetDefaultConfigPath()
	origConfigBytes, err := os.ReadFile(defaultPath)
	assert.NoError(t, err)

	config := filepath.Join(t.TempDir(), "ftl-project.toml")
	existing, err := os.ReadFile("testdata/ftl-project.toml")
	assert.NoError(t, err)
	err = os.WriteFile(config, existing, 0600)
	assert.NoError(t, err)

	t.Run("ExistingModule", func(t *testing.T) {
		setAndAssert(t, "echo", []string{config})
	})
	t.Run("NewModule", func(t *testing.T) {
		setAndAssert(t, "echooo", []string{config})
	})
	t.Run("MissingTOMLFile", func(t *testing.T) {
		err := os.Remove(config)
		assert.NoError(t, err)
		setAndAssert(t, "echooooo", []string{})
		err = os.WriteFile(defaultPath, origConfigBytes, 0600)
		assert.NoError(t, err)
	})
}

func setAndAssert(t *testing.T, module string, config []string) {
	t.Helper()

	ctx := log.ContextWithNewDefaultLogger(context.Background())

	cf, err := New(ctx,
		ProjectConfigResolver[Configuration]{Config: config},
		[]Provider[Configuration]{
			EnvarProvider[Configuration]{},
			InlineProvider[Configuration]{Inline: true}, // Writer
		})
	assert.NoError(t, err)

	var got *url.URL
	want := URL("inline://asdfasdf")
	err = cf.Set(ctx, Ref{Module: optional.Some[string](module), Name: "default"}, want)
	assert.NoError(t, err)
	err = cf.Get(ctx, Ref{Module: optional.Some[string](module), Name: "default"}, &got)
	assert.NoError(t, err)
	assert.Equal(t, want, got)
}
