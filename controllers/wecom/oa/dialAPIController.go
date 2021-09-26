package oa

import (
  "github.com/ArtisanCloud/power-wechat/src/work/oa/dial/request"
  "github.com/gin-gonic/gin"
  "net/http"
  "power-wechat-tutorial/services"
)

func APIDialGetDialRecord(c *gin.Context) {

  options := &request.RequestDialGetDialRecord{
    MeetingID:       1536508800,
    Title:           1536940800,
    MeetingStart:    0,
    MeetingDuration: 100,
  }
  res, err := services.WeComApp.OADial.GetDialRecord(options)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, res)
}
