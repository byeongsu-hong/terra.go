package service

import (
	"context"
	"fmt"
	"net/http"

	encoding "github.com/terra-money/core/app/params"

	"github.com/frostornge/terra-go/httpclient"

	cosmostypes "github.com/cosmos/cosmos-sdk/types"
	cosmosauth "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/pkg/errors"
)

//go:generate mockgen -destination ../../../test/mocks/terra/service/service_auth.go . AuthService
type AuthService interface {
	GetAccountInfo(ctx context.Context, addr cosmostypes.AccAddress) (cosmosauth.AccountI, error)
}

type authService struct {
	codec  encoding.EncodingConfig
	client httpclient.Client
}

func NewAuthService(client httpclient.Client) AuthService {
	return authService{codec: client.Codec(), client: client}
}

func (svc authService) GetAccountInfo(
	ctx context.Context,
	addr cosmostypes.AccAddress,
) (cosmosauth.AccountI, error) {
	var payload = httpclient.RequestPayload{
		Context: ctx,
		Method:  http.MethodGet,
		Path:    fmt.Sprintf("/auth/accounts/%s", addr.String()),
	}

	var body struct {
		Height string              `json:"height"`
		Result cosmosauth.AccountI `json:"result"`
	}
	if err := svc.client.RequestJSON(payload, &body); err != nil {
		return nil, errors.Wrap(err, "request json")
	}
	return body.Result, nil
}
