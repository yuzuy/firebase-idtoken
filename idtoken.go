package idtoken

import (
	"context"

	"firebase.google.com/go/auth"
	"google.golang.org/api/identitytoolkit/v3"
	"google.golang.org/api/option"
)

// Generate generates id token. opt needs a client option to call googleapis
func Generate(ctx context.Context, client *auth.Client, uid string, opt ...option.ClientOption) (string, error) {
	cToken, err := client.CustomToken(ctx, uid)
	if err != nil {
		return "", err
	}
	itkService, err := identitytoolkit.NewService(ctx, opt...)
	if err != nil {
		return "", err
	}

	resp, err := itkService.Relyingparty.VerifyCustomToken(
		&identitytoolkit.IdentitytoolkitRelyingpartyVerifyCustomTokenRequest{Token: cToken},
	).Do()
	if err != nil {
		return "", err
	}

	return resp.IdToken, nil
}
