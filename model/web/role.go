package web

type RoleCreate struct {
	Name string `json:"name,omitempty" validate:"required,gte=4"`
}

type RoleResponse struct {
	ID        string `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	CreatedAt int64  `json:"created_at,omitempty"`
	UpdatedAt int64  `json:"updated_at,omitempty"`
}
