package miniprogram

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/power"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"power-wechat-tutorial/services"
)

// 添加管理直播间小助手
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/livebroadcast/liveBroadcast.addAssistant.html
func APILiveAddAssistant(c *gin.Context) {
	var err error
	mediaPath := "./resource/cloud.jpg"
	value, err := ioutil.ReadFile(mediaPath)

	rs, err := services.MiniProgramApp.Broadcast.AddAssistant(&power.HashMap{
		"name":  "cloud.jpg", // 请确保文件名有准确的文件类型
		"value": value,
	})

	if err != nil {
		panic(err)
	}

	c.JSON(200, rs)
}


// 直播间导入商品
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/livebroadcast/liveBroadcast.addGoods.html
func APILiveAddGoods(c *gin.Context) {
	var err error
	mediaPath := "./resource/cloud.jpg"
	value, err := ioutil.ReadFile(mediaPath)

	rs, err := services.MiniProgramApp.Broadcast.AddGoods(&power.HashMap{
		"name":  "cloud.jpg", // 请确保文件名有准确的文件类型
		"value": value,
	})

	if err != nil {
		panic(err)
	}

	c.JSON(200, rs)
}

// 设置成员角色
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/livebroadcast/liveBroadcast.addRole.html
func APILiveAddRole(c *gin.Context) {
	var err error
	mediaPath := "./resource/cloud.jpg"
	value, err := ioutil.ReadFile(mediaPath)

	rs, err := services.MiniProgramApp.Broadcast.AddRole(&power.HashMap{
		"name":  "cloud.jpg", // 请确保文件名有准确的文件类型
		"value": value,
	})

	if err != nil {
		panic(err)
	}

	c.JSON(200, rs)
}

// 添加主播副号
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/livebroadcast/liveBroadcast.addSubAnchor.html
func APILiveAddSubAnchor(c *gin.Context) {
	var err error
	mediaPath := "./resource/cloud.jpg"
	value, err := ioutil.ReadFile(mediaPath)

	rs, err := services.MiniProgramApp.Broadcast.AddSubAnchor(&power.HashMap{
		"name":  "cloud.jpg", // 请确保文件名有准确的文件类型
		"value": value,
	})

	if err != nil {
		panic(err)
	}

	c.JSON(200, rs)
}

// 创建直播间
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/livebroadcast/liveBroadcast.createRoom.html
func APILiveCreateRoom(c *gin.Context) {
	var err error
	mediaPath := "./resource/cloud.jpg"
	value, err := ioutil.ReadFile(mediaPath)

	rs, err := services.MiniProgramApp.Broadcast.CreateRoom(&power.HashMap{
		"name":  "cloud.jpg", // 请确保文件名有准确的文件类型
		"value": value,
	})

	if err != nil {
		panic(err)
	}

	c.JSON(200, rs)
}

// 解除成员角色
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/livebroadcast/liveBroadcast.deleteRole.html
func APILiveDeleteRole(c *gin.Context) {
	var err error
	mediaPath := "./resource/cloud.jpg"
	value, err := ioutil.ReadFile(mediaPath)

	rs, err := services.MiniProgramApp.Broadcast.DeleteRole(&power.HashMap{
		"name":  "cloud.jpg", // 请确保文件名有准确的文件类型
		"value": value,
	})

	if err != nil {
		panic(err)
	}

	c.JSON(200, rs)
}

// 删除直播间
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/livebroadcast/liveBroadcast.deleteRoom.html
func APILiveDeleteRoom(c *gin.Context) {
	var err error
	mediaPath := "./resource/cloud.jpg"
	value, err := ioutil.ReadFile(mediaPath)

	rs, err := services.MiniProgramApp.Broadcast.DeleteRoom(&power.HashMap{
		"name":  "cloud.jpg", // 请确保文件名有准确的文件类型
		"value": value,
	})

	if err != nil {
		panic(err)
	}

	c.JSON(200, rs)
}

