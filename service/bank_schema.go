package service

import cosmostypes "github.com/cosmos/cosmos-sdk/types"

type GetBalanceResponse struct {
	Height  uint64            `json:"height"`
	Balance cosmostypes.Coins `json:"balance"`
}
