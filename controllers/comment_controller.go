package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tgm-tmy/go-api/controllers/services"
	"github.com/tgm-tmy/go-api/models"
)

type CommentController struct {
	service services.CommentServicer
}

func NewCommentController(s services.CommentServicer) *CommentController {
	return &CommentController{service: s}
}

func (c *CommentController) PostCommentHandler(ctx *gin.Context) {
	var reqComment models.Comment
	if err := ctx.ShouldBindJSON(&reqComment); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "fail to decode json"})
		return
	}

	comment, err := c.service.PostCommentService(reqComment)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "fail internal exec"})
		return
	}
	ctx.JSON(http.StatusOK, comment)
}
