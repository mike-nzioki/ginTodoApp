package models

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	Id       string `json:"id"`
	Task     string `json:"task"`
	Complete bool   `json:"complete"`
}
