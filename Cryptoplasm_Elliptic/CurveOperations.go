package Cryptoplasm_Elliptic

import (
    "crypto/rand"
    "math/big"
)

var (
    //Easy big.Int variables 0,1 and 2.
    Zero = big.NewInt(0)
    One  = big.NewInt(1)
    Two = big.NewInt(2)

    CurveE521 = ParamE521()
    InfinityPoint = ExtendedCoordinates {Zero, One,Zero,Zero}
)


type TwistedEdwardsCurve interface {
    // I - Mod P Methods
    AddMod		(a,b *big.Int)					*big.Int
    SubMod		(a,b *big.Int)					*big.Int
    MulMod		(a,b *big.Int)					*big.Int
    
    // II - Coordinates Conversion Methods
    Affine2Extended 	(InputP AffineCoordinates) 			(OutputP ExtendedCoordinates)
    Affine2Inverted 	(InputP AffineCoordinates) 			(OutputP InvertedCoordinates)
    Affine2Projective 	(InputP AffineCoordinates) 			(OutputP ProjectiveCoordinates)
    Extended2Affine 	(InputP ExtendedCoordinates) 			(OutputP AffineCoordinates)
    Inverted2Affine 	(InputP InvertedCoordinates) 			(OutputP AffineCoordinates)
    Projective2Affine 	(InputP ProjectiveCoordinates) 			(OutputP AffineCoordinates)

    // III - Boolean Methods
    IsInfinityPoint 	(InputP ExtendedCoordinates) 			bool
    IsInverseOnCurve	(P1, P2 ExtendedCoordinates) 			bool
    IsOnCurve 		(InputP ExtendedCoordinates) 			bool
    
    // IV - Basic Point Operations Methods
    DoubleWithZOne 	(InputP ExtendedCoordinates) 			(OutputP ExtendedCoordinates)
    Double 		(InputP ExtendedCoordinates) 			(OutputP ExtendedCoordinates)
    Triple 		(InputP ExtendedCoordinates) 			(OutputP ExtendedCoordinates)
    AdditionZ2OneV2 	(P1, P2 ExtendedCoordinates) 			(OutputP ExtendedCoordinates)
    AdditionV2 		(P1, P2 ExtendedCoordinates) 			(OutputP ExtendedCoordinates)
    
    // V - Complex Point Operations Methods
    FortyNiner	 	(InputP ExtendedCoordinates) 			(OutputP ExtendedCoordinates)
    PrecomputingMatrix 	() 						[7][7]ExtendedCoordinates

    // VI - Key Generation Methods
    GetRandomOnCurve 	() 						*big.Int
    ScalarMultiplier	(Scalar *big.Int, InputP AffineCoordinates)	(OutputP AffineCoordinates)
    GetKeys 		() 						(Keys CPKeyPair)
    PrivKey2PubKey	(PrivateKey string) 				(PublicKey string)
    PrivKeyInt2PubKey 	(PrivateKeyInt *big.Int) 			(PublicKey string)

    // VII - Schnorr Signature Methods
    CPSchnorrSign 	(Keys CPKeyPair, Hash []byte) 			(Signature Schnurr)
    CPSchnorrVerify 	(Sigma Schnurr, PublicKey string, Hash []byte)	bool
}

// TwistedEdwardsCurve Methods
// I - Mod P Methods
func (k *FiniteFieldEllipticCurve) AddMod (a,b *big.Int) *big.Int{
    var result = new(big.Int)
    result.Add(a,b)
    result.Mod(result,&k.P)
    return result
}

func (k *FiniteFieldEllipticCurve) SubMod (a,b *big.Int) *big.Int{
    var result = new(big.Int)
    result.Sub(a,b)
    result.Mod(result,&k.P)
    return result
}

func (k *FiniteFieldEllipticCurve) MulMod (a,b *big.Int) *big.Int{
    var result = new(big.Int)
    result.Mul(a,b)
    result.Mod(result,&k.P)
    return result
}

