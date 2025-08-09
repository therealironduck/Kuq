package global

import (
	"github.com/therealironduck/kuq/internal/database"
	"gorm.io/gorm"
)

type App struct {
	WorkspacePath string
	DB            *gorm.DB
}

type NewAppOptions struct {
	pathResolver   func(relative string) (path string, err error)
	dbPathResolver func(relative string) (path string, err error)
}

func NewApp(relativePath string, options *NewAppOptions) (*App, error) {
	if options == nil {
		options = &NewAppOptions{}
	}

	if options.pathResolver == nil {
		options.pathResolver = database.ResolveWorkspacePath
	}

	if options.dbPathResolver == nil {
		options.dbPathResolver = database.GetDatabasePath
	}

	workspacePath, err := options.pathResolver(relativePath)
	if err != nil {
		return nil, err
	}

	dbPath, err := options.dbPathResolver(relativePath)
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
