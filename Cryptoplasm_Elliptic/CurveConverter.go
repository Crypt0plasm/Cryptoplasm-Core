package Cryptoplasm_Elliptic

import (
    "math/big"
)

//Basic Modulus Operations
//Addition Modulo prime
func AddModulus (prime, a,b *big.Int) *big.Int{
    var result = new(big.Int)
    result.Add(a,b)
    result.Mod(result,prime)
    return result
}
//Subtraction Modulo prime
func SubModulus (prime, a,b *big.Int) *big.Int{
    var result = new(big.Int)
    result.Sub(a,b)
    result.Mod(result,prime)
    return result
}
//Multiplication Modulo prime
func MulModulus (prime, a,b *big.Int) *big.Int{
    var result = new(big.Int)
    result.Mul(a,b)
    result.Mod(result,prime)
    return result
}
//Division Modulo prime
func QuoModulus (prime, a,b *big.Int) *big.Int{
    var (
    	result = new(big.Int)
	mmi = new(big.Int)
    )
    mmi.ModInverse(b,prime)
    result = MulModulus(prime,a,mmi)
    return result
}

func TEC2Montgomery (prime, E, D *big.Int) (MB, MA *big.Int) {
    v1 := AddModulus(prime,E,D)
    v2 := MulModulus(prime,Two,v1)
    v3 := SubModulus(prime,E,D)
    MA = QuoModulus(prime,v2,v3)
    MB = QuoModulus(prime,Four,v3)
    //fmt.Printf("TWISTED EDUARDS: %v*x^2 + y^2 = 1 + %v*x^2*y^2 (mod %v)\n", E,D,prime)
    //fmt.Printf("     MONTGOMERY: %v*y^2 = x^3 + %v*x^2 + x (mod %v)\n", MB,MA,prime)
    return MB, MA
}

func Montgomery2ShortWeierstrass (prime, B, A *big.Int) (W1, W2 *big.Int) {
    var(
        v1 = new(big.Int)
	v2 = new(big.Int)
	v3 = new(big.Int)
	v4 = new(big.Int)
	Nine = big.NewInt(9)
	TwentySeven = big.NewInt(27)
    )
    v1.Exp(A,Two,prime)
    v2.Exp(A,Three,prime)
    v3.Exp(B,Two,prime)
    v4.Exp(B,Three,prime)

    v5 := SubModulus(prime,Three,v1)
    v6 := MulModulus(prime,Three,v3)
    W1 = QuoModulus(prime,v5,v6)

    v7 := MulModulus(prime,Two,v2)
    v8 := MulModulus(prime,Nine,A)
    v9 := SubModulus(prime,v7,v8)
    v10 := MulModulus(prime,TwentySeven,v4)
    W2 = QuoModulus(prime,v9,v10)

    //fmt.Printf("     MONTGOMERY: %v*y^2 = x^3 + %v*x^2 + x (mod %v)\n", B,A,prime)
    //fmt.Printf("  s.WEIERSTRASS: y^2 = x^3 + %v*x + %v (mod %v)\n", W1,W2,prime)
    return W1, W2
}

func DIShortWeierstrass(prime, W1,W2 *big.Int) (Delta, J *big.Int) {
    Delta,J = DILongWeierstrass(prime,Zero,Zero,Zero,W1,W2)
    return Delta,J

}

func DILongWeierstrass(prime, a1,a2,a3,a4,a6 *big.Int) (Delta, J *big.Int) {
    var(
	v1 = new(big.Int)
	v5 = new(big.Int)
	v12 = new(big.Int)
	v13 = new(big.Int)
	v15 = new(big.Int)
	v17 = new(big.Int)
	v24 = new(big.Int)
	Eight = big.NewInt(8)
	TwentySeven = big.NewInt(27)
	TwentyFour = big.NewInt(24)
	Nine = big.NewInt(9)
    )


    v1.Exp(a1,Two,prime)
    v2 := MulModulus(prime,Four,a2)
    b2 := AddModulus(prime,v1,v2)

    v3 := MulModulus(prime,a1,a3)
    v4 := MulModulus(prime,Two,a4)
    b4 := AddModulus(prime,v3,v4)

    v5.Exp(a3,Two,prime)
    v6 := MulModulus(prime,Four,a6)
    b6 := AddModulus(prime,v5,v6)

    v7 := MulModulus(prime,v1,a6)
    v8 := MulModulus(prime,v3,a4)
    v9 := MulModulus(prime,a2,a6)
    v10 := MulModulus(prime,Four,v9)

    v11 := MulModulus(prime,v5,a2)
    v12.Exp(a4,Two,prime)

    b8 := SubModulus(prime,v7,v8)
    b8 = AddModulus(prime,b8,v10)
    b8 = AddModulus(prime,b8,v11)
    b8 = SubModulus(prime,b8,v12)

    v13.Exp(b2,Two,prime)
    v14 := MulModulus(prime,v13,b8)

    v15.Exp(b4,Three,prime)
    v16 := MulModulus(prime,Eight,v15)

    v17.Exp(b6,Two,prime)
    v18 := MulModulus(prime,TwentySeven,v17)

    v19 := MulModulus(prime,Nine,b2)
    v20 := MulModulus(prime,b4,b6)
    v21 := MulModulus(prime,v19,v20)

    Delta = SubModulus(prime,v21,v18)
    Delta = SubModulus(prime,Delta,v16)
    Delta = SubModulus(prime,Delta,v14)

    v22 := MulModulus(prime,TwentyFour,b4)
    v23 := SubModulus(prime,v13,v22)

    v24.Exp(v23,Three,prime)
    J = QuoModulus(prime,v24,Delta)

    return Delta, J

}