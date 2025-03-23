package schemas

type AuthRegisterRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	PassWord string `json:"password" binding:"required"`
}

type AuthCheckRegistrationRequest struct {
	Email string `json:"email" binding:"required"`
	Code  string `json:"code" binding:"required"`
}

type AuthLoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"code" binding:"required"`
}
