package coinswap

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/irismod/coinswap/keeper"
	"github.com/irismod/coinswap/types"
)

// InitGenesis new coinswap genesis
func InitGenesis(ctx sdk.Context, k keeper.Keeper, data types.GenesisState) {
	if err := types.ValidateGenesis(data); err != nil {
		panic(fmt.Errorf("panic for ValidateGenesis,%v", err))
	}
	k.SetParams(ctx, data.Params)
}

// ExportGenesis returns a GenesisState for a given context and keeper.
func ExportGenesis(ctx sdk.Context, keeper keeper.Keeper) *types.GenesisState {
	return types.NewGenesisState(keeper.GetParams(ctx))
}
