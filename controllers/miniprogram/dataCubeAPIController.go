package miniprogram

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-module/carbon"
	"power-wechat-tutorial/services"
)

// 获取用户访问小程序日留存
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/data-analysis/visit-retain/analysis.getDailyRetain.html
func APIGetDailyRetain(c *gin.Context) {

	now := carbon.Now()
	from := now.AddDays(-1).ToFormatString("YMD", "UTC")
	to := now.ToFormatString("YMD", "UTC")

	rs, err := services.AppMiniProgram.DataCube.DailyRetainInfo(from, to)

	if err != nil {
		panic(err)
	}

	c.JSON(200, rs)

}

// 获取用户访问小程序日留存
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/data-analysis/visit-retain/analysis.getDailyRetain.html
func APIGetMonthlyRetain(c *gin.Context) {

	now := carbon.Now()
	from := now.AddDays(-25).ToFormatString("YMD", "UTC")
	to := now.ToFormatString("YMD", "UTC")

	rs, err := services.AppMiniProgram.DataCube.DailyRetainInfo(from, to)

	if err != nil {
		panic(err)
	}

	c.JSON(200, rs)

}

// 获取用户访问小程序周留存
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/data-analysis/visit-retain/analysis.getWeeklyRetain.html
func APIGetWeeklyRetain(c *gin.Context) {

	now := carbon.Now()
	from := now.AddDays(-7).ToFormatString("YMD", "UTC")
	to := now.ToFormatString("YMD", "UTC")

	rs, err := services.AppMiniProgram.DataCube.WeeklyRetainInfo(from, to)

	if err != nil {
		panic(err)
	}

	c.JSON(200, rs)

}

// 获取用户访问小程序数据概况
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/data-analysis/analysis.getDailySummary.html
func APIGetDailySummary(c *gin.Context) {

	now := carbon.Now()
	from := now.AddDays(-7).ToFormatString("YMD", "UTC")
	to := now.ToFormatString("YMD", "UTC")

	rs, err := services.AppMiniProgram.DataCube.SummaryTrend(from, to)

	if err != nil {
		panic(err)
	}

	c.JSON(200, rs)
}

// 获取小程序启动性能，运行性能等数据
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/data-analysis/analysis.getPerformanceData.html
func APIGetPerformanceData(c *gin.Context) {

	rs, err := services.AppMiniProgram.DataCube.GetPerformanceData(nil)

	if err != nil {
		panic(err)
	}

	c.JSON(200, rs)

}


// 获取用户访问小程序数据日趋势
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/data-analysis/visit-trend/analysis.getDailyVisitTrend.html
func APIGetDailyVisitTrend(c *gin.Context) {

	now := carbon.Now()
	from := now.AddDays(-1).ToFormatString("YMD", "UTC")
	to := now.ToFormatString("YMD", "UTC")

	rs, err := services.AppMiniProgram.DataCube.DailyVisitTrend(from, to)

	if err != nil {
		panic(err)
	}

	c.JSON(200, rs)

}

// 获取用户访问小程序数据月趋势(能查询到的最新数据为上一个自然月的数据)
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/data-analysis/visit-trend/analysis.getMonthlyVisitTrend.html
func APIGetMonthlyVisitTrend(c *gin.Context) {

	now := carbon.Now()
	from := now.AddDays(-25).ToFormatString("YMD", "UTC")
	to := now.ToFormatString("YMD", "UTC")

	rs, err := services.AppMiniProgram.DataCube.MonthlyVisitTrend(from, to)

	if err != nil {
		panic(err)
	}

	c.JSON(200, rs)

}

// 获取用户访问小程序数据周趋势
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/data-analysis/visit-trend/analysis.getWeeklyVisitTrend.html
func APIGetWeeklyVisitTrend(c *gin.Context) {
	now := carbon.Now()
	from := now.AddDays(-7).ToFormatString("YMD", "UTC")
	to := now.ToFormatString("YMD", "UTC")

	rs, err := services.AppMiniProgram.DataCube.WeeklyVisitTrend(from, to)

	if err != nil {
		panic(err)
	}

	c.JSON(200, rs)

}

// 获取小程序新增或活跃用户的画像分布数据
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/data-analysis/analysis.getUserPortrait.html
func APIGetUserPortrait(c *gin.Context) {

	now := carbon.Now()
	from := now.AddDays(-7).ToFormatString("YMD", "UTC")
	to := now.ToFormatString("YMD", "UTC")

	rs, err := services.AppMiniProgram.DataCube.UserPortrait(from, to)

	if err != nil {
		panic(err)
	}

	c.JSON(200, rs)

}

// 获取小程序新增或活跃用户的画像分布数据
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/data-analysis/analysis.getUserPortrait.html
func APIGetVisitPage(c *gin.Context) {

	now := carbon.Now()
	from := now.AddDays(-7).ToFormatString("YMD", "UTC")
	to := now.ToFormatString("YMD", "UTC")

	rs, err := services.AppMiniProgram.DataCube.VisitPage(from, to)

	if err != nil {
		panic(err)
	}

	c.JSON(200, rs)

}
