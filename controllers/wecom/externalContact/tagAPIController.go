package externalContact
import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/power"
	"github.com/ArtisanCloud/power-wechat/src/work/externalContact/tag/request"
	"github.com/gin-gonic/gin"
	"net/http"
	"power-wechat-tutorial/services"
)

// 获取企业标签库
// https://work.weixin.qq.com/api/doc/90000/90135/92117
func APIExternalContactGetCorpTagList(c *gin.Context) {

	tagIDs := []string{
		"etXXXXXXXXXX",
		"etYYYYYYYYYY",
	}

	groupIDs := []string{
		"etZZZZZZZZZZZZZ",
		"etYYYYYYYYYYYYY",
	}

	res, err := services.WeComContactApp.ExternalContactTag.GetCorpTagList(tagIDs, groupIDs)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)
}

func APIExternalContactAddCorpTag(c *gin.Context) {

	options := &request.RequestTagAddCorpTag{
		GroupID:   c.DefaultQuery("groupID", "GROUP_ID"),
		GroupName: c.DefaultQuery("groupName", "GROUP_NAME"),
		Order:     1,
		Tag: []*power.HashMap{
			&power.HashMap{
				"name":  "TAG_NAME_1",
				"order": 1,
			},
			&power.HashMap{
				"name":  "TAG_NAME_2",
				"order": 2,
			},
		},
	}

	res, err := services.WeComContactApp.ExternalContactTag.AddCorpTag(options)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)
}

func APIExternalContactEditCorpTag(c *gin.Context) {

	options := &request.RequestTagEditCorpTag{
		ID:        c.DefaultQuery("tagID", "TAG_ID"),
		GroupName: c.DefaultQuery("groupName", "NEW_TAG_NAME"),
		Order:     1,
		AgentID:   1000014,
	}
	res, err := services.WeComContactApp.ExternalContactTag.EditCorpTag(options)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)
}

func APIExternalContactDelCorpTag(c *gin.Context) {

	options := &request.RequestTagDelCorpTag{
		TagID: []string{
			c.DefaultQuery("tagID1", "TAG_ID_1"),
			c.DefaultQuery("tagID2", "TAG_ID_2"),
		},
		GroupID: []string{
			c.DefaultQuery("groupID1", "GROUP_ID_1"),
			c.DefaultQuery("groupID2", "GROUP_ID_2"),
		},
		AgentID: 1000014,
	}

	res, err := services.WeComContactApp.ExternalContactTag.DelCorpTag(options)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)
}

func APIExternalContactGetStrategyTagList(c *gin.Context) {

	options := &request.RequestTagGetStrategyTagList{
		TagID: []string{
			c.DefaultQuery("tagID1", "etXXXXXXXXXX"),
			c.DefaultQuery("tagID2", "etYYYYYYYYYY"),
		},
		GroupID: []string{
			c.DefaultQuery("groupID1", "etZZZZZZZZZZZZZ"),
			c.DefaultQuery("groupID2", "etYYYYYYYYYYYYY"),
		},
		StrategyID: 1,
	}

	res, err := services.WeComContactApp.ExternalContactTag.GetStrategyTagList(options)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)
}

func APIExternalContactAddStrategyTag(c *gin.Context) {

	options := &request.RequestTagAddStrategyTag{
		StrategyID: 1,
		GroupID:    c.DefaultQuery("groupID", "GROUP_ID"),
		GroupName:  c.DefaultQuery("groupName", "GROUP_NAME"),
		Order:      1,
		Tag: []*power.HashMap{
			&power.HashMap{
				"name":  "TAG_NAME_1",
				"order": 1,
			},
			&power.HashMap{
				"name":  "TAG_NAME_2",
				"order": 2,
			},
		},
	}

	res, err := services.WeComContactApp.ExternalContactTag.AddStrategyTag(options)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)
}

func APIExternalContactEditStrategyTag(c *gin.Context) {

	options := &request.RequestTagEditStrategyTag{
		ID:    c.DefaultQuery("groupID", "GROUP_ID"),
		Name:  c.DefaultQuery("groupName", "GROUP_NAME"),
		Order: 1,
	}

	res, err := services.WeComContactApp.ExternalContactTag.EditStrategyTag(options)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)
}
