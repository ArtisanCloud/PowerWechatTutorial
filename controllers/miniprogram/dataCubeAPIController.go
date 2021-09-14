package miniprogram

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/power"
	"github.com/gin-gonic/gin"
	"github.com/golang-module/carbon"
	"power-wechat-tutorial/services"
)

// 获取用户访问小程序日留存
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/data-analysis/visit-retain/analysis.getDailyRetain.html
func APIGetDailyRetain(c *gin.Context) {

	now := carbon.Now().AddDays(-5)
	from := now.ToFormatString("Ymd", services.TIMEZONE)
	to := now.ToFormatString("Ymd", services.TIMEZONE)

	rs, err := services.MiniProgramApp.DataCube.DailyRetainInfo(from, to)

	if err != nil {
		panic(err)
	}

	c.JSON(200, rs)

}

// 获取用户访问小程序月留存
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/data-analysis/visit-retain/analysis.getMonthlyRetain.html
func APIGetMonthlyRetain(c *gin.Context) {

	now := carbon.Now().AddMonths(-1)
	from := now.StartOfMonth().ToFormatString("Ymd", services.TIMEZONE)
	to := now.EndOfMonth().ToFormatString("Ymd", services.TIMEZONE)

	rs, err := services.MiniProgramApp.DataCube.MonthlyRetainInfo(from, to)

	if err != nil {
		panic(err)
	}

	c.JSON(200, rs)

}

// 获取用户访问小程序周留存
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/data-analysis/visit-retain/analysis.getWeeklyRetain.html
func APIGetWeeklyRetain(c *gin.Context) {

	now := carbon.Now().AddWeeks(-2)

	from := now.SetWeekStartsAt(carbon.Monday).StartOfWeek().ToFormatString("Ymd", services.TIMEZONE)
	to := now.SetWeekStartsAt(carbon.Monday).EndOfWeek().ToFormatString("Ymd", services.TIMEZONE)

	rs, err := services.MiniProgramApp.DataCube.WeeklyRetainInfo(from, to)

	if err != nil {
		panic(err)
	}

	c.JSON(200, rs)

}

// 获取用户访问小程序数据概况
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/data-analysis/analysis.getDailySummary.html
func APIGetDailySummary(c *gin.Context) {

	now := carbon.Now().AddDays(-5)
	from := now.ToFormatString("Ymd", services.TIMEZONE)
	to := now.ToFormatString("Ymd", services.TIMEZONE)

	rs, err := services.MiniProgramApp.DataCube.SummaryTrend(from, to)

	if err != nil {
		panic(err)
	}

	c.JSON(200, rs)
}

// 获取小程序启动性能，运行性能等数据
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/data-analysis/analysis.getPerformanceData.html
func APIGetPerformanceData(c *gin.Context) {
	now := carbon.Now().AddDays(-5)
	beginTimestamp := now.Timestamp()
	endTimestamp := now.AddDays(2).Timestamp()

	options := &power.HashMap{
		"time": power.HashMap{
			"end_timestamp":   endTimestamp,
			"begin_timestamp": beginTimestamp,
		},
		"module": "10022",
		"params": []power.StringMap{
			power.StringMap{

				"field": "networktype",
				"value": "wifi",
			},
			power.StringMap{
				"field": "device_level",
				"value": "1",
			},
			power.StringMap{
				"field": "device",
				"value": "1",
			},
		},
	}

	rs, err := services.MiniProgramApp.DataCube.GetPerformanceData(options)

	if err != nil {
		panic(err)
	}

	c.JSON(200, rs)
}

// 获取用户访问小程序数据日趋势
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/data-analysis/visit-trend/analysis.getDailyVisitTrend.html
func APIGetDailyVisitTrend(c *gin.Context) {

	now := carbon.Now().AddDays(-5)
	from := now.ToFormatString("Ymd", services.TIMEZONE)
	to := now.ToFormatString("Ymd", services.TIMEZONE)

	rs, err := services.MiniProgramApp.DataCube.DailyVisitTrend(from, to)

	if err != nil {
		panic(err)
	}

	c.JSON(200, rs)

}

// 获取用户访问小程序数据月趋势(能查询到的最新数据为上一个自然月的数据)
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/data-analysis/visit-trend/analysis.getMonthlyVisitTrend.html
func APIGetMonthlyVisitTrend(c *gin.Context) {

	now := carbon.Now().AddMonths(-1)
	from := now.StartOfMonth().ToFormatString("Ymd", services.TIMEZONE)
	to := now.EndOfMonth().ToFormatString("Ymd", services.TIMEZONE)

	rs, err := services.MiniProgramApp.DataCube.MonthlyVisitTrend(from, to)

	if err != nil {
		panic(err)
	}

	c.JSON(200, rs)

}

// 获取用户访问小程序数据周趋势
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/data-analysis/visit-trend/analysis.getWeeklyVisitTrend.html
func APIGetWeeklyVisitTrend(c *gin.Context) {

	now := carbon.Now().AddWeeks(-2)
	from := now.SetWeekStartsAt(carbon.Monday).StartOfWeek().ToFormatString("Ymd", services.TIMEZONE)
	to := now.SetWeekStartsAt(carbon.Monday).EndOfWeek().ToFormatString("Ymd", services.TIMEZONE)

	rs, err := services.MiniProgramApp.DataCube.WeeklyVisitTrend(from, to)

	if err != nil {
		panic(err)
	}

	c.JSON(200, rs)

}

// 获取小程序新增或活跃用户的画像分布数据
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/data-analysis/analysis.getUserPortrait.html
func APIGetUserPortrait(c *gin.Context) {

	now := carbon.Now().AddWeeks(-1)
	from := now.SetWeekStartsAt(carbon.Monday).StartOfWeek().ToFormatString("Ymd", services.TIMEZONE)
	to := now.SetWeekStartsAt(carbon.Monday).EndOfWeek().ToFormatString("Ymd", services.TIMEZONE)

	rs, err := services.MiniProgramApp.DataCube.UserPortrait(from, to)

	if err != nil {
		panic(err)
	}

	c.JSON(200, rs)

}

// 获取小程序新增或活跃用户的画像分布数据
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/data-analysis/analysis.getUserPortrait.html
func APIGetVisitPage(c *gin.Context) {

	now := carbon.Now().AddWeeks(-1)
	from := now.SetWeekStartsAt(carbon.Monday).StartOfWeek().ToFormatString("Ymd", services.TIMEZONE)
	to := now.SetWeekStartsAt(carbon.Monday).EndOfWeek().ToFormatString("Ymd", services.TIMEZONE)

	rs, err := services.MiniProgramApp.DataCube.VisitPage(from, to)

	if err != nil {
		panic(err)
	}

	c.JSON(200, rs)

}
