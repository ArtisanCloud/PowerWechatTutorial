package open_platform

import (
	"bytes"
	"encoding/xml"
	"github.com/ArtisanCloud/PowerLibs/v2/fmt"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel"
	openplatform "github.com/ArtisanCloud/PowerWeChat/v2/src/openPlatform/server/callbacks"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"power-wechat-tutorial/services"
)

func GetPreAuthorizationUrl(ctx *gin.Context) {
	//services.OpenPlatformApp.GetMobilePreAuthorizationURL()
	account, _ := services.OpenPlatformApp.OfficialAccount("", "", nil)
	a := account.Account
	a.Create()
}

// 验证票据 https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/Before_Develop/component_verify_ticket.html
func APIOpenPlatformCallback(context *gin.Context) {
	requestXML, _ := ioutil.ReadAll(context.Request.Body)
	context.Request.Body = ioutil.NopCloser(bytes.NewBuffer(requestXML))
	//println(string(requestXML))

	var err error

	rs, err := services.OpenPlatformApp.Server.Notify(context.Request, func(event *openplatform.Callback, decrypted []byte) (result interface{}) {

		result = kernel.SUCCESS_EMPTY_RESPONSE
		fmt.Dump("event: ", event, decrypted)
		msg := &openplatform.EventVerifyTicket{}
		err = xml.Unmarshal(decrypted, msg)
		if err != nil {
			log.Println(err)
			return err
		}
		fmt.Dump("msg: ", msg)

		return result
	})

	if err != nil {
		panic(err)
	}

	err = rs.Send(context.Writer)
	if err != nil {
		panic(err)
	}

}
