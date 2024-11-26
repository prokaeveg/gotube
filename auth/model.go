package auth

import "context"

type UserAuthRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}
type UserAuthResponse struct {
	ID           int    `json:"id"`
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

type UserRepository interface {
	FindUserIdByCredentials(ctx context.Context, username string, password string) (int, error)
}
