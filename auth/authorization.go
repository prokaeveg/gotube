package auth

import (
	"context"
	"errors"
)

func AuthUser(ctx context.Context, repo UserRepository, request UserAuthRequest) (UserAuthResponse, error) {
	userID, err := repo.FindUserIdByCredentials(ctx, request.Login, request.Password)
	if err != nil {
		return UserAuthResponse{}, err
	}

	if userID == 0 {
		return UserAuthResponse{}, errors.New("пользователь не найден")
	}

	token, refreshToken := CreateTokenForUser(userID)

	return UserAuthResponse{ID: userID, Token: token, RefreshToken: refreshToken}, nil
}
