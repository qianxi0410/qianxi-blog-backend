package controller

import "github.com/gin-gonic/gin"

type PostController struct{}

func (p PostController) Pong(c *gin.Context) {
	c.JSON(200, map[int]int{
		1: 2,
	})
}
