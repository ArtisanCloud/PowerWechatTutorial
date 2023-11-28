package miniprogram

import (
	"github.com/ArtisanCloud/PowerLibs/v3/fmt"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/miniProgram/industry/miniDrama/request"
	"github.com/gin-gonic/gin"
	"power-wechat-tutorial/services"
	"time"
)

func APIVideoMediaUploadByURL(ctx *gin.Context) {

	var params []*request.VideoMediaUploadByURLRequest

	params = append(params, &request.VideoMediaUploadByURLRequest{
		MediaUrl:  "xxxx.mp4",
		CoverUrl:  "xxxx",
		MediaName: "xxxx",
	})
	params = append(params, &request.VideoMediaUploadByURLRequest{
		MediaUrl:  "xxxx",
		CoverUrl:  "xxxx",
		MediaName: "xxx",
	})

	for _, param := range params {

		result, err := services.MiniProgramApp.MiniDramaVOD.VideoMediaUploadByURL(ctx, param)

		if err != nil {
			panic(err)
		}

		fmt.Dump(result)
	}

}

func APISearchMediaByTaskId(ctx *gin.Context) {

	taskIDs := []int64{111, 222, 333, 444, 555}

	for _, task := range taskIDs {

		result, err := services.MiniProgramApp.MiniDramaVOD.SearchMediaByTaskId(ctx, task)

		if err != nil {
			panic(err)
		}

		fmt.Dump(result)
	}

}

func APIGetMediaList(ctx *gin.Context) {

	result, err := services.MiniProgramApp.MiniDramaVOD.GetMediaList(ctx, &request.MediaListRequest{})

	if err != nil {
		panic(err)
	}

	fmt.Dump(result)

}

func APIGetMediaInfo(ctx *gin.Context) {
	MediaID := int64(111)

	result, err := services.MiniProgramApp.MiniDramaVOD.GetMediaInfo(ctx, MediaID)

	if err != nil {
		panic(err)
	}

	fmt.Dump(result)

}

// 需要提审后才能调用该接口
func APIGetMediaLink(ctx *gin.Context) {

	in := &request.GetMediaLinkRequest{
		MediaId: 1111,
		T:       time.Now().Unix() + 60*60*2, // 最大2小时过期
		Expr:    10,                          // 试看时常
	}

	result, err := services.MiniProgramApp.MiniDramaVOD.GetMediaLink(ctx, in)

	if err != nil {
		panic(err)
	}

	fmt.Dump(result)
}

func APIDeleteMedia(ctx *gin.Context) {

	MediaIds := []int64{111, 222, 333}

	for _, MediaID := range MediaIds {

		result, err := services.MiniProgramApp.MiniDramaVOD.DeleteMedia(ctx, MediaID)

		if err != nil {
			panic(err)
		}

		fmt.Dump(result)
	}

}

func UploadMediaMp(ctx *gin.Context) {
	result, err := services.MiniProgramApp.CustomerServiceMessage.UploadTempMedia(ctx, "image", "xxx.png", nil)

	if err != nil {

		panic(err)
	}

	fmt.Dump(result)
}

func APISubmitAudit(ctx *gin.Context) {

	result, err := services.MiniProgramApp.MiniDramaVOD.SubmitAudit(ctx, &request.SubmitAuditRequest{
		Name:               "xxxx",
		MediaCount:         1,
		MediaIdList:        []int64{111, 222, 333, 444, 555},
		Producer:           "xxxxx",    // 作者
		CoverMaterialId:    "xxxxxxxx", // 封面临时素材id （三天的那个 这里要调用《小程序》的上传临时素材接口）
		RegistrationNumber: "xxxxxx",   // 备案号
	})

	if err != nil {
		panic(err)
	}

	fmt.Dump(result)

}

func APIListDramas(ctx *gin.Context) {
	result, err := services.MiniProgramApp.MiniDramaVOD.GetDramaList(ctx, &request.GetDramaListRequest{})

	if err != nil {
		panic(err)
	}

	fmt.Dump(result)

}

func APIUploadByFile(ctx *gin.Context) {
	result, err := services.MiniProgramApp.MiniDramaVOD.VideoMediaUploadByFile(ctx, nil)

	if err != nil {
		panic(err)
	}

	fmt.Dump(result)

}
