package user

type (
	UserRepository interface {
		Create(*User) error
	}
)
