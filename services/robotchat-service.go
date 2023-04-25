package services

import (
	"github.com/artisancloud/openai"
	"github.com/bwmarrin/discordgo"
	"power-wechat-tutorial/config"
)

var RobotChatApp *RobotChat

type RobotChat struct {
	GPTClient      *openai.Client
	DiscordSession *discordgo.Session
}

func NewRobotChatService(conf *config.Configuration) (*RobotChat, error) {

	// 创建一个OpenAI API客户端
	config := &openai.ClientConfig{
		V1: openai.V1Config{
			OpenAPIKey:   conf.OpenAI.OpenAPIKey,   // 必填, 你的OpenAI API密钥
			Organization: conf.OpenAI.Organization, // 可选, 你的组织id
			HttpDebug:    conf.OpenAI.HttpDebug,    // 可选, 是否开启调试模式(使用默认log输出HTTP请求信息)
			ProxyURL:     conf.OpenAI.ProxyURL,     // 可选, 代理地址, 如果包含基本鉴权, 格式为http://username:password@xxxxxx:port
		},
	}
	client, err := openai.NewClient(config)
	if err != nil {
		panic(err)
	}

	// 创建一个Discord客户端
	discordSession, err := discordgo.New("Bot " + conf.Discord.Token)

	app := &RobotChat{
		GPTClient:      client,
		DiscordSession: discordSession,
	}

	return app, err
}
