module github.com/frostornge/terra-go

go 1.15

replace github.com/cosmos/ledger-cosmos-go => github.com/terra-money/ledger-terra-go v0.11.2

replace google.golang.org/grpc => google.golang.org/grpc v1.33.2

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1

replace github.com/99designs/keyring => github.com/cosmos/keyring v1.1.7-0.20210622111912-ef00f8ac3d76

require (
	github.com/airbloc/logger v1.4.7
	github.com/aws/aws-sdk-go v1.40.37 // indirect
	github.com/cosmos/cosmos-sdk v0.44.0
	github.com/maruel/panicparse v1.6.1 // indirect
	github.com/pkg/errors v0.9.1
	github.com/tendermint/tendermint v0.34.12
	github.com/terra-money/core v0.5.2
	gopkg.in/yaml.v2 v2.4.0 // indirect
)
