package oa

import (
  "github.com/ArtisanCloud/PowerWeChat/src/kernel/power"
  "github.com/ArtisanCloud/PowerWeChat/src/work/oa/request"
  "github.com/gin-gonic/gin"
  "net/http"
  "power-wechat-tutorial/services"
)

// 获取企业所有打卡规则
// https://work.weixin.qq.com/api/doc/90000/90135/93384
func APICheckinGetCorpCheckinOption(c *gin.Context) {

  res, err := services.WeComApp.OA.GetCorpCheckInOption()

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, res)
}

// 获取员工打卡规则
// https://work.weixin.qq.com/api/doc/90000/90135/90263
func APICheckinGetCheckinOption(c *gin.Context) {

  datetime := 1511971200
  userIDList := []string{c.DefaultQuery("userID", "matrix-x")}

  res, err := services.WeComApp.OA.GetCheckInOption(datetime, userIDList)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, res)
}


// 获取打卡记录数据
// https://work.weixin.qq.com/api/doc/90000/90135/90262
func APICheckinGetCheckinData(c *gin.Context) {

  userIDList := []string{c.DefaultQuery("userID", "matrix-x")}

  res, err := services.WeComApp.OA.GetCheckinData(3, 1492617600, 1492790400, userIDList)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, res)
}

// 获取打卡日报数据
// https://work.weixin.qq.com/api/doc/90000/90135/93374
func APICheckinGetCheckinDayData(c *gin.Context) {
  userIDList := []string{c.DefaultQuery("userID", "matrix-x")}

  res, err := services.WeComApp.OA.GetCheckinDayData(1599062400, 1599062400, userIDList)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, res)
}

// 获取打卡月报数据
// https://work.weixin.qq.com/api/doc/90000/90135/93387
func APICheckinGetCheckinMonthData(c *gin.Context) {
  userIDList := []string{c.DefaultQuery("userID", "matrix-x")}

  res, err := services.WeComApp.OA.GetCheckInMonthData(1599062400, 1599408000, userIDList)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, res)
}

// 获取打卡人员排班信息
// https://work.weixin.qq.com/api/doc/90000/90135/93380
func APICheckinGetCheckinSchedulist(c *gin.Context) {
  userIDList := []string{c.DefaultQuery("userID", "matrix-x")}

  res, err := services.WeComApp.OA.GetCheckInScheduleList(1492617600, 1492790400, userIDList)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, res)
}

// 为打卡人员排班
// https://work.weixin.qq.com/api/doc/90000/90135/93385
func APICheckinSetCheckinSchedulist(c *gin.Context) {

  options := &request.RequestCheckInSetScheduleList{
    GroupID: 226,
    Items: []*power.HashMap{
      &power.HashMap{
        "userid":      "james",
        "day":         5,
        "schedule_id": 234,
      },
    },
    YearMonth: 202012,
  }

  res, err := services.WeComApp.OA.SetCheckInScheduleList(options)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, res)

}

// 录入打卡人员人脸信息
// https://work.weixin.qq.com/api/doc/90000/90135/93378
func APICheckinAddCheckinUserFace(c *gin.Context) {
  userID := c.DefaultQuery("userID", "matrix-x")
  userFace := c.DefaultQuery("userFace", "PLACE_HOLDER")

  res, err := services.WeComApp.OA.AddCheckInUserFace(userID, userFace)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, res)
}
