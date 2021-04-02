package Cryptoplasm_Cardinality

import (
    aux "Cryptoplasm-Core/Auxiliary"
    b "Cryptoplasm-Core/Cryptoplasm_Blockchain_Constants"
    "fmt"
    p "github.com/Crypt0plasm/Firefly-APD"
    "math/big"
)

//Cryptoplasm Cardinality Package attempts to code the Schoof Algorithm for
//computing Prime Finite Field Defined Elliptic Curve Cardinality

const (
    ThousandK 	= 1000000
    Prime1M	= 15485865
)

var (
    MillionPrimes = MakeMillionPrimes()
    c = b.CryptoplasmPrecisionContext
)

func SieveOfEratosthenes(N int) (primes []int) {
    Truthiness := make([]bool, N)
    for i := 2; i < N; i++ {
	if Truthiness[i] == true {
	    continue
	}
	primes = append(primes, i)
	for k := i * i; k < N; k += i {
	    Truthiness[k] = true
	}
    }
    return
}

//Creates a Slice with the first 1 million Prime numbers as big.Int
func MakeMillionPrimes () [ThousandK]*big.Int {

    var PrimesBigInt [ThousandK]*big.Int
    PrimesInt := SieveOfEratosthenes(Prime1M)
    for i:=0; i<len(PrimesInt); i++ {
	PrimesBigInt[i] = big.NewInt(int64(PrimesInt[i]))
    }
    return PrimesBigInt
}

//Counts how many primes are needed to compute T modulo these primes,
//for the given Elliptic Finite Prime Field P
//We need as many primes as to reach a product between them greater than 4*Sqrt(PrimeField)
func CountPrimes (P *big.Int) uint64 {
    var (
    	Count uint64
    	Product = p.NFS("1")
    )
    PrimeDecimal := p.NFBI(P)
    ProductLimit := aux.FourSqrtQ(PrimeDecimal)
    //PrimeMatrix := MakeMillionPrimes()

    for i:=0; i<len(MillionPrimes); i++ {
        MatrixPrimeDecimal := p.NFBI(MillionPrimes[i])
        Product = b.MULxc(Product,MatrixPrimeDecimal)
        Count = Count + 1
        if b.DecimalGreaterThanOrEqual(Product,ProductLimit) == true {
            break
	}
    }
    return Count
}

func MakeSchoofSet (PrimeNumber *big.Int) []*big.Int {
    NeededElements := CountPrimes(PrimeNumber)
    SmallArray := MillionPrimes[0:NeededElements]
    //SmallArray := make([]*big.Int,NeededElements)
    //copy(LargeArray[0:NeededElements], SmallArray[:])
    return SmallArray
}

func ComputeEllipseTrace(Prime, W1, W2 *big.Int) *big.Int {
    var(
	AbsoluteChineseTrace = new(p.Decimal)
	FinalTrace = new(big.Int)
	PrimesSetProduct = big.NewInt(1)
    )
    PrimesSet := MakeSchoofSet(Prime)
    ResultSet := make([]*big.Int,len(PrimesSet))

    //Computing the ResultSet; This is the core of the Schoof Algorithm,
    //the most computing Intensive Function
    for i:=0; i<len(PrimesSet); i++ {
        ResultSet[i] = TraceModPrime(PrimesSet[i],Prime,W1,W2)
    }
    ChineseTrace,_ := aux.ChineseRemainderTheorem(ResultSet,PrimesSet)

    //Getting The final Trace by making sure the result of the ChineseTrace fits in absolute value
    //within 2*Sqrt(Prime)
    _,_ = c.Abs(AbsoluteChineseTrace,p.NFBI(ChineseTrace))
    if b.DecimalGreaterThanOrEqual(AbsoluteChineseTrace,aux.TwoSqrtQ(p.NFBI(Prime))) == true {
        FinalTrace = ChineseTrace
    } else {
	for i:=0; i<len(PrimesSet); i++ {
	    PrimesSetProduct.Mul(PrimesSetProduct,PrimesSet[i])
	}
	FinalTrace.Sub(ChineseTrace,PrimesSetProduct)
    }

    return FinalTrace
}

func TraceModPrime (SmallPrime, BigPrime, Coeff1, Coeff2 *big.Int) *big.Int {
    var Result = new(big.Int)

    //If SmallPrime is 2, then the Result is Zero, because all checked Elliptic Curves have Cofactor 4 or 8,
    //and therefore have at least one Point of order 2.
    Cmp := SmallPrime.Cmp(Two)
    if Cmp == 0 {
    	Result = big.NewInt(0)
    } else {
        //Here we need the function that computes Trace mod smallPrime from 3 upward, 3,5,7,11,etc...
        Result = SchoofCore(SmallPrime,BigPrime,Coeff1,Coeff2)
        fmt.Println("Result is unreachable...")
    }

    return Result
}

//The Core of the Schoof Algorithm, Computing Trace mod many smaller Primes.
func SchoofCore (SmallPrime, BigPrime, W1, W2 *big.Int) *big.Int {
    var Result = new(big.Int)
    return Result
}