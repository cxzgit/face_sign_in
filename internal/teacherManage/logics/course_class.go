package logics

import "face-signIn/internal/teacherManage/repositories"

func BindCourseClass(courseID uint, classIDs []uint) error {
	for _, classID := range classIDs {
		if err := repositories.BindCourseClass(courseID, classID); err != nil {
			return err
		}
	}
	return nil
}

func UnbindCourseClass(courseID, classID uint) error {
	return repositories.UnbindCourseClass(courseID, classID)
}

func GetClassesByCourseID(courseID uint) ([]uint, error) {
	classes, err := repositories.GetClassesByCourseID(courseID)
	if err != nil {
		return nil, err
	}
	var classIDs []uint
	for _, c := range classes {
		classIDs = append(classIDs, c.ID)
	}
	return classIDs, nil
}
