package user

import (
	"fmt"
	domainUser "yt-d-rep/github.com/event-sourcing-go/domain/user"
)

type (
	userRepository struct{}
)

func (r *userRepository) Create(u *domainUser.User) error {
	fmt.Println("User created: id=", u.ID)
	return nil
}
