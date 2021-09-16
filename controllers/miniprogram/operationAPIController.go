package miniprogram

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/power"
	"github.com/gin-gonic/gin"
	"github.com/golang-module/carbon"
	"io/ioutil"
	"net/http"
	"power-wechat-tutorial/services"
)

// 查询域名配置
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/operation/operation.getDomainInfo.html
func APIOperationGetDomainInfo(c *gin.Context) {
	rs, err := services.MiniprogramApp.Operation.GetDomainInfo("")

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, rs)
}

// 获取用户反馈列表
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/operation/operation.getFeedback.html
func APIOperationGetFeedback(c *gin.Context) {
	rs, err := services.MiniprogramApp.Operation.GetFeedback(1, 1, 3)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, rs)
}

// 获取 mediaId 图片
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/operation/operation.getFeedbackmedia.html
func APIOperationGetFeedbackMedia(c *gin.Context) {

	mediaID, exist := c.GetQuery("mediaID")
	if !exist {
		panic("parameter media id expected")
	}

	rs, err := services.MiniprogramApp.Operation.GetFeedbackMedia(1, mediaID)

	if err != nil {
		panic(err)
	}

	content, _ := ioutil.ReadAll(rs.Body)
	//fmt.Dump("content-type:",rs.Header.Get("Content-Type"))
	c.Header("Content-Type", rs.Header.Get("Content-Type"))
	c.Header("Content-Disposition", rs.Header.Get("attachment;filename=\""+rs.Header.Get("filename")+"\""))
	c.Data(http.StatusOK, rs.Header.Get("Content-Type"), content)

}

// 查询当前分阶段发布详情
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/operation/operation.getGrayReleasePlan.html
func APIOperationGetGrayReleasePlan(c *gin.Context) {
	rs, err := services.MiniprogramApp.Operation.GetGrayReleasePlan()

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, rs)
}

// 错误查询详情
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/operation/operation.getJsErrDetail.html
func APIOperationGetJsErrDetail(c *gin.Context) {

	options := &power.HashMap{
		"startTime":     "2021-01-25",
		"endTime":       "2021-01-26",
		"errorMsgMd5":   "f2fb4f8cd638466ad0e7607b01b7d0ca",
		"errorStackMd5": "795a63b70ce5755c7103611d93077603",
		"appVersion":    "0",
		"sdkVersion":    "0",
		"osName":        "2",
		"clientVersion": "0",
		"openid":        "",
		"offset":        0,
		"limit":         10,
		"desc":          "0",
	}

	rs, err := services.MiniprogramApp.Operation.GetJsErrDetail(options)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, rs)
}

// 错误查询列表
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/operation/operation.getJsErrList.html
func APIOperationGetJsErrList(c *gin.Context) {

	options := &power.HashMap{
		"startTime":  "2021-01-25",
		"endTime":    "2021-01-26",
		"errType":    "0",
		"appVersion": "0",
		"openid":     "",
		"keyword":    "",
		"orderby":    "uv",
		"desc":       "2",
		"offset":     0,
		"limit":      1,
	}

	rs, err := services.MiniprogramApp.Operation.GetJsErrList(options)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, rs)
}

// 错误查询, 接口即将废弃
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/operation/operation.getJsErrSearch.html
func APIOperationGetJsErrSearch(c *gin.Context) {

	options := &power.HashMap{
		"errmsg_keyword": "",
		"type":           1,
		"client_version": "",
		"start_time":     1587021734,
		"end_time":       1587626534,
		"start":          1,
		"limit":          1,
		"sceneDesc":      "测试数据",
	}

	rs, err := services.MiniprogramApp.Operation.GetJsErrSearch(options)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, rs)
}

// 性能监控
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/operation/operation.getPerformance.html
func APIOperationGetPerformance(c *gin.Context) {

	options := &power.HashMap{
		"cost_time_type":     2,
		"default_start_time": 1572339403,
		"default_end_time":   1574931403,
		"device":             "@_all",
		"networktype":        "@_all",
		"scene":              "@_all",
		"is_download_code":   "@_all",
	}

	rs, err := services.MiniprogramApp.Operation.GetPerformance(options)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, rs)
}

// 获取访问来源
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/operation/operation.getSceneList.html
func APIOperationGetSceneList(c *gin.Context) {
	rs, err := services.MiniprogramApp.Operation.GetSceneList()

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, rs)
}

// 获取客户端版本
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/operation/operation.getVersionList.html
func APIOperationGetVersionList(c *gin.Context) {
	rs, err := services.MiniprogramApp.Operation.GetVersionList()

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, rs)
}

// 获取客户端版本
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/operation/operation.getVersionList.html
func APIOperationRealtimelogSearch(c *gin.Context) {

	dateFormat := "Ymd"
	now := carbon.Now().AddWeeks(-2)
	from := now.SetWeekStartsAt(carbon.Monday).StartOfWeek().ToFormatString(dateFormat, services.TIMEZONE)
	to := now.SetWeekStartsAt(carbon.Monday).EndOfWeek().ToFormatString(dateFormat, services.TIMEZONE)

	options := &power.HashMap{
		"date":      "YYYYMMDD",
		"begintime": from,
		"endtime":   to,
	}

	rs, err := services.MiniprogramApp.Operation.RealTimeLogSearch(options)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, rs)
}
