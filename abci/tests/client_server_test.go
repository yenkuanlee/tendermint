package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"

	abciclientent "github.com/yenkuanlee/tendermint/abci/client"
	"github.com/yenkuanlee/tendermint/abci/example/kvstore"
	abciserver "github.com/yenkuanlee/tendermint/abci/server"
)

func TestClientServerNoAddrPrefix(t *testing.T) {
	addr := "localhost:26658"
	transport := "socket"
	app := kvstore.NewApplication()

	server, err := abciserver.NewServer(addr, transport, app)
	assert.NoError(t, err, "expected no error on NewServer")
	err = server.Start()
	assert.NoError(t, err, "expected no error on server.Start")

	client, err := abciclientent.NewClient(addr, transport, true)
	assert.NoError(t, err, "expected no error on NewClient")
	err = client.Start()
	assert.NoError(t, err, "expected no error on client.Start")
}
