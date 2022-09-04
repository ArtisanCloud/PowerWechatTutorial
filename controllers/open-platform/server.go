package open_platform

import (
	"bytes"
	"encoding/xml"
	"github.com/ArtisanCloud/PowerLibs/v2/fmt"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel"
	openplatform "github.com/ArtisanCloud/PowerWeChat/v2/src/openPlatform/server/callbacks"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"power-wechat-tutorial/services"
)

func APIOpenPlatformCallbackVerifyTicket(context *gin.Context) {

	requestXML, _ := ioutil.ReadAll(context.Request.Body)
	context.Request.Body = ioutil.NopCloser(bytes.NewBuffer(requestXML))
	println(string(requestXML))

	var err error

	rs, err := services.OpenPlatformApp.Server.Notify(context.Request, func(event *openplatform.Callback, decrypted []byte, infoType string) (result interface{}) {

		result = kernel.SUCCESS_EMPTY_RESPONSE
		//fmt.Dump(event)
		msg := &openplatform.EventVerifyTicket{}
		err = xml.Unmarshal(decrypted, msg)
		if err != nil {
			return err
		}
		fmt.Dump(msg)

		return result
	})

	if err != nil {
		panic(err)
	}

	err = rs.Send(context.Writer)
	if err != nil {
		panic(err)
	}

}

func APIOpenPlatformCallback(context *gin.Context) {

	requestXML, _ := ioutil.ReadAll(context.Request.Body)
	context.Request.Body = ioutil.NopCloser(bytes.NewBuffer(requestXML))
	println(string(requestXML))

	var err error

	rs, err := services.OpenPlatformApp.Server.Notify(context.Request, func(event *openplatform.Callback, decrypted []byte, infoType string) (result interface{}) {

		result = kernel.SUCCESS_EMPTY_RESPONSE
		//fmt.Dump(event)
		msg := &openplatform.EventAuthorization{}
		err = xml.Unmarshal(decrypted, msg)
		if err != nil {
			return err
		}
		fmt.Dump(msg)

		return result
	})

	if err != nil {
		panic(err)
	}

	err = rs.Send(context.Writer)
	if err != nil {
		panic(err)
	}

}
