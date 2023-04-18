package routes

import (
	"github.com/gin-gonic/gin"
	robot_chat "power-wechat-tutorial/controllers/robot-chat"
)

func InitRobotChatAPIRoutes(r *gin.Engine) {

	apiRouterOpenAI := r.Group("/robot-chat")
	{
		// Handle the request route
		apiRouterOpenAI.GET("/openai/chat", robot_chat.APIChatRequest)

	}

}
