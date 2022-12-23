package wecom

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/power"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"power-wechat-tutorial/services"
)

// 上传临时素材
// https://work.weixin.qq.com/api/doc/90000/90135/90253
func APIMediaUploadByURL(c *gin.Context) {
	mediaPath := "./resource/qrcode.png"
	rs, err := services.WeComApp.Media.UploadTempFile(mediaPath, nil)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, rs)
}

// 上传临时素材
// https://work.weixin.qq.com/api/doc/90000/90135/90253
func APIMediaUploadByData(c *gin.Context) {
	var err error
	mediaPath := "./resource/cloud.jpg"
	value, err := ioutil.ReadFile(mediaPath)

	rs, err := services.WeComApp.Media.UploadTempFile("", &power.HashMap{
		"name":  "temp.jpg", // 请确保文件名有准确的文件类型
		"value": value,
	})

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, rs)
}

// 上传图片
// https://work.weixin.qq.com/api/doc/90000/90135/90256
func APIMediaUploadImgByPath(c *gin.Context) {
	var err error
	mediaPath := "./resource/cloud.jpg"

	rs, err := services.WeComApp.Media.UploadImage(mediaPath, nil)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, rs)
}

// 上传图片
// https://work.weixin.qq.com/api/doc/90000/90135/90256
func APIMediaUploadImgByData(c *gin.Context) {
	var err error
	mediaPath := "./resource/cloud.jpg"
	value, err := ioutil.ReadFile(mediaPath)

	rs, err := services.WeComApp.Media.UploadImage("", &power.HashMap{
		"name":  "image.jpg", // 请确保文件名有准确的文件类型
		"value": value,
	})

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, rs)
}

// 获取临时素材
// https://work.weixin.qq.com/api/doc/90000/90135/90254
func APIMediaGet(c *gin.Context) {
	var err error

	mediaID := c.DefaultQuery("mediaID", "MEDIAID")

	rs, err := services.WeComApp.Media.Get(mediaID)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, rs)
}

// 获取高清语音素材
// https://work.weixin.qq.com/api/doc/90000/90135/90255
func APIMediaGetJSSDK(c *gin.Context) {
	var err error

	mediaID := c.DefaultQuery("mediaID", "MEDIAID")

	rs, err := services.WeComApp.Media.GetJSSDK(mediaID)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, rs)
}
