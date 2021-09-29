package routes

import (
  "github.com/gin-gonic/gin"
  "power-wechat-tutorial/controllers/wecom"
  "power-wechat-tutorial/controllers/wecom/accountService"
  "power-wechat-tutorial/controllers/wecom/externalContact"
  "power-wechat-tutorial/controllers/wecom/message"
  "power-wechat-tutorial/controllers/wecom/oa"
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
    wecomRouter.POST("externalContact/delStrategyTag", externalContact.APIExternalContactDelStrategyTag)
    wecomRouter.POST("externalContact/markTag", externalContact.APIExternalContactMarkTag)

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
    wecomRouter.POST("kf/customer/cancel_upgrade_service", accountService.APIAccountServiceCustomerCancelUpgradeService)

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
    wecomRouter.POST("message/send/text", message.APIMessageSendText)
    wecomRouter.POST("message/send/image", message.APIMessageSendImage)
    wecomRouter.POST("message/send/voice", message.APIMessageSendVoice)
    wecomRouter.POST("message/send/video", message.APIMessageSendVideo)
    wecomRouter.POST("message/send/file", message.APIMessageSendFile)
    wecomRouter.POST("message/send/textcard", message.APIMessageSendTextcard)
    wecomRouter.POST("message/send/news", message.APIMessageSendNews)
    wecomRouter.POST("message/send/mpnews", message.APIMessageSendMPNews)
    wecomRouter.POST("message/send/markdown", message.APIMessageSendMarkdown)
    wecomRouter.POST("message/send/miniprogram_notice", message.APIMessageSendMiniProgramNotice)

    // Handle template card route
    wecomRouter.POST("message/send/template_card/text_notice", message.APIMessageSendTemplateCardTextNotice)
    wecomRouter.POST("message/send/template_card/news_notice", message.APIMessageSendTemplateCardNewsNotice)
    wecomRouter.POST("message/send/template_card/button_interaction", message.APIMessageSendTemplateCardButtonInteraction)
    wecomRouter.POST("message/send/template_card/vote_interaction", message.APIMessageSendTemplateCardVoteInteraction)
    wecomRouter.POST("message/send/template_card/multiple_interaction", message.APIMessageSendTemplateCardMultipleInteraction)

    wecomRouter.POST("message/update/button/", message.APIMessageUpdateButton)
    wecomRouter.POST("message/update/template_card/text_notice", message.APIMessageUpdateTemplateCardTextNotice)
    wecomRouter.POST("message/update/template_card/news_notice", message.APIMessageUpdateTemplateCardNewsNotice)
    wecomRouter.POST("message/update/template_card/button_interaction", message.APIMessageUpdateTemplateCardButtonInteraction)
    wecomRouter.POST("message/update/template_card/vote_interaction", message.APIMessageUpdateTemplateCardVoteInteraction)
    wecomRouter.POST("message/update/template_card/multiple_interaction", message.APIMessageUpdateTemplateCardMultipleInteraction)

    // Handle message route
    wecomRouter.POST("message/recall", message.APIMessageRecall)

    // Handle app chat route
    wecomRouter.POST("appchat/create", message.APIAppChatCreate)
    wecomRouter.POST("appchat/update", message.APIAppChatUpdate)
    wecomRouter.POST("appchat/get", message.APIAppChatGet)

    wecomRouter.POST("appchat/send/text", message.APIAppChatSendText)
    wecomRouter.POST("appchat/send/image", message.APIAppChatSendImage)
    wecomRouter.POST("appchat/send/voice", message.APIAppChatSendVoice)
    wecomRouter.POST("appchat/send/video", message.APIAppChatSendVideo)
    wecomRouter.POST("appchat/send/file", message.APIAppChatSendFile)
    wecomRouter.POST("appchat/send/textcard", message.APIAppChatSendTextcard)
    wecomRouter.POST("appchat/send/news", message.APIAppChatSendNews)
    wecomRouter.POST("appchat/send/mpnews", message.APIAppChatSendMPNews)
    wecomRouter.POST("appchat/send/markdown", message.APIAppChatSendMarkdown)

    wecomRouter.POST("appchat/linkedcorp/send/text", message.APIAppChatLinkedCorpSendText)
    wecomRouter.POST("appchat/linkedcorp/send/image", message.APIAppChatLinkedCorpSendImage)
    wecomRouter.POST("appchat/linkedcorp/send/voice", message.APIAppChatLinkedCorpSendVoice)
    wecomRouter.POST("appchat/linkedcorp/send/video", message.APIAppChatLinkedCorpSendVideo)
    wecomRouter.POST("appchat/linkedcorp/send/file", message.APIAppChatLinkedCorpSendFile)
    wecomRouter.POST("appchat/linkedcorp/send/textcard", message.APIAppChatLinkedCorpSendTextcard)
    wecomRouter.POST("appchat/linkedcorp/send/news", message.APIAppChatLinkedCorpSendNews)
    wecomRouter.POST("appchat/linkedcorp/send/mpnews", message.APIAppChatLinkedCorpSendMPNews)
    wecomRouter.POST("appchat/linkedcorp/send/markdown", message.APIAppChatLinkedCorpSendMarkdown)
    wecomRouter.POST("appchat/linkedcorp/send/miniporgramnotice", message.APIAppChatLinkedCorpSendMiniProgramNotice)

    wecomRouter.POST("externalcontact/message/send/text", message.APIExternalContactMessageSendText)
    wecomRouter.POST("externalcontact/message/send/image", message.APIExternalContactMessageSendImage)
    wecomRouter.POST("externalcontact/message/send/voice", message.APIExternalContactMessageSendVoice)
    wecomRouter.POST("externalcontact/message/send/video", message.APIExternalContactMessageSendVideo)
    wecomRouter.POST("externalcontact/message/send/file", message.APIExternalContactMessageSendFile)
    wecomRouter.POST("externalcontact/message/send/textcard", message.APIExternalContactMessageSendTextCard)
    wecomRouter.POST("externalcontact/message/send/news", message.APIExternalContactMessageSendNews)
    wecomRouter.POST("externalcontact/message/send/mpnews", message.APIExternalContactMessageSendMPNews)
    wecomRouter.POST("externalcontact/message/send/miniprogram", message.APIExternalContactMessageSendMiniProgram)

    // Handle media route
    wecomRouter.POST("media/temp/upload/file", wecom.APIMediaUploadByURL)
    wecomRouter.POST("media/temp/upload/data", wecom.APIMediaUploadByData)
    wecomRouter.POST("media/upload/image/file", wecom.APIMediaUploadImgByPath)
    wecomRouter.POST("media/upload/image/data", wecom.APIMediaUploadImgByData)
    wecomRouter.GET("media/get", wecom.APIMediaGet)

    // Handle oa checkin route
    wecomRouter.POST("oa/checkin/getCorpCheckinOption", oa.APICheckinGetCorpCheckinOption)
    wecomRouter.POST("oa/checkin/getCheckinOption", oa.APICheckinGetCheckinOption)
    wecomRouter.POST("oa/checkin/getCheckinData", oa.APICheckinGetCheckinData)
    wecomRouter.POST("oa/checkin/getCheckinDayData", oa.APICheckinGetCheckinDayData)
    wecomRouter.POST("oa/checkin/getCheckinMonthData", oa.APICheckinGetCheckinMonthData)
    wecomRouter.POST("oa/checkin/getCheckinSchedulist", oa.APICheckinGetCheckinSchedulist)
    wecomRouter.POST("oa/checkin/setCheckinSchedulist", oa.APICheckinSetCheckinSchedulist)
    wecomRouter.POST("oa/checkin/addCheckinUserFace", oa.APICheckinAddCheckinUserFace)

    // Handle oa approval route
    wecomRouter.POST("oa/gettemplatedetail", oa.APIApprovalOaGetTemplateDetail)
    wecomRouter.POST("oa/getapprovalinfo", oa.APIApprovalOaGetApprovalInfo)
    wecomRouter.POST("oa/getapprovaldetail", oa.APIApprovalOaGetApprovalDetail)
    wecomRouter.POST("oa/getapprovaldata", oa.APIApprovalOaGetApprovalData)
    wecomRouter.POST("oa/vacation/getcorpconf", oa.APIApprovalVacationGetCorpConf)
    wecomRouter.POST("oa/vacation/getuservacationquota", oa.APIApprovalVacationGetUserVacationQuota)


    // Handle oa journal route
    wecomRouter.POST("oa/journal/get_record_list", oa.APIJournalGetRecordList)
    wecomRouter.POST("oa/journal/get_record_detail", oa.APIJournalGetRecordDetail)
    wecomRouter.POST("oa/journal/get_stat_list", oa.APIJournalGetStatList)



    // Handle oa meeting room route
    wecomRouter.POST("oa/meetingroom/add", oa.APIMeetingRoomAdd)
    wecomRouter.POST("oa/meetingroom/list", oa.APIMeetingRoomList)
    wecomRouter.POST("oa/meetingroom/edit", oa.APIMeetingRoomEdit)
    wecomRouter.POST("oa/meetingroom/del", oa.APIMeetingRoomDel)

    wecomRouter.POST("oa/meetingroom/get_booking_info", oa.APIMeetingRoomGetBookingInfo)
    wecomRouter.POST("oa/meetingroom/book", oa.APIMeetingRoomBook)
    wecomRouter.POST("oa/meetingroom/cancel_book", oa.APIMeetingRoomCancelBook)

    // Handle oa pstncc route
    wecomRouter.POST("pstncc/call", oa.APIPSTNCCCall)
    wecomRouter.POST("pstncc/getstates", oa.APIPSTNCCGetStates)


    // Handle oa calendar route
    wecomRouter.POST("oa/calendar/add", oa.APICalendarAdd)
    wecomRouter.POST("oa/calendar/update", oa.APICalendarUpdate)
    wecomRouter.POST("oa/calendar/get", oa.APICalendarGet)
    wecomRouter.POST("oa/calendar/del", oa.APICalendarDel)

    // Handle oa schedule route
    wecomRouter.POST("oa/schedule/add", oa.APIScheduleAdd)
    wecomRouter.POST("oa/schedule/update", oa.APIScheduleUpdate)
    wecomRouter.POST("oa/schedule/get", oa.APIScheduleGet)
    wecomRouter.POST("oa/schedule/del", oa.APIScheduleDel)

    // Handle meeting route
    wecomRouter.POST("meeting/create", oa.APIMeetingCreate)
    wecomRouter.POST("meeting/update", oa.APIMeetingUpdate)
    wecomRouter.POST("meeting/cancel", oa.APIMeetingCancel)
    wecomRouter.POST("meeting/getUserMeetingID", oa.APIMeetingGetUserMeetingID)

    // Handle living route
    wecomRouter.POST("living/create", oa.APILivingCreate)
    wecomRouter.POST("living/modify", oa.APILivingModify)
    wecomRouter.POST("living/cancel", oa.APILivingCancel)
    wecomRouter.POST("living/delete_replay_data", oa.APILivingDeleteReplayData)
    wecomRouter.POST("living/get_living_code", oa.APILivingGetLivingCode)
    wecomRouter.POST("living/get_user_all_livingid", oa.APILivingGetUserAllLivingID)
    wecomRouter.POST("living/get_living_info", oa.APILivingGetLivingInfo)
    wecomRouter.POST("living/get_watch_stat", oa.APILivingGetWatchStat)
    wecomRouter.POST("living/get_living_share_info", oa.APILivingGetLivingShareInfo)

    // Handle wedrive route
    wecomRouter.POST("wedrive/space_create", oa.APIWebDriveSpaceCreate)
    wecomRouter.POST("wedrive/space_acl_add", oa.APIWebDriveSpaceAcAdd)
    wecomRouter.POST("wedrive/file_list", oa.APIWebDriveFileList)
    wecomRouter.POST("wedrive/file_acl_add", oa.APIWebDriveFileAclAdd)

    // to be test with permission
    // Handle dial route
    wecomRouter.POST("dial/get_dial_record", oa.APIDialGetDialRecord)


    // to be test with permission
    // Handle corp group route
    wecomRouter.POST("corpgroup/corp/list_app_share_info", wecom.APICorpGroupCorpListAppShareInfo)
    wecomRouter.POST("corpgroup/corp/gettoken", wecom.APICorpGroupCorpGetToken)
    wecomRouter.POST("miniprogram/transfer_session", wecom.APICorpGroupMiniProgramTransferSession)

    // to be test with permission
    // Handle msg audit route
    wecomRouter.POST("msgaudit/get_permit_user_list", wecom.APIMsgAuditGetPermitUserList)
    wecomRouter.POST("msgaudit/check_single_agree", wecom.APIMsgAuditCheckSingleAgree)
    wecomRouter.POST("msgaudit/check_room_agree", wecom.APIMsgAuditCheckRoomAgree)
    wecomRouter.POST("msgaudit/groupchat/get", wecom.APIMsgAuditGroupChatGet)

    // to be test with permission
    // Handle invoice route
    wecomRouter.POST("invoice/reimburse/getinvoiceinfo", wecom.APIInvoiceReimburseGetInvoiceInfo)
    wecomRouter.POST("invoice/reimburse/updateinvoicestatus", wecom.APIInvoiceReimburseUpdateInvoiceStatus)
    wecomRouter.POST("invoice/reimburse/updatestatusbatch", wecom.APIInvoiceReimburseUpdateStatusBatch)
    wecomRouter.POST("invoice/reimburse/getinvoiceinfobatch", wecom.APIInvoiceReimburseGetInvoiceInfoBatch)

    // group robot
    // 群机器人
    wecomRouter.POST("/webhook/send/text", wecom.GroupRobotSendText)
    wecomRouter.POST("/webhook/send/markdown", wecom.GroupRobotSendMarkdown)
    wecomRouter.POST("/webhook/send/image", wecom.GroupRobotSendImage)
    wecomRouter.POST("/webhook/send/news-articles", wecom.GroupRobotSendNewsArticles)
    wecomRouter.POST("/webhook/send/file", wecom.GroupRobotSendFile)
    wecomRouter.POST("/webhook/send/template-card", wecom.GroupRobotSendTemplateCard)

  }
}
