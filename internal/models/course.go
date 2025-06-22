package models

import "gorm.io/gorm"

type Course struct {
	gorm.Model
	Name        string `json:"name"`
	TeacherID   uint   `json:"teacher_id"`
	ClassName   string `json:"class_name"`
	Description string `json:"description"`
}
