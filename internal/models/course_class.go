package models

import "gorm.io/gorm"

type CourseClass struct {
	gorm.Model
	CourseID uint
	ClassID  uint
}
