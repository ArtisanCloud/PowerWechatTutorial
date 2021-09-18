package miniprogram

import (
  "github.com/ArtisanCloud/power-wechat/src/kernel/power"
  "github.com/ArtisanCloud/power-wechat/src/miniProgram/liveBroadcast/request"
  "github.com/gin-gonic/gin"
  "net/http"
  "power-wechat-tutorial/services"
)

// APILiveAddAssistant 添加管理直播间小助手
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/livebroadcast/liveBroadcast.addAssistant.html
func APILiveAddAssistant(c *gin.Context) {
  rs, err := services.MiniProgramApp.Broadcast.AddAssistant(&request.RequestBroadcastAddAssistant{
    RoomID: 0,
    Users:  nil,
  })

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, rs)
}

// APILiveAddGoods 直播间导入商品
// https://developers.weixin.qq.com/miniprogram/dev/platform-capabilities/industry/liveplayer/studio-api.html#4
func APILiveAddGoods(c *gin.Context) {

  rs, err := services.MiniProgramApp.Broadcast.AddGoods(&request.RequestBroadcastAddGoods{
    IDs:    "",
    RoomID: 0,
  })

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, rs)
}

// APILiveAddRole 设置成员角色
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/livebroadcast/liveBroadcast.addRole.html
func APILiveAddRole(c *gin.Context) {
  rs, err := services.MiniProgramApp.Broadcast.AddRole("robot1", 1)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, rs)
}

// 添加主播副号
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/livebroadcast/liveBroadcast.addSubAnchor.html
func APILiveAddSubAnchor(c *gin.Context) {
  
  rs, err := services.MiniProgramApp.Broadcast.AddSubAnchor(3, "WalleAI")

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, rs)
}

// 创建直播间
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/livebroadcast/liveBroadcast.createRoom.html
func APILiveCreateRoom(c *gin.Context) {
  rs, err := services.MiniProgramApp.Broadcast.CreateRoom(&request.RequestBroadcastCreateRoom{
    Name:          "直播测试3",
    CoverImg:      "xisnqd5vEly5mjX9r3VjW1XGinnZjyEHcv0czEKYw8SBff7lixnbCIrE0QkgStj4",
    StartTime:     1631923200,
    EndTime:       1631928200,
    AnchorName:    "Ange",
    AnchorWechat:  "WalleAI",
    ShareImg:      "xisnqd5vEly5mjX9r3VjW1XGinnZjyEHcv0czEKYw8SBff7lixnbCIrE0QkgStj4",
    CloseLike:     0,
    CloseGoods:    0,
    CloseComment:  0,
    IsFeedsPublic: 0,
    CloseReplay:   0,
    CloseShare:    0,
    FeedsImg:      "xisnqd5vEly5mjX9r3VjW1XGinnZjyEHcv0czEKYw8SBff7lixnbCIrE0QkgStj4",
  })

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, rs)
}

// 解除成员角色
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/livebroadcast/liveBroadcast.deleteRole.html
func APILiveDeleteRole(c *gin.Context) {
  rs, err := services.MiniProgramApp.Broadcast.DeleteRole("walle", 1)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, rs)
}

// 删除直播间
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/livebroadcast/liveBroadcast.deleteRoom.html
func APILiveDeleteRoom(c *gin.Context) {
  rs, err := services.MiniProgramApp.Broadcast.DeleteRoom(1)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, rs)
}

// 删除主播副号
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/livebroadcast/liveBroadcast.deleteSubAnchor.html
func APILiveDeleteSubAnchor(c *gin.Context) {
  rs, err := services.MiniProgramApp.Broadcast.DeleteSubAnchor(1)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, rs)
}

// 编辑直播间
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/livebroadcast/liveBroadcast.editRoom.html
func APILiveEditRoom(c *gin.Context) {
  rs, err := services.MiniProgramApp.Broadcast.EditRoom(&request.RequestBroadcastEditRoom{
    Id:            2,
    Name:          "直播测试",
    CoverImg:      "xisnqd5vEly5mjX9r3VjW1XGinnZjyEHcv0czEKYw8SBff7lixnbCIrE0QkgStj4",
    StartTime:     1631923200,
    EndTime:       1631928200,
    AnchorName:    "Ange",
    AnchorWechat:  "WalleAI",
    ShareImg:      "xisnqd5vEly5mjX9r3VjW1XGinnZjyEHcv0czEKYw8SBff7lixnbCIrE0QkgStj4",
    CloseLike:     0,
    CloseGoods:    0,
    CloseComment:  0,
    IsFeedsPublic: 0,
    CloseReplay:   0,
    CloseShare:    0,
    CloseKF:       0,
    FeedsImg:      "xisnqd5vEly5mjX9r3VjW1XGinnZjyEHcv0czEKYw8SBff7lixnbCIrE0QkgStj4",
  })

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, rs)
}

