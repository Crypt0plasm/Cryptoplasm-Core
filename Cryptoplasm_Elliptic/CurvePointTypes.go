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
// II.1
func (k *FiniteFieldEllipticCurve) Affine2Extended (InputP AffineCoordinates) (OutputP ExtendedCoordinates) {
    //Infinity Point Test can be removed. If point is of form {0,1}, this function correctly returns
    //  the Extended Point {0,1,1,0}, which is an Infinity Point

    OutputP.EX = InputP.AX
    OutputP.EY = InputP.AY
    OutputP.ET = k.MulModP(InputP.AX,InputP.AY)
    OutputP.EZ = One

    return OutputP
}
// II.2
func (k *FiniteFieldEllipticCurve) Affine2Inverted (InputP AffineCoordinates) (OutputP InvertedCoordinates) {
    //Infinity Test must be added, hasn't been added since no work is done with Inverted Coordinates
    OutputP.IX = k.QuoModP(One,InputP.AX)
    OutputP.IY = k.QuoModP(One,InputP.AY)
    OutputP.IZ = One
    return OutputP
}
// II.3
func (k *FiniteFieldEllipticCurve) Affine2Projective (InputP AffineCoordinates) (OutputP ProjectiveCoordinates) {
    //Infinity Test must be added, hasn't been added since no work is done with Projective Coordinates
    OutputP.PX = InputP.AX
    OutputP.PY = InputP.AY
    OutputP.PZ = One
    return OutputP
}
// 4
func (k *FiniteFieldEllipticCurve) Extended2Affine (InputP ExtendedCoordinates) (OutputP AffineCoordinates) {
    //Infinity Point Test can be removed. If point is of form {0,m,m,0}, this function correctly returns
    //  the Affine Point {0,1}
    OutputP.AX = k.QuoModP(InputP.EX,InputP.EZ)
    OutputP.AY = k.QuoModP(InputP.EY,InputP.EZ)
    return OutputP
}
// II.5
func (k *FiniteFieldEllipticCurve) Inverted2Affine (InputP InvertedCoordinates) (OutputP AffineCoordinates) {
    //Infinity Test must be added, hasn't been added since no work is done with Inverted Coordinates
    OutputP.AX = k.QuoModP(InputP.IZ,InputP.IX)
    OutputP.AY = k.QuoModP(InputP.IZ,InputP.IY)
    return OutputP
}
// II.6
func (k *FiniteFieldEllipticCurve) Projective2Affine (InputP ProjectiveCoordinates) (OutputP AffineCoordinates) {
    //Infinity Test must be added, hasn't been added since no work is done with Projective Coordinates
    OutputP.AX = k.QuoModP(InputP.PX,InputP.PZ)
    OutputP.AY = k.QuoModP(InputP.PY,InputP.PZ)
    return OutputP
}