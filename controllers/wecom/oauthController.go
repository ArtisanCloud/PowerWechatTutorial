package wecom

import (
  "github.com/ArtisanCloud/power-wechat/src/kernel/power"
  "github.com/ArtisanCloud/power-wechat/src/work/oauth/request"
  "github.com/gin-gonic/gin"
  "net/http"
  "power-wechat-tutorial/services"
)

func WebAuthorizeUser(context *gin.Context) {

  callbackURI := "/callback/authorized/user"

  // $callbackUrl 为授权回调地址
  callbackUrl := services.WeComApp.GetConfig().GetString("oauth.callback", "artisan-cloud.com") + callbackURI // 需设置可信域名
  services.WeComApp.OAuth.Provider.WithRedirectURL(callbackUrl)
  // 返回一个 redirect 实例
  redirectURL, _ := services.WeComApp.OAuth.Provider.GetAuthURL()

  // 请在企业微信客户端打开链接
  http.Redirect(context.Writer, context.Request, redirectURL, http.StatusFound)

}

func WebAuthorizedUser(context *gin.Context) {

  params, _ := context.Get("params")
  para := params.(request.ParaOAuthCallback)
  cachedToken, err := services.WeComApp.AccessToken.GetToken(false)
  user, err := services.WeComApp.OAuth.Provider.WithApiAccessToken(cachedToken.AccessToken).UserFromCode(para.Code)
  if err != nil {
    panic(err)
  }

  rs := &power.HashMap{
    "deviceID": user.GetDeviceID(),
    "userID":   user.GetID(),
  }

  //// 正常返回json
  context.JSON(http.StatusOK, rs)

}

func WebAuthorizeContact(context *gin.Context) {

  callbackURI := "/callback/authorized/contact"

  // $callbackUrl 为授权回调地址
  callbackUrl := services.WeComApp.GetConfig().GetString("oauth.callback", "artisan-cloud.com") + callbackURI // 需设置可信域名

  // 返回一个 redirect 实例
  services.WeComApp.OAuth.Provider.WithRedirectURL(callbackUrl)
  redirectURL, _ := services.WeComApp.OAuth.Provider.GetQrConnectURL()

  // 直接跳转到企业微信授权
  http.Redirect(context.Writer, context.Request, redirectURL, http.StatusFound)

}

func WebAuthorizedContact(context *gin.Context) {

  params, _ := context.Get("params")
  para := params.(request.ParaOAuthCallback)
  cachedToken, err := services.WeComApp.AccessToken.GetToken(false)
  contact, err := services.WeComApp.OAuth.Provider.WithApiAccessToken(cachedToken.AccessToken).ContactFromCode(para.Code)
  if err != nil {
    panic(err)
  }

  rs := &power.HashMap{
    "openID":   contact.GetOpenID(),
  }

  //// 正常返回json
  context.JSON(http.StatusOK, rs)

}
