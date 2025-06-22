package models

import "gorm.io/gorm"

type SignInTask struct {
	gorm.Model
	ID          uint   `gorm:"primaryKey" json:"id"`
	CourseID    uint   `json:"course_id"`
	TeacherID   uint   `json:"teacher_id"`
	ClassID     uint   `json:"class_id"`
	StartTime   int64  `json:"start_time"`
	EndTime     int64  `json:"end_time"`
	Description string `json:"description"`
}
