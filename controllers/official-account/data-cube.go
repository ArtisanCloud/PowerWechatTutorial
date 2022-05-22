package official_account

import (
  "github.com/gin-gonic/gin"
  "net/http"
  "power-wechat-tutorial/services"
)

// 用户分析

// 获取用户增减数据
func GetUserSummary(ctx *gin.Context) {
  from := ctx.Query("from")
  to := ctx.Query("to")
  data, err := services.OfficialAccountApp.DataCube.GetDailySummary(from, to)
  if err != nil {
    ctx.JSON(http.StatusBadRequest, err)
    return
  }
  ctx.JSON(http.StatusOK, data)
}

// 获取累计用户数据

// 图文分析

// 获取图文群发每日数据

// 获取图文群发总数据

// 获取图文统计数据

// 获取图文统计分时数据（getuserreadhour）

// 获取图文分享转发数据

// 获取图文分享转发分时数据

// 消息分析

// 获取消息发送概况数据

// 获取消息发送分时数据

// 获取消息发送周数据

// 获取消息发送月数据

// 获取消息发送分布数据

// 获取消息发送分布周数据

// 获取消息发送分布月数据

// 广告分析

// 获取公众号分广告位数据

// 获取公众号返佣商品数据

// 获取公众号结算收入数据及结算主体信息
