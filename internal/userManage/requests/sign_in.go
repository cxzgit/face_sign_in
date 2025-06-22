package requests

// StudentSignInRequest 学生签到请求
type StudentSignInRequest struct {
	SignInTaskID uint   `json:"sign_in_task_id" binding:"required"`
	FaceImage    string `json:"face_image" binding:"required"` // Base64编码的人脸图片
}
