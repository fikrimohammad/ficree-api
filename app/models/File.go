package models

import (
	"net/url"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/fikrimohammad/ficree-api/infrastructures/storage"
)

type IFile interface {
	BaseName() string
	Path() string
	Size() int
	PublicURL() string
	UploadURL() string
	DownloadURL() string
}

func NewFile(uri string) (IFile, error) {
	var isURL bool = true
	var isStored bool = true
	storageSVC := storage.Instance()

	domainPattern := regexp.MustCompile(`s3-([\w-]+).amazonaws.com`)
	uri = strings.TrimSpace(uri)
	uri = domainPattern.ReplaceAllString(uri, "s3.\\1.amazonaws.com")

	u, parseErr := url.Parse(uri)
	if parseErr != nil {
		return &UnStored{URL: uri}, parseErr
	}

	if u.Scheme == "" {
		isURL = false
	}

	bucketURL := storageSVC.Bucket.URL
	if !strings.HasPrefix(uri, bucketURL) {
		isStored = false
	}

	if isURL && !isStored {
		return &UnStored{URL: uri}, nil
	}

	path := strings.Replace(uri, bucketURL, "", -1)
	path = regexp.MustCompile(`\A\/`).ReplaceAllString(path, "")
	path = regexp.MustCompile(`\?.*\z`).ReplaceAllString(path, "")

	return &Stored{Key: path}, nil
}

type Stored struct {
	Key string
}

func (f *Stored) BaseName() string {
	return filepath.Base(f.Key)
}

func (f *Stored) Path() string {
	return f.Key
}

func (f *Stored) PublicURL() string {
	storageSVC := storage.Instance()
	bucketURL := storageSVC.Bucket.URL

	return bucketURL + "/" + f.Key
}

func (f *Stored) DownloadURL() string {
	storageSVC := storage.Instance()

	sess, sessErr := session.NewSession(storageSVC.Config)
	if sessErr != nil {
		panic(sessErr)
	}

	s3SVC := s3.New(sess)
	req, _ := s3SVC.GetObjectRequest(
		&s3.GetObjectInput{
			Bucket: aws.String(storageSVC.Bucket.Name),
			Key:    aws.String(f.Key),
		},
	)

	presignedURL, presignErr := req.Presign(15 * time.Minute)
	if presignErr != nil {
		panic(presignErr)
	}

	return presignedURL
}

func (f *Stored) UploadURL() string {
	storageSVC := storage.Instance()

	sess, sessErr := session.NewSession(storageSVC.Config)
	if sessErr != nil {
		panic(sessErr)
	}

	s3SVC := s3.New(sess)
	req, _ := s3SVC.PutObjectRequest(
		&s3.PutObjectInput{
			Bucket: aws.String(storageSVC.Bucket.Name),
			Key:    aws.String(f.Key),
		},
	)

	presignedURL, presignErr := req.Presign(15 * time.Minute)
	if presignErr != nil {
		panic(presignErr)
	}

	return presignedURL
}

func (f *Stored) Size() int {
	return 0
}

type UnStored struct {
	URL string
}

func (f *UnStored) BaseName() string {
	return filepath.Base(f.URL)
}

func (f *UnStored) PublicURL() string {
	return f.URL
}

func (f *UnStored) DownloadURL() string {
	return f.URL
}

func (f *UnStored) UploadURL() string {
	return f.URL
}

func (f *UnStored) Path() string {
	u, _ := url.Parse(f.URL)
	return u.Path
}

func (f *UnStored) Size() int {
	return 0
}
