package robot_chat

import (
	v1 "github.com/artisancloud/openai/api/v1"
	"github.com/bwmarrin/discordgo"
	"github.com/gin-gonic/gin"
	"log"
	"power-wechat-tutorial/services"
)

func APIChatGPTRequest(c *gin.Context) {

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
	result, err := services.RobotChatApp.GPTClient.V1.Chat.CreateChatCompletion(&req)
	if err != nil {
		panic(err)
		return
	}

	c.JSON(200, result)
}

func APIChatMJRequest(c *gin.Context) {

	//if err := services.RobotChatApp.DiscordSession.Open(); err != nil {
	//	panic(err)
	//}
	//services.RobotChatApp.DiscordSession.UpdateStatusComplex(discordgo.UpdateStatusData{
	//	AFK:    false,
	//	Status: string(discordgo.StatusOnline),
	//})
	//services.RobotChatApp.DiscordSession.Identify.Presence.Status = string(discordgo.StatusOnline)

	// Create an new Application
	ap := &discordgo.Application{}
	ap.Name = "TestApp"
	ap.Description = "TestDesc"
	ap, err := services.RobotChatApp.DiscordSession.ApplicationCreate(ap)
	log.Printf("ApplicationCreate: err: %+v, app: %+v\n", err, ap)

	//// Get a specific Application by it's ID
	//ap, err = services.RobotChatApp.DiscordSession.Application(ap.ID)
	//log.Printf("Application: err: %+v, app: %+v\n", err, ap)
	//
	//apps, err := services.RobotChatApp.DiscordSession.Applications()
	//log.Printf("Applications: err: %+v, apps : %+v\n", err, apps)
	////for k, v := range apps {
	////	log.Printf("Applications: %d : %+v\n", k, v)
	////}

	// Wait here until CTRL-C or other term signal is received.
	//fmt.Println("Bot is now running. Press CTRL-C to exit.")
	//sc := make(chan os.Signal, 1)
	//signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM)
	//<-sc
	//
	//services.RobotChatApp.DiscordSession.Close()

	c.JSON(200, true)
}
