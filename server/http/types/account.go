package types

type LoginRequest struct {
	Username string `json:"username" validate:"required,min=4,max=72"`
	Password string `json:"password" validate:"required,min=4,max=72"`
}

type AuthResponse struct {
	Token string `json:"token"`
}

type RegisterRequest struct {
	Username  string `json:"username" validate:"required,min=4,max=72"`
	Password  string `json:"password" validate:"required,min=4,max=72"`
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
}
