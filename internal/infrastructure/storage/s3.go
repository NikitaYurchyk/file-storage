package storage

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"

	"file_storage/internal/domain"
)

type S3Storage struct {
    client *s3.Client
    bucket string
    region string
}

func NewS3Storage(bucket, region string) (domain.FileStorage, error) {
    cfg, err := config.LoadDefaultConfig(context.Background(),
        config.WithRegion(region),
    )
    if err != nil {
        return nil, fmt.Errorf("failed to load AWS config: %w", err)
    }

    return &S3Storage{
        client: s3.NewFromConfig(cfg),
        bucket: bucket,
        region: region,
    }, nil
}

func (s *S3Storage) Upload(ctx context.Context, key string, reader io.Reader, contentType string) (string, error) {
    _, err := s.client.PutObject(ctx, &s3.PutObjectInput{
        Bucket:      aws.String(s.bucket),
        Key:         aws.String(key),
        Body:        reader,
        ContentType: aws.String(contentType),
    })
    if err != nil {
        return "", fmt.Errorf("failed to upload to S3: %w", err)
    }

    url := fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", s.bucket, s.region, key)
    return url, nil
}

func (s *S3Storage) Download(ctx context.Context, key string) (io.ReadCloser, error) {
    result, err := s.client.GetObject(ctx, &s3.GetObjectInput{
        Bucket: aws.String(s.bucket),
        Key:    aws.String(key),
    })
    if err != nil {
        return nil, fmt.Errorf("failed to download from S3: %w", err)
    }
    return result.Body, nil
}

func (s *S3Storage) Delete(ctx context.Context, key string) error {
    _, err := s.client.DeleteObject(ctx, &s3.DeleteObjectInput{
        Bucket: aws.String(s.bucket),
        Key:    aws.String(key),
    })
    return err
}

func (s *S3Storage) GetSignedURL(ctx context.Context, key string, filename string, expiry time.Duration) (string, error) {
    presignClient := s3.NewPresignClient(s.client)

    disposition := fmt.Sprintf("attachment; filename=\"%s\"", filename)
    result, err := presignClient.PresignGetObject(ctx, &s3.GetObjectInput{
        Bucket:                     aws.String(s.bucket),
        Key:                        aws.String(key),
        ResponseContentDisposition: aws.String(disposition),
    }, s3.WithPresignExpires(expiry))
    if err != nil {
        return "", err
    }
    return result.URL, nil
}