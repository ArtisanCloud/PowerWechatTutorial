# PowerWeChat Tutorial 使用实例

我们单独写了一个项目[PowerWechatTutorial](https://github.com/ArtisanCloud/PowerWechatTutorial)，基本上覆盖了所有的API使用，希望能够帮助大家更快的上手Golang WeChat开发。

### 下载项目
```bash
git clone https://github.com/ArtisanCloud/PowerWechatTutorial.git
```

### Insomnia配置（可选）
可以下载insomnia的配置文件，便于在Tutorial里使用insomnia调试。
[insomnia.json](./insomnia.json)

### 项目配置
在项目根目录下，新建一个`config.yaml`, 把下面字段内容复制进去， 然后执行`go run main.go`。
如果程序正常启动，访问 [http://localhost:8888](http://localhost:8888) 会返回一个`Hello, PowerWechat`

```yaml
# 微信支付配置文档： https://powerwechat.artisan-cloud.com/zh/payment/index.html#userconfig%E5%8F%82%E6%95%B0%E8%AF%B4%E6%98%8E%EF%BC%9A
payment:
  appid: xxxxx # 公众号appid、小程序appid、企业微信corpid等
  mchid: 100000 # ArtisanCloud商户号
  certpath: certs/apiclient_cert.pem # 证书路径
  keypath: certs/apiclient_key.pem # 证书私钥路径
  serialno: 55D06F99FF64CF1759FDE5B77A0BEC8B67A78C2E
  key: xxxxx
  mchapiv3key: xxxxx
  notifyurl: https://www.artisancloud.cn
  redisaddr: localhost:6379

# 小程序配置文档： https://powerwechat.artisan-cloud.com/zh/mini-program/index.html
miniprogram:
  appid: xxxxx
  secret: xxxxx
  redisaddr: localhost:6379
  messagetoken: xxxxx
  messageaeskey: xxxxx

# 企业微信配置文档： https://powerwechat.artisan-cloud.com/zh/wecom/index.html
wecom:
  corpid: ww454dfb9d6f6d432a

  # ----- powerwechat-docs-demo start ------
  agent_id: 1000000
  secret: xxxxx
  messagetoken: xxxxx
  messageaeskey: xxxxx
  messagecallback: https://www.artisan-cloud.com/message/callback
  oauthcallback: https://www.artisan-cloud.com/oauth/callback
  # ----- powerwechat-docs-demo end ------

  # 联系人配置
  contactsecret: xxxxx
  contacttoken: xxxxx
  contactaeskey: xxxxx
  contactcallback: https://www.artisan-cloud.com/api/wx/customer

  redisaddr: localhost:6379

# 公众号配置文档： https://powerwechat.artisan-cloud.com/zh/official-account/index.html
offiaccount:
  appid: xxxxx
  appsecret: xxxxx
  redisaddr: localhost:6379
  messagetoken: xxxxx
  messageaeskey: xxxxx
```

在`main.go`里面，你可以选择性的注释掉不需要的模块，避免没有配置的时候报错影响运行。
