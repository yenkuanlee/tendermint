package core_types

import (
	amino "github.com/tendermint/go-amino"
	"github.com/yenkuanlee/tendermint/types"
)

func RegisterAmino(cdc *amino.Codec) {
	types.RegisterEventDatas(cdc)
	types.RegisterBlockAmino(cdc)
}
