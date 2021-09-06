package miniprogram

import (
	"github.com/gin-gonic/gin"
	"power-wechat-tutorial/services"
)

func APISNSSession(c *gin.Context)  {

	rs , err:=services.AppMiniProgram.Auth.Session("123")

	if err != nil {
		panic(err)
	}

	c.JSON(200, rs)

}


func APICheckEncryptedData(c *gin.Context)  {

	rs , err:=services.AppMiniProgram.Base.CheckEncryptedData("hsSuSUsePBqSQw2rYMtf9Nvha603xX8f2BMQBcYRoJiMNwOqt/UEhrqekebG5ar0LFNAm5MD4Uz6zorRwiXJwbySJ/FEJHav4NsobBIU1PwdjbJWVQLFy7+YFkHB32OnQXWMh6ugW7Dyk2KS5BXp1f5lniKPp1KNLyNLlFlNZ2mgJCJmWvHj5AI7BLpWwoRvqRyZvVXo+9FsWqvBdxmAPA==")

	if err != nil {
		panic(err)
	}

	c.JSON(200, rs)

}

func APIGetPaidUnionID(c *gin.Context)  {

	rs , err:=services.AppMiniProgram.Base.GetPaidUnionID("", nil)

	if err != nil {
		panic(err)
	}

	c.JSON(200, rs)

}

