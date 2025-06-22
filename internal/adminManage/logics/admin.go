package logics

import (
	"errors"
	"face-signIn/internal/adminManage/repositories"
	"face-signIn/internal/models"
	"golang.org/x/crypto/bcrypt"
)

func LoginAdmin(adminID, password string) (*models.Admin, error) {
	admin, err := repositories.GetAdminByAdminID(adminID)
	if err != nil {
		return nil, errors.New("管理员账号或密码错误")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(password)); err != nil {
		return nil, errors.New("管理员账号或密码错误")
	}
	return admin, nil
}
