package types

type User struct {
	Username string
	Password string
}

func (u *User) CleanUp() {
	u.Username = ""
	u.Password = ""
}
