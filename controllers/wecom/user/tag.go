package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"power-wechat-tutorial/services"
)

const defaultTagId = 100

// APITagCreate 创建标签
// https://open.work.weixin.qq.com/api/doc/90000/90135/90210
func APITagCreate(c *gin.Context) {
	res, err := services.WeComApp.UserTag.Create(c.Request.Context(), "TestTag", defaultTagId)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, res)
}

// APITagUpdate 更新标签名字
// https://open.work.weixin.qq.com/api/doc/90000/90135/90211
func APITagUpdate(c *gin.Context) {
	res, err := services.WeComApp.UserTag.Update(c.Request.Context(), "TestTag1", defaultTagId)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, res)
}

// APITagDelete 删除标签
// https://open.work.weixin.qq.com/api/doc/90000/90135/90212
func APITagDelete(c *gin.Context) {
	res, err := services.WeComApp.UserTag.Delete(c.Request.Context(), defaultTagId)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)
}

// TagList 获取标签列表
// https://open.work.weixin.qq.com/api/doc/90000/90135/90216
func APITagList(c *gin.Context) {
	res, err := services.WeComApp.UserTag.List(c.Request.Context())
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, res)
}

// TagUserGet 获取标签成员
// https://open.work.weixin.qq.com/api/doc/90000/90135/90213
func APITagUserGet(c *gin.Context) {
	res, err := services.WeComApp.UserTag.Get(c.Request.Context(), defaultTagId)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, res)
}

// TagUserAdd 增加标签成员
// https://open.work.weixin.qq.com/api/doc/90000/90135/90214
func APITagUserAdd(c *gin.Context) {
	//tagId := c.DefaultQuery("tagId", string(rune(defaultTagId)))
	userId := c.DefaultQuery("tagId", "walle")
	res, err := services.WeComApp.UserTag.TagUsers(c.Request.Context(), defaultTagId, []string{userId})
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, res)
}

// TagUserDel 删除标签成员
// https://open.work.weixin.qq.com/api/doc/90000/90135/90215
func APITagUserDel(c *gin.Context) {
	//tagId := c.DefaultQuery("tagId", string(rune(defaultTagId)))
	userId := c.DefaultQuery("tagId", "walle")
	res, err := services.WeComApp.UserTag.TagUsers(c.Request.Context(), defaultTagId, []string{userId})
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, res)
}
