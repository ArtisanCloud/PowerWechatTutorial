package routes

import (
	"github.com/gin-gonic/gin"
	art_bot "power-wechat-tutorial/controllers/robot-chat/art-bot"
	"power-wechat-tutorial/controllers/robot-chat/chat-bot"
)

func InitRobotChatAPIRoutes(r *gin.Engine) {

	apiRouterOpenAI := r.Group("/robot-chat")
	{
		// Handle the request route
		apiRouterOpenAI.GET("/chat-bot/openai/sendMessage", chat_bot.APIChatGPTSendMessageRequest)
		apiRouterOpenAI.GET("/chat-bot/openai/generateAnswer", chat_bot.APIChatGPTGenerateAnswerRequest)
		apiRouterOpenAI.GET("/art-bot/sd/txt2img", art_bot.APISDTxt2Img)
		apiRouterOpenAI.GET("/art-bot/sd/img2img", art_bot.APISDImg2Img)

	}

}
