package wecom

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"power-wechat-tutorial/services"
)

func APITicketGet(c *gin.Context) {

	res, err := services.WeComApp.JSSDK.GetTicket(c.Request.Context())

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)
}
