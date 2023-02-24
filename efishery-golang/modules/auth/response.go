package auth

type LoginResponse struct {
	Token string `json:"token"`
}

type RegisterResponse struct {
	ID       int    `json:"id"`
	Phone    string `json:"phone"`
	Name     string `json:"name"`
	Role     string `json:"role"`
	Password string `json:"password"`
}
