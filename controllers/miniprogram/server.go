package miniprogram

import (
	"github.com/ArtisanCloud/PowerLibs/v3/fmt"
	"github.com/ArtisanCloud/PowerLibs/v3/http/helper"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/contract"
	models2 "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/models"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/officialAccount/server/handlers/models"
	"github.com/gin-gonic/gin"
	"power-wechat-tutorial/services"
)

// 回调配置
// https://work.weixin.qq.com/api/doc/90000/90135/90930
func CallbackVerify(c *gin.Context) {
	rs, err := services.MiniProgramApp.Server.VerifyURL(c.Request)
	if err != nil {
		panic(err)
	}

	// 选择1
	//text, _ := ioutil.ReadAll(rs.Body)
	//c.String(http.StatusOK, string(text))

	// 选择2
	err = helper.HttpResponseSend(rs, c.Writer)

}

// 回调配置
// https://work.weixin.qq.com/api/doc/90000/90135/90930
func CallbackNotify(c *gin.Context) {
	//requestXML, _ := ioutil.ReadAll(c.Request.Body)
	//c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(requestXML))
	//println(string(requestXML))

	rs, err := services.MiniProgramApp.Server.Notify(c.Request, func(event contract.EventInterface) interface{} {
		fmt.Dump("event", event)
		//return  "handle callback"

		switch event.GetMsgType() {
		case models2.CALLBACK_MSG_TYPE_TEXT:
			msg := models.MessageText{}
			err := event.ReadMessage(&msg)
			if err != nil {
				println(err.Error())
				return "error"
			}
			fmt.Dump(msg)
		}

		return kernel.SUCCESS_EMPTY_RESPONSE

	})
	if err != nil {
		panic(err)
	}

	err = helper.HttpResponseSend(rs, c.Writer)

	if err != nil {
		panic(err)
	}

}
