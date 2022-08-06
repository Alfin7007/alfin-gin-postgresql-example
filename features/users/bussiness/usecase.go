package bussiness

import "myexample/go-gin/features/users"

type userUserCase struct {
	userData users.Data
}

func NewUserBussiness(userData users.Data) users.Bussiness {
	return &userUserCase{
		userData: userData,
	}
}

func (uc userUserCase) Login(userCore users.Core) (token string, err error) {
	return token, err
}

func (uc userUserCase) Register(userCore users.Core) (err error) {
	err = uc.userData.InsertData(userCore)
	return err
}
