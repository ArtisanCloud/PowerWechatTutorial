package externalContact

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/power"
	"github.com/ArtisanCloud/power-wechat/src/work/externalContact/contactWay/request"
	request2 "github.com/ArtisanCloud/power-wechat/src/work/externalContact/customerStrategy/request"
	"github.com/gin-gonic/gin"
	"net/http"
	"power-wechat-tutorial/services"
)

// 获取配置了客户联系功能的成员列表.
// https://work.weixin.qq.com/api/doc/90000/90135/92571
func APIExternalContactGetFollowUserList(c *gin.Context) {

	res, err := services.WeComContactApp.ExternalContact.GetFollowUsers()

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)

}

// 配置客户联系「联系我」方式.
// https://open.work.weixin.qq.com/api/doc/90000/90135/92572
func APIExternalContactAddContactWay(c *gin.Context) {

	options := &request.RequestAddContactWay{
		Type:          1,
		Scene:         1,
		Style:         1,
		Remark:        "渠道客户",
		SkipVerify:    true,
		State:         "teststate",
		User:          []string{"zhangsan", "lisi", "wangwu"},
		Party:         []int{2, 3},
		IsTemp:        true,
		ExpiresIn:     86400,
		ChatExpiresIn: 86400,
		UnionID:       "oxTWIuGaIt6gTKsQRLau2M0AAAA",
		Conclusions: &power.HashMap{
			"text": power.HashMap{
				"content": "文本消息内容",
			},
			"image": power.HashMap{
				"media_id": "MEDIA_ID",
			},
			"link": power.HashMap{
				"title":  "消息标题",
				"picurl": "https://example.pic.com/path",
				"desc":   "消息描述",
				"url":    "https://example.link.com/path",
			},
			"miniprogram": power.HashMap{
				"title":        "消息标题",
				"pic_media_id": "MEDIA_ID",
				"appid":        "wx8bd80126147dfAAA",
				"page":         "/path/index.html",
			},
		},
	}

	res, err := services.WeComContactApp.ExternalContactContactWay.Add(options)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)
}

// 获取企业已配置的「联系我」方式
// https://work.weixin.qq.com/api/doc/90000/90135/92572
func APIExternalContactGetContactWay(c *gin.Context) {
	configID := c.DefaultQuery("configID", "42b34949e138eb6e027c123cba77fad7")

	res, err := services.WeComContactApp.ExternalContactContactWay.Get(configID)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)
}

// 获取企业已配置的「联系我」列表
// https://work.weixin.qq.com/api/doc/90000/90135/92572
func APIExternalContactListContactWay(c *gin.Context) {
	options := &request.RequestListContactWay{
		1622476800,
		1625068800,
		"CURSOR",
		1000,
	}
	res, err := services.WeComContactApp.ExternalContactContactWay.List(options)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)
}

// 获取企业已配置的「联系我」列表
// https://work.weixin.qq.com/api/doc/90000/90135/92572
func APIExternalContactUpdateContactWay(c *gin.Context) {

	configID := c.DefaultQuery("configID", "42b34949e138eb6e027c123cba77fAAA")
	options := &power.HashMap{
		"config_id":       configID,
		"remark":          "渠道客户",
		"skip_verify":     true,
		"style":           1,
		"state":           "teststate",
		"user":            []string{"zhangsan", "lisi", "wangwu"},
		"party":           []int{2, 3},
		"expires_in":      86400,
		"chat_expires_in": 86400,
		"unionid":         "oxTWIuGaIt6gTKsQRLau2M0AAAA",
		"conclusions": power.HashMap{
			"text": power.StringMap{
				"content": "文本消息内容",
			},
			"image": power.StringMap{
				"media_id": "MEDIA_ID",
			},
			"link": power.StringMap{
				"title":  "消息标题",
				"picurl": "https://example.pic.com/path",
				"desc":   "消息描述",
				"url":    "https://example.link.com/path",
			},
			"miniprogram": power.StringMap{
				"title":        "消息标题",
				"pic_media_id": "MEDIA_ID",
				"appid":        "wx8bd80126147dfAAA",
				"page":         "/path/index",
			},
		},
	}

	res, err := services.WeComContactApp.ExternalContactContactWay.Update(configID, options)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)
}

func APIExternalContactDelContactWay(c *gin.Context) {

	configID := c.DefaultQuery("configID", "42b34949e138eb6e027c123cba77fad7")

	res, err := services.WeComContactApp.ExternalContactContactWay.Delete(configID)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)
}

func APIExternalContactCloseTempChat(c *gin.Context) {

	userID := c.DefaultQuery("userID", "matrix-x")
	externalUserID := c.DefaultQuery("externalUserID", "matrix-x")

	res, err := services.WeComContactApp.ExternalContactContactWay.CloseTempChat(userID, externalUserID)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)
}

func APIExternalContactList(c *gin.Context) {

	userID := c.DefaultQuery("userID", "matrix-x")

	res, err := services.WeComContactApp.ExternalContact.List(userID)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)
}

func APIExternalContactGet(c *gin.Context) {

	userID := c.DefaultQuery("userID", "matrix-x")

	res, err := services.WeComContactApp.ExternalContact.Get(userID, 1)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)
}

func APIExternalContactBatchGetByUser(c *gin.Context) {
	userIDs := []string{c.DefaultQuery("userID", "matrix-x")}

	res, err := services.WeComContactApp.ExternalContact.BatchGet(userIDs, "", 100)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)
}

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

func APIExternalContactCustomerStrategyList(c *gin.Context) {

	cursor := c.DefaultQuery("cursor", "CURSOR")
	res, err := services.WeComContactApp.ExternalContactCustomerStrategy.List(cursor, 1000)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)
}

func APIExternalContactCustomerStrategyGet(c *gin.Context) {

	res, err := services.WeComContactApp.ExternalContactCustomerStrategy.Get(1)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)
}

func APIExternalContactCustomerStrategyGetRange(c *gin.Context) {
	cursor := c.DefaultQuery("cursor", "CURSOR")

	res, err := services.WeComContactApp.ExternalContactCustomerStrategy.GetRange(1, cursor, 1000)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)
}

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

func APIExternalContactCustomerStrategyDel(c *gin.Context) {

	res, err := services.WeComContactApp.ExternalContactCustomerStrategy.Del(1)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)
}
