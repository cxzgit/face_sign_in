package models

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	StudentID string `gorm:"unique;not null" json:"student_id"`
	Name      string `json:"name"`
	Password  string `json:"password"`
	ClassID   uint   `json:"class_id"` // 班级ID
}
