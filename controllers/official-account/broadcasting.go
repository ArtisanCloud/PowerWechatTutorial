package official_account

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/officialAccount/broadcasting/request"
	"github.com/gin-gonic/gin"
	"net/http"
	"power-wechat-tutorial/services"
)

var testUserOpenID = "oF-S35jfUZNlcJ8RvtpqKaRbUON4"

func BroadcastSendText(ctx *gin.Context) {
	openID := ctx.DefaultPostForm("openID", testUserOpenID)
	message := ctx.DefaultPostForm("message", "hello, broadcasting test...")
	data, err := services.OfficialAccountApp.Broadcasting.SendText(message, &request.Reception{
		ToUser: []string{openID},
		Filter: &request.Filter{
			IsToAll: false,
			TagID:   0,
		},
	}, &power.HashMap{})

	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, data)
}

func BroadcastSendImage(ctx *gin.Context) {
	openID := ctx.DefaultPostForm("openID", testUserOpenID)
	mediaID := ctx.DefaultPostForm("mediaID", "")
	data, err := services.OfficialAccountApp.Broadcasting.SendImage(mediaID, &request.Reception{
		ToUser: []string{openID},
		Filter: &request.Filter{
			IsToAll: false,
			TagID:   0,
		},
	}, &power.HashMap{})

	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, data)
}

func BroadcastSendNews(ctx *gin.Context) {
	openID := ctx.DefaultPostForm("openID", testUserOpenID)
	mediaID := ctx.DefaultPostForm("mediaID", "")
	data, err := services.OfficialAccountApp.Broadcasting.SendNews(mediaID, &request.Reception{
		ToUser: []string{openID},
		Filter: &request.Filter{
			IsToAll: false,
			TagID:   0,
		},
	}, &power.HashMap{})

	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, data)
}

func BroadcastSendVoice(ctx *gin.Context) {
	openID := ctx.DefaultPostForm("openID", testUserOpenID)
	mediaID := ctx.DefaultPostForm("mediaID", "")
	data, err := services.OfficialAccountApp.Broadcasting.SendVoice(mediaID, &request.Reception{
		ToUser: []string{openID},
		Filter: &request.Filter{
			IsToAll: false,
			TagID:   0,
		},
	}, &power.HashMap{})

	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, data)
}

func BroadcastSendVideo(ctx *gin.Context) {
	openID := ctx.DefaultPostForm("openID", testUserOpenID)
	mediaID := ctx.DefaultPostForm("mediaID", "")
	data, err := services.OfficialAccountApp.Broadcasting.SendVideo(mediaID, &request.Reception{
		ToUser: []string{openID},
		Filter: &request.Filter{
			IsToAll: false,
			TagID:   0,
		},
	}, &power.HashMap{})

	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, data)
}

func BroadcastSendCard(ctx *gin.Context) {
	openID := ctx.DefaultPostForm("openID", testUserOpenID)
	mediaID := ctx.DefaultPostForm("mediaID", "")
	data, err := services.OfficialAccountApp.Broadcasting.SendCard(mediaID, &request.Reception{
		ToUser: []string{openID},
		Filter: &request.Filter{
			IsToAll: false,
			TagID:   0,
		},
	}, &power.HashMap{})

	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, data)
}

func BroadcastSendPreview(ctx *gin.Context) {
	openID := ctx.DefaultPostForm("openID", testUserOpenID)
	message := ctx.DefaultPostForm("message", "")
	data, err := services.OfficialAccountApp.Broadcasting.PreviewText(message, &request.Reception{
		ToUser: []string{openID},
		Filter: &request.Filter{
			IsToAll: false,
			TagID:   0,
		},
	}, "")

	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, data)
}

func BroadcastDelete(ctx *gin.Context) {
	msgID := ctx.Query("msgID")
	data, err := services.OfficialAccountApp.Broadcasting.Delete(msgID, 0)

	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, data)
}

func BroadcastStatus(ctx *gin.Context) {
	msgID := ctx.Query("msgID")
	data, err := services.OfficialAccountApp.Broadcasting.Status(msgID)

	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, data)
}
