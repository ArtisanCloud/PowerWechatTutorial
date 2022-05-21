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

    // 用户管理
    officialRouter.GET("/user/get", official_account.GetUserInfo)
    officialRouter.GET("/user/batchGet", official_account.GetBatchUserInfo)
    officialRouter.POST("/user/remark", official_account.UserRemark)
    officialRouter.GET("/user/getBlockList", official_account.GetUserBlacklist)
    officialRouter.POST("/user/block", official_account.UserBlock)
    officialRouter.POST("/user/unBlock", official_account.UserUnBlock)
    officialRouter.POST("/user/changeOpenID", official_account.UserChangeOpenID)

    // 用户标签管理
    officialRouter.GET("/userTag/list", official_account.GetUserTagList)
    officialRouter.POST("/userTag", official_account.UserTagCreate)
    officialRouter.PUT("/userTag", official_account.UserTagUpdate)
    officialRouter.DELETE("/userTag", official_account.UserTagDelete)
    officialRouter.GET("/userTag/getUserTagsByOpenID", official_account.GetUserTagsByOpenID)
    officialRouter.GET("/userTag/getUsersOfTag", official_account.GetUsersOfTag)
    officialRouter.POST("/userTag/batchTagUser", official_account.UserTagBatchTagUsers)
    officialRouter.POST("/userTag/unBatchTagUser", official_account.UserTagBatchUnTagUsers)

    // 客服消息
    officialRouter.GET("/customerService/list", official_account.GetCustomerList)
    officialRouter.GET("/customerService/online", official_account.GetCustomerOnline)
    officialRouter.POST("/customerService/create", official_account.CustomerCreate)
    officialRouter.PUT("/customerService/update", official_account.CustomerUpdate)
    officialRouter.DELETE("/customerService/delete", official_account.CustomerDelete)
    officialRouter.POST("/customerService/setAvatar", official_account.CustomerSetAvatar)
    officialRouter.POST("/customerService/messages", official_account.CustomerMessages)
    officialRouter.POST("/customerService/send", official_account.CustomerMessageSend)
    officialRouter.POST("/customerService/invite", official_account.CustomerInvite)

    officialRouter.POST("/customerServiceSession/create", official_account.CustomerSessionCreate)
    officialRouter.DELETE("/customerServiceSession/close", official_account.CustomerSessionClose)
    officialRouter.GET("/customerServiceSession/get", official_account.GetCustomerSession)
    officialRouter.GET("/customerServiceSession/list", official_account.CustomerSessionList)
    officialRouter.GET("/customerServiceSession/waiting", official_account.CustomerSessionWaiting)
  }
}
