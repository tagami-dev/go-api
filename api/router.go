package api

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/tgm-tmy/go-api/api/middlewares"
	"github.com/tgm-tmy/go-api/controllers"
	"github.com/tgm-tmy/go-api/services"
)

func NewRouter(db *sql.DB) *gin.Engine {
	ser := services.NewMyAppService(db)
	aCon := controllers.NewArticleController(ser)
	cCon := controllers.NewCommentController(ser)

	r := gin.Default()

	r.POST("/article", aCon.PostArticleHandler)
	r.GET("/article/list", aCon.ArticleListHandler)
	r.GET("/article/:id", aCon.ArticleDetailHandler)
	r.POST("/article/nice", aCon.PostNiceHandler)
	r.POST("/comment", cCon.PostCommentHandler)

	r.Use(middlewares.LoggingMiddleware())

	return r
}
