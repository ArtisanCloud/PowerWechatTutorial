package official_account

import (
  "github.com/gin-gonic/gin"
  "net/http"
  "power-wechat-tutorial/services"
)

func ClearQuota(ctx *gin.Context) {
  data, err := services.OfficialAccountApp.Base.ClearQuota()
  if err != nil {
    ctx.JSON(http.StatusBadRequest, err)
    return
  }
  ctx.JSON(http.StatusOK, data)
}

func GetCallbackIP(ctx *gin.Context) {
  data, err := services.OfficialAccountApp.Base.GetCallbackIP()
  if err != nil {
    ctx.JSON(http.StatusBadRequest, err)
    return
  }
  ctx.JSON(http.StatusOK, data)
}
