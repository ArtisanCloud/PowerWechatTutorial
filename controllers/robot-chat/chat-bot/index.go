package chat_bot

import (
	"github.com/ArtisanCloud/PowerLibs/v3/fmt"
	"github.com/ArtisanCloud/RobotChat/robots/kernel/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"power-wechat-tutorial/services"
)

func APIChatGPTSendMessageRequest(c *gin.Context) {

	prompt := c.DefaultQuery("prompt", "what's time now?")
	//services.ChatBotApp.ConversationManager.CreateConversation("123")
	//answer, err := services.ChatBotApp.Client.GenerateAnswer(c.Request.Context(), "what's time?")

	msg, err := services.ChatBotApp.CreateTextMessage(prompt)

	answer, err := services.ChatBotApp.Client.CreateChatCompletion(c.Request.Context(), msg, model.UserRole)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	c.JSON(http.StatusOK, answer)
}

func APIChatGPTGenerateAnswerRequest(c *gin.Context) {
	prompt := c.DefaultQuery("prompt", "what's time now?")
	//services.ChatBotApp.ConversationManager.CreateConversation("123")
	fmt.Dump(services.ChatBotApp.Client.GetConfig())

	msg, err := services.ChatBotApp.CreateTextMessage(prompt)

	answer, err := services.ChatBotApp.Client.CreateCompletion(c.Request.Context(), msg)

	//answer, err := services.ChatBotApp.Client.SendMessage(c.Request.Context(), "what's time?", "")

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	c.JSON(http.StatusOK, answer)
}
