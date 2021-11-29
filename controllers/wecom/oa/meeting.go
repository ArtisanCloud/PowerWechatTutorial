package oa

import (
  "github.com/ArtisanCloud/PowerWeChat/src/kernel/power"
  "github.com/ArtisanCloud/PowerWeChat/src/work/oa/meeting/request"
  "github.com/gin-gonic/gin"
  "net/http"
  "power-wechat-tutorial/services"
)

// 创建预约会议
// https://work.weixin.qq.com/api/doc/90000/90135/93627
func APIMeetingCreate(c *gin.Context) {

  options := &request.RequestMeetingCreate{
    CreatorUserID:   "zhangsan",
    Title:           "新建会议",
    MeetingStart:    1600000000,
    MeetingDuration: 3600,
    Description:     "新建会议描述",
    Type:            1,
    RemindTime:      60,
    AgentID:         1000014,
    Attendees: &power.HashMap{
      "userid": []string{
        "lisi",
        "wangwu",
      },
      "external_userid": []string{
        "woabc",
        "woced",
      },
      "device_sn": []string{
        "devsn1",
        "devsn2",
      },
    },
  }
  res, err := services.WeComApp.OAMeeting.Create(options)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, res)
}

func APIMeetingUpdate(c *gin.Context) {

  options := &request.RequestMeetingUpdate{
    MeetingID:       "XXXXXXXXX",
    Title:           "新需求",
    MeetingStart:    1600000000,
    MeetingDuration: 10000,
    Description:     "test",
    Type:            1,
    RemindTime:      60,
    Attendees: &power.HashMap{
      "userid": []string{
        "lisi",
        "wangwu",
      },

      "external_userid": []string{
        "woabc",
        "woced",
      },
      "device_sn": []string{
        "devsn1",
        "devsn2",
      },
    },
  }
  res, err := services.WeComApp.OAMeeting.Update(options)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, res)
}

func APIMeetingCancel(c *gin.Context) {
  meetingID := c.DefaultQuery("meetingID", "xxxxxx")
  res, err := services.WeComApp.OAMeeting.Cancel(meetingID)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, res)
}

func APIMeetingGetUserMeetingID(c *gin.Context) {

  userID := c.DefaultQuery("userID", "")
  cursor := "cursor"
  beginTime := int64(1586136317)
  endTime := int64(1586236317)
  limit := 100

  res, err := services.WeComApp.OAMeeting.GetUserMeetingID(userID, cursor, beginTime, endTime, limit)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, res)
}
