package bussiness

import (
	"errors"
	"fmt"
	"myexample/go-gin/features/users"
	"myexample/go-gin/middlewares"
)

type userUserCase struct {
	userData users.Data
}

func NewUserBussiness(userData users.Data) users.Bussiness {
	return &userUserCase{
		userData: userData,
	}
}

func (uc userUserCase) Login(userCore users.Core) (token string, err error) {
	result, errLogin := uc.userData.Login(userCore.Email)
	if errLogin != nil {
		return "", errLogin
	}
	if result.Password != userCore.Password {
		return "", errors.New("wrong password")
	}

	token, err = middlewares.GenerateToken(result.UserID)
	fmt.Println(err)
	return token, err
}

func (uc userUserCase) Register(userCore users.Core) (err error) {
	err = uc.userData.InsertData(userCore)
	return err
}
