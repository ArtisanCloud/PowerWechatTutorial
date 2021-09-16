package miniprogram

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/power"
	"github.com/gin-gonic/gin"
	"net/http"
	"power-wechat-tutorial/services"
)

// 本接口提供基于小程序的站内搜商品图片搜索能力
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/search/search.imageSearch.html
func APISearchImageSearch(c *gin.Context) {

	options := []*power.HashMap{
		&power.HashMap{
			"title":   "image title1",
			"img_url": "",
			"price":   "123",
			"path":    "path",
		},
		&power.HashMap{
			"title":   "image title2",
			"img_url": "",
			"price":   "123",
			"path":    "path",
		},
	}

	rs, err := services.MiniprogramApp.Search.ImageSearch(options)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, rs)
}

// 小程序内部搜索API提供针对页面的查询能力
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/search/search.siteSearch.html
func APISearchSiteSearch(c *gin.Context) {

	rs, err := services.MiniprogramApp.Search.SiteSearch("test", "pages/index/index")

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, rs)

}

// 小程序开发者可以通过本接口提交小程序页面url及参数信息(不要推送webview页面)
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/search/search.submitPages.html
func APISearchSubmitPages(c *gin.Context) {
	options := []*power.HashMap{
		&power.HashMap{
			"path":  "pages/index/index",
			"query": "userName=wechat_user",
		},
		&power.HashMap{
			"path":  "pages/video/index",
			"query": "vid=123",
		},
	}

	rs, err := services.MiniprogramApp.Search.SubmitPages(options)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, rs)

}
