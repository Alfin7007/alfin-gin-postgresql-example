package factory

import (
	userBussiness "myexample/go-gin/features/users/bussiness"
	userData "myexample/go-gin/features/users/data"
	userPresenter "myexample/go-gin/features/users/presentation"

	articleBussiness "myexample/go-gin/features/articles/bussiness"
	articleData "myexample/go-gin/features/articles/data"
	articlePresenter "myexample/go-gin/features/articles/presentation"

	"gorm.io/gorm"
)

type Presenter struct {
	UserPresenter    *userPresenter.UserHandler
	ArticlePresenter *articlePresenter.ArticleHandler
}

func InitFactory(db *gorm.DB) Presenter {
	varUserData := userData.NewUserRepo(db)
	varUserBussiness := userBussiness.NewUserBussiness(varUserData)
	varUserPresentation := userPresenter.NewUserHandler(varUserBussiness)

	varArticleData := articleData.NewArticleRepo(db)
	varArticleBussiness := articleBussiness.NewArticleBussiness(varArticleData)
	varArticlePresentation := articlePresenter.NewArticleHandler(varArticleBussiness)

	return Presenter{
		UserPresenter:    varUserPresentation,
		ArticlePresenter: varArticlePresentation,
	}
}
