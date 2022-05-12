package main

import(
	// "os"
	"fmt"
	"time"
	"math/big"
	c "github.com/kyoungbinkim/seal/crypto"
	// "github.com/kyoungbinkim/seal/seal"
	// filestore "github.com/kyoungbinkim/seal/filestore"
)


func main(){
	// storePath := "./"
	// filePath := filestore.Path("../lotus/filecoin-payload.bin")

	

	//32768
	//131072+1
	key, _ := c.KeyGen(131072+1)

	// fmt.Println("N : ", key.N)
	// fmt.Println("N : ", key.E)
	// fmt.Println("N : ", key.D)

	msg := big.NewInt(14)
	Enc := new(big.Int)
	Dec := new(big.Int)

	tic := time.Now()
	Enc.Exp(msg, key.E, key.N)
	// fmt.Println("Enc : ", Enc)
	fmt.Println("Enc time : ", time.Since(tic))

	tic = time.Now()	
	Dec.Exp(Enc, key.D, key.N)
	fmt.Println("Dec : ", Dec)
	fmt.Println("Dec time : ", time.Since(tic))


	// sk := c.LoadSk()
	// pk := sk.PublicKey
	// fmt.Println("sk D : ", sk.D)
	// fmt.Println("pk E : ", pk.E)
	
	// f := seal.Seal(storePath, filePath, key)
	// fmt.Println("file size : ", len(f))
}