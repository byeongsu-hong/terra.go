package cw20

import (
	"context"

	"github.com/frostornge/terra-go/types"

	cosmostypes "github.com/cosmos/cosmos-sdk/types"
)

var _ Querier = (*token)(nil)

type Querier interface {
	GetBalance(ctx context.Context, owner cosmostypes.AccAddress) (GetTokenBalanceResponse, error)
	GetTokenInfo(ctx context.Context) (GetTokenInfoResponse, error)
	GetMinter(ctx context.Context) (GetMinterResponse, error)
	GetAllowance(ctx context.Context, owner cosmostypes.AccAddress, spender cosmostypes.AccAddress) (GetAllowanceResponse, error)
	GetAllAllowances(ctx context.Context, owner cosmostypes.AccAddress, startAfter *cosmostypes.AccAddress, limit *uint32) ([]GetAllowanceResponse, error)
	GetAllAccounts(ctx context.Context, startAfter *cosmostypes.AccAddress, limit *uint32) ([]cosmostypes.AccAddress, error)
}

func (t token) GetBalance(ctx context.Context, owner cosmostypes.AccAddress) (GetTokenBalanceResponse, error) {
	var query = types.Q{"address": owner.String()}

	var resp struct {
		Height cosmostypes.Uint `json:"height"`
		Result struct {
			Balance cosmostypes.Int `json:"balance"`
		} `json:"result"`
	}

	if err := t.Query(ctx, types.Q{"balance": query}, &resp); err != nil {
		return GetTokenBalanceResponse{}, err
	}

	return GetTokenBalanceResponse{
		Height:  resp.Height.Uint64(),
		Balance: resp.Result.Balance,
	}, nil
}

func (t token) GetTokenInfo(ctx context.Context) (GetTokenInfoResponse, error) {
	var query = types.Q{}

	var resp struct {
		Height cosmostypes.Uint     `json:"height"`
		Result GetTokenInfoResponse `json:"result"`
	}

	if err := t.Query(ctx, types.Q{"token_info": query}, &resp); err != nil {
		return GetTokenInfoResponse{}, err
	}
	return resp.Result, nil
}

func (t token) GetMinter(ctx context.Context) (GetMinterResponse, error) {
	var query = types.Q{}

	var resp struct {
		Height cosmostypes.Uint  `json:"height"`
		Result GetMinterResponse `json:"result"`
	}

	if err := t.Query(ctx, types.Q{"minter": query}, &resp); err != nil {
		return GetMinterResponse{}, err
	}
	return resp.Result, nil
}

func (t token) GetAllowance(
	ctx context.Context,
	owner cosmostypes.AccAddress,
	spender cosmostypes.AccAddress,
) (GetAllowanceResponse, error) {
	var query = types.Q{
		"owner":   owner.String(),
		"spender": spender.String(),
	}

	var resp struct {
		Height cosmostypes.Uint     `json:"height"`
		Result GetAllowanceResponse `json:"result"`
	}

	if err := t.Query(ctx, types.Q{"allowance": query}, &resp); err != nil {
		return GetAllowanceResponse{}, err
	}
	return resp.Result, nil
}

func (t token) GetAllAllowances(
	ctx context.Context,
	owner cosmostypes.AccAddress,
	startAfter *cosmostypes.AccAddress,
	limit *uint32,
) ([]GetAllowanceResponse, error) {
	var query = types.Q{"owner": owner.String()}
	if startAfter != nil {
		query["start_after"] = (*startAfter).String()
	}
	if limit != nil {
		query["limit"] = *limit
	}

	var resp struct {
		Height cosmostypes.Uint `json:"height"`
		Result struct {
			Allowances []GetAllowanceResponse `json:"allowances"`
		} `json:"result"`
	}

	if err := t.Query(ctx, types.Q{"all_allowances": query}, &resp); err != nil {
		return nil, err
	}
	return resp.Result.Allowances, nil
}

func (t token) GetAllAccounts(
	ctx context.Context,
	startAfter *cosmostypes.AccAddress,
	limit *uint32,
) ([]cosmostypes.AccAddress, error) {
	var query = types.Q{}
	if startAfter != nil {
		query["start_after"] = (*startAfter).String()
	}
	if limit != nil {
		query["limit"] = *limit
	}

	var resp struct {
		Height string `json:"height"`
		Result struct {
			Accounts []cosmostypes.AccAddress `json:"accounts"`
		} `json:"result"`
	}

	if err := t.Query(ctx, types.Q{"all_accounts": query}, &resp); err != nil {
		return nil, err
	}
	return resp.Result.Accounts, nil
}
