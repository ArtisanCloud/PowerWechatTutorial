package miniprogram

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/power"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"power-wechat-tutorial/services"
)

// 同步校验图片违法违规内容
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/sec-check/security.imgSecCheck.html
func APISecurityImgSecCheckByPath(c *gin.Context) {

	mediaPath := "./resource/cloud.jpg"
	rs, err := services.MiniprogramApp.Security.ImgSecCheck(mediaPath, nil)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, rs)

}

// 同步校验图片违法违规内容
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/sec-check/security.imgSecCheck.html
func APISecurityImgSecCheckByData(c *gin.Context) {

	var err error
	mediaPath := "./resource/cloud.jpg"
	value, err := ioutil.ReadFile(mediaPath)

	rs, err := services.MiniprogramApp.Security.ImgSecCheck("", &power.HashMap{
		"name":  "cloud.jpg", // 请确保文件名有准确的文件类型
		"value": value,
	})

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, rs)

}

// 异步校验图片/音频是否含有违法违规内容
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/sec-check/security.mediaCheckAsync.html
func APISecurityMediaCheckAsync(c *gin.Context) {
	mediaURL := c.DefaultQuery("mediaURL", "https://www.w3school.com.cn/i/horse.mp3")
	openID, exist := c.GetQuery("openid")

	if !exist {
		panic("parameter open id expected")
	}

	rs, err := services.MiniprogramApp.Security.MediaCheckAsync(
		mediaURL,
		1,
		2,
		openID,
		1,
	)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, rs)

}



// 检查一段文本是否含有违法违规内容
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/sec-check/security.msgSecCheck.html
func APISecurityMsgSecCheck(c *gin.Context) {
	content := c.DefaultQuery("content", "Hello, World")
	openID, exist := c.GetQuery("openid")
	if !exist {
		panic("parameter open id expected")
	}

	rs, err := services.MiniprogramApp.Security.MsgSecCheck(
		openID,
		1,
		2,
		content,
		"test name",
		"test title",
		"test sign",
	)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, rs)

}