// 删除主播副号
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/livebroadcast/liveBroadcast.deleteSubAnchor.html
func APILiveDeleteSubAnchor(c *gin.Context) {
	var err error
	mediaPath := "./resource/cloud.jpg"
	value, err := ioutil.ReadFile(mediaPath)

	rs, err := services.MiniProgramApp.Broadcast.DeleteSubAnchor(&power.HashMap{
		"name":  "cloud.jpg", // 请确保文件名有准确的文件类型
		"value": value,
	})

	if err != nil {
		panic(err)
	}

	c.JSON(200, rs)
}

// 编辑直播间
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/livebroadcast/liveBroadcast.editRoom.html
func APILiveEditRoom(c *gin.Context) {
	var err error
	mediaPath := "./resource/cloud.jpg"
	value, err := ioutil.ReadFile(mediaPath)

	rs, err := services.MiniProgramApp.Broadcast.EditRoom(&power.HashMap{
		"name":  "cloud.jpg", // 请确保文件名有准确的文件类型
		"value": value,
	})

	if err != nil {
		panic(err)
	}

	c.JSON(200, rs)
}

// 查询管理直播间小助手
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/livebroadcast/liveBroadcast.getAssistantList.html
func APILiveGetAssistantList(c *gin.Context) {
	var err error
	mediaPath := "./resource/cloud.jpg"
	value, err := ioutil.ReadFile(mediaPath)

	rs, err := services.MiniProgramApp.Broadcast.GetAssistantList(&power.HashMap{
		"name":  "cloud.jpg", // 请确保文件名有准确的文件类型
		"value": value,
	})

	if err != nil {
		panic(err)
	}

	c.JSON(200, rs)
}

// 获取长期订阅用户
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/livebroadcast/liveBroadcast.getFollowers.html
func APILiveGetFollowers(c *gin.Context) {
	var err error
	mediaPath := "./resource/cloud.jpg"
	value, err := ioutil.ReadFile(mediaPath)

	rs, err := services.MiniProgramApp.Broadcast.GetFollowers(&power.HashMap{
		"name":  "cloud.jpg", // 请确保文件名有准确的文件类型
		"value": value,
	})

	if err != nil {
		panic(err)
	}

	c.JSON(200, rs)
}

// 获取直播间列表及直播间信息
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/livebroadcast/liveBroadcast.getLiveInfo.html
func APILiveGetLiveInfo(c *gin.Context) {
	var err error
	mediaPath := "./resource/cloud.jpg"
	value, err := ioutil.ReadFile(mediaPath)

	rs, err := services.MiniProgramApp.Broadcast.GetLiveInfo(&power.HashMap{
		"name":  "cloud.jpg", // 请确保文件名有准确的文件类型
		"value": value,
	})

	if err != nil {
		panic(err)
	}

	c.JSON(200, rs)
}

// 获取直播间推流地址
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/livebroadcast/liveBroadcast.getPushUrl.html
func APILiveGetPushUrl(c *gin.Context) {
	var err error
	mediaPath := "./resource/cloud.jpg"
	value, err := ioutil.ReadFile(mediaPath)

	rs, err := services.MiniProgramApp.Broadcast.GetPushUrl(&power.HashMap{
		"name":  "cloud.jpg", // 请确保文件名有准确的文件类型
		"value": value,
	})

	if err != nil {
		panic(err)
	}

	c.JSON(200, rs)
}

// 获取直播间推流地址
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/livebroadcast/liveBroadcast.getPushUrl.html
func APILiveGetRoleList(c *gin.Context) {
	var err error
	mediaPath := "./resource/cloud.jpg"
	value, err := ioutil.ReadFile(mediaPath)

	rs, err := services.MiniProgramApp.Broadcast.GetRoleList(&power.HashMap{
		"name":  "cloud.jpg", // 请确保文件名有准确的文件类型
		"value": value,
	})

	if err != nil {
		panic(err)
	}

	c.JSON(200, rs)
}

