package models

import "gorm.io/gorm"

type Class struct {
	gorm.Model
	Name      string `json:"name" gorm:"unique;not null"` // 班级名
	TeacherID uint   `json:"teacher_id"`                  // 班主任/负责教师ID
}
