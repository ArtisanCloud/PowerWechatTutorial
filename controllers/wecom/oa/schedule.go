package oa

import (
  "github.com/ArtisanCloud/PowerWeChat/src/kernel/power"
  "github.com/gin-gonic/gin"
  "net/http"
  "power-wechat-tutorial/services"
  "strconv"
)

// 创建日程
// https://work.weixin.qq.com/api/doc/90000/90135/93648
func APIScheduleAdd(c *gin.Context) {

  agentId := c.DefaultQuery("agentID", "1000014")
  agentID, _ := strconv.Atoi(agentId)

  schedule := &power.HashMap{
    "organizer":  "userid1",
    "start_time": 1571274600,
    "end_time":   1571320210,
    "attendees": []power.HashMap{
      power.HashMap{
        "userid": "userid2",
      },
    },
    "summary":     "需求评审会议",
    "description": "2.0版本需求初步评审",
    "reminders": power.HashMap{
      "is_remind":                1,
      "remind_before_event_secs": 3600,
      "is_repeat":                1,
      "repeat_type":              7,
      "repeat_until":             1606976813,
      "is_custom_repeat":         1,
      "repeat_interval":          1,
      "repeat_day_of_week":       []int{3, 7},
      "repeat_day_of_month":      []int{10, 21},
      "timezone":                 8,
    },
    "location": "广州国际媒体港10楼1005会议室",
    "cal_id":   "wcjgewCwAAqeJcPI1d8Pwbjt7nttzAAA",
  }

  res, err := services.WeComApp.OASchedule.Add(schedule, agentID)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, res)
}

// 更新日程
// https://work.weixin.qq.com/api/doc/90000/90135/93648
func APIScheduleUpdate(c *gin.Context) {

  schedule := &power.HashMap{
    "organizer":   "userid1",
    "schedule_id": "17c7d2bd9f20d652840f72f59e796AAA",
    "start_time":  1571274600,
    "end_time":    1571320210,
    "attendees": []power.HashMap{
      {
        "userid": "userid2",
      },
    },
    "summary":     "test_summary",
    "description": "test_description",
    "reminders": power.HashMap{
      "is_remind":                1,
      "remind_before_event_secs": 3600,
      "is_repeat":                1,
      "repeat_type":              7,
      "repeat_until":             1606976813,
      "is_custom_repeat":         1,
      "repeat_interval":          1,
      "repeat_day_of_week":       []int{3, 7},
      "repeat_day_of_month":      []int{10, 21},
      "timezone":                 8,
    },
    "location": "test_place",
  }
  res, err := services.WeComApp.OASchedule.Update(schedule)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, res)
}

// 获取日程详情
// https://work.weixin.qq.com/api/doc/90000/90135/93648
func APIScheduleGet(c *gin.Context) {
  scheduleIDList := []string{
    c.DefaultQuery("scheduleID", "17c7d2bd9f20d652840f72f59e796AAA"),
  }
  res, err := services.WeComApp.OASchedule.Get(scheduleIDList)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, res)
}

// 删除日程
// https://work.weixin.qq.com/api/doc/90000/90135/93648
func APIScheduleDel(c *gin.Context) {
  scheduleID := c.DefaultQuery("scheduleID", "17c7d2bd9f20d652840f72f59e796AAA")
  res, err := services.WeComApp.OASchedule.Del(scheduleID)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, res)
}
