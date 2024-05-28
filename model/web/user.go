package web

type UserCreate struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

type UserResponse struct {
	ID        string `json:"id,omitempty"`
	Username  string `json:"username,omitempty"`
	CreatedAt int64  `json:"created_at,omitempty"`
	UpdatedAt int64  `json:"updated_at,omitempty"`
}

type UserResponseLogin struct {
	User  *UserResponse `json:"user,omitempty"`
	Token string        `json:"token,omitempty"`
}