// 查询管理直播间小助手
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/livebroadcast/liveBroadcast.getAssistantList.html
func APILiveGetAssistantList(c *gin.Context) {
  rs, err := services.MiniProgramApp.Broadcast.GetAssistantList(2)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, rs)
}

// 获取长期订阅用户
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/livebroadcast/liveBroadcast.getFollowers.html
func APILiveGetFollowers(c *gin.Context) {
  rs, err := services.MiniProgramApp.Broadcast.GetFollowers(1, 10)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, rs)
}

// 获取直播间列表及直播间信息
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/livebroadcast/liveBroadcast.getLiveInfo.html
func APILiveGetLiveInfo(c *gin.Context) {
  rs, err := services.MiniProgramApp.Broadcast.GetLiveInfo(&request.RequestBroadcastGetLiveInfo{
    Start: 0,
    Limit: 0,
  })

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, rs)
}

// 获取直播间推流地址
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/livebroadcast/liveBroadcast.getPushUrl.html
func APILiveGetPushUrl(c *gin.Context) {
  rs, err := services.MiniProgramApp.Broadcast.GetPushUrl(3)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, rs)
}

// 获取直播间推流地址
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/livebroadcast/liveBroadcast.getPushUrl.html
func APILiveGetRoleList(c *gin.Context) {
  rs, err := services.MiniProgramApp.Broadcast.GetRoleList(&request.RequestBroadcastGetRoleList{
    Role:    0,
    Offset:  0,
    Limit:   0,
    Keyword: "",
  })

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, rs)
}

// 获取直播间分享二维码
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/livebroadcast/liveBroadcast.getSharedCode.html
func APILiveGetSharedCode(c *gin.Context) {
  rs, err := services.MiniProgramApp.Broadcast.GetSharedCode(2, "a=1&b=2")

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, rs)
}

// 获取主播副号
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/livebroadcast/liveBroadcast.getSubAnchor.html
func APILiveGetSubAnchor(c *gin.Context) {
  rs, err := services.MiniProgramApp.Broadcast.GetSubAnchor(1)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, rs)
}

// 商品添加并提审
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/livebroadcast/liveBroadcast.goodsAdd.html
func APILiveGoodsAdd(c *gin.Context) {
  rs, err := services.MiniProgramApp.Broadcast.GoodsAdd(&request.RequestBroadcastGoodsAdd{
    GoodsInfo: &power.HashMap{
      "coverImgUrl": "Hbgk4dmhoGouU2gCTgFJBV1GrFjgrHGM_3o5JODtJspgotaXrG7YIBUUfFaJ-bdT",
      "name":"TIT茶杯",
      "priceType": 1,
      "price": 99.5,
      // "price2": 150.5, priceType为2或3时必填
      "url":"pages/index/index",
    },
  })

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, rs)
}

// APILiveGoodsAudit 重新提交审核
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/livebroadcast/liveBroadcast.goodsAudit.html
func APILiveGoodsAudit(c *gin.Context) {
  rs, err := services.MiniProgramApp.Broadcast.GoodsAudit(4)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, rs)
}

// 重新提交审核
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/livebroadcast/liveBroadcast.goodsAudit.html
func APILiveGoodsDelete(c *gin.Context) {
  rs, err := services.MiniProgramApp.Broadcast.GoodsDelete(4)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, rs)
}

// 获取商品状态
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/livebroadcast/liveBroadcast.goodsInfo.html
func APILiveGoodsInfo(c *gin.Context) {
  rs, err := services.MiniProgramApp.Broadcast.GoodsInfo([]int{1})

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, rs)
}

// 获取商品列表
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/livebroadcast/liveBroadcast.goodsList.html
func APILiveGoodsList(c *gin.Context) {
  offset := c.DefaultQuery("offset", "0")
  count := c.DefaultQuery("count", "0")
  status := c.DefaultQuery("status", "2")
  rs, err := services.MiniProgramApp.Broadcast.GoodsList(offset, count, status)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, rs)
}

