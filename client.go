package terra

import (
	"github.com/frostornge/terra-go/httpclient"
	"github.com/frostornge/terra-go/service"
)

var _ Client = (*terraClient)(nil)

//go:generate mockgen -destination ../../test/mocks/terra/client.go . Client
type Client interface {
	Auth() service.AuthService
	Bank() service.BankService
	Contract() service.ContractService
	Treasury() service.TreasuryService
	Tendermint() service.TendermintService
	Transaction() service.TransactionService
}

type terraClient struct {
	client httpclient.Client

	auth        service.AuthService
	bank        service.BankService
	contract    service.ContractService
	treasury    service.TreasuryService
	tendermint  service.TendermintService
	transaction service.TransactionService
}

func (c terraClient) Auth() service.AuthService               { return c.auth }
func (c terraClient) Bank() service.BankService               { return c.bank }
func (c terraClient) Contract() service.ContractService       { return c.contract }
func (c terraClient) Treasury() service.TreasuryService       { return c.treasury }
func (c terraClient) Tendermint() service.TendermintService   { return c.tendermint }
func (c terraClient) Transaction() service.TransactionService { return c.transaction }

func NewClient(client httpclient.Client) Client {
	return terraClient{
		auth:        service.NewAuthService(client),
		bank:        service.NewBankService(client),
		contract:    service.NewContractService(client),
		treasury:    service.NewTreasuryService(client),
		tendermint:  service.NewTendermintService(client),
		transaction: service.NewTransactionService(client),
	}
}
