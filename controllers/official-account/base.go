package official_account

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"power-wechat-tutorial/services"
)

func ClearQuota(ctx *gin.Context) {
	data, err := services.OfficialAccountApp.Base.ClearQuota(ctx.Request.Context())
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, data)
}

func GetCallbackIP(ctx *gin.Context) {
	//services.OfficialAccountApp.AccessToken.SetCacheKey("123")
	data, err := services.OfficialAccountApp.Base.GetCallbackIP(ctx.Request.Context())
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, data)
}
