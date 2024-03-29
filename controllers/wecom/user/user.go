package user

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/user/request"
	"github.com/gin-gonic/gin"
	"net/http"
	"power-wechat-tutorial/services"
	"strconv"
)

// UserCreate 创建成员
// https://open.work.weixin.qq.com/api/doc/90000/90135/90195
func APIUserCreate(c *gin.Context) {
	userDetail, _ := c.Get("params")

	res, err := services.WeComContactApp.User.Create(c.Request.Context(), userDetail.(*request.RequestUserDetail))

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)
}

// UserGet 读取成员
// https://open.work.weixin.qq.com/api/doc/90000/90135/90196
func APIUserGet(c *gin.Context) {
	userId := c.DefaultQuery("userId", "TestUserId")
	res, err := services.WeComContactApp.User.Get(c.Request.Context(), userId)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)
}

// UserUpdate 更新成员
// https://open.work.weixin.qq.com/api/doc/90000/90135/90197
func APIUserUpdate(c *gin.Context) {
	userId := c.DefaultQuery("userId", "TestUserId")
	name := c.DefaultQuery("name", "TestUserName2")
	res, err := services.WeComContactApp.User.Update(c.Request.Context(), &request.RequestUserDetail{
		Userid: userId,
		Name:   name,
	})
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)
}

// UserDelete 删除成员
// https://open.work.weixin.qq.com/api/doc/90000/90135/90198
func APIUserDelete(c *gin.Context) {
	userId := c.DefaultQuery("userId", "TestUserId")
	res, err := services.WeComContactApp.User.Delete(c.Request.Context(), userId)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, res)
}

// UserBatchDelete 批量删除成员
// https://open.work.weixin.qq.com/api/doc/90000/90135/90199
func APIUserBatchDelete(c *gin.Context) {
	userId := c.DefaultQuery("userId", "TestUserId")
	res, err := services.WeComContactApp.User.BatchDelete(c.Request.Context(), []string{userId})
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, res)
}

// UserSimpleList 获取部门成员
// https://open.work.weixin.qq.com/api/doc/90000/90135/90200
func APIUserSimpleList(c *gin.Context) {
	departmentId := c.DefaultQuery("departmentId", "0")
	id, _ := strconv.Atoi(departmentId)
	res, err := services.WeComContactApp.User.GetDepartmentUsers(c.Request.Context(), id, 1)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, res)
}

// UserDetailList 获取部门成员详情
// https://open.work.weixin.qq.com/api/doc/90000/90135/90201
func APIUserDetailList(c *gin.Context) {
	departmentId := c.DefaultQuery("departmentId", "0")
	id, _ := strconv.Atoi(departmentId)
	res, err := services.WeComContactApp.User.GetDetailedDepartmentUsers(c.Request.Context(), id, 1)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, res)
}

// UserIdToOpenId UserId转OpenId
// https://open.work.weixin.qq.com/api/doc/90000/90135/90202
func APIUserUserIDToOpenID(c *gin.Context) {
	userId := c.DefaultQuery("userId", "walle")
	res, err := services.WeComContactApp.User.UserIdToOpenID(c.Request.Context(), userId)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, res)
}

// OpenIdToUserId OpenId转UserId
// https://open.work.weixin.qq.com/api/doc/90000/90135/90202
func APIUserOpenIDToUserID(c *gin.Context) {
	openId := c.DefaultQuery("openId", "walle")
	res, err := services.WeComContactApp.User.OpenIDToUserID(c.Request.Context(), openId)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, res)
}

// AuthAccept 二次验证
// https://open.work.weixin.qq.com/api/doc/90000/90135/90203
func APIUserAuthAccept(c *gin.Context) {
	userId := c.DefaultQuery("userId", "walle")
	res, err := services.WeComContactApp.User.Accept(c.Request.Context(), userId)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, res)
}

// BatchInvite 邀请成员
// https://open.work.weixin.qq.com/api/doc/90000/90135/90975
func APIUserBatchInvite(c *gin.Context) {
	userId := c.DefaultQuery("userId", "TestUserId")
	res, err := services.WeComContactApp.User.Invite(c.Request.Context(), &power.HashMap{
		"user": []string{userId},
	})
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, res)
}

// GetJoinQrCode 获取加入企业二维码
// https://open.work.weixin.qq.com/api/doc/90000/90135/91714
func APIUserGetJoinQrCode(c *gin.Context) {
	res, err := services.WeComContactApp.User.GetJoinQrCode(c.Request.Context(), 3)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, res)
}

// GetActiveStat 获取企业活跃成员数
// https://open.work.weixin.qq.com/api/doc/90000/90135/92714
func APIUserGetActiveStat(c *gin.Context) {
	date := c.DefaultQuery("date", "2021-09-13")
	res, err := services.WeComContactApp.User.GetActiveStat(c.Request.Context(), date)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, res)
}

