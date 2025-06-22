package repositories

import (
	"face-signIn/internal/models"
	"face-signIn/pkg/globals"
)

func GetStudentsByClass(classID uint) ([]models.Student, error) {
	var students []models.Student
	if err := globals.DB.Where("class_id = ?", classID).Find(&students).Error; err != nil {
		return nil, err
	}
	return students, nil
}

func CreateStudent(student *models.Student) error {
	return globals.DB.Create(student).Error
}

func UpdateStudent(student *models.Student) error {
	return globals.DB.Save(student).Error
}

func DeleteStudentByID(id uint) error {
	return globals.DB.Delete(&models.Student{}, id).Error
}

func GetStudentByID(id uint) (*models.Student, error) {
	var student models.Student
	if err := globals.DB.First(&student, id).Error; err != nil {
		return nil, err
	}
	return &student, nil
}
