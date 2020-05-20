package executor

import (
	"fmt"
	"testing"

	ctypes "github.com/binance-chain/go-sdk/common/types"
	"github.com/stretchr/testify/require"
	rpcclient "github.com/tendermint/tendermint/rpc/client"
)

func TestBBCExecutor_GetSequence(t *testing.T) {
	bbcExecutor, err := NewBBCExecutor(cfg, ctypes.TmpTestNetwork)
	require.NoError(t, err)

	rpc  := rpcclient.NewHTTP(BBCRpc, "/websocket")
	require.NoError(t, err)
	abciInfo, err := rpc.ABCIInfo()
	require.NoError(t, err)

	height := abciInfo.Response.LastBlockHeight - 1

	nextBindSequence, err := bbcExecutor.GetNextSequence(BindChannelID, height)
	require.NoError(t, err)
	t.Log(fmt.Sprintf("channel %d, sequence %d", BindChannelID, nextBindSequence))

	nextTransferSequence, err := bbcExecutor.GetNextSequence(TransferChannelID, height)
	require.NoError(t, err)
	t.Log(fmt.Sprintf("channel %d, sequence %d", TransferChannelID, nextTransferSequence))

	nextRefundSequence, err := bbcExecutor.GetNextSequence(RefundChannelID, height)
	require.NoError(t, err)
	t.Log(fmt.Sprintf("channel %d, sequence %d", RefundChannelID, nextRefundSequence))

	nextStakingSequence, err := bbcExecutor.GetNextSequence(StakingChannelID, height)
	require.NoError(t, err)
	t.Log(fmt.Sprintf("channel %d, sequence %d", StakingChannelID, nextStakingSequence))

	nextGovSequence, err := bbcExecutor.GetNextSequence(GovChannelID, height)
	require.NoError(t, err)
	t.Log(fmt.Sprintf("channel %d, sequence %d", GovChannelID, nextGovSequence))
}