/*
Package server is used to start a new ABCI server.

It contains two server implementation:
 * gRPC server
 * socket server

*/
package server

import (
	"fmt"

	"github.com/yenkuanlee/tendermint/abci/types"
	"github.com/yenkuanlee/tendermint/libs/log"
	"github.com/yenkuanlee/tendermint/libs/service"
)

func NewServer(logger log.Logger, protoAddr, transport string, app types.Application) (service.Service, error) {
	var s service.Service
	var err error
	switch transport {
	case "socket":
		s = NewSocketServer(logger, protoAddr, app)
	case "grpc":
		s = NewGRPCServer(logger, protoAddr, types.NewGRPCApplication(app))
	default:
		err = fmt.Errorf("unknown server type %s", transport)
	}
	return s, err
}
