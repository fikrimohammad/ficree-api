package storage

import (
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/fikrimohammad/ficree-api/config"
)

type Config struct {
	AccessKey                  string `mapstructure:"aws_access_key"`
	SecretKey                  string `mapstructure:"aws_secret_key"`
	SessionToken               string `mapstructure:"aws_session_token"`
	Region                     string `mapstructure:"aws_region"`
	BucketName                 string `mapstructure:"aws_bucket_name"`
	PresignedURLExpirationTime int    `mapstructure:"aws_presigned_url_expiration_time"`
	TransferAcceleration       bool   `mapstructure:"aws_transfer_acceleration"`
	MinioEndpoint              string `mapstructure:"minio_endpoint"`
	ForcePathStyle             bool   `mapstructure:"force_path_style"`
}

type Bucket struct {
	Name string
	URL  string
}

type Storage struct {
	Config *aws.Config
	Bucket *Bucket
}

var storageInstance *Storage

// Instance is a function to fetch storage configuration
func Instance() *Storage {
	return storageInstance
}

// Init is a function to initialize storage service
func Init() {
	var storageConfig Config
	configErr := config.Instance().Unmarshal(&storageConfig)
	if configErr != nil {
		panic(configErr.Error())
	}

	awsCredentials := credentials.NewStaticCredentials(
		storageConfig.AccessKey,
		storageConfig.SecretKey,
		storageConfig.SessionToken,
	)

	awsConfig := &aws.Config{}
	awsConfig.Credentials = awsCredentials
	awsConfig.Region = aws.String(storageConfig.Region)
	awsConfig.S3UseAccelerate = aws.Bool(storageConfig.TransferAcceleration)
	if storageConfig.MinioEndpoint != "" {
		awsConfig.Endpoint = aws.String(storageConfig.MinioEndpoint)
		awsConfig.S3ForcePathStyle = aws.Bool(true)
	}

	storageBucket := &Bucket{}
	storageBucket.Name = storageConfig.BucketName
	storageBucket.URL = buildBucketURL(storageConfig)

	storageInstance = &Storage{}
	storageInstance.Config = awsConfig
	storageInstance.Bucket = storageBucket
}

func buildBucketURL(config Config) string {
	var bucketURL string

	if config.ForcePathStyle {
		bucketURL = "http://s3.REGIONNAME.amazonaws.com/BUCKETNAME"
	} else {
		bucketURL = "http://BUCKETNAME.s3.REGIONNAME.amazonaws.com"
	}

	if config.MinioEndpoint != "" {
		bucketURL = strings.Replace(
			bucketURL,
			"s3.REGIONNAME.amazonaws.com",
			config.MinioEndpoint,
			-1,
		)
	}

	bucketURL = strings.Replace(bucketURL, "BUCKETNAME", config.BucketName, -1)
	bucketURL = strings.Replace(bucketURL, "REGIONNAME", config.Region, -1)

	return bucketURL
}
