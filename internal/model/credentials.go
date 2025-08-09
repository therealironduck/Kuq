package model

import "gorm.io/gorm"

type Credential struct {
	gorm.Model
	Name   string `gorm:"uniqueIndex"`
	SSHKey string
}
