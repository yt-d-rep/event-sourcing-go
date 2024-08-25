package usercommand

import (
	domainUser "yt-d-rep/github.com/event-sourcing-go/domain/user"
)

type (
	Usecase interface {
		Create(*createArgs) error
	}
	usecase struct {
		uRepo domainUser.UserRepository
	}
	createArgs struct {
		name string
	}
)

func (u *usecase) Create(args *createArgs) error {
	name, err := domainUser.NewUserName(args.name)
	if err != nil {
		return err
	}
	newUser := domainUser.NewUser(name)
	if err = u.uRepo.Create(newUser); err != nil {
		return err
	}
	return nil
}
