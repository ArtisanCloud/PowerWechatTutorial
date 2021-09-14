package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"power-wechat-tutorial/controllers/miniprogram"
	"power-wechat-tutorial/controllers/payment"
	"power-wechat-tutorial/services"
)

var Host string = ""
var Port string = "8888"

func main() {

	var err error
	services.PaymentService, err = services.NewWXPaymentService(nil)
	if err != nil || services.PaymentService == nil {
		panic(err)
	}

	services.AppMiniProgram, err = services.NewMiniMiniProgramService()
	if err != nil || services.AppMiniProgram == nil {
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
		routerMiniProgram.GET("/auth/checkEncryptedData", miniprogram.APICheckEncryptedData)
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
<<<<<<< HEAD

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

=======

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

>>>>>>> 457d0eaa4eabb26487f7aceb91b06832746534c3
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
		routerMiniProgram.GET("/liveBroadcast/addSubAnchor", miniprogram.APILiveAddSubAnchor)
		routerMiniProgram.GET("/liveBroadcast/createRoom", miniprogram.APILiveCreateRoom)
		routerMiniProgram.GET("/liveBroadcast/deleteRole", miniprogram.APILiveDeleteRole)
		routerMiniProgram.GET("/liveBroadcast/deleteRoom", miniprogram.APILiveDeleteRoom)
		routerMiniProgram.GET("/liveBroadcast/deleteSubAnchor", miniprogram.APILiveDeleteSubAnchor)
		routerMiniProgram.GET("/liveBroadcast/editRoom", miniprogram.APILiveEditRoom)
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

	log.Fatalln(r.Run(Host + ":" + Port))

}
