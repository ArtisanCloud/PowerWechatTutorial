package wecom

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"power-wechat-tutorial/services"
)

func Code2Session(c *gin.Context) {
	code := c.DefaultQuery("code", "")

	miniProgramApp, err := services.WeComApp.MiniProgram()
	if err != nil {
		panic(err)
	}
	res, err := miniProgramApp.Auth.Session(c.Request.Context(), code)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)
}
