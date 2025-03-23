package schemas

type AuthSignUpRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	PassWord string `json:"password" binding:"required"`
}

type AuthVerifySignUpRequest struct {
	Email string `json:"email" binding:"required"`
	Code  string `json:"code" binding:"required"`
}

type AuthSignInRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"code" binding:"required"`
}
