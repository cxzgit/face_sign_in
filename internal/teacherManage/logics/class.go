package logics

import (
	"face-signIn/internal/models"
	"face-signIn/internal/teacherManage/repositories"
)

// GetAllClasses 获取所有班级
func GetAllClasses() ([]models.Class, error) {
	return repositories.GetAllClasses()
}
