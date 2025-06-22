package repositories

import (
	"face-signIn/internal/models"
	"face-signIn/pkg/globals"
)

// 查询学生待签到任务
func GetPendingSignInTasksForStudent(classID, studentID uint, now int64) ([]models.SignInTask, error) {
	var tasks []models.SignInTask
	subQuery := globals.DB.Model(&models.SignInRecord{}).
		Select("sign_in_task_id").
		Where("student_id = ?", studentID)
	if err := globals.DB.Model(&models.SignInTask{}).
		Where("start_time <= ? AND end_time >= ?", now, now).
		Where("course_id IN (SELECT id FROM courses WHERE class_id = ?)", classID).
		Where("id NOT IN (?)", subQuery).
		Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

// 获取签到记录（可用于学生查询历史）
func GetSignInRecordsByStudentID(studentID uint) ([]models.SignInRecord, error) {
	var records []models.SignInRecord
	if err := globals.DB.Where("student_id = ?", studentID).Find(&records).Error; err != nil {
		return nil, err
	}
	return records, nil
}

// GetSignInRecordByTaskIDAndStudentID 根据任务ID和学生ID查询签到记录
func GetSignInRecordByTaskIDAndStudentID(taskID, studentID uint) (*models.SignInRecord, error) {
	var record models.SignInRecord
	if err := globals.DB.Where("sign_in_task_id = ? AND student_id = ?", taskID, studentID).First(&record).Error; err != nil {
		return nil, err
	}
	return &record, nil
}

// CreateSignInRecord 创建一条签到记录
func CreateSignInRecord(record *models.SignInRecord) error {
	return globals.DB.Create(record).Error
}