// III - Boolean Methods
// 7
func (k *FiniteFieldEllipticCurve) IsInfinityPoint (InputP ExtendedCoordinates) bool {
    var result bool

    Cmp1 := InputP.EX.Cmp(Zero)
    Cmp2 := InputP.EY.Cmp(One)
    Cmp3 := InputP.EZ.Cmp(Zero)
    Cmp4 := InputP.ET.Cmp(Zero)

    if Cmp1 == 0 && Cmp2 == 0 && Cmp3 == 0 && Cmp4 == 0 {
        result = true
    } else {
        result = false
    }
    return result
}

// 8
func (k *FiniteFieldEllipticCurve) IsInverseOnCurve (P1, P2 ExtendedCoordinates) bool {
    var result bool
    Point1Aff := k.Extended2Affine(P1)
    Point2Aff := k.Extended2Affine(P2)
    Cmp := Point1Aff.AX.Cmp(Point2Aff.AX)

    if Cmp == 0  {
	result = true
    } else {
	result = false
    }
    return result
}

// 9
func (k *FiniteFieldEllipticCurve) IsOnCurve (InputP ExtendedCoordinates) bool {
    var (
	result bool
	PointAffine = k.Extended2Affine(InputP)
	A 	= new(big.Int)
	B 	= new(big.Int)
    )

    A.Exp(PointAffine.AX,Two,&k.P)
    B.Exp(PointAffine.AY,Two,&k.P)
    Left := k.AddMod(A,B)
    C := k.MulMod(A,B)
    D := k.MulMod(C,&k.D)
    Right := k.AddMod(One,D)
    //fmt.Println("Left  is",Left)
    //fmt.Println("Right is",Right)
    CompareResult := Left.Cmp(Right)
    if CompareResult == 0 {
        result = true
    } else {
        result = false
    }
    return result
}

// IV - Basic Point Operations Methods
// 10
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
func (k *FiniteFieldEllipticCurve) DoubleWithZOne (InputP ExtendedCoordinates) (OutputP ExtendedCoordinates) {
    var (
	A 	= new(big.Int)
	B 	= new(big.Int)
	v1	= new(big.Int)
    )
    //When Z is one, Point cant be Infinity-Point, therefore no check is needed.
    A.Exp(InputP.EX,Two,&k.P)
    B.Exp(InputP.EY,Two,&k.P)
    E := k.MulMod(Two,InputP.ET)
    G := k.AddMod(A,B)
    F := k.SubMod(G,Two)
    H := k.SubMod(A,B)
    OutputP.EX = k.MulMod(E,F)
    OutputP.EY = k.MulMod(G,H)
    v1.Exp(G,Two,&k.P)
    v2 := k.MulMod(Two,G)
    OutputP.EZ = k.SubMod(v1,v2)
    OutputP.ET = k.MulMod(E,H)
    return OutputP
}

//11
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
func (k *FiniteFieldEllipticCurve) Double (InputP ExtendedCoordinates) (OutputP ExtendedCoordinates) {
    var (
	A 	= new(big.Int)
	B 	= new(big.Int)
	C 	= new(big.Int)
	v1	= new(big.Int)
    )

    if k.IsInfinityPoint(InputP) == true {
	OutputP = InfinityPoint
    } else {
	A.Exp(InputP.EX,Two,&k.P)
	B.Exp(InputP.EY,Two,&k.P)
	C.Exp(InputP.EZ,Two,&k.P)
	C = k.MulMod(C,Two)
	v2 := k.AddMod(InputP.EX,InputP.EY)
	v1.Exp(v2,Two,&k.P)
	G := k.AddMod(A,B)
	E := k.SubMod(v1,G)
	F := k.SubMod(G,C)
	H := k.SubMod(A,B)
	OutputP.EX = k.MulMod(E,F)
	OutputP.EY = k.MulMod(G,H)
	OutputP.EZ = k.MulMod(F,G)
	OutputP.ET = k.MulMod(E,H)
    }
    return OutputP
}

