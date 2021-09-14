package miniprogram

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/power"
	"github.com/gin-gonic/gin"
	"power-wechat-tutorial/services"
)

func APIUpdatableMessageCreateActivityID(c *gin.Context) {

	openID, exist := c.GetQuery("openID")
	if !exist {
		panic("parameter open id expected")
	}

	rs, err := services.MiniProgramApp.UpdatableMessage.CreateActivityID("", openID)

	if err != nil {
		panic(err)
	}

	c.JSON(200, rs)
}

func APIUpdatableMessageUpdatableMessage(c *gin.Context) {

	activityID, exist := c.GetQuery("activityID")
	if !exist {
		panic("parameter open id expected")
	}

	rs, err := services.MiniProgramApp.UpdatableMessage.SetUpdatableMsg(activityID, 0, &power.HashMap{
		"parameter_list": []power.StringMap{
			power.StringMap{
				"name": "member_count",
				"value": "2",
			},
			power.StringMap{
				"name":  "room_limit",
				"value": "5",
			},
		},
	})

	if err != nil {
		panic(err)
	}

	c.JSON(200, rs)
}
