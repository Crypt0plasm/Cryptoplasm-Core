package Auxiliary

import (
    b "Cryptoplasm-Core/Cryptoplasm_Blockchain_Constants"
    p "github.com/Crypt0plasm/Firefly-APD"
    "math/big"
    "unicode/utf8"
)

func SqrtP1 (Number *p.Decimal) *p.Decimal {
    var(
	One = p.NFI(1)
	Sqrt = new(p.Decimal)
	c = b.CryptoplasmPrecisionContext
    )
    _, _ = c.Sqrt(Sqrt, Number)
    Result := b.ADDxc(Sqrt,One)
    return Result
}

func ReverseBigIntSlice (s []*big.Int) []*big.Int {
    for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
	s[i], s[j] = s[j], s[i]
    }
    return s
}

func ListDivisors(Number *big.Int) []*big.Int {
    var(
	Zero        = p.NFI(0)
	One         = p.NFI(1)
	Start 	    = p.NFI(1)
	End 	    = SqrtP1(p.NFBI(Number))
	Slice1      []*big.Int
	Slice2      []*big.Int
	Divisors    []*big.Int
    )

    for i:=new(p.Decimal).Set(Start); b.DecimalLessThan(i,End) == true; i=b.ADDxc(i,One) {
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

func TrimFirstRune (s string) string {
    _, i := utf8.DecodeRuneInString(s)
    return s[i:]
}