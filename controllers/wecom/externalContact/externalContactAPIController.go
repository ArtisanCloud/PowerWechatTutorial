package externalContact


import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/power"
	request2 "github.com/ArtisanCloud/power-wechat/src/work/externalContact/customerStrategy/request"
	"github.com/gin-gonic/gin"
	"net/http"
	"power-wechat-tutorial/services"
)


// 获取外部联系人列表.
// https://work.weixin.qq.com/api/doc/90000/90135/92113
func APIExternalContactList(c *gin.Context) {

	userID := c.DefaultQuery("userID", "matrix-x")

	res, err := services.WeComContactApp.ExternalContact.List(userID)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)
}

// 获取外部联系人详情.
// https://work.weixin.qq.com/api/doc/90000/90135/92114
func APIExternalContactGet(c *gin.Context) {

	userID := c.DefaultQuery("userID", "matrix-x")

	res, err := services.WeComContactApp.ExternalContact.Get(userID, 1)

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

	options := &power.HashMap{
		"userid":          c.DefaultQuery("userID", "matrix-x"),
		"external_userid": c.DefaultQuery("externalUserID", "woAJ2GCAAAd1asdasdjO4wKmE8Aabj9AAA"),
		"remark":          "备注信息",
		"description":     "描述信息",
		"remark_company":  "腾讯科技",
		"remark_mobiles": []string{
			"13800000001",
			"13800000002",
		},
		"remark_pic_mediaid": "MEDIAID",
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

		Privilege: &power.HashMap{
			"view_customer_list":         true,
			"view_customer_data":         true,
			"view_room_list":             true,
			"contact_me":                 true,
			"join_room":                  true,
			"share_customer":             false,
			"oper_resign_customer":       true,
			"send_customer_msg":          true,
			"edit_welcome_msg":           true,
			"view_behavior_data":         true,
			"view_room_data":             true,
			"send_group_msg":             true,
			"room_deduplication":         true,
			"rapid_reply":                true,
			"onjob_customer_transfer":    true,
			"edit_anti_spam_rule":        true,
			"export_customer_list":       true,
			"export_customer_data":       true,
			"export_customer_group_list": true,
			"manage_customer_tag":        true,
		},
		Range: []*power.HashMap{

			&power.HashMap{
				"type":   1,
				"userid": "zhangsan",
			},
			&power.HashMap{
				"type":    2,
				"partyid": 1,
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
		Privilege: &power.HashMap{
			"view_customer_list":         true,
			"view_customer_data":         true,
			"view_room_list":             true,
			"contact_me":                 true,
			"join_room":                  true,
			"share_customer":             false,
			"oper_resign_customer":       true,
			"oper_resign_group":          true,
			"send_customer_msg":          true,
			"edit_welcome_msg":           true,
			"view_behavior_data":         true,
			"view_room_data":             true,
			"send_group_msg":             true,
			"room_deduplication":         true,
			"rapid_reply":                true,
			"onjob_customer_transfer":    true,
			"edit_anti_spam_rule":        true,
			"export_customer_list":       true,
			"export_customer_data":       true,
			"export_customer_group_list": true,
			"manage_customer_tag":        true,
		},
		RangeAdd: []*power.HashMap{

			&power.HashMap{
				"type":   1,
				"userid": "zhangsan",
			},
			&power.HashMap{
				"type":    2,
				"partyid": 1,
			},
		},
		RangeDel: []*power.HashMap{

			&power.HashMap{
				"type":   1,
				"userid": "lisi",
			},
			&power.HashMap{
				"type":    2,
				"partyid": 2,
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