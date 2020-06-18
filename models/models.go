package models

import (
	"github.com/jinzhu/gorm"
)

type Project struct {
	gorm.Model
	Name string
	Desc string 
	Tags string
	RepoLink string
	ComChannel string
}