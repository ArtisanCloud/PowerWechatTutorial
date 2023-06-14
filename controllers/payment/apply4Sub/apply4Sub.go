package partner

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/payment/apply4Sub/request"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"power-wechat-tutorial/services"
)

func APIApplyFor(c *gin.Context) {
	//traceNo := c.Query("traceNo")
	//log.Printf("traceNo: %s", traceNo)

	para := &request.RequestApplyForBusiness{}

	rs, err := services.PaymentApp.Apply4Sub.ApplyForBusiness(c.Request.Context(), para)
	if err != nil {
		log.Println("出错了： ", err)
		c.String(400, err.Error())
		return
	}
	c.JSON(http.StatusOK, rs)

}
