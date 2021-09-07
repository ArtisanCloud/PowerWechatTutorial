package miniprogram

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/power"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"power-wechat-tutorial/services"
)

func APICustomerServiceMessageSend(c *gin.Context) {

	openID, exist := c.GetQuery("openID")
	if !exist {
		panic("parameter open id expected")
	}

	rs, err := services.AppMiniProgram.CustomerServiceMessage.Send(openID, "text", &power.HashMap{
		"content": "Hello World",
	})

	if err != nil {
		panic(err)
	}

	c.JSON(200, rs)

}

func APICustomerServiceMessageSetTyping(c *gin.Context) {

	openID, exist := c.GetQuery("openID")
	if !exist {
		panic("parameter open id expected")
	}
	command, exist := c.GetQuery("command")
	if !exist {
		panic("parameter command expected")
	}

	rs, err := services.AppMiniProgram.CustomerServiceMessage.SetTyping(openID, command)

	if err != nil {
		panic(err)
	}

	c.JSON(200, rs)

}

func APICustomerServiceMessageUploadTempMediaByFile(c *gin.Context) {

	mediaPath := "./resource/cloud.jpg"
	rs, err := services.AppMiniProgram.CustomerServiceMessage.UploadTempMedia("image", mediaPath, nil)

	if err != nil {
		panic(err)
	}

	c.JSON(200, rs)

}

func APICustomerServiceMessageUploadTempMediaByData(c *gin.Context) {

	var err error
	mediaPath := "./resource/cloud.jpg"
	value, err := ioutil.ReadFile(mediaPath)

	rs, err := services.AppMiniProgram.CustomerServiceMessage.UploadTempMedia("image", "", &power.HashMap{
		"name":     "cloud.jpg", // 请确保文件名有准确的文件类型
		"value":        value,
	})

	if err != nil {
		panic(err)
	}

	c.JSON(200, rs)

}

func APICustomerServiceMessageGetTempMedia(c *gin.Context) {

	mediaID, exist := c.GetQuery("mediaID")
	if !exist {
		panic("parameter media id expected")
	}

	rs, err := services.AppMiniProgram.CustomerServiceMessage.GetTempMedia(mediaID)

	if err != nil {
		panic(err)
	}

	content, _ := ioutil.ReadAll(rs.Body)
	//fmt.Dump("content-type:",rs.Header.Get("Content-Type"))
	c.Header("Content-Type", rs.Header.Get("Content-Type"))
	c.Header("Content-Disposition", rs.Header.Get("attachment;filename=\""+rs.Header.Get("filename")+"\""))
	c.Data(http.StatusOK, rs.Header.Get("Content-Type"), content)

}
