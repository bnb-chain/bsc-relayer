package common

const (
	chainIDLength              uint64 = 32
	heightLength               uint64 = 8
	validatorSetHashLength     uint64 = 32
	validatorPubkeyLength      uint64 = 32
	validatorVotingPowerLength uint64 = 8
	appHashLength              uint64 = 32
	maxConsensusStateLength    uint64 = 32 * (128 - 1) // maximum validator quantity 99
)
