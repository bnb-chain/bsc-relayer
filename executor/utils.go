package executor

import (
	"encoding/binary"

	"github.com/bnb-chain/bsc-relayer/common"
)

func buildCrossChainPackageKey(sourceChainID, destinationChainID common.CrossChainID, channelID common.CrossChainChannelID, sequence uint64) []byte {
	key := make([]byte, totalPackageKeyLength)

	copy(key[:prefixLength], prefixForCrossChainPackageKey)
	binary.BigEndian.PutUint16(key[prefixLength:sourceChainIDLength+prefixLength], uint16(sourceChainID))
	binary.BigEndian.PutUint16(key[prefixLength+sourceChainIDLength:prefixLength+sourceChainIDLength+destChainIDLength], uint16(destinationChainID))
	copy(key[prefixLength+sourceChainIDLength+destChainIDLength:], []byte{byte(channelID)})
	binary.BigEndian.PutUint64(key[prefixLength+sourceChainIDLength+destChainIDLength+channelIDLength:], sequence)

	return key
}

func isHeaderNonExistingErr(err error) bool {
	if err != nil && err.Error() == "RPC error -32603 - Internal error: Height must be less than or equal to the current blockchain height" {
		return true
	}
	return false
}

func buildChannelSequenceKey(destCrossChainID common.CrossChainID, channelID common.CrossChainChannelID) []byte {
	key := make([]byte, prefixLength+destChainIDLength+channelIDLength)

	copy(key[:prefixLength], prefixForSequenceKey)
	binary.BigEndian.PutUint16(key[prefixLength:prefixLength+destChainIDLength], uint16(destCrossChainID))
	copy(key[prefixLength+destChainIDLength:], []byte{byte(channelID)})

	return key
}
