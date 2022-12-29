package wecom

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/power"
	"github.com/gin-gonic/gin"
	"net/http"
	"power-wechat-tutorial/services"
)

// 获取会话内容存档开启成员列表
// https://open.work.weixin.qq.com/api/doc/90000/90135/91614
func APIMsgAuditGetPermitUserList(c *gin.Context) {

	msgType := "1"
	res, err := services.WeComApp.MsgAudit.GetPermitUsersList(c.Request.Context(), msgType)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)
}

// 获取会话同意情况
// https://open.work.weixin.qq.com/api/doc/90000/90135/91782
func APIMsgAuditCheckSingleAgree(c *gin.Context) {
	info := []*power.StringMap{
		{"userid": "XuJinSheng", "exteranalopenid": "wmeDKaCQAAGd9oGiQWxVsAKwV2HxNAAA"},
		{"userid": "XuJinSheng", "exteranalopenid": "wmeDKaCQAAIQ_p7ACn_jpLVBJSGocAAA"},
		{"userid": "XuJinSheng", "exteranalopenid": "wmeDKaCQAAPE_p7ABnxkpLBBJSGocAAA"},
	}
	res, err := services.WeComApp.MsgAudit.CheckSingleAgree(c.Request.Context(), info)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)
}

// 获取会话同意情况
// https://open.work.weixin.qq.com/api/doc/90000/90135/91782
func APIMsgAuditCheckRoomAgree(c *gin.Context) {

	roomID := c.DefaultQuery("roomID", "wrjc7bDwAASxc8tZvBErFE02BtPWyAAA")

	res, err := services.WeComApp.MsgAudit.CheckRoomAgree(c.Request.Context(), roomID)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)
}

func APIMsgAuditGroupChatGet(c *gin.Context) {

	roomID := c.DefaultQuery("roomID", "wrjc7bDwAASxc8tZvBErFE02BtPWyAAA")
	res, err := services.WeComApp.MsgAudit.GroupChatGet(c.Request.Context(), roomID)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)
}
