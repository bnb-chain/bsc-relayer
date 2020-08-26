package common

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/tendermint/tendermint/crypto/ed25519"
	tmtypes "github.com/tendermint/tendermint/types"
)

type CrossChainChannelID int8
type CrossChainID uint16

type ConsensusState struct {
	ChainID             string
	Height              uint64
	AppHash             []byte
	CurValidatorSetHash []byte
	NextValidatorSet    *tmtypes.ValidatorSet
}

// output:
// | chainID   | height   | appHash  | curValidatorSetHash | [{validator pubkey, voting power}] |
// | 32 bytes  | 8 bytes  | 32 bytes | 32 bytes            | [{32 bytes, 8 bytes}]              |
func (cs ConsensusState) EncodeConsensusState() ([]byte, error) {
	validatorSetLength := uint64(len(cs.NextValidatorSet.Validators))
	serializeLength := chainIDLength + heightLength + appHashLength + validatorSetHashLength + validatorSetLength*(validatorPubkeyLength+validatorVotingPowerLength)
	if serializeLength > maxConsensusStateLength {
		return nil, fmt.Errorf("too many validators %d, consensus state bytes should not exceed %d", len(cs.NextValidatorSet.Validators), maxConsensusStateLength)
	}

	encodingBytes := make([]byte, serializeLength)

	pos := uint64(0)
	if uint64(len(cs.ChainID)) > chainIDLength {
		return nil, fmt.Errorf("chainID length should be no more than 32")
	}
	copy(encodingBytes[pos:pos+chainIDLength], cs.ChainID)
	pos += chainIDLength

	binary.BigEndian.PutUint64(encodingBytes[pos:pos+heightLength], uint64(cs.Height))
	pos += heightLength

	copy(encodingBytes[pos:pos+appHashLength], cs.AppHash)
	pos += appHashLength

	copy(encodingBytes[pos:pos+validatorSetHashLength], cs.CurValidatorSetHash)
	pos += validatorSetHashLength

	for index := uint64(0); index < validatorSetLength; index++ {
		validator := cs.NextValidatorSet.Validators[index]
		pubkey, ok := validator.PubKey.(ed25519.PubKeyEd25519)
		if !ok {
			return nil, fmt.Errorf("invalid pubkey type")
		}

		copy(encodingBytes[pos:pos+validatorPubkeyLength], pubkey[:])
		pos += validatorPubkeyLength

		binary.BigEndian.PutUint64(encodingBytes[pos:pos+validatorVotingPowerLength], uint64(validator.VotingPower))
		pos += validatorVotingPowerLength
	}

	return encodingBytes, nil
}

type Header struct {
	tmtypes.SignedHeader
	ValidatorSet     *tmtypes.ValidatorSet `json:"validator_set"`
	NextValidatorSet *tmtypes.ValidatorSet `json:"next_validator_set"`
}

func (h *Header) Validate(chainID string) error {
	if err := h.SignedHeader.ValidateBasic(chainID); err != nil {
		return err
	}
	if h.ValidatorSet == nil {
		return fmt.Errorf("invalid header: validator set is nil")
	}
	if h.NextValidatorSet == nil {
		return fmt.Errorf("invalid header: next validator set is nil")
	}
	if !bytes.Equal(h.ValidatorsHash, h.ValidatorSet.Hash()) {
		return fmt.Errorf("invalid header: validator set does not match hash")
	}
	if !bytes.Equal(h.NextValidatorsHash, h.NextValidatorSet.Hash()) {
		return fmt.Errorf("invalid header: next validator set does not match hash")
	}
	return nil
}

func (h *Header) EncodeHeader() ([]byte, error) {
	bz, err := Cdc.MarshalBinaryLengthPrefixed(h)
	if err != nil {
		return nil, err
	}
	return bz, nil
}
