package miniprogram

import (
	"crypto/sha256"
	"fmt"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/miniProgram/base/request"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"power-wechat-tutorial/services"
)

func APISNSSession(c *gin.Context) {

	//code, exist := c.GetQuery("code")
	//if !exist {
	//	panic("parameter code expected")
	//}
	//
	//rs, err := services.MiniProgramApp.Auth.Session(c.Request.Context(), code)
	//services.MiniProgramApp.AccessToken.Refresh()

	//{"session_key":"7o91EfeHO4PEnoeVzJxvqw==","openid":"o4QEk5Kc_y8QTrENCpKoxYhS4jkg","unionid":"orLIIs_Tfr0DLoG2iMwSq7RuaYRg"}

	encrypt := "DRdLarcsm8c2jsdSOubaLWh/FuNgC38OszQly2n5OTchOMdTMJb6FqwfSKD3PV3B5UKfljlwpkdQxFAd3q8orqfBK/X9+lth5zXvSZ+OxICncxqbycpJFf+o695N2aCE66oC+GPebSpYld0aiUYtbPZQdwr1uqg1toCIDkZCKuuXkFFl5QeKgwg6VJVZk5w5IM2mDzAnUr8pIQllvlA/4A=="
	sessionKey := "7o91EfeHO4PEnoeVzJxvqw=="
	//openId:="o4QEk5Kc_y8QTrENCpKoxYhS4jkg"
	//unionId:="orLIIs_Tfr0DLoG2iMwSq7RuaYRg"
	iv := "JhojkHhkdgqHPkqD9uYZYQ=="
	rs, err := services.MiniProgramApp.Encryptor.DecryptData(encrypt, sessionKey, iv)
	//services.MiniProgramApp.Base.CheckEncryptedData(c.Request.Context())

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, rs)

}

func APICheckEncryptedData(c *gin.Context) {
	encryptedData := c.DefaultQuery("encryptedData", "sTWzm26PrbsXlSA8AoW+GpiyNLJP0H5p2UT4dXKwLSvXv8aU4wIiJcZUcM/IzNXnoFtERY3BDRbZh6bwd0ZGENVhucqDPXmchTqseryIZnJiKsiNMHCpAkCA2Yl00q4UpOZYtGMuTX5BTuo1yB3bOOuIfDu6neHV3D158CofGB9m7TxFQ8A/JcauWzhvmEAPygfFaqCgDTEmluLu7S8wMA==")
	hashByte := sha256.Sum256([]byte(encryptedData))
	hash := hashByte[:]
	rs, err := services.MiniProgramApp.Base.CheckEncryptedData(c.Request.Context(), fmt.Sprintf("%x", hash))

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, rs)

}

func APIGetPaidUnionID(c *gin.Context) {
	openid := c.DefaultQuery("openid", "")
	log.Printf("openid: %s\n", openid)
	rs, err := services.MiniProgramApp.Base.GetPaidUnionID(c.Request.Context(), &request.RequestGetPaidUnionID{
		OpenID: openid,
		// TransactionID: "",
		// MchID:         "",
		// OutTradeNo:    "",
	})

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, rs)

}
