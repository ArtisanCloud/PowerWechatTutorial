package official_account

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"power-wechat-tutorial/services"
)

// 用户分析

// GetUserSummary 获取用户增减数据
func GetUserSummary(ctx *gin.Context) {
	from := ctx.Query("from")
	to := ctx.Query("to")
	data, err := services.OfficialAccountApp.DataCube.GetUserSummary(ctx.Request.Context(), from, to)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, data)
}

// 获取累计用户数据
func GetUserCumulate(ctx *gin.Context) {
	from := ctx.Query("from")
	to := ctx.Query("to")
	data, err := services.OfficialAccountApp.DataCube.GetUserCumulate(ctx.Request.Context(), from, to)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, data)
}

// 图文分析

// 获取图文群发每日数据
func ArticleSummary(ctx *gin.Context) {
	from := ctx.Query("from")
	to := ctx.Query("to")
	data, err := services.OfficialAccountApp.DataCube.ArticleSummary(ctx.Request.Context(), from, to)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, data)
}

// 获取图文群发总数据
func ArticleTotal(ctx *gin.Context) {
	from := ctx.Query("from")
	to := ctx.Query("to")
	data, err := services.OfficialAccountApp.DataCube.ArticleTotal(ctx.Request.Context(), from, to)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, data)
}

// 获取图文统计数据
func UserReadSummary(ctx *gin.Context) {
	from := ctx.Query("from")
	to := ctx.Query("to")
	data, err := services.OfficialAccountApp.DataCube.UserReadSummary(ctx.Request.Context(), from, to)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, data)
}

// TODO: 缺少getuserreadhour

// 获取图文统计分时数据（getuserreadhour）
//func UserShareHourly(ctx *gin.Context) {
//  from := ctx.Query("from")
//  to := ctx.Query("to")
//  data, err := services.OfficialAccountApp.DataCube.UserShareHourly(from, to)
//  if err != nil {
//    ctx.String(http.StatusBadRequest, err.Error())
//    return
//  }
//  ctx.JSON(http.StatusOK, data)
//}

// 获取图文分享转发数据
func UserShareSummary(ctx *gin.Context) {
	from := ctx.Query("from")
	to := ctx.Query("to")
	data, err := services.OfficialAccountApp.DataCube.UpstreamMessageSummary(ctx.Request.Context(), from, to)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, data)
}

// 获取图文分享转发分时数据
func UserShareHourly(ctx *gin.Context) {
	from := ctx.Query("from")
	to := ctx.Query("to")
	data, err := services.OfficialAccountApp.DataCube.UserShareHourly(ctx.Request.Context(), from, to)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, data)
}

// 消息分析

// 获取消息发送概况数据
func UpstreamMessageSummary(ctx *gin.Context) {
	from := ctx.Query("from")
	to := ctx.Query("to")
	data, err := services.OfficialAccountApp.DataCube.UpstreamMessageSummary(ctx.Request.Context(), from, to)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, data)
}

// 获取消息发送分时数据
func UpstreamMessageHourly(ctx *gin.Context) {
	from := ctx.Query("from")
	to := ctx.Query("to")
	data, err := services.OfficialAccountApp.DataCube.UpstreamMessageHourly(ctx.Request.Context(), from, to)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, data)
}

// 获取消息发送周数据
func UpstreamMessageWeekly(ctx *gin.Context) {
	from := ctx.Query("from")
	to := ctx.Query("to")
	data, err := services.OfficialAccountApp.DataCube.UpstreamMessageWeekly(ctx.Request.Context(), from, to)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, data)
}

// 获取消息发送月数据
func UpstreamMessageMonthly(ctx *gin.Context) {
	from := ctx.Query("from")
	to := ctx.Query("to")
	data, err := services.OfficialAccountApp.DataCube.UpstreamMessageMonthly(ctx.Request.Context(), from, to)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, data)
}

// 获取消息发送分布数据
func UpstreamMessageDistSummary(ctx *gin.Context) {
	from := ctx.Query("from")
	to := ctx.Query("to")
	data, err := services.OfficialAccountApp.DataCube.UpstreamMessageDistSummary(ctx.Request.Context(), from, to)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, data)
}

// 获取消息发送分布周数据
func UpstreamMessageDistWeekly(ctx *gin.Context) {
	from := ctx.Query("from")
	to := ctx.Query("to")
	data, err := services.OfficialAccountApp.DataCube.UpstreamMessageDistWeekly(ctx.Request.Context(), from, to)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, data)
}

// 获取消息发送分布月数据
func UpstreamMessageDistMonthly(ctx *gin.Context) {
	from := ctx.Query("from")
	to := ctx.Query("to")
	data, err := services.OfficialAccountApp.DataCube.UpstreamMessageDistMonthly(ctx.Request.Context(), from, to)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, data)
}

// 广告分析

// 获取公众号分广告位数据
//func GetUserCumulate(ctx *gin.Context) {
//  from := ctx.Query("from")
//  to := ctx.Query("to")
//  data, err := services.OfficialAccountApp.DataCube.GetUserCumulate(from, to)
//  if err != nil {
//    ctx.String(http.StatusBadRequest, err.Error())
//    return
//  }
//  ctx.JSON(http.StatusOK, data)
//}
//
//// 获取公众号返佣商品数据
//func GetUserCumulate(ctx *gin.Context) {
//  from := ctx.Query("from")
//  to := ctx.Query("to")
//  data, err := services.OfficialAccountApp.DataCube.GetUserCumulate(from, to)
//  if err != nil {
//    ctx.String(http.StatusBadRequest, err.Error())
//    return
//  }
//  ctx.JSON(http.StatusOK, data)
//}
//
//// 获取公众号结算收入数据及结算主体信息
//func GetUserCumulate(ctx *gin.Context) {
//  from := ctx.Query("from")
//  to := ctx.Query("to")
//  data, err := services.OfficialAccountApp.DataCube.GetUserCumulate(from, to)
//  if err != nil {
//    ctx.String(http.StatusBadRequest, err.Error())
//    return
//  }
//  ctx.JSON(http.StatusOK, data)
//}

// 接口分析

// 获取接口分析数据
func InterfaceSummary(ctx *gin.Context) {
	from := ctx.Query("from")
	to := ctx.Query("to")
	data, err := services.OfficialAccountApp.DataCube.InterfaceSummary(ctx.Request.Context(), from, to)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, data)
}

// 获取接口分析分时数据
func InterfaceSummaryHourly(ctx *gin.Context) {
	from := ctx.Query("from")
	to := ctx.Query("to")
	data, err := services.OfficialAccountApp.DataCube.InterfaceSummaryHourly(ctx.Request.Context(), from, to)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, data)
}
