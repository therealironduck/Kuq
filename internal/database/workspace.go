package database

import (
	"path/filepath"
)

var abs = filepath.Abs

func ResolveWorkspacePath(relative string) (path string, err error) {
	path, err = abs(relative)
	return
}

func GetDatabasePath(relative string) (path string, err error) {
	workspacePath, err := ResolveWorkspacePath(relative)
	if err != nil {
		return
	}

	path = filepath.Join(workspacePath, fileName)

	return
}
