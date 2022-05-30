package official_account

import (
  "github.com/gin-gonic/gin"
  "net/http"
  "power-wechat-tutorial/services"
)

func JSSDKBuildConfig(ctx *gin.Context) {
  url := "https://www.artisan-cloud.com/"
  jsapiList := []string{"chooseImage"}
  debug := false
  beta := false
  openTagList := []string{"wx-open-launch-app", "wx-open-launch-weapp"}
  data, err := services.OfficialAccountApp.JSSDK.BuildConfig(jsapiList, debug, beta, openTagList, url)
  if err != nil {
    ctx.String(http.StatusBadRequest, err.Error())
    return
  }
  ctx.JSON(http.StatusOK, data)
}
