package usercommand

import (
	domainUser "yt-d-rep/github.com/event-sourcing-go/domain/user"
)

type (
	Usecase interface {
		Create() error
	}
	usecase struct {
		uRepo domainUser.UserRepository
	}
)

func (u *usecase) Create() error {
	newUser := domainUser.NewUser(1, "John Doe")
	err := u.uRepo.Create(newUser)
	if err != nil {
		return err
	}
	return nil
}
