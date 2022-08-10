package config

type QiniuConfig struct {
	AccessKey   string `env:"QN_ACCESS_KEY"`
	SecretKey   string `env:"QN_SECRET_KEY"`
	ImageBucket string `env:"QN_IMAGE_BUCKET"`
	ImageDomain string `env:"QN_IMAGE_DOMAIN"`
	MediaBucket string `env:"QN_MEDIA_BUCKET"`
	MediaDomain string `env:"QN_MEDIA_DOMAIN"`
}
