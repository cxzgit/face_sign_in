package responses

type StudentInfo struct {
	ID        uint   `json:"id"`
	StudentID string `json:"student_id"`
	Name      string `json:"name"`
	ClassID   uint   `json:"class_id"`
}

func NewStudentInfo(id uint, studentID, name string, classID uint) *StudentInfo {
	return &StudentInfo{
		ID:        id,
		StudentID: studentID,
		Name:      name,
		ClassID:   classID,
	}
}
