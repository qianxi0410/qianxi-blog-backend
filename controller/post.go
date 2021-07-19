package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/qianxi/blog-backend/service"
	"github.com/qianxi/blog-backend/util"
)

type PostController struct{}

var postService service.PostService

func (p PostController) GetPostById(c *gin.Context) {
	id := c.Param("id")
	result, err := postService.Get(id)
	if err != nil {
		c.JSON(200, util.Response{
			Code: util.ERROR,
			Msg:  err.Error(),
			Data: nil,
		})
		return
	}
	c.JSON(200, util.Response{
		Code: util.OK,
		Msg:  "success",
		Data: result,
	})
}

func (p PostController) GetPostByPageQuery(c *gin.Context) {
	page, size := c.Param("page"), c.Param("size")
	result, err := postService.GetPostByPageQuery(page, size)
	if err != nil {
		c.JSON(200, util.Response{
			Code: util.ERROR,
			Msg:  err.Error(),
			Data: nil,
		})
		return
	}
	c.JSON(200, util.Response{
		Code: util.OK,
		Msg:  "success",
		Data: result,
	})
}
