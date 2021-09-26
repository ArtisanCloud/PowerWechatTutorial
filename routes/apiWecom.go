package routes

import (
  "github.com/gin-gonic/gin"
  "power-wechat-tutorial/controllers/wecom"
  "power-wechat-tutorial/controllers/wecom/accountService"
  "power-wechat-tutorial/controllers/wecom/externalContact"
  "power-wechat-tutorial/controllers/wecom/message"
  "power-wechat-tutorial/controllers/wecom/user"
)

func InitWecomAPIRoutes(r *gin.Engine) {

  wecomRouter := r.Group("/wecom")
  {

    // Handle user route
    wecomRouter.POST("/user/create", user.APIUserCreate)
    wecomRouter.GET("/user/get", user.APIUserGet)
    wecomRouter.PUT("/user/update", user.APIUserUpdate)
    wecomRouter.DELETE("/user/delete", user.APIUserDelete)
    wecomRouter.DELETE("/user/batch", user.APIUserBatchDelete)
    wecomRouter.GET("/users/simple", user.APIUserSimpleList)
    wecomRouter.GET("/users/detail", user.APIUserDetailList)
    wecomRouter.POST("/user/userIDToOpenID", user.APIUserUserIDToOpenID)
    wecomRouter.POST("/user/openIDToUserID", user.APIUserOpenIDToUserID)
    wecomRouter.GET("/user/authsucc", user.APIUserAuthAccept)
    wecomRouter.GET("/batch/invite", user.APIUserBatchInvite)
    wecomRouter.GET("/corp/qrcode", user.APIUserGetJoinQrCode)
    wecomRouter.GET("/user/getActiveStat", user.APIUserGetActiveStat)

    // Handle department route
    wecomRouter.POST("/department/create", externalContact.APIDepartmentCreate)
    wecomRouter.PUT("/department/update", externalContact.APIDepartmentUpdate)
    wecomRouter.DELETE("/department/delete", externalContact.APIDepartmentDelete)
    wecomRouter.GET("/department/list", externalContact.APIDepartmentList)

    // Handle tag route
    wecomRouter.POST("/tag/create", user.APITagCreate)
    wecomRouter.PUT("/tag/update", user.APITagUpdate)
    wecomRouter.DELETE("/tag/delete", user.APITagDelete)
    wecomRouter.GET("/tag/get", user.APITagUserGet)
    wecomRouter.POST("/tag/addTagUsers", user.APITagUserAdd)
    wecomRouter.DELETE("/tag/delTagUsers", user.APITagUserDel)
    wecomRouter.GET("/tag/list", user.APITagList)

    // Handle batch route
    wecomRouter.POST("/batch/syncUser", user.APIBatchSyncUser)
    wecomRouter.POST("/batch/replaceUser", user.APIBatchReplaceUser)
    wecomRouter.POST("/batch/replaceParty", user.APIBatchReplaceParty)
    wecomRouter.GET("/batch/getResult", user.APIBatchGetResult)

    // Handle linked corp route
    wecomRouter.POST("/linkedcorp/agent/getPermList", user.APILinkedCorpAgentGetPermList)
    wecomRouter.POST("/linkedcorp/user/get", user.APILinkedCorpUserGet)
    wecomRouter.POST("/linkedcorp/user/simplelist", user.APILinkedCorpUserSimpleList)
    wecomRouter.POST("/linkedcorp/user/list", user.APILinkedCorpUserList)
    wecomRouter.POST("/linkedcorp/department/list", user.APILinkedCorpDepartmentList)

    // Handle linked corp route
    wecomRouter.POST("export/simpleUser", user.APIExportSimpleUser)
    wecomRouter.POST("export/user", user.APIExportUser)
    wecomRouter.POST("export/department", user.APIExportDepartment)
    wecomRouter.POST("export/tagUser", user.APIExportTagUser)
    wecomRouter.GET("export/getResult", user.APIExportGetResult)

    // Handle external contact route
    wecomRouter.POST("externalContact/addContactWay", externalContact.APIExternalContactGetFollowUserList)
    wecomRouter.POST("externalContact/getContactWay", externalContact.APIExternalContactGetContactWay)
    wecomRouter.POST("externalContact/listContactWay", externalContact.APIExternalContactListContactWay)
    wecomRouter.POST("externalContact/updateContactWay", externalContact.APIExternalContactUpdateContactWay)
    wecomRouter.POST("externalContact/getFollowUserList", externalContact.APIExternalContactListContactWay)
    wecomRouter.POST("externalContact/delContactWay", externalContact.APIExternalContactDelContactWay)
    wecomRouter.POST("externalContact/closeTempChat", externalContact.APIExternalContactCloseTempChat)
    wecomRouter.POST("externalContact/list", externalContact.APIExternalContactList)
    wecomRouter.POST("externalContact/get", externalContact.APIExternalContactGet)
    wecomRouter.POST("externalContact/batch/get_by_user", externalContact.APIExternalContactBatchGetByUser)
    wecomRouter.POST("externalContact/remark", externalContact.APIExternalContactRemark)
    wecomRouter.POST("externalContact/customerStrategy/list", externalContact.APIExternalContactCustomerStrategyList)
    wecomRouter.POST("externalContact/customerStrategy/get", externalContact.APIExternalContactCustomerStrategyGet)
    wecomRouter.POST("externalContact/customerStrategy/get_range", externalContact.APIExternalContactCustomerStrategyGetRange)
    wecomRouter.POST("externalContact/customerStrategy/create", externalContact.APIExternalContactCustomerStrategyCreate)
    wecomRouter.POST("externalContact/customerStrategy/edit", externalContact.APIExternalContactCustomerStrategyEdit)
    wecomRouter.POST("externalContact/customerStrategy/del", externalContact.APIExternalContactCustomerStrategyDel)

    // Handle external contact tag route
    wecomRouter.POST("externalContact/getCorpTagList", externalContact.APIExternalContactGetCorpTagList)
    wecomRouter.POST("externalContact/addCorpTag", externalContact.APIExternalContactAddCorpTag)
    wecomRouter.POST("externalContact/editCorpTag", externalContact.APIExternalContactEditCorpTag)
    wecomRouter.POST("externalContact/delCXzorpTag", externalContact.APIExternalContactDelCorpTag)
    wecomRouter.POST("externalContact/getStrategyTagList", externalContact.APIExternalContactGetStrategyTagList)
    wecomRouter.POST("externalContact/addStrategyTag", externalContact.APIExternalContactAddStrategyTag)
    wecomRouter.POST("externalContact/editStrategyTag", externalContact.APIExternalContactEditStrategyTag)

    // Handle external contact transfer route
    wecomRouter.POST("externalContact/transferCustomer", externalContact.APIExternalContactTransferCustomer)
    wecomRouter.POST("externalContact/transferResult", externalContact.APIExternalContactTransferResult)
    wecomRouter.POST("externalContact/getUnassignedList", externalContact.APIExternalContactGetUnassignedList)
    wecomRouter.POST("externalContact/resignedTransferCustomer", externalContact.APIExternalContactResignedTransferCustomer)
    wecomRouter.POST("externalContact/resignedTransferResult", externalContact.APIExternalContactResignedTransferResult)
    wecomRouter.POST("externalContact/groupChatTransfer", externalContact.APIExternalContactGroupChatTransfer)

    // Handle external contact group chat route
    wecomRouter.POST("externalContact/groupChat/list", externalContact.APIExternalContactGroupChatList)
    wecomRouter.POST("externalContact/groupChat/get", externalContact.APIExternalContactGroupChatGet)
    wecomRouter.POST("externalContact/openGIDToChatID", externalContact.APIExternalContactOpenGIDToChatID)

    // Handle external contact moment route
    wecomRouter.POST("externalContact/getMomentList", externalContact.APIExternalContactGetMomentList)
    wecomRouter.POST("externalContact/getMomentTask", externalContact.APIExternalContactGetMomentTask)
    wecomRouter.POST("externalContact/getMomentCustomerList", externalContact.APIExternalContactGetMomentCustomerList)
    wecomRouter.POST("externalContact/getMomentSendResult", externalContact.APIExternalContactGetMomentSendResult)
    wecomRouter.POST("externalContact/getMomentComments", externalContact.APIExternalContactGetMomentComments)
    wecomRouter.POST("externalContact/momentStrategy/list", externalContact.APIExternalContactMomentStrategyList)
    wecomRouter.POST("externalContact/momentStrategy/get", externalContact.APIExternalContactMomentStrategyGet)
    wecomRouter.POST("externalContact/momentStrategy/get_range", externalContact.APIExternalContactMomentStrategyGetRange)
    wecomRouter.POST("externalContact/momentStrategy/create", externalContact.APIExternalContactMomentStrategyCreate)
    wecomRouter.POST("externalContact/momentStrategy/edit", externalContact.APIExternalContactMomentStrategyEdit)
    wecomRouter.POST("externalContact/momentStrategy/del", externalContact.APIExternalContactMomentStrategyDel)

    // Handle external contact message route
    wecomRouter.POST("externalContact/addMsgTemplate", externalContact.APIExternalContactAddMsgTemplate)
    wecomRouter.POST("externalContact/getGroupMsgListV2", externalContact.APIExternalContactGetGroupMsgListV2)
    wecomRouter.POST("externalContact/getGroupMsgTask", externalContact.APIExternalContactGetGroupMsgTask)
    wecomRouter.POST("externalContact/getGroupMsgSendResult", externalContact.APIExternalContactGetGroupMsgSendResult)
    wecomRouter.POST("externalContact/sendWelcomeMsg", externalContact.APIExternalContactSendWelcomeMsg)
    wecomRouter.POST("externalContact/groupWelcomeTemplate", externalContact.APIExternalContactGroupWelcomeTemplateAdd)
    wecomRouter.POST("externalContact/groupWelcomeTemplate/edit", externalContact.APIExternalContactGroupWelcomeTemplateEdit)
    wecomRouter.POST("externalContact/groupWelcomeTemplate/get", externalContact.APIExternalContactGroupWelcomeTemplateGet)
    wecomRouter.POST("externalContact/groupWelcomeTemplate/del", externalContact.APIExternalContactGroupWelcomeTemplateDel)

    // Handle external contact statics route
    wecomRouter.POST("externalContact/getUserBehaviorData", externalContact.APIExternalContactGetUserBehaviorData)
    wecomRouter.POST("externalContact/groupChat/statistic", externalContact.APIExternalContactGroupChatStatistic)

    // Handle account service route
    wecomRouter.POST("kf/account/add", accountService.APIAccountServiceAccountAdd)
    wecomRouter.POST("kf/account/del", accountService.APIAccountServiceAccountDel)
    wecomRouter.POST("kf/account/update", accountService.APIAccountServiceAccountUpdate)
    wecomRouter.POST("kf/account/list", accountService.APIAccountServiceAccountList)
    wecomRouter.POST("kf/add_contact_way", accountService.APIAccountServiceAddContactWay)

    // Handle account service servicer route
    wecomRouter.POST("kf/servicer/add", accountService.APIAccountServiceServicerAdd)
    wecomRouter.POST("kf/servicer/del", accountService.APIAccountServiceServicerDel)
    wecomRouter.POST("kf/servicer/list", accountService.APIAccountServiceServicerList)

    // Handle account service state route
    wecomRouter.POST("kf/service_state/get", accountService.APIAccountServiceStateGet)
    wecomRouter.POST("kf/service_state/trans", accountService.APIAccountServiceStateTrans)

    // Handle account service sync message route
    wecomRouter.POST("kf/sync_msg", accountService.APIAccountServiceSyncMsg)
    wecomRouter.POST("kf/send_msg", accountService.APIAccountServiceSendMsg)
    wecomRouter.POST("kf/send_msg_on_event", accountService.APIAccountServiceSendMsgOnEvent)

    // Handle account service customer route
    wecomRouter.POST("kf/customer/batchget", accountService.APIAccountServiceCustomerBatchGet)
    wecomRouter.POST("kf/customer/get_upgrade_service_config", accountService.APIAccountServiceCustomerGetUpgradeServiceConfig)
    wecomRouter.POST("kf/customer/upgrade_service", accountService.APIAccountServiceCustomerUpgradeService)

    // Handle agent route
    wecomRouter.POST("agent/get", wecom.APIAgentGet)
    wecomRouter.POST("agent/list", wecom.APIAgentList)
    wecomRouter.POST("agent/set", wecom.APIAgentSet)

    // Handle menu route
    wecomRouter.POST("menu/create", wecom.APIAgentMenuCreate)
    wecomRouter.POST("menu/get", wecom.APIAgentMenuGet)
    wecomRouter.POST("menu/delete", wecom.APIAgentMenuDelete)

    // Handle message route
    wecomRouter.POST("agent/set_workbench_template", wecom.APIAgentSetWorkbenchTemplate)
    wecomRouter.POST("agent/get_workbench_template", wecom.APIAgentGetWorkbenchTemplate)
    wecomRouter.POST("agent/set_workbench_data", wecom.APIAgentSetWorkbenchData)

    // Handle message send route
    wecomRouter.POST("qy/message/send/text", message.APIQYMessageSendText)
    wecomRouter.POST("qy/message/send/image", message.APIQYMessageSendImage)
    wecomRouter.POST("qy/message/send/voice", message.APIQYMessageSendVoice)
    wecomRouter.POST("qy/message/send/video", message.APIQYMessageSendVideo)
    wecomRouter.POST("qy/message/send/file", message.APIQYMessageSendFile)
    wecomRouter.POST("qy/message/send/textcard", message.APIQYMessageSendTextcard)
    wecomRouter.POST("qy/message/send/news", message.APIQYMessageSendNews)
    wecomRouter.POST("qy/message/send/mpnews", message.APIQYMessageSendMPNews)
    wecomRouter.POST("qy/message/send/markdown", message.APIQYMessageSendMarkdown)
    wecomRouter.POST("qy/message/send/miniprogram_notice", message.APIQYMessageSendMiniProgramNotice)

    // Handle template card route
    wecomRouter.POST("qy/message/send/template_card/text_notice", message.APIQYMessageSendTemplateCardTextNotice)
    wecomRouter.POST("qy/message/send/template_card/news_notice", message.APIQYMessageSendTemplateCardNewsNotice)
    wecomRouter.POST("qy/message/send/template_card/button_interaction", message.APIQYMessageSendTemplateCardButtonInteraction)
    wecomRouter.POST("qy/message/send/template_card/vote_interaction", message.APIQYMessageSendTemplateCardVoteInteraction)
    wecomRouter.POST("qy/message/send/template_card/multiple_interaction", message.APIQYMessageSendTemplateCardMultipleInteraction)

    wecomRouter.POST("qy/message/update/button/", message.APIQYMessageUpdateButton)
    wecomRouter.POST("qy/message/update/template_card/text_notice", message.APIQYMessageUpdateTemplateCardTextNotice)
    wecomRouter.POST("qy/message/update/template_card/news_notice", message.APIQYMessageUpdateTemplateCardNewsNotice)
    wecomRouter.POST("qy/message/update/template_card/button_interaction", message.APIQYMessageUpdateTemplateCardButtonInteraction)
    wecomRouter.POST("qy/message/update/template_card/vote_interaction", message.APIQYMessageUpdateTemplateCardVoteInteraction)
    wecomRouter.POST("qy/message/update/template_card/multiple_interaction", message.APIQYMessageUpdateTemplateCardMultipleInteraction)

    // Handle message route
    wecomRouter.POST("qy/message/recall", message.APIQYMessageRecall)

    // Handle app chat route
    wecomRouter.POST("qy/appchat/create", message.APIQYAppChatCreate)
    wecomRouter.POST("qy/appchat/update", message.APIQYAppChatUpdate)
    wecomRouter.POST("qy/appchat/get", message.APIQYAppChatGet)

    wecomRouter.POST("qy/appchat/send/text", message.APIQYAppChatSendText)
    wecomRouter.POST("qy/appchat/send/image", message.APIQYAppChatSendImage)
    wecomRouter.POST("qy/appchat/send/voice", message.APIQYAppChatSendVoice)
    wecomRouter.POST("qy/appchat/send/video", message.APIQYAppChatSendVideo)
    wecomRouter.POST("qy/appchat/send/file", message.APIQYAppChatSendFile)
    wecomRouter.POST("qy/appchat/send/textcard", message.APIQYAppChatSendTextcard)
    wecomRouter.POST("qy/appchat/send/news", message.APIQYAppChatSendNews)
    wecomRouter.POST("qy/appchat/send/mpnews", message.APIQYAppChatSendMpnews)
    wecomRouter.POST("qy/appchat/send/markdown", message.APIQYAppChatSendMarkdown)

    wecomRouter.POST("qy/appchat/linkedcorp/send/text", message.APIQAppChatLinkedCorpSendText)
    wecomRouter.POST("qy/appchat/linkedcorp/send/image", message.APIQAppChatLinkedCorpSendImage)
    wecomRouter.POST("qy/appchat/linkedcorp/send/voice", message.APIQAppChatLinkedCorpSendVoice)
    wecomRouter.POST("qy/appchat/linkedcorp/send/video", message.APIQAppChatLinkedCorpSendVideo)
    wecomRouter.POST("qy/appchat/linkedcorp/send/file", message.APIQAppChatLinkedCorpSendFile)
    wecomRouter.POST("qy/appchat/linkedcorp/send/textcard", message.APIQAppChatLinkedCorpSendTextcard)
    wecomRouter.POST("qy/appchat/linkedcorp/send/news", message.APIQAppChatLinkedCorpSendNews)
    wecomRouter.POST("qy/appchat/linkedcorp/send/mpnews", message.APIQAppChatLinkedCorpSendMpnews)
    wecomRouter.POST("qy/appchat/linkedcorp/send/markdown", message.APIQAppChatLinkedCorpSendMarkdown)

    wecomRouter.POST("qy/externalcontact/message/send/text", message.APIQYExternalContactMessageSendText)
    wecomRouter.POST("qy/externalcontact/message/send/image", message.APIQYExternalContactMessageSendImage)
    wecomRouter.POST("qy/externalcontact/message/send/voice", message.APIQYExternalContactMessageSendVoice)
    wecomRouter.POST("qy/externalcontact/message/send/video", message.APIQYExternalContactMessageSendVideo)
    wecomRouter.POST("qy/externalcontact/message/send/file", message.APIQYExternalContactMessageSendFile)
    wecomRouter.POST("qy/externalcontact/message/send/textcard", message.APIQYExternalContactMessageSendTextcard)
    wecomRouter.POST("qy/externalcontact/message/send/news", message.APIQYExternalContactMessageSendNews)
    wecomRouter.POST("qy/externalcontact/message/send/mpnews", message.APIQYExternalContactMessageSendMpnews)
    wecomRouter.POST("qy/externalcontact/message/send/markdown", message.APIQYExternalContactMessageSendMarkdown)


    // Handle media route
    wecomRouter.POST("qy/media/temp/upload/file", wecom.APIQYMediaUploadByURL)
    wecomRouter.POST("qy/media/temp/upload/data", wecom.APIQYMediaUploadByData)
    wecomRouter.POST("qy/media/upload/image", wecom.APIQYMediaUploadImgByPath)
    wecomRouter.POST("qy/media/upload/image", wecom.APIQYMediaUploadImgByData)
    wecomRouter.GET("qy/media/get", wecom.APIQYMediaGet)

    wecomRouter.POST("/")
  }
}
