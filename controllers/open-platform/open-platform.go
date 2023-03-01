package open_platform

import (
	"github.com/ArtisanCloud/PowerLibs/v3/fmt"
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/gin-gonic/gin"
	"power-wechat-tutorial/services"
	"strconv"
)

func GetPreAuthorizationUrl(ctx *gin.Context) {
	appID := ctx.DefaultQuery("app_id", "wx3c7e1c9f9f1f9f9f")
	refreshToken := ctx.DefaultQuery("refresh_token", "wx3c7e1c9f9f1f9f9f")
	//services.OpenPlatformApp.GetMobilePreAuthorizationURL()
	account, _ := services.OpenPlatformApp.OfficialAccount(appID, refreshToken, nil)
	a := account.Account
	a.Create(ctx.Request.Context())
}

func APIOpenPlatformPreAuthCode(ctx *gin.Context) {
	// 如果之前在回调的时候，已经将获取的ComponentVerifyTicket通过SetTicket缓存后，则可以注销代码，跳过这个环节
	ticket := ctx.Query("ticket")
	err := services.OpenPlatformApp.VerifyTicket.SetTicket(ticket)
	if err != nil {
		panic(err)
	}

	rs, err := services.OpenPlatformApp.Base.CreatePreAuthorizationCode(ctx.Request.Context())
	if err != nil {
		panic(err)
	}
	fmt.Dump(rs)
}

func GetFastRegistrationURL(ctx *gin.Context) {
	// 如果之前在回调的时候，已经将获取的ComponentVerifyTicket通过SetTicket缓存后，则可以注销代码，跳过这个环节
	ticket, _ := services.OpenPlatformApp.VerifyTicket.GetTicket()
	if ticket == "" {
		ticket = ctx.Query("ticket")
		err := services.OpenPlatformApp.VerifyTicket.SetTicket(ticket)
		if err != nil {
			panic(err)
		}
	}

	url, err := services.OpenPlatformApp.GetFastRegistrationURL(ctx.Request.Context(), "https://michael.debug.artisancloud.cn/wx/api/openplatform/callback", &object.StringMap{
		"auth_type": "1",
	})

	if err != nil {
		panic(err)
	}

	ctx.JSON(200, url)
}

// HandleAuthorize Code换调用凭据信息
func HandleAuthorize(ctx *gin.Context) {
	authCode := ctx.DefaultQuery("auth_code", "")
	res, err := services.OpenPlatformApp.Base.HandleAuthorize(ctx.Request.Context(), authCode)
	if err != nil {
		panic(err)
	}
	ctx.JSON(200, res)
}

// GetAuthorizer 获取授权方的帐号基本信息
func GetAuthorizer(ctx *gin.Context) {
	appID := ctx.DefaultQuery("app_id", "wx3c7e1c9f9f1f9f9f")
	res, err := services.OpenPlatformApp.Base.GetAuthorizer(ctx.Request.Context(), appID)
	if err != nil {
		panic(err)
	}
	ctx.JSON(200, res)
}

// GetAuthorizerOption 获取授权方的选项设置信息
func GetAuthorizerOption(ctx *gin.Context) {
	appID := ctx.DefaultQuery("app_id", "wx3c7e1c9f9f1f9f9f")
	name := ctx.DefaultQuery("name", "location_report")
	res, err := services.OpenPlatformApp.Base.GetAuthorizerOption(ctx.Request.Context(), appID, name)
	if err != nil {
		panic(err)
	}
	ctx.JSON(200, res)
}

// SetAuthorizerOption 设置授权方的选项信息
func SetAuthorizerOption(ctx *gin.Context) {
	appID := ctx.DefaultQuery("app_id", "wx3c7e1c9f9f1f9f9f")
	name := ctx.DefaultQuery("name", "location_report")
	value := ctx.DefaultQuery("value", "1")
	res, err := services.OpenPlatformApp.Base.SetAuthorizerOption(ctx.Request.Context(), appID, name, value)
	if err != nil {
		panic(err)
	}
	ctx.JSON(200, res)
}

// GetAuthorizers 获取授权方的帐号列表
func GetAuthorizers(ctx *gin.Context) {
	offset := ctx.DefaultQuery("offset", "0")
	count := ctx.DefaultQuery("count", "10")
	offsetInt, _ := strconv.Atoi(offset)
	countInt, _ := strconv.Atoi(count)
	res, err := services.OpenPlatformApp.Base.GetAuthorizers(ctx.Request.Context(), offsetInt, countInt)
	if err != nil {
		panic(err)
	}
	ctx.JSON(200, res)
}
