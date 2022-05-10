package seal

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"math"
	"math/big" 

	"encoding/hex"
	// "crypto/rsa"

	filestore "github.com/kyoungbinkim/seal/filestore"
)

// type piece struct{
// 	sectorNum	sectorNum
// 	pieceNum	pieceNum
// 	data		[]byte
// }

func byteToBigInt(data []byte) *big.Int{
	ret := new(big.Int)
	ret.SetBytes(data)
	return ret
}

// func seal_core(d big.Int, pk big.Int, g big.Int){
// }

func ReadFile(path string, ) ([]byte, string) {
	f , err := ioutil.ReadFile(path)
	if err != nil{
		panic(err)
	}
	p, _ := filepath.Abs(path)
	
	return f, p
}


func Seal(storePath string, filePath filestore.Path) []byte {
	p, _ := filepath.Abs(string(storePath))
	fmt.Println("Seal path : ", p)
	fs, err := filestore.NewLocalFileStore(storePath)
	if err != nil{
		panic(err)
	}

	f , err := fs.Open(filePath)
	if err != nil{
		panic(err)
	}
	fmt.Println("f Path : ", f.Path())

	f.ToBigInt()
	var ret []byte

	fmt.Println(f)

	SectorBitSize := math.Pow(2,16)
	pieceSize := int( math.Pow(2,14) )// 16KiB
	pieceNum := len(f.ToBinary()) / int(pieceSize)
	fmt.Println(f.ToBinary(), hex.EncodeToString(f.ToBinary()))
	fmt.Println("file : ",len(f.ToBinary()),"bytes")
	fmt.Println("Bit Size : ", SectorBitSize)
	fmt.Println("pieceSize(2^14) : ", int(pieceSize) * int(pieceNum))
	
	// d := f.ToBinary()
	// // var seala []*big.Int
	// for i := 0; i < pieceNum; i++ {
	// 	p := d[i*pieceSize : (i+1) * pieceSize]
	// 	fmt.Println(i, " piceSize : ", len(p))
	// 	pieceBigInt := byteToBigInt(p)
	// 	fmt.Println(i, " big Int :", pieceBigInt)
		
	// 	// go func(a big.Int, b)
	// }

	return ret
}


func Sealing(sector filestore.File, fs filestore.filestore, pk *rsa.PublicKey){
	sectorBytes := sector.ToBinary()
}