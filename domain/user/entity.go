package user

import "github.com/google/uuid"

type (
	User struct {
		id    string
		name  userName
		event userEventEnum
	}
)

func NewUser(name userName) *User {
	id := uuid.New().String()
	return &User{id: id, name: name, event: UserCreated}
}

// User: Get Props Method
func (u *User) GetID() string {
	return u.id
}
func (u *User) GetName() string {
	return u.name.String()
}
func (u *User) GetEventString() string {
	return u.event.String()
}

// User: Set Props Method
func (u *User) Resign() {
	u.event = UserResigned
}
func (u *User) Return() {
	u.event = UserReturned
}
func (u *User) ChangeName(name userName) {
	u.name = name
	u.event = UserChangedName
}
