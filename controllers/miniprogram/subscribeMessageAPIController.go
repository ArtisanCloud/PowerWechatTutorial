package miniprogram

import (
  "github.com/ArtisanCloud/PowerWeChat/src/miniProgram/subscribeMessage/request"
  "github.com/gin-gonic/gin"
  "net/http"
  "power-wechat-tutorial/services"
)

// 组合模板并添加至帐号下的个人模板库
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/subscribe-message/subscribeMessage.addTemplate.html
func APISubscribeMessageAddTemplate(c *gin.Context) {
  tID, exist := c.GetQuery("tid")
  if !exist {
    panic("parameter tid expected")
  }

  rs, err := services.MiniProgramApp.SubscribeMessage.AddTemplate(tID, []int{1, 2}, "测试数据")

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, rs)
}

// 删除帐号下的个人模板
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/subscribe-message/subscribeMessage.deleteTemplate.html
func APISubscribeMessageDeleteTemplate(c *gin.Context) {
  priTmplID, exist := c.GetQuery("priTmplID")
  if !exist {
    panic("parameter tid expected")
  }

  rs, err := services.MiniProgramApp.SubscribeMessage.DeleteTemplate(priTmplID)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, rs)
}

// 获取小程序账号的类目
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/subscribe-message/subscribeMessage.getCategory.html
func APISubscribeMessageGetCategory(c *gin.Context) {
  rs, err := services.MiniProgramApp.SubscribeMessage.GetCategory()

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, rs)
}

// 获取模板标题下的关键词列表
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/subscribe-message/subscribeMessage.getPubTemplateKeyWordsById.html
func APISubscribeMessageGetPubTemplateKeyWordsByID(c *gin.Context) {
  tID, exist := c.GetQuery("tid")
  if !exist {
    panic("parameter tid expected")
  }

  rs, err := services.MiniProgramApp.SubscribeMessage.GetPubTemplateKeyWordsByID(tID)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, rs)
}

// 获取帐号所属类目下的公共模板标题
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/subscribe-message/subscribeMessage.getPubTemplateTitleList.html
func APISubscribeMessageGetPubTemplateTitleList(c *gin.Context) {
  rs, err := services.MiniProgramApp.SubscribeMessage.GetPubTemplateTitleList("676", 0, 10)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, rs)
}

// 获取当前帐号下的个人模板列表
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/subscribe-message/subscribeMessage.getTemplateList.html
func APISubscribeMessageGetTemplateList(c *gin.Context) {
  rs, err := services.MiniProgramApp.SubscribeMessage.GetTemplateList()

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, rs)
}

// 发送订阅消息
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/subscribe-message/subscribeMessage.send.html
func APISubscribeMessageSend(c *gin.Context) {

  toUser, exist := c.GetQuery("openid")
  if !exist {
    panic("parameter open id expected")
  }

  templateID := c.DefaultQuery("templateID", "Y1471771tIQyEogSHjqCgD1P7iy52N_JYH-q0Sw7EvQ")

  page := "page/index/index"
  miniprogramState := "developer"
  lang := "zh_CN"

  data := &power.HashMap{
    "phrase1": power.StringMap{
      "value": "已预约",
    },
    "thing2": power.StringMap{
      "value": "Cycle",
    },
    "time3": power.StringMap{
      "value": "15:30",
    },
    "thing4": power.StringMap{
      "value": "兴业太古汇",
    },
    "thing5": power.StringMap{
      "value": "开课3小时前",
    },
  }

  rs, err := services.MiniProgramApp.SubscribeMessage.Send(&request.RequestSubscribeMessageSend{
    ToUser:           toUser,
    TemplateID:       templateID,
    Page:             page,
    MiniProgramState: miniprogramState,
    Lang:             lang,
    Data:             data,
  })

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, rs)
}
