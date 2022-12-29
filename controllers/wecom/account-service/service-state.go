package account_service

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"power-wechat-tutorial/services"
	"strconv"
)

// 获取会话状态
// https://work.weixin.qq.com/api/doc/90000/90135/94669
func APIAccountServiceStateGet(c *gin.Context) {

	openKFID := c.DefaultQuery("openKFID", "kfxxxxxxxxxxxxxx")
	externalUserID := c.DefaultQuery("externalUserID", "wmxxxxxxxxxxxxxxxxxx")

	res, err := services.WeComApp.AccountServiceState.Get(c.Request.Context(), openKFID, externalUserID)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)
}

// 变更会话状态
// https://work.weixin.qq.com/api/doc/90000/90135/94669
func APIAccountServiceStateTrans(c *gin.Context) {

	openKFID := c.DefaultQuery("openKFID", "kfxxxxxxxxxxxxxx")
	externalUserID := c.DefaultQuery("externalUserID", "kfxxxxxxxxxxxxxx")
	serviceState := c.DefaultQuery("serviceState", "kfxxxxxxxxxxxxxx")
	servicerUserID := c.DefaultQuery("servicerUserID", "kfxxxxxxxxxxxxxx")

	state, _ := strconv.Atoi(serviceState)

	res, err := services.WeComApp.AccountServiceState.Trans(c.Request.Context(), openKFID, externalUserID, state, servicerUserID)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)
}
