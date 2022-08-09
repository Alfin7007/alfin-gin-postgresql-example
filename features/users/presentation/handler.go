package presentation

import (
	"myexample/go-gin/features/users"
	"myexample/go-gin/features/users/presentation/request"
	"myexample/go-gin/features/users/presentation/response"
	"myexample/go-gin/helper"
	"myexample/go-gin/middlewares"

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

	c.JSON(helper.AuthOK(userID, token))
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
		c.JSON(helper.FailedBadRequestWithMSG(err.Error()))
		c.Abort()
		return
	}

	c.JSON(helper.SuccessCreateNoData())
}

func (uh UserHandler) GetUser(c *gin.Context) {
	userID, _, errJWT := middlewares.JWTTokenCheck(c)
	if errJWT != nil {
		c.JSON(helper.FailedBadRequestWithMSG("invalid or exp jwt"))
		return
	}
	result, err := uh.userBussiness.GetData(userID)
	if err != nil {
		c.JSON(helper.FailedBadRequestWithMSG(err.Error()))
		return
	}

	c.JSON(helper.SuccessGetData(response.FromCore(result)))
}
