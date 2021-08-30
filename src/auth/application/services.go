package application

import (
	"context"
	"fmt"
	"sagara-test/src/auth/infrastructure/helper"
)

func loginService(ctx context.Context, request VMAuthRequest) (response VMAuthResponse, err error) {
	if request.Username != "username" && request.Password != "password" {
		err = fmt.Errorf("Invalid username and password")
		return
	}

	token, expireAt, err := helper.GenerateToken("", 1, 1)
	if err != nil {
		return
	}

	response = VMAuthResponse{
		AccessToken:  token["access_token"],
		RefreshToken: token["refresh_token"],
		ExpiresIn:    expireAt,
	}

	return
}
