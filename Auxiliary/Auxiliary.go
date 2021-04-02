package Auxiliary

import (
    b "Cryptoplasm-Core/Cryptoplasm_Blockchain_Constants"
    "fmt"
    p "github.com/Crypt0plasm/Firefly-APD"
    "math/big"
    "unicode/utf8"
)

var (
    IOne = big.NewInt(1)
    DOne = p.NFI(1)
    DTwo = p.NFI(2)
    DFour = p.NFI(4)
    c = b.CryptoplasmPrecisionContext
)

// Computes 2*Sqrt(x)
func TwoSqrtQ (Number *p.Decimal) *p.Decimal {
    var(
	Sqrt = new(p.Decimal)
    )
    _, _ = c.Sqrt(Sqrt, Number)
    Result := b.MULxc(Sqrt,DTwo)
    return Result
}

// Computes 4*Sqrt(x)
func FourSqrtQ (Number *p.Decimal) *p.Decimal {
    var(
	Sqrt = new(p.Decimal)
    )
    _, _ = c.Sqrt(Sqrt, Number)
    Result := b.MULxc(Sqrt,DFour)
    return Result
}

// Computes Sqrt(x)+1
func SqrtP1 (Number *p.Decimal) *p.Decimal {
    var(
	Sqrt = new(p.Decimal)
	c = b.CryptoplasmPrecisionContext
    )
    _, _ = c.Sqrt(Sqrt, Number)
    Result := b.ADDxc(Sqrt,DOne)
    return Result
}

func ReverseBigIntSlice (s []*big.Int) []*big.Int {
    for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
	s[i], s[j] = s[j], s[i]
    }
    return s
}

func ReverseStringSlice (s []string) []string {
    for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
	s[i], s[j] = s[j], s[i]
    }
    return s
}

func ListDivisors(Number *big.Int) []*big.Int {
    var(
	Zero        = p.NFI(0)
	Start 	    = p.NFI(1)
	End 	    = SqrtP1(p.NFBI(Number))
	Slice1      []*big.Int
	Slice2      []*big.Int
	Divisors    []*big.Int
    )

    for i:=new(p.Decimal).Set(Start); b.DecimalLessThan(i,End) == true; i=b.ADDxc(i,DOne) {
	if b.DecimalEqual(b.DivMod(p.NFBI(Number),i),Zero) == true {
	    if b.DecimalEqual(b.DivInt(p.NFBI(Number),i),i) == true {
		Slice1 = append(Slice1, &(i.Coeff))
	    } else {
		Slice1 = append(Slice1, &(i.Coeff))
		d := b.DivInt(p.NFBI(Number),i)
		Slice2 = append(Slice2, &(d.Coeff))
	    }
	}
    }
    Divisors = append(Slice1,ReverseBigIntSlice(Slice2)...)
    return Divisors
}

func ChineseRemainderTheorem (ModuloResults, PrimesSet []*big.Int) (*big.Int, error) {
    FirstPrime := new(big.Int).Set(PrimesSet[0])
    for _, n1 := range PrimesSet[1:] {
	FirstPrime.Mul(FirstPrime, n1)
    }
    var x, q, s, z big.Int
    for i, n1 := range PrimesSet {
	q.Div(FirstPrime, n1)
	z.GCD(nil, &s, n1, &q)
	if z.Cmp(IOne) != 0 {
	    return nil, fmt.Errorf("%d not coprime", n1)
	}
	x.Add(&x, s.Mul(ModuloResults[i], s.Mul(&s, &q)))
    }
    return x.Mod(&x, FirstPrime), nil
}

func TrimFirstRune (s string) string {
    _, i := utf8.DecodeRuneInString(s)
    return s[i:]
}