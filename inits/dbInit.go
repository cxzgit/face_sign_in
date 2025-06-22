package inits

import (
	"face-signIn/internal/models"
	"face-signIn/pkg/globals"
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// DBInit 初始化mysql
func DBInit() {
	if err := viper.UnmarshalKey("database", &globals.AppConfig.Database); err != nil {
		globals.Log.Panicf("无法解码为结构: %s", err)
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		globals.AppConfig.Database.User,
		globals.AppConfig.Database.Password,
		globals.AppConfig.Database.Host,
		globals.AppConfig.Database.Port,
		globals.AppConfig.Database.Name,
	)

	var err error
	globals.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true, // 取消外键约束
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "fs_", // 设置表前缀
		},
	})
	if err != nil {
		globals.Log.Panicf("Failed to connect to database: %v", err)
	}
}

// TableInit 初始化表
func TableInit() {
	err := globals.DB.AutoMigrate(
		&models.Student{},
		&models.Teacher{},
		&models.Admin{},
		&models.Class{},
		&models.Course{},
		&models.CourseClass{},
		&models.SignInTask{},
		&models.SignInRecord{},
	)
	if err != nil {
		globals.Log.Panicf("db.AutoMigrate err = %s", err)
		return
	}
}
