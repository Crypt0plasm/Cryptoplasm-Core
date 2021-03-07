package Cryptoplasm_Elliptic

import "math/big"

// Structures defining different types of TEC (Twisted-Edwards-Curves) Point Coordinates
//
//
// Affine Coordinates, represent the normal X and Y point coordinates.
// They are designated with aX and aY
//
// Type0 - Affine Coordinates
type AffineCoordinates struct {
    AX *big.Int
    AY *big.Int
}
// XZ Coordinates represents x as X and Z satisfying the following equations
//
// x = X/Z
//
//Type1 - XZ Coordinates
type XZCoordinates struct {
    X *big.Int
    Y *big.Int
    Z *big.Int
}
// Extended Coordinates represent aX and aY as eX, eY, eZ and eT satisfying the following equations:
//
// aX = eX/eZ
// aY = eY/eZ
// aX*aY = eT/eZ
// with eZ != 0
//
// Type2 - Extended Coordinates
type ExtendedCoordinates struct {
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
// Type3 - Inverted Coordinates
type InvertedCoordinates struct {
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
// Type4 - Projective Coordinates
type ProjectiveCoordinates struct {
    PX *big.Int
    PY *big.Int
    PZ *big.Int
}

// II - Coordinates Conversion Methods
// 1
func (k *FiniteFieldEllipticCurve) Affine2Extended (InputP AffineCoordinates) (OutputP ExtendedCoordinates) {
    Cmp1 := InputP.AX.Cmp(Zero)
    Cmp2 := InputP.AY.Cmp(Zero)

    //Infinity Point Test
    if Cmp1 == 0 && Cmp2 == 0 {
        OutputP = InfinityPoint
    } else {
        OutputP.EX = InputP.AX
        OutputP.EY = InputP.AY
        OutputP.ET = k.MulModP(InputP.AX,InputP.AY)
        OutputP.EZ = One
    }

    return OutputP
}
// 2
func (k *FiniteFieldEllipticCurve) Affine2Inverted (InputP AffineCoordinates) (OutputP InvertedCoordinates) {
    OutputP.IX = k.QuoModP(One,InputP.AX)
    OutputP.IY = k.QuoModP(One,InputP.AY)
    OutputP.IZ = One
    return OutputP
}
// 3
func (k *FiniteFieldEllipticCurve) Affine2Projective (InputP AffineCoordinates) (OutputP ProjectiveCoordinates) {
    //Infinity Test must be added
    OutputP.PX = InputP.AX
    OutputP.PY = InputP.AY
    OutputP.PZ = One
    return OutputP
}
// 4
func (k *FiniteFieldEllipticCurve) Extended2Affine (InputP ExtendedCoordinates) (OutputP AffineCoordinates) {
    //Infinity Point Test
    if k.IsInfinityPoint(InputP) == true {
        OutputP.AX = Zero
        OutputP.AY = Zero
    } else {
        OutputP.AX = k.QuoModP(InputP.EX,InputP.EZ)
        OutputP.AY = k.QuoModP(InputP.EY,InputP.EZ)
    }
    return OutputP
}
// 5
func (k *FiniteFieldEllipticCurve) Inverted2Affine (InputP InvertedCoordinates) (OutputP AffineCoordinates) {
    //Infinity Test must be added
    OutputP.AX = k.QuoModP(InputP.IZ,InputP.IX)
    OutputP.AY = k.QuoModP(InputP.IZ,InputP.IY)
    return OutputP
}
// 6
func (k *FiniteFieldEllipticCurve) Projective2Affine (InputP ProjectiveCoordinates) (OutputP AffineCoordinates) {
    //Infinity Test must be added
    OutputP.AX = k.QuoModP(InputP.PX,InputP.PZ)
    OutputP.AY = k.QuoModP(InputP.PY,InputP.PZ)
    return OutputP
}