package overseer

import (
	"context"

	"github.com/frostornge/terra-go/types"

	cosmostypes "github.com/cosmos/cosmos-sdk/types"
)

var _ Querier = (*overseer)(nil)

type Querier interface {
	GetConfig(ctx context.Context) (GetConfigResponse, error)
	GetEpochState(ctx context.Context) (GetEpochStateResponse, error)
	GetWhitelist(
		ctx context.Context,
		collateral *cosmostypes.AccAddress,
		startAfter *cosmostypes.AccAddress,
		limit *uint32,
	) (GetWhitelistResponse, error)
	GetCollaterals(
		ctx context.Context,
		borrower cosmostypes.AccAddress,
	) (GetCollateralsResponse, error)
	GetAllCollaterals(
		ctx context.Context,
		startAfter *cosmostypes.AccAddress,
		limit *uint32,
	) (GetAllCollateralsResponse, error)
	GetDistributionParams(ctx context.Context) (GetDistributionParamsResponse, error)
	GetBorrowLimit(
		ctx context.Context,
		borrower cosmostypes.AccAddress,
		blockTime *uint64,
	) (GetBorrowLimitResponse, error)
}

func (o overseer) GetConfig(ctx context.Context) (GetConfigResponse, error) {
	var query = types.Q{}

	var resp struct {
		Height cosmostypes.Uint  `json:"height"`
		Result GetConfigResponse `json:"result"`
	}

	if err := o.Query(ctx, types.Q{"config": query}, &resp); err != nil {
		return GetConfigResponse{}, err
	}
	resp.Result.Height = resp.Height
	return resp.Result, nil
}

func (o overseer) GetEpochState(ctx context.Context) (GetEpochStateResponse, error) {
	var query = types.Q{}

	var resp struct {
		Height cosmostypes.Uint      `json:"height"`
		Result GetEpochStateResponse `json:"result"`
	}

	if err := o.Query(ctx, types.Q{"epoch_state": query}, &resp); err != nil {
		return GetEpochStateResponse{}, err
	}
	resp.Result.Height = resp.Height
	return resp.Result, nil
}

func (o overseer) GetWhitelist(
	ctx context.Context,
	collateral *cosmostypes.AccAddress,
	startAfter *cosmostypes.AccAddress,
	limit *uint32,
) (GetWhitelistResponse, error) {
	var query = types.Q{}
	if collateral != nil {
		query["collateral_token"] = collateral.String()
	}
	if startAfter != nil {
		query["start_after"] = startAfter.String()
	}
	if limit != nil {
		query["limit"] = *limit
	}

	var resp struct {
		Height cosmostypes.Uint     `json:"height"`
		Result GetWhitelistResponse `json:"result"`
	}

	if err := o.Query(ctx, types.Q{"whitelist": query}, &resp); err != nil {
		return GetWhitelistResponse{}, err
	}
	resp.Result.Height = resp.Height
	return resp.Result, nil
}

func (o overseer) GetCollaterals(
	ctx context.Context,
	borrower cosmostypes.AccAddress,
) (GetCollateralsResponse, error) {
	var query = types.Q{"borrower": borrower.String()}

	var resp struct {
		Height cosmostypes.Uint       `json:"height"`
		Result GetCollateralsResponse `json:"result"`
	}

	if err := o.Query(ctx, types.Q{"collaterals": query}, &resp); err != nil {
		return GetCollateralsResponse{}, err
	}
	resp.Result.Height = resp.Height
	return resp.Result, nil
}

func (o overseer) GetAllCollaterals(
	ctx context.Context,
	startAfter *cosmostypes.AccAddress,
	limit *uint32,
) (GetAllCollateralsResponse, error) {
	var query = types.Q{}
	if startAfter != nil {
		query["start_after"] = startAfter.String()
	}
	if limit != nil {
		query["limit"] = *limit
	}

	var resp struct {
		Height cosmostypes.Uint          `json:"height"`
		Result GetAllCollateralsResponse `json:"result"`
	}

	if err := o.Query(ctx, types.Q{"all_collaterals": query}, &resp); err != nil {
		return GetAllCollateralsResponse{}, err
	}
	resp.Result.Height = resp.Height
	return resp.Result, nil
}

func (o overseer) GetDistributionParams(ctx context.Context) (GetDistributionParamsResponse, error) {
	var query = types.Q{}

	var resp struct {
		Height cosmostypes.Uint              `json:"height"`
		Result GetDistributionParamsResponse `json:"result"`
	}

	if err := o.Query(ctx, types.Q{"distribution_params": query}, &resp); err != nil {
		return GetDistributionParamsResponse{}, err
	}
	resp.Result.Height = resp.Height
	return resp.Result, nil
}

func (o overseer) GetBorrowLimit(ctx context.Context, borrower cosmostypes.AccAddress, blockTime *uint64) (GetBorrowLimitResponse, error) {
	var query = types.Q{"borrower": borrower.String()}
	if blockTime != nil {
		query["block_time"] = *blockTime
	}

	var resp struct {
		Height cosmostypes.Uint       `json:"height"`
		Result GetBorrowLimitResponse `json:"result"`
	}

	if err := o.Query(ctx, types.Q{"borrow_limit": query}, &resp); err != nil {
		return GetBorrowLimitResponse{}, err
	}
	resp.Result.Height = resp.Height
	return resp.Result, nil
}
