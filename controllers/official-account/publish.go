package official_account

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/officialAccount/publish/request"
	"github.com/gin-gonic/gin"
	"net/http"
	"power-wechat-tutorial/services"
	"strconv"
)

// GetCustomerList 获取所有客服
func APIDraftAdd(ctx *gin.Context) {
	data, err := services.OfficialAccountApp.Publish.DraftAdd(ctx.Request.Context(), &request.RequestDraftAdd{
		Articles: []*request.Article{
			&request.Article{
				Title:              "testTitle",
				Author:             "testAuthor",
				Digest:             "testDigest",
				Content:            "testContent",
				ContentSourceUrl:   "test url",
				ThumbMediaId:       "test",
				NeedOpenComment:    0,
				OnlyFansCanComment: 1,
			},
		},
	})
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, data)
}

func APIDraftGet(ctx *gin.Context) {
	mediaID := ctx.Query("mediaID")
	data, err := services.OfficialAccountApp.Publish.DraftGet(ctx.Request.Context(), mediaID)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, data)
}

func APIDraftDelete(ctx *gin.Context) {

	mediaID := ctx.Query("mediaID")

	data, err := services.OfficialAccountApp.Publish.DraftDelete(ctx.Request.Context(), mediaID)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, data)
}

func APIDraftUpdate(ctx *gin.Context) {
	data, err := services.OfficialAccountApp.Publish.DraftUpdate(ctx.Request.Context(), &request.RequestDraftUpdate{
		MediaId: "123",
		Index:   1,
		Articles: &request.Article{
			Title:              "testTitle",
			Author:             "testAuthor",
			Digest:             "testDigest",
			Content:            "testContent",
			ContentSourceUrl:   "test url",
			ThumbMediaId:       "test",
			NeedOpenComment:    0,
			OnlyFansCanComment: 1,
		},
	})
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, data)
}

func APIDraftCount(ctx *gin.Context) {
	data, err := services.OfficialAccountApp.Publish.DraftCount(ctx.Request.Context())
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, data)
}

func APIDraftBatchGet(ctx *gin.Context) {
	offset, _ := strconv.Atoi(ctx.Query("offset"))
	count, _ := strconv.Atoi(ctx.Query("count"))
	noContent, _ := strconv.Atoi(ctx.Query("noContent"))

	data, err := services.OfficialAccountApp.Publish.DraftBatchGet(ctx.Request.Context(), &request.RequestBatchGet{
		Offset:    offset,
		Count:     count,
		NoContent: noContent,
	})
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, data)
}

func APIDraftSwitch(ctx *gin.Context) {
	data, err := services.OfficialAccountApp.Publish.DraftSwitch(ctx.Request.Context())
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, data)
}

func APIDraftCheckSwitch(ctx *gin.Context) {
	data, err := services.OfficialAccountApp.Publish.DraftCheckSwitch(ctx.Request.Context())
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, data)
}

func APIPublishSubmit(ctx *gin.Context) {
	data, err := services.OfficialAccountApp.Publish.PublishSubmit(ctx.Request.Context(), "123")
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, data)
}

func APIPublishGet(ctx *gin.Context) {
	publishID := ctx.Query("publishID")
	data, err := services.OfficialAccountApp.Publish.PublishGet(ctx.Request.Context(), publishID)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, data)
}

func APIPublishDelete(ctx *gin.Context) {
	articleID := ctx.Query("articleID")
	data, err := services.OfficialAccountApp.Publish.PublishDelete(ctx.Request.Context(), articleID, 1)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, data)
}

func APIPublishGetArticle(ctx *gin.Context) {

	articleID := ctx.Query("articleID")

	data, err := services.OfficialAccountApp.Publish.PublishGetArticle(ctx.Request.Context(), articleID)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, data)
}

func APIPublishBatchGet(ctx *gin.Context) {
	data, err := services.OfficialAccountApp.Publish.PublishBatchGet(ctx.Request.Context(), &request.RequestBatchGet{
		Offset:    0,
		Count:     10,
		NoContent: 1,
	})
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, data)
}
