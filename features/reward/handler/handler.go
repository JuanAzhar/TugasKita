package handler

import (
	"net/http"
	"tugaskita/features/reward/dto"
	"tugaskita/features/reward/entity"
	middleware "tugaskita/utils/jwt"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type RewardController struct {
	rewardUsecase entity.RewardUseCaseInterface
}

func New(rewardUC entity.RewardUseCaseInterface) *RewardController {
	return &RewardController{
		rewardUsecase: rewardUC,
	}
}

func (handler *RewardController) AddReward(e echo.Context) error {
	_, role, err := middleware.ExtractTokenUserId(e)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": err.Error(),
		})
	}

	if role != "admin" {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": "access denied",
		})
	}

	input := new(dto.RewardRequest)
	errBind := e.Bind(&input)
	if errBind != nil {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": "error bind data",
		})
	}

	data := entity.RewardCore{
		Name:  input.Name,
		Stock: input.Stock,
		Price: input.Price,
		Image: input.Image,
	}

	errTask := handler.rewardUsecase.CreateReward(data)
	if errTask != nil {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": "error create reward",
			"error":   errTask.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]any{
		"message": "succes create reward",
	})
}

func (handler *RewardController) ReadAllReward(e echo.Context) error {
	data, err := handler.rewardUsecase.FindAllReward()
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": "error get all reward",
		})
	}

	dataList := []dto.RewardResponse{}
	for _, v := range data {
		result := dto.RewardResponse{
			Id:    v.ID.String(),
			Name:  v.Name,
			Stock: v.Stock,
			Price: v.Price,
			Image: v.Image,
		}
		dataList = append(dataList, result)
	}

	return e.JSON(http.StatusOK, map[string]any{
		"message": "get all reward",
		"data":    dataList,
	})
}

func (handler *RewardController) ReadSpecificReward(e echo.Context) error {

	idParamstr := e.Param("id")

	idParams, err := uuid.Parse(idParamstr)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": "reward not found",
		})
	}

	data, err := handler.rewardUsecase.FindById(idParams.String())
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": "error get specific reward",
		})
	}

	response := dto.RewardResponse{
		Id:    data.ID.String(),
		Name:  data.Name,
		Stock: data.Stock,
		Price: data.Price,
		Image: data.Image,
	}

	return e.JSON(http.StatusOK, map[string]any{
		"message": "get reward",
		"data":    response,
	})
}

func (handler *RewardController) DeleteReward(e echo.Context) error {
	_, role, errRole := middleware.ExtractTokenUserId(e)
	if errRole != nil {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": errRole.Error(),
		})
	}

	if role != "admin" {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": "access denied",
		})
	}

	idParams := e.Param("id")
	err := handler.rewardUsecase.DeleteReward(idParams)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Error deleting reward",
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Reward deleted successfully",
	})
}

func (handler *RewardController) UpdateReward(e echo.Context) error {
	_, role, err := middleware.ExtractTokenUserId(e)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": err.Error(),
		})
	}

	if role != "admin" {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": "access denied",
		})
	}

	idParams := e.Param("id")

	data := new(dto.RewardRequest)
	if errBind := e.Bind(data); errBind != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Error binding data",
		})
	}

	rewardData := entity.RewardCore{
		Name:  data.Name,
		Stock: data.Stock,
		Price: data.Price,
		Image: data.Image,
	}

	errUpdate := handler.rewardUsecase.UpdateReward(idParams, rewardData)
	if errUpdate != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Error updating reward",
			"error":   errUpdate.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "reward updated successfully",
	})
}

func (handler *RewardController) FindAllRewardHistory(e echo.Context) error {
	userId, _, err := middleware.ExtractTokenUserId(e)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": err.Error(),
		})
	}

	data, err := handler.rewardUsecase.FindAllRewardHistory(userId)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": "error get history reward",
		})
	}

	dataList := []dto.RewardRequestResponse{}
	for _, v := range data {
		result := dto.RewardRequestResponse{
			Id:       v.Id.String(),
			RewardId: v.RewardId,
			UserId:   v.UserId,
			Status:   v.Status,
		}
		dataList = append(dataList, result)
	}

	return e.JSON(http.StatusOK, map[string]any{
		"message": "get all reward history",
		"data":    dataList,
	})
}

func (handler *RewardController) FindAllUploadReward(e echo.Context) error {
	_, role, err := middleware.ExtractTokenUserId(e)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": err.Error(),
		})
	}

	if role != "admin" {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": "access denied",
		})
	}

	data, err := handler.rewardUsecase.FindAllUploadReward()
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": "error get all user reward",
		})
	}

	dataList := []entity.UserRewardRequestCore{}
	for _, v := range data {
		result := entity.UserRewardRequestCore{
			Id:       v.Id,
			RewardId: v.RewardId,
			UserId:   v.UserId,
			Status:   v.Status,
		}
		dataList = append(dataList, result)
	}

	return e.JSON(http.StatusOK, map[string]any{
		"message": "get all user reward",
		"data":    dataList,
	})
}

func (handler *RewardController) UploadRewardRequest(e echo.Context) error {
	input := new(dto.RewardReqRequest)
	errBind := e.Bind(&input)
	if errBind != nil {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": "error bind data",
		})
	}

	userId, _, errRole := middleware.ExtractTokenUserId(e)
	if errRole != nil {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": errRole.Error(),
		})
	}

	dataInput := entity.UserRewardRequestCore{
		RewardId: input.RewardId,
		UserId:   userId,
	}

	err := handler.rewardUsecase.UploadRewardRequest(dataInput)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": "error upload request reward",
			"error":   err.Error(),
		})
	}

	dataRespon := dto.RewardRequestResponse{
		RewardId: input.RewardId,
		UserId:   userId,
	}

	return e.JSON(http.StatusOK, map[string]any{
		"message": "succes upload request reward",
		"data":    dataRespon,
	})
}

func (handler *RewardController) FindUserRewardById(e echo.Context) error {
	_, role, err := middleware.ExtractTokenUserId(e)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": err.Error(),
		})
	}

	if role != "admin" {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": "access denied",
		})
	}

	idParams := e.Param("id")

	data, err := handler.rewardUsecase.FindUserRewardById(idParams)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": "error get specific reward",
		})
	}

	response := dto.RewardRequestResponse{
		Id:       data.Id.String(),
		UserId:   data.UserId,
		RewardId: data.RewardId,
		Status:   data.Status,
	}

	return e.JSON(http.StatusOK, map[string]any{
		"message": "get reward",
		"data":    response,
	})
}
