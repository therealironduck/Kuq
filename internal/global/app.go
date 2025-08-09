package global

import (
	"github.com/therealironduck/kuq/internal/database"
	"gorm.io/gorm"
)

type App struct {
	WorkspacePath string
	DB            *gorm.DB
}

var pathResolver = database.ResolveWorkspacePath
var dbPathResolver = database.GetDatabasePath

func New(relativePath string) (*App, error) {
	workspacePath, err := pathResolver(relativePath)
	if err != nil {
		return nil, err
	}

	dbPath, err := dbPathResolver(relativePath)
	if err != nil {
		return nil, err
	}

	db, err := database.New(dbPath)
	if err != nil {
		return nil, err
	}

	return &App{
		WorkspacePath: workspacePath,
		DB:            db,
	}, nil
}
