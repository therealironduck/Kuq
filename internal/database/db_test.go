package database

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

const getTablesQuery = `
SELECT 
    name
FROM 
    sqlite_schema
WHERE 
    type ='table' AND 
    name NOT LIKE 'sqlite_%';
`

type tableResult struct {
	Name string
}

func TestDatabaseMigrations(t *testing.T) {
	t.Parallel()

	t.Run("credentials", func(t *testing.T) {
		t.Parallel()

		db, err := New(":memory:")
		require.NoError(t, err)
		require.NotNil(t, db)

		result, err := gorm.G[tableResult](db).Raw(getTablesQuery).Find(context.Background())
		require.NoError(t, err)
		found := false

		for _, table := range result {
			if table.Name == "credentials" {
				found = true
				break
			}
		}

		require.True(t, found, "Table 'credentials' not found. Existing tables: %v", result)
	})
}

func TestDatabaseConnection(t *testing.T) {
	t.Parallel()

	t.Run("returns an error if the database can't be created", func(t *testing.T) {
		t.Parallel()

		_, err := New("invalid/path.db")
		require.Error(t, err)
	})
}
