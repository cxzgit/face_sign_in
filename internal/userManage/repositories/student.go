package repositories

import (
	"face-signIn/internal/models"
	"face-signIn/pkg/globals"
)

// GetStudentByID 根据主键ID获取学生
func GetStudentByID(id uint) (*models.Student, error) {
	var student models.Student
	if err := globals.DB.First(&student, id).Error; err != nil {
		return nil, err
	}
	return &student, nil
}

// GetStudentByStudentID 根据学号获取学生
func GetStudentByStudentID(studentID string) (*models.Student, error) {
	var student models.Student
	if err := globals.DB.Where("student_id = ?", studentID).First(&student).Error; err != nil {
		return nil, err
	}
	return &student, nil
}

func CreateStudent(student *models.Student) error {
	return globals.DB.Create(student).Error
}
