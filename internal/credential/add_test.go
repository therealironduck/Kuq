package credential

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/therealironduck/kuq/internal/database"
	"github.com/therealironduck/kuq/internal/global"
	"github.com/therealironduck/kuq/internal/model"
	"gorm.io/gorm"
)

func NewMockApp(t *testing.T) (*global.App, func()) {
	t.Helper()

	tmpFile, err := os.MkdirTemp("", "testdir")
	require.NoError(t, err)

	db, err := database.New(":memory:")
	require.NoError(t, err)

	app := &global.App{
		WorkspacePath: tmpFile,
		DB:            db,
	}

	cleanup := func() {
		os.RemoveAll(tmpFile)
	}

	return app, cleanup
}

func TestAddCredentials(t *testing.T) {
	t.Parallel()

	t.Run("adds a valid new credential", func(t *testing.T) {
		t.Parallel()

		app, cleanup := NewMockApp(t)
		t.Cleanup(cleanup)

		id, err := Add(context.Background(), app, AddOptions{Name: "Jordan's Key", SSHKey: "ssh-123"})
		require.NoError(t, err)

		credential, err := gorm.G[model.Credential](app.DB).Where("id = ?", id).First(context.Background())
		require.NoError(t, err)

		assert.Equal(t, "Jordan's Key", credential.Name)
		assert.Equal(t, "ssh-123", credential.SSHKey)
		assert.Equal(t, id, credential.ID)
	})

	t.Run("the name is unique", func(t *testing.T) {
		t.Parallel()

		app, cleanup := NewMockApp(t)
		t.Cleanup(cleanup)

		_, err := Add(context.Background(), app, AddOptions{Name: "Jordan's Key", SSHKey: "ssh-123"})
		require.NoError(t, err)

		id, err := Add(context.Background(), app, AddOptions{Name: "Jordan's Key", SSHKey: "another-ssh-key"})
		require.Error(t, err)
		require.Zero(t, id)

		count, err := gorm.G[model.Credential](app.DB).Count(context.Background(), "id")
		require.NoError(t, err)
		require.Equal(t, int64(1), count)
	})

	// TODO: Validate ssh key format
	// @see https://askubuntu.com/questions/883943/check-if-private-key-is-malformed
}
