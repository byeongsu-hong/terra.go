package service

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"time"

	cosmostypes "github.com/cosmos/cosmos-sdk/types"
	cosmostx "github.com/cosmos/cosmos-sdk/types/tx"
	"github.com/frostornge/terra-go/httpclient"
	"github.com/pkg/errors"
	abcitypes "github.com/tendermint/tendermint/abci/types"
	encoding "github.com/terra-money/core/app/params"
)

//go:generate mockgen -destination ../../../test/mocks/terra/service/service_transaction.go . TransactionService
type TransactionService interface {
	GetTx(ctx context.Context, txHash string) (cosmostypes.TxResponse, error)
	BroadcastTx(
		ctx context.Context,
		tx cosmostx.Tx,
		mode cosmostx.BroadcastMode,
	) (cosmostypes.TxResponse, error)
	Simulate(
		ctx context.Context,
		txBytes []byte,
	) (cosmostypes.GasInfo, error)
}

type transactionService struct {
	codec  encoding.EncodingConfig
	client httpclient.Client
}

func NewTransactionService(client httpclient.Client) TransactionService {
	return transactionService{codec: client.Codec(), client: client}
}

func (svc transactionService) GetTx(
	ctx context.Context,
	txHash string,
) (cosmostypes.TxResponse, error) {
	var payload = httpclient.RequestPayload{
		Context: ctx,
		Method:  http.MethodGet,
		Path:    fmt.Sprintf("/cosmos/tx/v1beta1/txs/{%s}", txHash),
	}

	var body cosmostx.GetTxResponse
	if err := svc.client.RequestJSON(payload, &body); err != nil {
		return cosmostypes.TxResponse{}, errors.Wrap(err, "request json")
	}
	return *body.GetTxResponse(), nil
}

func (svc transactionService) BroadcastTx(
	ctx context.Context,
	tx cosmostx.Tx,
	mode cosmostx.BroadcastMode,
) (cosmostypes.TxResponse, error) {
	txBytes, err := tx.Marshal()
	if err != nil {
		return cosmostypes.TxResponse{}, errors.Wrap(err, "marshal tx")
	}

	var req = cosmostx.BroadcastTxRequest{TxBytes: txBytes, Mode: mode}
	rawPayloadBody, err := svc.codec.Marshaler.MarshalJSON(&req)
	if err != nil {
		return cosmostypes.TxResponse{}, errors.Wrap(err, "marshal request body")
	}

	var payload = httpclient.RequestPayload{
		Context: ctx,
		Method:  http.MethodPost,
		Path:    "/cosmos/tx/v1beta1/txs",
		Body:    bytes.NewReader(rawPayloadBody),
	}

	var body cosmostx.BroadcastTxResponse
	if err := svc.client.RequestJSON(payload, &body); err != nil {
		return cosmostypes.TxResponse{}, errors.Wrap(err, "request json")
	}
	time.Sleep(1 * time.Second) // wait for lcd

	if body.TxResponse.Code != abcitypes.CodeTypeOK {
		return cosmostypes.TxResponse{}, errors.New(body.TxResponse.RawLog)
	}
	return *body.GetTxResponse(), nil
}

func (svc transactionService) Simulate(
	ctx context.Context,
	txBytes []byte,
) (cosmostypes.GasInfo, error) {
	var req = cosmostx.SimulateRequest{TxBytes: txBytes}
	rawPayloadBody, err := svc.codec.Marshaler.MarshalJSON(&req)
	if err != nil {
		return cosmostypes.GasInfo{}, errors.Wrap(err, "marshal request body")
	}

	var payload = httpclient.RequestPayload{
		Context: ctx,
		Method:  http.MethodPost,
		Path:    "/cosmos/tx/v1beta1/simulate",
		Body:    bytes.NewReader(rawPayloadBody),
	}

	var body cosmostx.SimulateResponse
	if err := svc.client.RequestJSON(payload, &body); err != nil {
		return cosmostypes.GasInfo{}, errors.Wrap(err, "request json")
	}
	return *body.GetGasInfo(), nil
}
