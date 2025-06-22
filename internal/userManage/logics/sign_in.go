package logics

import (
	"errors"
	"face-signIn/internal/models"
	"face-signIn/internal/userManage/repositories"
	"face-signIn/pkg/utils"
	"gorm.io/gorm"
	"time"
)

// GetPendingSignInTasksForStudent 获取学生待签到任务
func GetPendingSignInTasksForStudent(studentID, classID uint) ([]models.SignInTask, error) {
	now := time.Now().Unix()
	return repositories.GetPendingSignInTasksForStudent(classID, studentID, now)
}

// GetSignInRecordsByStudentID 获取学生签到历史
func GetSignInRecordsByStudentID(studentID uint) ([]models.SignInRecord, error) {
	return repositories.GetSignInRecordsByStudentID(studentID)
}

// StudentSignIn 学生签到逻辑
func StudentSignIn(studentID, signInTaskID uint, faceImage string) error {
	// 1. 检查是否在签到时间范围内
	task, err := repositories.GetSignInTaskByID(signInTaskID)
	if err != nil {
		return errors.New("签到任务不存在")
	}
	now := time.Now().Unix()
	if now < task.StartTime || now > task.EndTime {
		return errors.New("不在签到时间范围内")
	}

	// 2. 检查是否重复签到
	_, err = repositories.GetSignInRecordByTaskIDAndStudentID(signInTaskID, studentID)
	if err == nil {
		return errors.New("你已经签到过了，请勿重复签到")
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err // 其他数据库错误
	}

	// 3. 人脸识别验证
	ok, err := utils.FaceVerify(faceImage)
	if err != nil {
		return err // 返回人脸识别的具体错误
	}
	if !ok {
		return errors.New("人脸识别失败，请重试")
	}

	// 4. 创建签到记录
	record := &models.SignInRecord{
		SignInTaskID: signInTaskID,
		StudentID:    studentID,
		SignInTime:   time.Now().Unix(),
		FaceImage:    faceImage, // 出于隐私和存储考虑，生产环境可考虑不存或存脱敏URL
		Status:       1,         // 1-已签到
	}

	return repositories.CreateSignInRecord(record)
}
