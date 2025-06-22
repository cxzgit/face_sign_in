package repositories

import (
	"face-signIn/internal/models"
	"face-signIn/pkg/globals"
)

func GetAllTeachers() ([]models.Teacher, error) {
	var teachers []models.Teacher
	if err := globals.DB.Find(&teachers).Error; err != nil {
		return nil, err
	}
	return teachers, nil
}

func CreateTeacher(teacher *models.Teacher) error {
	return globals.DB.Create(teacher).Error
}

func UpdateTeacher(teacher *models.Teacher) error {
	return globals.DB.Save(teacher).Error
}

func DeleteTeacherByID(id uint) error {
	return globals.DB.Delete(&models.Teacher{}, id).Error
}

func GetTeacherByID(id uint) (*models.Teacher, error) {
	var teacher models.Teacher
	if err := globals.DB.First(&teacher, id).Error; err != nil {
		return nil, err
	}
	return &teacher, nil
}
