package logics

import (
	"errors"
	"face-signIn/internal/models"
	"face-signIn/internal/teacherManage/repositories"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func RegisterTeacher(teacherID, name, password string) error {
	if _, err := repositories.GetTeacherByTeacherID(teacherID); err == nil {
		return errors.New("教师号已注册")
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	teacher := &models.Teacher{
		TeacherID: teacherID,
		Name:      name,
		Password:  string(hash),
	}
	return repositories.CreateTeacher(teacher)
}

func LoginTeacher(teacherID, password string) (*models.Teacher, error) {
	teacher, err := repositories.GetTeacherByTeacherID(teacherID)
	if err != nil {
		return nil, errors.New("教师号或密码错误")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(teacher.Password), []byte(password)); err != nil {
		return nil, errors.New("教师号或密码错误")
	}
	return teacher, nil
}
