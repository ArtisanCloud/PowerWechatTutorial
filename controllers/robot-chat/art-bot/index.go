package art_bot

import (
	"github.com/gin-gonic/gin"
)

func APISDTxt2Img(c *gin.Context) {

	////prompt := c.DefaultQuery("prompt", "a pretty girl")
	////services.ChatBotApp.ConversationManager.CreateConversation("123")
	//////answer, err := services.ChatBotApp.Client.GenerateAnswer(c.Request.Context(), "what's time?")
	////
	////answer, err := services.ArtBotApp.Client.Text2Image(c.Request.Context(), &request2.Text2Image{
	////	Prompt:          prompt,
	////	DoNotSendImages: true,
	////	BatchSize:       1,
	////	BatchCount:      1,
	////	Steps:           5,
	////})
	////
	////if err != nil {
	////	c.JSON(http.StatusBadRequest, err.Error())
	////}
	//
	//c.JSON(http.StatusOK, answer)
}

func APISDImg2Img(c *gin.Context) {
	//prompt := c.DefaultQuery("prompt", "what's time now?")
	//services.ChatBotApp.ConversationManager.CreateConversation("123")
	//fmt.Dump(services.ChatBotApp.Client.GetConfig())
	//answer, err := services.ArtBotApp.Client.Text2Image(c.Request.Context(), &request2.Text2Image{
	//	Prompt:          prompt,
	//	DoNotSendImages: true,
	//	BatchSize:       1,
	//	BatchCount:      1,
	//	Steps:           5,
	//})
	////answer, err := services.ChatBotApp.Client.SendMessage(c.Request.Context(), "what's time?", "")
	//
	//if err != nil {
	//	c.JSON(http.StatusBadRequest, err.Error())
	//}
	//
	//c.JSON(http.StatusOK, answer)
}
