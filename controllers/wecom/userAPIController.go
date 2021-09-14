package wecom

import (
	"github.com/ArtisanCloud/power-wechat/src/work/user/request"
	"github.com/gin-gonic/gin"
	"power-wechat-tutorial/services"
)

// UserCreate 创建成员
// https://open.work.weixin.qq.com/api/doc/90000/90135/90195
func APIUserCreate(c *gin.Context) {
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
func APIUserGet(c *gin.Context) {
	userId := c.DefaultQuery("userId", "TestUserId")
	res, err := services.WeComApp.User.Get(userId)
	if err != nil {
		panic(err)
	}

	c.JSON(200, res)
}

// UserUpdate 更新成员
// https://open.work.weixin.qq.com/api/doc/90000/90135/90197
func APIUserUpdate(c *gin.Context) {
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
func APIUserDelete(c *gin.Context) {
	userId := c.DefaultQuery("userId", "TestUserId")
	res, err := services.WeComApp.User.Delete(userId)
	if err != nil {
		panic(err)
	}
	c.JSON(200, res)
}

// UserBatchDelete 批量删除成员
// https://open.work.weixin.qq.com/api/doc/90000/90135/90199
func APIUserBatchDelete(c *gin.Context) {
	userId := c.DefaultQuery("userId", "TestUserId")
	res, err := services.WeComApp.User.BatchDelete([]string{userId})
	if err != nil {
		panic(err)
	}
	c.JSON(200, res)
}