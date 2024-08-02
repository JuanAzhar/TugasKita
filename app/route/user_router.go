package route

import (
	"tugaskita/features/user/handler"
	"tugaskita/features/user/repository"
	"tugaskita/features/user/service"
	m "tugaskita/utils/jwt"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func UserRouter(db *gorm.DB, e *echo.Group) {
	userRepository := repository.New(db)
	userUseCase := service.New(userRepository)
	userController := handler.New(userUseCase)

	e.POST("", userController.Register)
	e.POST("/login", userController.Login)
	e.GET("", userController.ReadAllUser,m.JWTMiddleware())
	e.GET("/profile", userController.ReadProfileUser, m.JWTMiddleware())
	e.GET("/:id", userController.ReadSpecificUser, m.JWTMiddleware())
	e.DELETE("/:id", userController.DeleteUser, m.JWTMiddleware())
}
