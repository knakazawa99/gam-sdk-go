package gam

import (
	"context"
	"fmt"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"golang.org/x/oauth2/jwt"
)

const scope = "https://www.googleapis.com/auth/dfp"

type AccessTokenInput struct {
	PrivateKeyPemFilePath *string
	PrivateKeyPemString   *string
	PrivateKeyID          string
	ServiceAccountID      string
}

func GetAccessToken(ctx context.Context, config *jwt.Config) (*oauth2.Token, error) {
	token, err := config.TokenSource(ctx).Token()
	if err != nil {
		return nil, fmt.Errorf("failed to get token: %w", err)
	}
	return token, nil
}

func GetConfig(ctx context.Context, input AccessTokenInput) (*jwt.Config, error) {
	var privateKeyPem []byte

	if input.PrivateKeyPemFilePath != nil {
		var err error
		privateKeyPem, err = os.ReadFile(*input.PrivateKeyPemFilePath)
		if err != nil {
			return nil, fmt.Errorf("failed to read private key file %w", err)
		}
	}
	if input.PrivateKeyPemString != nil {
		privateKeyPem = []byte(*input.PrivateKeyPemString)
	}

	return &jwt.Config{
		Email:        input.ServiceAccountID,
		PrivateKey:   privateKeyPem,
		PrivateKeyID: input.PrivateKeyID,
		TokenURL:     google.JWTTokenURL,
		Scopes:       []string{scope},
	}, nil
}
