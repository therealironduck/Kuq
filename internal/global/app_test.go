package global

import (
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewApp(t *testing.T) { //nolint tparallel
	t.Run("initializes the database and sets the variables", func(t *testing.T) {
		t.Parallel()

		originalPathResolver := pathResolver
		originalDbPathResolver := dbPathResolver

		t.Cleanup(func() {
			pathResolver = originalPathResolver
			dbPathResolver = originalDbPathResolver
		})

		tmpDir, err := os.MkdirTemp("", "testdir")
		require.NoError(t, err)

		t.Cleanup(func() {
			err := os.Remove(tmpDir)
			assert.NoError(t, err)
		})

		pathResolver = func(relative string) (path string, err error) {
			assert.Equal(t, "my/cool/path", relative)
			path = tmpDir

			return
		}

		dbPathResolver = func(relative string) (path string, err error) {
			assert.Equal(t, "my/cool/path", relative)
			path = ":memory:"

			return
		}

		app, err := New("my/cool/path")
		require.NoError(t, err)

		require.Equal(t, tmpDir, app.WorkspacePath)
		require.NotNil(t, app.DB)
	})

	t.Run("error handling", func(t *testing.T) { //nolint tparallel
		t.Run("return error if workspace path cant be resolved", func(t *testing.T) { //nolint tparallel
			originalPathResolver := pathResolver

			t.Cleanup(func() {
				pathResolver = originalPathResolver
			})

			pathResolver = func(_ string) (path string, err error) {
				err = errors.New("forced error") //nolint
				return
			}

			app, err := New("my/cool/path")
			require.Error(t, err)
			require.Nil(t, app)
		})

		t.Run("return error if database path cant be resolved", func(t *testing.T) { //nolint tparallel
			originalDbPathResolver := dbPathResolver

			t.Cleanup(func() {
				dbPathResolver = originalDbPathResolver
			})

			dbPathResolver = func(_ string) (path string, err error) {
				err = errors.New("forced error") //nolint
				return
			}

			app, err := New("my/cool/path")
			require.Error(t, err)
			require.Nil(t, app)
		})

		t.Run("returns error if database cant be initialzed", func(t *testing.T) { //nolint tparallel
			originalDbPathResolver := dbPathResolver

			t.Cleanup(func() {
				dbPathResolver = originalDbPathResolver
			})

			dbPathResolver = func(_ string) (path string, err error) {
				path = "invalid/data/path"
				return
			}

			app, err := New("my/cool/path")
			require.Error(t, err)
			require.Nil(t, app)
		})
	})
}
