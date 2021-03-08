package service

import (
	"context"
	"fmt"
	"net/http"

	"github.com/frostornge/terra-go/httpclient"

	"github.com/cosmos/cosmos-sdk/codec"
	cosmostypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/pkg/errors"
)

//go:generate mockgen -destination ../../../test/mocks/terra/service/service_bank.go . BankService
type BankService interface {
	GetBalance(ctx context.Context, acc cosmostypes.AccAddress) (GetBalanceResponse, error)
}

type bankService struct {
	codec  *codec.Codec
	client httpclient.Client
}

func NewBankService(client httpclient.Client) BankService {
	return bankService{codec: client.Codec(), client: client}
}

func (svc bankService) GetBalance(ctx context.Context, acc cosmostypes.AccAddress) (GetBalanceResponse, error) {
	var payload = httpclient.RequestPayload{
		Context: ctx,
		Method:  http.MethodGet,
		Path:    fmt.Sprintf("/bank/balances/%s", acc.String()),
	}

	var body struct {
		Height cosmostypes.Uint  `json:"height"`
		Result cosmostypes.Coins `json:"result"`
	}
	if err := svc.client.RequestJSON(payload, &body); err != nil {
		return GetBalanceResponse{}, errors.Wrap(err, "request json")
	}
	return GetBalanceResponse{
		Height:  body.Height.Uint64(),
		Balance: body.Result,
	}, nil
}
