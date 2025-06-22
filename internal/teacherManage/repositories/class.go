package repositories

import (
	"face-signIn/internal/models"
	"face-signIn/pkg/globals"
)

// GetAllClasses 获取所有班级
func GetAllClasses() ([]models.Class, error) {
	var classes []models.Class
	if err := globals.DB.Find(&classes).Error; err != nil {
		return nil, err
	}
	return classes, nil
}
