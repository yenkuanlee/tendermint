package core

import (
	"github.com/yenkuanlee/tendermint/rpc/coretypes"
	rpctypes "github.com/yenkuanlee/tendermint/rpc/jsonrpc/types"
)

// Health gets node health. Returns empty result (200 OK) on success, no
// response - in case of an error.
// More: https://docs.tendermint.com/master/rpc/#/Info/health
func (env *Environment) Health(ctx *rpctypes.Context) (*coretypes.ResultHealth, error) {
	return &coretypes.ResultHealth{}, nil
}
