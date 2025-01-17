package keeper

import (
	"fmt"

	"github.com/cometbft/cometbft/libs/log"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"

	"github.com/osmosis-labs/osmosis/v23/x/bridge/types"
)

type Keeper struct {
	storeKey   storetypes.StoreKey
	paramSpace paramtypes.Subspace

	accountKeeper      types.AccountKeeper
	tokenFactoryKeeper types.TokenFactoryKeeper

	govModuleAddr string
}

// NewKeeper returns a new instance of the x/bridge keeper.
func NewKeeper(
	storeKey storetypes.StoreKey,
	paramSpace paramtypes.Subspace,
	accountKeeper types.AccountKeeper,
	tokenFactoryKeeper types.TokenFactoryKeeper,
	govModuleAddr string,
) Keeper {
	// ensure bridge module account is set
	if addr := accountKeeper.GetModuleAddress(types.ModuleName); addr == nil {
		panic("the bridge module account has not been set")
	}

	if !paramSpace.HasKeyTable() {
		paramSpace = paramSpace.WithKeyTable(types.ParamKeyTable())
	}

	return Keeper{
		storeKey:           storeKey,
		paramSpace:         paramSpace,
		accountKeeper:      accountKeeper,
		tokenFactoryKeeper: tokenFactoryKeeper,
		govModuleAddr:      govModuleAddr,
	}
}

// Logger returns a logger for the x/bridge module.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
