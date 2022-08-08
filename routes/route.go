package routes

import (
	"myexample/go-gin/factory"

	"github.com/gin-gonic/gin"
)

func New(presenter factory.Presenter) *gin.Engine {
	router := gin.Default()
	router.POST("/users", presenter.UserPresenter.Register)
	router.POST("/login", presenter.UserPresenter.Login)
	return router
}
