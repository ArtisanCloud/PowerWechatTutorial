package payment

import (
	"github.com/ArtisanCloud/go-libs/http/response"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path"
	"power-wechat-tutorial/services"
)

// 下载账单API为通用接口，交易/资金账单都可以通过该接口获取到对应的账单
// https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter3_1_8.shtml
func APIBillDownloadURL(c *gin.Context) {

	requestDownload := &response.RequestDownload{
		"SHA1",
		"442c2363a7e014b7f7cf3e2c558375bcf385951d",
		"https://cdn.pixabay.com/photo/2015/04/23/22/00/tree-736885_1280.jpg",
	}
	homePath, _ := os.UserHomeDir()
	filePath := path.Join(homePath, "Desktop/download-url")

	rs, err := services.PaymentService.Bill.DownloadBill(requestDownload, filePath)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, rs)

}
