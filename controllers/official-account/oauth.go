package official_account

import (
	"github.com/ArtisanCloud/PowerLibs/v3/fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"power-wechat-tutorial/services"
)

func GetAuthCode(ctx *gin.Context) {

	//result, err := services.OfficialAccountApp.JSSDK.ConfigSignature(ctx, "text", "", 0)
	//if err != nil {
	//	panic(err.Error())
	//}
	//ctx.JSON(http.StatusOK, gin.H{"result": result})
	//return
	code := ctx.Query("code")
	state := ctx.Query("state")

	ctx.JSON(http.StatusOK, gin.H{"code": code, "state": state})
}

func UserFromCode(ctx *gin.Context) {
	code := ctx.Query("code")
	services.OfficialAccountApp.OAuth.SetScopes([]string{"snsapi_base"})
	user, err := services.OfficialAccountApp.OAuth.UserFromCode(code)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func UserFromToken(ctx *gin.Context) {
	token := ctx.Query("token")
	openID := ctx.Query("openID")
	user, err := services.OfficialAccountApp.OAuth.UserFromToken(token, openID)
	rsToken := user.GetTokenResponse()
	fmt.Dump(rsToken, (*rsToken)["openid"])
	log.Println(err)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, user)
}
