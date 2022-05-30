package user

import (
  "github.com/ArtisanCloud/PowerLibs/v2/fmt"
  "github.com/ArtisanCloud/PowerWeChat/v2/src/kernel"
  "github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/contract"
  "github.com/ArtisanCloud/PowerWeChat/v2/src/work/server/handlers/models"
  "github.com/gin-gonic/gin"
  "io/ioutil"
  "net/http"
  "power-wechat-tutorial/services"
)

// 回调配置
// https://work.weixin.qq.com/api/doc/90000/90135/90930
func CallbackVerify(c *gin.Context) {
  rs, err := services.WeComContactApp.Server.Serve(c.Request)
  if err != nil {
    panic(err)
  }

  text, _ := ioutil.ReadAll(rs.Body)
  c.String(http.StatusOK, string(text))

}

// 回调配置
// https://work.weixin.qq.com/api/doc/90000/90135/90930
func CallbackNotify(c *gin.Context) {

  rs, err := services.WeComContactApp.Server.Notify(c.Request, func(event contract.EventInterface) interface{} {
    fmt.Dump("event", event)
    //return  "handle callback"

    if event.GetEvent() == models.CALLBACK_EVENT_CHANGE_CONTACT && event.GetChangeType() == models.CALLBACK_EVENT_CHANGE_TYPE_CREATE_PARTY {
      msg := models.EventPartyCreate{}
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

  err = rs.Send(c.Writer)

  if err != nil {
    panic(err)
  }

}
