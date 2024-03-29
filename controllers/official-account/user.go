package official_account

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/officialAccount/user/request"
	"github.com/gin-gonic/gin"
	"net/http"
	"power-wechat-tutorial/services"
)

// GetUserInfo 获取单个用户信息
func GetUserInfo(ctx *gin.Context) {
	openID := ctx.Query("openID")
	lang := ctx.Query("lang")
	data, err := services.OfficialAccountApp.User.Get(ctx.Request.Context(), openID, lang)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, data)
}

// 获取多个用户信息
func GetBatchUserInfo(ctx *gin.Context) {
	openID := ctx.Query("openID")
	data, err := services.OfficialAccountApp.User.BatchGet(ctx.Request.Context(), &request.RequestBatchGetUserInfo{
		UserList: []*request.UserList{
			{
				Openid: openID,
			},
		},
	})
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, data)
}

// GetUserList 获取用户列表
func GetUserList(ctx *gin.Context) {
	nextOpenId := ctx.Query("nextOpenId")
	data, err := services.OfficialAccountApp.User.List(ctx.Request.Context(), nextOpenId)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, data)
}

// UserRemark 修改用户备注
func UserRemark(ctx *gin.Context) {
	openID := ctx.Query("openID")
	remark := ctx.Query("remark")
	data, err := services.OfficialAccountApp.User.Remark(ctx.Request.Context(), openID, remark)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, data)
}

// UserBlock 拉黑用户
func UserBlock(ctx *gin.Context) {
	openID := ctx.Query("openID")
	data, err := services.OfficialAccountApp.User.Block(ctx.Request.Context(), []string{openID})
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, data)
}

// UserUnBlock 取消拉黑用户
func UserUnBlock(ctx *gin.Context) {
	openID := ctx.Query("openID")
	data, err := services.OfficialAccountApp.User.Unblock(ctx.Request.Context(), []string{openID})
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, data)
}

// GetUserBlacklist 获取用户列表
func GetUserBlacklist(ctx *gin.Context) {
	beginOpenid := ctx.Query("beginOpenid")
	data, err := services.OfficialAccountApp.User.Blacklist(ctx.Request.Context(), beginOpenid)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, data)
}

// UserChangeOpenID 账号迁移 openid 转换
func UserChangeOpenID(ctx *gin.Context) {
	oldAppId := ctx.Query("oldAppId")
	data, err := services.OfficialAccountApp.User.ChangeOpenID(ctx.Request.Context(), oldAppId, []string{})
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, data)
}
