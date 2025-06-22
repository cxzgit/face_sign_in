package repositories

import (
	"face-signIn/internal/models"
	"face-signIn/pkg/globals"
)

func CreateCourse(course *models.Course) error {
	return globals.DB.Create(course).Error
}

func GetCourseByID(id uint) (*models.Course, error) {
	var course models.Course
	if err := globals.DB.First(&course, id).Error; err != nil {
		return nil, err
	}
	return &course, nil
}

func GetCoursesByTeacherID(teacherID uint) ([]models.Course, error) {
	var courses []models.Course
	if err := globals.DB.Where("teacher_id = ?", teacherID).Find(&courses).Error; err != nil {
		return nil, err
	}
	return courses, nil
}

func UpdateCourse(course *models.Course) error {
	return globals.DB.Save(course).Error
}

func DeleteCourse(id uint) error {
	return globals.DB.Delete(&models.Course{}, id).Error
}
