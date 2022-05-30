package external_contact

import (
  "github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/power"
  "github.com/ArtisanCloud/PowerWeChat/v2/src/work/externalContact/statistics/request"
  "github.com/gin-gonic/gin"
  "net/http"
  "power-wechat-tutorial/services"
)

// 获取「联系客户统计」数据
// https://work.weixin.qq.com/api/doc/90000/90135/92132
func APIExternalContactGetUserBehaviorData(c *gin.Context) {

  options := &request.RequestGetUserBehaviorData{
    UserID:    []string{c.DefaultQuery("userID", "matrix-x")},
    PartyID:   []int{1001},
    StartTime: 1536508800,
    EndTime:   1536595200,
  }

  res, err := services.WeComContactApp.ExternalContactStatistics.GetUserBehaviorData(options)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, res)
}

// 获取「群聊数据统计」数据
// https://work.weixin.qq.com/api/doc/90000/90135/92133
func APIExternalContactGroupChatStatistic(c *gin.Context) {
  options := &request.RequestStatistic{
    DayBeginTime: 1600272000,
    DayEndTime:   1600444800,
    OwnerFilter: &power.HashMap{
      "userid_list": []string{c.DefaultQuery("userID", "matrix-x")},
    },
    OrderBy:  2,
    OrderAsc: 0,
    Offset:   0,
    Limit:    1000,
  }
  res, err := services.WeComContactApp.ExternalContactStatistics.Statistic(options)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, res)
}
