package presentation

import (
	"myexample/go-gin/features/users"
	"myexample/go-gin/features/users/presentation/request"
	"net/http"

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
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "login fail => " + errBind.Error()})
		c.Abort()
		return
	}
	userCore := request.ToCore(userRequest)
	token, errLogin := uh.userBussiness.Login(userCore)
	if errLogin != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
		c.Abort()
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "success", "token": token})
}

func (uh UserHandler) Register(c *gin.Context) {
	userRequest := request.UserRequest{}

	errBind := c.Bind(&userRequest)
	if errBind != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "register fail => " + errBind.Error()})
		c.Abort()
		return
	}

	userCore := request.ToCore(userRequest)
	err := uh.userBussiness.Register(userCore)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "register fail => " + err.Error()})
	}
	c.IndentedJSON(http.StatusCreated, gin.H{"message": "success"})
}
