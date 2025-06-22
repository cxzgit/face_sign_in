package responses

import "face-signIn/internal/models"

// SignInRecordInfo 定义了返回给前端的单条签到记录的结构
type SignInRecordInfo struct {
	ID           uint   `json:"id"`
	SignInTaskID uint   `json:"sign_in_task_id"`
	StudentID    uint   `json:"student_id"`
	SignInTime   int64  `json:"sign_in_time"`
	Status       int    `json:"status"`
	TaskDesc     string `json:"task_desc,omitempty"` // 关联的签到任务描述
}

// NewSignInRecordInfo 创建一个新的签到记录响应对象
func NewSignInRecordInfo(record models.SignInRecord, taskDesc string) *SignInRecordInfo {
	return &SignInRecordInfo{
		ID:           record.ID,
		SignInTaskID: record.SignInTaskID,
		StudentID:    record.StudentID,
		SignInTime:   record.SignInTime,
		Status:       record.Status,
		TaskDesc:     taskDesc,
	}
}
