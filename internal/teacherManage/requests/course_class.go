package requests

type BindCourseClassRequest struct {
	CourseID uint   `json:"course_id" binding:"required"`
	ClassIDs []uint `json:"class_ids" binding:"required"` // 支持批量绑定
}

type UnbindCourseClassRequest struct {
	CourseID uint `json:"course_id" binding:"required"`
	ClassID  uint `json:"class_id" binding:"required"`
}
