package routes

import (
  "github.com/gin-gonic/gin"
  "power-wechat-tutorial/controllers/official-account"
)

func InitOfficialAPIRoutes(r *gin.Engine) {

  officialRouter := r.Group("/official")
  {
    // Base
    officialRouter.GET("/base/clearQuota", official_account.ClearQuota)
    officialRouter.GET("/base/callbackIP", official_account.GetCallbackIP)

    // jssdk
    officialRouter.GET("/jssdk/config", official_account.JSSDKBuildConfig)

    // 临时素材管理
    officialRouter.POST("/media/uploadImage", official_account.APIUploadImage)
    officialRouter.POST("/media/uploadVoice", official_account.APIUploadVoice)
    officialRouter.POST("/media/uploadVideo", official_account.APIUploadVideo)
    officialRouter.POST("/media/uploadThumb", official_account.APIUploadThumb)
    officialRouter.GET("/media/get", official_account.APIGetMedia)

    // 永久素材管理
    officialRouter.POST("/material/uploadImage", official_account.APIUploadMaterialImage)
    officialRouter.POST("/material/uploadVoice", official_account.APIUploadMaterialVoice)
    officialRouter.POST("/material/uploadVideo", official_account.APIUploadMaterialVideo)
    officialRouter.POST("/material/uploadThumb", official_account.APIUploadMaterialThumb)
    officialRouter.GET("/material/list", official_account.APIGetMaterialList)
    officialRouter.GET("/material/get", official_account.APIGetMaterial)

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
    officialRouter.POST("/customerService/invite", official_account.CustomerInvite)
    officialRouter.POST("/customerService/sendText", official_account.CustomerMessageSendText)
    officialRouter.POST("/customerService/sendImage", official_account.CustomerMessageSendImage)
    officialRouter.POST("/customerService/sendVideo", official_account.CustomerMessageSendVideo)
    officialRouter.POST("/customerService/sendVoice", official_account.CustomerMessageSendVoice)
    officialRouter.POST("/customerService/sendLink", official_account.CustomerMessageSendLink)
    officialRouter.POST("/customerService/sendMusic", official_account.CustomerMessageSendMusic)
    officialRouter.POST("/customerService/sendNews", official_account.CustomerMessageSendMusic)
    officialRouter.POST("/customerService/sendRaw", official_account.CustomerMessageSendRaw)
    officialRouter.POST("/customerService/send", official_account.CustomerMessageSend)

    // 客服会话控制
    officialRouter.POST("/customerServiceSession/create", official_account.CustomerSessionCreate)
    officialRouter.DELETE("/customerServiceSession/close", official_account.CustomerSessionClose)
    officialRouter.GET("/customerServiceSession/get", official_account.GetCustomerSession)
    officialRouter.GET("/customerServiceSession/list", official_account.CustomerSessionList)
    officialRouter.GET("/customerServiceSession/waiting", official_account.CustomerSessionWaiting)

    // 数据统计
    officialRouter.GET("/dateCube/getUserSummary", official_account.GetUserSummary)
    officialRouter.GET("/dateCube/getUserCumulate", official_account.GetUserCumulate)
    officialRouter.GET("/dateCube/articleSummary", official_account.ArticleSummary)
    officialRouter.GET("/dateCube/articleTotal", official_account.ArticleTotal)
    officialRouter.GET("/dateCube/userReadSummary", official_account.UserReadSummary)
    officialRouter.GET("/dateCube/userShareSummary", official_account.UserShareSummary)
    officialRouter.GET("/dateCube/userShareHourly", official_account.UserShareHourly)
    officialRouter.GET("/dateCube/upstreamMessageSummary", official_account.UpstreamMessageSummary)
    officialRouter.GET("/dateCube/upstreamMessageHourly", official_account.UpstreamMessageHourly)
    officialRouter.GET("/dateCube/upstreamMessageWeekly", official_account.UpstreamMessageWeekly)
    officialRouter.GET("/dateCube/upstreamMessageMonthly", official_account.UpstreamMessageMonthly)
    officialRouter.GET("/dateCube/upstreamMessageDistSummary", official_account.UpstreamMessageDistSummary)
    officialRouter.GET("/dateCube/upstreamMessageDistWeekly", official_account.UpstreamMessageDistWeekly)
    officialRouter.GET("/dateCube/upstreamMessageDistMonthly", official_account.UpstreamMessageDistMonthly)

    // 生成二维码
    officialRouter.GET("/qrcode/temp", official_account.GetTempQrCode)
    officialRouter.GET("/qrcode/forever", official_account.GetForeverQrCode)

    // 短key托管
    officialRouter.GET("/shorten/gen", official_account.ShortGenKey)
    officialRouter.GET("/shorten/fetch", official_account.FetchShortGen)

    // 自动回复
    officialRouter.GET("/autoReply/current", official_account.AutoReplyCurrent)

    // OAUTH2
    officialRouter.GET("/oauth/userFromCode", official_account.UserFromCode)
    officialRouter.GET("/oauth/userFromToken", official_account.UserFromCode)
  }
}
