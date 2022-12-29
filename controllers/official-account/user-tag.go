package official_account

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"power-wechat-tutorial/services"
)

// UserTagList 获取用户标签列表
func GetUserTagList(ctx *gin.Context) {
	data, err := services.OfficialAccountApp.UserTag.List(ctx.Request.Context())
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, data)
}

// UserTagCreate 创建用户标签
func UserTagCreate(ctx *gin.Context) {
	tagName, _ := ctx.GetPostForm("tagName")
	data, err := services.OfficialAccountApp.UserTag.Create(ctx.Request.Context(), tagName)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, data)
}

// UserTagUpdate 更新用户标签
func UserTagUpdate(ctx *gin.Context) {
	tagID, _ := ctx.GetPostForm("tagID")
	tagName, _ := ctx.GetPostForm("tagName")
	data, err := services.OfficialAccountApp.UserTag.Update(ctx.Request.Context(), tagID, tagName)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, data)
}

// UserTagDelete 删除用户标签
func UserTagDelete(ctx *gin.Context) {
	tagID, _ := ctx.GetPostForm("tagID")
	data, err := services.OfficialAccountApp.UserTag.Delete(ctx.Request.Context(), tagID)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, data)
}

// GetUserTagsByOpenID 获取用户身上的标签列表
func GetUserTagsByOpenID(ctx *gin.Context) {
	openID := ctx.Query("openID")
	data, err := services.OfficialAccountApp.UserTag.UserTags(ctx.Request.Context(), openID)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, data)
}

// GetUsersOfTag 获取标签下粉丝列表
func GetUsersOfTag(ctx *gin.Context) {
	tagID := ctx.Query("tagID")
	nextOpenID := ctx.Query("nextOpenID")
	data, err := services.OfficialAccountApp.UserTag.UsersOfTag(ctx.Request.Context(), tagID, nextOpenID)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, data)
}

// UserTagBatchTagUsers 批量为用户打标签
func UserTagBatchTagUsers(ctx *gin.Context) {
	openID := ctx.Query("openID")
	tagID := ctx.Query("tagID")
	data, err := services.OfficialAccountApp.UserTag.TagUsers(ctx.Request.Context(), []string{openID}, tagID)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, data)
}

// UserTagBatchUnTagUsers 批量为用户取消标签
func UserTagBatchUnTagUsers(ctx *gin.Context) {
	openID := ctx.Query("openID")
	tagID := ctx.Query("tagID")
	data, err := services.OfficialAccountApp.UserTag.UntagUsers(ctx.Request.Context(), []string{openID}, tagID)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, data)
}
