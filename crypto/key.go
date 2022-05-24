/*
	https://docs.gnark.consensys.net/en/latest/ 
	go snark 확인해보기


	https://cs.opensource.google/go/go/+/refs/tags/go1.18.1:src/crypto/rsa/rsa.go;l=261;drc=cf697253abb781e8a3e8825b7a4b5b96a534b907;bpv=1;bpt=1
	https://pkg.go.dev/crypto/internal/boring 
*/


package c

import(

	"fmt"
	"io"
	"os"
	"encoding/gob"
	"math"
	"math/big"
	"crypto/rsa"
	"crypto/rand"
	"time"
	"sync"
	"errors"

	fs "github.com/kyoungbinkim/seal/filestore"
)


// type PublicKey struct {
// 	N *big.Int // modulus
// 	E int      // public exponent
// }

// type PrivateKey struct {
// 	PublicKey            // public part.
// 	D         *big.Int   // private exponent
// 	Primes    []*big.Int // prime factors of N, has >= 2 elements.
// }

type Key struct{
	N *big.Int
	E *big.Int
	D *big.Int
}

var bigZero = big.NewInt(0)
var bigOne = big.NewInt(1)

var (
	closedChanOnce sync.Once
	closedChan     chan struct{}
)

func KeyGenX(bits int) (*Key, error){
	tic := time.Now()
	// sk, err := rsa.GenerateKey(rand.Reader, bits) //65536
	key := new(Key)
	key.D = big.NewInt(3)
	key.E = new(big.Int)
	random := rand.Reader
	// tmp := new(big.Int)
	order := new(big.Int)
	MaybeReadByte(random)
	fmt.Println("D : ", key.D)
	for {
		s := time.Now()
		key.N, _ = rand.Prime(random, bits)
		fmt.Println("pick Prime time : ", time.Since(s))
		order.Sub(key.N, bigOne)
		if  key.E.ModInverse(key.D, order) != nil {			
			break
		}
		
	}

	toc := time.Since(tic)
	fmt.Println("keyGen takes : ", toc)
	
	fmt.Println("N : ", key.N.BitLen())
	fmt.Println("d : ", key.D.BitLen())
	// fmt.Println("primes", sk.Primes[0].BitLen(), sk.Primes[1].BitLen())
	fmt.Println("e : ", key.E.BitLen())

	keyFile, err := os.Create(fmt.Sprint(bits) + "bits.key")
	if err != nil{
		panic(err)
		os.Exit(1)
	}

	keyEncoder := gob.NewEncoder(keyFile)
	keyEncoder.Encode(*key)
	keyFile.Close()

	return key, err
}

func LoadSk(bits int) (*rsa.PrivateKey) {
	keyFile,err := os.Open(fmt.Sprint(bits)+"bits_private.key")
	if err != nil{
		os.Exit(1)
		panic(err)
	}
	var sk rsa.PrivateKey
	skDecoder := gob.NewDecoder(keyFile)
	skDecoder.Decode(&sk)
	
	fmt.Println(sk.D)

	return &sk
}

func KeyGen(f fs.FileStore, keyPath fs.Path) *Key {
	keyFile,err := f.Open(keyPath)
	if err != nil {
		panic(err)
	}
	k := Key{
		keyFile.ToBigInt(),
		big.NewInt(1),
		big.NewInt(1),
	}

	order :=big.NewInt(0)
	order.Sub(k.N, bigOne)
	bigTwo := big.NewInt(2)
	fmt.Println("order : ", order)
	for {
		
		for{
			k.D.Add(k.D, bigTwo)
			if k.D.ProbablyPrime(0) {
				break
			}
		}

		k.E.ModInverse(k.D, order)
		
		if k.E.Cmp(bigOne) != 0 {
			fmt.Println("k.E : ", k.E)
			break
		}
	}	
	fmt.Println("k.D : " , k.D)
	fmt.Println(fs.Path("./key.bin"))
	// ffs, err :=fs.FileStore.Create()
	myKeyFile, err := f.Create("./key.bin")
	if err != nil {
		panic(err)
	}

	keyEncoder := gob.NewEncoder(myKeyFile)
	keyEncoder.Encode(k)
	myKeyFile.Close()
	

	return &k;
}

