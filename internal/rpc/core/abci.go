package core

import (
	"context"

	abci "github.com/yenkuanlee/tendermint/abci/types"
	"github.com/yenkuanlee/tendermint/internal/proxy"
	"github.com/yenkuanlee/tendermint/libs/bytes"
	"github.com/yenkuanlee/tendermint/rpc/coretypes"
)

// ABCIQuery queries the application for some information.
// More: https://docs.tendermint.com/master/rpc/#/ABCI/abci_query
func (env *Environment) ABCIQuery(
	ctx context.Context,
	path string,
	data bytes.HexBytes,
	height int64,
	prove bool,
) (*coretypes.ResultABCIQuery, error) {
	resQuery, err := env.ProxyApp.Query(ctx, abci.RequestQuery{
		Path:   path,
		Data:   data,
		Height: height,
		Prove:  prove,
	})
	if err != nil {
		return nil, err
	}

	return &coretypes.ResultABCIQuery{Response: *resQuery}, nil
}

// ABCIInfo gets some info about the application.
// More: https://docs.tendermint.com/master/rpc/#/ABCI/abci_info
func (env *Environment) ABCIInfo(ctx context.Context) (*coretypes.ResultABCIInfo, error) {
	resInfo, err := env.ProxyApp.Info(ctx, proxy.RequestInfo)
	if err != nil {
		return nil, err
	}

	return &coretypes.ResultABCIInfo{Response: *resInfo}, nil
}
