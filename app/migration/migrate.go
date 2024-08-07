package migration

import (
	users "tugaskita/features/user/model"
	task "tugaskita/features/task/model"
	reward "tugaskita/features/reward/model"

	"gorm.io/gorm"
)

func InitMigration(db *gorm.DB){
	db.AutoMigrate(&users.Users{})
	db.AutoMigrate(&task.Task{})
	db.AutoMigrate(&task.UserTaskUpload{})
	db.AutoMigrate(&task.UserTaskSubmission{})
	db.AutoMigrate(&reward.Reward{})
	db.AutoMigrate(&reward.UserRewardRequest{})
}