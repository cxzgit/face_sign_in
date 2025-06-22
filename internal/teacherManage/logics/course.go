package logics

import (
	"errors"
	"face-signIn/internal/models"
	"face-signIn/internal/teacherManage/repositories"
)

func CreateCourse(teacherID uint, name, className, description string) error {
	course := &models.Course{
		Name:        name,
		TeacherID:   teacherID,
		ClassName:   className,
		Description: description,
	}
	return repositories.CreateCourse(course)
}

func GetCoursesByTeacherID(teacherID uint) ([]models.Course, error) {
	return repositories.GetCoursesByTeacherID(teacherID)
}

func UpdateCourse(teacherID, id uint, name, className, description string) error {
	course, err := repositories.GetCourseByID(id)
	if err != nil {
		return err
	}
	if course.TeacherID != teacherID {
		return errors.New("无权限修改该课程")
	}
	course.Name = name
	course.ClassName = className
	course.Description = description
	return repositories.UpdateCourse(course)
}

func DeleteCourse(teacherID, id uint) error {
	course, err := repositories.GetCourseByID(id)
	if err != nil {
		return err
	}
	if course.TeacherID != teacherID {
		return errors.New("无权限删除该课程")
	}
	return repositories.DeleteCourse(id)
}
