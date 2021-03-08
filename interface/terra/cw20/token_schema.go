package cw20

import cosmostypes "github.com/cosmos/cosmos-sdk/types"

// ================= Query ================= //

type GetTokenBalanceResponse struct {
	Height  uint64          `json:"height"`
	Balance cosmostypes.Int `json:"balance"`
}

type GetTokenInfoResponse struct {
	Decimals    int8            `json:"decimals"`
	Name        string          `json:"name"`
	Symbol      string          `json:"symbol"`
	TotalSupply cosmostypes.Int `json:"total_supply"`
}

type GetMinterResponse struct {
	Cap    *cosmostypes.Int       `json:"cap"`
	Minter cosmostypes.AccAddress `json:"minter"`
}

type GetAllowanceResponse struct {
	Allowance cosmostypes.Int        `json:"allowance"`
	Expires   map[string]interface{} `json:"expires"`
}
