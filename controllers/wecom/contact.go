package wecom

import (
  "github.com/ArtisanCloud/power-wechat/src/kernel/power"
  "github.com/ArtisanCloud/power-wechat/src/work/user/request"
  "github.com/gin-gonic/gin"
  "power-wechat-tutorial/services"
  "strconv"
)

func DepartmentList(c *gin.Context) {
  res, err := services.WeComContactApp.Department.List()

  if err != nil {
    panic(err)
  }

  c.JSON(200, res)
}

const defaultDepartmentId = "3000001"

// DepartmentCreate 创建部门
// https://open.work.weixin.qq.com/api/doc/90000/90135/90205
func DepartmentCreate(c *gin.Context) {
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

  c.JSON(200, res)
}

// DepartmentUpdate 更新部门
// https://open.work.weixin.qq.com/api/doc/90000/90135/90206
func DepartmentUpdate(c *gin.Context) {
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

  c.JSON(200, res)
}

// DepartmentDelete 删除部门
// https://open.work.weixin.qq.com/api/doc/90000/90135/90207
func DepartmentDelete(c *gin.Context) {
  idStr := c.DefaultQuery("id", defaultDepartmentId)
  id, _ := strconv.Atoi(idStr)
  res, err := services.WeComContactApp.Department.Delete(id)

  if err != nil {
    panic(err)
  }

  c.JSON(200, res)
}



// UserCreate 创建成员
// https://open.work.weixin.qq.com/api/doc/90000/90135/90195
func UserCreate(c *gin.Context) {
  name := c.DefaultQuery("name", "TestUserName")
  userId := c.DefaultQuery("userId", "TestUserId")
  mobile := c.DefaultQuery("mobile", "13184237578")
  res, err := services.WeComApp.User.Create(&request.RequestUserDetail{
    UserID:     userId,
    Name:       name,
    Mobile:     mobile,
    Department: []int{0},
  })

  if err != nil {
    panic(err)
  }

  c.JSON(200, res)
}

// UserGet 读取成员
// https://open.work.weixin.qq.com/api/doc/90000/90135/90196
func UserGet(c *gin.Context) {
  userId := c.DefaultQuery("userId", "TestUserId")
  res, err := services.WeComApp.User.Get(userId)
  if err != nil {
    panic(err)
  }

  c.JSON(200, res)
}

// UserUpdate 更新成员
// https://open.work.weixin.qq.com/api/doc/90000/90135/90197
func UserUpdate(c *gin.Context) {
  userId := c.DefaultQuery("userId", "TestUserId")
  name := c.DefaultQuery("name", "TestUserName2")
  res, err := services.WeComApp.User.Update(&request.RequestUserDetail{
    UserID: userId,
    Name:   name,
  })
  if err != nil {
    panic(err)
  }

  c.JSON(200, res)
}


// UserDelete 删除成员
// https://open.work.weixin.qq.com/api/doc/90000/90135/90198
func UserDelete(c *gin.Context) {
  userId := c.DefaultQuery("userId", "TestUserId")
  res, err := services.WeComApp.User.Delete(userId)
  if err != nil {
    panic(err)
  }
  c.JSON(200, res)
}


// UserBatchDelete 批量删除成员
// https://open.work.weixin.qq.com/api/doc/90000/90135/90199
func UserBatchDelete(c *gin.Context) {
  userId := c.DefaultQuery("userId", "TestUserId")
  res, err := services.WeComApp.User.BatchDelete([]string{userId})
  if err != nil {
    panic(err)
  }
  c.JSON(200, res)
}

// UserSimpleList 获取部门成员
// https://open.work.weixin.qq.com/api/doc/90000/90135/90200
func UserSimpleList(c *gin.Context)  {
  departmentId := c.DefaultQuery("departmentId", "0")
  id, _ := strconv.Atoi(departmentId)
  res, err := services.WeComApp.User.GetDepartmentUsers(id, 1)
  if err != nil {
    panic(err)
  }
  c.JSON(200, res)
}

// UserDetailList 获取部门成员详情
// https://open.work.weixin.qq.com/api/doc/90000/90135/90201
func UserDetailList(c *gin.Context)  {
  departmentId := c.DefaultQuery("departmentId", "0")
  id, _ := strconv.Atoi(departmentId)
  res, err := services.WeComApp.User.GetDetailedDepartmentUsers(id, 1)
  if err != nil {
    panic(err)
  }
  c.JSON(200, res)
}

// UserIdToOpenId UserId转OpenId
// https://open.work.weixin.qq.com/api/doc/90000/90135/90202
func UserIdToOpenId(c *gin.Context)  {
  userId := c.DefaultQuery("userId", "walle")
  res, err := services.WeComApp.User.UserIdToOpenid(userId)
  if err != nil {
    panic(err)
  }
  c.JSON(200, res)
}

