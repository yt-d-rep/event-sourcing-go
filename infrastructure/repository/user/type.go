package user

import (
	"time"
	domainUser "yt-d-rep/github.com/event-sourcing-go/domain/user"
)

type User struct {
	ID        string `dynamodbav:"ID"`
	Name      string `dynamodbav:"Name"`
	Event     string `dynamodbav:"Event"`
	CreatedAt string `dynamodbav:"CreatedAt"`
}

func convertUser(dUser *domainUser.User) *User {
	return &User{
		ID:        dUser.GetID(),
		Name:      dUser.GetName(),
		Event:     dUser.GetEventString(),
		CreatedAt: time.Now().Format(time.RFC3339),
	}
}