//12
// Formulas used are these: http://hyperelliptic.org/EFD/g1p/data/twisted/extended/tripling/tpl-2015-c
// Source 2015 Chuengsatiansup
// Standard Elliptic Curve in Twisted Edward form: a * x^2 + y^2 = 1 +      d * x^2 * y^2
// Since E521 Equation is:			       x^2 + y^2 = 1 - 376014 * x^2 * y^2,
// It is assumed that a = 1, and XX is used instead of aXX
func (k *FiniteFieldEllipticCurve) Triple (InputP ExtendedCoordinates) (OutputP ExtendedCoordinates) {
    var (
	YY	= new(big.Int)
	XX	= new(big.Int)
	ZZ	= new(big.Int)
    )

    if k.IsInfinityPoint(InputP) == true {
	OutputP = InfinityPoint
    } else {
	YY.Exp(InputP.EY, Two,&k.P)
	XX.Exp(InputP.EX, Two,&k.P)
	ZZ.Exp(InputP.EZ, Two,&k.P)
	Ap := k.AddMod(YY,XX)
	twoZZ := k.MulMod(Two,ZZ)
	Diff1 := k.SubMod(twoZZ,Ap)
	B := k.MulMod(Two,Diff1)
	xB := k.MulMod(XX,B)
	yB := k.MulMod(YY,B)
	Diff2 := k.SubMod(YY,XX)
	AA := k.MulMod(Ap,Diff2)
	F := k.SubMod(AA,yB)
	G := k.AddMod(AA,xB)
	Sum := k.AddMod(yB,AA)
	xE := k.MulMod(InputP.EX,Sum)
	Diff3 := k.SubMod(xB,AA)
	yH := k.MulMod(InputP.EY,Diff3)
	zF := k.MulMod(InputP.EZ,F)
	zG := k.MulMod(InputP.EZ,G)
	OutputP.EX = k.MulMod(xE,zF)
	OutputP.EY = k.MulMod(zG,yH)
	OutputP.EZ = k.MulMod(zF,zG)
	OutputP.ET = k.MulMod(xE,yH)
    }
    return OutputP
}

//13
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
func (k *FiniteFieldEllipticCurve) AdditionZ2OneV2 (P1, P2 ExtendedCoordinates) (OutputP ExtendedCoordinates) {
    //When Z2 is one, Point2 cant be an InfinityPoint, therefore only a check for Point1 is needed.
    //Also a test is made if Point1 and Point2 are inverse to one another. In this case the sum is the point at Infinity
    //Addition is defined as below when one point is InfinityPoint or when both points are inverse.
    //These are the special cases when added two inverse points or when adding with InfinityPoint

    if k.IsInfinityPoint(P1) == true {
	OutputP = P2
    } else if k.IsInverseOnCurve(P1,P2) == true {
	OutputP = InfinityPoint
    } else {
	A := k.MulMod(P1.EX,P2.EX)
	B := k.MulMod(P1.EY,P2.EY)
	dC:= k.MulMod(P2.ET,&k.D)
	C := k.MulMod(P1.ET,dC)
	v1 := k.AddMod(P1.EX,P1.EY)
	v2 := k.AddMod(P2.EX,P2.EY)
	v3 := k.AddMod(A,B)
	v4 := k.MulMod(v1,v2)
	E := k.SubMod(v4,v3)
	F := k.SubMod(P1.EZ,C)
	G := k.AddMod(P1.EZ,C)
	H := k.SubMod(B,A)
	OutputP.EX = k.MulMod(E,F)
	OutputP.EY = k.MulMod(G,H)
	OutputP.EZ = k.MulMod(F,G)
	OutputP.ET = k.MulMod(E,H)
    }
    return OutputP
}

