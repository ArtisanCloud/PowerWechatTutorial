package official_account

import (
  "github.com/gin-gonic/gin"
  "net/http"
  "power-wechat-tutorial/services"
)

func UserFromCode(ctx *gin.Context) {
  code := ctx.Query("code")
  user, err := services.OfficialAccountApp.OAuth.UserFromCode(code)
  if err != nil {
    ctx.JSON(http.StatusBadRequest, err)
    return
  }
  ctx.JSON(http.StatusOK, user)
}

func UserFromToken(ctx *gin.Context) {
  token := ctx.Query("token")
  user, err := services.OfficialAccountApp.OAuth.UserFromToken(token)
  if err != nil {
    ctx.JSON(http.StatusBadRequest, err)
    return
  }
  ctx.JSON(http.StatusOK, user)
}
