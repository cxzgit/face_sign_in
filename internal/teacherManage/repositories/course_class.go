package repositories

import (
	"face-signIn/internal/models"
	"face-signIn/pkg/globals"
)

// 绑定课程和班级
func BindCourseClass(courseID, classID uint) error {
	cc := &models.CourseClass{CourseID: courseID, ClassID: classID}
	return globals.DB.Create(cc).Error
}

// 解绑课程和班级
func UnbindCourseClass(courseID, classID uint) error {
	return globals.DB.Where("course_id = ? AND class_id = ?", courseID, classID).Delete(&models.CourseClass{}).Error
}

// 查询课程绑定的所有班级
func GetClassesByCourseID(courseID uint) ([]models.Class, error) {
	var classes []models.Class
	var courseClasses []models.CourseClass
	if err := globals.DB.Where("course_id = ?", courseID).Find(&courseClasses).Error; err != nil {
		return nil, err
	}
	var classIDs []uint
	for _, cc := range courseClasses {
		classIDs = append(classIDs, cc.ClassID)
	}
	if len(classIDs) == 0 {
		return classes, nil
	}
	if err := globals.DB.Where("id IN ?", classIDs).Find(&classes).Error; err != nil {
		return nil, err
	}
	return classes, nil
}
