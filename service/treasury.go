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

//go:generate mockgen -destination ../../../test/mocks/terra/service/service_treasury.go . TreasuryService
type TreasuryService interface {
	CalculateTax(ctx context.Context, coin cosmostypes.Coin) (cosmostypes.Int, error)
	GetTaxRate(ctx context.Context) (GetTaxRateResponse, error)
	GetTaxCap(ctx context.Context, denom string) (GetTaxCapResponse, error)
}

type treasuryService struct {
	codec  *codec.Codec
	client httpclient.Client
}

func NewTreasuryService(client httpclient.Client) TreasuryService {
	return treasuryService{codec: client.Codec(), client: client}
}

func (svc treasuryService) CalculateTax(ctx context.Context, coin cosmostypes.Coin) (cosmostypes.Int, error) {
	taxRateResp, err := svc.GetTaxRate(ctx)
	if err != nil {
		return cosmostypes.Int{}, errors.Wrap(err, "fetch tax rate")
	}
	taxRate := taxRateResp.TaxRate

	taxCapResp, err := svc.GetTaxCap(ctx, coin.Denom)
	if err != nil {
		return cosmostypes.Int{}, errors.Wrapf(err, "fetch tax cap of %s", coin.Denom)
	}
	taxCap := taxCapResp.TaxCap

	tax := cosmostypes.NewDecFromInt(coin.Amount).Mul(taxRate).TruncateInt()
	if tax.GT(taxCap) {
		tax = taxCap
	}
	return tax, nil
}

func (svc treasuryService) GetTaxRate(ctx context.Context) (GetTaxRateResponse, error) {
	var payload = httpclient.RequestPayload{
		Context: ctx,
		Method:  http.MethodGet,
		Path:    "/treasury/tax_rate",
	}

	var body struct {
		Height cosmostypes.Uint `json:"height"`
		Result cosmostypes.Dec  `json:"result"`
	}
	if err := svc.client.RequestJSON(payload, &body); err != nil {
		return GetTaxRateResponse{}, errors.Wrap(err, "request json")
	}
	return GetTaxRateResponse{
		Height:  body.Height.Uint64(),
		TaxRate: body.Result,
	}, nil
}

func (svc treasuryService) GetTaxCap(ctx context.Context, denom string) (GetTaxCapResponse, error) {
	var payload = httpclient.RequestPayload{
		Context: ctx,
		Method:  http.MethodGet,
		Path:    fmt.Sprintf("/treasury/tax_cap/%s", denom),
	}

	var body struct {
		Height cosmostypes.Uint `json:"height"`
		Result cosmostypes.Int  `json:"result"`
	}
	if err := svc.client.RequestJSON(payload, &body); err != nil {
		return GetTaxCapResponse{}, errors.Wrap(err, "request json")
	}
	return GetTaxCapResponse{
		Height: body.Height.Uint64(),
		TaxCap: body.Result,
	}, nil
}
