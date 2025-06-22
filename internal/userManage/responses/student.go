package responses

type StudentInfo struct {
	StudentID string `json:"student_id"`
	Name      string `json:"name"`
}

func NewStudentInfo(studentID, name string) *StudentInfo {
	return &StudentInfo{
		StudentID: studentID,
		Name:      name,
	}
}
