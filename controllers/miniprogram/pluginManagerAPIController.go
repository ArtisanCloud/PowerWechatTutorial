package miniprogram

import (
	"github.com/gin-gonic/gin"
	"power-wechat-tutorial/services"
)

// 向插件开发者发起使用插件的申请
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/plugin-management/pluginManager.applyPlugin.html
func APIPluginManagerApplyPlugin(c *gin.Context) {

	pluginAppID, exist := c.GetQuery("pluginAppID")
	if !exist {
		panic("parameter plugin app id expected")
	}

	rs, err := services.AppMiniProgram.Plugin.ApplyPlugin(pluginAppID, "test reason")

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, rs)
}

// 获取当前所有插件使用方（供插件开发者调用）
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/plugin-management/pluginManager.getPluginDevApplyList.html
func APIPluginManagerGetPluginDevApplyList(c *gin.Context) {

	rs, err := services.AppMiniProgram.Plugin.GetPluginDevApplyList("dev_apply_list", 1, 1)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, rs)
}

// 查询已添加的插件
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/plugin-management/pluginManager.getPluginList.html
func APIPluginManagerGetPluginList(c *gin.Context) {

	rs, err := services.AppMiniProgram.Plugin.GetPluginList()

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

	rs, err := services.AppMiniProgram.Plugin.SetDevPluginApplyStatus("dev_agree", pluginAppID, "test reason")

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

	rs, err := services.AppMiniProgram.Plugin.UnbindPlugin(pluginAppID)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, rs)
}
