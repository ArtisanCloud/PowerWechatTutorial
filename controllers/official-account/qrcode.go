package official_account

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"power-wechat-tutorial/services"
)

// 创建临时二维码
func GetTempQrCode(ctx *gin.Context) {
	data, err := services.OfficialAccountApp.QRCode.Temporary(ctx.Request.Context(), "val1", 30*24*3600)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, data)
}

// GetForeverQrCode 创建永久二维码
func GetForeverQrCode(ctx *gin.Context) {
	data, err := services.OfficialAccountApp.QRCode.Forever(ctx.Request.Context(), "val1")
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, data)
}

// 获取二维码网址
func GetQrCodeUrl(ctx *gin.Context) {
	url := services.OfficialAccountApp.QRCode.URL("from")
	ctx.String(http.StatusOK, url)
}
