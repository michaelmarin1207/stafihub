package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	relayertypes "github.com/stafihub/stafihub/x/relayers/types"
)

type SudoKeeper interface {
	IsAdmin(ctx sdk.Context, address string) bool
	GetAdmin(ctx sdk.Context) sdk.AccAddress
}

// BankKeeper defines the contract needed to be fulfilled for banking and supply
// dependencies.
type BankKeeper interface {
	//HasDenomMetaData(ctx sdk.Context, denom string) bool
	GetDenomMetaData(ctx sdk.Context, denom string) (banktypes.Metadata, bool)
	GetSupply(ctx sdk.Context, denom string) sdk.Coin
	SendCoinsFromModuleToAccount(ctx sdk.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error
	SendCoinsFromModuleToModule(ctx sdk.Context, senderModule, recipientModule string, amt sdk.Coins) error
	SendCoinsFromAccountToModule(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
	MintCoins(ctx sdk.Context, name string, amt sdk.Coins) error
	BurnCoins(ctx sdk.Context, moduleName string, amt sdk.Coins) error
	GetBalance(ctx sdk.Context, addr sdk.AccAddress, denom string) sdk.Coin
	SendCoins(ctx sdk.Context, fromAddr sdk.AccAddress, toAddr sdk.AccAddress, amt sdk.Coins) error
}

type RelayerKeeper interface {
	IsRelayer(ctx sdk.Context, denom, address string) bool
	LastVoter(ctx sdk.Context, denom string) (val relayertypes.LastVoter, found bool)
}
