package user

type UserCreateInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}