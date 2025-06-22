package logics

import (
	"face-signIn/internal/models"
	"face-signIn/internal/teacherManage/repositories"
	"golang.org/x/crypto/bcrypt"
)

func GetStudentsByClass(classID uint) ([]models.Student, error) {
	return repositories.GetStudentsByClass(classID)
}

func CreateStudent(studentID, name, password string, classID uint) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	student := &models.Student{
		StudentID: studentID,
		Name:      name,
		Password:  string(hash),
		ClassID:   classID,
	}
	return repositories.CreateStudent(student)
}

func UpdateStudent(id uint, name, password string, classID uint) error {
	student, err := repositories.GetStudentByID(id)
	if err != nil {
		return err
	}
	if name != "" {
		student.Name = name
	}
	if password != "" {
		hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		student.Password = string(hash)
	}
	if classID != 0 {
		student.ClassID = classID
	}
	return repositories.UpdateStudent(student)
}

func DeleteStudentByID(id uint) error {
	return repositories.DeleteStudentByID(id)
}
