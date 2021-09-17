package externalContact

import (
	"github.com/ArtisanCloud/power-wechat/src/work/externalContact/moment/request"
	"github.com/gin-gonic/gin"
	"net/http"
	"power-wechat-tutorial/services"
)

func APIExternalContactGetMomentList(c *gin.Context) {

	options := &request.RequestGetMomentList{
		StartTime:  c.get,
		EndTime:    c.GetInt64("endTime"),
		Creator:    "walle",
		FilterType: 1,
		Cursor:     c.DefaultQuery("cursor", "CURSOR"),
		Limit:      10,
	}

	res, err := services.WeComContactApp.ExternalContactMoment.GetMomentList(options)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)
}

func APIExternalContactGetMomentTask(c *gin.Context) {


	res, err := services.WeComContactApp.ExternalContactMoment.Mom(options)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)

}

func APIExternalContactGetMomentCustomerList(c *gin.Context) {

}

func APIExternalContactGetMomentSendResult(c *gin.Context) {

}

func APIExternalContactGetMomentComments(c *gin.Context) {

}

func APIExternalContactMomentStrategyList(c *gin.Context) {

}

func APIExternalContactMomentStrategyGet(c *gin.Context) {

}

func APIExternalContactMomentStrategyGetRange(c *gin.Context) {

}

func APIExternalContactMomentStrategyCreate(c *gin.Context) {

}

func APIExternalContactMomentStrategyEdit(c *gin.Context) {

}

func APIExternalContactMomentStrategyDel(c *gin.Context) {

}
