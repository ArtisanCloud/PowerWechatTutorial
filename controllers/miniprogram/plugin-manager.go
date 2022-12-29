package miniprogram

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"power-wechat-tutorial/services"
)

// 向插件开发者发起使用插件的申请
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/plugin-management/pluginManager.applyPlugin.html
func APIPluginManagerApplyPlugin(c *gin.Context) {

	pluginAppID, exist := c.GetQuery("pluginAppID")
	if !exist {
		panic("parameter plugin app id expected")
	}

	rs, err := services.MiniProgramApp.Plugin.ApplyPlugin(c.Request.Context(), pluginAppID, "test reason")

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, rs)
}

// 获取当前所有插件使用方（供插件开发者调用）
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/plugin-management/pluginManager.getPluginDevApplyList.html
func APIPluginManagerGetPluginDevApplyList(c *gin.Context) {

	rs, err := services.MiniProgramApp.Plugin.GetPluginDevApplyList(c.Request.Context(), "dev_apply_list", 1, 1)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, rs)
}

// 查询已添加的插件
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/plugin-management/pluginManager.getPluginList.html
func APIPluginManagerGetPluginList(c *gin.Context) {

	rs, err := services.MiniProgramApp.Plugin.GetPluginList(c.Request.Context())

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, rs)
}

// 修改插件使用申请的状态（供插件开发者调用）
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/plugin-management/pluginManager.setDevPluginApplyStatus.html
func APIPluginManagerSetDevPluginApplyStatus(c *gin.Context) {

	pluginAppID, exist := c.GetQuery("pluginAppID")
	if !exist {
		panic("parameter plugin app id expected")
	}

	rs, err := services.MiniProgramApp.Plugin.SetDevPluginApplyStatus(c.Request.Context(), "dev_agree", pluginAppID, "test reason")

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, rs)
}

// 删除已添加的插件
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/plugin-management/pluginManager.unbindPlugin.html
func APIPluginManagerUnbindPlugin(c *gin.Context) {

	pluginAppID, exist := c.GetQuery("pluginAppID")
	if !exist {
		panic("parameter plugin app id expected")
	}

	rs, err := services.MiniProgramApp.Plugin.UnbindPlugin(c.Request.Context(), pluginAppID)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, rs)
}
