package routes

import (
	"net/http"
	"os"
	"power-wechat-tutorial/controllers/wecom"
	account_service "power-wechat-tutorial/controllers/wecom/account-service"
	external_contact "power-wechat-tutorial/controllers/wecom/external-contact"
	"power-wechat-tutorial/controllers/wecom/message"
	"power-wechat-tutorial/controllers/wecom/oa"
	"power-wechat-tutorial/controllers/wecom/user"
	"power-wechat-tutorial/controllers/wecom/user/validate"

	"github.com/gin-gonic/gin"
)

func InitWecomAPIRoutes(r *gin.Engine) {

	r.GET("/WW_verify_L7QyPRfjldgxXN5t.txt", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, os.Getenv("app_oauth_verify_code"))
	})

	wecomRouter := r.Group("/wecom")
	{

		// 内部应用OAuth授权
		wecomRouter.GET("/oauth/authorize/user", wecom.WebAuthorizeUser)
		wecomRouter.GET("/callback/authorized/user", wecom.WebAuthorizedUser)
		wecomRouter.GET("/callback/authorized/v2/user", wecom.WebAuthorizedUserV2)
		wecomRouter.GET("/oauth/authorize/contact", wecom.WebAuthorizeContact)
		//r.GET("/callback/authorized/contact", request.ValidateRequestOAuthCallbackQRCode, wecom.WebAuthorizedContact)

		// Handle user route
		wecomRouter.POST("/user/create", validate.ValidateUserCreate, user.APIUserCreate)
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
		wecomRouter.POST("/department/create", user.APIDepartmentCreate)
		wecomRouter.PUT("/department/update", user.APIDepartmentUpdate)
		wecomRouter.DELETE("/department/delete", user.APIDepartmentDelete)
		wecomRouter.GET("/department/list", user.APIDepartmentList)
		wecomRouter.GET("/department/simplelist", user.APIDepartmentSimpleList)

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
		wecomRouter.POST("externalContact/addContactWay", external_contact.APIExternalContactGetFollowUserList)
		wecomRouter.POST("externalContact/getContactWay", external_contact.APIExternalContactGetContactWay)
		wecomRouter.POST("externalContact/listContactWay", external_contact.APIExternalContactListContactWay)
		wecomRouter.POST("externalContact/updateContactWay", external_contact.APIExternalContactUpdateContactWay)
		wecomRouter.POST("externalContact/getFollowUserList", external_contact.APIExternalContactListContactWay)
		wecomRouter.POST("externalContact/delContactWay", external_contact.APIExternalContactDelContactWay)
		wecomRouter.POST("externalContact/closeTempChat", external_contact.APIExternalContactCloseTempChat)
		wecomRouter.POST("externalContact/list", external_contact.APIExternalContactList)
		wecomRouter.POST("externalContact/get", external_contact.APIExternalContactGet)
		wecomRouter.POST("externalContact/batch/get_by_user", external_contact.APIExternalContactBatchGetByUser)
		wecomRouter.POST("externalContact/remark", external_contact.APIExternalContactRemark)
		wecomRouter.POST("externalContact/customerStrategy/list", external_contact.APIExternalContactCustomerStrategyList)
		wecomRouter.POST("externalContact/customerStrategy/get", external_contact.APIExternalContactCustomerStrategyGet)
		wecomRouter.POST("externalContact/customerStrategy/get_range", external_contact.APIExternalContactCustomerStrategyGetRange)
		wecomRouter.POST("externalContact/customerStrategy/create", external_contact.APIExternalContactCustomerStrategyCreate)
		wecomRouter.POST("externalContact/customerStrategy/edit", external_contact.APIExternalContactCustomerStrategyEdit)
		wecomRouter.POST("externalContact/customerStrategy/del", external_contact.APIExternalContactCustomerStrategyDel)

		// Handle external contact tag route
		wecomRouter.POST("externalContact/getCorpTagList", external_contact.APIExternalContactGetCorpTagList)
		wecomRouter.POST("externalContact/addCorpTag", external_contact.APIExternalContactAddCorpTag)
		wecomRouter.POST("externalContact/editCorpTag", external_contact.APIExternalContactEditCorpTag)
		wecomRouter.POST("externalContact/delCorpTag", external_contact.APIExternalContactDelCorpTag)
		wecomRouter.POST("externalContact/getStrategyTagList", external_contact.APIExternalContactGetStrategyTagList)
		wecomRouter.POST("externalContact/addStrategyTag", external_contact.APIExternalContactAddStrategyTag)
		wecomRouter.POST("externalContact/editStrategyTag", external_contact.APIExternalContactEditStrategyTag)
		wecomRouter.POST("externalContact/delStrategyTag", external_contact.APIExternalContactDelStrategyTag)
		wecomRouter.POST("externalContact/markTag", external_contact.APIExternalContactMarkTag)

		// Handle external contact transfer route
		wecomRouter.POST("externalContact/transferCustomer", external_contact.APIExternalContactTransferCustomer)
		wecomRouter.POST("externalContact/transferResult", external_contact.APIExternalContactTransferResult)
		wecomRouter.POST("externalContact/getUnassignedList", external_contact.APIExternalContactGetUnassignedList)
		wecomRouter.POST("externalContact/resignedTransferCustomer", external_contact.APIExternalContactResignedTransferCustomer)
		wecomRouter.POST("externalContact/resignedTransferResult", external_contact.APIExternalContactResignedTransferResult)
		wecomRouter.POST("externalContact/groupChatTransfer", external_contact.APIExternalContactGroupChatTransfer)

		// Handle external contact group chat route
		wecomRouter.POST("externalContact/groupChat/list", external_contact.APIExternalContactGroupChatList)
		wecomRouter.POST("externalContact/groupChat/get", external_contact.APIExternalContactGroupChatGet)
		wecomRouter.POST("externalContact/openGIDToChatID", external_contact.APIExternalContactOpenGIDToChatID)

		// Handle external contact moment route
		wecomRouter.POST("externalContact/getMomentList", external_contact.APIExternalContactGetMomentList)
		wecomRouter.POST("externalContact/getMomentTask", external_contact.APIExternalContactGetMomentTask)
		wecomRouter.POST("externalContact/getMomentCustomerList", external_contact.APIExternalContactGetMomentCustomerList)
		wecomRouter.POST("externalContact/getMomentSendResult", external_contact.APIExternalContactGetMomentSendResult)
		wecomRouter.POST("externalContact/getMomentComments", external_contact.APIExternalContactGetMomentComments)
		wecomRouter.POST("externalContact/momentStrategy/list", external_contact.APIExternalContactMomentStrategyList)
		wecomRouter.POST("externalContact/momentStrategy/get", external_contact.APIExternalContactMomentStrategyGet)
		wecomRouter.POST("externalContact/momentStrategy/get_range", external_contact.APIExternalContactMomentStrategyGetRange)
		wecomRouter.POST("externalContact/momentStrategy/create", external_contact.APIExternalContactMomentStrategyCreate)
		wecomRouter.POST("externalContact/momentStrategy/edit", external_contact.APIExternalContactMomentStrategyEdit)
		wecomRouter.POST("externalContact/momentStrategy/del", external_contact.APIExternalContactMomentStrategyDel)

		// Handle external contact message route
		wecomRouter.POST("externalContact/addMsgTemplate", external_contact.APIExternalContactAddMsgTemplate)
		wecomRouter.POST("externalContact/getGroupMsgListV2", external_contact.APIExternalContactGetGroupMsgListV2)
		wecomRouter.POST("externalContact/getGroupMsgTask", external_contact.APIExternalContactGetGroupMsgTask)
		wecomRouter.POST("externalContact/getGroupMsgSendResult", external_contact.APIExternalContactGetGroupMsgSendResult)
		wecomRouter.POST("externalContact/sendWelcomeMsg", external_contact.APIExternalContactSendWelcomeMsg)
		wecomRouter.POST("externalContact/groupWelcomeTemplate", external_contact.APIExternalContactGroupWelcomeTemplateAdd)
		wecomRouter.POST("externalContact/groupWelcomeTemplate/edit", external_contact.APIExternalContactGroupWelcomeTemplateEdit)
		wecomRouter.POST("externalContact/groupWelcomeTemplate/get", external_contact.APIExternalContactGroupWelcomeTemplateGet)
		wecomRouter.POST("externalContact/groupWelcomeTemplate/del", external_contact.APIExternalContactGroupWelcomeTemplateDel)

		// Handle external contact statics route
		wecomRouter.POST("externalContact/getUserBehaviorData", external_contact.APIExternalContactGetUserBehaviorData)
		wecomRouter.POST("externalContact/groupChat/statistic", external_contact.APIExternalContactGroupChatStatistic)

		// Handle account service route
		wecomRouter.POST("kf/account/add", account_service.APIAccountServiceAccountAdd)
		wecomRouter.POST("kf/account/del", account_service.APIAccountServiceAccountDel)
		wecomRouter.POST("kf/account/update", account_service.APIAccountServiceAccountUpdate)
		wecomRouter.POST("kf/account/list", account_service.APIAccountServiceAccountList)
		wecomRouter.POST("kf/add_contact_way", account_service.APIAccountServiceAddContactWay)

		// Handle account service servicer route
		wecomRouter.POST("kf/servicer/add", account_service.APIAccountServiceServicerAdd)
		wecomRouter.POST("kf/servicer/del", account_service.APIAccountServiceServicerDel)
		wecomRouter.POST("kf/servicer/list", account_service.APIAccountServiceServicerList)

		// Handle account service state route
		wecomRouter.POST("kf/service_state/get", account_service.APIAccountServiceStateGet)
		wecomRouter.POST("kf/service_state/trans", account_service.APIAccountServiceStateTrans)

		// Handle account service sync message route
		wecomRouter.POST("kf/sync_msg", account_service.APIAccountServiceSyncMsg)
		wecomRouter.POST("kf/send_msg", account_service.APIAccountServiceSendMsg)
		wecomRouter.POST("kf/send_msg_on_event", account_service.APIAccountServiceSendMsgOnEvent)

		// Handle account service customer route
		wecomRouter.POST("kf/customer/batchget", account_service.APIAccountServiceCustomerBatchGet)
		wecomRouter.POST("kf/customer/get_upgrade_service_config", account_service.APIAccountServiceCustomerGetUpgradeServiceConfig)
		wecomRouter.POST("kf/customer/upgrade_service", account_service.APIAccountServiceCustomerUpgradeService)
		wecomRouter.POST("kf/customer/cancel_upgrade_service", account_service.APIAccountServiceCustomerCancelUpgradeService)

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
		wecomRouter.POST("message/appchat/create", message.APIAppChatCreate)
		wecomRouter.POST("message/appchat/update", message.APIAppChatUpdate)
		wecomRouter.POST("message/appchat/get", message.APIAppChatGet)

		wecomRouter.POST("message/appchat/send/text", message.APIAppChatSendText)
		wecomRouter.POST("message/appchat/send/image", message.APIAppChatSendImage)
		wecomRouter.POST("message/appchat/send/voice", message.APIAppChatSendVoice)
		wecomRouter.POST("message/appchat/send/video", message.APIAppChatSendVideo)
		wecomRouter.POST("message/appchat/send/file", message.APIAppChatSendFile)
		wecomRouter.POST("message/appchat/send/textcard", message.APIAppChatSendTextcard)
		wecomRouter.POST("message/appchat/send/news", message.APIAppChatSendNews)
		wecomRouter.POST("message/appchat/send/mpnews", message.APIAppChatSendMPNews)
		wecomRouter.POST("message/appchat/send/markdown", message.APIAppChatSendMarkdown)

		wecomRouter.POST("message/appchat/linkedcorp/send/text", message.APIAppChatLinkedCorpSendText)
		wecomRouter.POST("message/appchat/linkedcorp/send/image", message.APIAppChatLinkedCorpSendImage)
		wecomRouter.POST("message/appchat/linkedcorp/send/voice", message.APIAppChatLinkedCorpSendVoice)
		wecomRouter.POST("message/appchat/linkedcorp/send/video", message.APIAppChatLinkedCorpSendVideo)
		wecomRouter.POST("message/appchat/linkedcorp/send/file", message.APIAppChatLinkedCorpSendFile)
		wecomRouter.POST("message/appchat/linkedcorp/send/textcard", message.APIAppChatLinkedCorpSendTextcard)
		wecomRouter.POST("message/appchat/linkedcorp/send/news", message.APIAppChatLinkedCorpSendNews)
		wecomRouter.POST("message/appchat/linkedcorp/send/mpnews", message.APIAppChatLinkedCorpSendMPNews)
		wecomRouter.POST("message/appchat/linkedcorp/send/markdown", message.APIAppChatLinkedCorpSendMarkdown)
		wecomRouter.POST("message/appchat/linkedcorp/send/miniporgramnotice", message.APIAppChatLinkedCorpSendMiniProgramNotice)

		wecomRouter.POST("message/externalcontact/send/text", message.APIExternalContactMessageSendText)
		wecomRouter.POST("message/externalcontact/send/image", message.APIExternalContactMessageSendImage)
		wecomRouter.POST("message/externalcontact/send/voice", message.APIExternalContactMessageSendVoice)
		wecomRouter.POST("message/externalcontact/send/video", message.APIExternalContactMessageSendVideo)
		wecomRouter.POST("message/externalcontact/send/file", message.APIExternalContactMessageSendFile)
		wecomRouter.POST("message/externalcontact/send/textcard", message.APIExternalContactMessageSendTextCard)
		wecomRouter.POST("message/externalcontact/send/news", message.APIExternalContactMessageSendNews)
		wecomRouter.POST("message/externalcontact/send/mpnews", message.APIExternalContactMessageSendMPNews)
		wecomRouter.POST("message/externalcontact/send/miniprogram", message.APIExternalContactMessageSendMiniProgram)

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

		// 企业微信小程序 MiniProgram
		wecomRouter.POST("/miniprogram/code2Session", wecom.Code2Session)
	}

	// Handle user callback route
	r.GET("user/callback", user.CallbackVerify)
	r.POST("user/callback", user.CallbackNotify)

	// Handle message callback route
	r.GET("message/callback", message.CallbackVerify)
	r.POST("message/callback", message.CallbackNotify)
	r.POST("message/testbuffer", message.TestBuffer)

}
