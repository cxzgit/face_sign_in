package requests

// 发起签到任务
// 支持多班级
// class_ids: 班级ID数组
//
type CreateSignInTaskRequest struct {
	CourseID    uint   `json:"course_id" binding:"required"`
	ClassIDs    []uint `json:"class_ids" binding:"required"` // 多班级
	StartTime   int64  `json:"start_time" binding:"required"`
	EndTime     int64  `json:"end_time" binding:"required"`
	Description string `json:"description"`
}

// 学生签到
type StudentSignInRequest struct {
	SignInTaskID uint   `json:"sign_in_task_id" binding:"required"`
	FaceImage    string `json:"face_image" binding:"required"`
}

// 老师手动代签到
// TeacherManualSignInRequest 用于老师为学生手动代签到
// 只需指定签到任务ID和学生ID
// 可选备注字段
// 不需要人脸图片
//
type TeacherManualSignInRequest struct {
	SignInTaskID uint   `json:"sign_in_task_id" binding:"required"`
	StudentID    uint   `json:"student_id" binding:"required"`
	Remark       string `json:"remark"`
}
