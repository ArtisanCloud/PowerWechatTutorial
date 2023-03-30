package payment

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/payment/profitSharing/request"
	"github.com/gin-gonic/gin"
	"net/http"
	"power-wechat-tutorial/services"
)

func APIOrders(c *gin.Context) {

	transactionID := c.DefaultQuery("transactionID", "")
	outOrderNo := c.DefaultQuery("OutRefundNo", "")

	config := services.MiniProgramApp.GetConfig()
	RSAPublicKeyPath := config.GetString("rsa_public_key_path", "")
	services.PaymentApp.Base.BaseClient.RsaOAEP.RSAEncryptor.PublicKeyPath = RSAPublicKeyPath
	_, err := services.PaymentApp.Base.BaseClient.RsaOAEP.RSAEncryptor.LoadPublicKeyByPath()
	if err != nil {
		panic(err)
	}

	bufferEncryptedName, err := services.PaymentApp.Base.BaseClient.RsaOAEP.RSAEncryptor.Encrypt([]byte("hu89ohu89ohu89o"))
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

	//config := services.MiniProgramApp.GetConfig()
	//RSAPublicKeyPath := config.GetString("rsa_public_key_path", "")
	//services.PaymentApp.Base.BaseClient.RsaOAEP.RSAEncryptor.PublicKeyPath = RSAPublicKeyPath
	//_, err := services.PaymentApp.Base.BaseClient.RsaOAEP.RSAEncryptor.LoadPublicKeyByPath()
	//if err != nil {
	//	panic(err)
	//}

	bufferEncryptedName, err := services.PaymentApp.Base.BaseClient.RsaOAEP.RSAEncryptor.Encrypt([]byte("hu89ohu89ohu89o"))
	if err != nil {
		panic(err)
	}

	receiverType := "MERCHANT_ID"
	account := "86693852"
	name := string(bufferEncryptedName)
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
