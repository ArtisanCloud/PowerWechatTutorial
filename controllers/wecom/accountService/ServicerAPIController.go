package accountService

import (
  "github.com/gin-gonic/gin"
  "net/http"
  "power-wechat-tutorial/services"
)

// 添加接待人员
// https://work.weixin.qq.com/api/doc/90000/90135/94646
func APIAccountServiceServicerAdd(c *gin.Context) {

  openKFID := c.DefaultQuery("openKFID", "kfxxxxxxxxxxxxxx")
  userIDList := []string{c.DefaultQuery("userIDList", "matrix-x")}

  res, err := services.WeComApp.AccountServiceServicer.Add(openKFID, userIDList)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, res)
}

// 删除接待人员
// https://work.weixin.qq.com/api/doc/90000/90135/94647
func APIAccountServiceServicerDel(c *gin.Context) {

  openKFID := c.DefaultQuery("openKFID", "kfxxxxxxxxxxxxxx")
  userIDList := []string{c.DefaultQuery("userIDList", "matrix-x")}

  res, err := services.WeComApp.AccountServiceServicer.Del(openKFID, userIDList)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, res)
}

// 获取接待人员列表
// https://work.weixin.qq.com/api/doc/90000/90135/94645
func APIAccountServiceServicerList(c *gin.Context) {
  openKFID := c.DefaultQuery("openKFID", "kfxxxxxxxxxxxxxx")

  res, err := services.WeComApp.AccountServiceServicer.List(openKFID)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, res)
}
