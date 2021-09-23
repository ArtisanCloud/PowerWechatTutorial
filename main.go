package main

import (
  "github.com/gin-gonic/gin"
  "log"
  "power-wechat-tutorial/controllers/miniprogram"
  "power-wechat-tutorial/controllers/payment"
  "power-wechat-tutorial/controllers/wecom"
  "power-wechat-tutorial/controllers/wecom/accountService"
  "power-wechat-tutorial/controllers/wecom/externalContact"
  "power-wechat-tutorial/services"
)

var Host = ""
var Port = "8888"

func main() {

  var err error
  services.PaymentApp, err = services.NewWXPaymentApp(nil)
  if err != nil || services.PaymentApp == nil {
    panic(err)
  }

  services.MiniProgramApp, err = services.NewMiniMiniProgramService()
  if err != nil || services.MiniProgramApp == nil {
    panic(err)
  }

  services.WeComApp, err = services.NewWeComService()
  if err != nil || services.WeComApp == nil {
    panic(err)
  }
  services.WeComContactApp, err = services.NewWeComContactService()
  if err != nil || services.WeComContactApp == nil {
    panic(err)
  }

  r := gin.Default()

  // Payment App Router
  apiRouterPayment := r.Group("/payment")
  {
    // Handle the pay route
    apiRouterPayment.GET("/order/make", payment.APIMakeOrder)
    apiRouterPayment.POST("/wx/notify", payment.CallbackWXNotify)
    apiRouterPayment.GET("/order/query", payment.APIQueryOrder)
    apiRouterPayment.GET("/order/close", payment.APICloseOrder)
    apiRouterPayment.Static("/wx/payment", "./web")

    // Handle the bill route
    apiRouterPayment.GET("/bill/downloadURL", payment.APIBillDownloadURL)

  }

  // MiniProgram App Router
  routerMiniProgram := r.Group("/miniprogram")
  {
    // Handle the auth route
    routerMiniProgram.GET("/auth", miniprogram.APISNSSession)
    routerMiniProgram.POST("/auth/checkEncryptedData", miniprogram.APICheckEncryptedData)
    routerMiniProgram.GET("/auth/getPaidUnionID", miniprogram.APIGetPaidUnionID)

    // Handle the data cube analysis route
    routerMiniProgram.GET("/datacube/getDailyRetain", miniprogram.APIGetDailyRetain)
    routerMiniProgram.GET("/datacube/getMonthlyRetain", miniprogram.APIGetMonthlyRetain)
    routerMiniProgram.GET("/datacube/getWeeklyRetain", miniprogram.APIGetWeeklyRetain)
    routerMiniProgram.GET("/datacube/getDailySummary", miniprogram.APIGetDailySummary)

    routerMiniProgram.GET("/datacube/getDailyVisitTrend", miniprogram.APIGetDailyVisitTrend)
    routerMiniProgram.GET("/datacube/getMonthlyVisitTrend", miniprogram.APIGetMonthlyVisitTrend)
    routerMiniProgram.GET("/datacube/getWeeklyVisitTrend", miniprogram.APIGetWeeklyVisitTrend)
    routerMiniProgram.GET("/datacube/getPerformanceData", miniprogram.APIGetPerformanceData)
    routerMiniProgram.GET("/datacube/getUserPortrait", miniprogram.APIGetUserPortrait)
    routerMiniProgram.GET("/datacube/getVisitDistribution", miniprogram.APIGetVisitDistribution)
    routerMiniProgram.GET("/datacube/getVisitPage", miniprogram.APIGetVisitPage)

    // Handle the customer service message  route
    routerMiniProgram.GET("/customerServiceMessage/send", miniprogram.APICustomerServiceMessageSend)
    routerMiniProgram.GET("/customerServiceMessage/setTyping", miniprogram.APICustomerServiceMessageSetTyping)
    routerMiniProgram.GET("/customerServiceMessage/uploadTempMediaByFile", miniprogram.APICustomerServiceMessageUploadTempMediaByFile)
    routerMiniProgram.GET("/customerServiceMessage/uploadTempMediaByData", miniprogram.APICustomerServiceMessageUploadTempMediaByData)
    routerMiniProgram.GET("/customerServiceMessage/getTempMedia", miniprogram.APICustomerServiceMessageGetTempMedia)

    // Handle the uniform message  route
    routerMiniProgram.GET("/uniformMessage/send", miniprogram.APIUniformMessageSend)

    // Handle the updatable message  route
    routerMiniProgram.GET("/updatableMessage/createActivityID", miniprogram.APIUpdatableMessageCreateActivityID)
    routerMiniProgram.GET("/updatableMessage/updatableMessage", miniprogram.APIUpdatableMessageUpdatableMessage)

    // Handle the plugin manager route
    routerMiniProgram.GET("/pluginManager/applyPlugin", miniprogram.APIPluginManagerApplyPlugin)
    routerMiniProgram.GET("/pluginManager/getPluginDevApplyList", miniprogram.APIPluginManagerGetPluginDevApplyList)
    routerMiniProgram.GET("/pluginManager/getPluginList", miniprogram.APIPluginManagerGetPluginList)
    routerMiniProgram.GET("/pluginManager/setDevPluginApplyStatus", miniprogram.APIPluginManagerSetDevPluginApplyStatus)
    routerMiniProgram.GET("/pluginManager/unbindPlugin", miniprogram.APIPluginManagerUnbindPlugin)

    // Handle the nearby Poi route
    routerMiniProgram.GET("/nearbyPoi/add", miniprogram.APINearbyPoiAdd)
    routerMiniProgram.GET("/nearbyPoi/delete", miniprogram.APINearbyPoiDelete)
    routerMiniProgram.GET("/nearbyPoi/getList", miniprogram.APINearbyPoiGetList)
    routerMiniProgram.GET("/nearbyPoi/setShowStatus", miniprogram.APINearbySetShowStatus)

    // Handle the wxa code route
    routerMiniProgram.GET("/wxaCode/createQRCode", miniprogram.APIWXACodeCreateQRCode)
    routerMiniProgram.GET("/wxaCode/get", miniprogram.APIWXACodeGet)
    routerMiniProgram.GET("/wxaCode/getUnlimited", miniprogram.APIWXACodeGetUnlimited)

    // Handle the url scheme route
    routerMiniProgram.GET("/urlScheme/generate", miniprogram.APIURLSchemeGenerate)

    // Handle the url link route
    routerMiniProgram.GET("/urlLink/generate", miniprogram.APIURLLinkGenerate)

    // Handle the security route
    routerMiniProgram.GET("/security/imgSecCheckByPath", miniprogram.APISecurityImgSecCheckByPath)
    routerMiniProgram.GET("/security/imgSecCheckByData", miniprogram.APISecurityImgSecCheckByData)
    routerMiniProgram.GET("/security/mediaCheckAsync", miniprogram.APISecurityMediaCheckAsync)
    routerMiniProgram.GET("/security/msgSecCheck", miniprogram.APISecurityMsgSecCheck)

    // Handle the image route
    routerMiniProgram.GET("/image/aiCropByURL", miniprogram.APIImgAICropByURL)
    routerMiniProgram.GET("/image/aiCropByData", miniprogram.APIImgAICropByData)
    routerMiniProgram.GET("/image/scanQRCodeByURL", miniprogram.APIImgScanQRCodeByURL)
    routerMiniProgram.GET("/image/scanQRCodeByData", miniprogram.APIImgScanQRCodeByData)
    routerMiniProgram.GET("/image/superResolutionByURL", miniprogram.APIImgSuperResolutionByURL)
    routerMiniProgram.GET("/image/superResolutionByData", miniprogram.APIImgSuperResolutionByData)

    // Handle the image route
    routerMiniProgram.GET("/delivery/abnormalConfirm", miniprogram.APIImmediateDeliveryAbnormalConfirm)
    routerMiniProgram.GET("/delivery/addOrder", miniprogram.APIImmediateDeliveryAddOrder)
    routerMiniProgram.GET("/delivery/addTip", miniprogram.APIImmediateDeliveryAddTip)
    routerMiniProgram.GET("/delivery/bindAccount", miniprogram.APIImmediateDeliveryBindAccount)
    routerMiniProgram.GET("/delivery/cancelOrder", miniprogram.APIImmediateDeliveryCancelOrder)
    routerMiniProgram.GET("/delivery/getAllImmeDelivery", miniprogram.APIImmediateDeliveryGetAllImmeDelivery)
    routerMiniProgram.GET("/delivery/getBindAccount", miniprogram.APIImmediateDeliveryGetBindAccount)
    routerMiniProgram.GET("/delivery/getOrder", miniprogram.APIImmediateDeliveryGetOrder)
    routerMiniProgram.GET("/delivery/mockUpdateOrder", miniprogram.APIImmediateDeliveryMockUpdateOrder)
    routerMiniProgram.GET("/delivery/openDelivery", miniprogram.APIImmediateDeliveryOpenDelivery)
    routerMiniProgram.GET("/delivery/preAddOrder", miniprogram.APIImmediateDeliveryPreAddOrder)
    routerMiniProgram.GET("/delivery/preCancelOrder", miniprogram.APIImmediateDeliveryPreCancelOrder)
    routerMiniProgram.GET("/delivery/realMockUpdateOrder", miniprogram.APIImmediateDeliveryRealMockUpdateOrder)
    routerMiniProgram.GET("/delivery/reOrder", miniprogram.APIImmediateDeliveryReOrder)

    // Handle the internet route
    routerMiniProgram.GET("/internet/getUserEncryptKey", miniprogram.APIInternetGetUserEncryptKey)

    // Handle the live broadcast route
    routerMiniProgram.GET("/liveBroadcast/addAssistant", miniprogram.APILiveAddAssistant)
    routerMiniProgram.GET("/liveBroadcast/addGoods", miniprogram.APILiveAddGoods)
    routerMiniProgram.GET("/liveBroadcast/addRole", miniprogram.APILiveAddRole)
    routerMiniProgram.POST("/liveBroadcast/addSubAnchor", miniprogram.APILiveAddSubAnchor)
    routerMiniProgram.POST("/liveBroadcast/createRoom", miniprogram.APILiveCreateRoom)
    routerMiniProgram.DELETE("/liveBroadcast/deleteRole", miniprogram.APILiveDeleteRole)
    routerMiniProgram.DELETE("/liveBroadcast/deleteRoom", miniprogram.APILiveDeleteRoom)
    routerMiniProgram.DELETE("/liveBroadcast/deleteSubAnchor", miniprogram.APILiveDeleteSubAnchor)
    routerMiniProgram.POST("/liveBroadcast/editRoom", miniprogram.APILiveEditRoom)
    routerMiniProgram.GET("/liveBroadcast/getAssistantList", miniprogram.APILiveGetAssistantList)
    routerMiniProgram.GET("/liveBroadcast/getFollowers", miniprogram.APILiveGetFollowers)
    routerMiniProgram.GET("/liveBroadcast/getLiveInfo", miniprogram.APILiveGetLiveInfo)
    routerMiniProgram.GET("/liveBroadcast/getPushUrl", miniprogram.APILiveGetPushUrl)
    routerMiniProgram.GET("/liveBroadcast/getRoleList", miniprogram.APILiveGetRoleList)
    routerMiniProgram.GET("/liveBroadcast/getSharedCode", miniprogram.APILiveGetSharedCode)
    routerMiniProgram.GET("/liveBroadcast/getSubAnchor", miniprogram.APILiveGetSubAnchor)
    routerMiniProgram.GET("/liveBroadcast/goodsAdd", miniprogram.APILiveGoodsAdd)
    routerMiniProgram.GET("/liveBroadcast/goodsAudit", miniprogram.APILiveGoodsAudit)
    routerMiniProgram.GET("/liveBroadcast/goodsDelete", miniprogram.APILiveGoodsDelete)
    routerMiniProgram.GET("/liveBroadcast/goodsInfo", miniprogram.APILiveGoodsInfo)
    routerMiniProgram.GET("/liveBroadcast/goodsList", miniprogram.APILiveGoodsList)
    routerMiniProgram.GET("/liveBroadcast/goodsPush", miniprogram.APILiveGoodsPush)
    routerMiniProgram.GET("/liveBroadcast/goodsResetAudit", miniprogram.APILiveGoodsResetAudit)
    routerMiniProgram.GET("/liveBroadcast/goodsSale", miniprogram.APILiveGoodsSale)
    routerMiniProgram.GET("/liveBroadcast/goodsSort", miniprogram.APILiveGoodsSort)
    routerMiniProgram.GET("/liveBroadcast/goodsUpdate", miniprogram.APILiveGoodsUpdate)
    routerMiniProgram.GET("/liveBroadcast/goodsVideo", miniprogram.APILiveGoodsVideo)
    routerMiniProgram.GET("/liveBroadcast/modifyAssistant", miniprogram.APILiveModifyAssistant)
    routerMiniProgram.GET("/liveBroadcast/modifySubAnchor", miniprogram.APILiveModifySubAnchor)
    routerMiniProgram.GET("/liveBroadcast/pushMessage", miniprogram.APILivePushMessage)
    routerMiniProgram.GET("/liveBroadcast/removeAssistant", miniprogram.APILiveRemoveAssistant)
    routerMiniProgram.GET("/liveBroadcast/updateComment", miniprogram.APILiveUpdateComment)
    routerMiniProgram.GET("/liveBroadcast/updateFeedPublic", miniprogram.APILiveUpdateFeedPublic)
    routerMiniProgram.GET("/liveBroadcast/updateKF", miniprogram.APILiveUpdateKF)
    routerMiniProgram.GET("/liveBroadcast/updateReplay", miniprogram.APILiveUpdateReplay)

    // Handle the express route
    routerMiniProgram.GET("/express/addOrder", miniprogram.APIExpressAddOrder)
    routerMiniProgram.GET("/express/batchGetOrder", miniprogram.APIExpressBatchGetOrder)
    routerMiniProgram.GET("/express/bindAccount", miniprogram.APIExpressBindAccount)
    routerMiniProgram.GET("/express/cancelOrder", miniprogram.APIExpressCancelOrder)
    routerMiniProgram.GET("/express/getAllAccount", miniprogram.APIExpressGetAllAccount)
    routerMiniProgram.GET("/express/getAllDelivery", miniprogram.APIExpressGetAllDelivery)
    routerMiniProgram.GET("/express/getOrder", miniprogram.APIExpressGetOrder)
    routerMiniProgram.GET("/express/getPath", miniprogram.APIExpressGetPath)
    routerMiniProgram.GET("/express/getPrinter", miniprogram.APIExpressGetPrinter)
    routerMiniProgram.GET("/express/getQuota", miniprogram.APIExpressGetQuota)
    routerMiniProgram.GET("/express/testUpdateOrder", miniprogram.APIExpressTestUpdateOrder)
    routerMiniProgram.GET("/express/updatePrinter", miniprogram.APIExpressUpdatePrinter)
    routerMiniProgram.GET("/express/getContact", miniprogram.APIExpressGetContact)
    routerMiniProgram.GET("/express/previewTemplate", miniprogram.APIExpressPreviewTemplate)
    routerMiniProgram.GET("/express/updateBusiness", miniprogram.APIExpressUpdateBusiness)
    routerMiniProgram.GET("/express/updatePath", miniprogram.APIExpressUpdatePath)

    // Handle the ocr route
    routerMiniProgram.GET("/ocr/bankcardByURL", miniprogram.APIOCRBankCardByURL)
    routerMiniProgram.GET("/ocr/bankcardByData", miniprogram.APIOCRBankCardByData)
    routerMiniProgram.GET("/ocr/businessLicenseByURL", miniprogram.APIOCRBusinessLicenseByURL)
    routerMiniProgram.GET("/ocr/businessLicenseByData", miniprogram.APIOCRBusinessLicenseByData)
    routerMiniProgram.GET("/ocr/driverLicenseByURL", miniprogram.APIOCRDriverLicenseByURL)
    routerMiniProgram.GET("/ocr/driverLicenseByData", miniprogram.APIOCRDriverLicenseByData)
    routerMiniProgram.GET("/ocr/idcardByURL", miniprogram.APIOCRIDCardByURL)
    routerMiniProgram.GET("/ocr/idcardByData", miniprogram.APIOCRIDCardByData)
    routerMiniProgram.GET("/ocr/printedTextByURL", miniprogram.APIOCRPrintedTextByURL)
    routerMiniProgram.GET("/ocr/printedTextByData", miniprogram.APIOCRPrintedTextByData)
    routerMiniProgram.GET("/ocr/vehicleLicenseByURL", miniprogram.APIOCRVehicleLicenseByURL)
    routerMiniProgram.GET("/ocr/vehicleLicenseByData", miniprogram.APIOCRVehicleLicenseByData)

    // Handle operation route
    routerMiniProgram.GET("/operation/getDomainInfo", miniprogram.APIOperationGetDomainInfo)
    routerMiniProgram.GET("/operation/getFeedback", miniprogram.APIOperationGetFeedback)
    routerMiniProgram.GET("/operation/getFeedbackmedia", miniprogram.APIOperationGetFeedbackMedia)
    routerMiniProgram.GET("/operation/getGrayReleasePlan", miniprogram.APIOperationGetGrayReleasePlan)
    routerMiniProgram.GET("/operation/getJsErrDetail", miniprogram.APIOperationGetJsErrDetail)
    routerMiniProgram.GET("/operation/getJsErrList", miniprogram.APIOperationGetJsErrList)
    routerMiniProgram.GET("/operation/getJsErrSearch", miniprogram.APIOperationGetJsErrSearch)
    routerMiniProgram.GET("/operation/getPerformance", miniprogram.APIOperationGetPerformance)
    routerMiniProgram.GET("/operation/getSceneList", miniprogram.APIOperationGetSceneList)
    routerMiniProgram.GET("/operation/getVersionList", miniprogram.APIOperationGetVersionList)
    routerMiniProgram.GET("/operation/realtimelogSearch", miniprogram.APIOperationRealtimelogSearch)

    // Handle risk control route
    routerMiniProgram.GET("/riskControl/getUserRiskRank", miniprogram.APIRiskControlGetUserRiskRank)

    // Handle search route
    routerMiniProgram.GET("/search/imageSearch", miniprogram.APISearchImageSearch)
    routerMiniProgram.GET("/search/siteSearch", miniprogram.APISearchSiteSearch)
    routerMiniProgram.GET("/search/submitPages", miniprogram.APISearchSubmitPages)

    // Handle service market route
    routerMiniProgram.GET("/serviceMarket/invokeService", miniprogram.APIServiceMarketInvokeService)

    // Handle short link route
    routerMiniProgram.GET("/shortLink/generate", miniprogram.APIShortLinkGenerate)

    // Handle soter route
    routerMiniProgram.GET("/soter/verifySignature", miniprogram.APISoterVerifySignature)

    // Handle subscribe message route
    routerMiniProgram.GET("/subscribeMessage/addTemplate", miniprogram.APISubscribeMessageAddTemplate)
    routerMiniProgram.GET("/subscribeMessage/deleteTemplate", miniprogram.APISubscribeMessageDeleteTemplate)
    routerMiniProgram.GET("/subscribeMessage/getCategory", miniprogram.APISubscribeMessageGetCategory)
    routerMiniProgram.GET("/subscribeMessage/getPubTemplateKeyWordsByID", miniprogram.APISubscribeMessageGetPubTemplateKeyWordsByID)
    routerMiniProgram.GET("/subscribeMessage/getPubTemplateTitleList", miniprogram.APISubscribeMessageGetPubTemplateTitleList)
    routerMiniProgram.GET("/subscribeMessage/getTemplateList", miniprogram.APISubscribeMessageGetTemplateList)
    routerMiniProgram.GET("/subscribeMessage/send", miniprogram.APISubscribeMessageSend)

  }

  // WeCom App Router
  wecomRouter := r.Group("/wecom")
  {

    // Handle user route
    wecomRouter.POST("/user/create", wecom.APIUserCreate)
    wecomRouter.GET("/user/get", wecom.APIUserGet)
    wecomRouter.PUT("/user/update", wecom.APIUserUpdate)
    wecomRouter.DELETE("/user/delete", wecom.APIUserDelete)
    wecomRouter.DELETE("/user/batch", wecom.APIUserBatchDelete)
    wecomRouter.GET("/users/simple", wecom.APIUserSimpleList)
    wecomRouter.GET("/users/detail", wecom.APIUserDetailList)
    wecomRouter.POST("/user/userIDToOpenID", wecom.APIUserUserIDToOpenID)
    wecomRouter.POST("/user/openIDToUserID", wecom.APIUserOpenIDToUserID)
    wecomRouter.GET("/user/authsucc", wecom.APIUserAuthAccept)
    wecomRouter.GET("/batch/invite", wecom.APIUserBatchInvite)
    wecomRouter.GET("/corp/qrcode", wecom.APIUserGetJoinQrCode)
    wecomRouter.GET("/user/getActiveStat", wecom.APIUserGetActiveStat)

    // Handle department route
    wecomRouter.POST("/department/create", wecom.APIDepartmentCreate)
    wecomRouter.PUT("/department/update", wecom.APIDepartmentUpdate)
    wecomRouter.DELETE("/department/delete", wecom.APIDepartmentDelete)
    wecomRouter.GET("/department/list", wecom.APIDepartmentList)

    // Handle tag route
    wecomRouter.POST("/tag/create", wecom.APITagCreate)
    wecomRouter.PUT("/tag/update", wecom.APITagUpdate)
    wecomRouter.DELETE("/tag/delete", wecom.APITagDelete)
    wecomRouter.GET("/tag/get", wecom.APITagUserGet)
    wecomRouter.POST("/tag/addTagUsers", wecom.APITagUserAdd)
    wecomRouter.DELETE("/tag/delTagUsers", wecom.APITagUserDel)
    wecomRouter.GET("/tag/list", wecom.APITagList)

    // Handle batch route
    wecomRouter.POST("/batch/syncUser", wecom.APIBatchSyncUser)
    wecomRouter.POST("/batch/replaceUser", wecom.APIBatchReplaceUser)
    wecomRouter.POST("/batch/replaceParty", wecom.APIBatchReplaceParty)
    wecomRouter.GET("/batch/getResult", wecom.APIBatchGetResult)

    // Handle linked corp route
    wecomRouter.POST("/linkedcorp/agent/getPermList", wecom.APILinkedCorpAgentGetPermList)
    wecomRouter.POST("/linkedcorp/user/get", wecom.APILinkedCorpUserGet)
    wecomRouter.POST("/linkedcorp/user/simplelist", wecom.APILinkedCorpUserSimpleList)
    wecomRouter.POST("/linkedcorp/user/list", wecom.APILinkedCorpUserList)
    wecomRouter.POST("/linkedcorp/department/list", wecom.APILinkedCorpDepartmentList)

    // Handle linked corp route
    wecomRouter.POST("export/simpleUser", wecom.APIExportSimpleUser)
    wecomRouter.POST("export/user", wecom.APIExportUser)
    wecomRouter.POST("export/department", wecom.APIExportDepartment)
    wecomRouter.POST("export/tagUser", wecom.APIExportTagUser)
    wecomRouter.GET("export/getResult", wecom.APIExportGetResult)

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
    wecomRouter.POST("externalContact/groupWelcomeTemplate", externalContact.APIExternalContactGroupWelcomeTemplate)
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
    wecomRouter.POST("kf/service_state/trans",accountService.APIAccountServiceStateTrans)

    // Handle account service sync message route
    wecomRouter.POST("kf/sync_msg", accountService.APIAccountServiceSyncMsg)
    wecomRouter.POST("kf/send_msg", accountService.APIAccountServiceSendMsg)
    wecomRouter.POST("kf/send_msg_on_event", accountService.APIAccountServiceSendMsgOnEvent)

    // Handle account service customer route
    wecomRouter.POST("kf/customer/batchget", accountService.APIAccountServiceCustomerBatchGet)
    wecomRouter.POST("kf/customer/get_upgrade_service_config", accountService.APIAccountServiceCustomerGetUpgradeServiceConfig)
    wecomRouter.POST("kf/customer/upgrade_service",accountService.APIAccountServiceCustomerUpgradeService)


    // Handle agent route
    wecomRouter.POST("agent/get", wecom.APIAgentGet)
    wecomRouter.POST("agent/list", wecom.APIAgentList)
    wecomRouter.POST("agent/set", wecom.APIAgentSet)

    // Handle menu route
    wecomRouter.POST("menu/create", wecom.APIAgentMenuCreate)
    wecomRouter.POST("menu/get", wecom.APIAgentMenuGet)
    wecomRouter.POST("menu/delete", wecom.APIAgentMenuDelete)

    // Handle agent workbench route
    wecomRouter.POST("agent/set_workbench_template", wecom.APIAgentSetWorkbenchTemplate)
    wecomRouter.POST("agent/get_workbench_template", wecom.APIAgentGetWorkbenchTemplate)
    wecomRouter.POST("agent/set_workbench_data", wecom.APIAgentSetWorkbenchData)





    // Handle message route
    wecomRouter.POST("/message/send", wecom.APISendTextMsg)
    wecomRouter.POST("/message/recall", wecom.APIRecallMsg)

    wecomRouter.POST("/")

  }

  log.Fatalln(r.Run(Host + ":" + Port))

}