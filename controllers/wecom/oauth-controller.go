package wecom

import (
  "github.com/ArtisanCloud/PowerWeChat/src/kernel/power"
  "github.com/gin-gonic/gin"
  "log"
  "net/http"
  "power-wechat-tutorial/services"
)

func WebAuthorizeUser(ctx *gin.Context) {

  // $callbackUrl 为授权回调地址
  callbackUrl := services.WeComApp.GetConfig().GetString("oauth.callback", "artisan-cloud.com") +"/callback/authorized/user"// 需设置可信域名
  services.WeComApp.OAuth.Provider.WithRedirectURL(callbackUrl)

  // 返回一个 redirect 实例
  redirectURL, _ := services.WeComApp.OAuth.Provider.GetAuthURL()

  log.Println("redirectURL: ", redirectURL)

  // 请在企业微信客户端打开链接
  //http.Redirect(ctx.Writer, ctx.Request, redirectURL, http.StatusFound)
  ctx.Redirect(http.StatusFound, redirectURL)
}

func WebAuthorizedUser(ctx *gin.Context) {

  code := ctx.Query("code")
  user, err := services.WeComApp.OAuth.Provider.ContactFromCode(code)

  if err != nil {
    panic(err)
  }

  rs := &power.HashMap{
    "deviceID": user.GetDeviceID(),
    "userID":   user.GetID(),
    "openID":  user.GetOpenID(),
  }

  //// 正常返回json
  ctx.JSON(http.StatusOK, rs)
}

func WebAuthorizeContact(ctx *gin.Context) {

  // $callbackUrl 为授权回调地址
  // 需设置可信域名
  callbackUrl := services.WeComApp.GetConfig().GetString("oauth.callback", "")+"/callback/authorized/contact"

  // 返回一个 redirect 实例
  services.WeComApp.OAuth.Provider.WithRedirectURL(callbackUrl)
  redirectURL, _ := services.WeComApp.OAuth.Provider.GetQrConnectURL()

  log.Println("redirectURL: ", redirectURL)

  // 直接跳转到企业微信授权
  http.Redirect(ctx.Writer, ctx.Request, redirectURL, http.StatusFound)

}

func WebAuthorizedContact(ctx *gin.Context) {

  code := ctx.Query("code")
  user, err := services.WeComApp.OAuth.Provider.Detailed().ContactFromCode(code)
  if err != nil {
    panic(err)
  }

  rs := &power.HashMap{
    "openID": user.GetOpenID(),
    "userID": user.GetID(),
    "name": user.GetName(),
    "avatar": user.GetAvatar(),
  }

  //// 正常返回json
  ctx.JSON(http.StatusOK, rs)

}
