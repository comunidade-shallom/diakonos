package support

import (
	"crypto/rand"
	"math/big"
)

func RandInt(max int64) int64 {
	nBig, err := rand.Int(rand.Reader, big.NewInt(max))
	if err != nil {
		panic(err)
	}

	return nBig.Int64()
}
