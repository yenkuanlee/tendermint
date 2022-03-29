package v0

import (
	"context"

	abciclient "github.com/yenkuanlee/tendermint/abci/client"
	"github.com/yenkuanlee/tendermint/abci/example/kvstore"
	"github.com/yenkuanlee/tendermint/config"
	"github.com/yenkuanlee/tendermint/internal/mempool"
	mempoolv0 "github.com/yenkuanlee/tendermint/internal/mempool/v0"
)

var mp mempool.Mempool

func init() {
	app := kvstore.NewApplication()
	cc := abciclient.NewLocalCreator(app)
	appConnMem, _ := cc()
	err := appConnMem.Start()
	if err != nil {
		panic(err)
	}

	cfg := config.DefaultMempoolConfig()
	cfg.Broadcast = false

	mp = mempoolv0.NewCListMempool(cfg, appConnMem, 0)
}

func Fuzz(data []byte) int {
	err := mp.CheckTx(context.Background(), data, nil, mempool.TxInfo{})
	if err != nil {
		return 0
	}

	return 1
}
