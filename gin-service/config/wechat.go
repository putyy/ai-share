package config

type WechatConfig struct {
	AppID     string `env:"WECHAT_APP_ID"`
	AppSecret string `env:"WECHAT_APP_SECRET"`
}
