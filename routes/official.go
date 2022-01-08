package routes

import (
  "github.com/gin-gonic/gin"
  "power-wechat-tutorial/controllers/official-account"
)

func InitOfficialAPIRoutes(r *gin.Engine) {

  officialRouter := r.Group("/official")
  {

    officialRouter.POST("/media/uploadTempImage", official_account.APIUploadTempImage)
    officialRouter.POST("/media/uploadTempVoice", official_account.APIUploadTempVoice)
    officialRouter.POST("/media/uploadTempVideo", official_account.APIUploadTempVideo)
    officialRouter.POST("/media/uploadTempFile", official_account.APIUploadTempFile)
    //officialRouter.POST("/media/uploadImage", officialAccount.APIUploadImage)
    officialRouter.POST("/media/get", official_account.APIGetMedia)

  }
}
