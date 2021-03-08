package bind

import (
	"context"

	"github.com/frostornge/terra-go"

	cosmostypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/pkg/errors"
)

type BaseContract interface {
	GetAddress() cosmostypes.AccAddress
	GetClient() terra.Client

	MakeMessage(
		acc terra.Account,
		method string,
		payload interface{},
		coins cosmostypes.Coins,
	) ([]cosmostypes.Msg, error)

	Execute(
		ctx context.Context,
		acc terra.Account,
		msgs []cosmostypes.Msg,
		mode *terra.BroadcastMode,
		opts *terra.CreateTxOptions,
	) (cosmostypes.TxResponse, error)

	Query(ctx context.Context, query terra.Q, resp interface{}) error
}

type baseContract struct {
	addr   cosmostypes.AccAddress
	client terra.Client
}

func NewBaseContract(addr cosmostypes.AccAddress, client terra.Client) BaseContract {
	return baseContract{
		addr:   addr,
		client: client,
	}
}

func (b baseContract) GetAddress() cosmostypes.AccAddress { return b.addr }
func (b baseContract) GetClient() terra.Client            { return b.client }

func (b baseContract) MakeMessage(
	acc terra.Account,
	method string,
	payload interface{},
	coins cosmostypes.Coins,
) ([]cosmostypes.Msg, error) {
	if payload == nil {
		payload = terra.Q{}
	}

	executeMsg, err := terra.GenerateExecuteMsg(
		acc.GetAddress(),
		b.addr,
		terra.Q{method: payload},
		coins,
	)
	if err != nil {
		return nil, errors.Wrap(err, "make message")
	}
	return []cosmostypes.Msg{executeMsg}, nil
}

func (b baseContract) Execute(
	ctx context.Context,
	acc terra.Account,
	msgs []cosmostypes.Msg,
	mode *terra.BroadcastMode,
	opts *terra.CreateTxOptions,
) (cosmostypes.TxResponse, error) {
	createOption := terra.CreateTxOptions{GasAdjustment: 1.2}
	if opts != nil {
		createOption = *opts
	}
	createOption.Msgs = msgs

	tx, err := acc.CreateAndSignTx(ctx, createOption)
	if err != nil {
		return cosmostypes.TxResponse{}, errors.Wrap(err, "sign tx")
	}

	broadcastMode := terra.ModeBlock
	if mode != nil {
		broadcastMode = *mode
	}

	resp, err := b.client.Transaction().BroadcastTx(ctx, tx, broadcastMode)
	if err != nil {
		return cosmostypes.TxResponse{}, errors.Wrap(err, "broadcast tx")
	}
	return resp, nil
}

func (b baseContract) Query(ctx context.Context, query terra.Q, resp interface{}) error {
	if err := b.client.Contract().QueryContractStore(ctx, b.addr, query, resp); err != nil {
		return errors.Wrap(err, "query contract store")
	}
	return nil
}
