package robot_chat

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"power-wechat-tutorial/services"
)

func APIChatGPTRequest(c *gin.Context) {
	services.ChatBotApp.ConversationManager.CreateConversation("123")
	//answer, err := services.ChatBotApp.Client.GenerateAnswer(c.Request.Context(), "what's time?")

	answer, err := services.ChatBotApp.Client.SendMessage(c.Request.Context(), "what's time?", "")

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	c.JSON(http.StatusOK, answer)
}
