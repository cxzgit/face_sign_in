package repositories

import (
	"face-signIn/internal/models"
	"face-signIn/pkg/globals"
)

// 签到任务相关
func CreateSignInTask(task *models.SignInTask) error {
	return globals.DB.Create(task).Error
}

func GetSignInTaskByID(id uint) (*models.SignInTask, error) {
	var task models.SignInTask
	if err := globals.DB.First(&task, id).Error; err != nil {
		return nil, err
	}
	return &task, nil
}

func GetSignInTasksByCourseID(courseID uint) ([]models.SignInTask, error) {
	var tasks []models.SignInTask
	if err := globals.DB.Where("course_id = ?", courseID).Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

// 签到记录相关
func CreateSignInRecord(record *models.SignInRecord) error {
	return globals.DB.Create(record).Error
}

func GetSignInRecordsByTaskID(taskID uint) ([]models.SignInRecord, error) {
	var records []models.SignInRecord
	if err := globals.DB.Where("sign_in_task_id = ?", taskID).Find(&records).Error; err != nil {
		return nil, err
	}
	return records, nil
}

func GetSignInRecordByTaskIDAndStudentID(taskID, studentID uint) (*models.SignInRecord, error) {
	var record models.SignInRecord
	if err := globals.DB.Where("sign_in_task_id = ? AND student_id = ?", taskID, studentID).First(&record).Error; err != nil {
		return nil, err
	}
	return &record, nil
}
