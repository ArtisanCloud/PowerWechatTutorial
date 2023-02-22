package open_platform

import (
	"github.com/ArtisanCloud/PowerLibs/v3/fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"power-wechat-tutorial/services"
)

// GetAuthorizerOfficialAccount 获取授权方实例 - 公众号
func GetAuthorizerOfficialAccount(ctx *gin.Context) {
	appID := ctx.DefaultQuery("app_id", "")
	refreshToken := ctx.DefaultQuery("refresh_token", "")
	if appID == "" || refreshToken == "" {
		panic("app_id or refresh_token is empty ")
	}
	officialAccount, err := services.OpenPlatformApp.OfficialAccount(appID, refreshToken, nil)
	if err != nil {
		panic(err)
	}
	ipList, err := officialAccount.OfficialAccount.Base.GetCallbackIP(ctx)
	if err != nil {
		panic(err)
	}
	ctx.JSON(http.StatusOK, gin.H{"msg": "success", "data": ipList})

}

// GetAuthorizerOfficialAccountUser 获取授权方实例 - 公众号
func GetAuthorizerOfficialAccountUser(ctx *gin.Context) {
	appID := ctx.DefaultQuery("app_id", "")
	code := ctx.DefaultQuery("code", "")
	refreshToken := ctx.DefaultQuery("refresh_token", "")
	if appID == "" || refreshToken == "" || code == "" {
		panic("app_id or refresh_token is empty ")
	}
	officialAccount, err := services.OpenPlatformApp.OfficialAccount(appID, refreshToken, nil)
	if err != nil {
		panic(err)
	}

	user, err := officialAccount.OfficialAccount.JSSDK.ConfigSignature(ctx, "ttt", "", 0)
	if err != nil {
		panic(err)
	}
	ctx.JSON(http.StatusOK, gin.H{"msg": "success", "data": user})

}

// GetAuthorizerMiniProgram 获取授权方实例 - 小程序
func GetAuthorizerMiniProgram(ctx *gin.Context) {
	appID := ctx.DefaultQuery("app_id", "")
	refreshToken := ctx.DefaultQuery("refresh_token", "")
	if appID == "" || refreshToken == "" {
		panic("app_id or refresh_token is empty ")
	}

	miniProgram, err := services.OpenPlatformApp.MiniProgram(appID, refreshToken, nil)
	if err != nil {
		panic(err)
	}
	ctx.JSON(http.StatusOK, gin.H{"msg": "success", "data": miniProgram})
}

// AuthorizerManage 创建开放平台账号
// 并绑定公众号或小程序
func AuthorizerAccountCreate(ctx *gin.Context) {
	appID := ctx.DefaultQuery("app_id", "")
	refreshToken := ctx.DefaultQuery("refresh_token", "")
	if appID == "" || refreshToken == "" {
		panic("app_id or refresh_token is empty ")
	}

	officialAccount, err := services.OpenPlatformApp.OfficialAccount(appID, refreshToken, nil)
	if err != nil {
		panic(err)
	}
	// 代公众号实现业务
	account := officialAccount.Account
	result, err := account.Create(ctx.Request.Context())
	if err != nil {
		panic(err)
	}
	ctx.JSON(http.StatusOK, gin.H{"msg": "success", "data": result})
}

// AuthorizerAccountBind 将公众号或小程序绑定到指定开放平台帐号下
func AuthorizerAccountBind(ctx *gin.Context) {
	appID := ctx.DefaultQuery("app_id", "")
	refreshToken := ctx.DefaultQuery("refresh_token", "")
	openAppID := ctx.DefaultQuery("open_app_id", "")

	if appID == "" || refreshToken == "" || openAppID == "" {
		panic("app_id or refresh_token is empty ")
	}

	officialAccount, err := services.OpenPlatformApp.OfficialAccount(appID, refreshToken, nil)
	if err != nil {
		panic(err)
	}
	// 将公众号或小程序绑定到指定开放平台帐号下
	account := officialAccount.Account
	result, err := account.BindTo(ctx.Request.Context(), openAppID)
	if err != nil {
		panic(err)
	}
	ctx.JSON(http.StatusOK, gin.H{"msg": "success", "data": result})
}

// AuthorizerAccountUnbind 将公众号或小程序从指定开放平台帐号下解绑
func AuthorizerAccountUnbind(ctx *gin.Context) {
	appID := ctx.DefaultQuery("app_id", "")
	refreshToken := ctx.DefaultQuery("refresh_token", "")
	openAppID := ctx.DefaultQuery("open_app_id", "")
	if appID == "" || refreshToken == "" || openAppID == "" {
		panic("app_id or refresh_token is empty ")
	}

	officialAccount, err := services.OpenPlatformApp.OfficialAccount(appID, refreshToken, nil)
	if err != nil {
		panic(err)
	}
	account := officialAccount.Account

	// 将公众号或小程序从指定开放平台帐号下解绑
	result, err := account.UnbindFrom(ctx.Request.Context(), openAppID)
	if err != nil {
		panic(err)
	}
	ctx.JSON(http.StatusOK, gin.H{"msg": "success", "data": result})
}

// AuthorizerAccountGet 获取公众号/小程序所绑定的开放平台帐号
func AuthorizerAccountGet(ctx *gin.Context) {
	appID := ctx.DefaultQuery("app_id", "")
	refreshToken := ctx.DefaultQuery("refresh_token", "")
	if appID == "" || refreshToken == "" {
		panic("app_id or refresh_token is empty ")
	}
	officialAccount, err := services.OpenPlatformApp.OfficialAccount(appID, refreshToken, nil)
	if err != nil {
		panic(err)
	}
	account := officialAccount.Account

	// 获取开放平台帐号下绑定的公众号或小程序
	result, err := account.GetBinding(ctx.Request.Context())
	if err != nil {
		panic(err)
	}
	ctx.JSON(http.StatusOK, gin.H{"msg": "success", "data": result})
}

// OfficialAccountDemo 调用授权方业务例子
func OfficialAccountDemo(ctx *gin.Context) {
	appID := ctx.DefaultQuery("app_id", "")
	refreshToken := ctx.DefaultQuery("refresh_token", "")
	if appID == "" || refreshToken == "" {
		panic("app_id or refresh_token is empty ")
	}
	officialAccount, err := services.OpenPlatformApp.OfficialAccount(appID, refreshToken, nil)
	if err != nil {
		panic(err)
	}
	// 获取用户列表
	list, err := officialAccount.User.List(ctx.Request.Context(), "")
	if err != nil {
		panic(err)
	}
	fmt.Dump(list)
	ctx.JSON(http.StatusOK, gin.H{"msg": "success", "data": list})
}

// MiniProgramDemo 调用授权方业务例子
func MiniProgramDemo(ctx *gin.Context) {
	appID := ctx.DefaultQuery("app_id", "")
	refreshToken := ctx.DefaultQuery("refresh_token", "")
	code := ctx.DefaultQuery("code", "")
	if appID == "" || refreshToken == "" || code == "" {
		panic("app_id or refresh_token is empty ")
	}

	miniProgram, err := services.OpenPlatformApp.MiniProgram(appID, refreshToken, nil)
	if err != nil {
		panic(err)
	}
	// 根据 code 获取 session
	data, err := miniProgram.Auth.Session(ctx.Request.Context(), code)
	if err != nil {
		panic(err)
	}
	ctx.JSON(http.StatusOK, gin.H{"msg": "success", "data": data})
}
