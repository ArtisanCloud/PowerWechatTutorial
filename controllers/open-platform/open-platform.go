package open_platform

import (
	"github.com/ArtisanCloud/PowerLibs/v2/fmt"
	"github.com/ArtisanCloud/PowerLibs/v2/object"
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
	a.Create()
}

func APIOpenPlatformPreAuthCode(context *gin.Context) {
	ticket := context.Query("ticket")

	err := services.OpenPlatformApp.VerifyTicket.SetTicket(ticket)
	if err != nil {
		panic(err)
	}

	token, err := services.OpenPlatformApp.AccessToken.GetToken(false)
	if err != nil {
		panic(err)
	}
	fmt.Dump(token)

	rs, err := services.OpenPlatformApp.Base.CreatePreAuthorizationCode()
	if err != nil {
		panic(err)
	}
	fmt.Dump(rs)
}

// HandleAuthorize Code换调用凭据信息
func HandleAuthorize(ctx *gin.Context) {
	authCode := ctx.DefaultQuery("authCode", "")
	res, err := services.OpenPlatformApp.Base.HandleAuthorize(authCode)
	if err != nil {
		panic(err)
	}
	ctx.JSON(200, res)
}

func GetFastRegistrationURL(context *gin.Context) {
	url := services.OpenPlatformApp.GetFastRegistrationURL("https://test.com", &object.StringMap{
		"auth_type": "1",
	})

	fmt.Dump(url)
}

// GetAuthorizer 获取授权方的帐号基本信息
func GetAuthorizer(ctx *gin.Context) {
	appID := ctx.DefaultQuery("app_id", "wx3c7e1c9f9f1f9f9f")
	res, err := services.OpenPlatformApp.Base.GetAuthorizer(appID)
	if err != nil {
		panic(err)
	}
	ctx.JSON(200, res)
}

// GetAuthorizerOption 获取授权方的选项设置信息
func GetAuthorizerOption(ctx *gin.Context) {
	appID := ctx.DefaultQuery("app_id", "wx3c7e1c9f9f1f9f9f")
	name := ctx.DefaultQuery("name", "location_report")
	res, err := services.OpenPlatformApp.Base.GetAuthorizerOption(appID, name)
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
	res, err := services.OpenPlatformApp.Base.SetAuthorizerOption(appID, name, value)
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
	res, err := services.OpenPlatformApp.Base.GetAuthorizers(offsetInt, countInt)
	if err != nil {
		panic(err)
	}
	ctx.JSON(200, res)
}
