package external_contact

import (
	"net/http"
	"power-wechat-tutorial/services"

	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/externalContact/tag/request"
	"github.com/gin-gonic/gin"
)

// 获取企业标签库
// https://work.weixin.qq.com/api/doc/90000/90135/92117
func APIExternalContactGetCorpTagList(c *gin.Context) {

	tagIDs := []string{}

	groupIDs := []string{}

	res, err := services.WeComContactApp.ExternalContactTag.GetCorpTagList(tagIDs, groupIDs)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)
}

func APIExternalContactAddCorpTag(c *gin.Context) {

	options := &request.RequestTagAddCorpTag{
		GroupID:   power.String(c.DefaultQuery("groupID", "GROUP_ID")),
		GroupName: c.DefaultQuery("groupName", "GROUP_NAME"),
		Order:     1,
		Tag: []request.RequestTagAddCorpTagFieldTag{
			{
				Name:  c.DefaultQuery("tagName", "tagName"),
				Order: 1,
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
		ID:      c.DefaultQuery("tagID", "TAG_ID"),
		Name:    c.DefaultQuery("groupName", "NEW_TAG_NAME"),
		Order:   1,
		AgentID: power.Int64(1000016),
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
			//c.DefaultQuery("tagID2", "TAG_ID_2"),
		},
		GroupID: []string{
			//c.DefaultQuery("groupID1", "GROUP_ID_1"),
			//c.DefaultQuery("groupID2", "GROUP_ID_2"),
		},
		AgentID: power.Int64(1000016),
	}

	res, err := services.WeComContactApp.ExternalContactTag.DelCorpTag(options)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)
}

// APIExternalContactGetStrategyTagList 获取指定规则组下的企业客户标签
// https://work.weixin.qq.com/api/doc/90000/90135/94882
func APIExternalContactGetStrategyTagList(c *gin.Context) {

	options := &request.RequestTagGetStrategyTagList{
		TagID: []string{
			//c.DefaultQuery("tagID1", "etXXXXXXXXXX"),
			//c.DefaultQuery("tagID2", "etYYYYYYYYYY"),
		},
		GroupID: []string{
			//c.DefaultQuery("groupID1", "etZZZZZZZZZZZZZ"),
			//c.DefaultQuery("groupID2", "etYYYYYYYYYYYYY"),
		},
		StrategyID: 1,
	}

	res, err := services.WeComContactApp.ExternalContactTag.GetStrategyTagList(options)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)
}

// APIExternalContactAddStrategyTag 为指定规则组创建企业客户标签
func APIExternalContactAddStrategyTag(c *gin.Context) {

	options := &request.RequestTagAddStrategyTag{
		StrategyID: 1,
		GroupID:    c.DefaultQuery("groupID", "GROUP_ID"),
		GroupName:  c.DefaultQuery("groupName", "GROUP_NAME"),
		Order:      1,
		Tag: []request.RequestTagAddStrategyTagFieldTag{
			{
				Name:  "TAG_NAME_1",
				Order: 1,
			},
			{
				Name:  "TAG_NAME_2",
				Order: 2,
			},
		},
	}

	res, err := services.WeComContactApp.ExternalContactTag.AddStrategyTag(options)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)
}

// APIExternalContactEditStrategyTag 编辑指定规则组下的企业客户标签
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

// APIExternalContactDelStrategyTag 删除指定规则组下的企业客户标签
func APIExternalContactDelStrategyTag(c *gin.Context) {

	options := &request.RequestTagDelStrategyTag{
		TagID: []string{
			"TAG_ID_1",
			"TAG_ID_2",
		},
		GroupID: []string{
			"GROUP_ID_1",
			"GROUP_ID_2",
		},
	}

	res, err := services.WeComContactApp.ExternalContactTag.DelStrategyTag(options)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)
}

// APIExternalContactMarkTag 编辑客户企业标签
func APIExternalContactMarkTag(c *gin.Context) {
	userID := c.Query("userID")
	externalUserID := c.Query("externalUserID")
	tagID := c.Query("tagID")

	options := &request.RequestTagMarkTag{
		UserID:         userID,
		ExternalUserID: externalUserID,
		AddTag: []string{
			tagID,
		},
		RemoveTag: []string{
			"etsdnEDAAA90dZT43kQpqLGDU8wjDhbw",
		},
	}

	res, err := services.WeComContactApp.ExternalContactTag.MarkTag(options)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)
}
