package external_contact

import (
  "github.com/ArtisanCloud/PowerWeChat/src/kernel/power"
  "github.com/gin-gonic/gin"
  "net/http"
  "power-wechat-tutorial/services"
  "strconv"
)

const defaultDepartmentId = "3000001"

// DepartmentCreate 创建部门
// https://open.work.weixin.qq.com/api/doc/90000/90135/90205
func APIDepartmentCreate(c *gin.Context) {
  name := c.DefaultQuery("name", "IT支持部")
  idStr := c.DefaultQuery("id", defaultDepartmentId)
  id, _ := strconv.Atoi(idStr)

  res, err := services.WeComContactApp.Department.Create(&power.HashMap{
    "name":     name,
    "parentid": 1,
    "id":       id,
  })

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, res)
}

// DepartmentUpdate 更新部门
// https://open.work.weixin.qq.com/api/doc/90000/90135/90206
func APIDepartmentUpdate(c *gin.Context) {
  name := c.DefaultQuery("name", "IT支持部1")
  idStr := c.DefaultQuery("id", defaultDepartmentId)
  id, _ := strconv.Atoi(idStr)
  res, err := services.WeComContactApp.Department.Update(id, &power.HashMap{
    "name":     name,
    "parentid": 1,
  })

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, res)
}

// DepartmentDelete 删除部门
// https://open.work.weixin.qq.com/api/doc/90000/90135/90207
func APIDepartmentDelete(c *gin.Context) {
  idStr := c.DefaultQuery("id", defaultDepartmentId)
  id, _ := strconv.Atoi(idStr)
  res, err := services.WeComContactApp.Department.Delete(id)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, res)
}

// 获取部门列表
// https://open.work.weixin.qq.com/api/doc/90000/90135/90208
func APIDepartmentList(c *gin.Context) {
  res, err := services.WeComContactApp.Department.List()

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, res)
}
