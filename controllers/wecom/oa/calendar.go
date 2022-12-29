package oa

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/power"
	"github.com/gin-gonic/gin"
	"net/http"
	"power-wechat-tutorial/services"
	"strconv"
)

// 创建日历
// https://work.weixin.qq.com/api/doc/90000/90135/93647
func APICalendarAdd(c *gin.Context) {
	agentId := c.DefaultQuery("agentID", "1000014")
	agentID, _ := strconv.Atoi(agentId)

	calendar := &power.HashMap{
		"organizer":      "userid1",
		"readonly":       1,
		"set_as_default": 1,
		"summary":        "test_summary",
		"color":          "#FF3030",
		"description":    "test_describe",
		"shares": []power.HashMap{
			power.HashMap{
				"userid": "userid2",
			},
			power.HashMap{
				"userid":   "userid3",
				"readonly": 1,
			},
		},
	}

	res, err := services.WeComApp.OACalendar.Add(c.Request.Context(), calendar, agentID)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)
}

// 更新日历
// https://work.weixin.qq.com/api/doc/90000/90135/93647
func APICalendarUpdate(c *gin.Context) {

	calendar := &power.HashMap{
		"cal_id":      "wcjgewCwAAqeJcPI1d8Pwbjt7nttzAAA",
		"readonly":    1,
		"summary":     "test_summary",
		"color":       "#FF3030",
		"description": "test_describe_1",
		"shares": []power.HashMap{
			power.HashMap{
				"userid": "userid1",
			},
			power.HashMap{
				"userid":   "userid2",
				"readonly": 1,
			},
		},
	}

	res, err := services.WeComApp.OACalendar.Update(c.Request.Context(), calendar)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)
}

// 获取日历详情
// https://work.weixin.qq.com/api/doc/90000/90135/93647
func APICalendarGet(c *gin.Context) {
	calIDList := []string{
		c.DefaultQuery("calID", "wcjgewCwAAqeJcPI1d8Pwbjt7nttzAAA"),
	}
	res, err := services.WeComApp.OACalendar.Get(c.Request.Context(), calIDList)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)
}

// 删除日历
// https://work.weixin.qq.com/api/doc/90000/90135/93647
func APICalendarDel(c *gin.Context) {
	calID := c.DefaultQuery("calID", "wcjgewCwAAqeJcPI1d8Pwbjt7nttzAAA")
	res, err := services.WeComApp.OACalendar.Del(c.Request.Context(), calID)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)
}
