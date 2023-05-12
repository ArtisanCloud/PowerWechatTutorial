package official_account

import (
	"github.com/ArtisanCloud/PowerLibs/v3/fmt"
	"github.com/ArtisanCloud/PowerLibs/v3/http/helper"
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/contract"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/messages"
	models2 "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/models"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/officialAccount/server/handlers/models"
	"github.com/gin-gonic/gin"
	"power-wechat-tutorial/services"
)

// 回调配置
// https://work.weixin.qq.com/api/doc/90000/90135/90930
func CallbackVerify(c *gin.Context) {
	rs, err := services.OfficialAccountApp.Server.VerifyURL(c.Request)
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

	rs, err := services.OfficialAccountApp.Server.Notify(c.Request, func(event contract.EventInterface) interface{} {
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
		case models.CALLBACK_EVENT_SCAN:
			msg := models.EventScanCodePush{}
			err := event.ReadMessage(&msg)
			if err != nil {
				println(err.Error())
				return "error"
			}
		}

		//// 调用GPT3.5模型
		//req := v1.CreateChatCompletionRequest{
		//	Model: "gpt-3.5-turbo",
		//	Messages: []v1.Message{
		//		{
		//			Role:    "user",
		//			Content: "Hello!",
		//		},
		//	},
		//}
		//result, err := services.RobotChatApp.GPTClient.V1.Chat.CreateChatCompletion(&req)
		//if err != nil {
		//	fmt2.Printf("error: %v", err)
		//	return err.Error()
		//}

		//return messages.NewText(result.Object)
		replyData := messages.NewNews([]*object.HashMap{
			{
				"title":       "test",
				"description": "tttttt",
				"url":         "",
				"image":       "",
			},
		})

		return replyData
		//return messages.NewText("ok")

		//media_id := "JRzcFCs0neDADadmOep2YOszEXI0ZFesCRP75VgIX7UgLzy7Uqk2YeYcwyHtOmAe"
		//return messages.NewImage(media_id, &power.HashMap{})
		//return kernel.SUCCESS_EMPTY_RESPONSE

	})
	if err != nil {
		panic(err)
	}

	err = helper.HttpResponseSend(rs, c.Writer)

	if err != nil {
		panic(err)
	}

}