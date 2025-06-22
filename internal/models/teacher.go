package models

import "gorm.io/gorm"

type Teacher struct {
	gorm.Model
	TeacherID string `gorm:"unique;not null" json:"teacher_id"`
	Name      string `json:"name"`
	Password  string `json:"password"`
	// ClassID   uint   `json:"class_id"` // 可选：如需支持班主任/班级归属
}
