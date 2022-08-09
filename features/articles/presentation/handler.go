package presentation

import (
	"myexample/go-gin/features/articles"
	"myexample/go-gin/features/articles/presentation/request"
	"myexample/go-gin/features/articles/presentation/response"
	"myexample/go-gin/helper"
	"myexample/go-gin/middlewares"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ArticleHandler struct {
	artBussiness articles.Bussiness
}

func NewArticleHandler(artBussiness articles.Bussiness) *ArticleHandler {
	return &ArticleHandler{
		artBussiness: artBussiness,
	}
}

func (ah ArticleHandler) InsertArticle(c *gin.Context) {
	userID, _, errJWT := middlewares.JWTTokenCheck(c)
	if errJWT != nil {
		c.JSON(helper.FailedBadRequestWithMSG("invalid pr exp JWT"))
		return
	}

	artRequest := request.Article{}
	errBind := c.Bind(&artRequest)
	if errBind != nil {
		c.JSON(helper.FailedBadRequest())
		return
	}
	artRequest.UserID = userID

	err := ah.artBussiness.InsertArticle(request.ToCore(artRequest))
	if err != nil {
		c.JSON(helper.FailedBadRequest())
		return
	}
	c.JSON(helper.SuccessCreateNoData())
}

func (ah ArticleHandler) GetArticle(c *gin.Context) {
	articleID, errParam := strconv.Atoi(c.Param("id"))
	if errParam != nil {
		c.JSON(helper.FailedBadRequestWithMSG("invalid param"))
		return
	}
	result, err := ah.artBussiness.GetArticle(articleID)
	if err != nil {
		c.JSON(helper.FailedBadRequest())
	}

	c.JSON(helper.SuccessGetData(response.FromSingleCore(result)))
}

func (ah ArticleHandler) GetAllArticle(c *gin.Context) {

	result, err := ah.artBussiness.GetAll()
	if err != nil {
		c.JSON(helper.FailedBadRequest())
	}

	c.JSON(helper.SuccessGetData(response.FromCoreList(result)))
}
