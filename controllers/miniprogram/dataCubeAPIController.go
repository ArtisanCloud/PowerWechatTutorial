package miniprogram

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/power"
	"github.com/gin-gonic/gin"
	"net/http"
	"power-wechat-tutorial/services"
	"time"
)

// 获取用户访问小程序日留存
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/data-analysis/visit-retain/analysis.getDailyRetain.html
func APIGetDailyRetain(c *gin.Context) {

	now := time.Now().Add(-5 * 24 * time.Hour)

	from := now.Format(services.DATETIME_FORMAT)
	to := now.Format(services.DATETIME_FORMAT)

	rs, err := services.AppMiniProgram.DataCube.GetDailyRetainInfo(from, to)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, rs)

}

// 获取用户访问小程序周留存
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/data-analysis/visit-retain/analysis.getWeeklyRetain.html
func APIGetWeeklyRetain(c *gin.Context) {

	from := "20210906"
	to := "20210912"

	rs, err := services.AppMiniProgram.DataCube.GetWeeklyRetainInfo(from, to)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, rs)

}

// 获取用户访问小程序月留存
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/data-analysis/visit-retain/analysis.getMonthlyRetain.html
func APIGetMonthlyRetain(c *gin.Context) {

	from := "20170201"
	to := "20170228"

	rs, err := services.AppMiniProgram.DataCube.GetMonthlyRetainInfo(from, to)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, rs)

}

// 获取用户访问小程序数据概况
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/data-analysis/analysis.getDailySummary.html
func APIGetDailySummary(c *gin.Context) {

	now := time.Now().Add(-5 * 24 * time.Hour)

	from := now.Format(services.DATETIME_FORMAT)
	to := now.Format(services.DATETIME_FORMAT)

	rs, err := services.AppMiniProgram.DataCube.GetDailySummary(from, to)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, rs)
}

// 获取用户访问小程序数据日趋势
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/data-analysis/visit-trend/analysis.getDailyVisitTrend.html
func APIGetDailyVisitTrend(c *gin.Context) {

	now := time.Now().Add(-5 * 24 * time.Hour)

	from := now.Format(services.DATETIME_FORMAT)
	to := now.Format(services.DATETIME_FORMAT)

	rs, err := services.AppMiniProgram.DataCube.GetDailyVisitTrend(from, to)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, rs)

}

// 获取用户访问小程序数据月趋势(能查询到的最新数据为上一个自然月的数据)
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/data-analysis/visit-trend/analysis.getMonthlyVisitTrend.html
func APIGetMonthlyVisitTrend(c *gin.Context) {

	from := "20210831"
	to := "20210930"

	rs, err := services.AppMiniProgram.DataCube.GetMonthlyVisitTrend(from, to)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, rs)

}

// 获取用户访问小程序数据周趋势
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/data-analysis/visit-trend/analysis.getWeeklyVisitTrend.html
func APIGetWeeklyVisitTrend(c *gin.Context) {

	from := "20210906"
	to := "20210912"

	rs, err := services.AppMiniProgram.DataCube.GetWeeklyVisitTrend(from, to)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, rs)

}

// 获取小程序启动性能，运行性能等数据
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/data-analysis/analysis.getPerformanceData.html
func APIGetPerformanceData(c *gin.Context) {

	now := time.Now().Add(-5 * 24 * time.Hour)
	beginTimestamp := now.Unix()
	endTimestamp := now.Add(2 * 24 * time.Hour).Unix()

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

	rs, err := services.AppMiniProgram.DataCube.GetPerformanceData(options)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, rs)
}

// 获取小程序新增或活跃用户的画像分布数据
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/data-analysis/analysis.getUserPortrait.html
func APIGetUserPortrait(c *gin.Context) {

	from := "20210906"
	to := "20210912"

	rs, err := services.AppMiniProgram.DataCube.GetUserPortrait(from, to)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, rs)

}

// 获取用户小程序访问分布数据
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/data-analysis/analysis.getVisitDistribution.html
func APIGetVisitDistribution(c *gin.Context) {

	from := "20210906"
	to := "20210906"

	rs, err := services.AppMiniProgram.DataCube.GetVisitDistribution(from, to)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, rs)

}

// 获取小程序新增或活跃用户的画像分布数据
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/data-analysis/analysis.getUserPortrait.html
func APIGetVisitPage(c *gin.Context) {

	from := "20210906"
	to := "20210912"

	rs, err := services.AppMiniProgram.DataCube.GetVisitPage(from, to)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, rs)

}
