package database

import (
	"github.com/therealironduck/kuq/internal/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const fileName = "kuq.db"

func New(path string) (db *gorm.DB, err error) {
	db, err = gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		return
	}

	db.AutoMigrate(&model.Credential{})

	return
}
