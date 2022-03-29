package factory

import (
	"encoding/hex"
	"strings"

	"github.com/yenkuanlee/tendermint/libs/rand"
	"github.com/yenkuanlee/tendermint/types"
)

// NodeID returns a valid NodeID based on an inputted string
func NodeID(str string) types.NodeID {
	id, err := types.NewNodeID(strings.Repeat(str, 2*types.NodeIDByteLength))
	if err != nil {
		panic(err)
	}
	return id
}

// RandomNodeID returns a randomly generated valid NodeID
func RandomNodeID() types.NodeID {
	id, err := types.NewNodeID(hex.EncodeToString(rand.Bytes(types.NodeIDByteLength)))
	if err != nil {
		panic(err)
	}
	return id
}
