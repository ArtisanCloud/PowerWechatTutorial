package official_account

import (
  "github.com/gin-gonic/gin"
  "log"
  "net/http"
  "power-wechat-tutorial/services"
)

func UserFromCode(ctx *gin.Context) {
  code := ctx.Query("code")
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
  log.Println(err)
  if err != nil {
    ctx.String(http.StatusBadRequest, err.Error())
    return
  }
  ctx.JSON(http.StatusOK, user)
}
