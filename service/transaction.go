package service

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/frostornge/terra-go/httpclient"
	"github.com/frostornge/terra-go/types"

	"github.com/cosmos/cosmos-sdk/codec"
	cosmostypes "github.com/cosmos/cosmos-sdk/types"
	cosmosauthrest "github.com/cosmos/cosmos-sdk/x/auth/client/rest"
	"github.com/pkg/errors"
	abcitypes "github.com/tendermint/tendermint/abci/types"
	terraauth "github.com/terra-project/core/x/auth"
	terraauthutils "github.com/terra-project/core/x/auth/client/utils"
)

//go:generate mockgen -destination ../../../test/mocks/terra/service/service_transaction.go . TransactionService
type TransactionService interface {
	GetTxByHash(ctx context.Context, txHash string) (cosmostypes.TxResponse, error)
	QueryTx(ctx context.Context, req QueryTxRequest) (QueryTxResponse, error)
	BroadcastTx(
		ctx context.Context,
		tx terraauth.StdTx,
		mode types.BroadcastMode,
	) (cosmostypes.TxResponse, error)
	EstimateFee(
		ctx context.Context,
		tx terraauth.StdTx,
		gasAdjustment string,
		gasPrices cosmostypes.DecCoins,
	) (terraauth.StdFee, error)
}

type transactionService struct {
	codec  *codec.Codec
	client httpclient.Client
}

func NewTransactionService(client httpclient.Client) TransactionService {
	return transactionService{codec: client.Codec(), client: client}
}

func (svc transactionService) GetTxByHash(
	ctx context.Context,
	txHash string,
) (cosmostypes.TxResponse, error) {
	var payload = httpclient.RequestPayload{
		Context: ctx,
		Method:  http.MethodGet,
		Path:    fmt.Sprintf("/txs/%s", txHash),
	}

	var body cosmostypes.TxResponse
	if err := svc.client.RequestJSON(payload, &body); err != nil {
		return cosmostypes.TxResponse{}, errors.Wrap(err, "request json")
	}
	return body, nil
}

func (svc transactionService) QueryTx(ctx context.Context, req QueryTxRequest) (QueryTxResponse, error) {
	var payload = httpclient.RequestPayload{
		Context: ctx,
		Method:  http.MethodGet,
		Path:    "/txs",
		Query:   make(map[string]string),
	}

	if req.Page != nil {
		req.Query["page"] = *req.Page
	}
	if req.Limit != nil {
		req.Query["limit"] = *req.Limit
	}

	for k, v := range req.Query {
		payload.Query[k] = fmt.Sprintf("%v", v)
	}

	var body QueryTxResponse
	if err := svc.client.RequestJSON(payload, &body); err != nil {
		return QueryTxResponse{}, errors.Wrap(err, "request json")
	}
	return body, nil
}

func (svc transactionService) BroadcastTx(
	ctx context.Context,
	tx terraauth.StdTx,
	mode types.BroadcastMode,
) (cosmostypes.TxResponse, error) {
	var req = cosmosauthrest.BroadcastReq{
		Tx:   tx,
		Mode: string(mode),
	}

	rawPayloadBody, err := svc.codec.MarshalJSON(req)
	if err != nil {
		return cosmostypes.TxResponse{}, errors.Wrap(err, "marshal request body")
	}

	var payload = httpclient.RequestPayload{
		Context: ctx,
		Method:  http.MethodPost,
		Path:    "/txs",
		Body:    bytes.NewReader(rawPayloadBody),
	}

	var body cosmostypes.TxResponse
	if err := svc.client.RequestJSON(payload, &body); err != nil {
		return cosmostypes.TxResponse{}, errors.Wrap(err, "request json")
	}
	time.Sleep(1 * time.Second) // wait for lcd

	if body.Code != abcitypes.CodeTypeOK {
		return body, errors.New(body.RawLog)
	}
	return body, nil
}

func (svc transactionService) EstimateFee(
	ctx context.Context,
	tx terraauth.StdTx,
	gasAdjustment string,
	gasPrices cosmostypes.DecCoins,
) (terraauth.StdFee, error) {
	var req = terraauthutils.EstimateFeeReq{
		Tx:            tx,
		GasAdjustment: gasAdjustment,
		GasPrices:     gasPrices,
	}

	rawPayloadBody, err := svc.codec.MarshalJSON(req)
	if err != nil {
		return terraauth.StdFee{}, errors.Wrap(err, "marshal request body")
	}

	var payload = httpclient.RequestPayload{
		Context: ctx,
		Method:  http.MethodPost,
		Path:    "/txs/estimate_fee",
		Body:    bytes.NewReader(rawPayloadBody),
	}

	var body struct {
		Height string                         `json:"height"`
		Result terraauthutils.EstimateFeeResp `json:"result"`
	}
	if err := svc.client.RequestJSON(payload, &body); err != nil {
		return terraauth.StdFee{}, errors.Wrap(err, "request json")
	}
	return terraauth.NewStdFee(body.Result.Gas, body.Result.Fees), nil
}
