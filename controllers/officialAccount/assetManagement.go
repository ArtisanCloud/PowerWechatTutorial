package officialAccount

import (
  "github.com/gin-gonic/gin"
  "power-wechat-tutorial/services"
)

func APIMediaUploadTempFile(c *gin.Context)  {
  res, err := services.OfficialAccountApp.Media.UploadTempFile("./resource/cloud.jpg",nil)
  if err != nil {
    panic(err)
  }

  c.JSON(200, res)
}