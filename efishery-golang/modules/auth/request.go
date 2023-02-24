package auth

type LoginRequest struct {
	Phone    string `form:"phone"`
	Password string `form:"password"`
}

type RegisterRequest struct {
	Phone    string `form:"phone"`
	Password string `form:"password"`
	Role     string `form:"role"`
	Name     string `form:"name"`
}
