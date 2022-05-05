package main

import(
	// "os"
	"fmt"

	"github.com/kyoungbinkim/seal"
	filestore "github.com/filecoin-project/go-fil-filestore"
)


func main(){
	storePath := "./"
	filePath := filestore.Path("../lotus/filecoin-payload.bin")

	fmt.Println("hello")
	f := seal.Seal(storePath, filePath)

	fmt.Println("file size : ", len(f))
	// b := []byte()
	// n,err := f.Read(b)
	// fmt.Println("b",b,n,err);

	// fmt.Println("Chdir : ", f.Chdir())
}