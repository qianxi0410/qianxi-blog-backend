package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/qianxi/blog-backend/service"
	"github.com/qianxi/blog-backend/util"
)

type OAuth2Controller struct{}

var oauth2Service service.OAuth2Service

func (oc OAuth2Controller) GetUserInfo(c *gin.Context) {
	var result string

	code := c.Param("code")
	result, err := oauth2Service.GetUserInfo(code)

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
		Data: result,
	})
}
