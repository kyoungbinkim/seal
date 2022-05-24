package seal

import (
	"fmt"
	// "io/ioutil"
	// "path/filepath"
	// "math"
	"math/big" 

	// "encoding/hex"
	// "crypto/rsa"

	// filestore "github.com/kyoungbinkim/seal/filestore"
	c "github.com/kyoungbinkim/seal/crypto"
)



// func Seal(datafs,keyfs filestore.FileStore, dataPath filestore.Path) *big.Int {
// 	ret := new(big.Int)
// 	File,err := datafs.Open(dataPath)
// 	key := c.LoadKey(keyfs)
// 	return ret
// }

func ChunkSeal(data *big.Int, key *c.Key) *big.Int {
	ret := new(big.Int)
	ret.Exp(data, key.E, key.N)
	return ret
}

func ChunkDec(replica *big.Int, key *c.Key) *big.Int{
	fmt.Println("decoding key : ", key.D)
	ret := new(big.Int)
	ret.Exp(replica, key.D, key.N)
	return ret
}