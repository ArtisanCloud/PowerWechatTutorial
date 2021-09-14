package miniprogram

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/power"
	"github.com/gin-gonic/gin"
	"power-wechat-tutorial/services"
)

// 根据提交的用户信息数据获取用户的安全等级 risk_rank，无需用户授权。
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/safety-control-capability/riskControl.getUserRiskRank.html
func APIRiskControlGetUserRiskRank(c *gin.Context) {

	appID, exist := c.GetQuery("appID")
	if !exist {
		panic("parameter app id expected")
	}

	openID, exist := c.GetQuery("openID")
	if !exist {
		panic("parameter open id expected")
	}



	options := &power.HashMap{
		"appid":         appID,
		"openid":        openID,
		"scene":         1,
		"mobile_no":     "12345678",
		"bank_card_no":  "******",
		"cert_no":       "*******",
		"client_ip":     "******",
		"email_address": "***@qq.com",
		"extended_info": "",
	}

	rs, err := services.MiniProgramApp.RiskControl.GetUserRiskRank(options)

	if err != nil {
		panic(err)
	}

	c.JSON(200, rs)
}
