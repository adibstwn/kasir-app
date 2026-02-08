package dto

type LoginUserReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateUserReq struct {
	Name     string `json:"name"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password"`
}
