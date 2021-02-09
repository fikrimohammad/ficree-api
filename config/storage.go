package config

import (
	"github.com/spf13/viper"
)

// Storage ...
type Storage struct {
	AWSAccessKey                  string `mapstructure:"aws_access_key"`
	AWSSecretKey                  string `mapstructure:"aws_secret_key"`
	AWSSessionToken               string `mapstructure:"aws_session_token"`
	AWSRegion                     string `mapstructure:"aws_region"`
	AWSBucketName                 string `mapstructure:"aws_bucket_name"`
	AWSPresignedURLExpirationTime int    `mapstructure:"aws_presigned_url_expiration_time"`
	AWSTransferAcceleration       bool   `mapstructure:"aws_transfer_acceleration"`
	AWSForcePathStyle             bool   `mapstructure:"aws_force_path_style"`
	MinioEndpoint                 string `mapstructure:"minio_endpoint"`
}

// LoadDefaultStorageConfig is a function to load default value for AWS S3 configs
func LoadDefaultStorageConfig(provider *viper.Viper) {
	provider.SetDefault("AWS_ACCESS_KEY", "admin")
	provider.SetDefault("AWS_SECRET_KEY", "Password01")
	provider.SetDefault("AWS_SESSION_TOKEN", "")
	provider.SetDefault("AWS_REGION", "ap-southeast-1")
	provider.SetDefault("AWS_BUCKET_NAME", "ficree-dev")
	provider.SetDefault("AWS_PRESIGNED_URL_EXPIRATION_TIME", 900)
	provider.SetDefault("AWS_TRANSFER_ACCELERATION", false)
	provider.SetDefault("AWS_FORCE_PATH_STYLE", false)
	provider.SetDefault("MINIO_ENDPOINT", "")
}
