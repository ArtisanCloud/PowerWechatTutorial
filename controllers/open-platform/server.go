package open_platform

import (
	"bytes"
	"encoding/xml"
	"github.com/ArtisanCloud/PowerLibs/v3/fmt"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	openplatform "github.com/ArtisanCloud/PowerWeChat/v3/src/openPlatform/server/callbacks"
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
	"power-wechat-tutorial/services"
)

func APIOpenPlatformCallback(context *gin.Context) {

	requestXML, _ := io.ReadAll(context.Request.Body)
	context.Request.Body = io.NopCloser(bytes.NewBuffer(requestXML))
	println(string(requestXML))

	var err error

	rs, err := services.OpenPlatformApp.Server.Notify(context.Request, func(event *openplatform.Callback, decrypted []byte, infoType string) (result interface{}) {

		result = kernel.SUCCESS_EMPTY_RESPONSE

		switch infoType {
		case openplatform.EVENT_COMPONENT_VERIFY_TICKET:
			msg := &openplatform.EventVerifyTicket{}
			err = xml.Unmarshal(decrypted, msg)
			//fmt.Dump(event)
			if err != nil {
				return err
			}
			// set ticket in redis
			err = services.OpenPlatformApp.VerifyTicket.SetTicket(msg.ComponentVerifyTicket)
			if err != nil {
				return err
			}

			fmt.Dump(msg)
		case openplatform.EVENT_AUTHORIZED:

		}
		return result
	})

	if err != nil {
		panic(err)
	}

	err = rs.Write(context.Writer)
	if err != nil {
		panic(err)
	}

}

func APIOpenPlatformCallbackWithApp(context *gin.Context) {

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

	err = rs.Write(context.Writer)
	if err != nil {
		panic(err)
	}

}
