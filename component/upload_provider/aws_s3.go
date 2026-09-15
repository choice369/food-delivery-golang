package upload_provider

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"net/http"

	"food_delivery/common"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type s3Provider struct {
	bucketName string
	region     string
	apiKey     string
	secret     string
	domain     string
	session    *session.Session
}

func NewS3Provider(bucketName string, region string, apiKey string, domain string, secret string) *s3Provider {
	provider := &s3Provider{
		bucketName: bucketName,
		region:     region,
		apiKey:     apiKey,
		secret:     secret,
		domain:     domain,
	}

	s3Session, err := session.NewSession(&aws.Config{
		Region:      aws.String(provider.region),
		Credentials: credentials.NewStaticCredentials(provider.apiKey, provider.secret, ""),
	})
	if err != nil {
		log.Fatal(err)
	}

	provider.session = s3Session

	return provider
}

func (p *s3Provider) SaveFileUploaded(ctx context.Context, data []byte, dst string) (*common.Image, error) {
	fileBytes := bytes.NewReader(data)
	fileType := http.DetectContentType(data)

	_, err := s3.New(p.session).PutObject(&s3.PutObjectInput{
		Bucket:      aws.String(p.bucketName),
		Key:         aws.String(dst),
		ACL:         aws.String("private"),
		ContentType: aws.String(fileType),
		Body:        fileBytes,
	})
	//req, _ := s3.New(p.session).PutObjectRequest(&s3.PutObjectInput{
	//	Bucket: aws.String(p.bucketName),
	//	Key:    aws.String(dst),
	//	ACL:    aws.String("private"),
	//})
	//
	//req.Presign(time.Second * 5) // return URL
	if err != nil {
		return nil, err
	}

	img := &common.Image{
		Url:       fmt.Sprintf("%d%s", p.domain, dst),
		CloudName: "s3",
	}

	return img, nil
}