// 推送商品
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/livebroadcast/liveBroadcast.goodsPush.html
func APILiveGoodsPush(c *gin.Context) {
  rs, err := services.MiniProgramApp.Broadcast.GoodsPush(1, 1)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, rs)
}

// 撤回商品审核
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/livebroadcast/liveBroadcast.goodsResetAudit.html
func APILiveGoodsResetAudit(c *gin.Context) {
  rs, err := services.MiniProgramApp.Broadcast.GoodsResetAudit(&request.RequestBroadcastGoodsResetAudit{
    AuditID: 0,
    GoodsID: 0,
  })

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, rs)
}

// 上下架商品
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/livebroadcast/liveBroadcast.goodsSale.html
func APILiveGoodsSale(c *gin.Context) {
  rs, err := services.MiniProgramApp.Broadcast.GoodsSale(1, 1,1)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, rs)
}

// 直播间商品排序
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/livebroadcast/liveBroadcast.goodsSort.html
func APILiveGoodsSort(c *gin.Context) {
  rs, err := services.MiniProgramApp.Broadcast.GoodsSort(1, nil)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, rs)
}

// 更新商品
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/livebroadcast/liveBroadcast.goodsUpdate.html
func APILiveGoodsUpdate(c *gin.Context) {
  rs, err := services.MiniProgramApp.Broadcast.GoodsUpdate(&request.RequestBroadcastGoodsUpdate{
    GoodsInfo: nil,
  })

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, rs)
}

// 下载商品讲解视频
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/livebroadcast/liveBroadcast.goodsVideo.html
func APILiveGoodsVideo(c *gin.Context) {
  rs, err := services.MiniProgramApp.Broadcast.GoodsVideo(1, 1)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, rs)
}

// 修改管理直播间小助手
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/livebroadcast/liveBroadcast.modifyAssistant.html
func APILiveModifyAssistant(c *gin.Context) {

  rs, err := services.MiniProgramApp.Broadcast.ModifyAssistant(&request.RequestBroadcastModifyAssistant{
    RoomID:   0,
    UserName: "",
    NickName: "",
  })

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, rs)
}

// 修改主播副号
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/livebroadcast/liveBroadcast.modifySubAnchor.html
func APILiveModifySubAnchor(c *gin.Context) {
  rs, err := services.MiniProgramApp.Broadcast.ModifySubAnchor(1, "WalleAI")

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, rs)
}

// 向长期订阅用户群发直播间开始事件
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/livebroadcast/liveBroadcast.pushMessage.html
func APILivePushMessage(c *gin.Context) {
  rs, err := services.MiniProgramApp.Broadcast.PushMessage(2, []string{"oC-vT5KfvgxATPoVl88oeTE-hnfE"})

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, rs)
}

// 删除管理直播间小助手
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/livebroadcast/liveBroadcast.removeAssistant.html
func APILiveRemoveAssistant(c *gin.Context) {
  rs, err := services.MiniProgramApp.Broadcast.RemoveAssistant(&request.RequestBroadcastRemoveAssistant{
    RoomID:   0,
    UserName: "",
  })

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, rs)
}

// 开启/关闭直播间全局禁言
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/livebroadcast/liveBroadcast.updateComment.html
func APILiveUpdateComment(c *gin.Context) {
  rs, err := services.MiniProgramApp.Broadcast.UpdateComment(2, 1)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, rs)
}

// 开启/关闭直播间官方收录
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/livebroadcast/liveBroadcast.updateFeedPublic.html
func APILiveUpdateFeedPublic(c *gin.Context) {
  rs, err := services.MiniProgramApp.Broadcast.UpdateFeedPublic(2, 0)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, rs)
}

// 开启/关闭客服功能
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/livebroadcast/liveBroadcast.updateKF.html
func APILiveUpdateKF(c *gin.Context) {
  rs, err := services.MiniProgramApp.Broadcast.UpdateKF(2,1)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, rs)
}

// 开启/关闭回放功能
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/livebroadcast/liveBroadcast.updateReplay.html
func APILiveUpdateReplay(c *gin.Context) {
  rs, err := services.MiniProgramApp.Broadcast.UpdateReplay(1, 1)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, rs)
}
