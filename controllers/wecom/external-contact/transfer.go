package external_contact

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/externalContact/transfer/request"
	"github.com/gin-gonic/gin"
	"net/http"
	"power-wechat-tutorial/services"
)

// 分配在职成员的客户
// https://work.weixin.qq.com/api/doc/90000/90135/92125
func APIExternalContactTransferCustomer(c *gin.Context) {
	options := &request.RequestTransferCustomer{
		HandoverUserID:     c.DefaultQuery("handoverUserID", "walle"),
		TakeoverUserID:     c.DefaultQuery("takeoverUserID", "matrix-x"),
		ExternalUserID:     []string{c.DefaultQuery("externalUserID", "woAJ2GCAAAXtWyujaWJHDDGi0mACAAAA")},
		TransferSuccessMsg: "您好，您的服务已升级，后续将由我的同事李四@腾讯接替我的工作，继续为您服务。",
	}

	res, err := services.WeComContactApp.ExternalContactTransfer.TransferCustomer(c.Request.Context(), options)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)
}

// 查询客户接替状态
// https://work.weixin.qq.com/api/doc/90000/90135/94088
func APIExternalContactTransferResult(c *gin.Context) {

	options := &request.RequestTransferResult{
		HandoverUserID: "walle",
		TakeoverUserID: "matrix-x",
		Cursor:         c.DefaultQuery("cursor", "CURSOR"),
	}
	res, err := services.WeComContactApp.ExternalContactTransfer.TransferResult(c.Request.Context(), options)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)
}

// 获取待分配的离职成员列表
// https://work.weixin.qq.com/api/doc/90000/90135/92125
func APIExternalContactGetUnassignedList(c *gin.Context) {

	pageID := 1
	pageSize := 1000
	cursor := ""

	res, err := services.WeComContactApp.ExternalContactTransfer.GetUnassignedList(c.Request.Context(), pageID, cursor, pageSize)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)
}

// 分配离职成员的客户
// https://work.weixin.qq.com/api/doc/90000/90135/94081
func APIExternalContactResignedTransferCustomer(c *gin.Context) {

	options := &request.RequestResignedTransferCustomer{
		HandoverUserID: "walle",
		TakeoverUserID: "matrix-x",
		ExternalUserID: []string{"woAJ2GCAAAXtWyujaWJHDDGi0mACBBBB"},
	}

	res, err := services.WeComContactApp.ExternalContactTransfer.ResignedTransferCustomer(c.Request.Context(), options)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)
}

// 查询客户接替状态
// https://work.weixin.qq.com/api/doc/90000/90135/94082
func APIExternalContactResignedTransferResult(c *gin.Context) {
	options := &request.RequestResignedTransferResult{
		HandoverUserID: "walle",
		TakeoverUserID: "matrix-x",
		Cursor:         c.DefaultQuery("cursor", "CURSOR"),
	}

	res, err := services.WeComContactApp.ExternalContactTransfer.ResignedTransferResult(c.Request.Context(), options)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)
}

// 分配离职成员的客户群
// https://work.weixin.qq.com/api/doc/90000/90135/92127
func APIExternalContactGroupChatTransfer(c *gin.Context) {

	options := &request.RequestGroupChatTransfer{
		ChatIDList: []string{"wrOgQhDgAAcwMTB7YmDkbeBsgT_AAAA", "wrOgQhDgAAMYQiS5ol9G7gK9JVQUAAAA"},
		NewOwner:   "walle",
	}

	res, err := services.WeComContactApp.ExternalContactTransfer.GroupChatTransfer(c.Request.Context(), options)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)
}
