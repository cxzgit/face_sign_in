package requests

type CreateStudentRequest struct {
	StudentID string `json:"student_id" binding:"required"`
	Name      string `json:"name" binding:"required"`
	Password  string `json:"password" binding:"required"`
	ClassID   uint   `json:"class_id" binding:"required"`
}

type UpdateStudentRequest struct {
	ID       uint   `json:"id" binding:"required"`
	Name     string `json:"name"`
	Password string `json:"password"`
	ClassID  uint   `json:"class_id"`
}
