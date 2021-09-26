package routes

import (
  "github.com/gin-gonic/gin"
  "power-wechat-tutorial/controllers/officialAccount"
)

func InitOfficialAPIRoutes(r *gin.Engine) {

  officialRouter := r.Group("/official")
  {

    officialRouter.POST("/media/uploadTempImage", officialAccount.APIUploadTempImage)
    officialRouter.POST("/media/uploadTempVoice", officialAccount.APIUploadTempVoice)
    officialRouter.POST("/media/uploadTempVideo", officialAccount.APIUploadTempVideo)
    officialRouter.POST("/media/uploadTempFile", officialAccount.APIUploadTempFile)
    //officialRouter.POST("/media/uploadImage", officialAccount.APIUploadImage)
    officialRouter.POST("/media/get", officialAccount.APIGetMedia)

  }
}
