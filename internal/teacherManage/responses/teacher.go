package responses

type TeacherInfo struct {
	TeacherID string `json:"teacher_id"`
	Name      string `json:"name"`
}

func NewTeacherInfo(teacherID, name string) *TeacherInfo {
	return &TeacherInfo{
		TeacherID: teacherID,
		Name:      name,
	}
}
