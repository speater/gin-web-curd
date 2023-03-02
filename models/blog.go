package models

import "gorm.io/gorm"

type Blog struct {
	gorm.Model
	Date    string
	Content string
	Author  string
}
