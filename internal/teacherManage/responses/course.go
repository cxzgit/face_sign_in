package responses

type CourseInfo struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	TeacherID   uint   `json:"teacher_id"`
	ClassName   string `json:"class_name"`
	Description string `json:"description"`
}

func NewCourseInfo(id uint, name string, teacherID uint, className, description string) *CourseInfo {
	return &CourseInfo{
		ID:          id,
		Name:        name,
		TeacherID:   teacherID,
		ClassName:   className,
		Description: description,
	}
}
