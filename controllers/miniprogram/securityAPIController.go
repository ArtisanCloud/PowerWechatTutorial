package miniprogram

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/power"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"power-wechat-tutorial/services"
)

// 向插件开发者发起使用插件的申请
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/sec-check/security.imgSecCheck.html
func APISecurityImgSecCheckByPath(c *gin.Context) {

	mediaPath := "./resource/cloud.jpg"
	rs, err := services.AppMiniProgram.Security.ImgSecCheck(mediaPath, nil)

	if err != nil {
		panic(err)
	}

	c.JSON(200, rs)

}

// 向插件开发者发起使用插件的申请
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/sec-check/security.imgSecCheck.html
func APISecurityImgSecCheckByData(c *gin.Context) {

	var err error
	mediaPath := "./resource/cloud.jpg"
	value, err := ioutil.ReadFile(mediaPath)

	rs, err := services.AppMiniProgram.Security.ImgSecCheck("", &power.HashMap{
		"name":  "cloud.jpg", // 请确保文件名有准确的文件类型
		"value": value,
	})

	if err != nil {
		panic(err)
	}

	c.JSON(200, rs)

}

// 异步校验图片/音频是否含有违法违规内容
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/sec-check/security.mediaCheckAsync.html
func APISecurityMediaCheckAsync(c *gin.Context) {

	openID, exist := c.GetQuery("openID")
	if !exist {
		panic("parameter open id expected")
	}

	rs, err := services.AppMiniProgram.Security.MediaCheckAsync(
		"https://developers.weixin.qq.com/miniprogram/assets/images/head_global_z_@all.png",
		2,
		1,
		openID,
		1,
	)

	if err != nil {
		panic(err)
	}

	c.JSON(200, rs)

}

// 检查一段文本是否含有违法违规内容
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/sec-check/security.msgSecCheck.html
func APISecurityMsgSecCheck(c *gin.Context) {

	openID, exist := c.GetQuery("openID")
	if !exist {
		panic("parameter open id expected")
	}

	rs, err := services.AppMiniProgram.Security.MsgSecCheck(
		openID,
		1,
		2,
		"hello world",
		"test name",
		"test title",
		"test sign",
	)

	if err != nil {
		panic(err)
	}

	c.JSON(200, rs)

}
