package services

import (
  "github.com/ArtisanCloud/PowerWeChat/src/kernel/response"
  "github.com/ArtisanCloud/PowerWeChat/src/miniProgram"
  "log"
  "os"
  "power-wechat-tutorial/config"
)

var MiniProgramApp *miniProgram.MiniProgram

const TIMEZONE = "asia/shanghai"
const DATETIME_FORMAT = "20060102"

func NewMiniMiniProgramService(conf *config.Configuration) (*miniProgram.MiniProgram, error) {
  log.Printf("miniprogram app_id: %s", os.Getenv("miniprogram_app_id"))
  app, err := miniProgram.NewMiniProgram(&miniProgram.UserConfig{

    AppID:  conf.MiniProgram.AppID,  // 小程序、公众号或者企业微信的appid
    Secret: conf.MiniProgram.Secret, // 商户号 appID

    ResponseType: response.TYPE_MAP,
    Log: miniProgram.Log{
      Level: "debug",
      File:  "./wechat.log",
    },

    HttpDebug: true,
    Debug:     false,
    //"sandbox": true,

  })

  return app, err
}
