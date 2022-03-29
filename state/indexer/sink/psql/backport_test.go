package psql

import (
	"github.com/yenkuanlee/tendermint/state/indexer"
	"github.com/yenkuanlee/tendermint/state/txindex"
)

var (
	_ indexer.BlockIndexer = BackportBlockIndexer{}
	_ txindex.TxIndexer    = BackportTxIndexer{}
)
