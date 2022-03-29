package types

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/yenkuanlee/tendermint/config"
	"github.com/yenkuanlee/tendermint/crypto/tmhash"
	"github.com/yenkuanlee/tendermint/internal/test/factory"
	tmrand "github.com/yenkuanlee/tendermint/libs/rand"
	tmtime "github.com/yenkuanlee/tendermint/libs/time"
	tmproto "github.com/yenkuanlee/tendermint/proto/tendermint/types"
	"github.com/yenkuanlee/tendermint/types"
)

func TestPeerCatchupRounds(t *testing.T) {
	cfg, err := config.ResetTestRoot(t.TempDir(), "consensus_height_vote_set_test")
	if err != nil {
		t.Fatal(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	valSet, privVals := factory.ValidatorSet(ctx, t, 10, 1)

	chainID := cfg.ChainID()
	hvs := NewHeightVoteSet(chainID, 1, valSet)

	vote999_0 := makeVoteHR(ctx, t, 1, 0, 999, privVals, chainID)
	added, err := hvs.AddVote(vote999_0, "peer1")
	if !added || err != nil {
		t.Error("Expected to successfully add vote from peer", added, err)
	}

	vote1000_0 := makeVoteHR(ctx, t, 1, 0, 1000, privVals, chainID)
	added, err = hvs.AddVote(vote1000_0, "peer1")
	if !added || err != nil {
		t.Error("Expected to successfully add vote from peer", added, err)
	}

	vote1001_0 := makeVoteHR(ctx, t, 1, 0, 1001, privVals, chainID)
	added, err = hvs.AddVote(vote1001_0, "peer1")
	if err != ErrGotVoteFromUnwantedRound {
		t.Errorf("expected GotVoteFromUnwantedRoundError, but got %v", err)
	}
	if added {
		t.Error("Expected to *not* add vote from peer, too many catchup rounds.")
	}

	added, err = hvs.AddVote(vote1001_0, "peer2")
	if !added || err != nil {
		t.Error("Expected to successfully add vote from another peer")
	}

}

func makeVoteHR(
	ctx context.Context,
	t *testing.T,
	height int64,
	valIndex, round int32,
	privVals []types.PrivValidator,
	chainID string,
) *types.Vote {
	t.Helper()

	privVal := privVals[valIndex]
	pubKey, err := privVal.GetPubKey(ctx)
	require.NoError(t, err)

	randBytes := tmrand.Bytes(tmhash.Size)

	vote := &types.Vote{
		ValidatorAddress: pubKey.Address(),
		ValidatorIndex:   valIndex,
		Height:           height,
		Round:            round,
		Timestamp:        tmtime.Now(),
		Type:             tmproto.PrecommitType,
		BlockID:          types.BlockID{Hash: randBytes, PartSetHeader: types.PartSetHeader{}},
	}

	v := vote.ToProto()
	err = privVal.SignVote(ctx, chainID, v)
	require.NoError(t, err, "Error signing vote")

	vote.Signature = v.Signature

	return vote
}
