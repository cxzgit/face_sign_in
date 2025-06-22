package responses

type SignInTaskInfo struct {
	ID          uint   `json:"id"`
	CourseID    uint   `json:"course_id"`
	TeacherID   uint   `json:"teacher_id"`
	ClassID     uint   `json:"class_id"`
	StartTime   int64  `json:"start_time"`
	EndTime     int64  `json:"end_time"`
	Description string `json:"description"`
}

func NewSignInTaskInfo(id, courseID, teacherID, classID uint, startTime, endTime int64, description string) *SignInTaskInfo {
	return &SignInTaskInfo{
		ID:          id,
		CourseID:    courseID,
		TeacherID:   teacherID,
		ClassID:     classID,
		StartTime:   startTime,
		EndTime:     endTime,
		Description: description,
	}
}

type SignInRecordInfo struct {
	ID           uint   `json:"id"`
	SignInTaskID uint   `json:"sign_in_task_id"`
	StudentID    uint   `json:"student_id"`
	SignInTime   int64  `json:"sign_in_time"`
	FaceImage    string `json:"face_image"`
	Status       int    `json:"status"`
}

func NewSignInRecordInfo(id, signInTaskID, studentID uint, signInTime int64, faceImage string, status int) *SignInRecordInfo {
	return &SignInRecordInfo{
		ID:           id,
		SignInTaskID: signInTaskID,
		StudentID:    studentID,
		SignInTime:   signInTime,
		FaceImage:    faceImage,
		Status:       status,
	}
}
