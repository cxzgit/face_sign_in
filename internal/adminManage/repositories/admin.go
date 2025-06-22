package repositories

import (
	"face-signIn/internal/models"
	"face-signIn/pkg/globals"
)

func GetAdminByAdminID(adminID string) (*models.Admin, error) {
	var admin models.Admin
	if err := globals.DB.Where("admin_id = ?", adminID).First(&admin).Error; err != nil {
		return nil, err
	}
	return &admin, nil
}
