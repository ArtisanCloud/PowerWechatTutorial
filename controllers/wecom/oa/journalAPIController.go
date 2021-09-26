package oa

import (
  "github.com/ArtisanCloud/power-wechat/src/kernel/power"
  "github.com/gin-gonic/gin"
  "net/http"
  "power-wechat-tutorial/services"
)

// 批量获取汇报记录单号
// https://work.weixin.qq.com/api/doc/90000/90135/93393
func APIJournalGetRecordList(c *gin.Context) {

  starttime := 1606230000
  endtime := 1606361304
  cursor := 0
  limit := 10
  filters := []*power.StringMap{
    &power.StringMap{
      "key":   "creator",
      "value": "kele",
    },
    &power.StringMap{
      "key":   "department",
      "value": "1",
    },
    &power.StringMap{
      "key":   "template_id",
      "value": "3TmALk1ogfgKiQE3e3jRwnTUhMTh8vca1N8zUVNUx",
    },
  }

  res, err := services.WeComApp.OAJournal.GetRecordList(starttime, endtime, cursor, limit, filters)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, res)
}

// 获取汇报记录详情
// https://work.weixin.qq.com/api/doc/90000/90135/93394
func APIJournalGetRecordDetail(c *gin.Context) {

  journalUUID := c.DefaultQuery("journalUUID", "41eJejN57EJNzr8HrZfmKyCN7xwKw1qRxCZUxCVuo9fsWVMSKac6nk4q8rARTDaVNdx")
  res, err := services.WeComApp.OAJournal.GetRecordDetail(journalUUID)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, res)
}

// 获取汇报记录详情
// https://work.weixin.qq.com/api/doc/90000/90135/93395
func APIJournalGetStatList(c *gin.Context) {

  templateID := c.DefaultQuery("templateID", "41eJejN57EJNzr8HrZfmKyCN7xwKw1qRxCZUxCVuo9fsWVMSKac6nk4q8rARTDaVNdx")
  startTime := 1604160000
  endTime := 1606363092
  res, err := services.WeComApp.OAJournal.GetStatList(templateID, startTime, endTime)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, res)
}
