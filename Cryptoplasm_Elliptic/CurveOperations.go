package Cryptoplasm_Elliptic

import (
    aux "Cryptoplasm-Core/Auxiliary"
    "crypto/rand"
    "fmt"
    "math/big"
)

var (
    //Easy big.Int variables 0,1 and 2.
    Zero = big.NewInt(0)
    One  = big.NewInt(1)
    Two = big.NewInt(2)
    Three = big.NewInt(3)
    Four = big.NewInt(4)
    Eight = big.NewInt(8)
    Sixteen = big.NewInt(16)

    Seven = big.NewInt(7)
    SquaredSeven = big.NewInt(49)
    SeventySeven = big.NewInt(77)
    SevenHundredSeventySeven = big.NewInt(777)
    SevenTimesSeven = big.NewInt(7777777)


    //Technically, InfinityPoint in Extended Coordinates is of the form {0,m,m,0}.
    //Initially an Extended Point in the form of {0,1,0,0} has been artificially imposed as infinity Point.
    //Now the simplest Infinity-Point in Extended Coordinates is defined as {0,1,1,0}. This, as all Extended Points in
    //the Form of {0,m,m,0}, converts to the Affine Point (0,1), which represents the Infinity Point.
    //However we define more Infinity Points as for observing and aesthetic and troubleshooting purposes.
    //
    //Although Elliptic Curves defined over Prime Field in other Forms, dont seem to have an Affine Representation for
    //the Infinity Point, TEC-Curves seem to have, because the Point (0,1) behaves likes 0 in Elliptic Point Addition,
    //therefore it can be called Infinity-Point.
    //

    //InfinityPointSingle results from Additions when two Points are added that are inverse to one another.
    //	See IsInverseOnCurve Function for more Details:
    //InfinityPointSquared results from Additions when two Infinity Points are added.
    //InfinityPointDouble results from Doubling an Infinity Point.
    //InfinityPointTriple results from Tripling an Infinity Point.
    //InfinityPointSevenFold is used as Infinity Point in the Scalar Multiplication Function.

    InfinityPointSingle = ExtendedCoordinates {Zero, Seven,Seven,Zero}
    InfinityPointSquared = ExtendedCoordinates {Zero, SquaredSeven,SquaredSeven,Zero}
    InfinityPointDouble = ExtendedCoordinates {Zero, SeventySeven,SeventySeven,Zero}
    InfinityPointTriple = ExtendedCoordinates {Zero, SevenHundredSeventySeven,SevenHundredSeventySeven,Zero}
    InfinityPointSevenFold = ExtendedCoordinates {Zero, SevenTimesSeven,SevenTimesSeven,Zero}

    //Every TEC curve with Cofactor 4 (Cofactor 8 Curves haven't been studied this way) has in Affine Coordinates
    //4 Permanent Points, these are:
    //PP1 = {0,1}							//Order 1
    //PP2 = {0,P-1}, with P the Prime Field of the Elliptic Curve	//Order 2
    //PP3 = {1,0}							//Order 4
    //PP4 = {P-1,0}							//Order 4
    //
    //Through experimentation it has been observed these points have some special properties, namely:
    //
    //PP1 + Point(x,y) = Point(x,y)
    //Therefore PP1 can designated the "Infinity-Point" of the Elliptic Curve.
    //Further it has been observed that Points in Extended Coordinates of the form {0,m,m,0} convert to PP1 in Affine Coordinates
    //Therefore Extended Coordinates in form of {0,m,m,0} are considered Infinity Points.
    //PP1 Point is the only point with Order 1
    //
    //PP2 + Point1(x,y) = Point2(u,v), with x+u=y+v=P, with P the Prime Field of the Elliptic Curve
    //Therefore PP2 is designated the "Filler-Point" of the Elliptic Curve
    //Further it has been observed that Points in Extended Coordinates of the form {0,m,n,0}, where m+n=P, with P the
    //Prime Field of the Elliptic Curve, convert to PP2 in Affine Coordinates. Therefore such points are "Filler-Points"
    //in extended Coordinates.
    //PP2 is the only Point with Order 2
    //
    //PP3 + Point1(x,y) = Point2(y,v), with x+v=P, with P the Prime Field of the Elliptic Curve.
    //Since the x coordinate of Point2 is equal to the y coordinate of Point1, PP3 is designated as "Translator-Point".
    //PP3 and PP4 have Order 4
    //
    //PP4 + Point1(x,y) = Point2(u,x), with y+u=P, with P the Prime Field of the Elliptic Curve.
    //Since the y coordinate of Point2 is equal to the x coordinate of Point1, PP3 is designated as "Reverse-Translator-Point".
    //PP3 and PP4 have Order 4
)


