package identity

type (
	UserID string
)

type User struct {
	id   UserID
	name string
}

func NewUser(id UserID, name string) *User {
	return &User{id: id, name: name}
}

func (u User) Id() UserID {
	return u.id
}

func (u User) Name() string {
	return u.name
}
