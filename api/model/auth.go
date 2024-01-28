package models

type SignUpRequest struct {
	Name       string `json:"name"`
	SecondName string `json:"second_name"`
	Email      string `json:"email" example:"example@email.com"`
	Password   string `json:"password"`
}

type LoginRequest struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token"`
}

type RefreshTokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
