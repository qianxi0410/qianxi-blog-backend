package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/qianxi/blog-backend/model"
	"github.com/qianxi/blog-backend/service"
	"github.com/qianxi/blog-backend/util"
)

type CommentController struct{}

var commentService service.CommentService

func (cs CommentController) Save(c *gin.Context) {
	var comment model.Comment

	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusOK, util.Response{
			Code: util.ERROR,
			Msg:  err.Error(),
			Data: nil,
		})
		return
	}

	id, err := commentService.Save(comment)
	if err != nil {
		c.JSON(http.StatusOK, util.Response{
			Code: util.ERROR,
			Msg:  err.Error(),
			Data: nil,
		})
		return
	}
	c.JSON(http.StatusOK, util.Response{
		Code: util.OK,
		Msg:  "success",
		Data: id,
	})
}

func (cs CommentController) Delete(c *gin.Context) {
	id := c.Param("id")
	err := commentService.Delete(id)
	if err != nil {
		c.JSON(http.StatusOK, util.Response{
			Code: util.ERROR,
			Msg:  err.Error(),
			Data: nil,
		})
		return
	}
	c.JSON(http.StatusOK, util.Response{
		Code: util.OK,
		Msg:  "success",
		Data: nil,
	})
}