// 获取直播间分享二维码
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/livebroadcast/liveBroadcast.getSharedCode.html
func APILiveGetSharedCode(c *gin.Context) {
	var err error
	mediaPath := "./resource/cloud.jpg"
	value, err := ioutil.ReadFile(mediaPath)

	rs, err := services.MiniProgramApp.Broadcast.GetSharedCode(&power.HashMap{
		"name":  "cloud.jpg", // 请确保文件名有准确的文件类型
		"value": value,
	})

	if err != nil {
		panic(err)
	}

	c.JSON(200, rs)
}

// 获取主播副号
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/livebroadcast/liveBroadcast.getSubAnchor.html
func APILiveGetSubAnchor(c *gin.Context) {
	var err error
	mediaPath := "./resource/cloud.jpg"
	value, err := ioutil.ReadFile(mediaPath)

	rs, err := services.MiniProgramApp.Broadcast.GetSubAnchor(&power.HashMap{
		"name":  "cloud.jpg", // 请确保文件名有准确的文件类型
		"value": value,
	})

	if err != nil {
		panic(err)
	}

	c.JSON(200, rs)
}

// 商品添加并提审
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/livebroadcast/liveBroadcast.goodsAdd.html
func APILiveGoodsAdd(c *gin.Context) {
	var err error
	mediaPath := "./resource/cloud.jpg"
	value, err := ioutil.ReadFile(mediaPath)

	rs, err := services.MiniProgramApp.Broadcast.GoodsAdd(&power.HashMap{
		"name":  "cloud.jpg", // 请确保文件名有准确的文件类型
		"value": value,
	})

	if err != nil {
		panic(err)
	}

	c.JSON(200, rs)
}

// 重新提交审核
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/livebroadcast/liveBroadcast.goodsAudit.html
func APILiveGoodsAudit(c *gin.Context) {
	var err error
	mediaPath := "./resource/cloud.jpg"
	value, err := ioutil.ReadFile(mediaPath)

	rs, err := services.MiniProgramApp.Broadcast.GoodsAudit(&power.HashMap{
		"name":  "cloud.jpg", // 请确保文件名有准确的文件类型
		"value": value,
	})

	if err != nil {
		panic(err)
	}

	c.JSON(200, rs)
}

// 重新提交审核
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/livebroadcast/liveBroadcast.goodsAudit.html
func APILiveGoodsDelete(c *gin.Context) {
	var err error
	mediaPath := "./resource/cloud.jpg"
	value, err := ioutil.ReadFile(mediaPath)

	rs, err := services.MiniProgramApp.Broadcast.GoodsDelete(&power.HashMap{
		"name":  "cloud.jpg", // 请确保文件名有准确的文件类型
		"value": value,
	})

	if err != nil {
		panic(err)
	}

	c.JSON(200, rs)
}

// 获取商品状态
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/livebroadcast/liveBroadcast.goodsInfo.html
func APILiveGoodsInfo(c *gin.Context) {
	var err error
	mediaPath := "./resource/cloud.jpg"
	value, err := ioutil.ReadFile(mediaPath)

	rs, err := services.MiniProgramApp.Broadcast.GoodsInfo(&power.HashMap{
		"name":  "cloud.jpg", // 请确保文件名有准确的文件类型
		"value": value,
	})

	if err != nil {
		panic(err)
	}

	c.JSON(200, rs)
}

// 获取商品列表
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/livebroadcast/liveBroadcast.goodsList.html
func APILiveGoodsList(c *gin.Context) {
	var err error
	mediaPath := "./resource/cloud.jpg"
	value, err := ioutil.ReadFile(mediaPath)

	rs, err := services.MiniProgramApp.Broadcast.GoodsList(&power.HashMap{
		"name":  "cloud.jpg", // 请确保文件名有准确的文件类型
		"value": value,
	})

	if err != nil {
		panic(err)
	}

	c.JSON(200, rs)
}

