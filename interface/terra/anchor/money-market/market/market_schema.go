package market

import cosmostypes "github.com/cosmos/cosmos-sdk/types"

// ================= Query ================= //

type GetConfigResponse struct {
	OwnerAddr         cosmostypes.AccAddress `json:"owner_addr"`
	ATerraContract    cosmostypes.AccAddress `json:"aterra_contract"`
	InterestModel     cosmostypes.AccAddress `json:"interest_model"`
	DistributionModel cosmostypes.AccAddress `json:"distribution_model"`
	OverseerContract  cosmostypes.AccAddress `json:"overseer_contract"`
	CollectorContract cosmostypes.AccAddress `json:"collector_contract"`
	FaucetContract    cosmostypes.AccAddress `json:"faucet_contract"`
	StableDenom       string                 `json:"stable_denom"`
	ReserveFactor     cosmostypes.Dec        `json:"reserve_factor"`
	MaxBorrowFactor   cosmostypes.Dec        `json:"max_borrow_factor"`
}

type GetStateResponse struct {
	TotalLiabilities    cosmostypes.Dec `json:"total_liabilities"`
	TotalReserves       cosmostypes.Dec `json:"total_reserves"`
	LastInterestUpdated uint64          `json:"last_interest_updated"`
	LastRewardUpdated   uint64          `json:"last_reward_updated"`
	GlobalInterestIndex cosmostypes.Dec `json:"global_interest_index"`
	GlobalRewardIndex   cosmostypes.Dec `json:"global_reward_index"`
	ANCEmissionRate     cosmostypes.Dec `json:"anc_emission_rate"`
}

type GetEpochStateResponse struct {
	ExchangeRate cosmostypes.Dec `json:"exchange_rate"`
	ATokenSupply cosmostypes.Int `json:"a_token_supply"`
}

type GetBorrowerInfoResponse struct {
	Borrower       cosmostypes.AccAddress `json:"borrower"`
	InterestIndex  cosmostypes.Dec        `json:"interest_index"`
	RewardIndex    cosmostypes.Dec        `json:"reward_index"`
	LoanAmount     cosmostypes.Int        `json:"loan_amount"`
	PendingRewards cosmostypes.Dec        `json:"pending_rewards"`
}

type GetBorrowerInfosResponse struct {
	BorrowerInfos []GetBorrowerInfoResponse `json:"borrower_infos"`
}
