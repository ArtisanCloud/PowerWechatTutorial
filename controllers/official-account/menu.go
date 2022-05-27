package official_account

import (
  "github.com/ArtisanCloud/PowerWeChat/src/officialAccount/menu/request"
  "github.com/gin-gonic/gin"
  "net/http"
  "power-wechat-tutorial/services"
)

func MenuList(ctx *gin.Context) {
  data, err := services.OfficialAccountApp.Menu.Get()
  if err != nil {
    ctx.JSON(http.StatusBadRequest, err)
    return
  }
  ctx.JSON(http.StatusOK, data)
}

func MenuCurrent(ctx *gin.Context) {
  data, err := services.OfficialAccountApp.Menu.CurrentSelfMenu()
  if err != nil {
    ctx.JSON(http.StatusBadRequest, err)
    return
  }
  ctx.JSON(http.StatusOK, data)
}

func MenuCreate(ctx *gin.Context) {
  data, err := services.OfficialAccountApp.Menu.Create(&request.RequestMenuCreate{
    Buttons: []*request.Button{
      {
        Type: "click",
        Name: "今日歌曲",
        Key:  "V1001_TODAY_MUSIC",
      },
      {
        Name: "Menu1",
        SubButtons: []request.SubButton{
          {
            Type: "view",
            Name: "搜索",
            URL:  "http://www.soso.com/",
          },
          {
            Type:     "miniprogram",
            Name:     "wxa",
            URL:      "http://mp.weixin.qq.com",
            AppID:    "wx286b93c14bbf93aa",
            PagePath: "pages/lunar/index",
          },
          {
            Type: "click",
            Name: "赞一下我们",
            Key:  "V1001_GOOD",
          },
        },
      },
    },
  })
  if err != nil {
    ctx.JSON(http.StatusBadRequest, err)
    return
  }
  ctx.JSON(http.StatusOK, data)
}

func MenuCreateConditional(ctx *gin.Context) {
  data, err := services.OfficialAccountApp.Menu.CreateConditional(&request.RequestMenuCreate{
    Buttons: []*request.Button{
      {
        Type: "click",
        Name: "今日歌曲",
        Key:  "V1001_TODAY_MUSIC",
      },
    },
  }, &request.RequestMatchRule{
    Sex:                "1",
    Country:            "中国",
    Province:           "广东",
    City:               "广州",
    ClientPlatformType: "2",
  })
  if err != nil {
    ctx.JSON(http.StatusBadRequest, err)
    return
  }
  ctx.JSON(http.StatusOK, data)
}

func MenuDelete(ctx *gin.Context) {
  data, err := services.OfficialAccountApp.Menu.Delete()
  if err != nil {
    ctx.JSON(http.StatusBadRequest, err)
    return
  }
  ctx.JSON(http.StatusOK, data)
}

func MenuDeleteConditional(ctx *gin.Context) {
  data, err := services.OfficialAccountApp.Menu.DeleteConditional(1)
  if err != nil {
    ctx.JSON(http.StatusBadRequest, err)
    return
  }
  ctx.JSON(http.StatusOK, data)
}

func MenuMatch(ctx *gin.Context) {
  userID := ctx.Query("userID")
  data, err := services.OfficialAccountApp.Menu.TryMatch(userID)
  if err != nil {
    ctx.JSON(http.StatusBadRequest, err)
    return
  }
  ctx.JSON(http.StatusOK, data)
}
