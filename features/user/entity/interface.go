package entity


type UserDataInterface interface{
	Register(data UserCore) (row int, err error)
	Login(email, password string) (UserCore, string, error)
	ReadAllUser()([]UserCore, error)
	ReadSpecificUser(id string) (user UserCore, err error)
	DeleteUser(id string) (err error)
	UpdatePoint(id string, data UserCore) error
}

type UserUseCaseInterface interface{
	Register(data UserCore) (row int, err error)
	Login(email, password string) (UserCore, string, error)
	ReadAllUser()([]UserCore, error)
	ReadSpecificUser(id string) (user UserCore, err error)
	DeleteUser(id string) (err error)
}