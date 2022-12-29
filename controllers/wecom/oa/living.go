package oa

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/oa/living/request"
	"github.com/gin-gonic/gin"
	"net/http"
	"power-wechat-tutorial/services"
)

// 创建预约直播
// https://work.weixin.qq.com/api/doc/90000/90135/93637
func APILivingCreate(c *gin.Context) {

	options := &request.RequestLivingCreate{
		AnchorUserID:         "zhangsan",
		Theme:                "theme",
		LivingStart:          1600000000,
		LivingDuration:       3600,
		Description:          "test description",
		Type:                 4,
		AgentID:              1000014,
		RemindTime:           60,
		ActivityCoverMediaID: "MEDIA_ID",
		ActivityShareMediaID: "MEDIA_ID",
		ActivityDetail: &power.HashMap{
			"description": "活动描述，非活动类型的直播不用传",
			"image_list": []string{
				"MEDIA_ID_1",
				"MEDIA_ID_2",
			},
		},
	}
	res, err := services.WeComApp.OALiving.Create(c.Request.Context(), options)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)
}

func APILivingModify(c *gin.Context) {

	options := &request.RequestLivingModify{
		LivingID:       c.DefaultQuery("livingID", "XXXXXXXXX"),
		Theme:          "theme",
		LivingStart:    1600100000,
		LivingDuration: 3600,
		Description:    "test description",
		Type:           1,
		RemindTime:     60,
	}
	res, err := services.WeComApp.OALiving.Modify(c.Request.Context(), options)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)
}

func APILivingCancel(c *gin.Context) {
	livingID := c.DefaultQuery("livingID", "XXXXXXXXX")
	res, err := services.WeComApp.OALiving.Cancel(c.Request.Context(), livingID)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)
}

func APILivingDeleteReplayData(c *gin.Context) {
	livingID := c.DefaultQuery("livingID", "XXXXXXXXX")
	res, err := services.WeComApp.OALiving.DeleteReplayData(c.Request.Context(), livingID)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)
}

func APILivingGetLivingCode(c *gin.Context) {
	livingID := c.DefaultQuery("livingID", "XXXXXXXXX")
	openID := c.DefaultQuery("openID", "XXXXXXXXX")
	res, err := services.WeComApp.OALiving.GetLivingCode(c.Request.Context(), livingID, openID)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)
}

func APILivingGetUserAllLivingID(c *gin.Context) {
	userID := c.DefaultQuery("userID", "XXXXXXXXX")
	cursor := "CURSOR"
	limit := 20
	res, err := services.WeComApp.OALiving.GetUserAllLivingID(c.Request.Context(), userID, cursor, limit)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)
}

func APILivingGetLivingInfo(c *gin.Context) {
	livingID := c.DefaultQuery("livingID", "XXXXXXXXX")
	res, err := services.WeComApp.OALiving.GetLivingInfo(c.Request.Context(), livingID)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)
}

func APILivingGetWatchStat(c *gin.Context) {
	livingID := c.DefaultQuery("livingID", "XXXXXXXXX")
	nextKey := c.DefaultQuery("nextKey", "XXXXXXXXX")

	res, err := services.WeComApp.OALiving.GetWatchStat(c.Request.Context(), livingID, nextKey)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)
}

func APILivingGetLivingShareInfo(c *gin.Context) {
	wwShareCode := c.DefaultQuery("wwShareCode", "XXXXXXXXX")

	res, err := services.WeComApp.OALiving.GetLivingShareInfo(c.Request.Context(), wwShareCode)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)
}