type FiniteFieldEllipticCurveMethods interface {
    // Ia - Mod P Methods
    AddModP		(a,b *big.Int)					*big.Int			// Ia.1
    SubModP		(a,b *big.Int)					*big.Int			// Ia.2
    MulModP		(a,b *big.Int)					*big.Int			// Ia.3
    QuoModP		(a,b *big.Int)					*big.Int			// Ia.4
    // Ib - Mod Q Methods
    AddModQ		(a,b *big.Int)					*big.Int			// Ib.1
    SubModQ		(a,b *big.Int)					*big.Int			// Ib.2
    MulModQ		(a,b *big.Int)					*big.Int			// Ib.3
    QuoModQ		(a,b *big.Int)					*big.Int			// Ib.4

    // II - Coordinates Conversion Methods
    Affine2Extended 	(InputP AffineCoordinates) 			(OutputP ExtendedCoordinates)	// II.1
    Affine2Inverted 	(InputP AffineCoordinates) 			(OutputP InvertedCoordinates)	// II.2
    Affine2Projective 	(InputP AffineCoordinates) 			(OutputP ProjectiveCoordinates)	// II.3
    Extended2Affine 	(InputP ExtendedCoordinates) 			(OutputP AffineCoordinates)	// II.4
    Inverted2Affine 	(InputP InvertedCoordinates) 			(OutputP AffineCoordinates)	// II.5
    Projective2Affine 	(InputP ProjectiveCoordinates) 			(OutputP AffineCoordinates)	// II.6

    // III - Boolean Methods
    IsInfinityPoint 	(InputP ExtendedCoordinates) 			bool				// III.1
    IsInverseOnCurve	(P1, P2 ExtendedCoordinates) 			bool				// III.2
    IsOnCurve 		(InputP ExtendedCoordinates) 			(OnCurve bool, Infinity bool)	// III.3
    ArePointsEqualEx	(P1, P2 ExtendedCoordinates) 			bool				// III.4
    ArePointsEqualAf	(P1, P2 AffineCoordinates) 			bool				// III.5

    // IV - Basic Point Operations Methods
    DoubleWithZOne 	(InputP ExtendedCoordinates) 			(OutputP ExtendedCoordinates)	// IV.1
    Double 		(InputP ExtendedCoordinates) 			(OutputP ExtendedCoordinates)	// IV.2
    Triple 		(InputP ExtendedCoordinates) 			(OutputP ExtendedCoordinates)	// IV.3
    AdditionZ2OneV2 	(P1, P2 ExtendedCoordinates) 			(OutputP ExtendedCoordinates)	// IV.4
    AdditionV2 		(P1, P2 ExtendedCoordinates) 			(OutputP ExtendedCoordinates)	// IV.5

    // V - Complex Point Operations Methods
    FortyNiner	 	(InputP ExtendedCoordinates) 			(OutputP ExtendedCoordinates)	// V.1
    PrecomputingMatrixG	() 						[7][7]ExtendedCoordinates	// V.2
    PrecomputingMatrixPt(InputP ExtendedCoordinates) 			[7][7]ExtendedCoordinates	// V.3

    //==================================================================================================================
    // VIa - Basic Key Generation Methods - Part I
    GetRandomBitsOnCurve() 						string				// VIa.1
    ClampBitString	(BitString string)				*big.Int			// VIa.2
    ValidateBitString	(BitString string)				(bool, bool, bool)		// VIa.3
    ValidateBigBitString(BitString string)				(bool, bool, bool)		// VIa.4
    GetY		(X *big.Int)					*big.Int			// VIa.5
    ScalarMultiplierG	(Scalar *big.Int)				(OutputP ExtendedCoordinates)	// VIa.6
    ScalarMultiplierPt	(Scalar *big.Int, InputP ExtendedCoordinates)	(OutputP ExtendedCoordinates)	// VIa.7

    // VIb - Complex Key Generation Methods - Part II
    RBS2CRYPTOPLASM	()						(Keys CPKeyPair, Address string)// VIb.1
    GetKeys 		(BitString string)				(Keys CPKeyPair)		// VIb.2
    GetKeysInt 		(Number *big.Int) 				(Keys CPKeyPair)		// VIb.3
    PrivKey2PubKey	(PrivateKey string) 				(PublicKey string)		// VIb.4
    PrivKey2BitString	(PrivateKey string)				(BitString string)		// VIb.5
    PrivKeyInt2PubKey 	(PrivateKeyInt *big.Int) 			(PublicKey string)		// VIb.6

    // VIc - Complete Key Generation Methods - Part III
    MakeCryptoplasmKeys ()										// VIc.0
    CPFromRandomBits	()						(Keys CPKeyPair, Address string)// VIc.1
    CPFromManualBits	()						(Keys CPKeyPair, Address string)// VIc.2
    CPFromNumber	(Base int)					(Keys CPKeyPair, Address string)// VIc.3
    CPFromSeed		()						(Keys CPKeyPair, Address string)// VIc.4
    StringToBitString	(Word string)					string				// VIc.5
    SaveBitString	(BitString string)				()				// VIc.6
    ExportPrivateKey	(BitString, KeyName, Password string)		()				// VIc.7
    OpenKeys		()						(string, error)			// VIc.8

    //==================================================================================================================

    // VII - Schnorr Signature Methods
    SchnorrHash 	(r *big.Int, PublicKey string, Message []byte) 	*big.Int			// VII.1
    SchnorrSign 	(Keys CPKeyPair, Hash []byte) 			(Signature Schnurr)		// VII.2
    SchnorrVerify 	(Sigma Schnurr, PublicKey string, Hash []byte) 	bool				// VII.3

    // VIII - Generator Creating Methods (used for the novel TEC curves to create a Generator)
    GetPointOrder 	(InputP AffineCoordinates) 			*big.Int			// VIII.1
    MakeGenMin		()						(OutputP AffineCoordinates)	// VIII.2
}
//
// FiniteFieldEllipticCurve Methods
//
// Ia - Mod P Methods
//
// Ia.1
func (k *FiniteFieldEllipticCurve) AddModP (a,b *big.Int) *big.Int{
    result := AddModulus(&k.P,a,b)
    return result
}
// Ia.2
func (k *FiniteFieldEllipticCurve) AddModQ (a,b *big.Int) *big.Int{
    result := AddModulus(&k.Q,a,b)
    return result
}
// Ia.3
func (k *FiniteFieldEllipticCurve) SubModP (a,b *big.Int) *big.Int{
    result := SubModulus(&k.P,a,b)
    return result
}
// Ia.4
func (k *FiniteFieldEllipticCurve) SubModQ (a,b *big.Int) *big.Int{
    result := SubModulus(&k.Q,a,b)
    return result
}
// Ib - Mod Q Methods
// Ib.1
func (k *FiniteFieldEllipticCurve) MulModP (a,b *big.Int) *big.Int{
    result := MulModulus(&k.P,a,b)
    return result
}
// Ib.2
func (k *FiniteFieldEllipticCurve) MulModQ (a,b *big.Int) *big.Int{
    result := MulModulus(&k.Q,a,b)
    return result
}
// Ib.3
func (k *FiniteFieldEllipticCurve) QuoModP (a,b *big.Int) *big.Int{
    result := QuoModulus(&k.P,a,b)
    return result
}
// Ib.4
func (k *FiniteFieldEllipticCurve) QuoModQ (a,b *big.Int) *big.Int{
    result := QuoModulus(&k.Q,a,b)
    return result
}
//
// III - Boolean Methods
//
// III.1
func (k *FiniteFieldEllipticCurve) IsInfinityPoint (InputP ExtendedCoordinates) bool {
    var result bool

    Cmp1 := InputP.EX.Cmp(Zero)
    Cmp2 := InputP.ET.Cmp(Zero)
    Cmp3 := InputP.EY.Cmp(InputP.EZ)

    if Cmp1 == 0 && Cmp2 == 0 && Cmp3 == 0 {
	result = true
    } else {
	result = false
    }
    return result
}
// III.2
func (k *FiniteFieldEllipticCurve) IsInverseOnCurve (P1, P2 ExtendedCoordinates) bool {
    //Normally points with similar X values are named inverse for one another. However it has been observed that such
    //Points which have equal X Affine Coordinates, when added up, dont result in an Infinity Point.
    //Paradoxically, it has been observed that Points with equal Y Affine Coordinates do result in an Infinity Point
    //when added up. (Is this a property of Twisted Edwards Curves?)
    //Therefore testing if two Points are inverse on curve involves testing whether their Affine Y coordinates are equal.
    var result bool
    Point1Aff := k.Extended2Affine(P1)
    Point2Aff := k.Extended2Affine(P2)
    //fmt.Println("Point1 is ",Point1Aff.AY)
    //fmt.Println("Point2 is ",Point2Aff.AY)
    Cmp := Point1Aff.AY.Cmp(Point2Aff.AY)


    if Cmp == 0  {
	result = true
    } else {
	result = false
    }
    return result
}
// III.3
func (k *FiniteFieldEllipticCurve) IsOnCurve (InputP ExtendedCoordinates) (OnCurve bool, Infinity bool) {
    var (
	PointAffine = k.Extended2Affine(InputP)
	A 	= new(big.Int)
	B 	= new(big.Int)
    )

    if k.IsInfinityPoint(InputP) == true {
	Infinity = true
    } else {
	Infinity = false
    }

    A.Exp(PointAffine.AX,Two,&k.P)
    B.Exp(PointAffine.AY,Two,&k.P)
    Left := k.AddModP(A,B)
    C := k.MulModP(A,B)
    D := k.MulModP(C,&k.D)
    Right := k.AddModP(One,D)
    //fmt.Println("Left  is",Left)
    //fmt.Println("Right is",Right)
    CompareResult := Left.Cmp(Right)
    if CompareResult == 0 || k.IsInfinityPoint(InputP) == true {
	OnCurve = true
    } else {
	OnCurve = false
    }
    return OnCurve, Infinity
}
// III.4
func (k *FiniteFieldEllipticCurve) ArePointsEqualEx (P1, P2 ExtendedCoordinates) bool {
    var result bool
    Point1Aff := k.Extended2Affine(P1)
    Point2Aff := k.Extended2Affine(P2)
    result = k.ArePointsEqualAf(Point1Aff,Point2Aff)
    return result
}
// III.5
func (k *FiniteFieldEllipticCurve) ArePointsEqualAf (P1, P2 AffineCoordinates) bool {
    var result bool

    Cmp1 := P1.AX.Cmp(P2.AX)
    Cmp2 := P1.AY.Cmp(P2.AY)

    if Cmp1 == 0 && Cmp2 == 0 {
	result = true
    } else {
	result = false
    }

    return result
}
//
// IV - Basic Point Operations Methods
//
// As described in https://cr.yp.to/ecdh/curve41417-20140706.pdf Appendix A
// Since CurveE521 has the same Structure (Twisted Edward Curve) as Curve41417,
// the same formulas hold true for CurveE521 as well.
//
// Formulas used are these: http://hyperelliptic.org/EFD/g1p/data/twisted/extended/doubling/mdbl-2008-hwcd
// Source: 2008 Hisil--Wong--Carter--Dawson, http://eprint.iacr.org/2008/522, Section 3.3, plus assumption Z1=1, plus standard simplification
// PDF - "Twisted Edwards Curves Revisited": https://eprint.iacr.org/2008/522.pdf
//
// The DoubleWithZOne function assumes Z = 1 and considers a = 1
// Standard Elliptic Curve in Twisted Edward form: a * x^2 + y^2 = 1 +      d * x^2 * y^2
// Since E521 Equation is:			       x^2 + y^2 = 1 - 376014 * x^2 * y^2
//
// IV.1
func (k *FiniteFieldEllipticCurve) DoubleWithZOne (InputP ExtendedCoordinates) (OutputP ExtendedCoordinates) {
    var (
	A 	= new(big.Int)
	B 	= new(big.Int)
	v1	= new(big.Int)
    )
    //When Z is one, Point cant be Infinity-Point, therefore no check is needed.
    A.Exp(InputP.EX,Two,&k.P)
    B.Exp(InputP.EY,Two,&k.P)
    E := k.MulModP(Two,InputP.ET)
    G := k.AddModP(A,B)
    F := k.SubModP(G,Two)
    H := k.SubModP(A,B)
    OutputP.EX = k.MulModP(E,F)
    OutputP.EY = k.MulModP(G,H)
    v1.Exp(G,Two,&k.P)
    v2 := k.MulModP(Two,G)
    OutputP.EZ = k.SubModP(v1,v2)
    OutputP.ET = k.MulModP(E,H)
    return OutputP
}
// As described in https://cr.yp.to/ecdh/curve41417-20140706.pdf Appendix A
// Since CurveE521 has the same Structure (Twisted Edward Curve) as Curve41417,
// the same formulas hold true for CurveE521 as well.
//
// Formulas used are these: http://hyperelliptic.org/EFD/g1p/data/twisted/extended/doubling/dbl-2008-hwcd
// Source 2008 Hisil--Wong--Carter--Dawson, http://eprint.iacr.org/2008/522, Section 3.3
// PDF - "Twisted Edwards Curves Revisited": https://eprint.iacr.org/2008/522.pdf
//
// The Double function doesnt make any assumptions regarding Z and considers a = 1
// Standard Elliptic Curve in Twisted Edward form: a * x^2 + y^2 = 1 +      d * x^2 * y^2
// Since E521 Equation is:			       x^2 + y^2 = 1 - 376014 * x^2 * y^2
//
// IV.2
func (k *FiniteFieldEllipticCurve) Double (InputP ExtendedCoordinates) (OutputP ExtendedCoordinates) {
    var (
	A 	= new(big.Int)
	B 	= new(big.Int)
	C 	= new(big.Int)
	v1	= new(big.Int)
    )

    if k.IsInfinityPoint(InputP) == true {
	OutputP = InfinityPointDouble
    } else {
	A.Exp(InputP.EX,Two,&k.P)
	B.Exp(InputP.EY,Two,&k.P)
	C.Exp(InputP.EZ,Two,&k.P)
	C = k.MulModP(C,Two)
	v2 := k.AddModP(InputP.EX,InputP.EY)
	v1.Exp(v2,Two,&k.P)
	G := k.AddModP(A,B)
	E := k.SubModP(v1,G)
	F := k.SubModP(G,C)
	H := k.SubModP(A,B)
	OutputP.EX = k.MulModP(E,F)
	OutputP.EY = k.MulModP(G,H)
	OutputP.EZ = k.MulModP(F,G)
	OutputP.ET = k.MulModP(E,H)
    }
    return OutputP
}
// Formulas used are these: http://hyperelliptic.org/EFD/g1p/data/twisted/extended/tripling/tpl-2015-c
// Source 2015 Chuengsatiansup
// Standard Elliptic Curve in Twisted Edward form: a * x^2 + y^2 = 1 +      d * x^2 * y^2
// Since E521 Equation is:			       x^2 + y^2 = 1 - 376014 * x^2 * y^2,
// It is assumed that a = 1, and XX is used instead of aXX
//
// IV.3
func (k *FiniteFieldEllipticCurve) Triple (InputP ExtendedCoordinates) (OutputP ExtendedCoordinates) {
    var (
	YY	= new(big.Int)
	XX	= new(big.Int)
	ZZ	= new(big.Int)
    )

    if k.IsInfinityPoint(InputP) == true {
	OutputP = InfinityPointTriple
    } else {
	YY.Exp(InputP.EY, Two,&k.P)
	XX.Exp(InputP.EX, Two,&k.P)
	ZZ.Exp(InputP.EZ, Two,&k.P)
	Ap := k.AddModP(YY,XX)
	twoZZ := k.MulModP(Two,ZZ)
	Diff1 := k.SubModP(twoZZ,Ap)
	B := k.MulModP(Two,Diff1)
	xB := k.MulModP(XX,B)
	yB := k.MulModP(YY,B)
	Diff2 := k.SubModP(YY,XX)
	AA := k.MulModP(Ap,Diff2)
	F := k.SubModP(AA,yB)
	G := k.AddModP(AA,xB)
	Sum := k.AddModP(yB,AA)
	xE := k.MulModP(InputP.EX,Sum)
	Diff3 := k.SubModP(xB,AA)
	yH := k.MulModP(InputP.EY,Diff3)
	zF := k.MulModP(InputP.EZ,F)
	zG := k.MulModP(InputP.EZ,G)
	OutputP.EX = k.MulModP(xE,zF)
	OutputP.EY = k.MulModP(zG,yH)
	OutputP.EZ = k.MulModP(zF,zG)
	OutputP.ET = k.MulModP(xE,yH)
    }
    return OutputP
}
// As described in https://cr.yp.to/ecdh/curve41417-20140706.pdf Appendix A
// Since CurveE521 has the same Structure (Twisted Edward Curve) as Curve41417,
// the same formulas hold true for CurveE521 as well.
// The Precomputing Addition always adds the BasePoint in ExtCoordinates, therefore Z is always 1.
//
// Formulas used are these: http://hyperelliptic.org/EFD/g1p/data/twisted/extended/addition/madd-2008-hwcd
// Source: 2008 Hisil--Wong--Carter--Dawson, http://eprint.iacr.org/2008/522, Section 3.1
// PDF - "Twisted Edwards Curves Revisited": https://eprint.iacr.org/2008/522.pdf
//
// This Addition Variant assumes Z2 = 1 and considers D != 0 and a = 1
// Standard Elliptic Curve in Twisted Edward form: a * x^2 + y^2 = 1 +      d * x^2 * y^2
// Since E521 Equation is:			       x^2 + y^2 = 1 - 376014 * x^2 * y^2
//
// IV.4
func (k *FiniteFieldEllipticCurve) AdditionZ2OneV2 (P1, P2 ExtendedCoordinates) (OutputP ExtendedCoordinates) {
    //When Z2 is one, Point2 cant be an InfinityPoint, therefore only a check for Point1 is needed.
    //Also a test is made if Point1 and Point2 are inverse to one another. In this case the sum is the point at Infinity
    //
    //Although these Test are not needed, when they hold true, no further computation has to be done,
    //If no Test were to be done, and the computation would proceed as coded in the last part, it would still produce a
    //correct result.
    //Therefore the Tests only spare an unnecessary  computation.

    if k.IsInfinityPoint(P1) == true {
	OutputP = P2
    } else if k.IsInverseOnCurve(P1,P2) == true {
	OutputP = InfinityPointSingle
    } else {
	A := k.MulModP(P1.EX,P2.EX)
	B := k.MulModP(P1.EY,P2.EY)
	dC:= k.MulModP(P2.ET,&k.D)
	C := k.MulModP(P1.ET,dC)
	v1 := k.AddModP(P1.EX,P1.EY)
	v2 := k.AddModP(P2.EX,P2.EY)
	v3 := k.AddModP(A,B)
	v4 := k.MulModP(v1,v2)
	E := k.SubModP(v4,v3)
	F := k.SubModP(P1.EZ,C)
	G := k.AddModP(P1.EZ,C)
	H := k.SubModP(B,A)
	OutputP.EX = k.MulModP(E,F)
	OutputP.EY = k.MulModP(G,H)
	OutputP.EZ = k.MulModP(F,G)
	OutputP.ET = k.MulModP(E,H)
    }
    return OutputP
}
// As described in https://cr.yp.to/ecdh/curve41417-20140706.pdf Appendix A
// Since CurveE521 has the same Structure (Twisted Edward Curve) as Curve41417,
// the same formulas hold true for CurveE521 as well.
// The Precomputing Addition always adds the BasePoint in ExtCoordinates, therefore Z is always 1.
//
// Formulas used are these: http://hyperelliptic.org/EFD/g1p/data/twisted/extended/addition/add-2008-hwcd
// Source: 2008 Hisil--Wong--Carter--Dawson, http://eprint.iacr.org/2008/522, Section 3.1
// PDF - "Twisted Edwards Curves Revisited": https://eprint.iacr.org/2008/522.pdf
//
// This Addition Variant doesnt make any assumptions regarding Z2 and considers D != 0, and a = 1
// Standard Elliptic Curve in Twisted Edward form: a * x^2 + y^2 = 1 +      d * x^2 * y^2
// Since E521 Equation is:			       x^2 + y^2 = 1 - 376014 * x^2 * y^2
//
// IV.5
func (k *FiniteFieldEllipticCurve) AdditionV2 (P1, P2 ExtendedCoordinates) (OutputP ExtendedCoordinates) {
    //Both Points are tested if they are each or both InfinityPoints, or if the are inverse to one another.
    //
    //Although these Test are not needed, when they hold true, no further computation has to be done,
    //If no Test were to be done, and the computation would proceed as coded in the last part, it would still produce a
    //correct result.
    //Therefore the Tests only spare an unnecessary  computation.

    if k.IsInfinityPoint(P1) == true {
	OutputP = P2
    } else if k.IsInfinityPoint(P2) == true {
	OutputP = P1
    } else if k.IsInfinityPoint(P1) == true && k.IsInfinityPoint(P2) == true {
	OutputP = InfinityPointSquared
    } else if k.IsInverseOnCurve(P1,P2) == true {
	OutputP = InfinityPointSingle
    } else {
	A := k.MulModP(P1.EX,P2.EX)
	B := k.MulModP(P1.EY,P2.EY)
	dC:= k.MulModP(P2.ET,&k.D)
	C := k.MulModP(P1.ET,dC)
	D := k.MulModP(P1.EZ,P2.EZ)
	v1:= k.AddModP(P1.EX,P1.EY)
	v2:= k.AddModP(P2.EX,P2.EY)
	v3:= k.AddModP(A,B)
	v4:= k.MulModP(v1,v2)
	E := k.SubModP(v4,v3)
	F := k.SubModP(D,C)
	G := k.AddModP(D,C)
	H := k.SubModP(B,A)
	OutputP.EX = k.MulModP(E,F)
	OutputP.EY = k.MulModP(G,H)
	OutputP.EZ = k.MulModP(F,G)
	OutputP.ET = k.MulModP(E,H)
    }
    return OutputP
}
//
// V - Complex Point Operations Methods
//
// V.1
// Adds a Point to itself 49 times
func (k *FiniteFieldEllipticCurve) FortyNiner (InputP ExtendedCoordinates) (OutputP ExtendedCoordinates) {
    Point03 := k.Triple(InputP)
    Point06 := k.Double(Point03)
    Point12 := k.Double(Point06)
    Point24 := k.Double(Point12)
    Point48 := k.Double(Point24)
    Point49 := k.AdditionV2(Point48,InputP)

    return Point49
}
// V.2
// Creates a precomputing Matrix with the curve Generator as is Base Point
func (k *FiniteFieldEllipticCurve) PrecomputingMatrixG () [7][7]ExtendedCoordinates {
    CurveGenerator 	:= AffineCoordinates{AX: &k.PBX, AY: &k.PBY}
    CurveGeneratorExt 	:= k.Affine2Extended(CurveGenerator)
    Result 		:= k.PrecomputingMatrixPt(CurveGeneratorExt)
    return Result
}
// V.3
// Creates a precomputing Matrix with a given Point as is Base Point
func (k *FiniteFieldEllipticCurve) PrecomputingMatrixPt (InputP ExtendedCoordinates) [7][7]ExtendedCoordinates {
    //start := time.Now()
    BasePointExt02 	:= k.Double(InputP)
    BasePointExt03 	:= k.AdditionV2(BasePointExt02,InputP)
    BasePointExt04 	:= k.Double(BasePointExt02)
    BasePointExt05 	:= k.AdditionV2(BasePointExt04,InputP)
    BasePointExt06	:= k.Double(BasePointExt03)
    BasePointExt07	:= k.AdditionV2(BasePointExt06,InputP)
    BasePointExt08	:= k.Double(BasePointExt04)
    BasePointExt09	:= k.AdditionV2(BasePointExt08,InputP)
    BasePointExt10	:= k.Double(BasePointExt05)
    BasePointExt11	:= k.AdditionV2(BasePointExt10,InputP)
    BasePointExt12	:= k.Double(BasePointExt06)
    BasePointExt13	:= k.AdditionV2(BasePointExt12,InputP)
    BasePointExt14	:= k.Double(BasePointExt07)
    BasePointExt15	:= k.AdditionV2(BasePointExt14,InputP)
    BasePointExt16	:= k.Double(BasePointExt08)
    BasePointExt17	:= k.AdditionV2(BasePointExt16,InputP)
    BasePointExt18	:= k.Double(BasePointExt09)
    BasePointExt19	:= k.AdditionV2(BasePointExt18,InputP)
    BasePointExt20	:= k.Double(BasePointExt10)
    BasePointExt21	:= k.AdditionV2(BasePointExt20,InputP)
    BasePointExt22	:= k.Double(BasePointExt11)
    BasePointExt23	:= k.AdditionV2(BasePointExt22,InputP)
    BasePointExt24	:= k.Double(BasePointExt12)
    BasePointExt25	:= k.AdditionV2(BasePointExt24,InputP)
    BasePointExt26	:= k.Double(BasePointExt13)
    BasePointExt27	:= k.AdditionV2(BasePointExt26,InputP)
    BasePointExt28	:= k.Double(BasePointExt14)
    BasePointExt29	:= k.AdditionV2(BasePointExt28,InputP)
    BasePointExt30	:= k.Double(BasePointExt15)
    BasePointExt31	:= k.AdditionV2(BasePointExt30,InputP)
    BasePointExt32	:= k.Double(BasePointExt16)
    BasePointExt33	:= k.AdditionV2(BasePointExt32,InputP)
    BasePointExt34	:= k.Double(BasePointExt17)
    BasePointExt35	:= k.AdditionV2(BasePointExt34,InputP)
    BasePointExt36	:= k.Double(BasePointExt18)
    BasePointExt37	:= k.AdditionV2(BasePointExt36,InputP)
    BasePointExt38	:= k.Double(BasePointExt19)
    BasePointExt39	:= k.AdditionV2(BasePointExt38,InputP)
    BasePointExt40	:= k.Double(BasePointExt20)
    BasePointExt41	:= k.AdditionV2(BasePointExt40,InputP)
    BasePointExt42	:= k.Double(BasePointExt21)
    BasePointExt43	:= k.AdditionV2(BasePointExt42,InputP)
    BasePointExt44	:= k.Double(BasePointExt22)
    BasePointExt45	:= k.AdditionV2(BasePointExt44,InputP)
    BasePointExt46	:= k.Double(BasePointExt23)
    BasePointExt47	:= k.AdditionV2(BasePointExt46,InputP)
    BasePointExt48	:= k.Double(BasePointExt24)
    BasePointExt49	:= k.AdditionV2(BasePointExt48,InputP)
    //Point49 isn't used in the pre-computation, it is only created to fill the Matrix.

    MatrixRow0 := [...]ExtendedCoordinates{InputP,BasePointExt02,BasePointExt03,BasePointExt04,BasePointExt05,BasePointExt06,BasePointExt07}
    MatrixRow1 := [...]ExtendedCoordinates{BasePointExt08,BasePointExt09,BasePointExt10,BasePointExt11,BasePointExt12,BasePointExt13,BasePointExt14}
    MatrixRow2 := [...]ExtendedCoordinates{BasePointExt15,BasePointExt16,BasePointExt17,BasePointExt18,BasePointExt19,BasePointExt20,BasePointExt21}
    MatrixRow3 := [...]ExtendedCoordinates{BasePointExt22,BasePointExt23,BasePointExt24,BasePointExt25,BasePointExt26,BasePointExt27,BasePointExt28}
    MatrixRow4 := [...]ExtendedCoordinates{BasePointExt29,BasePointExt30,BasePointExt31,BasePointExt32,BasePointExt33,BasePointExt34,BasePointExt35}
    MatrixRow5 := [...]ExtendedCoordinates{BasePointExt36,BasePointExt37,BasePointExt38,BasePointExt39,BasePointExt40,BasePointExt41,BasePointExt42}
    MatrixRow6 := [...]ExtendedCoordinates{BasePointExt43,BasePointExt44,BasePointExt45,BasePointExt46,BasePointExt47,BasePointExt48,BasePointExt49}

    Matrix := [7][7]ExtendedCoordinates{MatrixRow0,MatrixRow1,MatrixRow2,MatrixRow3,MatrixRow4,MatrixRow5,MatrixRow6}

    //elapsed := time.Since(start)
    //fmt.Println("")
    //fmt.Println("Computing and verifying 47 points took", elapsed)
    return Matrix
}
//
// VIa - Basic Key Generation Methods - Part I
//
// VIa.1
// Generates a Random BitString according to curve Safe-Scalar-Size
func (k *FiniteFieldEllipticCurve) GetRandomBitsOnCurve () string {
    var (
	CoreSlice = make([]string, k.S)
	BinaryString string
    )
    //Generating k.S random bits.
    for i := uint64(0); i < k.S; i++ {
	RandomNumber,_ := rand.Int(rand.Reader,Two)
	String := RandomNumber.Text(2)
	CoreSlice[i] = String
    }
    //Converting the Slice of 0s and 1s to a String
    for i := 0; i < len(CoreSlice); i++ {
	BinaryString = BinaryString + CoreSlice[i]
    }
    return BinaryString
}
//
// VIa.2
// Clamps a BitString according to Curve Cofactor. It is assumed the the BitString is of Safe-Scalar-Size size.
func (k *FiniteFieldEllipticCurve) ClampBitString (BitString string) *big.Int {
    var (
	BinaryString string
	Scalar = new(big.Int)
    )
    Truth,_,_ := k.ValidateBitString(BitString)
    if Truth == true {
	//Creating the final zeros of the string
	BinaryCofactor := k.R.Text(2)
	TrimmedBinaryCofactor := aux.TrimFirstRune(BinaryCofactor)
	BinaryString = "1" + BitString + TrimmedBinaryCofactor
    } else {
	fmt.Println("The BitString is invalid with the ",k.Name,"Curve")
    }
    //Converting the BinaryString to big.Int
    Scalar.SetString(BinaryString,2)
    return Scalar
}
//
// VIa.3
// Checks if a given String is composed of only 0 and 1, and if it's size is according to Safe-Scalar-Size size.
func (k *FiniteFieldEllipticCurve) ValidateBitString (BitString string) (bool, bool, bool) {
    var (
	Count uint64
	r = []rune(BitString)
	LengthTruth,StructureTruth,TotalTruth bool
    )
    Length := uint64(len(r))

    if Length == k.S {
	LengthTruth = true
    } else {
	LengthTruth = false
    }

    for i:=0; i<len(r); i++{
	if string(r[i]) == "0" || string(r[i]) == "1" {
	    Count = Count + 1
	} else {
	    Count = Count + 0
	}
    }
    if Count == Length {
	StructureTruth = true
    } else {
	StructureTruth = false
    }

    if LengthTruth == true && StructureTruth == true {
	TotalTruth = true
    } else {
	TotalTruth = false
    }

    return TotalTruth,LengthTruth,StructureTruth
}
//
// VIa.4
// Checks if a string of 0s and 1s has a proper form compatible with a clamped Scalar for the given Curve.
func (k *FiniteFieldEllipticCurve) ValidateBigBitString (BitString string) (bool, bool, bool) {
    var (
        Count uint64
	r = []rune(BitString)
	T1,T2 bool
	LengthTruth,StructureTruth,TotalTruth bool
    )
    //Getting the Cofactor Bits
    BinaryCofactor := k.R.Text(2)
    TrimmedBinaryCofactor := aux.TrimFirstRune(BinaryCofactor)
    r2 := []rune(TrimmedBinaryCofactor)

    //Assessing Length Truth
    if 1+k.S+uint64(len(r2)) == uint64(len(r)) {
	LengthTruth = true
    }

    //Assessing if the first character is 1(T1), and if the last characters are 0s according to Curve Cofactor(T2)
    if string(r[0]) == "1" {
        T1 = true
    }

    for i:=1; i<=len(r2); i++{
        if string(r[len(r)-i]) == "0" {
            Count = Count + 1
	} else {
	    Count = Count + 0
	}
    }
    if Count == uint64(len(r2)) {
	T2 = true
    } else {
	T2 = false
    }
    if T1 == true && T2 == true {
	StructureTruth = true
    }

    //If both Condition are true, then the BitString is validated
    if LengthTruth == true && StructureTruth == true {
        TotalTruth = true
    } else {
        TotalTruth = false
    }
    return TotalTruth,LengthTruth,StructureTruth
}
//
// VIa.5
// Computes the Y value of a Point given its X value. Works only for Twisted Edwards Curves and their specific formula.
func (k *FiniteFieldEllipticCurve) GetY (X *big.Int) *big.Int {
    var (
    	XSq = new(big.Int)
    	Y = new(big.Int)
    )
    XSq.Exp(X,Two,&k.P)
    M := k.MulModP(&k.A,XSq)
    N := k.MulModP(&k.D,XSq)
    MM := k.SubModP(One,M)
    NN := k.SubModP(One,N)
    YSq := k.QuoModP(MM,NN)

    Y.ModSqrt(YSq,&k.P)
    return Y
}
//
// VIa.6
// Multiplies the Base Point of the Curve with a Scalar (which has been properly clamped previously)
func (k *FiniteFieldEllipticCurve) ScalarMultiplierG (Scalar *big.Int) (OutputP ExtendedCoordinates) {
    CurveGenerator 	:= AffineCoordinates{AX: &k.PBX, AY: &k.PBY}
    CurveGeneratorExt 	:= k.Affine2Extended(CurveGenerator)
    Result 		:= k.ScalarMultiplierPt(Scalar, CurveGeneratorExt)
    return Result
}
//
// VIa.7
// Multiplies a given Point on the Curve with a Scalar (which has been properly clamped previously)
func (k *FiniteFieldEllipticCurve) ScalarMultiplierPt (Scalar *big.Int, InputP ExtendedCoordinates) (OutputP ExtendedCoordinates) {
    var (
	//start = time.Now()
	PrivKey49 		= Scalar.Text(49)
	PrivKey49SliceRune 	= []rune(PrivKey49)
	PrivKey49SliceString 	= make([]string,len(PrivKey49))
	ZeroPoint		= InfinityPointSevenFold
	Result			ExtendedCoordinates
    )
    //start := time.Now()
    for i := 0; i < len(PrivKey49); i++ {
	PrivKey49SliceString[i] = string(PrivKey49SliceRune[i])
    }
    Result = ZeroPoint

    //If Scalar Mod Q is Zero, Multiplication is InfinityPoint, if the multiplied point has order Q
    //However the Scalar will never be chosen as to be divisible by Q

    PrecMatrix := k.PrecomputingMatrixPt(InputP)
    for i := 0; i < len(PrivKey49SliceString); i++ {
	Character := PrivKey49SliceString[i]
	switch Character {
	//At last slice element, a 49x Point multiplication isn't executed.
	//49x Point Multiplication occurs on if i is not the last element in the slice
	case "0":
	    Result = k.AdditionV2(Result,ZeroPoint)
	    if i != len(PrivKey49SliceString) - 1 {
		Result = k.FortyNiner(Result)
	    }
	case "1":
	    Result = k.AdditionV2(Result,PrecMatrix[0][0])
	    if i != len(PrivKey49SliceString) - 1 {
		Result = k.FortyNiner(Result)
	    }
	case "2":
	    Result = k.AdditionV2(Result,PrecMatrix[0][1])
	    if i != len(PrivKey49SliceString) - 1 {
		Result = k.FortyNiner(Result)
	    }
	case "3":
	    Result = k.AdditionV2(Result,PrecMatrix[0][2])
	    if i != len(PrivKey49SliceString) - 1 {
		Result = k.FortyNiner(Result)
	    }
	case "4":
	    Result = k.AdditionV2(Result,PrecMatrix[0][3])
	    if i != len(PrivKey49SliceString) - 1 {
		Result = k.FortyNiner(Result)
	    }
	case "5":
	    Result = k.AdditionV2(Result,PrecMatrix[0][4])
	    if i != len(PrivKey49SliceString) - 1 {
		Result = k.FortyNiner(Result)
	    }
	case "6":
	    Result = k.AdditionV2(Result,PrecMatrix[0][5])
	    if i != len(PrivKey49SliceString) - 1 {
		Result = k.FortyNiner(Result)
	    }
	case "7":
	    Result = k.AdditionV2(Result,PrecMatrix[0][6])
	    if i != len(PrivKey49SliceString) - 1 {
		Result = k.FortyNiner(Result)
	    }
	case "8":
	    Result = k.AdditionV2(Result,PrecMatrix[1][0])
	    if i != len(PrivKey49SliceString) - 1 {
		Result = k.FortyNiner(Result)
	    }
	case "9":
	    Result = k.AdditionV2(Result,PrecMatrix[1][1])
	    if i != len(PrivKey49SliceString) - 1 {
		Result = k.FortyNiner(Result)
	    }
	case "a":
	    Result = k.AdditionV2(Result,PrecMatrix[1][2])
	    if i != len(PrivKey49SliceString) - 1 {
		Result = k.FortyNiner(Result)
	    }
	case "b":
	    Result = k.AdditionV2(Result,PrecMatrix[1][3])
	    if i != len(PrivKey49SliceString) - 1 {
		Result = k.FortyNiner(Result)
	    }
	case "c":
	    Result = k.AdditionV2(Result,PrecMatrix[1][4])
	    if i != len(PrivKey49SliceString) - 1 {
		Result = k.FortyNiner(Result)
	    }
	case "d":
	    Result = k.AdditionV2(Result,PrecMatrix[1][5])
	    if i != len(PrivKey49SliceString) - 1 {
		Result = k.FortyNiner(Result)
	    }
	case "e":
	    Result = k.AdditionV2(Result,PrecMatrix[1][6])
	    if i != len(PrivKey49SliceString) - 1 {
		Result = k.FortyNiner(Result)
	    }
	case "f":
	    Result = k.AdditionV2(Result,PrecMatrix[2][0])
	    if i != len(PrivKey49SliceString) - 1 {
		Result = k.FortyNiner(Result)
	    }
	case "g":
	    Result = k.AdditionV2(Result,PrecMatrix[2][1])
	    if i != len(PrivKey49SliceString) - 1 {
		Result = k.FortyNiner(Result)
	    }
	case "h":
	    Result = k.AdditionV2(Result,PrecMatrix[2][2])
	    if i != len(PrivKey49SliceString) - 1 {
		Result = k.FortyNiner(Result)
	    }
	case "i":
	    Result = k.AdditionV2(Result,PrecMatrix[2][3])
	    if i != len(PrivKey49SliceString) - 1 {
		Result = k.FortyNiner(Result)
	    }
	case "j":
	    Result = k.AdditionV2(Result,PrecMatrix[2][4])
	    if i != len(PrivKey49SliceString) - 1 {
		Result = k.FortyNiner(Result)
	    }
	case "k":
	    Result = k.AdditionV2(Result,PrecMatrix[2][5])
	    if i != len(PrivKey49SliceString) - 1 {
		Result = k.FortyNiner(Result)
	    }
	case "l":
	    Result = k.AdditionV2(Result,PrecMatrix[2][6])
	    if i != len(PrivKey49SliceString) - 1 {
		Result = k.FortyNiner(Result)
	    }
	case "m":
	    Result = k.AdditionV2(Result,PrecMatrix[3][0])
	    if i != len(PrivKey49SliceString) - 1 {
		Result = k.FortyNiner(Result)
	    }
	case "n":
	    Result = k.AdditionV2(Result,PrecMatrix[3][1])
	    if i != len(PrivKey49SliceString) - 1 {
		Result = k.FortyNiner(Result)
	    }
	case "o":
	    Result = k.AdditionV2(Result,PrecMatrix[3][2])
	    if i != len(PrivKey49SliceString) - 1 {
		Result = k.FortyNiner(Result)
	    }
	case "p":
	    Result = k.AdditionV2(Result,PrecMatrix[3][3])
	    if i != len(PrivKey49SliceString) - 1 {
		Result = k.FortyNiner(Result)
	    }
	case "q":
	    Result = k.AdditionV2(Result,PrecMatrix[3][4])
	    if i != len(PrivKey49SliceString) - 1 {
		Result = k.FortyNiner(Result)
	    }
	case "r":
	    Result = k.AdditionV2(Result,PrecMatrix[3][5])
	    if i != len(PrivKey49SliceString) - 1 {
		Result = k.FortyNiner(Result)
	    }
	case "s":
	    Result = k.AdditionV2(Result,PrecMatrix[3][6])
	    if i != len(PrivKey49SliceString) - 1 {
		Result = k.FortyNiner(Result)
	    }
	case "t":
	    Result = k.AdditionV2(Result,PrecMatrix[4][0])
	    if i != len(PrivKey49SliceString) - 1 {
		Result = k.FortyNiner(Result)
	    }
	case "u":
	    Result = k.AdditionV2(Result,PrecMatrix[4][1])
	    if i != len(PrivKey49SliceString) - 1 {
		Result = k.FortyNiner(Result)
	    }
	case "v":
	    Result = k.AdditionV2(Result,PrecMatrix[4][2])
	    if i != len(PrivKey49SliceString) - 1 {
		Result = k.FortyNiner(Result)
	    }
	case "w":
	    Result = k.AdditionV2(Result,PrecMatrix[4][3])
	    if i != len(PrivKey49SliceString) - 1 {
		Result = k.FortyNiner(Result)
	    }
	case "x":
	    Result = k.AdditionV2(Result,PrecMatrix[4][4])
	    if i != len(PrivKey49SliceString) - 1 {
		Result = k.FortyNiner(Result)
	    }
	case "y":
	    Result = k.AdditionV2(Result,PrecMatrix[4][5])
	    if i != len(PrivKey49SliceString) - 1 {
		Result = k.FortyNiner(Result)
	    }
	case "z":
	    Result = k.AdditionV2(Result,PrecMatrix[4][6])
	    if i != len(PrivKey49SliceString) - 1 {
		Result = k.FortyNiner(Result)
	    }
	case "A":
	    Result = k.AdditionV2(Result,PrecMatrix[5][0])
	    if i != len(PrivKey49SliceString) - 1 {
		Result = k.FortyNiner(Result)
	    }
	case "B":
	    Result = k.AdditionV2(Result,PrecMatrix[5][1])
	    if i != len(PrivKey49SliceString) - 1 {
		Result = k.FortyNiner(Result)
	    }
	case "C":
	    Result = k.AdditionV2(Result,PrecMatrix[5][2])
	    if i != len(PrivKey49SliceString) - 1 {
		Result = k.FortyNiner(Result)
	    }
	case "D":
	    Result = k.AdditionV2(Result,PrecMatrix[5][3])
	    if i != len(PrivKey49SliceString) - 1 {
		Result = k.FortyNiner(Result)
	    }
	case "E":
	    Result = k.AdditionV2(Result,PrecMatrix[5][4])
	    if i != len(PrivKey49SliceString) - 1 {
		Result = k.FortyNiner(Result)
	    }
	case "F":
	    Result = k.AdditionV2(Result,PrecMatrix[5][5])
	    if i != len(PrivKey49SliceString) - 1 {
		Result = k.FortyNiner(Result)
	    }
	case "G":
	    Result = k.AdditionV2(Result,PrecMatrix[5][6])
	    if i != len(PrivKey49SliceString) - 1 {
		Result = k.FortyNiner(Result)
	    }
	case "H":
	    Result = k.AdditionV2(Result,PrecMatrix[6][0])
	    if i != len(PrivKey49SliceString) - 1 {
		Result = k.FortyNiner(Result)
	    }
	case "I":
	    Result = k.AdditionV2(Result,PrecMatrix[6][1])
	    if i != len(PrivKey49SliceString) - 1 {
		Result = k.FortyNiner(Result)
	    }
	case "J":
	    Result = k.AdditionV2(Result,PrecMatrix[6][2])
	    if i != len(PrivKey49SliceString) - 1 {
		Result = k.FortyNiner(Result)
	    }
	case "K":
	    Result = k.AdditionV2(Result,PrecMatrix[6][3])
	    if i != len(PrivKey49SliceString) - 1 {
		Result = k.FortyNiner(Result)
	    }
	case "L":
	    Result = k.AdditionV2(Result,PrecMatrix[6][4])
	    if i != len(PrivKey49SliceString) - 1 {
		Result = k.FortyNiner(Result)
	    }
	case "M":
	    Result = k.AdditionV2(Result,PrecMatrix[6][5])
	    if i != len(PrivKey49SliceString) - 1 {
		Result = k.FortyNiner(Result)
	    }
	}
    }
    OutputP = Result
    //elapsed := time.Since(start)
    //fmt.Println("")
    //fmt.Println("Computing PublicKey points took:", elapsed)
    return OutputP
}
//
// VIII - Generator Creating Methods (used for the novel TEC curves to create a Generator)
//
// VIII.1
// Creates a Generator Point for the given curve with a minimum X value, such as the Point resulted has order Q
func (k *FiniteFieldEllipticCurve) MakeGenMin () (OutputP AffineCoordinates) {
    var(
	ComputedP AffineCoordinates
	Start 	= big.NewInt(1)
	End 	= k.Q
	X	= new(big.Int)
	Y	= new(big.Int)
    )

    // i must be a new big.Int so it does not overwrite Start
    for i:=new(big.Int).Set(Start); i.Cmp(&End) < 0; i.Add(i,One) {
	Y = k.GetY(i)
	if Y.Cmp(Zero) != 0 {
	    X = i
	    ComputedP = AffineCoordinates{X,Y}
	    if k.GetPointOrder(ComputedP) == "Q" {
		break
	    }
	}

    }
    IsPointOnCurve,_ := k.IsOnCurve(k.Affine2Extended(ComputedP))
    if IsPointOnCurve == true {
	OutputP = ComputedP
    }
    return OutputP
}
//
// VIII.2
// Returns the order of a given Point for the given Curve
func (k *FiniteFieldEllipticCurve) GetPointOrder (InputP AffineCoordinates) string {
    var (
	PointOrderS 	string
	PointOrder 	= big.NewInt(-1)
	MinusOne	= big.NewInt(-1)
	QScalar		= new(big.Int)
    )
    PointExt := k.Affine2Extended(InputP)
    CofactorDivisors := aux.ListDivisors(&k.R)

    fmt.Println("Cofactor Divisors are",CofactorDivisors)
    for i:=0; i<len(CofactorDivisors); i++ {
	I := k.ScalarMultiplierPt(CofactorDivisors[i],PointExt)
	fmt.Println("i",i,"is mult by",CofactorDivisors[i],"and results is",I)
	if k.IsInfinityPoint(I) == true {
	    PointOrder = CofactorDivisors[i]
	    PointOrderS = PointOrder.Text(10)
	    break
	}
    }

    Compare := PointOrder.Cmp(MinusOne)
    if Compare == 0 {
	for j:=0; j<len(CofactorDivisors); j++ {
	    if j == 0 {
		J := k.ScalarMultiplierPt(&k.Q,PointExt)
		fmt.Println("j",j,"is mult by Q and results is",J)
		if k.IsInfinityPoint(J) == true {
		    PointOrder = &k.Q
		    PointOrderS = "Q"
		    break
		}
	    } else {
		QScalar.Mul(&k.Q,CofactorDivisors[j])
		J := k.ScalarMultiplierPt(QScalar,PointExt)
		js := CofactorDivisors[j].Text(10)
		Ss := js+"Q"
		fmt.Println("j",j,"is mult by",Ss," and results is",J)
		if k.IsInfinityPoint(J) == true {
		    PointOrder = QScalar
		    PointOrderS = Ss
		    break
		}
	    }
	}
    }
    return PointOrderS
}