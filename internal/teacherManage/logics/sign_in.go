package logics

import (
	"errors"
	"face-signIn/internal/models"
	"face-signIn/internal/teacherManage/repositories"
	"face-signIn/pkg/utils"
	"time"
)

// 发起签到任务，支持多班级
func CreateSignInTask(teacherID, courseID uint, classIDs []uint, startTime, endTime int64, description string) error {
	// 校验课程-班级绑定关系
	boundClassIDs, err := GetClassesByCourseID(courseID)
	if err != nil {
		return err
	}
	for _, classID := range classIDs {
		found := false
		for _, boundID := range boundClassIDs {
			if classID == boundID {
				found = true
				break
			}
		}
		if !found {
			return errors.New("班级未绑定到该课程")
		}
	}
	// 为每个班级分别创建签到任务
	for _, classID := range classIDs {
		task := &models.SignInTask{
			CourseID:    courseID,
			TeacherID:   teacherID,
			ClassID:     classID,
			StartTime:   startTime,
			EndTime:     endTime,
			Description: description,
		}
		if err := repositories.CreateSignInTask(task); err != nil {
			return err
		}
	}
	return nil
}

func GetSignInTasksByCourseID(courseID uint) ([]models.SignInTask, error) {
	return repositories.GetSignInTasksByCourseID(courseID)
}

func GetSignInTaskByID(id uint) (*models.SignInTask, error) {
	return repositories.GetSignInTaskByID(id)
}

// 学生签到
func StudentSignIn(studentID, signInTaskID uint, faceImage string) error {
	// 校验是否已签到
	if record, _ := repositories.GetSignInRecordByTaskIDAndStudentID(signInTaskID, studentID); record != nil {
		return errors.New("已签到，无需重复签到")
	}
	// 校验人脸
	ok, err := utils.FaceVerify(faceImage)
	if err != nil || !ok {
		return errors.New("人脸识别失败")
	}
	// 记录签到
	record := &models.SignInRecord{
		SignInTaskID: signInTaskID,
		StudentID:    studentID,
		SignInTime:   time.Now().Unix(),
		FaceImage:    faceImage,
		Status:       1,
	}
	return repositories.CreateSignInRecord(record)
}

func GetSignInRecordsByTaskID(taskID uint) ([]models.SignInRecord, error) {
	return repositories.GetSignInRecordsByTaskID(taskID)
}

// TeacherManualSignIn 老师手动为学生代签到
func TeacherManualSignIn(teacherID, signInTaskID, studentID uint, remark string) error {
	// 校验是否已签到
	if record, _ := repositories.GetSignInRecordByTaskIDAndStudentID(signInTaskID, studentID); record != nil {
		return errors.New("该学生已签到，无需重复代签")
	}
	// 校验该签到任务是否属于该老师
	task, err := repositories.GetSignInTaskByID(signInTaskID)
	if err != nil {
		return errors.New("签到任务不存在")
	}
	if task.TeacherID != teacherID {
		return errors.New("无权限为该签到任务代签")
	}
	// 记录代签到
	record := &models.SignInRecord{
		SignInTaskID: signInTaskID,
		StudentID:    studentID,
		SignInTime:   time.Now().Unix(),
		FaceImage:    "", // 代签无图片
		Status:       1,
	}
	return repositories.CreateSignInRecord(record)
}
