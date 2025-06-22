package models

import "gorm.io/gorm"

type SignInRecord struct {
	gorm.Model
	SignInTaskID uint   `json:"sign_in_task_id"`
	StudentID    uint   `json:"student_id"`
	SignInTime   int64  `json:"sign_in_time"`
	FaceImage    string `json:"face_image"`
	Status       int    `json:"status"` // 0未签到 1已签到
}
