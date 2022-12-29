package message

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/power"
	"github.com/gin-gonic/gin"
	"net/http"
	"power-wechat-tutorial/services"
)

func APIExternalContactMessageSendText(c *gin.Context) {

	options := &power.HashMap{
		"to_external_user":  []string{"external_userid1", "external_userid2"},
		"to_parent_userid":  []string{"parent_userid1", "parent_userid2"},
		"to_student_userid": []string{"student_userid1", "student_userid2"},
		"to_party":          []string{"partyid1", "partyid2"},
		"toall":             0,
		"msgtype":           "text",
		"agentid":           1,
		"text": power.HashMap{
			"content": "你的快递已到，请携带工卡前往邮件中心领取。\n出发前可查看<a href=\"http://work.weixin.qq.com\">邮件中心视频实况</a>，聪明避开排队。",
		},
		"enable_id_trans":          0,
		"enable_duplicate_check":   0,
		"duplicate_check_interval": 1800,
	}

	res, err := services.WeComApp.MessageExternalContact.Send(c.Request.Context(), options)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)
}

func APIExternalContactMessageSendImage(c *gin.Context) {
	options := &power.HashMap{
		"to_external_user":  []string{"external_userid1", "external_userid2"},
		"to_parent_userid":  []string{"parent_userid1", "parent_userid2"},
		"to_student_userid": []string{"student_userid1", "student_userid2"},
		"to_party":          []string{"partyid1", "partyid2"},
		"toall":             0,
		"msgtype":           "image",
		"agentid":           1,
		"image": power.HashMap{
			"media_id": "MEDIA_ID",
		},
		"enable_duplicate_check":   0,
		"duplicate_check_interval": 1800,
	}
	res, err := services.WeComApp.MessageExternalContact.Send(c.Request.Context(), options)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)
}

func APIExternalContactMessageSendVoice(c *gin.Context) {

	options := &power.HashMap{
		"to_external_user":  []string{"external_userid1", "external_userid2"},
		"to_parent_userid":  []string{"parent_userid1", "parent_userid2"},
		"to_student_userid": []string{"student_userid1", "student_userid2"},
		"to_party":          []string{"partyid1", "partyid2"},
		"toall":             0,
		"msgtype":           "voice",
		"agentid":           1,
		"voice": power.HashMap{
			"media_id": "MEDIA_ID",
		},
		"enable_duplicate_check":   0,
		"duplicate_check_interval": 1800,
	}
	res, err := services.WeComApp.MessageExternalContact.Send(c.Request.Context(), options)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)
}

func APIExternalContactMessageSendVideo(c *gin.Context) {

	options := &power.HashMap{
		"to_external_user":  []string{"external_userid1", "external_userid2"},
		"to_parent_userid":  []string{"parent_userid1", "parent_userid2"},
		"to_student_userid": []string{"student_userid1", "student_userid2"},
		"to_party":          []string{"partyid1", "partyid2"},
		"toall":             0,
		"msgtype":           "video",
		"agentid":           1,
		"video": power.HashMap{
			"media_id":    "MEDIA_ID",
			"title":       "Title",
			"description": "Description",
		},
		"enable_duplicate_check":   0,
		"duplicate_check_interval": 1800,
	}
	res, err := services.WeComApp.MessageExternalContact.Send(c.Request.Context(), options)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)
}

func APIExternalContactMessageSendFile(c *gin.Context) {
	options := &power.HashMap{
		"to_external_user":  []string{"external_userid1", "external_userid2"},
		"to_parent_userid":  []string{"parent_userid1", "parent_userid2"},
		"to_student_userid": []string{"student_userid1", "student_userid2"},
		"to_party":          []string{"partyid1", "partyid2"},
		"toall":             0,
		"msgtype":           "file",
		"agentid":           1,
		"file": power.HashMap{
			"media_id": "1Yv-zXfHjSjU-7LH-GwtYqDGS-zz6w22KmWAT5COgP7o",
		},
		"enable_duplicate_check":   0,
		"duplicate_check_interval": 1800,
	}
	res, err := services.WeComApp.MessageExternalContact.Send(c.Request.Context(), options)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)
}

func APIExternalContactMessageSendTextCard(c *gin.Context) {

}

func APIExternalContactMessageSendNews(c *gin.Context) {

	options := &power.HashMap{
		"to_external_user":  []string{"external_userid1", "external_userid2"},
		"to_parent_userid":  []string{"parent_userid1", "parent_userid2"},
		"to_student_userid": []string{"student_userid1", "student_userid2"},
		"to_party":          []string{"partyid1", "partyid2"},
		"toall":             0,
		"msgtype":           "news",
		"agentid":           1,
		"news": power.HashMap{
			"articles": []power.HashMap{
				{
					"title":       "中秋节礼品领取",
					"description": "今年中秋节公司有豪礼相送",
					"url":         "URL",
					"picurl":      "http://res.mail.qq.com/node/ww/wwopenmng/images/independent/doc/test_pic_msg1.png",
				},
			},
		},
		"enable_id_trans":          0,
		"enable_duplicate_check":   0,
		"duplicate_check_interval": 1800,
	}
	res, err := services.WeComApp.MessageExternalContact.Send(c.Request.Context(), options)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)

}

func APIExternalContactMessageSendMPNews(c *gin.Context) {

	options := &power.HashMap{
		"to_external_user":  []string{"external_userid1", "external_userid2"},
		"to_parent_userid":  []string{"parent_userid1", "parent_userid2"},
		"to_student_userid": []string{"student_userid1", "student_userid2"},
		"to_party":          []string{"partyid1", "partyid2"},
		"toall":             0,
		"msgtype":           "mpnews",
		"agentid":           1,
		"mpnews": power.HashMap{
			"articles": []power.HashMap{
				{
					"title":              "Title",
					"thumb_media_id":     "MEDIA_ID",
					"author":             "Author",
					"content_source_url": "URL",
					"content":            "Content",
					"digest":             "Digest description",
				},
			},
		},
		"enable_id_trans":          0,
		"enable_duplicate_check":   0,
		"duplicate_check_interval": 1800,
	}
	res, err := services.WeComApp.MessageExternalContact.Send(c.Request.Context(), options)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)

}

func APIExternalContactMessageSendMiniProgram(c *gin.Context) {

	options := &power.HashMap{
		"to_external_user":  []string{"external_userid1", "external_userid2"},
		"to_parent_userid":  []string{"parent_userid1", "parent_userid2"},
		"to_student_userid": []string{"student_userid1", "student_userid2"},
		"to_party":          []string{"partyid1", "partyid2"},
		"toall":             0,
		"agentid":           1,
		"msgtype":           "miniprogram",
		"miniprogram": power.HashMap{
			"appid":          "APPID",
			"title":          "欢迎报名夏令营",
			"thumb_media_id": "MEDIA_ID",
			"pagepath":       "PAGE_PATH",
		},
		"enable_id_trans":          0,
		"enable_duplicate_check":   0,
		"duplicate_check_interval": 1800,
	}
	res, err := services.WeComApp.MessageExternalContact.Send(c.Request.Context(), options)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)

}
