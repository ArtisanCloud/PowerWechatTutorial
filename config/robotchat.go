package config

type RobotChat struct {
	ArtBot  `env:"art_bot"`
	ChatBot `env:"chat_bot"`
}
type ArtBot struct {
	Channel         string `env:"Channel"`
	StableDiffusion `env:"stale_diffusion"`
	Queue           Queue `env:"queue"`
	Log             Log   `env:"log"`
}
type ChatBot struct {
	Channel string `env:"Channel"`
	ChatGPT `env:"chat_gpt"`
	//XFYun    XFYun    `env:"xfyun"`
	//THUDMGLM THUDMGLM `env:"thudm_glm"`
	Queue Queue `env:"queue"`
	Log   Log   `env:"log"`
}

type StableDiffusion struct {
	Token     string `env:"token"`
	BaseUrl   string `env:"base_url"`
	PrefixUri string `env:"prefix_uri"`
	Version   string `env:"version"`
	HttpDebug bool   `env:"http_debug"`
	ProxyURL  string `env:"proxy_url"`
}

type ChatGPT struct {
	OpenAPIKey   string `env:"open_api_key"`
	Organization string `env:"organization"`
	Model        string `env:"model"`
	HttpDebug    bool   `env:"http_debug"`
	BaseURL      string `env:"base_url"`
	APIType      string `env:"api_type"`
	APIVersion   string `env:"api_version"`
}

type Redis struct {
	Addr       string `env:"addr"`
	ClientName string `env:"client_name"`
	Username   string `env:"user_name"`
	Password   string `env:"password"`
	DB         int    `env:"db"`
	MaxRetries int    `env:"max_retries"`
}

type Queue struct {
	Driver    string `env:"driver"`
	NotifyUrl string `env:"notify_url"`
	Redis     Redis  `env:"redis"`
}

type Log struct {
	Driver    string `env:"driver"`
	Env       string `env:"env"`
	InfoLog   string `env:"info_log"`
	ErrorLog  string `env:"error_log"`
	HttpDebug bool   `env:"httpde_bug"`
}
