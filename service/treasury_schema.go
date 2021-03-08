package service

import cosmostypes "github.com/cosmos/cosmos-sdk/types"

type GetTaxRateResponse struct {
	Height  uint64          `json:"height"`
	TaxRate cosmostypes.Dec `json:"tax_rate"`
}

type GetTaxCapResponse struct {
	Height uint64          `json:"height"`
	TaxCap cosmostypes.Int `json:"tax_cap"`
}
