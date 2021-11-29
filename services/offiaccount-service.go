package services

import (
  "github.com/ArtisanCloud/PowerWeChat/src/officialAccount"
  "log"
  "os"
  "power-wechat-tutorial/config"
)

var OfficialAccountApp *officialAccount.OfficialAccount


func NewOfficialAccountAppService(conf *config.Configuration) (*officialAccount.OfficialAccount, error) {
  log.Printf("officialAccount app_id: %s", os.Getenv("miniprogram_app_id"))
  app, err := officialAccount.NewOfficialAccount(&officialAccount.UserConfig{

    AppID:  conf.MiniProgram.AppID, // 小程序、公众号或者企业微信的appid
    Secret: conf.MiniProgram.Secret, // 商户号 appID

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
