package robot_chat

import (
	v1 "github.com/artisancloud/openai/api/v1"
	"github.com/gin-gonic/gin"
	"power-wechat-tutorial/services"
)

func APIChatRequest(c *gin.Context) {

	// 调用GPT3.5模型
	req := v1.CreateChatCompletionRequest{
		Model: "gpt-3.5-turbo",
		Messages: []v1.Message{
			{
				Role:    "user",
				Content: "Hello!",
			},
		},
	}
	result, err := services.RobotChatApp.Client.V1.Chat.CreateChatCompletion(&req)
	if err != nil {
		panic(err)
		return
	}

	c.JSON(200, result)
}
