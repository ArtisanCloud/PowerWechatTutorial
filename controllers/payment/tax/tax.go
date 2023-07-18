package tax

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/payment/tax/request"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"power-wechat-tutorial/services"
)

func APIApplyForCardTemplate(c *gin.Context) {
	//traceNo := c.Query("traceNo")
	//log.Printf("traceNo: %s", traceNo)

	para := &request.RequestApplyForCardTemplate{
		CardAppid: "wxa5fa5b1adab7aa1f",
		CardTemplateInformation: &request.CardTemplateInformation{
			PayeeName: "123",
			LogoUrl:   "123",
			CustomCell: &request.CustomCell{
				Words:               "123",
				Description:         "321",
				JumpUrl:             "3213",
				MiniProgramUserName: "1231",
				MiniProgramPath:     "123",
			},
		},
	}

	rs, err := services.PaymentApp.Tax.ApplyForCardTemplate(c.Request.Context(), para)
	if err != nil {
		log.Println("出错了： ", err)
		c.String(400, err.Error())
		return
	}
	c.JSON(http.StatusOK, rs)

}
