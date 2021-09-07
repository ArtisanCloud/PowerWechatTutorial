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

	rs, err := services.AppMiniProgram.CustomerServiceMessage.UploadTempMedia("image", "", &power.HashMap{
		"Content-Type": "image/jpeg",
		"filename":     "test-cloud",
		"value":        "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",
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
