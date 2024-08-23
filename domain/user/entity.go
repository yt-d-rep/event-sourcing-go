package user

type (
	User struct {
		ID int
	}
)

func NewUser(id int) *User {
	return &User{ID: id}
}
