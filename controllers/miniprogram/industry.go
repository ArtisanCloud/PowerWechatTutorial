package miniprogram

import (
	"github.com/ArtisanCloud/PowerLibs/v3/fmt"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/miniProgram/industry/miniDrama/request"
	"github.com/gin-gonic/gin"
	"power-wechat-tutorial/services"
	"time"
)

func APIVideoMediaUploadByURL(ctx *gin.Context) {

	params := &request.VideoMediaUploadByURLRequest{
		MediaUrl:  "https://db.bbwhcmjx.cn/mp40609diyishenyi78.mp4",
		CoverUrl:  "https://qiniu.rongjuwh.cn/FmDzqVBJtV-RF7_2aA8m44Uyw8Vh",
		MediaName: "女富婆的第一神医78",
	}

	result, err := services.MiniProgramApp.MiniDramaVOD.VideoMediaUploadByURL(ctx, params)

	if err != nil {
		panic(err)
	}

	fmt.Dump(result)

	//  "task_id": 100003

}

func APISearchMediaByTaskId(ctx *gin.Context) {

	result, err := services.MiniProgramApp.MiniDramaVOD.SearchMediaByTaskId(ctx, 100005)

	if err != nil {
		panic(err)
	}

	fmt.Dump(result)

	// "media_id": 100004

}

func APIGetMediaList(ctx *gin.Context) {

	result, err := services.MiniProgramApp.MiniDramaVOD.GetMediaList(ctx, &request.MediaListRequest{})

	if err != nil {
		panic(err)
	}

	fmt.Dump(result)

	/**

	{
	        "errcode": 0,
	        "errmsg": "ok",
	        "media_info_list": [
	                {
	                        "media_id": 100004,
	                        "create_time": 1686634625,
	                        "expire_time": 1694410428,
	                        "drama_id": 0,
	                        "file_size": 9436765,
	                        "duration": 61,
	                        "name": "超级洗脚妹 - 第1集",
	                        "description": "",
	                        "cover_url": "https://1500022159.vod-qcloud.com/cbb03642vodsh1500022159/9a40a6493270835009675841781/cover.jpeg",
	                        "original_url": "https://1500022159.vod-qcloud.com/cbb03642vodsh1500022159/9a40a6493270835009675841781/f0.mp4?t=64881cbc\u0026us=64880eac0a2cdd5c4fcd6768\u0026sign=17cd41a58e5ad2c3b1d392a5e16953d6",
	                        "mp4_url": "",
	                        "hls_url": "",
	                        "audit_detail": {
	                                "status": 0,
	                                "create_time": 0,
	                                "audit_time": 0,
	                                "reason": "",
	                                "evidence_material_id_list": null
	                        }
	                }
	        ]
	}


	*/

}

func APIGetMediaInfo(ctx *gin.Context) {

	result, err := services.MiniProgramApp.MiniDramaVOD.GetMediaInfo(ctx, 100005)

	if err != nil {
		panic(err)
	}

	fmt.Dump(result)

	/**

	{
	        "errcode": 0,
	        "errmsg": "ok",
	        "media_info": {
	                "media_id": 100004,
	                "create_time": 1686634625,
	                "expire_time": 1694410428,
	                "drama_id": 0,
	                "file_size": 9436765,
	                "duration": 61,
	                "name": "超级洗脚妹 - 第1集",
	                "description": "",
	                "cover_url": "https://1500022159.vod-qcloud.com/cbb03642vodsh1500022159/9a40a6493270835009675841781/cover.jpeg",
	                "original_url": "https://1500022159.vod-qcloud.com/cbb03642vodsh1500022159/9a40a6493270835009675841781/f0.mp4?t=648824a8\u0026us=6488169889cca8d74cc63b3d\u0026sign=d4e5cd32b99b68c33e80dc9c0e8c0ae7",
	                "mp4_url": "",
	                "hls_url": ""
	        }
	}


	*/

}

// 需要提审后才能调用该接口
func APIGetMediaLink(ctx *gin.Context) {

	in := &request.GetMediaLinkRequest{
		MediaId: "100006",
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

	MediaIds := []int64{100004, 100006}

	for _, MediaID := range MediaIds {

		result, err := services.MiniProgramApp.MiniDramaVOD.DeleteMedia(ctx, MediaID)

		if err != nil {
			panic(err)
		}

		fmt.Dump(result)
	}

	/**

	{
	        "errcode": 0,
	        "errmsg": "ok"
	}

	*/

}

func APISubmitAudit(ctx *gin.Context) {

	result, err := services.MiniProgramApp.MiniDramaVOD.SubmitAudit(ctx, &request.SubmitAuditRequest{
		Name:                 "爱你是人间妄想",
		MediaCount:           3,
		MediaIdList:          []int64{},
		Producer:             "上海凡酷文化传媒有限公司",
		CoverMaterialId:      "", // 封面临时素材id （三天的那个
		AuthorizedMaterialId: "",
		RegistrationNumber:   "", // 备案号

	})

	if err != nil {
		panic(err)
	}

	fmt.Dump(result)

	/**

	{
	        "errcode": 0,
	        "errmsg": "ok",
	        "task_id": 100003
	}

	*/

}
