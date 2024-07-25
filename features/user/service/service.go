package service

import (
	"errors"
	"regexp"
	"tugaskita/features/user/entity"
	crypt "tugaskita/utils/bcrypt"
)

type userUseCase struct {
	userRepository entity.UserDataInterface
}

func New(userUCase entity.UserDataInterface) entity.UserUseCaseInterface {
	return &userUseCase{
		userRepository: userUCase,
	}
}

// DeleteUser implements entity.UserUseCaseInterface.
func (userUC *userUseCase) DeleteUser(id string) (err error) {
	if id == "" {
		return errors.New("insert user id")
	}

	_, errFind := userUC.userRepository.ReadSpecificUser(id)
	if errFind != nil {
		return errors.New("user not found")
	}

	errDelete := userUC.userRepository.DeleteUser(id)
	if errDelete != nil {
		return errors.New("can't delete user")
	}

	return nil

}

// Login implements entity.UserUseCaseInterface.
func (userUC *userUseCase) Login(email string, password string) (entity.UserCore, string, error) {
	if email == "" || password == "" {
		return entity.UserCore{}, "", errors.New("error, email or password can't be empty")
	}

	loginData, token, err := userUC.userRepository.Login(email, password)
	if err != nil {
		return entity.UserCore{}, "", err
	}

	if crypt.CheckPasswordHash(loginData.Password, password) {
		return loginData, token, nil
	}

	return entity.UserCore{}, "", errors.New("Login Failed")
}

// ReadSpecificUser implements entity.UserUseCaseInterface.
func (userUC *userUseCase) ReadSpecificUser(id string) (user entity.UserCore, err error) {
	if id == "" {
		return entity.UserCore{}, errors.New("event ID is required")
	}

	user, err = userUC.userRepository.ReadSpecificUser(id)
	if err != nil {
		return entity.UserCore{}, err
	}

	return user, nil
}

// Register implements entity.UserUseCaseInterface.
func (userUC *userUseCase) Register(data entity.UserCore) (row int, err error) {
	if data.Email == "" || data.Password == "" {
		return 0, errors.New("error, email or password can't be empty")
	}
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	match, _ := regexp.MatchString(emailRegex, data.Email)
	if !match {
		return 0, errors.New("error. email format not valid")
	}

	errRegister, err := userUC.userRepository.Register(data)
	if err != nil {
		return 0, err
	}

	return errRegister, nil
}

// ReadAllUser implements entity.UserUseCaseInterface.
func (userUC *userUseCase) ReadAllUser() ([]entity.UserCore, error) {
	users, err := userUC.userRepository.ReadAllUser()
	if err != nil {
		return nil, errors.New("error get data")
	}

	return users, nil
}