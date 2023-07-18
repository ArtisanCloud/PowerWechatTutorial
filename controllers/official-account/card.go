package official_account

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"power-wechat-tutorial/services"
)

func APIUpdate(ctx *gin.Context) {
	data, err := services.OfficialAccountApp.Card.Update(ctx.Request.Context(),
		"ph_gmt7cUVrlRk8swPwx7aDyF-pg",
		"member_card",
		nil,
	)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, data)
}
