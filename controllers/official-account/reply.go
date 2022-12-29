package official_account

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"power-wechat-tutorial/services"
)

// AutoReply 获取当前设置的回复规则
func AutoReplyCurrent(ctx *gin.Context) {
	data, err := services.OfficialAccountApp.AutoReply.Current(ctx.Request.Context())
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
	}
	ctx.JSON(http.StatusOK, data)
}
