package user

type (
	User struct {
		id    int
		name  string
		event userEventEnum
	}
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

func NewUser(id int, name string) *User {
	return &User{id: id, name: name, event: UserCreated}
}

// User: Get Props Method
func (u *User) GetID() int {
	return u.id
}
func (u *User) GetName() string {
	return u.name
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
func (u *User) ChangeName(name string) {
	u.name = name
	u.event = UserChangedName
}
