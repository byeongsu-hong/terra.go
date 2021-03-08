package service

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/frostornge/terra-go/httpclient"

	"github.com/cosmos/cosmos-sdk/codec"
	cosmostypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/pkg/errors"
	terrawasm "github.com/terra-project/core/x/wasm"
)

//go:generate mockgen -destination ../../../test/mocks/terra/service/service_contract.go . ContractService
type ContractService interface {
	GetCodeID(ctx context.Context, codeId uint64) (terrawasm.CodeInfo, error)
	GetContractInfo(ctx context.Context, addr cosmostypes.AccAddress) (terrawasm.ContractInfo, error)
	QueryContractStore(ctx context.Context, addr cosmostypes.AccAddress, query interface{}, resp interface{}) error
}

type contractService struct {
	codec  *codec.Codec
	client httpclient.Client
}

func NewContractService(client httpclient.Client) ContractService {
	return contractService{codec: client.Codec(), client: client}
}

func (svc contractService) GetCodeID(ctx context.Context, codeId uint64) (terrawasm.CodeInfo, error) {
	var payload = httpclient.RequestPayload{
		Context: ctx,
		Method:  http.MethodGet,
		Path:    fmt.Sprintf("/wasm/codes/%d", codeId),
	}

	var body struct {
		Height string             `json:"height"`
		Result terrawasm.CodeInfo `json:"result"`
	}
	if err := svc.client.RequestJSON(payload, &body); err != nil {
		return terrawasm.CodeInfo{}, errors.Wrap(err, "request json")
	}
	return body.Result, nil
}

func (svc contractService) GetContractInfo(
	ctx context.Context,
	addr cosmostypes.AccAddress,
) (terrawasm.ContractInfo, error) {
	var payload = httpclient.RequestPayload{
		Context: ctx,
		Method:  http.MethodGet,
		Path:    fmt.Sprintf("/wasm/contracts/%s", addr.String()),
	}

	var body struct {
		Height string                 `json:"height"`
		Result terrawasm.ContractInfo `json:"result"`
	}
	if err := svc.client.RequestJSON(payload, &body); err != nil {
		return terrawasm.ContractInfo{}, errors.Wrap(err, "request json")
	}
	return body.Result, nil
}

func (svc contractService) QueryContractStore(
	ctx context.Context,
	addr cosmostypes.AccAddress,
	query interface{},
	resp interface{},
) error {
	jsonQuery, err := json.Marshal(query)
	if err != nil {
		return errors.Wrap(err, "marshal query message")
	}

	var payload = httpclient.RequestPayload{
		Context: ctx,
		Method:  http.MethodGet,
		Path:    fmt.Sprintf("/wasm/contracts/%s/store", addr.String()),
		Query:   map[string]string{"query_msg": string(jsonQuery)},
	}

	if err := svc.client.RequestJSON(payload, resp); err != nil {
		return errors.Wrap(err, "request json")
	}
	return nil
}
