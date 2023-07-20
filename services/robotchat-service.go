package services

import (
	"github.com/ArtisanCloud/RobotChat/rcconfig"
	"github.com/ArtisanCloud/RobotChat/robots/artBot"
	config3 "github.com/ArtisanCloud/RobotChat/robots/artBot/config"
	"github.com/ArtisanCloud/RobotChat/robots/artBot/driver/Meonako"
	"github.com/ArtisanCloud/RobotChat/robots/chatBot"
	config2 "github.com/ArtisanCloud/RobotChat/robots/chatBot/config"
	go_openai "github.com/ArtisanCloud/RobotChat/robots/chatBot/driver/go-openai"
	"power-wechat-tutorial/config"
)

var ChatBotApp *chatBot.ChatBot
var ArtBotApp *artBot.ArtBot

//type RobotChat struct {
//	GPTClient      *Client
//	DiscordSession *discordgo.Session
//}

func NewChatBotService(conf config.ChatBot) (*chatBot.ChatBot, error) {

	driver := go_openai.NewDriver(&rcconfig.ChatBot{
		ChatGPT: config2.ChatGPT{
			OpenAPIKey:   conf.ChatGPT.OpenAPIKey,
			Model:        conf.ChatGPT.Model,
			Organization: conf.ChatGPT.Organization,
			HttpDebug:    conf.ChatGPT.HttpDebug,
			BaseUrl:      conf.ChatGPT.BaseURL,
			APIType:      conf.ChatGPT.APIType,
			APIVersion:   conf.ChatGPT.APIVersion,
		},
	})

	return chatBot.NewChatBot(driver)
}

func NewArtBotService(conf config.ArtBot) (*artBot.ArtBot, error) {

	driver := Meonako.NewDriver(&rcconfig.ArtBot{
		StableDiffusion: config3.StableDiffusion{
			HttpDebug: conf.HttpDebug,
			BaseUrl:   conf.BaseUrl,
		},
	})

	return artBot.NewArtBot(driver)

}

//func NewRobotChatService(conf *config.Configuration) (*RobotChat, error) {
//
//	//RCConfi
//
//	driver := NewDrive(&ChatRobot)
//
//	robot, err := NewChatRobot(driver)
//	if err != nil {
//		t.Error(err)
//	}
//
//}

//func NewRobotChatService(conf *config.Configuration) (*RobotChat, error) {
//
//	// 创建一个OpenAI API客户端
//	config := &openai.ClientConfig{
//		V1: openai.V1Config{
//			OpenAPIKey:   conf.OpenAI.OpenAPIKey,   // 必填, 你的OpenAI API密钥
//			Organization: conf.OpenAI.Organization, // 可选, 你的组织id
//			HttpDebug:    conf.OpenAI.HttpDebug,    // 可选, 是否开启调试模式(使用默认log输出HTTP请求信息)
//			ProxyURL:     conf.OpenAI.ProxyURL,     // 可选, 代理地址, 如果包含基本鉴权, 格式为http://username:password@xxxxxx:port
//		},
//	}
//	client, err := openai.NewClient(config)
//	if err != nil {
//		panic(err)
//	}
//
//	// 创建一个Discord客户端
//	discordSession, err := discordgo.New("Bot " + conf.Discord.Token)
//
//	app := &RobotChat{
//		GPTClient:      client,
//		DiscordSession: discordSession,
//	}
//
//	return app, err
//}
