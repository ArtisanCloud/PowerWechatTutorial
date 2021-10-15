package message

import (
  "github.com/ArtisanCloud/go-libs/fmt"
  "github.com/gin-gonic/gin"
  "io/ioutil"
  "net/http"
  "power-wechat-tutorial/services"
)

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
  requestBody, _ := ioutil.ReadAll(c.Request.Body)
  fmt.Dump(requestBody)
}
