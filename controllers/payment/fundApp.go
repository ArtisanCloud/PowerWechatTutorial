package payment

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/payment/fundApp/request"
	"github.com/gin-gonic/gin"
	"net/http"
	"power-wechat-tutorial/services"
)

func APITransferBills(c *gin.Context) {

	req := &request.RequestTransferBills{
		Appid:              "Appid",
		OutBillNo:          "OutBillNo",
		TransferSceneId:    "TransferSceneId",
		Openid:             "Openid",
		UserName:           "UserName",
		TransferAmount:     1,
		TransferRemark:     "TransferRemark",
		NotifyUrl:          "NotifyUrl",
		UserRecvPerception: "UserRecvPerception",
		TransferSceneReportInfos: []request.TransferSceneReportInfo{
			{
				InfoType:    "InfoType",
				InfoContent: "InfoContent",
			},
		},
	}
	ctx := c.Request.Context()
	//fmt.Dump(ctx)
	rs, err := services.PaymentApp.FundApp.TransferBills(ctx, req)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, rs)

}