// 增量更新成员
// https://work.weixin.qq.com/api/doc/90000/90135/90980
func APIBatchSyncUser(c *gin.Context) {
	mediaID := c.DefaultQuery("mediaID", "xxx")
	callback := &power.StringMap{
		"url":            "xxx",
		"token":          "xxx",
		"encodingaeskey": "xxx",
	}

	res, err := services.WeComContactApp.UserBatchJobs.SyncUser(c.Request.Context(), mediaID, true, callback)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, res)
}

// 全量覆盖成员
// https://work.weixin.qq.com/api/doc/90000/90135/90981
func APIBatchReplaceUser(c *gin.Context) {
	callback := &power.StringMap{
		"url":            "xxx",
		"token":          "xxx",
		"encodingaeskey": "xxx",
	}

	res, err := services.WeComContactApp.UserBatchJobs.ReplaceUser(c.Request.Context(), "xxx", true, callback)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, res)
}

// 全量覆盖部门
// https://work.weixin.qq.com/api/doc/90000/90135/90982
func APIBatchReplaceParty(c *gin.Context) {
	callback := &power.StringMap{
		"url":            "xxx",
		"token":          "xxx",
		"encodingaeskey": "xxx",
	}

	res, err := services.WeComContactApp.UserBatchJobs.ReplaceParty(c.Request.Context(), "xxx", true, callback)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, res)
}

// 获取异步任务结果
// https://work.weixin.qq.com/api/doc/90000/90135/90983
func APIBatchGetResult(c *gin.Context) {
	jobID, exist := c.GetQuery("jobID")
	if !exist {
		panic("parameter job id expected")
	}

	res, err := services.WeComContactApp.UserBatchJobs.GetBatchResult(c.Request.Context(), jobID)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, res)
}

// APIExportSimpleUser 获取异步任务结果
// https://open.work.weixin.qq.com/api/doc/90000/90135/94849
func APIExportSimpleUser(c *gin.Context) {

	encodingAESKey := c.DefaultQuery("encryptedMsgHash", "hsSuSUsePBqSQw2rYMtf9Nvha603xX8f2BMQBcYRoJiMNwOqt/UEhrqekebG5ar0LFNAm5MD4Uz6zorRwiXJwbySJ/FEJHav4NsobBIU1PwdjbJWVQLFy7+YFkHB32OnQXWMh6ugW7Dyk2KS5BXp1f5lniKPp1KNLyNLlFlNZ2mgJCJmWvHj5AI7BLpWwoRvqRyZvVXo+9FsWqvBdxmAPA==")
	blockSize := int64(123)
	res, err := services.WeComContactApp.UserExportJobs.SimpleUser(c.Request.Context(), encodingAESKey, blockSize)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, res)
}

// APIExportUser 导出成员详情
// https://open.work.weixin.qq.com/api/doc/90000/90135/94851
func APIExportUser(c *gin.Context) {
	encodingAESKey := c.DefaultQuery("encryptedMsgHash", "IJUiXNpvGbODwKEBSEsAeOAPAhkqHqNCF6g19t9wfg2")
	blockSize := int64(1000000)

	res, err := services.WeComContactApp.UserExportJobs.User(c.Request.Context(), encodingAESKey, blockSize)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, res)
}

// APIExportDepartment 异步导出部门
// https://open.work.weixin.qq.com/api/doc/90000/90135/94852
func APIExportDepartment(c *gin.Context) {
	encodingAESKey := c.DefaultQuery("encryptedMsgHash", "IJUiXNpvGbODwKEBSEsAeOAPAhkqHqNCF6g19t9wfg2")
	blockSize := int64(1000000)

	res, err := services.WeComContactApp.UserExportJobs.Department(c.Request.Context(), encodingAESKey, blockSize)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, res)
}

// APIExportTagUser 异步导出标签成员
// https://open.work.weixin.qq.com/api/doc/90000/90135/94853
func APIExportTagUser(c *gin.Context) {

	tagID := 1
	encodingAESKey := c.DefaultQuery("encryptedMsgHash", "IJUiXNpvGbODwKEBSEsAeOAPAhkqHqNCF6g19t9wfg2")
	blockSize := int64(1000000)

	res, err := services.WeComContactApp.UserExportJobs.TagUser(c.Request.Context(), tagID, encodingAESKey, blockSize)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, res)
}

// APIExportGetResult 获取导出结果
// https://open.work.weixin.qq.com/api/doc/90000/90135/94854
func APIExportGetResult(c *gin.Context) {
	jobID := c.DefaultQuery("jobID", "jobid_xxxxxxxxxxxxxxx")

	res, err := services.WeComContactApp.UserExportJobs.GetExportResult(c.Request.Context(), jobID)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, res)
}
