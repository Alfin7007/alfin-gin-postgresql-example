package bussiness

import (
	"errors"
	"myexample/go-gin/features/users"
	"myexample/go-gin/middlewares"

	"golang.org/x/crypto/bcrypt"
)

type userUserCase struct {
	userData users.Data
}

func NewUserBussiness(userData users.Data) users.Bussiness {
	return &userUserCase{
		userData: userData,
	}
}

func (uc userUserCase) Login(userCore users.Core) (id int, token string, err error) {
	result, errLogin := uc.userData.FindUser(userCore.Email)
	if errLogin != nil {
		return 0, "", errLogin
	}
	passCompare := bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(userCore.Password))
	if passCompare != nil {
		return 0, "", errors.New("wrong password")
	}

	token, _ = middlewares.GenerateToken(result.UserID)
	return result.UserID, token, err
}

func (uc userUserCase) Register(userCore users.Core) (err error) {
	_, userCheck := uc.userData.FindUser(userCore.Email)
	if userCheck == nil {
		return errors.New("email existing")
	}
	bytePass := []byte(userCore.Password)
	hashPass, _ := bcrypt.GenerateFromPassword(bytePass, bcrypt.DefaultCost)
	userCore.Password = string(hashPass)
	errInsert := uc.userData.InsertData(userCore)
	if errInsert != nil {
		return errors.New("failed insert data")
	}
	return errors.New("success")
}
