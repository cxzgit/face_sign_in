package models

import "gorm.io/gorm"

type Admin struct {
	gorm.Model
	AdminID  string `gorm:"unique;not null" json:"admin_id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}
