package Cryptoplasm_Elliptic

import "math/big"

// Structures defining different types of TEC (Twisted-Edwards-Curves) Point Coordinates
//
//
// Affine Coordinates, represent the normal X and Y point coordinates.
// They are designated with aX and aY
//
// Type0 - Affine Coordinates
type TecAffineCoordinates struct {
    AX *big.Int
    AY *big.Int
}
// Extended Coordinates represent aX and aY as eX, eY, eZ and eT satisfying the following equations:
//
// aX = eX/eZ
// aY = eY/eZ
// aX*aY = eT/eZ
// with eZ != 0
//
// Type1 - Extended Coordinates
type TecExtendedCoordinates struct {
    EX *big.Int
    EY *big.Int
    EZ *big.Int
    ET *big.Int
}
//
// Inverted Coordinates represent aX and aY as iX, iY and iZ satisfying the following equations:
//
// aX = iZ/iX
// aY = iZ/iY
//
// Type2 - Inverted Coordinates
type TecInvertedCoordinates struct {
    IX *big.Int
    IY *big.Int
    IZ *big.Int
}
//
// Projective Coordinates represent aX and aY as pX, pY and pZ satisfying the following equations:
//
// aX = iX/iZ
// aY = iY/iZ
// with eZ != 0
//
// Type3 - Projective Coordinates
type TecProjectiveCoordinates struct {
    PX *big.Int
    PY *big.Int
    PZ *big.Int
}

// Point Conversions From Affine Coordinates
//
// Type 1: Affine --> Extended
func TecAffine2TecExtended (point TecAffineCoordinates, prime *big.Int) TecExtendedCoordinates {
    var (
    	result TecExtendedCoordinates
    )

    result.EX = point.AX
    result.EY = point.AY
    result.ET = MulMod(point.AX,point.AY,prime)
    result.EZ = One
    return result
}
//
// Type 2: Affine --> Inverted
// For this conversion the Prime number is required
func TecAffine2TecInverted (point TecAffineCoordinates, prime *big.Int) TecInvertedCoordinates {
    var (
        result TecInvertedCoordinates
        mmix = new(big.Int)
        mmiy = new(big.Int)
    )
    mmix.ModInverse(point.AX,prime)
    mmiy.ModInverse(point.AY,prime)

    result.IX = mmix
    result.IY = mmiy
    result.IZ = One
    return result
}
//
// Type 3: Affine --> Projective
func TecAffine2TecProjective (point TecAffineCoordinates) TecProjectiveCoordinates {
    var result TecProjectiveCoordinates

    result.PX = point.AX
    result.PY = point.AY
    result.PZ = One
    return result
}

// Point Conversions To Affine Coordinates
//
// These conversion are based on the modulo prime number the Elliptic is defined upon
//
// Type 1: Extended to Affine
func TecExtended2TecAffine (point TecExtendedCoordinates, prime *big.Int) TecAffineCoordinates {
    var (
        result TecAffineCoordinates
        mmi = new(big.Int)
    )

    //mmi = ModularMultiplicativeInverse(point.EZ,prime)
    mmi.ModInverse(point.EZ,prime)

    result.AX = MulMod(point.EX,mmi,prime)
    result.AY = MulMod(point.EY,mmi,prime)

    return result
}
//
// Type 2: Inverted to Affine
//
func TecInverted2TecAffine (point TecInvertedCoordinates, prime *big.Int) TecAffineCoordinates {
    var (
        result TecAffineCoordinates
        mmix = new(big.Int)
        mmiy = new(big.Int)
    )

    mmix.ModInverse(point.IX,prime)
    mmiy.ModInverse(point.IY,prime)

    result.AX = MulMod(point.IZ,mmix,prime)
    result.AY = MulMod(point.IZ,mmiy,prime)

    return result
}
//
// Type 3: Projective to Affine
//
func TecProjective2TecAffine (point TecProjectiveCoordinates, prime *big.Int) TecAffineCoordinates {
    var (
        result TecAffineCoordinates
        mmi = new(big.Int)
    )

    //mmi = ModularMultiplicativeInverse(point.EZ,prime)
    mmi.ModInverse(point.PZ,prime)

    result.AX = MulMod(point.PX,mmi,prime)
    result.AY = MulMod(point.PY,mmi,prime)

    return result
}

func ModularMultiplicativeInverse(a,b *big.Int) *big.Int {
    //function assumes b is prime, which means gcd(a,b) is 1
    //This means a ModularMultiplicativeInverse Exists
    //Therefore condition GCD!=1 is not tested
    var (
        result = new(big.Int)
        s =new(big.Int)
    )
    s.Sub(b, Two)
    result.Exp(a,s,b)
    return result
}

