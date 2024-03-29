package miniprogram

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/miniProgram/serviceMarket/request"
	"github.com/gin-gonic/gin"
	"net/http"
	"power-wechat-tutorial/services"
)

// 调用服务平台提供的服务
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/service-market/serviceMarket.invokeService.html
func APIServiceMarketInvokeService(c *gin.Context) {

	serviceID, exist := c.GetQuery("serviceID")
	if !exist {
		panic("parameter service id expected")
	}

	apiName, exist := c.GetQuery("apiName")
	if !exist {
		panic("parameter api name expected")
	}

	clientMsgID, exist := c.GetQuery("clientMsgID")
	if !exist {
		panic("parameter client msg id expected")
	}

	serviceData := &power.HashMap{
		"img_url":   "http://mmbiz.qpic.cn/mmbiz_jpg/7UFjuNbYxibu66xSqsQqKcuoGBZM77HIyibdiczeWibdMeA2XMt5oibWVQMgDibriazJSOibLqZxcO6DVVcZMxDKgeAtbQ/0",
		"data_type": 3,
		"ocr_type":  1,
	}

	rs, err := services.MiniProgramApp.ServiceMarket.InvokeService(c.Request.Context(),
		&request.RequestServiceMarket{
			Service:     serviceID,
			Api:         apiName,
			ClientMsgID: clientMsgID,
			Data:        serviceData,
		})

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, rs)
}
