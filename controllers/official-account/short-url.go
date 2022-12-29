package official_account

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"power-wechat-tutorial/services"
)

// ShortGenKey 短key托管
func ShortGenKey(ctx *gin.Context) {
	longData := ctx.DefaultQuery("longData", "longData test.....")
	data, err := services.OfficialAccountApp.URL.ShortGenKey(ctx.Request.Context(), longData, 30*24*3600)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
	}
	ctx.JSON(http.StatusOK, data)
}

// FetchShortGen 短key还原
func FetchShortGen(ctx *gin.Context) {
	shortKey := ctx.Query("shortKey")
	data, err := services.OfficialAccountApp.URL.FetchShorten(ctx.Request.Context(), shortKey)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
	}
	ctx.JSON(http.StatusOK, data)
}
