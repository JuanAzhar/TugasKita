package migration

import (
	users "tugaskita/features/user/model"

	"gorm.io/gorm"
)

func InitMigration(db *gorm.DB){
	db.AutoMigrate(&users.Users{})
}