package repositories

import (
	"face-signIn/internal/models"
	"face-signIn/pkg/globals"
)

// GetSignInTaskByID 根据ID获取签到任务详情
func GetSignInTaskByID(id uint) (*models.SignInTask, error) {
	var task models.SignInTask
	if err := globals.DB.First(&task, id).Error; err != nil {
		return nil, err
	}
	return &task, nil
}
