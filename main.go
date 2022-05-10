package main

import(
	// "os"
	"fmt"
	c "github.com/kyoungbinkim/seal/crypto"
	"github.com/kyoungbinkim/seal/seal"
	filestore "github.com/kyoungbinkim/seal/filestore"
)


func main(){
	storePath := "./"
	filePath := filestore.Path("../lotus/filecoin-payload.bin")

	fmt.Println("hello")
	f := seal.Seal(storePath, filePath)

	fmt.Println("file size : ", len(f))

	//32768
	c.KeyGen(65536)
	sk := c.LoadSk(65536)
	pk := sk.PublicKey
	fmt.Println("sk D : ", sk.D)
	fmt.Println("pk E : ", pk.E)
	// b := []byte()
	// n,err := f.Read(b)
	// fmt.Println("b",b,n,err);

	// fmt.Println("Chdir : ", f.Chdir())
}