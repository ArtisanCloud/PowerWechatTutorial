package external_contact

import (
	"net/http"
	"power-wechat-tutorial/services"
	"strconv"

	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/externalContact/moment/request"
	request2 "github.com/ArtisanCloud/PowerWeChat/v3/src/work/externalContact/momentStrategy/request"
	"github.com/gin-gonic/gin"
)

// 获取企业全部的发表列表
// https://work.weixin.qq.com/api/doc/90000/90135/93333
func APIExternalContactGetMomentList(c *gin.Context) {

	options := &request.RequestGetMomentList{
		StartTime:  1605000000,
		EndTime:    1605172726,
		Creator:    "walle",
		FilterType: 1,
		Cursor:     c.DefaultQuery("cursor", "CURSOR"),
		Limit:      10,
	}

	res, err := services.WeComContactApp.ExternalContactMoment.GetMomentList(options)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)
}

// 获取客户朋友圈企业发表的列表
// https://work.weixin.qq.com/api/doc/90000/90135/93333
func APIExternalContactGetMomentTask(c *gin.Context) {

	momentID := c.DefaultQuery("momentID", "wrOgQhDgAAMYQiS5ol9G7gK9JVAAAA")
	cursor := c.DefaultQuery("cursor", "CURSOR")
	limit := 100

	res, err := services.WeComContactApp.ExternalContactMoment.GetMomentTask(momentID, cursor, limit)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)

}

// 获取客户朋友圈发表时选择的可见范围
// https://work.weixin.qq.com/api/doc/90000/90135/93333
func APIExternalContactGetMomentCustomerList(c *gin.Context) {

	momentID := c.DefaultQuery("momentID", "wrOgQhDgAAMYQiS5ol9G7gK9JVAAAA")
	userID := c.DefaultQuery("userID", "matrix-x")
	cursor := c.DefaultQuery("cursor", "CURSOR")
	limit := 100

	res, err := services.WeComContactApp.ExternalContactMoment.GetMomentCustomerList(momentID, userID, cursor, limit)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)

}

// 获取客户朋友圈发表时选择的可见范围
// https://work.weixin.qq.com/api/doc/90000/90135/93333
func APIExternalContactGetMomentSendResult(c *gin.Context) {

	momentID := c.DefaultQuery("momentID", "wrOgQhDgAAMYQiS5ol9G7gK9JVAAAA")
	userID := c.DefaultQuery("userID", "matrix-x")
	cursor := c.DefaultQuery("cursor", "CURSOR")
	limit := 100

	res, err := services.WeComContactApp.ExternalContactMoment.GetMomentCustomerList(momentID, userID, cursor, limit)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)
}

// 获取客户朋友圈的互动数据
// https://work.weixin.qq.com/api/doc/90000/90135/93333
func APIExternalContactGetMomentComments(c *gin.Context) {

	momentID := c.DefaultQuery("momentID", "wrOgQhDgAAMYQiS5ol9G7gK9JVAAAA")
	userID := c.DefaultQuery("userID", "matrix-x")

	res, err := services.WeComContactApp.ExternalContactMoment.GetMomentComments(momentID, userID)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)
}

// 获取规则组列表
// https://work.weixin.qq.com/api/doc/90000/90135/94890
func APIExternalContactMomentStrategyList(c *gin.Context) {
	cursor := c.DefaultQuery("cursor", "CURSOR")
	limit := 100

	res, err := services.WeComContactApp.ExternalContactMomentStrategy.List(cursor, limit)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)
}

// 获取规则组详情
// https://work.weixin.qq.com/api/doc/90000/90135/94890
func APIExternalContactMomentStrategyGet(c *gin.Context) {
	strategyId := c.DefaultQuery("strategyID", "0")
	strategyID, _ := strconv.Atoi(strategyId)
	res, err := services.WeComContactApp.ExternalContactMomentStrategy.Get(strategyID)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)
}

// 获取规则组管理范围
// https://work.weixin.qq.com/api/doc/90000/90135/94890
func APIExternalContactMomentStrategyGetRange(c *gin.Context) {
	strategyId := c.DefaultQuery("strategyID", "0")
	strategyID, _ := strconv.Atoi(strategyId)
	cursor := c.DefaultQuery("cursor", "CURSOR")
	limit := 100

	res, err := services.WeComContactApp.ExternalContactMomentStrategy.GetRange(strategyID, cursor, limit)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)
}

// 创建新的规则组
// https://work.weixin.qq.com/api/doc/90000/90135/94890
func APIExternalContactMomentStrategyCreate(c *gin.Context) {

	options := &request2.RequestMomentStrategyCreate{
		ParentID:     0,
		StrategyName: "NAME",
		AdminList: []string{
			"zhangsan",
			"lisi",
		},
		Privilege: &power.HashMap{
			"send_moment":                  true,
			"view_moment_list":             true,
			"manage_moment_cover_and_sign": true,
		},
		Range: []*power.HashMap{
			{
				"type":   1,
				"userid": "zhangsan",
			},
			{
				"type":    2,
				"partyid": 1,
			},
		},
	}

	res, err := services.WeComContactApp.ExternalContactMomentStrategy.Create(options)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)
}

func APIExternalContactMomentStrategyEdit(c *gin.Context) {

	options := &request2.RequestMomentStrategyEdit{
		StrategyID:   1,
		StrategyName: "NAME",
		AdminList: []string{
			"zhangsan",
			"lisi",
		},
		Privilege: &power.HashMap{
			"send_moment":                  true,
			"view_moment_list":             true,
			"manage_moment_cover_and_sign": true,
		},
		RangeAdd: []*power.HashMap{
			{
				"type":   1,
				"userid": "zhangsan",
			},
			{
				"type":    2,
				"partyid": 1,
			},
		},
		RangeDel: []*power.HashMap{
			{
				"type":   1,
				"userid": "lisi",
			},
			{
				"type":    2,
				"partyid": 2,
			},
		},
	}

	res, err := services.WeComContactApp.ExternalContactMomentStrategy.Edit(options)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)
}

func APIExternalContactMomentStrategyDel(c *gin.Context) {

	strategyId := c.DefaultQuery("strategyID", "0")
	strategyID, _ := strconv.Atoi(strategyId)

	res, err := services.WeComContactApp.ExternalContactMomentStrategy.Del(strategyID)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)
}