//14
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
func (k *FiniteFieldEllipticCurve) AdditionV2 (P1, P2 ExtendedCoordinates) (OutputP ExtendedCoordinates) {
    //Both Points are tested if they are InfinityPoints, or if the are inverse to one another.
    //Addition is defined as below when one point is InfinityPoint or when both points are inverse.
    //These are the special cases when added two inverse points or when adding with InfinityPoint

    if k.IsInfinityPoint(P1) == true {
	OutputP = P2
    } else if k.IsInfinityPoint(P2) == true {
	OutputP = P1
    } else if k.IsInfinityPoint(P1) == true && k.IsInfinityPoint(P2) == true {
	OutputP = InfinityPoint
    } else if k.IsInverseOnCurve(P1,P2) == true {
	OutputP = InfinityPoint
    } else {
	A := k.MulMod(P1.EX,P2.EX)
	B := k.MulMod(P1.EY,P2.EY)
	dC:= k.MulMod(P2.ET,&k.D)
	C := k.MulMod(P1.ET,dC)
	D := k.MulMod(P1.EZ,P2.EZ)
	v1:= k.AddMod(P1.EX,P1.EY)
	v2:= k.AddMod(P2.EX,P2.EY)
	v3:= k.AddMod(A,B)
	v4:= k.MulMod(v1,v2)
	E := k.SubMod(v4,v3)
	F := k.SubMod(D,C)
	G := k.AddMod(D,C)
	H := k.SubMod(B,A)
	OutputP.EX = k.MulMod(E,F)
	OutputP.EY = k.MulMod(G,H)
	OutputP.EZ = k.MulMod(F,G)
	OutputP.ET = k.MulMod(E,H)
    }
    return OutputP
}

// V - Complex Point Operations Methods
// 15
func (k *FiniteFieldEllipticCurve) FortyNiner (InputP ExtendedCoordinates) (OutputP ExtendedCoordinates) {
    Point03 := k.Triple(InputP)
    Point06 := k.Double(Point03)
    Point12 := k.Double(Point06)
    Point24 := k.Double(Point12)
    Point48 := k.Double(Point24)
    Point49 := k.AdditionV2(Point48,InputP)

    return Point49
}

