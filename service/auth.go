package service

import (
	"context"
	"fmt"
	"net/http"

	"github.com/frostornge/terra-go/httpclient"

	"github.com/cosmos/cosmos-sdk/codec"
	cosmostypes "github.com/cosmos/cosmos-sdk/types"
	cosmosauth "github.com/cosmos/cosmos-sdk/x/auth/exported"
	"github.com/pkg/errors"
)

//go:generate mockgen -destination ../../../test/mocks/terra/service/service_auth.go . AuthService
type AuthService interface {
	GetAccountInfo(ctx context.Context, addr cosmostypes.AccAddress) (cosmosauth.Account, error)
}

type authService struct {
	codec  *codec.Codec
	client httpclient.Client
}

func NewAuthService(client httpclient.Client) AuthService {
	return authService{codec: client.Codec(), client: client}
}

func (svc authService) GetAccountInfo(
	ctx context.Context,
	addr cosmostypes.AccAddress,
) (cosmosauth.Account, error) {
	var payload = httpclient.RequestPayload{
		Context: ctx,
		Method:  http.MethodGet,
		Path:    fmt.Sprintf("/auth/accounts/%s", addr.String()),
	}

	var body struct {
		Height string             `json:"height"`
		Result cosmosauth.Account `json:"result"`
	}
	if err := svc.client.RequestJSON(payload, &body); err != nil {
		return nil, errors.Wrap(err, "request json")
	}
	return body.Result, nil
}
