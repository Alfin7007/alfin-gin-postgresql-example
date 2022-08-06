package presentation

import (
	"myexample/go-gin/features/users"
	"myexample/go-gin/features/users/presentation/request"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userBussiness users.Bussiness
}

func NewUserHandler(userBussiness users.Bussiness) *userHandler {
	return &userHandler{
		userBussiness: userBussiness,
	}
}

func (uh userHandler) Login(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "login"})
}

func (uh userHandler) Register(c *gin.Context) {
	userRequest := request.UserRequest{}

	errBind := c.Bind(&userRequest)
	if errBind != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "register fail => " + errBind.Error()})
	}

	userCore := request.ToCore(userRequest)
	err := uh.userBussiness.Register(userCore)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "register fail => " + err.Error()})
	}
	c.IndentedJSON(http.StatusCreated, gin.H{"message": "success"})
}