// 16
func (k *FiniteFieldEllipticCurve) PrecomputingMatrix (InputP AffineCoordinates) [7][7]ExtendedCoordinates {
    //start := time.Now()
    //Generator := AffineCoordinates {&k.PBX, &k.PBY}

    BasePointExt	:= k.Affine2Extended(InputP)
    BasePointExt02 	:= k.DoubleWithZOne(BasePointExt)
    BasePointExt03 	:= k.AdditionZ2OneV2(BasePointExt02,BasePointExt)
    BasePointExt04 	:= k.Double(BasePointExt02)
    BasePointExt05 	:= k.AdditionZ2OneV2(BasePointExt04,BasePointExt)
    BasePointExt06	:= k.Double(BasePointExt03)
    BasePointExt07	:= k.AdditionZ2OneV2(BasePointExt06,BasePointExt)
    BasePointExt08	:= k.Double(BasePointExt04)
    BasePointExt09	:= k.AdditionZ2OneV2(BasePointExt08,BasePointExt)
    BasePointExt10	:= k.Double(BasePointExt05)
    BasePointExt11	:= k.AdditionZ2OneV2(BasePointExt10,BasePointExt)
    BasePointExt12	:= k.Double(BasePointExt06)
    BasePointExt13	:= k.AdditionZ2OneV2(BasePointExt12,BasePointExt)
    BasePointExt14	:= k.Double(BasePointExt07)
    BasePointExt15	:= k.AdditionZ2OneV2(BasePointExt14,BasePointExt)
    BasePointExt16	:= k.Double(BasePointExt08)
    BasePointExt17	:= k.AdditionZ2OneV2(BasePointExt16,BasePointExt)
    BasePointExt18	:= k.Double(BasePointExt09)
    BasePointExt19	:= k.AdditionZ2OneV2(BasePointExt18,BasePointExt)
    BasePointExt20	:= k.Double(BasePointExt10)
    BasePointExt21	:= k.AdditionZ2OneV2(BasePointExt20,BasePointExt)
    BasePointExt22	:= k.Double(BasePointExt11)
    BasePointExt23	:= k.AdditionZ2OneV2(BasePointExt22,BasePointExt)
    BasePointExt24	:= k.Double(BasePointExt12)
    BasePointExt25	:= k.AdditionZ2OneV2(BasePointExt24,BasePointExt)
    BasePointExt26	:= k.Double(BasePointExt13)
    BasePointExt27	:= k.AdditionZ2OneV2(BasePointExt26,BasePointExt)
    BasePointExt28	:= k.Double(BasePointExt14)
    BasePointExt29	:= k.AdditionZ2OneV2(BasePointExt28,BasePointExt)
    BasePointExt30	:= k.Double(BasePointExt15)
    BasePointExt31	:= k.AdditionZ2OneV2(BasePointExt30,BasePointExt)
    BasePointExt32	:= k.Double(BasePointExt16)
    BasePointExt33	:= k.AdditionZ2OneV2(BasePointExt32,BasePointExt)
    BasePointExt34	:= k.Double(BasePointExt17)
    BasePointExt35	:= k.AdditionZ2OneV2(BasePointExt34,BasePointExt)
    BasePointExt36	:= k.Double(BasePointExt18)
    BasePointExt37	:= k.AdditionZ2OneV2(BasePointExt36,BasePointExt)
    BasePointExt38	:= k.Double(BasePointExt19)
    BasePointExt39	:= k.AdditionZ2OneV2(BasePointExt38,BasePointExt)
    BasePointExt40	:= k.Double(BasePointExt20)
    BasePointExt41	:= k.AdditionZ2OneV2(BasePointExt40,BasePointExt)
    BasePointExt42	:= k.Double(BasePointExt21)
    BasePointExt43	:= k.AdditionZ2OneV2(BasePointExt42,BasePointExt)
    BasePointExt44	:= k.Double(BasePointExt22)
    BasePointExt45	:= k.AdditionZ2OneV2(BasePointExt44,BasePointExt)
    BasePointExt46	:= k.Double(BasePointExt23)
    BasePointExt47	:= k.AdditionZ2OneV2(BasePointExt46,BasePointExt)
    BasePointExt48	:= k.Double(BasePointExt24)
    BasePointExt49	:= k.AdditionZ2OneV2(BasePointExt48,BasePointExt)
    //Point49 isn't used in the pre-computation, it is only created to fill the Matrix.

    MatrixRow0 := [...]ExtendedCoordinates{BasePointExt,BasePointExt02,BasePointExt03,BasePointExt04,BasePointExt05,BasePointExt06,BasePointExt07}
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

// VI - Key Generation Methods
// 17
func (k *FiniteFieldEllipticCurve) GetRandomOnCurve () *big.Int {
    RandomNumber,_ := rand.Int(rand.Reader,&k.P)
    return RandomNumber
}

// 18
func (k *FiniteFieldEllipticCurve) ScalarMultiplier (Scalar *big.Int, InputP AffineCoordinates) (OutputP AffineCoordinates) {
    var (
	//start = time.Now()
	PrivKey49 		= Scalar.Text(49)
	PrivKey49SliceRune 	= []rune(PrivKey49)
	PrivKey49SliceString 	= make([]string,len(PrivKey49))
	ZeroPoint		= InfinityPoint
	Result			ExtendedCoordinates
    )
    //start := time.Now()
    for i := 0; i < len(PrivKey49); i++ {
	PrivKey49SliceString[i] = string(PrivKey49SliceRune[i])
    }
    Result = ZeroPoint
    PrecMatrix := k.PrecomputingMatrix(InputP)
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
    OutputP = k.Extended2Affine(Result)

    //elapsed := time.Since(start)
    //fmt.Println("")
    //fmt.Println("Computing PublicKey points took:", elapsed)
    return OutputP
}