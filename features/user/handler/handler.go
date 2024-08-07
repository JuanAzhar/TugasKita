package handler

import (
	"net/http"
	dto "tugaskita/features/user/dto"
	"tugaskita/features/user/entity"
	middleware "tugaskita/utils/jwt"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type UserController struct {
	userUsecase entity.UserUseCaseInterface
}

func New(userUC entity.UserUseCaseInterface) *UserController {
	return &UserController{
		userUsecase: userUC,
	}
}

func (handler *UserController) Register(e echo.Context) error {
	input := dto.UserRequest{}
	errBind := e.Bind(&input)
	if errBind != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "error bind data" + errBind.Error(),
		})
	}
	image, err := e.FormFile("image")
	if err != nil {
		if err == http.ErrMissingFile {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "No file uploaded",
			})
		}
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Error uploading file",
		})
	}

	data := entity.UserCore{
		Name:     input.Name,
		Image:    input.Image,
		Email:    input.Email,
		Password: input.Password,
	}

	row, errUser := handler.userUsecase.Register(data, image)
	if errUser != nil {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": "error create user",
			"error":   errUser.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]any{
		"message": "succes create user",
		"data":    row,
	})
}

func (handler *UserController) Login(e echo.Context) error {
	input := new(dto.UserRequest)
	errBind := e.Bind(&input)
	if errBind != nil {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": "error bind data",
		})
	}

	data := entity.UserCore{
		Email:    input.Email,
		Password: input.Password,
	}

	data, token, err := handler.userUsecase.Login(data.Email, data.Password)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": "error login",
			"error":   err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]any{
		"message": "login success",
		"email":   data.Email,
		"token":   token,
	})
}

func (handler *UserController) DeleteUser(e echo.Context) error {
	idParams := e.Param("id")
	err := handler.userUsecase.DeleteUser(idParams)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Error deleting user",
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "User deleted successfully",
	})
}

func (handler *UserController) ReadSpecificUser(e echo.Context) error {

	idParamstr := e.Param("id")

	idParams, err := uuid.Parse(idParamstr)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": "user not found",
		})
	}

	data, err := handler.userUsecase.ReadSpecificUser(idParams.String())
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": "error get specific user",
		})
	}

	response := dto.UserResponse{
		Id:         data.ID,
		Name:       data.Name,
		Email:      data.Email,
		Point:      data.Point,
		TotalPoint: data.TotalPoint,
	}

	return e.JSON(http.StatusOK, map[string]any{
		"message": "get user",
		"data":    response,
	})
}

func (handler *UserController) ReadProfileUser(e echo.Context) error {
	userId, _, err := middleware.ExtractTokenUserId(e)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": err.Error(),
		})
	}

	println("user Id : ", userId)

	idCheck, err := uuid.Parse(userId)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": "user not found",
		})
	}

	data, err := handler.userUsecase.ReadSpecificUser(idCheck.String())
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": "error get profile user",
		})
	}

	response := dto.UserResponse{
		Id:         data.ID,
		Name:       data.Name,
		Email:      data.Email,
		Point:      data.Point,
		TotalPoint: data.TotalPoint,
	}

	return e.JSON(http.StatusOK, map[string]any{
		"message": "get user profile",
		"data":    response,
	})
}

func (handler *UserController) ReadAllUser(e echo.Context) error {
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

	data, err := handler.userUsecase.ReadAllUser()
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": "error get all user",
		})
	}

	dataList := []dto.UserResponse{}
	for _, v := range data {
		result := dto.UserResponse{
			Id:         v.ID,
			Name:       v.Name,
			Email:      v.Email,
			Point:      v.Point,
			TotalPoint: v.TotalPoint,
		}
		dataList = append(dataList, result)
	}

	return e.JSON(http.StatusOK, map[string]any{
		"message": "get all user",
		"data":    dataList,
	})
}

func (handler *UserController) GetRankUser(e echo.Context) error {
	data, err := handler.userUsecase.GetRankUser()
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": "error get all user",
		})
	}

	dataList := []dto.UserRankResponse{}
	for _, v := range data {
		result := dto.UserRankResponse{
			Name:  v.Name,
			Point: v.Point,
		}
		dataList = append(dataList, result)
	}

	return e.JSON(http.StatusOK, map[string]any{
		"message": "get all user rank",
		"data":    dataList,
	})
}

func (handler *UserController) ChangePassword(e echo.Context) error {
	userId, _, err := middleware.ExtractTokenUserId(e)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]any{
			"message": err.Error(),
		})
	}

	data := new(dto.UserRequest)
	if errBind := e.Bind(data); errBind != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Error binding data",
		})
	}

	userData := entity.UserCore{
		Password: data.Password,
	}

	errUpdate := handler.userUsecase.ChangePassword(userId, userData)
	if errUpdate != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Error updating password",
			"error":   errUpdate.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "password updated",
	})

}
