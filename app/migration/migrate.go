package migration

import (
	users "tugaskita/features/user/model"
	task "tugaskita/features/task/model"

	"gorm.io/gorm"
)

func InitMigration(db *gorm.DB){
	db.AutoMigrate(&users.Users{})
	db.AutoMigrate(&task.Task{})
	db.AutoMigrate(&task.UserTaskUpload{})
}