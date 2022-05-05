package filestore

import (
    "fmt"
    "os"
    "io/ioutil"
    "math/big"
)

type fd struct {
    *os.File
    filename string
}

func newFile(filename Path) (File, error) {
    var err error
    result := fd{ filename: string(filename) }
    result.File, err = os.OpenFile(result.filename, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
    if err != nil {
        return nil, err
    }
    return &result, nil
}

func (f fd) Path() Path {
    return Path(f.filename)
}

func (f fd) Size() int64 {
    info, err := os.Stat(f.filename)
    if err != nil {
        return -1
    }
    return info.Size()
}

func (f fd) ToBinary() []byte {
    fmt.Println("dir : ", f.filename)
    fb, err := ioutil.ReadFile(f.filename)
    if err != nil{
		panic(err)
	}
    // fmt.Println("fb", fb)
    return fb
}

func (f fd) ToBigInt() *big.Int {
    ret := new(big.Int)
    fb := f.ToBinary()
    fmt.Println("file binary len : ", len(fb))
    ret.SetBytes(fb)
    // fmt.Println("file bigInt", ret)
    return ret
}
