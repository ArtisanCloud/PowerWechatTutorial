package miniprogram

import (
  "github.com/ArtisanCloud/PowerWeChat/v2/src/miniProgram/updatableMessage/request"
  "github.com/gin-gonic/gin"
  "net/http"
  "power-wechat-tutorial/services"
)

// 创建被分享动态消息或私密消息的 activity_id
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/updatable-message/updatableMessage.createActivityId.html
func APIUpdatableMessageCreateActivityID(c *gin.Context) {

  openID, exist := c.GetQuery("openID")
  if !exist {
    panic("parameter open id expected")
  }

  rs, err := services.MiniProgramApp.UpdatableMessage.CreateActivityID("", openID)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, rs)
}

// 修改被分享的动态消息。
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/updatable-message/updatableMessage.setUpdatableMsg.html
func APIUpdatableMessageUpdatableMessage(c *gin.Context) {

  activityID, exist := c.GetQuery("activityID")
  if !exist {
    panic("parameter open id expected")
  }

  rs, err := services.MiniProgramApp.UpdatableMessage.SetUpdatableMsg(&request.RequestSetUpdatableMsg{
    ActivityID:  activityID,
    TargetState: 0,
    TemplateInfo: &request.TemplateInfo{
      ParameterList: []*request.ParameterListItem{
        {
          Name:  "member_count",
          Value: "2",
        },
        {
          Name:  "room_limit",
          Value: "5",
        },
      },
    },
  })

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, rs)
}
