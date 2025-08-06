package s3

import (
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/lyyyuna/storage-manager/pkg/config"
)

type Service struct {
	client *s3.S3
	bucket string
}

func NewService(cfg *config.Config) (*Service, error) {
	sess, err := session.NewSession(&aws.Config{
		Region:   aws.String(cfg.Qiniu.Region),
		Endpoint: aws.String(cfg.Qiniu.Endpoint),
		Credentials: credentials.NewStaticCredentials(
			cfg.Qiniu.Ak,
			cfg.Qiniu.Sk,
			"",
		),
		S3ForcePathStyle: aws.Bool(true),
	})
	if err != nil {
		return nil, err
	}

	return &Service{
		client: s3.New(sess),
		bucket: cfg.Qiniu.Bucket,
	}, nil
}

func (s *Service) GeneratePresignedURL(key string, expiration time.Duration) (string, error) {
	req, _ := s.client.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(key),
	})

	url, err := req.Presign(expiration)
	if err != nil {
		return "", err
	}

	return url, nil
}

func (s *Service) GenerateUploadURL(key string, expiration time.Duration) (string, error) {
	req, _ := s.client.PutObjectRequest(&s3.PutObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(key),
	})

	url, err := req.Presign(expiration)
	if err != nil {
		return "", err
	}

	return url, nil
}
