package entity

import (
	"net/url"
	"regexp"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/fikrimohammad/ficree-api/domain"
	"github.com/fikrimohammad/ficree-api/infrastructure/storage"
)

// NewAWSFile ......
func NewAWSFile(uri string) domain.File {
	var isURL bool = true
	var isStored bool = true
	storageSVC := storage.GetAWSInstance()

	url, err := url.Parse(uri)
	if err != nil {
		isURL = false
	}

	if url != nil && url.Scheme == "" {
		isURL = false
	}

	bucketURL := storageSVC.Bucket.URL
	if !strings.Contains(uri, bucketURL) {
		isStored = false
	}

	if isURL && !isStored {
		return &AWSUnStoredFile{URL: uri}
	}

	path := strings.Replace(uri, bucketURL, "", -1)
	path = regexp.MustCompile(`\A\/`).ReplaceAllString(path, "")
	path = regexp.MustCompile(`\?.*\z`).ReplaceAllString(path, "")

	return &AWSStoredFile{Key: path}
}

// AWSStoredFile ......
type AWSStoredFile struct {
	Key string
}

// PublicURL ......
func (f *AWSStoredFile) PublicURL() string {
	storageSVC := storage.GetAWSInstance()
	bucketURL := storageSVC.Bucket.URL
	publicURL, _ := url.Parse(bucketURL + "/" + f.Key)
	return publicURL.String()
}

// DownloadURL ......
func (f *AWSStoredFile) DownloadURL() (string, error) {
	storageSVC := storage.GetAWSInstance()

	sess, err := session.NewSession(storageSVC.Config)
	if err != nil {
		return "", err
	}

	s3SVC := s3.New(sess)
	req, _ := s3SVC.GetObjectRequest(
		&s3.GetObjectInput{
			Bucket: aws.String(storageSVC.Bucket.Name),
			Key:    aws.String(f.Key),
		},
	)

	presignedURL, err := req.Presign(15 * time.Minute)
	if err != nil {
		return "", err
	}

	return presignedURL, nil
}

// UploadURL .......
func (f *AWSStoredFile) UploadURL() (string, error) {
	storageSVC := storage.GetAWSInstance()

	sess, err := session.NewSession(storageSVC.Config)
	if err != nil {
		return "", err
	}

	s3SVC := s3.New(sess)
	req, _ := s3SVC.PutObjectRequest(
		&s3.PutObjectInput{
			Bucket: aws.String(storageSVC.Bucket.Name),
			Key:    aws.String(f.Key),
		},
	)

	presignedURL, err := req.Presign(15 * time.Minute)
	if err != nil {
		return "", err
	}

	return presignedURL, nil
}

func (f *AWSStoredFile) AsFileOutput() (*domain.FileOutput, error) {
	presignedUploadURL, err := f.UploadURL()
	if err != nil {
		return nil, err
	}

	output := &domain.FileOutput{
		PublicURL:          f.PublicURL(),
		PresignedUploadURL: presignedUploadURL,
	}
	return output, nil
}

// AWSUnStoredFile ....
type AWSUnStoredFile struct {
	URL string
}

// PublicURL ....
func (f *AWSUnStoredFile) PublicURL() string {
	return f.URL
}

// DownloadURL .....
func (f *AWSUnStoredFile) DownloadURL() (string, error) {
	return f.URL, nil
}

// UploadURL .....
func (f *AWSUnStoredFile) UploadURL() (string, error) {
	return "", nil
}

func (f *AWSUnStoredFile) AsFileOutput() (*domain.FileOutput, error) {
	presignedUploadURL, _ := f.UploadURL()
	output := &domain.FileOutput{
		PublicURL:          f.PublicURL(),
		PresignedUploadURL: presignedUploadURL,
	}
	return output, nil
}
