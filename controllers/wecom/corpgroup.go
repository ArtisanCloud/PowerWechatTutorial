package wecom

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"power-wechat-tutorial/services"
	"strconv"
)

// 获取应用共享信息
// https://open.work.weixin.qq.com/api/doc/90000/90135/93403
func APICorpGroupCorpListAppShareInfo(c *gin.Context) {

	agentId := c.DefaultQuery("agentID", "")
	agentID, _ := strconv.Atoi(agentId)

	res, err := services.WeComApp.CorpGroup.GetAppShareInfo(c.Request.Context(), agentID)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)
}

// 获取下级企业的access_token
// https://open.work.weixin.qq.com/api/doc/90000/90135/93359
func APICorpGroupCorpGetToken(c *gin.Context) {

	corpID := c.DefaultQuery("corpID", "")
	agentID := c.DefaultQuery("agentID", "")
	res, err := services.WeComApp.CorpGroup.GetToken(c.Request.Context(), corpID, agentID)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)
}

// 获取下级企业的小程序session
// https://open.work.weixin.qq.com/api/doc/90000/90135/93355
func APICorpGroupMiniProgramTransferSession(c *gin.Context) {

	userID := c.DefaultQuery("userID", "matrix-x")
	sessionKey := ""
	res, err := services.WeComApp.CorpGroup.GetMiniProgramTransferSession(c.Request.Context(), userID, sessionKey)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)
}
