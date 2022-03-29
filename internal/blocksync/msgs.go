package blocksync

import (
	bcproto "github.com/yenkuanlee/tendermint/proto/tendermint/blocksync"
	"github.com/yenkuanlee/tendermint/types"
)

const (
	MaxMsgSize = types.MaxBlockSizeBytes +
		bcproto.BlockResponseMessagePrefixSize +
		bcproto.BlockResponseMessageFieldKeySize
)
