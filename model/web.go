package model

type UserCreate struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

type UserResponse struct {
	ID       string `json:"id,omitempty"`
	Username string `json:"username,omitempty"`
}

type UserResponseLogin struct {
	User  *UserResponse `json:"user,omitempty"`
	Token string        `json:"token,omitempty"`
}

type WebResponse struct {
	Code   int
	Status string
	Data   any
}
