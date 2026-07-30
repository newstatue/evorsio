package s3

import (
	"context"
	"fmt"
	"io"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	awss3 "github.com/aws/aws-sdk-go-v2/service/s3"
)

type Client struct {
	client *awss3.Client
	bucket string
}

func New(
	endpoint,
	accessKey,
	secretKey,
	bucket string,
) (*Client, error) {
	if strings.TrimSpace(endpoint) == "" {
		return nil, fmt.Errorf("s3 endpoint required")
	}
	if strings.TrimSpace(bucket) == "" {
		return nil, fmt.Errorf("s3 bucket required")
	}
	if strings.TrimSpace(accessKey) == "" {
		return nil, fmt.Errorf("s3 access key required")
	}
	if strings.TrimSpace(secretKey) == "" {
		return nil, fmt.Errorf("s3 secret key required")
	}

	cfg, err := awsconfig.LoadDefaultConfig(
		context.Background(),
		awsconfig.WithRegion("us-east-1"),
		awsconfig.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(accessKey, secretKey, ""),
		),
	)

	if err != nil {
		return nil, fmt.Errorf("unable to load aws config: %w", err)
	}

	client := awss3.NewFromConfig(cfg, func(options *awss3.Options) {
		options.BaseEndpoint = aws.String(endpoint)
		options.UsePathStyle = true
	})

	return &Client{
		client: client,
		bucket: bucket,
	}, nil
}

func (c *Client) Upload(ctx context.Context, key string, body io.Reader, contentType string) (string, error) {
	object, err := c.client.PutObject(ctx, &awss3.PutObjectInput{
		Bucket:      aws.String(c.bucket),
		Key:         aws.String(key),
		Body:        body,
		ContentType: aws.String(contentType),
	})

	if err != nil {
		return "", fmt.Errorf("unable to upload object: %w", err)
	}

	return aws.ToString(object.ETag), nil
}

func (c *Client) Download(ctx context.Context, key string) (io.ReadCloser, error) {
	object, err := c.client.GetObject(ctx, &awss3.GetObjectInput{
		Bucket: aws.String(c.bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return nil, fmt.Errorf("unable to download object: %w", err)
	}

	return object.Body, nil
}

func (c *Client) Delete(ctx context.Context, key string) error {
	_, err := c.client.DeleteObject(ctx, &awss3.DeleteObjectInput{
		Bucket: aws.String(c.bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return fmt.Errorf("unable to delete object: %w", err)
	}
	return nil
}
