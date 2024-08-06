package route

import (
	"tugaskita/features/reward/handler"
	"tugaskita/features/reward/repository"
	"tugaskita/features/reward/service"
	m "tugaskita/utils/jwt"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RewardRouter(db *gorm.DB, e *echo.Group) {
	rewardRepository := repository.NewRewardRepository(db)
	rewardUseCase := service.NewRewardService(rewardRepository)
	rewardController := handler.New(rewardUseCase)

	reward := e.Group("/reward")
	reward.GET("", rewardController.ReadAllReward, m.JWTMiddleware())
	reward.GET("/:id", rewardController.ReadSpecificReward, m.JWTMiddleware())
	reward.POST("", rewardController.AddReward, m.JWTMiddleware())
	reward.PUT("/:id", rewardController.UpdateReward, m.JWTMiddleware())
	reward.DELETE("/:id", rewardController.DeleteReward, m.JWTMiddleware())
	
	reward.GET("/history", rewardController.FindAllRewardHistory,m.JWTMiddleware())
	reward.GET("/user/:id", rewardController.FindUserRewardById, m.JWTMiddleware())
	reward.GET("/user", rewardController.FindAllUploadReward, m.JWTMiddleware())
	reward.POST("/exchange", rewardController.UploadRewardRequest, m.JWTMiddleware())
}
