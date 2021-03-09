package overseer

import (
	"github.com/frostornge/terra-go/types"

	cosmostypes "github.com/cosmos/cosmos-sdk/types"
)

type GetConfigResponse struct {
	Height                 cosmostypes.Uint       `json:"-"`
	OwnerAddr              cosmostypes.AccAddress `json:"owner_addr"`
	OracleContract         cosmostypes.AccAddress `json:"oracle_contract"`
	MarketContract         cosmostypes.AccAddress `json:"market_contract"`
	LiquidationContract    cosmostypes.AccAddress `json:"liquidation_contract"`
	CollectorContract      cosmostypes.AccAddress `json:"collector_contract"`
	DistributionThreshold  cosmostypes.Dec        `json:"distribution_threshold"`
	TargetDepositRate      cosmostypes.Dec        `json:"target_deposit_rate"`
	BufferDistributionRate cosmostypes.Dec        `json:"buffer_distribution_rate"`
	ANCPurchaseFactor      cosmostypes.Dec        `json:"anc_purchase_factor"`
	StableDenom            string                 `json:"stable_denom"`
	EpochPeriod            cosmostypes.Uint       `json:"epoch_period"`
	PriceTimeFrame         cosmostypes.Uint       `json:"price_time_frame"`
}

type GetEpochStateResponse struct {
	Height             cosmostypes.Uint `json:"-"`
	DepositRate        cosmostypes.Dec  `json:"deposit_rate"`
	PrevATerraSupply   cosmostypes.Uint `json:"prev_a_terra_supply"`
	PrevExchangeRate   cosmostypes.Dec  `json:"prev_exchange_rate"`
	LastExecutedHeight cosmostypes.Uint `json:"last_executed_height"`
}

type GetWhitelistResponse struct {
	Height          cosmostypes.Uint       `json:"-"`
	Name            string                 `json:"name"`
	Symbol          string                 `json:"symbol"`
	MaxLTV          cosmostypes.Dec        `json:"max_ltv"`
	CustodyContract cosmostypes.AccAddress `json:"custody_contract"`
	CollateralToken cosmostypes.AccAddress `json:"collateral_token"`
}

type GetCollateralsResponse struct {
	Height      cosmostypes.Uint       `json:"-"`
	Borrower    cosmostypes.AccAddress `json:"borrower"`
	Collaterals []types.TokensHuman    `json:"collaterals"`
}

type GetAllCollateralsResponse struct {
	Height         cosmostypes.Uint `json:"-"`
	AllCollaterals []struct {
		Borrower    cosmostypes.AccAddress `json:"borrower"`
		Collaterals []types.TokensHuman    `json:"collaterals"`
	} `json:"all_collaterals"`
}

type GetDistributionParamsResponse struct {
	Height               cosmostypes.Uint `json:"-"`
	DepositRate          cosmostypes.Dec  `json:"deposit_rate"`
	TargetDepositRate    cosmostypes.Dec  `json:"target_deposit_rate"`
	ThresholdDepositRate cosmostypes.Dec  `json:"threshold_deposit_rate"`
}

type GetBorrowLimitResponse struct {
	Height      cosmostypes.Uint       `json:"-"`
	Borrower    cosmostypes.AccAddress `json:"borrower"`
	BorrowLimit cosmostypes.Int        `json:"borrow_limit"`
}
