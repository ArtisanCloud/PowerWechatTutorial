package user

import (
	"github.com/gin-gonic/gin"
  "net/http"
  "power-wechat-tutorial/services"
)

// APILinkedCorpAgentGetPermList 获取应用的可见范围
// https://open.work.weixin.qq.com/api/doc/90000/90135/93172
func APILinkedCorpAgentGetPermList(c *gin.Context) {

  res, err := services.WeComApp.UserLinkedCorp.GetPermList()
  if err != nil {
    panic(err)
  }
  c.JSON(http.StatusOK, res)
}

// APILinkedCorpUserGet 获取互联企业成员详细信息
// https://work.weixin.qq.com/api/doc/90000/90135/93171
func APILinkedCorpUserGet(c *gin.Context) {
  userID := c.DefaultQuery("userID", "walle")
  res, err := services.WeComApp.UserLinkedCorp.GetUser(userID)
  if err != nil {
    panic(err)
  }
  c.JSON(http.StatusOK, res)
}

// APILinkedCorpUserSimpleList 获取互联企业部门成员
// https://work.weixin.qq.com/api/doc/90000/90135/93168
func APILinkedCorpUserSimpleList(c *gin.Context) {
  departmentID := c.DefaultQuery("departmentID", "xxx")
  res, err := services.WeComApp.UserLinkedCorp.GetUserSimpleList(departmentID, true)
  if err != nil {
    panic(err)
  }
  c.JSON(http.StatusOK, res)

}

// APILinkedCorpUserList 获取互联企业部门成员详情
// https://work.weixin.qq.com/api/doc/90000/90135/93169
func APILinkedCorpUserList(c *gin.Context) {
  departmentID := c.DefaultQuery("departmentID", "xxx")
  res, err := services.WeComApp.UserLinkedCorp.GetUserList(departmentID, true)
  if err != nil {
    panic(err)
  }
  c.JSON(http.StatusOK, res)
}

// APILinkedCorpDepartmentList 获取互联企业部门列表
// https://work.weixin.qq.com/api/doc/90000/90135/93170
func APILinkedCorpDepartmentList(c *gin.Context) {
  departmentID := c.DefaultQuery("departmentID", "xxx")
  res, err := services.WeComApp.UserLinkedCorp.GetDepartmentList(departmentID)
  if err != nil {
    panic(err)
  }
  c.JSON(http.StatusOK, res)
}
