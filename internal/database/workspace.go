package database

import (
	"path/filepath"
)

var abs = filepath.Abs

func ResolveWorkspacePath(relative string) (path string, err error) {
	path, err = abs(relative)
	return
}

func GetDatabasePath(workspacePath string) string {
	return filepath.Join(workspacePath, fileName)
}