// 推送商品
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/livebroadcast/liveBroadcast.goodsPush.html
func APILiveGoodsPush(c *gin.Context) {
	var err error
	mediaPath := "./resource/cloud.jpg"
	value, err := ioutil.ReadFile(mediaPath)

	rs, err := services.MiniProgramApp.Broadcast.GoodsPush(&power.HashMap{
		"name":  "cloud.jpg", // 请确保文件名有准确的文件类型
		"value": value,
	})

	if err != nil {
		panic(err)
	}

	c.JSON(200, rs)
}

// 撤回商品审核
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/livebroadcast/liveBroadcast.goodsResetAudit.html
func APILiveGoodsResetAudit(c *gin.Context) {
	var err error
	mediaPath := "./resource/cloud.jpg"
	value, err := ioutil.ReadFile(mediaPath)

	rs, err := services.MiniProgramApp.Broadcast.GoodsResetAudit(&power.HashMap{
		"name":  "cloud.jpg", // 请确保文件名有准确的文件类型
		"value": value,
	})

	if err != nil {
		panic(err)
	}

	c.JSON(200, rs)
}

// 上下架商品
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/livebroadcast/liveBroadcast.goodsSale.html
func APILiveGoodsSale(c *gin.Context) {
	var err error
	mediaPath := "./resource/cloud.jpg"
	value, err := ioutil.ReadFile(mediaPath)

	rs, err := services.MiniProgramApp.Broadcast.GoodsSale(&power.HashMap{
		"name":  "cloud.jpg", // 请确保文件名有准确的文件类型
		"value": value,
	})

	if err != nil {
		panic(err)
	}

	c.JSON(200, rs)
}

// 直播间商品排序
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/livebroadcast/liveBroadcast.goodsSort.html
func APILiveGoodsSort(c *gin.Context) {
	var err error
	mediaPath := "./resource/cloud.jpg"
	value, err := ioutil.ReadFile(mediaPath)

	rs, err := services.MiniProgramApp.Broadcast.GoodsSort(&power.HashMap{
		"name":  "cloud.jpg", // 请确保文件名有准确的文件类型
		"value": value,
	})

	if err != nil {
		panic(err)
	}

	c.JSON(200, rs)
}

// 更新商品
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/livebroadcast/liveBroadcast.goodsUpdate.html
func APILiveGoodsUpdate(c *gin.Context) {
	var err error
	mediaPath := "./resource/cloud.jpg"
	value, err := ioutil.ReadFile(mediaPath)

	rs, err := services.MiniProgramApp.Broadcast.GoodsUpdate(&power.HashMap{
		"name":  "cloud.jpg", // 请确保文件名有准确的文件类型
		"value": value,
	})

	if err != nil {
		panic(err)
	}

	c.JSON(200, rs)
}

// 下载商品讲解视频
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/livebroadcast/liveBroadcast.goodsVideo.html
func APILiveGoodsVideo(c *gin.Context) {
	var err error
	mediaPath := "./resource/cloud.jpg"
	value, err := ioutil.ReadFile(mediaPath)

	rs, err := services.MiniProgramApp.Broadcast.GoodsVideo(&power.HashMap{
		"name":  "cloud.jpg", // 请确保文件名有准确的文件类型
		"value": value,
	})

	if err != nil {
		panic(err)
	}

	c.JSON(200, rs)
}


// 修改管理直播间小助手
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/livebroadcast/liveBroadcast.modifyAssistant.html
func APILiveModifyAssistant(c *gin.Context) {
	var err error
	mediaPath := "./resource/cloud.jpg"
	value, err := ioutil.ReadFile(mediaPath)

	rs, err := services.MiniProgramApp.Broadcast.ModifyAssistant(&power.HashMap{
		"name":  "cloud.jpg", // 请确保文件名有准确的文件类型
		"value": value,
	})

	if err != nil {
		panic(err)
	}

	c.JSON(200, rs)
}

