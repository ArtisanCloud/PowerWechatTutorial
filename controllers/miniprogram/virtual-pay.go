package miniprogram

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/miniProgram/virtualPayment/request"
	"github.com/gin-gonic/gin"
	"power-wechat-tutorial/services"
)

func APIStartUploadGoods(ctx *gin.Context) {

	params := &request.UploadProductsRequest{
		Env: 0,
		UploadItem: []*request.GoodItem{
			{
				Id:      "18",
				Name:    "1000K币",
				Price:   1000,
				Remake:  "1000K币",
				ItemUrl: "https://qiniu.rongjuwh.cn/applet_kb_20230801101836.png",
			},
		},
	}

	rs, err := services.MiniProgramApp.VirtualPayment.StartUploadGoods(ctx.Request.Context(), params)

	if err != nil {
		panic(err)
	}

	ctx.JSON(200, rs)
}
