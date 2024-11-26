package mock

import (
	"gotube/auth"
	"math/rand/v2"
)

func AuthUser(login string, password string) (auth.UserAuthResponse, error) {
	id := rand.IntN(10 << 16)
	token, refreshToken := auth.CreateTokenForUser(id)

	return auth.UserAuthResponse{
		ID:           id,
		Token:        token,
		RefreshToken: refreshToken,
	}, nil
}
