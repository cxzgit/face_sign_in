package repositories

import (
	"face-signIn/internal/models"
	"face-signIn/pkg/globals"
)

func GetTeacherByTeacherID(teacherID string) (*models.Teacher, error) {
	var teacher models.Teacher
	if err := globals.DB.Where("teacher_id = ?", teacherID).First(&teacher).Error; err != nil {
		return nil, err
	}
	return &teacher, nil
}

func CreateTeacher(teacher *models.Teacher) error {
	return globals.DB.Create(teacher).Error
}
