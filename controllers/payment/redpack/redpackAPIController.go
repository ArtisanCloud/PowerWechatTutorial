package redpack

import (
	"github.com/ArtisanCloud/PowerWeChat/src/payment/redpack/request"
	"github.com/gin-gonic/gin"
  "net/http"
  "power-wechat-tutorial/services"
)

func APISendNormal(c *gin.Context) {

	options := &request.RequestSendRedPack{
		MchBillno:   "0010010404201411170000046545",
		MchID:       "888",
		Wxappid:     "wxcbda96de0b165486",
		SendName:    "send_name",
		ReOpenid:    "onqOjjmM1tad-3ROpncN-yUfa6uI",
		TotalAmount: 200,
		TotalNum:    1,
		Wishing:     "恭喜发财",
		ClientIp:    "127.0.0.1",
		ActName:     "新年红包",
		Remark:      "新年红包",
		SceneID:     "PRODUCT_2",
		NonceStr:    "50780e0cca98c8c8e814883e5caa672e",
		RiskInfo:    "posttime%3d123123412%26clientversion%3d234134%26mobile%3d122344545%26deviceid%3dIOS",
	}

  services.PaymentApp.SetExternalRequest(c.Request)
	payConf, err := services.PaymentApp.RedPack.SendNormal(options)
	if err != nil {
		panic(err)
	}

	c.XML(http.StatusOK, payConf)
}
