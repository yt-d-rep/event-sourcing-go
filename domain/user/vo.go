package user

import "errors"

type (
	userName      string
	userEventEnum int
)

const (
	UserUnknown userEventEnum = iota
	UserCreated
	UserResigned
	UserReturned
	UserChangedName
)

func (u userEventEnum) String() string {
	switch u {
	case UserCreated:
		return "UserCreated"
	case UserResigned:
		return "UserResigned"
	case UserReturned:
		return "UserReturned"
	case UserChangedName:
		return "UserChangedName"
	default:
		return "UserUnknown"
	}
}

func NewUserName(name string) (userName, error) {
	u := userName(name)
	err := u.Validate()
	if err != nil {
		return "", err
	}
	return u, nil
}

func (u userName) String() string {
	return string(u)
}
func (u userName) Validate() error {
	if u == "" {
		return errors.New("user name is required and cannot be empty")
	}
	return nil
}
