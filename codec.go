package terra

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cosmostypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	distrclient "github.com/cosmos/cosmos-sdk/x/distribution/client"
	paramsclient "github.com/cosmos/cosmos-sdk/x/params/client"
	upgradeclient "github.com/cosmos/cosmos-sdk/x/upgrade/client"
	"github.com/frostornge/terra-go/types"
	"github.com/terra-project/core/x/auth"
	"github.com/terra-project/core/x/auth/vesting"
	"github.com/terra-project/core/x/bank"
	"github.com/terra-project/core/x/crisis"
	"github.com/terra-project/core/x/distribution"
	"github.com/terra-project/core/x/evidence"
	"github.com/terra-project/core/x/genutil"
	"github.com/terra-project/core/x/gov"
	"github.com/terra-project/core/x/market"
	"github.com/terra-project/core/x/mint"
	"github.com/terra-project/core/x/msgauth"
	"github.com/terra-project/core/x/oracle"
	"github.com/terra-project/core/x/params"
	"github.com/terra-project/core/x/slashing"
	"github.com/terra-project/core/x/staking"
	"github.com/terra-project/core/x/supply"
	"github.com/terra-project/core/x/treasury"
	treasuryclient "github.com/terra-project/core/x/treasury/client"
	"github.com/terra-project/core/x/upgrade"
	"github.com/terra-project/core/x/wasm"
)

func MakeCodec() *codec.Codec {
	var cdc = codec.New()
	cdc.RegisterConcrete(wasm.MsgStoreCode{}, "wasm/MsgStoreCode", nil)
	cdc.RegisterConcrete(wasm.MsgInstantiateContract{}, "wasm/MsgInstantiateContract", nil)
	cdc.RegisterConcrete(types.MsgExecuteContract{}, "wasm/MsgExecuteContract", nil)
	cdc.RegisterConcrete(wasm.MsgMigrateContract{}, "wasm/MsgMigrateContract", nil)
	cdc.RegisterConcrete(wasm.MsgUpdateContractOwner{}, "wasm/MsgUpdateContractOwner", nil)

	module.NewBasicManager(
		genutil.AppModuleBasic{},
		auth.AppModuleBasic{},
		bank.AppModuleBasic{},
		staking.AppModuleBasic{},
		mint.AppModuleBasic{},
		distribution.AppModuleBasic{},
		gov.NewAppModuleBasic(
			paramsclient.ProposalHandler,
			distrclient.ProposalHandler,
			upgradeclient.ProposalHandler,
			treasuryclient.TaxRateUpdateProposalHandler,
			treasuryclient.RewardWeightUpdateProposalHandler,
		),
		params.AppModuleBasic{},
		crisis.AppModuleBasic{},
		slashing.AppModuleBasic{},
		supply.AppModuleBasic{},
		evidence.AppModuleBasic{},
		upgrade.AppModuleBasic{},
		oracle.AppModuleBasic{},
		market.AppModuleBasic{},
		treasury.AppModuleBasic{},
		msgauth.AppModuleBasic{},
	).RegisterCodec(cdc)
	vesting.RegisterCodec(cdc)
	cosmostypes.RegisterCodec(cdc)
	codec.RegisterCrypto(cdc)
	codec.RegisterEvidences(cdc)

	return cdc
}
