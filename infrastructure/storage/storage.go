package storage

import (
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/fikrimohammad/ficree-api/config"
)

// AWSBucket .....
type AWSBucket struct {
	Name string
	URL  string
}

// AWSStorage .....
type AWSStorage struct {
	Config *aws.Config
	Bucket *AWSBucket
}

var awsInstance *AWSStorage

// GetAWSInstance is a function to fetch storage configuration
func GetAWSInstance() *AWSStorage {
	if awsInstance == nil {
		InitAWSInstance()
	}

	return awsInstance
}

// InitAWSInstance is a function to initialize AWS S3 service
func InitAWSInstance() {
	awsCredentials := credentials.NewStaticCredentials(
		config.Get().AWSAccessKey,
		config.Get().AWSSecretKey,
		config.Get().AWSSessionToken,
	)

	awsConfig := &aws.Config{}
	awsConfig.Credentials = awsCredentials
	awsConfig.Region = aws.String(config.Get().AWSRegion)
	awsConfig.S3UseAccelerate = aws.Bool(config.Get().AWSTransferAcceleration)
	awsConfig.S3ForcePathStyle = aws.Bool(config.Get().AWSForcePathStyle)
	if config.Get().MinioEndpoint != "" {
		awsConfig.Endpoint = aws.String(config.Get().MinioEndpoint)
		awsConfig.S3ForcePathStyle = aws.Bool(true)
	}

	awsBucket := &AWSBucket{}
	awsBucket.Name = config.Get().AWSBucketName
	awsBucket.URL = buildBucketURL()

	awsInstance = &AWSStorage{}
	awsInstance.Config = awsConfig
	awsInstance.Bucket = awsBucket
}

func buildBucketURL() string {
	var bucketURL string

	if config.Get().AWSForcePathStyle {
		bucketURL = "https://s3.REGIONNAME.amazonaws.com/BUCKETNAME"
	} else {
		bucketURL = "https://BUCKETNAME.s3.REGIONNAME.amazonaws.com"
	}

	if config.Get().MinioEndpoint != "" {
		bucketURL = strings.Replace(
			bucketURL,
			"https://s3.REGIONNAME.amazonaws.com",
			config.Get().MinioEndpoint,
			-1,
		)
	}

	bucketURL = strings.Replace(bucketURL, "BUCKETNAME", config.Get().AWSBucketName, -1)
	bucketURL = strings.Replace(bucketURL, "REGIONNAME", config.Get().AWSRegion, -1)

	return bucketURL
}
