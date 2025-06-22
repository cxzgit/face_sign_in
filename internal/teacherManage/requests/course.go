package requests

type CreateCourseRequest struct {
	Name        string `json:"name" binding:"required"`
	ClassName   string `json:"class_name" binding:"required"`
	Description string `json:"description"`
}

type UpdateCourseRequest struct {
	ID          uint   `json:"id" binding:"required"`
	Name        string `json:"name" binding:"required"`
	ClassName   string `json:"class_name" binding:"required"`
	Description string `json:"description"`
}
