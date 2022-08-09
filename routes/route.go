package routes

import (
	"myexample/go-gin/factory"

	"github.com/gin-gonic/gin"
)

func New(presenter factory.Presenter) *gin.Engine {
	router := gin.Default()
	router.POST("/login", presenter.UserPresenter.Login)
	router.POST("/users", presenter.UserPresenter.Register)
	router.GET("/users", presenter.UserPresenter.GetUser)

	router.POST("/articles", presenter.ArticlePresenter.InsertArticle)
	router.GET("/articles", presenter.ArticlePresenter.GetAllArticle)
	router.GET("/articles/:id", presenter.ArticlePresenter.GetArticle)
	return router
}
