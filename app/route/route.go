package route

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func New(e *echo.Echo, db *gorm.DB){
	user := e.Group("user")

	UserRouter(db, user)
}
