package officialAccount

import (
  "github.com/ArtisanCloud/power-wechat/src/kernel/power"
  "github.com/gin-gonic/gin"
  "power-wechat-tutorial/services"
)

func APIMediaUpload(c *gin.Context)  {
  var outRes interface{}
  res, err := services.OfficialAccountApp.Media.UploadTempFile("./resouce/cloud.jpg", &power.HashMap{
    "name": "cloud.jpg",
  }, &outRes)
  if err != nil {
    panic(err)
  }

  c.JSON(200, res)
}