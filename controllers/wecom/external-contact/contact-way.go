package external_contact

import (
	"net/http"
	"power-wechat-tutorial/services"

	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/externalContact/contactWay/request"
	"github.com/gin-gonic/gin"
)

// 获取配置了客户联系功能的成员列表.
// https://work.weixin.qq.com/api/doc/90000/90135/92571
func APIExternalContactGetFollowUserList(c *gin.Context) {

	res, err := services.WeComContactApp.ExternalContact.GetFollowUsers(c.Request.Context())

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)

}

// 配置客户联系「联系我」方式.
// https://open.work.weixin.qq.com/api/doc/90000/90135/92572
func APIExternalContactAddContactWay(c *gin.Context) {

	options := &request.RequestAddContactWay{
		Type:          1,
		Scene:         1,
		Style:         1,
		Remark:        "渠道客户",
		SkipVerify:    true,
		State:         "teststate",
		User:          []string{"zhangsan", "lisi", "wangwu"},
		Party:         []int{2, 3},
		IsTemp:        true,
		ExpiresIn:     86400,
		ChatExpiresIn: 86400,
		UnionID:       "oxTWIuGaIt6gTKsQRLau2M0AAAA",
		//Conclusions: &power.HashMap{
		//  "text": power.HashMap{
		//    "content": "文本消息内容",
		//  },
		//  "image": power.HashMap{
		//    "media_id": "MEDIA_ID",
		//  },
		//  "link": power.HashMap{
		//    "title":  "消息标题",
		//    "picurl": "https://example.pic.com/path",
		//    "desc":   "消息描述",
		//    "url":    "https://example.link.com/path",
		//  },
		//  "miniprogram": power.HashMap{
		//    "title":        "消息标题",
		//    "pic_media_id": "MEDIA_ID",
		//    "appid":        "wx8bd80126147dfAAA",
		//    "page":         "/path/index.html",
		//  },
		//},
	}

	res, err := services.WeComContactApp.ExternalContactContactWay.Add(c.Request.Context(), options)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)
}

// 获取企业已配置的「联系我」方式
// https://work.weixin.qq.com/api/doc/90000/90135/92572
func APIExternalContactGetContactWay(c *gin.Context) {
	configID := c.DefaultQuery("configID", "42b34949e138eb6e027c123cba77fad7")

	res, err := services.WeComContactApp.ExternalContactContactWay.Get(c.Request.Context(), configID)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)
}

// 获取企业已配置的「联系我」列表
// https://work.weixin.qq.com/api/doc/90000/90135/92572
func APIExternalContactListContactWay(c *gin.Context) {
	options := &request.RequestListContactWay{
		StartTime: 1622476800,
		EndTime:   1625068800,
		Cursor:    "CURSOR",
		Limit:     1000,
	}
	res, err := services.WeComContactApp.ExternalContactContactWay.List(c.Request.Context(), options)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)
}

// 更新企业已配置的「联系我」列表
// https://work.weixin.qq.com/api/doc/90000/90135/92572
func APIExternalContactUpdateContactWay(c *gin.Context) {

	configID := c.DefaultQuery("configID", "42b34949e138eb6e027c123cba77fAAA")
	options := &request.RequestUpdateContactWay{
		ConfigID:      configID,
		Style:         1,
		Remark:        "渠道客户",
		SkipVerify:    true,
		State:         "teststate",
		User:          []string{"zhangsan", "lisi", "wangwu"},
		Party:         []int{2, 3},
		ExpiresIn:     86400,
		ChatExpiresIn: 86400,
		UnionID:       "oxTWIuGaIt6gTKsQRLau2M0AAAA",
		//RequestAddContactWay: request.RequestAddContactWay{
		//Conclusions: &power.HashMap{
		//  "text": power.HashMap{
		//    "content": "文本消息内容",
		//  },
		//  "image": power.HashMap{
		//    "media_id": "MEDIA_ID",
		//  },
		//  "link": power.HashMap{
		//    "title":  "消息标题",
		//    "picurl": "https://example.pic.com/path",
		//    "desc":   "消息描述",
		//    "url":    "https://example.link.com/path",
		//  },
		//  "miniprogram": power.HashMap{
		//    "title":        "消息标题",
		//    "pic_media_id": "MEDIA_ID",
		//    "appid":        "wx8bd80126147dfAAA",
		//    "page":         "/path/index.html",
		//  },
		//},
		//},
	}

	res, err := services.WeComContactApp.ExternalContactContactWay.Update(c.Request.Context(), options)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)

}

func APIExternalContactDelContactWay(c *gin.Context) {

	configID := c.DefaultQuery("configID", "42b34949e138eb6e027c123cba77fad7")

	res, err := services.WeComContactApp.ExternalContactContactWay.Delete(c.Request.Context(), configID)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)
}

func APIExternalContactCloseTempChat(c *gin.Context) {

	userID := c.DefaultQuery("userID", "matrix-x")
	externalUserID := c.DefaultQuery("externalUserID", "matrix-x")

	res, err := services.WeComContactApp.ExternalContactContactWay.CloseTempChat(c.Request.Context(), userID, externalUserID)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)
}
