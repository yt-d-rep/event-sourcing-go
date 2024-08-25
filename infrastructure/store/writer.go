package store

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type (
	Writer interface {
		GetContext() context.Context
		GetClient() *dynamodb.Client
	}
	writer struct {
		ctx    context.Context
		client *dynamodb.Client
	}
)

func newWriter() (*writer, error) {
	// NOTE: ローカル実行用に固定している
	var ctx = context.Background()
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("localhost"),
		config.WithEndpointResolver(aws.EndpointResolverFunc(
			func(service, region string) (aws.Endpoint, error) {
				return aws.Endpoint{URL: "http://dynamodb:8000", SigningRegion: "localhost"}, nil // The SigningRegion key was what's was missing! D'oh.
			})),
		config.WithCredentialsProvider(credentials.StaticCredentialsProvider{
			Value: aws.Credentials{
				AccessKeyID: "dummy", SecretAccessKey: "dummy", SessionToken: "dummy",
				Source: "ローカル実行",
			},
		}),
	)
	if err != nil {
		return nil, err
	}
	client := dynamodb.NewFromConfig(cfg)
	return &writer{
		ctx:    ctx,
		client: client,
	}, nil
}

func (w *writer) GetContext() context.Context {
	return w.ctx
}

func (w *writer) GetClient() *dynamodb.Client {
	return w.client
}
