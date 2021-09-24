package services

import (
  "github.com/ArtisanCloud/power-wechat/src/officialAccount"
  "log"
  "os"
)

var OfficialAccountApp *officialAccount.OfficialAccount


func NewOfficialAccountAppService() (*officialAccount.OfficialAccount, error) {
  log.Printf("officialAccount app_id: %s", os.Getenv("miniprogram_app_id"))
  app, err := officialAccount.NewOfficialAccount(&officialAccount.UserConfig{

    AppID:  os.Getenv("miniprogram_app_id"), // 小程序、公众号或者企业微信的appid
    Secret: os.Getenv("miniprogram_secret"), // 商户号 appID

    ResponseType: os.Getenv("response_type"),
    Log: officialAccount.Log{
      Level: "debug",
      File:  "./wechat.log",
    },

    HttpDebug: true,
    Debug:     false,
    //"sandbox": true,

  })

  return app, err
}
