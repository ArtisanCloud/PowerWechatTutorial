package external_contact

import (
	request2 "github.com/ArtisanCloud/PowerWeChat/v3/src/work/externalContact/customerStrategy/request"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/externalContact/request"
	"github.com/gin-gonic/gin"
	"net/http"
	"power-wechat-tutorial/services"
)

// 获取客户列表
// https://work.weixin.qq.com/api/doc/90000/90135/92113
func APIExternalContactList(c *gin.Context) {

	userID := c.DefaultQuery("userID", "matrix-x")

	res, err := services.WeComContactApp.ExternalContact.List(userID)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)
}

// 获取客户详情
// https://work.weixin.qq.com/api/doc/90000/90135/92114
func APIExternalContactGet(c *gin.Context) {

	userID := c.DefaultQuery("userID", "matrix-x")

	res, err := services.WeComContactApp.ExternalContact.Get(userID, "")

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)
}

// 批量获取客户详情.
// https://open.work.weixin.qq.com/api/doc/90000/90135/92994
func APIExternalContactBatchGetByUser(c *gin.Context) {
	userIDs := []string{c.DefaultQuery("userID", "matrix-x")}

	res, err := services.WeComContactApp.ExternalContact.BatchGet(userIDs, "", 100)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)
}

// 修改客户备注信息.
// https://work.weixin.qq.com/api/doc/90000/90135/92115
func APIExternalContactRemark(c *gin.Context) {

	options := &request.RequestExternalContactRemark{
		UserID:         c.DefaultQuery("userID", "matrix-x"),
		ExternalUserID: c.DefaultQuery("externalUserID", "woAJ2GCAAAd1asdasdjO4wKmE8Aabj9AAA"),
		Remark:         "备注信息",
		Description:    "描述信息",
		//RemarkCompany:  "腾讯科技",
		//RemarkMobiles: []string{
		//	"13800000001",
		//	"13800000002",
		//},
		//RemarkPicMediaID: "MEDIAID",
	}

	res, err := services.WeComContactApp.ExternalContact.Remark(options)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)
}

// 获取规则组列表
// https://work.weixin.qq.com/api/doc/90000/90135/94883
func APIExternalContactCustomerStrategyList(c *gin.Context) {

	cursor := c.DefaultQuery("cursor", "CURSOR")
	res, err := services.WeComContactApp.ExternalContactCustomerStrategy.List(cursor, 1000)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)
}

// 获取规则组详情
// https://work.weixin.qq.com/api/doc/90000/90135/94883
func APIExternalContactCustomerStrategyGet(c *gin.Context) {

	res, err := services.WeComContactApp.ExternalContactCustomerStrategy.Get(1)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)
}

// 获取规则组管理范围
// https://work.weixin.qq.com/api/doc/90000/90135/94883
func APIExternalContactCustomerStrategyGetRange(c *gin.Context) {
	cursor := c.DefaultQuery("cursor", "CURSOR")

	res, err := services.WeComContactApp.ExternalContactCustomerStrategy.GetRange(1, cursor, 1000)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)
}

// 创建新的规则组
// https://work.weixin.qq.com/api/doc/90000/90135/94883
func APIExternalContactCustomerStrategyCreate(c *gin.Context) {
	options := &request2.RequestCustomerStrategyCreate{
		ParentID:     0,
		StrategyName: "NAME",
		AdminList: []string{
			"zhangsan",
			"lisi",
		},
		Privilege: request2.RequestCustomerStrategyPrivilege{
			ViewCustomerList:        true,
			ViewCustomerData:        true,
			ViewRoomList:            true,
			ContactMe:               true,
			JoinRoom:                true,
			ShareCustomer:           true,
			OperResignCustomer:      true,
			SendCustomerMsg:         true,
			EditWelcomeMsg:          true,
			ViewBehaviorData:        true,
			ViewRoomData:            true,
			SendGroupMsg:            true,
			RoomDeduplication:       true,
			RapidReply:              true,
			OnjobCustomerTransfer:   true,
			EditAntiSpamRule:        true,
			ExportCustomerList:      true,
			ExportCustomerData:      true,
			ExportCustomerGroupList: true,
			ManageCustomerTag:       true,
		},
		Range: []request2.RequestCustomerStrategyRange{
			{
				Type:   1,
				UserID: "zhangsan",
			},
			{
				Type:    2,
				PartyID: 1,
			},
		},
	}

	res, err := services.WeComContactApp.ExternalContactCustomerStrategy.Create(options)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)
}

// 编辑规则组及其管理范围
// https://work.weixin.qq.com/api/doc/90000/90135/94883
func APIExternalContactCustomerStrategyEdit(c *gin.Context) {
	options := &request2.RequestCustomerStrategyEdit{
		StrategyID:   1,
		StrategyName: "NAME",
		AdminList: []string{
			"zhangsan",
			"lisi",
		},
		Privilege: request2.RequestCustomerStrategyPrivilege{
			ViewCustomerList:        true,
			ViewCustomerData:        true,
			ViewRoomList:            true,
			ContactMe:               true,
			JoinRoom:                true,
			ShareCustomer:           true,
			OperResignCustomer:      true,
			SendCustomerMsg:         true,
			EditWelcomeMsg:          true,
			ViewBehaviorData:        true,
			ViewRoomData:            true,
			SendGroupMsg:            true,
			RoomDeduplication:       true,
			RapidReply:              true,
			OnjobCustomerTransfer:   true,
			EditAntiSpamRule:        true,
			ExportCustomerList:      true,
			ExportCustomerData:      true,
			ExportCustomerGroupList: true,
			ManageCustomerTag:       true,
		},
		RangeAdd: []request2.RequestCustomerStrategyRange{
			{
				Type:   1,
				UserID: "zhangsan",
			},
			{
				Type:    2,
				PartyID: 1,
			},
		},
		RangeDel: []request2.RequestCustomerStrategyRange{
			{
				Type:   1,
				UserID: "lisi",
			},
			{
				Type:    2,
				PartyID: 2,
			},
		},
	}

	res, err := services.WeComContactApp.ExternalContactCustomerStrategy.Edit(options)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)
}

// 删除规则组
// https://work.weixin.qq.com/api/doc/90000/90135/94883
func APIExternalContactCustomerStrategyDel(c *gin.Context) {

	res, err := services.WeComContactApp.ExternalContactCustomerStrategy.Del(1)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)
}
