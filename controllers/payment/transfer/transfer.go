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

	payConf, err := services.PaymentApp.Transfer.ToBalance(c.Request.Context(), options)
	if err != nil {
		panic(err)
	}

	c.XML(http.StatusOK, payConf)
}

func APIQueryBalanceOrder(c *gin.Context) {

	rs, err := services.PaymentApp.Transfer.QueryBalanceOrder(c.Request.Context(), "0010010404201411170000046545")
	if err != nil {
		panic(nil)
	}

	c.JSON(http.StatusOK, rs)
}

func APIToBankCard(c *gin.Context) {

	mchId := services.PaymentApp.GetConfig().GetString("mch_id", "")

	options := &request.RequestToBankCard{
		MchID:          mchId,
		BankCode:       "0010010404201411170000046545",
		Amount:         30,
		Desc:           "活动奖励",
		EncBankNO:      "123213",
		EncTrueName:    "321312",
		NonceStr:       "1231232131",
		PartnerTradeNO: "213313213212",
	}

	payConf, err := services.PaymentApp.Transfer.ToBankCard(c.Request.Context(), options)
	if err != nil {
		panic(err)
	}

	c.XML(http.StatusOK, payConf)
}
