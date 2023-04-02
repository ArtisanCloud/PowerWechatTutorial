package payment

import (
	"encoding/base64"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/payment/profitSharing/request"
	"github.com/gin-gonic/gin"
	"net/http"
	"power-wechat-tutorial/services"
)

func APIOrders(c *gin.Context) {

	transactionID := c.DefaultQuery("transactionID", "")
	outOrderNo := c.DefaultQuery("OutRefundNo", "")

	config := services.MiniProgramApp.GetConfig()

	cipherdata, err := services.PaymentApp.Base.BaseClient.RsaOAEP.RSAEncryptor.Encrypt([]byte("hu89ohu89ohu89o"))
	bufferEncryptedName := base64.StdEncoding.EncodeToString(cipherdata)
	if err != nil {
		panic(err)
	}

	para := &request.RequestShare{
		AppID:         config.GetString("app_id", ""),
		TransactionID: transactionID,
		OutOrderNO:    outOrderNo,
		Receivers: []*request.Receiver{
			&request.Receiver{
				Type:        "MERCHANT_ID",
				Account:     "86693852",
				Name:        string(bufferEncryptedName),
				Amount:      300,
				Description: "分给商户A",
			},
		},
		UnfreezeUnSplit: true,
	}

	rs, err := services.PaymentApp.ProfitSharing.Share(c.Request.Context(), para)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, rs)

}

func APIAddReceiver(c *gin.Context) {

	cipherdata, err := services.PaymentApp.Base.BaseClient.RsaOAEP.RSAEncryptor.Encrypt([]byte("hu89ohu89ohu89o"))
	bufferEncryptedName := base64.StdEncoding.EncodeToString(cipherdata)

	receiverType := "MERCHANT_ID"
	account := "86693852"
	name := bufferEncryptedName
	relationType := "STAFF"
	customRelation := "分给商户A"

	rs, err := services.PaymentApp.ProfitSharing.AddReceiver(c.Request.Context(),
		receiverType, account, name,
		relationType, customRelation)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, rs)

}
