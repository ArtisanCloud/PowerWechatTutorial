package message

import (
  "encoding/xml"
  "github.com/ArtisanCloud/go-libs/fmt"
  "github.com/ArtisanCloud/power-wechat/src/kernel"
  "github.com/ArtisanCloud/power-wechat/src/kernel/contract"
  "github.com/ArtisanCloud/power-wechat/src/work/server/handlers"
  "github.com/ArtisanCloud/power-wechat/src/work/server/handlers/models"
  "github.com/gin-gonic/gin"
  "io/ioutil"
  "net/http"
  "power-wechat-tutorial/services"
)

func TestBuffer(c *gin.Context) {

	textXML := "<xml><ToUserName><![CDATA[ww454dfb9d6f6d432a]]></ToUserName><FromUserName><![CDATA[WangChaoYi]]></FromUserName><CreateTime>1634401052</CreateTime><MsgType><![CDATA[text]]></MsgType><Content><![CDATA[thioutrrr]]></Content><MsgId>7019699067840561924</MsgId><AgentID>1000008</AgentID></xml>"
	var md interface{}
  md2 := models.MessageText{}
	err := xml.Unmarshal([]byte(textXML), &md2)
  md = md2
	fmt.Dump(md, err)
}

// 回调配置
// https://work.weixin.qq.com/api/doc/90000/90135/90930
func CallbackVerify(c *gin.Context) {
	rs, err := services.WeComApp.Server.Serve(c.Request)
	if err != nil {
		panic(err)
	}

	text, _ := ioutil.ReadAll(rs.Body)
	c.String(http.StatusOK, string(text))

}

// 回调配置
// https://work.weixin.qq.com/api/doc/90000/90135/90930
func CallbackNotify(c *gin.Context) {
	closure := handlers.NewServerCallbackHandler()
	closure.Callback = func(header contract.EventInterface, content interface{}) interface{} {
		fmt.Dump("header", header, "content", content)
		//return  "handle callback"
		return kernel.SUCCESS_EMPTY_RESPONSE

	}
	services.WeComApp.Server.Push(closure, 0)

	rs, err := services.WeComApp.Server.Serve(c.Request)
	if err != nil {
		panic(err)
	}

	err = rs.Send(c.Writer)

	if err != nil {
		panic(err)
	}

}
