package user

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}

func (User) IsNode() {}

