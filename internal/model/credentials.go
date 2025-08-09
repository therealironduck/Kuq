package model

import "gorm.io/gorm"

type Credential struct {
	gorm.Model
	Name   string
	SSHKey string
}
