package redpack

import (
	"github.com/ArtisanCloud/PowerLibs/v3/fmt"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/payment/redpack/request"
	"github.com/gin-gonic/gin"
	"net/http"
	"power-wechat-tutorial/services"
)

func APISendNormal(c *gin.Context) {

	mchId := services.PaymentApp.GetConfig().GetString("mch_id", "")
	appId := services.PaymentApp.GetConfig().GetString("app_id", "")

	options := &request.RequestSendRedPack{
		MchBillNO:   "0010010404201411170000046545",
		MchID:       mchId,
		WXAppID:     appId,
		SendName:    "ArtisanCloud",
		ReOpenID:    "oAuaP0TRUMwP169nQfg7XCEAw3HQ",
		TotalAmount: 100,
		TotalNum:    1,
		Wishing:     "恭喜发财",
		ClientIP:    "127.0.0.1",
		ActName:     "新年红包",
		Remark:      "新年红包",
		SceneID:     "PRODUCT_2",
		RiskInfo:    "posttime%3d123123412%26clientversion%3d234134%26mobile%3d122344545%26deviceid%3dIOS",
	}

	payConf, err := services.PaymentApp.RedPack.SendNormal(c.Request.Context(), options)
	if err != nil {
		fmt.Dump(err.Error())
		panic(err)
	}

	c.XML(http.StatusOK, payConf)
}

func APIQueryRedPack(c *gin.Context) {
	rs, err := services.PaymentApp.RedPack.Info(c.Request.Context(), "0010010404201411170000046545")
	if err != nil {
		panic(nil)
	}

	c.JSON(http.StatusOK, rs)
}
