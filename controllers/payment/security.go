package payment

import (
	"github.com/ArtisanCloud/PowerLibs/v3/fmt"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/support"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/payment/security/response"
	"github.com/gin-gonic/gin"
	"net/http"
	"power-wechat-tutorial/services"
)

// Get RSA Public Key.
// https://pay.weixin.qq.com/wiki/doc/api/tools/mch_pay_yhk.php?chapter=25_7&index=4
func APIGetRSAPublicKey(c *gin.Context) {

	rs, err := services.PaymentApp.Security.GetRSAPublicKey(c.Request.Context())
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, rs)

}

// 获取平台证书
// https://pay.weixin.qq.com/wiki/doc/apiv3/apis/wechatpay5_1.shtml
func APIGetCertificates(c *gin.Context) {

	rs, err := services.PaymentApp.Security.GetCertificates(c.Request.Context())
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, rs)

}

func APIDecryptCertificate(c *gin.Context) {

	form := &response.ResponseGetCertificates{}

	if err := c.ShouldBindJSON(form); err != nil {
		panic(err)
	}

	config := services.PaymentApp.GetConfig()
	v3AESKey := config.GetString("mch_api_v3_key", "")

	plainTXT, err := support.DecryptAES256GCM(
		v3AESKey,
		form.Data[0].EncryptCertificate.AssociatedData,
		form.Data[0].EncryptCertificate.Nonce,
		form.Data[0].EncryptCertificate.Ciphertext,
	)
	if err != nil {
		panic(err)
	}
	fmt.Dump(plainTXT)
	c.JSON(http.StatusOK, plainTXT)

}
