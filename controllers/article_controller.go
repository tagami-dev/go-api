package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tgm-tmy/go-api/apperrors"
	"github.com/tgm-tmy/go-api/controllers/services"
	"github.com/tgm-tmy/go-api/models"
)

type ArticleController struct {
	service services.ArticleServicer
}

func NewArticleController(s services.ArticleServicer) *ArticleController {
	return &ArticleController{service: s}
}

func (c *ArticleController) HelloHandler(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Hello, world!\n")
}

func (c *ArticleController) PostArticleHandler(ctx *gin.Context) {
	var reqArticle models.Article
	if err := ctx.ShouldBindJSON(&reqArticle); err != nil {
		err = apperrors.ReqBodyDecodeFailed.Wrap(err, "bad request body")
		apperrors.ErrorHandler(ctx, err)
		return
	}

	article, err := c.service.PostArticleService(reqArticle)
	if err != nil {
		apperrors.ErrorHandler(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, article)
}

func (c *ArticleController) ArticleListHandler(ctx *gin.Context) {
	page, err := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	if err != nil {
		err = apperrors.BadParam.Wrap(err, "queryparam must be number")
		apperrors.ErrorHandler(ctx, err)
		return
	}

	articleList, err := c.service.GetArticleListService(page)
	if err != nil {
		apperrors.ErrorHandler(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, articleList)
}

func (c *ArticleController) ArticleDetailHandler(ctx *gin.Context) {
	articleID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		err = apperrors.BadParam.Wrap(err, "pathparam must be number")
		apperrors.ErrorHandler(ctx, err)
		return
	}

	article, err := c.service.GetArticleService(articleID)
	if err != nil {
		apperrors.ErrorHandler(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, article)
}

func (c *ArticleController) PostNiceHandler(ctx *gin.Context) {
	var reqArticle models.Article
	if err := ctx.ShouldBindJSON(&reqArticle); err != nil {
		apperrors.ErrorHandler(ctx, err)
		return
	}

	article, err := c.service.PostNiceService(reqArticle)
	if err != nil {
		apperrors.ErrorHandler(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, article)
}
