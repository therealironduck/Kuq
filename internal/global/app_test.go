package global

import (
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewApp(t *testing.T) {
	t.Parallel()

	t.Run("initializes the database and sets the variables", func(t *testing.T) {
		t.Parallel()

		tmpDir, err := os.MkdirTemp("", "testdir")
		require.NoError(t, err)

		t.Cleanup(func() {
			err := os.Remove(tmpDir)
			assert.NoError(t, err)
		})

		pathResolver := func(relative string) (path string, err error) {
			assert.Equal(t, "my/cool/path", relative)
			path = tmpDir

			return
		}

		dbPathResolver := func(relative string) (path string, err error) {
			assert.Equal(t, "my/cool/path", relative)
			path = ":memory:"

			return
		}

		app, err := NewApp("my/cool/path", &NewAppOptions{
			pathResolver:   pathResolver,
			dbPathResolver: dbPathResolver,
		})
		require.NoError(t, err)

		require.Equal(t, tmpDir, app.WorkspacePath)
		require.NotNil(t, app.DB)
	})

	t.Run("error handling", func(t *testing.T) {
		t.Parallel()

		t.Run("return error if workspace path cant be resolved", func(t *testing.T) {
			t.Parallel()

			pathResolver := func(_ string) (path string, err error) {
				err = errors.New("forced error") //nolint
				return
			}

			app, err := NewApp("my/cool/path", &NewAppOptions{pathResolver: pathResolver})
			require.Error(t, err)
			require.Nil(t, app)
		})

		t.Run("return error if database path cant be resolved", func(t *testing.T) {
			t.Parallel()

			dbPathResolver := func(_ string) (path string, err error) {
				err = errors.New("forced error") //nolint
				return
			}

			app, err := NewApp("my/cool/path", &NewAppOptions{dbPathResolver: dbPathResolver})
			require.Error(t, err)
			require.Nil(t, app)
		})

		t.Run("returns error if database cant be initialzed", func(t *testing.T) {
			t.Parallel()

			dbPathResolver := func(_ string) (path string, err error) {
				path = "invalid/data/path"
				return
			}

			app, err := NewApp("my/cool/path", &NewAppOptions{dbPathResolver: dbPathResolver})
			require.Error(t, err)
			require.Nil(t, app)
		})
	})
}
