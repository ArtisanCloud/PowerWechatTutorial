package transfer

import (
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/payment/transfer/request"
	"github.com/gin-gonic/gin"
	"net/http"
	"power-wechat-tutorial/services"
)

func APIBatchTransfer(c *gin.Context) {

	appId := services.PaymentApp.GetConfig().GetString("app_id", "")

	transfer := &request.RequestTransferBatch{
		AppID:       appId,
		OutBatchNO:  "0010010404201411170000046522",
		BatchName:   "batch-1",
		BatchRemark: "batch-1-remark",
		TotalAmount: 30,
		TotalNum:    1,
		TransferDetailList: []*request.TransferDetail{
			&request.TransferDetail{
				OutDetailNO:    "00100104042014111700000465221",
				TransferAmount: 30,
				TransferRemark: "remark",
				OpenID:         "o4QEk5Kc_y8QTrENCpKoxYhS4jkg",
				UserName:       object.NewNullString("username", true),
			},
		},
	}

	payResult, err := services.PaymentApp.TransferBatch.Batch(c.Request.Context(), transfer)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, payResult)
}

func APIQueryBatchOrder(c *gin.Context) {

	rs, err := services.PaymentApp.TransferBatch.QueryBatch(
		c.Request.Context(),
		"{batchID}}",
		true,
		0,
		10,
		"")
	if err != nil {
		panic(nil)
	}

	c.JSON(http.StatusOK, rs)
}

func APIQueryBatchOrderDetail(c *gin.Context) {

	rs, err := services.PaymentApp.TransferBatch.QueryBatchDetail(
		c.Request.Context(),
		"{batchID}}",
		"{batchDetailID}}")

	if err != nil {
		panic(nil)
	}

	c.JSON(http.StatusOK, rs)
}
