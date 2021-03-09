package market

import (
	"context"

	"github.com/frostornge/terra-go/types"

	cosmostypes "github.com/cosmos/cosmos-sdk/types"
)

var _ Querier = (*market)(nil)

type Querier interface {
	GetConfig(ctx context.Context) (GetConfigResponse, error)
	GetState(ctx context.Context) (GetStateResponse, error)
	GetEpochState(ctx context.Context, height *uint64) (GetEpochStateResponse, error)
	GetBorrowerInfo(ctx context.Context, borrower cosmostypes.AccAddress, height *uint64) (GetBorrowerInfoResponse, error)
	GetBorrowerInfos(ctx context.Context, startAfter *cosmostypes.AccAddress, limit *uint32) ([]GetBorrowerInfoResponse, error)
}

func (m market) GetConfig(ctx context.Context) (GetConfigResponse, error) {
	var query = types.Q{}

	var resp struct {
		Height cosmostypes.Uint  `json:"height"`
		Result GetConfigResponse `json:"result"`
	}

	if err := m.Query(ctx, types.Q{"config": query}, &resp); err != nil {
		return GetConfigResponse{}, err
	}
	return resp.Result, nil
}

func (m market) GetState(ctx context.Context) (GetStateResponse, error) {
	var query = types.Q{}

	var resp struct {
		Height cosmostypes.Uint `json:"height"`
		Result GetStateResponse `json:"result"`
	}

	if err := m.Query(ctx, types.Q{"state": query}, &resp); err != nil {
		return GetStateResponse{}, err
	}
	return resp.Result, nil
}

func (m market) GetEpochState(ctx context.Context, height *uint64) (GetEpochStateResponse, error) {
	var query = types.Q{}
	if height != nil {
		query["block_height"] = *height
	}

	var resp struct {
		Height cosmostypes.Uint      `json:"height"`
		Result GetEpochStateResponse `json:"result"`
	}

	if err := m.Query(ctx, types.Q{"epoch_state": query}, &resp); err != nil {
		return GetEpochStateResponse{}, err
	}
	return resp.Result, nil
}

func (m market) GetBorrowerInfo(ctx context.Context, borrower cosmostypes.AccAddress, height *uint64) (GetBorrowerInfoResponse, error) {
	var query = types.Q{"borrower": borrower.String()}
	if height != nil {
		query["block_height"] = *height
	}

	var resp struct {
		Height cosmostypes.Uint        `json:"height"`
		Result GetBorrowerInfoResponse `json:"result"`
	}

	if err := m.Query(ctx, types.Q{"borrower_info": query}, &resp); err != nil {
		return GetBorrowerInfoResponse{}, err
	}
	return resp.Result, nil
}

func (m market) GetBorrowerInfos(ctx context.Context, startAfter *cosmostypes.AccAddress, limit *uint32) ([]GetBorrowerInfoResponse, error) {
	var query = types.Q{}
	if startAfter != nil {
		query["start_after"] = (*startAfter).String()
	}
	if limit != nil {
		query["limit"] = *limit
	}

	var resp struct {
		Height cosmostypes.Uint         `json:"height"`
		Result GetBorrowerInfosResponse `json:"result"`
	}

	if err := m.Query(ctx, types.Q{"borrower_infos": query}, &resp); err != nil {
		return nil, err
	}
	return resp.Result.BorrowerInfos, nil
}
