package factory

import (
	userBussiness "myexample/go-gin/features/users/bussiness"
	userData "myexample/go-gin/features/users/data"
	userPresenter "myexample/go-gin/features/users/presentation"

	"gorm.io/gorm"
)

type Presenter struct {
	UserPresenter *userPresenter.UserHandler
}

func InitFactory(db *gorm.DB) Presenter {
	varUserData := userData.NewUserRepo(db)
	varUserBussiness := userBussiness.NewUserBussiness(varUserData)
	varUserPresentation := userPresenter.NewUserHandler(varUserBussiness)

	return Presenter{
		UserPresenter: varUserPresentation,
	}
}
