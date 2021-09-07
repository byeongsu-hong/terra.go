package terra

import (
	"context"
	"sync"

	cosmosclienttx "github.com/cosmos/cosmos-sdk/client/tx"

	cosmoscryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	cosmostypes "github.com/cosmos/cosmos-sdk/types"
	cosmostx "github.com/cosmos/cosmos-sdk/types/tx"
	cosmosauth "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/pkg/errors"
	terraassets "github.com/terra-money/core/types/assets"
)

var (
	DefaultGasAdjustment = 1.2
	DefaultGasPrice      = cosmostypes.DecCoins{{
		Denom:  terraassets.MicroLunaDenom,
		Amount: cosmostypes.NewDecWithPrec(150000, 6),
	}}
)

type Account interface {
	GetClient() Client
	GetChainId() string
	Update(ctx context.Context) error
	CreateTx(ctx context.Context, opts CreateTxOptions) (cosmostx.Tx, error)
	CreateAndSignTx(ctx context.Context, opts CreateTxOptions) (cosmostx.Tx, error)

	cosmosauth.AccountI
}

type keyedAccount struct {
	key     Key
	client  Client
	chainId string
	mutex   sync.Mutex
	cosmosauth.AccountI
}

func NewAccount(ctx context.Context, client Client, key Key) (Account, error) {
	nodeInfo, err := client.Tendermint().GetNodeInfo(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "fetch node info")
	}

	acc := &keyedAccount{
		key:     key,
		client:  client,
		chainId: nodeInfo.Network,
	}
	if err := acc.Update(ctx); err != nil {
		return nil, errors.Wrap(err, "update account")
	}

	return acc, nil
}

func (a *keyedAccount) GetAddress() cosmostypes.AccAddress  { return a.key.AccAddress() }
func (a *keyedAccount) GetPubKey() cosmoscryptotypes.PubKey { return a.key.PubKey() }
func (a *keyedAccount) GetClient() Client                   { return a.client }
func (a *keyedAccount) GetChainId() string                  { return a.chainId }

func (a *keyedAccount) Update(ctx context.Context) error {
	a.mutex.Lock()
	defer a.mutex.Unlock()

	accInfo, err := a.client.Auth().GetAccountInfo(ctx, a.key.AccAddress())
	if err != nil {
		return errors.Wrap(err, "fetch account info")
	}

	a.AccountI = accInfo
	return nil
}

type CreateTxOptions struct {
	Msgs          []cosmostypes.Msg
	Fee           *cosmostx.Fee
	GasAdjustment float64
	GasPrices     cosmostypes.DecCoins
	Sequence      *uint64
	Memo          string
}

func (a *keyedAccount) CreateTx(ctx context.Context, opts CreateTxOptions) (cosmostx.Tx, error) {
	if opts.GasAdjustment == 0 {
		opts.GasAdjustment = DefaultGasAdjustment
	}
	if opts.GasPrices == nil || len(opts.GasPrices) == 0 {
		opts.GasPrices = DefaultGasPrice
	}
	if err := a.Update(ctx); err != nil {
		return cosmostx.Tx{}, errors.Wrap(err, "update account")
	}

	resp, err := a.client.Bank().GetBalance(ctx, a.GetAddress())
	if err != nil {
		return cosmostx.Tx{}, errors.Wrapf(err, "get balance of %s", a.GetAddress().String())
	}
	for index, coin := range resp.Balance {
		resp.Balance[index] = cosmostypes.NewCoin(
			coin.Denom,
			cosmostypes.NewInt(1),
		)
	}

	sequence := a.GetSequence()
	if opts.Sequence != nil {
		if sequence < *opts.Sequence {
			sequence = *opts.Sequence
		}
	}

	factory := cosmosclienttx.Factory{}.
		WithChainID(a.chainId).
		WithAccountNumber(a.GetAccountNumber()).
		WithSequence(sequence).
		WithMemo(opts.Memo).
		WithGasAdjustment(opts.GasAdjustment)

	var fee cosmostx.Fee
	if opts.Fee == nil {
		simTxBytes, err := factory.BuildSimTx(opts.Msgs...)
		if err != nil {
			return cosmostx.Tx{}, errors.Wrap(err, "build simulate tx")
		}

		gasInfo, err := a.client.Transaction().Simulate(ctx, simTxBytes)
		if err != nil {
			return cosmostx.Tx{}, errors.Wrap(err, "estimate fee")
		}
		factory.WithGas(uint64(float64(gasInfo.GetGasUsed()) * opts.GasAdjustment))
	} else {
		fee = *opts.Fee
	}

	return terraauth.StdSignMsg{
		ChainID:       a.chainId,
		AccountNumber: a.GetAccountNumber(),
		Sequence:      sequence,
		Fee:           fee,
		Msgs:          opts.Msgs,
		Memo:          opts.Memo,
	}, nil
}

func (a *keyedAccount) CreateAndSignTx(ctx context.Context, opts CreateTxOptions) (terraauth.StdTx, terraauth.StdSignMsg, error) {
	signMsg, err := a.CreateTx(ctx, opts)
	if err != nil {
		return terraauth.StdTx{}, terraauth.StdSignMsg{}, errors.Wrap(err, "create tx")
	}

	signedTx, err := a.key.SignTx(signMsg)
	if err != nil {
		return terraauth.StdTx{}, terraauth.StdSignMsg{}, errors.Wrap(err, "sign tx")
	}
	return signedTx, signMsg, nil
}
