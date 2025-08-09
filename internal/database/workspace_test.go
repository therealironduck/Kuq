package database

import (
	"errors"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWorkspaceResolve(t *testing.T) { // nolint
	t.Run("resolves the workspace absolute path when given a `.`", func(t *testing.T) { //nolint
		localPath, err := filepath.Abs(".")
		require.NoError(t, err)

		absolute, err := ResolveWorkspacePath(".")
		require.NoError(t, err)

		require.Equal(t, localPath, absolute)
	})

	t.Run("returns the absolute path if an absolute path was given", func(t *testing.T) { //nolint
		const path = "/home/duck/ducky"
		absolute, err := ResolveWorkspacePath(path)
		require.NoError(t, err)
		require.Equal(t, path, absolute)
	})

	t.Run("returns an error if something can't be resolved", func(t *testing.T) { //nolint
		originalAbs := abs
		t.Cleanup(func() {
			abs = originalAbs
		})

		abs = func(_ string) (string, error) {
			return "", errors.New("forced error") //nolint
		}

		_, err := ResolveWorkspacePath(".")
		require.Error(t, err)
	})
}

func TestDatabasePath(t *testing.T) { //nolint
	t.Run("converts and return the database path", func(t *testing.T) { //nolint
		localPath, err := filepath.Abs(".")
		require.NoError(t, err)

		result, err := GetDatabasePath(localPath)
		require.NoError(t, err)
		require.Equal(t, filepath.Join(localPath, fileName), result)
	})

	t.Run("returns errors while resolving", func(t *testing.T) { //nolint
		originalAbs := abs
		t.Cleanup(func() {
			abs = originalAbs
		})

		abs = func(_ string) (string, error) {
			return "", errors.New("forced error") //nolint
		}

		result, err := GetDatabasePath("invalid")
		require.Error(t, err)
		require.Empty(t, result)
	})
}
