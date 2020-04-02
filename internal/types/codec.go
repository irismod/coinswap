package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

// ModuleCdc defines the module codec
var ModuleCdc *codec.Codec

// RegisterCodec registers concrete types on the codec.
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgSwapOrder{}, "irismod/coinswap/MsgSwapOrder", nil)
	cdc.RegisterConcrete(MsgAddLiquidity{}, "irismod/coinswap/MsgAddLiquidity", nil)
	cdc.RegisterConcrete(MsgRemoveLiquidity{}, "irismod/coinswap/MsgRemoveLiquidity", nil)
	cdc.RegisterConcrete(&Params{}, "irismod/coinswap/Params", nil)
}

func init() {
	ModuleCdc = codec.New()
	RegisterCodec(ModuleCdc)
	ModuleCdc.Seal()
}
