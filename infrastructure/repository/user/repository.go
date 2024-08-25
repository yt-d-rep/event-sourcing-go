package user

import (
	domainUser "yt-d-rep/github.com/event-sourcing-go/domain/user"
	"yt-d-rep/github.com/event-sourcing-go/infrastructure/store"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type (
	userRepository struct {
		writer store.Writer
	}
)

func (r *userRepository) Create(u *domainUser.User) error {
	user := convertUser(u)
	av, err := attributevalue.MarshalMap(user)
	if err != nil {
		return err
	}

	ctx := r.writer.GetContext()
	client := r.writer.GetClient()

	_, err = client.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String("user"),
		Item:      av,
	})
	if err != nil {
		return err
	}
	return nil
}