// OpenIdToUserId OpenId转UserId
// https://open.work.weixin.qq.com/api/doc/90000/90135/90202
func OpenIdToUserId(c *gin.Context)  {
  openId := c.DefaultQuery("openId", "walle")
  res, err := services.WeComApp.User.OpenIDToUserID(openId)
  if err != nil {
    panic(err)
  }
  c.JSON(200, res)
}

// AuthAccept 二次验证
// https://open.work.weixin.qq.com/api/doc/90000/90135/90203
func AuthAccept(c *gin.Context)  {
  userId := c.DefaultQuery("userId", "walle")
  res, err := services.WeComApp.User.Accept(userId)
  if err != nil {
    panic(err)
  }
  c.JSON(200, res)
}

// BatchInvite 邀请成员
// https://open.work.weixin.qq.com/api/doc/90000/90135/90975
func BatchInvite(c *gin.Context)  {
  userId := c.DefaultQuery("userId", "TestUserId")
  res, err := services.WeComApp.User.Invite(&power.HashMap{
    "user": []string{userId},
  })
  if err != nil {
    panic(err)
  }
  c.JSON(200, res)
}

// GetJoinQrCode 获取加入企业二维码
// https://open.work.weixin.qq.com/api/doc/90000/90135/91714
func GetJoinQrCode(c *gin.Context)  {
  res, err := services.WeComApp.User.GetInvitationQrCode(3)
  if err != nil {
    panic(err)
  }
  c.JSON(200, res)
}

// GetActiveStat 获取企业活跃成员数
// https://open.work.weixin.qq.com/api/doc/90000/90135/92714
func GetActiveStat(c *gin.Context)  {
  date := c.DefaultQuery("date", "2021-09-13")
  res, err := services.WeComApp.User.GetActiveStat(date)
  if err != nil {
    panic(err)
  }
  c.JSON(200, res)
}

const defaultTagId = 100
// TagCreate 创建标签
// https://open.work.weixin.qq.com/api/doc/90000/90135/90210
func TagCreate(c *gin.Context)  {
  res, err := services.WeComApp.UserTag.Create("TestTag", defaultTagId)
  if err != nil {
    panic(err)
  }
  c.JSON(200, res)
}


// TagUpdate 创建标签
// https://open.work.weixin.qq.com/api/doc/90000/90135/90210
func TagUpdate(c *gin.Context)  {
  res, err := services.WeComApp.UserTag.Update("TestTag1", defaultTagId)
  if err != nil {
    panic(err)
  }
  c.JSON(200, res)
}

// TagDelete 删除标签
// https://open.work.weixin.qq.com/api/doc/90000/90135/90212
func TagDelete(c *gin.Context)  {
  res, err := services.WeComApp.UserTag.Delete(defaultDepartmentId)
  if err != nil {
    panic(err)
  }

  c.JSON(200, res)
}


// TagList 获取标签列表
// https://open.work.weixin.qq.com/api/doc/90000/90135/90216
func TagList(c *gin.Context)  {
  res, err := services.WeComApp.UserTag.List()
  if err != nil {
    panic(err)
  }
  c.JSON(200, res)
}

// TagUserGet 获取标签成员
// https://open.work.weixin.qq.com/api/doc/90000/90135/90213
func TagUserGet(c *gin.Context)  {
  tagId := c.DefaultQuery("tagId", string(rune(defaultTagId)))
  res, err := services.WeComApp.UserTag.Get(tagId)
  if err != nil {
    panic(err)
  }
  c.JSON(200, res)
}

// TagUserAdd 增加标签成员
func TagUserAdd(c *gin.Context)  {
  //tagId := c.DefaultQuery("tagId", string(rune(defaultTagId)))
  userId := c.DefaultQuery("tagId", "walle")
  res, err := services.WeComApp.UserTag.TagUsers(defaultTagId, []string{userId})
  if err != nil {
    panic(err)
  }
  c.JSON(200, res)
}

// TagUserDel 删除标签成员
// https://open.work.weixin.qq.com/api/doc/90000/90135/90215
func TagUserDel(c *gin.Context)  {
  //tagId := c.DefaultQuery("tagId", string(rune(defaultTagId)))
  userId := c.DefaultQuery("tagId", "walle")
  res, err := services.WeComApp.UserTag.TagUsers(defaultTagId, []string{userId})
  if err != nil {
    panic(err)
  }
  c.JSON(200, res)
}

// BatchSyncUser 增量更新成员
// TODO
func BatchSyncUser(c *gin.Context) {
  //mediaId := c.DefaultQuery("mediaId", "xxxxxxx")
  //res, err := services.WeComApp.UserBatchJobs
  //if err != nil {
  // panic(err)
  //}
  //c.JSON(200, res)
}
