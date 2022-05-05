/*
	https://docs.gnark.consensys.net/en/latest/ 
	go snark 확인해보기


	https://cs.opensource.google/go/go/+/refs/tags/go1.18.1:src/crypto/rsa/rsa.go;l=261;drc=cf697253abb781e8a3e8825b7a4b5b96a534b907;bpv=1;bpt=1
	https://pkg.go.dev/crypto/internal/boring 
*/


package crypto

import(
	"fmt"
	"math"
	"math/big"
)


type PublicKey struct {
	N *big.Int // modulus
	E int      // public exponent
}

type PrivateKey struct {
	PublicKey            // public part.
	D         *big.Int   // private exponent
	Primes    []*big.Int // prime factors of N, has >= 2 elements.
}
