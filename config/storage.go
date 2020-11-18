package config

import (
	"github.com/spf13/viper"
)

// LoadDefaultStorageConfig is a function to load default value for AWS S3 configs
func LoadDefaultStorageConfig(provider *viper.Viper) {
	provider.SetDefault("AWS_ACCESS_KEY", "admin")
	provider.SetDefault("AWS_SECRET_KEY", "Password01")
	provider.SetDefault("AWS_SESSION_TOKEN", "")
	provider.SetDefault("AWS_REGION", "ap-southeast-1")
	provider.SetDefault("AWS_BUCKET_NAME", "ficree-dev")
	provider.SetDefault("AWS_PRESIGNED_URL_EXPIRATION_TIME", 900)
	provider.SetDefault("AWS_TRANSFER_ACCELERATION", false)
	provider.SetDefault("MINIO_ENDPOINT", "")
}
