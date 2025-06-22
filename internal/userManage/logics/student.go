package logics

import (
	"errors"
	"face-signIn/internal/models"
	"face-signIn/internal/userManage/repositories"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func RegisterStudent(studentID, name, password string) error {
	if _, err := repositories.GetStudentByStudentID(studentID); err == nil {
		return errors.New("学号已注册")
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	student := &models.Student{
		StudentID: studentID,
		Name:      name,
		Password:  string(hash),
	}
	return repositories.CreateStudent(student)
}

func LoginStudent(studentID, password string) (*models.Student, error) {
	student, err := repositories.GetStudentByStudentID(studentID)
	if err != nil {
		return nil, errors.New("学号或密码错误")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(student.Password), []byte(password)); err != nil {
		return nil, errors.New("学号或密码错误")
	}
	return student, nil
}
