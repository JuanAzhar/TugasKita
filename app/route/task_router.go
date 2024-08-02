package route

import (
	"tugaskita/features/task/handler"
	"tugaskita/features/task/repository"
	userRepo "tugaskita/features/user/repository"
	"tugaskita/features/task/service"
	m "tugaskita/utils/jwt"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func TaskRouter(db *gorm.DB, e *echo.Group) {
	userRepository := userRepo.New(db)

	taskRepository := repository.NewTaskRepository(db, userRepository)
	taskUseCase := service.NewTaskService(taskRepository)
	taskController := handler.New(taskUseCase)

	user := e.Group("/user-task")
	user.GET("/:id", taskController.ReadSpecificTask, m.JWTMiddleware())
	user.GET("", taskController.ReadAllTask, m.JWTMiddleware())
	user.POST("", taskController.UploadTaskUser, m.JWTMiddleware())
	user.GET("/riwayat", taskController.ReadHistoryTaskUser, m.JWTMiddleware())

	admin := e.Group("/admin-task")
	admin.GET("/:id", taskController.ReadSpecificTask, m.JWTMiddleware())
	admin.PUT("/:id", taskController.UpdateTask, m.JWTMiddleware())
	admin.GET("", taskController.ReadAllTask, m.JWTMiddleware())
	admin.POST("", taskController.AddTask, m.JWTMiddleware())
	admin.DELETE("/:taskId", taskController.DeleteTask, m.JWTMiddleware())

	admin.GET("/user", taskController.FindAllUserTask, m.JWTMiddleware())
	admin.PUT("/user/:id", taskController.UpdateTaskStatus, m.JWTMiddleware())
	admin.GET("/user/:id", taskController.FindUserTaskById, m.JWTMiddleware())
}
