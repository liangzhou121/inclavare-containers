package sgx // import "github.com/inclavare-containers/rune/libenclave/attestation/sgx"

// RA Type
const (
	UnknownRaType = iota
	EPID
	DCAP
)

const (
	SpidLength = 16
)

const (
	AttestationEpid  = "sgx-epid"
	AttestationEcdsa = "sgx-ecdsa"
)
