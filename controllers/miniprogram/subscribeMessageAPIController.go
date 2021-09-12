package miniprogram

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/power"
	"github.com/gin-gonic/gin"
	"power-wechat-tutorial/services"
)

// 组合模板并添加至帐号下的个人模板库
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/subscribe-message/subscribeMessage.addTemplate.html
func APISubscribeMessageAddTemplate(c *gin.Context) {
	tID, exist := c.GetQuery("tID")
	if !exist {
		panic("parameter tid expected")
	}

	rs, err := services.AppMiniProgram.SubscribeMessage.AddTemplate(tID, []int{1, 2}, "测试数据")

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, rs)
}

// 删除帐号下的个人模板
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/subscribe-message/subscribeMessage.deleteTemplate.html
func APISubscribeMessageDeleteTemplate(c *gin.Context) {
	priTmplID, exist := c.GetQuery("priTmplID")
	if !exist {
		panic("parameter tid expected")
	}

	rs, err := services.AppMiniProgram.SubscribeMessage.DeleteTemplate(priTmplID)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, rs)
}

// 获取小程序账号的类目
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/subscribe-message/subscribeMessage.getCategory.html
func APISubscribeMessageGetCategory(c *gin.Context) {
	rs, err := services.AppMiniProgram.SubscribeMessage.GetCategory()

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, rs)
}

// 获取模板标题下的关键词列表
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/subscribe-message/subscribeMessage.getPubTemplateKeyWordsById.html
func APISubscribeMessageGetPubTemplateKeyWordsByID(c *gin.Context) {
	tID, exist := c.GetQuery("tID")
	if !exist {
		panic("parameter tid expected")
	}

	rs, err := services.AppMiniProgram.SubscribeMessage.GetPubTemplateKeyWordsByID(tID)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, rs)
}

// 获取帐号所属类目下的公共模板标题
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/subscribe-message/subscribeMessage.getPubTemplateTitleList.html
func APISubscribeMessageGetPubTemplateTitleList(c *gin.Context) {
	rs, err := services.AppMiniProgram.SubscribeMessage.GetPubTemplateTitleList("2,616", 0, 5)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, rs)
}

// 获取当前帐号下的个人模板列表
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/subscribe-message/subscribeMessage.getTemplateList.html
func APISubscribeMessageGetTemplateList(c *gin.Context) {
	rs, err := services.AppMiniProgram.SubscribeMessage.GetTemplateList()

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, rs)
}

// 发送订阅消息
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/subscribe-message/subscribeMessage.send.html
func APISubscribeMessageSend(c *gin.Context) {

	toUser, exist := c.GetQuery("openID")
	if !exist {
		panic("parameter open id expected")
	}

	templateID, exist := c.GetQuery("templateID")
	if !exist {
		panic("parameter template id expected")
	}

	page := "index"
	miniprogramState := "developer"
	lang := "zh_CN"

	data := &power.HashMap{
		"number01": power.StringMap{
			"value": "339208499",
		},
		"date01": power.StringMap{
			"value": "2015年01月05日",
		},
		"site01": power.StringMap{
			"value": "TIT创意园",
		},
		"site02": power.StringMap{
			"value": "广州市新港中路397号",
		},
	}

	rs, err := services.AppMiniProgram.SubscribeMessage.Send(toUser, templateID, page, miniprogramState, lang, data)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, rs)
}
