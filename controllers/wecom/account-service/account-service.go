package account_service

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/accountService/request"
	"github.com/gin-gonic/gin"
	"net/http"
	"power-wechat-tutorial/services"
)

// 添加客服帐号
// https://work.weixin.qq.com/api/doc/90000/90135/94648
func APIAccountServiceAccountAdd(c *gin.Context) {

	mediaID := c.DefaultQuery("mediaID", "294DpAog3YA5b9rTK4PjjfRfYLO0L5qpDHAJIzhhQ2jAEWjb9i661Q4lk8oFnPtmj")

	res, err := services.WeComApp.AccountService.Add(c.Request.Context(), "新建的客服帐号", mediaID)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)
}

// 删除客服帐号
// https://work.weixin.qq.com/api/doc/90000/90135/94663
func APIAccountServiceAccountDel(c *gin.Context) {

	openKFID := c.DefaultQuery("openKFID", "wkAJ2GCAAAZSfhHCt7IFSvLKtMPxyJTw")

	res, err := services.WeComApp.AccountService.Del(c.Request.Context(), openKFID)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)
}

// 修改客服帐号
// https://work.weixin.qq.com/api/doc/90000/90135/94664
func APIAccountServiceAccountUpdate(c *gin.Context) {

	options := &request.RequestAccountUpdate{
		OpenKFID: c.DefaultQuery("openKFID", "wkAJ2GCAAAZSfhHCt7IFSvLKtMPxyJTw"),
		Name:     "修改客服名",
		MediaID:  c.DefaultQuery("mediaID", "294DpAog3YA5b9rTK4PjjfRfYLO0L5qpDHAJIzhhQ2jAEWjb9i661Q4lk8oFnPtmj"),
	}

	res, err := services.WeComApp.AccountService.Update(c.Request.Context(), options)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)
}

// 获取客服帐号列表
// https://work.weixin.qq.com/api/doc/90000/90135/94661
func APIAccountServiceAccountList(c *gin.Context) {

	res, err := services.WeComApp.AccountService.List(c.Request.Context())

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)
}

// 获取客服帐号链接
// https://work.weixin.qq.com/api/doc/90000/90135/94665
func APIAccountServiceAddContactWay(c *gin.Context) {

	openKFID := c.DefaultQuery("openKFID", "wkAJ2GCAAAZSfhHCt7IFSvLKtMPxyJTw")
	scene := c.DefaultQuery("scene", "1234")

	res, err := services.WeComApp.AccountService.AddContactWay(c.Request.Context(), openKFID, scene)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)
}
