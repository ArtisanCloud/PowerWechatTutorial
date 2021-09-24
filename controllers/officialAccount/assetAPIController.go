package officialAccount

import (
  "github.com/gin-gonic/gin"
  "io"
  "net/http"
  "power-wechat-tutorial/services"
)

// APIUploadTempFile 新增临时文件
// https://developers.weixin.qq.com/doc/offiaccount/Asset_Management/New_temporary_materials.html
func APIUploadTempImage(c *gin.Context) {
  res, err := services.OfficialAccountApp.Media.UploadTempImage("./resource/cloud.jpg", nil)

  if err != nil {
    panic(err)
  }
  c.JSON(http.StatusOK, res)
}

func APIUploadTempVoice(c *gin.Context) {
  res, err := services.OfficialAccountApp.Media.UploadTempVoice("./resource/cha-cha-ender.mp3", nil)

  if err != nil {
    panic(err)
  }
  c.JSON(200, res)
}

func APIUploadTempVideo(c *gin.Context) {
  res, err := services.OfficialAccountApp.Media.UploadTempVideo("./resource/3d_ocean_1590675653.mp4", nil)
  if err != nil {
    panic(err)
  }

  c.JSON(200, res)
}
func APIUploadTempFile(c *gin.Context) {
  res, err := services.OfficialAccountApp.Media.UploadTempFile("./resource/3d_ocean_1590675653.mp4", nil)
  if err != nil {
    panic(err)
  }

  c.JSON(200, res)
}

// APIGetMedia 获取临时素材
func APIGetMedia(c *gin.Context) {
  res, err := services.OfficialAccountApp.Media.Get(c.DefaultQuery("mediaId", "YbE2OL2Wz5b09q8rw1FGhgeEPsQBDbSxzpZHmZ7Zk_Yz7eMzql7xfCy7U-9mcHFm"))
  if err != nil {
    panic(err)
  }

  io.Copy(c.Writer, res.GetBody())
}

//// APIUploadImage 新增永久文件
//// https://developers.weixin.qq.com/doc/offiaccount/Asset_Management/Adding_Permanent_Assets.html
//func APIUploadImage(c *gin.Context) {
//  res, err := services.OfficialAccountApp.Media.UploadImage("./resource/cloud.jpg", nil)
//  if err != nil {
//    panic(err)
//  }
//
//  c.JSON(200, res)
//}
//func APIUploadVoice(c *gin.Context) {
//  var outRes interface{}
//  _, err := services.OfficialAccountApp.Media.UploadImage("./resource/cloud.jpg", nil)
//  if err != nil {
//    panic(err)
//  }
//
//  c.JSON(200, outRes)
//}
//func APIUploadVideo(c *gin.Context) {
//  var outRes interface{}
//  _, err := services.OfficialAccountApp.Media.UploadImage("./resource/cloud.jpg", nil)
//  if err != nil {
//    panic(err)
//  }
//
//  c.JSON(200, outRes)
//}
