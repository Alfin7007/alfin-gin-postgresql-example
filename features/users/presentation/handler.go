package presentation

import (
	"myexample/go-gin/features/users"
	"myexample/go-gin/features/users/presentation/request"
	"myexample/go-gin/helper"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userBussiness users.Bussiness
}

func NewUserHandler(userBussiness users.Bussiness) *UserHandler {
	return &UserHandler{
		userBussiness: userBussiness,
	}
}

func (uh UserHandler) Login(c *gin.Context) {
	userRequest := request.UserRequest{}
	errBind := c.Bind(&userRequest)
	if errBind != nil {
		c.JSON(helper.FailedBadRequest())
		c.Abort()
		return
	}
	userCore := request.ToCore(userRequest)
	userID, token, errLogin := uh.userBussiness.Login(userCore)
	if errLogin != nil {
		c.JSON(helper.FailedNotFound())
		c.Abort()
		return
	}

	c.IndentedJSON(helper.AuthOK(userID, token))
}

func (uh UserHandler) Register(c *gin.Context) {
	userRequest := request.UserRequest{}

	errBind := c.Bind(&userRequest)
	if errBind != nil {
		c.IndentedJSON(helper.FailedBadRequest())
		c.Abort()
		return
	}

	userCore := request.ToCore(userRequest)
	err := uh.userBussiness.Register(userCore)
	if err != nil {
		c.IndentedJSON(helper.FailedBadRequestWithMSG(err.Error()))
		c.Abort()
		return
	}

	c.IndentedJSON(helper.SuccessCreateNoData())
}