func LoadKey(f fs.FileStore) *Key {
	fmt.Println("====== load Key ======")
	keyFile,err := f.Open(fs.Path("key.bin"))
	if err != nil{
		os.Exit(1)
		panic(err)
	}
	var k Key
	keyDecoder := gob.NewDecoder(keyFile)
	keyDecoder.Decode(&k)
	
	fmt.Println(k.D)

	return &k
}

// MaybeReadByte reads a single byte from r with ~50% probability. This is used
// to ensure that callers do not depend on non-guaranteed behaviour, e.g.
// assuming that rsa.GenerateKey is deterministic w.r.t. a given random stream.
//
// This does not affect tests that pass a stream of fixed bytes as the random
// source (e.g. a zeroReader).
func MaybeReadByte(r io.Reader) {
	closedChanOnce.Do(func() {
		closedChan = make(chan struct{})
		close(closedChan)
	})

	select {
	case <-closedChan:
		return
	case <-closedChan:
		var buf [1]byte
		r.Read(buf[:])
	}
}

func GenerateMultiPrimeKey(random io.Reader, nprimes int, bits int) (*rsa.PrivateKey, error) {
	MaybeReadByte(random)

	priv := new(rsa.PrivateKey)
	priv.E = 3

	if nprimes < 2 {
		return nil, errors.New("crypto/rsa: GenerateMultiPrimeKey: nprimes must be >= 2")
	}

	if bits < 64 {
		primeLimit := float64(uint64(1) << uint(bits/nprimes))
		// pi approximates the number of primes less than primeLimit
		pi := primeLimit / (math.Log(primeLimit) - 1)
		// Generated primes start with 11 (in binary) so we can only
		// use a quarter of them.
		pi /= 4
		// Use a factor of two to ensure that key generation terminates
		// in a reasonable amount of time.
		pi /= 2
		if pi <= float64(nprimes) {
			return nil, errors.New("crypto/rsa: too few primes of given length to generate an RSA key")
		}
	}

	primes := make([]*big.Int, nprimes)

NextSetOfPrimes:
	for {
		todo := bits
		// crypto/rand should set the top two bits in each prime.
		// Thus each prime has the form
		//   p_i = 2^bitlen(p_i) × 0.11... (in base 2).
		// And the product is:
		//   P = 2^todo × α
		// where α is the product of nprimes numbers of the form 0.11...
		//
		// If α < 1/2 (which can happen for nprimes > 2), we need to
		// shift todo to compensate for lost bits: the mean value of 0.11...
		// is 7/8, so todo + shift - nprimes * log2(7/8) ~= bits - 1/2
		// will give good results.
		if nprimes >= 7 {
			todo += (nprimes - 2) / 5
		}
		for i := 0; i < nprimes; i++ {
			var err error
			primes[i], err = rand.Prime(random, todo/(nprimes-i))
			if err != nil {
				return nil, err
			}
			todo -= primes[i].BitLen()
		}

		// Make sure that primes is pairwise unequal.
		for i, prime := range primes {
			for j := 0; j < i; j++ {
				if prime.Cmp(primes[j]) == 0 {
					continue NextSetOfPrimes
				}
			}
		}

		n := new(big.Int).Set(bigOne)
		totient := new(big.Int).Set(bigOne)
		pminus1 := new(big.Int)
		for _, prime := range primes {
			n.Mul(n, prime)
			pminus1.Sub(prime, bigOne)
			totient.Mul(totient, pminus1)
		}
		if n.BitLen() != bits {
			// This should never happen for nprimes == 2 because
			// crypto/rand should set the top two bits in each prime.
			// For nprimes > 2 we hope it does not happen often.
			continue NextSetOfPrimes
		}

		priv.D = new(big.Int)
		e := big.NewInt(int64(priv.E))
		ok := priv.D.ModInverse(e, totient)

		if ok != nil {
			priv.Primes = primes
			priv.N = n
			break
		}
	}

	priv.Precompute()
	return priv, nil
}
