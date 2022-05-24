package util

import (
	"fmt"
	"mat/big"
)

func Byte2BigInt(b []byte) *bit.Int {
	ret := new(big.Int)
	ret.SetBytes(b)

	return ret
}

