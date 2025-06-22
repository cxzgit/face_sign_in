package requests

type TeacherRegisterRequest struct {
	TeacherID string `json:"teacher_id" binding:"required"`
	Name      string `json:"name" binding:"required"`
	Password  string `json:"password" binding:"required"`
}

type TeacherLoginRequest struct {
	TeacherID string `json:"teacher_id" binding:"required"`
	Password  string `json:"password" binding:"required"`
}