// 修改主播副号
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/livebroadcast/liveBroadcast.modifySubAnchor.html
func APILiveModifySubAnchor(c *gin.Context) {
	var err error
	mediaPath := "./resource/cloud.jpg"
	value, err := ioutil.ReadFile(mediaPath)

	rs, err := services.MiniProgramApp.Broadcast.ModifySubAnchor(&power.HashMap{
		"name":  "cloud.jpg", // 请确保文件名有准确的文件类型
		"value": value,
	})

	if err != nil {
		panic(err)
	}

	c.JSON(200, rs)
}

// 向长期订阅用户群发直播间开始事件
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/livebroadcast/liveBroadcast.pushMessage.html
func APILivePushMessage(c *gin.Context) {
	var err error
	mediaPath := "./resource/cloud.jpg"
	value, err := ioutil.ReadFile(mediaPath)

	rs, err := services.MiniProgramApp.Broadcast.PushMessage(&power.HashMap{
		"name":  "cloud.jpg", // 请确保文件名有准确的文件类型
		"value": value,
	})

	if err != nil {
		panic(err)
	}

	c.JSON(200, rs)
}

// 删除管理直播间小助手
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/livebroadcast/liveBroadcast.removeAssistant.html
func APILiveRemoveAssistant(c *gin.Context) {
	var err error
	mediaPath := "./resource/cloud.jpg"
	value, err := ioutil.ReadFile(mediaPath)

	rs, err := services.MiniProgramApp.Broadcast.RemoveAssistant(&power.HashMap{
		"name":  "cloud.jpg", // 请确保文件名有准确的文件类型
		"value": value,
	})

	if err != nil {
		panic(err)
	}

	c.JSON(200, rs)
}

// 开启/关闭直播间全局禁言
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/livebroadcast/liveBroadcast.updateComment.html
func APILiveUpdateComment(c *gin.Context) {
	var err error
	mediaPath := "./resource/cloud.jpg"
	value, err := ioutil.ReadFile(mediaPath)

	rs, err := services.MiniProgramApp.Broadcast.UpdateComment(&power.HashMap{
		"name":  "cloud.jpg", // 请确保文件名有准确的文件类型
		"value": value,
	})

	if err != nil {
		panic(err)
	}

	c.JSON(200, rs)
}

// 开启/关闭直播间官方收录
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/livebroadcast/liveBroadcast.updateFeedPublic.html
func APILiveUpdateFeedPublic(c *gin.Context) {
	var err error
	mediaPath := "./resource/cloud.jpg"
	value, err := ioutil.ReadFile(mediaPath)

	rs, err := services.MiniProgramApp.Broadcast.UpdateFeedPublic(&power.HashMap{
		"name":  "cloud.jpg", // 请确保文件名有准确的文件类型
		"value": value,
	})

	if err != nil {
		panic(err)
	}

	c.JSON(200, rs)
}

// 开启/关闭客服功能
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/livebroadcast/liveBroadcast.updateKF.html
func APILiveUpdateKF(c *gin.Context) {
	var err error
	mediaPath := "./resource/cloud.jpg"
	value, err := ioutil.ReadFile(mediaPath)

	rs, err := services.MiniProgramApp.Broadcast.UpdateKF(&power.HashMap{
		"name":  "cloud.jpg", // 请确保文件名有准确的文件类型
		"value": value,
	})

	if err != nil {
		panic(err)
	}

	c.JSON(200, rs)
}

// 开启/关闭回放功能
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/livebroadcast/liveBroadcast.updateReplay.html
func APILiveUpdateReplay(c *gin.Context) {
	var err error
	mediaPath := "./resource/cloud.jpg"
	value, err := ioutil.ReadFile(mediaPath)

	rs, err := services.MiniProgramApp.Broadcast.UpdateReplay(&power.HashMap{
		"name":  "cloud.jpg", // 请确保文件名有准确的文件类型
		"value": value,
	})

	if err != nil {
		panic(err)
	}

	c.JSON(200, rs)
}
