package transfer

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/payment/transfer/request"
	"github.com/gin-gonic/gin"
	"net/http"
	"power-wechat-tutorial/services"
)

func APIToTransfer(c *gin.Context) {

	mchId := services.PaymentApp.GetConfig().GetString("mch_id", "")
	appId := services.PaymentApp.GetConfig().GetString("app_id", "")

	options := &request.RequestTransferToBalance{
		MchAppID:         appId,
		MchID:            mchId,
		DeviceInfo:       "",
		NonceStr:         "",
		PartnerTradeNo:   "0010010404201411170000046545",
		OpenID:           "o4QEk5Kc_y8QTrENCpKoxYhS4jkg",
		CheckName:        "NO_CHECK",
		ReUserName:       "",
		Amount:           30,
		Desc:             "活动奖励",
		SpBillCreateIP:   "",
		Scene:            "",
		BrandID:          0,
		FinderTemplateID: "",
	}

	payConf, err := services.PaymentApp.Transfer.ToBalance(options)
	if err != nil {
		panic(err)
	}

	c.XML(http.StatusOK, payConf)
}

func APIQueryBalanceOrder(c *gin.Context) {

	rs, err := services.PaymentApp.Transfer.QueryBalanceOrder("0010010404201411170000046545")
	if err != nil {
		panic(nil)
	}

	c.JSON(http.StatusOK, rs)
}
