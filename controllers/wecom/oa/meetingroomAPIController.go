package oa

import (
  "github.com/ArtisanCloud/power-wechat/src/kernel/power"
  "github.com/ArtisanCloud/power-wechat/src/work/oa/meetingroom/request"
  "github.com/gin-gonic/gin"
  "net/http"
  "power-wechat-tutorial/services"
)

// 添加会议室
// https://work.weixin.qq.com/api/doc/90000/90135/93619
func APIMeetingRoomAdd(c *gin.Context) {

  options := &request.RequestMeetingRoomAdd{
    Name:      "18F-会议室",
    Capacity:  10,
    City:      "深圳",
    Building:  "腾讯大厦",
    Floor:     "18F",
    Equipment: []int{1, 2, 3},
    Coordinate: &power.StringMap{
      "latitude":  "22.540503",
      "longitude": "113.934528",
    },
  }

  res, err := services.WeComApp.OAMeetingRoom.Add(options)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, res)
}

// 查询会议室
// https://work.weixin.qq.com/api/doc/90000/90135/93619
func APIMeetingRoomList(c *gin.Context) {

  options := &request.RequestMeetingRoomList{
    City:      "深圳",
    Building:  "腾讯大厦",
    Floor:     "18F",
    Equipment: []int{1, 2, 3},
  }

  res, err := services.WeComApp.OAMeetingRoom.List(options)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, res)
}

// 编辑会议室
// https://work.weixin.qq.com/api/doc/90000/90135/93619
func APIMeetingRoomEdit(c *gin.Context) {

  options := &request.RequestMeetingRoomEdit{
    MeetingRoomID: 2,
    Name:          "18F-会议室",
    Capacity:      10,
    City:          "深圳",
    Building:      "腾讯大厦",
    Floor:         "18F",
    Equipment:     []int{1, 2, 3},
    Coordinate: &power.StringMap{
      "latitude":  "22.540503",
      "longitude": "113.934528",
    },
  }

  res, err := services.WeComApp.OAMeetingRoom.Edit(options)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, res)
}

// 删除会议室
// https://work.weixin.qq.com/api/doc/90000/90135/93619
func APIMeetingRoomDel(c *gin.Context) {
  options := &request.RequestMeetingRoomDel{
    MeetingRoomID: 1,
  }
  res, err := services.WeComApp.OAMeetingRoom.Del(options)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, res)
}

// 查询会议室的预定信息
// https://work.weixin.qq.com/api/doc/90000/90135/93620#%E6%9F%A5%E8%AF%A2%E4%BC%9A%E8%AE%AE%E5%AE%A4%E7%9A%84%E9%A2%84%E5%AE%9A%E4%BF%A1%E6%81%AF
func APIMeetingRoomGetBookingInfo(c *gin.Context) {

  options := &request.RequestMeetingRoomGetBookingInfo{
    MeetingroomID: 1,
    StartTime:     1593532800,
    EndTime:       1593619200,
    City:          "深圳",
    Building:      "腾讯大厦",
    Floor:         "18F",
  }
  res, err := services.WeComApp.OAMeetingRoom.GetBookingInfo(options)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, res)
}

// 预定会议室
// https://work.weixin.qq.com/api/doc/90000/90135/93620
func APIMeetingRoomBook(c *gin.Context) {

  options := &request.RequestMeetingRoomBook{
    MeetingRoomID: 1,
    Subject:       "周会",
    StartTime:     1593532800,
    EndTime:       1593619200,
    Booker:        "zhangsan",
    Attendees:     []string{"lisi", "wangwu"},
  }

  res, err := services.WeComApp.OAMeetingRoom.Book(options)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, res)
}

func APIMeetingRoomCancelBook(c *gin.Context) {

  options := &request.RequestMeetingRoomCancelBook{
    MeetingID:    c.DefaultQuery("meetingID","mt42b34949gsaseb6e027c123cbafAAA"),
    KeepSchedule: 1,
  }

  res, err := services.WeComApp.OAMeetingRoom.CancelBook(options)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, res)
}
