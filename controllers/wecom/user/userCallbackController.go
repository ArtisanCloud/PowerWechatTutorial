package user

import (
  "github.com/ArtisanCloud/power-wechat/src/kernel/contract"
  "github.com/ArtisanCloud/power-wechat/src/work/server/handlers"
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
  closure := handlers.NewServerCallbackHandler()
  closure.Callback = func(header contract.EventInterface, content interface{}) interface{} {
    //fmt.Dump("payload", payload)
    return "handle callback"
  }
  services.WeComContactApp.Server.Push(closure, 0)

  rs, err := services.WeComContactApp.Server.Serve(c.Request)
  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, rs)

}
