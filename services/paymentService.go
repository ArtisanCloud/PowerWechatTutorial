package services

import (
  "github.com/ArtisanCloud/power-wechat/src/payment"
  "log"
  "net/http"
  "os"
)

const TRANSACTION_SUCCESS = "TRANSACTION.SUCCESS"
const TRANSACTION_FAILED = "TRANSACTION.FAILED"

var PaymentService *payment.Payment

func NewWXPaymentService(r *http.Request) (*payment.Payment, error) {
  log.Printf("app_id: %s", os.Getenv("app_id"))
  Payment, err := payment.NewPayment(&payment.UserConfig{
    //"corp_id":        os.Getenv("corp_id"),
    //"secret":         os.Getenv("secret"),
    AppID:       os.Getenv("app_id"),         // 小程序、公众号或者企业微信的appid
    MchID:       os.Getenv("mch_id"),         // 商户号 appID
    MchApiV3Key: os.Getenv("mch_api_v3_key"), //
    Key:         os.Getenv("key"),
    CertPath:    os.Getenv("wx_cert_path"),
    KeyPath:     os.Getenv("wx_key_path"),
    SerialNo:    os.Getenv("serial_no"),

    ResponseType: os.Getenv("array"),
    Log: payment.Log{
      Level: "debug",
      File:  "./wechat.log",
    },
    Http: payment.Http{
      Timeout: 30.0,
      BaseURI: "https://api.mch.weixin.qq.com",
    },

    NotifyURL: os.Getenv("notify_url"),
    HttpDebug: true,
    //Debug: true,
    //"sandbox": true,

    // server config
    //Token:            os.Getenv("token"),
    //AESKey:           os.Getenv("aes_key"),
  })

  return Payment, err
}
