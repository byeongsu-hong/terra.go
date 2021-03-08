package service

import (
	"context"
	"fmt"
	"net/http"

	"github.com/frostornge/terra-go/httpclient"

	cosmosrpc "github.com/cosmos/cosmos-sdk/client/rpc"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/pkg/errors"
	"github.com/tendermint/tendermint/p2p"
	tdmttypes "github.com/tendermint/tendermint/types"
)

//go:generate mockgen -destination ../../../test/mocks/terra/service/service_tendermint.go . TendermintService
type TendermintService interface {
	GetNodeInfo(ctx context.Context) (p2p.DefaultNodeInfo, error)
	GetSyncStatus(ctx context.Context) (cosmosrpc.SyncingResponse, error)
	GetBlockByHeight(ctx context.Context, height *uint64) (tdmttypes.BlockID, *tdmttypes.Block, error)
}

type tendermintService struct {
	codec  *codec.Codec
	client httpclient.Client
}

func NewTendermintService(client httpclient.Client) TendermintService {
	return tendermintService{codec: client.Codec(), client: client}
}

func (svc tendermintService) GetNodeInfo(ctx context.Context) (p2p.DefaultNodeInfo, error) {
	var payload = httpclient.RequestPayload{
		Context: ctx,
		Method:  http.MethodGet,
		Path:    "/node_info",
	}

	var body struct {
		NodeInfo p2p.DefaultNodeInfo `json:"node_info"`
	}
	if err := svc.client.RequestJSON(payload, &body); err != nil {
		return p2p.DefaultNodeInfo{}, errors.Wrap(err, "request json")
	}
	return body.NodeInfo, nil
}

func (svc tendermintService) GetSyncStatus(ctx context.Context) (cosmosrpc.SyncingResponse, error) {
	var payload = httpclient.RequestPayload{
		Context: ctx,
		Method:  http.MethodGet,
		Path:    "/syncing",
	}

	var body cosmosrpc.SyncingResponse
	if err := svc.client.RequestJSON(payload, &body); err != nil {
		return cosmosrpc.SyncingResponse{}, errors.Wrap(err, "request json")
	}
	return body, nil
}

func (svc tendermintService) GetBlockByHeight(ctx context.Context, height *uint64) (tdmttypes.BlockID, *tdmttypes.Block, error) {
	var path = "/blocks/latest"
	if height != nil {
		path = fmt.Sprintf("/blocks/%d", *height)
	}

	var payload = httpclient.RequestPayload{
		Context: ctx,
		Method:  http.MethodGet,
		Path:    path,
	}

	var body struct {
		BlockID tdmttypes.BlockID `json:"block_id"`
		Block   tdmttypes.Block   `json:"block"`
	}
	if err := svc.client.RequestJSON(payload, &body); err != nil {
		return tdmttypes.BlockID{}, nil, errors.Wrap(err, "request json")
	}
	return body.BlockID, &body.Block, nil
}
