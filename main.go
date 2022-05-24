package main

import(
	// "os"
	"fmt"
	"time"
	"sync"
	"math/big"
	"runtime"

	c "github.com/kyoungbinkim/seal/crypto"
	"github.com/kyoungbinkim/seal/seal"
	filestore "github.com/kyoungbinkim/seal/filestore"
)


func main(){
	// storePath := "./"
	// filePath := filestore.Path("../lotus/filecoin-payload.bin")

	keyPath := "./crypto/key/"
	fs, err := filestore.NewLocalFileStore(keyPath)
	if err != nil {
		panic(err)
	}

	f,err := fs.Open(filestore.Path("MersennePrime132049.bin"))
	if err != nil {
		panic(err)
	}
	fmt.Println(f.Size())
	fmt.Println(f.ToBinary())
	fmt.Println(f.ToBigInt())
	

	// kPath := filestore.Path("MersennePrime132049.bin")
	// c.KeyGen(fs, kPath)

	k  := c.LoadKey(fs)
	fmt.Println("k.N bitlen : ", k.N.BitLen())
	fmt.Println("k.E bitlen : ", k.E.BitLen())

	runtime.GOMAXPROCS(3)
	var wait sync.WaitGroup
    wait.Add(12)

	datafs, err := filestore.NewLocalFileStore("./")
	data, err := datafs.Open(filestore.Path("16kb-Test.bin"))

	dataBigInt := data.ToBigInt()
	dataBigInt.Add(dataBigInt, big.NewInt(996655))
	fmt.Println("data : ", dataBigInt)
	fmt.Println("data len : ", dataBigInt.BitLen() )

	tic := time.Now()

	for i:=0 ; i<12 ; i++ {
		dataBigInt.Add(dataBigInt, big.NewInt(17))
		go func(data *big.Int) {
			defer wait.Done()
			seal.ChunkSeal(data,k)
		}(dataBigInt)
	}

	// go func(data *big.Int){
	// 	defer wait.Done()
	// 	seal.ChunkSeal(dataBigInt,k)
	// }(dataBigInt)

	// dataBigInt.Add(dataBigInt, big.NewInt(996655))

	// go func(data *big.Int){ 
	// 	defer wait.Done()
		
	// 	seal.ChunkSeal(dataBigInt,k)
	// }(dataBigInt)
	wait.Wait()
	fmt.Println("seal takes ", time.Since(tic))
	
}