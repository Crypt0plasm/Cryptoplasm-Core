package Cryptoplasm_Elliptic

import (
    "fmt"
    "math/big"
)

var (
    CurveE521 = ParamE521()
    CurveE521InfinityPoint = TecExtendedCoordinates {big.NewInt(0), big.NewInt(1),Zero,Zero}
    CurveE521BasePoint = TecAffineCoordinates {&CurveE521.PBX, &CurveE521.PBY}
    CurveE521Prime = &CurveE521.P
    CurveE521Constant = &CurveE521.D
)

func IsInfinityPoint (Point TecExtendedCoordinates) bool {
    var (
    	result bool
	XInt64 = Point.EX.Int64()
	YInt64 = Point.EY.Int64()
	ZInt64 = Point.EZ.Int64()
	TInt64 = Point.ET.Int64()
    )
    //Conversion to int64 is needed as big.Int isn't comparable
    //Overflow doesnt matter since only Zero and One is searched for.

    if XInt64 == 0 && YInt64 == 1 && ZInt64 == 0 && TInt64 == 0 {
        result = true
    } else {
        result = false
    }
    return result
}

func IsInverseOnCurveE521 (Point1, Point2 TecExtendedCoordinates) bool {
    var (
	result bool
	Point1Aff = TecExtended2TecAffine(Point1,CurveE521Prime)
	Point2Aff = TecExtended2TecAffine(Point2,CurveE521Prime)
	X1Int64 = Point1Aff.AX.Int64()
	X2Int64 = Point2Aff.AX.Int64()
    )
    //Conversion to int64 is needed as big.Int isn't comparable
    //Overflow doesnt matter since an equality test is made.

    if X1Int64 == X2Int64  {
	result = true
    } else {
	result = false
    }
    return result
}

func MulMod (a,b,prime *big.Int) *big.Int{
    var result = new(big.Int)
    result.Mul(a,b)
    result.Mod(result,prime)
    return result
}

func AddMod (a,b,prime *big.Int) *big.Int{
    var result = new(big.Int)
    result.Add(a,b)
    result.Mod(result,prime)
    return result
}

func SubMod (a,b,prime *big.Int) *big.Int{
    var result = new(big.Int)
    result.Sub(a,b)
    result.Mod(result,prime)
    return result
}

func IsOnCurveE521 (Point TecExtendedCoordinates) bool {
    var (
	PointAffine = TecExtended2TecAffine(Point,CurveE521Prime)
	CompareResult int
	result bool
	A 	= new(big.Int)
	B 	= new(big.Int)
    )

    A.Exp(PointAffine.AX, Two,CurveE521Prime)
    B.Exp(PointAffine.AY, Two,CurveE521Prime)
    Left := AddMod(A,B,CurveE521Prime)
    C := MulMod(A,B,CurveE521Prime)
    D := MulMod(C,CurveE521Constant,CurveE521Prime)
    Right := AddMod(One,D,CurveE521Prime)
    //fmt.Println("Left  is",Left)
    //fmt.Println("Right is",Right)
    CompareResult = Left.Cmp(Right)
    if CompareResult == 0 {
        result = true
    } else {
        result = false
    }

    return result
}

func InversePointOnCurveE521 (Point TecExtendedCoordinates) TecExtendedCoordinates {
    var (
        PointAff = TecExtended2TecAffine(Point,CurveE521Prime)
        ResultAff TecAffineCoordinates
        ResultExt TecExtendedCoordinates
	y 	= new(big.Int)
    )

    y.Sub(Zero,PointAff.AY)
    ResultAff.AX = PointAff.AX
    ResultAff.AY = y
    ResultExt = TecAffine2TecExtended(ResultAff,CurveE521Prime)
    return ResultExt
}

func PrecomputingE521 () [7][7]TecExtendedCoordinates {
    //start := time.Now()

    BasePointExt	:= TecAffine2TecExtended(CurveE521BasePoint,CurveE521Prime)
    BasePointExt02 	:= DoubleWithZOne(BasePointExt,CurveE521Prime)
    BasePointExt03 	:= AdditionZ2OneV2(BasePointExt02,BasePointExt,CurveE521Prime,CurveE521Constant)
    BasePointExt04 	:= Double(BasePointExt02,CurveE521Prime)
    BasePointExt05 	:= AdditionZ2OneV2(BasePointExt04,BasePointExt,CurveE521Prime,CurveE521Constant)
    BasePointExt06	:= Double(BasePointExt03,CurveE521Prime)
    BasePointExt07	:= AdditionZ2OneV2(BasePointExt06,BasePointExt,CurveE521Prime,CurveE521Constant)
    BasePointExt08	:= Double(BasePointExt04,CurveE521Prime)
    BasePointExt09	:= AdditionZ2OneV2(BasePointExt08,BasePointExt,CurveE521Prime,CurveE521Constant)
    BasePointExt10	:= Double(BasePointExt05,CurveE521Prime)
    BasePointExt11	:= AdditionZ2OneV2(BasePointExt10,BasePointExt,CurveE521Prime,CurveE521Constant)
    BasePointExt12	:= Double(BasePointExt06,CurveE521Prime)
    BasePointExt13	:= AdditionZ2OneV2(BasePointExt12,BasePointExt,CurveE521Prime,CurveE521Constant)
    BasePointExt14	:= Double(BasePointExt07,CurveE521Prime)
    BasePointExt15	:= AdditionZ2OneV2(BasePointExt14,BasePointExt,CurveE521Prime,CurveE521Constant)
    BasePointExt16	:= Double(BasePointExt08,CurveE521Prime)
    BasePointExt17	:= AdditionZ2OneV2(BasePointExt16,BasePointExt,CurveE521Prime,CurveE521Constant)
    BasePointExt18	:= Double(BasePointExt09,CurveE521Prime)
    BasePointExt19	:= AdditionZ2OneV2(BasePointExt18,BasePointExt,CurveE521Prime,CurveE521Constant)
    BasePointExt20	:= Double(BasePointExt10,CurveE521Prime)
    BasePointExt21	:= AdditionZ2OneV2(BasePointExt20,BasePointExt,CurveE521Prime,CurveE521Constant)
    BasePointExt22	:= Double(BasePointExt11,CurveE521Prime)
    BasePointExt23	:= AdditionZ2OneV2(BasePointExt22,BasePointExt,CurveE521Prime,CurveE521Constant)
    BasePointExt24	:= Double(BasePointExt12,CurveE521Prime)
    BasePointExt25	:= AdditionZ2OneV2(BasePointExt24,BasePointExt,CurveE521Prime,CurveE521Constant)
    BasePointExt26	:= Double(BasePointExt13,CurveE521Prime)
    BasePointExt27	:= AdditionZ2OneV2(BasePointExt26,BasePointExt,CurveE521Prime,CurveE521Constant)
    BasePointExt28	:= Double(BasePointExt14,CurveE521Prime)
    BasePointExt29	:= AdditionZ2OneV2(BasePointExt28,BasePointExt,CurveE521Prime,CurveE521Constant)
    BasePointExt30	:= Double(BasePointExt15,CurveE521Prime)
    BasePointExt31	:= AdditionZ2OneV2(BasePointExt30,BasePointExt,CurveE521Prime,CurveE521Constant)
    BasePointExt32	:= Double(BasePointExt16,CurveE521Prime)
    BasePointExt33	:= AdditionZ2OneV2(BasePointExt32,BasePointExt,CurveE521Prime,CurveE521Constant)
    BasePointExt34	:= Double(BasePointExt17,CurveE521Prime)
    BasePointExt35	:= AdditionZ2OneV2(BasePointExt34,BasePointExt,CurveE521Prime,CurveE521Constant)
    BasePointExt36	:= Double(BasePointExt18,CurveE521Prime)
    BasePointExt37	:= AdditionZ2OneV2(BasePointExt36,BasePointExt,CurveE521Prime,CurveE521Constant)
    BasePointExt38	:= Double(BasePointExt19,CurveE521Prime)
    BasePointExt39	:= AdditionZ2OneV2(BasePointExt38,BasePointExt,CurveE521Prime,CurveE521Constant)
    BasePointExt40	:= Double(BasePointExt20,CurveE521Prime)
    BasePointExt41	:= AdditionZ2OneV2(BasePointExt40,BasePointExt,CurveE521Prime,CurveE521Constant)
    BasePointExt42	:= Double(BasePointExt21,CurveE521Prime)
    BasePointExt43	:= AdditionZ2OneV2(BasePointExt42,BasePointExt,CurveE521Prime,CurveE521Constant)
    BasePointExt44	:= Double(BasePointExt22,CurveE521Prime)
    BasePointExt45	:= AdditionZ2OneV2(BasePointExt44,BasePointExt,CurveE521Prime,CurveE521Constant)
    BasePointExt46	:= Double(BasePointExt23,CurveE521Prime)
    BasePointExt47	:= AdditionZ2OneV2(BasePointExt46,BasePointExt,CurveE521Prime,CurveE521Constant)
    BasePointExt48	:= Double(BasePointExt24,CurveE521Prime)
    BasePointExt49	:= AdditionZ2OneV2(BasePointExt48,BasePointExt,CurveE521Prime,CurveE521Constant)
    //Point49 isn't used in the pre-computation, it is only created to fill the Matrix.

    MatrixRow0 := [...]TecExtendedCoordinates{BasePointExt,BasePointExt02,BasePointExt03,BasePointExt04,BasePointExt05,BasePointExt06,BasePointExt07}
    MatrixRow1 := [...]TecExtendedCoordinates{BasePointExt08,BasePointExt09,BasePointExt10,BasePointExt11,BasePointExt12,BasePointExt13,BasePointExt14}
    MatrixRow2 := [...]TecExtendedCoordinates{BasePointExt15,BasePointExt16,BasePointExt17,BasePointExt18,BasePointExt19,BasePointExt20,BasePointExt21}
    MatrixRow3 := [...]TecExtendedCoordinates{BasePointExt22,BasePointExt23,BasePointExt24,BasePointExt25,BasePointExt26,BasePointExt27,BasePointExt28}
    MatrixRow4 := [...]TecExtendedCoordinates{BasePointExt29,BasePointExt30,BasePointExt31,BasePointExt32,BasePointExt33,BasePointExt34,BasePointExt35}
    MatrixRow5 := [...]TecExtendedCoordinates{BasePointExt36,BasePointExt37,BasePointExt38,BasePointExt39,BasePointExt40,BasePointExt41,BasePointExt42}
    MatrixRow6 := [...]TecExtendedCoordinates{BasePointExt43,BasePointExt44,BasePointExt45,BasePointExt46,BasePointExt47,BasePointExt48,BasePointExt49}

    PrecomputingMatrix := [7][7]TecExtendedCoordinates{MatrixRow0,MatrixRow1,MatrixRow2,MatrixRow3,MatrixRow4,MatrixRow5,MatrixRow6}

    //fmt.Println("Point02 is", BasePointExt02)
    //fmt.Println("Point03 is", BasePointExt03)
    //fmt.Println("Point04 is", BasePointExt04)
    //fmt.Println("Point05 is", BasePointExt05)
    //fmt.Println("Point06 is", BasePointExt06)
    //fmt.Println("Point07 is", BasePointExt07)
    //fmt.Println("Point08 is", BasePointExt08)
    //fmt.Println("Point09 is", BasePointExt09)
    //fmt.Println("Point10 is", BasePointExt10)
    //fmt.Println("Point11 is", BasePointExt11)
    //fmt.Println("Point12 is", BasePointExt12)
    //fmt.Println("Point13 is", BasePointExt13)
    //fmt.Println("Point14 is", BasePointExt14)
    //fmt.Println("Point15 is", BasePointExt15)
    //fmt.Println("Point16 is", BasePointExt16)
    //fmt.Println("Point17 is", BasePointExt17)
    //fmt.Println("Point18 is", BasePointExt18)
    //fmt.Println("Point19 is", BasePointExt19)
    //fmt.Println("Point20 is", BasePointExt20)
    //fmt.Println("Point21 is", BasePointExt21)
    //fmt.Println("Point22 is", BasePointExt22)
    //fmt.Println("Point23 is", BasePointExt23)
    //fmt.Println("Point24 is", BasePointExt24)
    //fmt.Println("Point25 is", BasePointExt25)
    //fmt.Println("Point26 is", BasePointExt26)
    //fmt.Println("Point27 is", BasePointExt27)
    //fmt.Println("Point28 is", BasePointExt28)
    //fmt.Println("Point29 is", BasePointExt29)
    //fmt.Println("Point30 is", BasePointExt30)
    //fmt.Println("Point31 is", BasePointExt31)
    //fmt.Println("Point32 is", BasePointExt32)
    //fmt.Println("Point33 is", BasePointExt33)
    //fmt.Println("Point34 is", BasePointExt34)
    //fmt.Println("Point35 is", BasePointExt35)
    //fmt.Println("Point36 is", BasePointExt36)
    //fmt.Println("Point37 is", BasePointExt37)
    //fmt.Println("Point38 is", BasePointExt38)
    //fmt.Println("Point39 is", BasePointExt39)
    //fmt.Println("Point40 is", BasePointExt40)
    //fmt.Println("Point41 is", BasePointExt41)
    //fmt.Println("Point42 is", BasePointExt42)
    //fmt.Println("Point43 is", BasePointExt43)
    //fmt.Println("Point44 is", BasePointExt44)
    //fmt.Println("Point45 is", BasePointExt45)
    //fmt.Println("Point46 is", BasePointExt46)
    //fmt.Println("Point47 is", BasePointExt47)
    //fmt.Println("Point48 is", BasePointExt48)
    //fmt.Println("Point49 is", BasePointExt49)

    fmt.Println("")
    //fmt.Println("Starting Point Verification")
    //fmt.Println("Point02 is",IsOnCurveE521(BasePointExt02))
    //fmt.Println("Point03 is",IsOnCurveE521(BasePointExt03))
    //fmt.Println("Point04 is",IsOnCurveE521(BasePointExt04))
    //fmt.Println("Point05 is",IsOnCurveE521(BasePointExt05))
    //fmt.Println("Point06 is",IsOnCurveE521(BasePointExt06))
    //fmt.Println("Point07 is",IsOnCurveE521(BasePointExt07))
    //fmt.Println("Point08 is",IsOnCurveE521(BasePointExt08))
    //fmt.Println("Point09 is",IsOnCurveE521(BasePointExt09))
    //fmt.Println("Point10 is",IsOnCurveE521(BasePointExt10))
    //fmt.Println("Point11 is",IsOnCurveE521(BasePointExt11))
    //fmt.Println("Point12 is",IsOnCurveE521(BasePointExt12))
    //fmt.Println("Point13 is",IsOnCurveE521(BasePointExt13))
    //fmt.Println("Point14 is",IsOnCurveE521(BasePointExt14))
    //fmt.Println("Point15 is",IsOnCurveE521(BasePointExt15))
    //fmt.Println("Point16 is",IsOnCurveE521(BasePointExt16))
    //fmt.Println("Point17 is",IsOnCurveE521(BasePointExt17))
    //fmt.Println("Point18 is",IsOnCurveE521(BasePointExt18))
    //fmt.Println("Point19 is",IsOnCurveE521(BasePointExt19))
    //fmt.Println("Point20 is",IsOnCurveE521(BasePointExt20))
    //fmt.Println("Point21 is",IsOnCurveE521(BasePointExt21))
    //fmt.Println("Point22 is",IsOnCurveE521(BasePointExt22))
    //fmt.Println("Point23 is",IsOnCurveE521(BasePointExt23))
    //fmt.Println("Point24 is",IsOnCurveE521(BasePointExt24))
    //fmt.Println("Point25 is",IsOnCurveE521(BasePointExt25))
    //fmt.Println("Point26 is",IsOnCurveE521(BasePointExt26))
    //fmt.Println("Point27 is",IsOnCurveE521(BasePointExt27))
    //fmt.Println("Point28 is",IsOnCurveE521(BasePointExt28))
    //fmt.Println("Point29 is",IsOnCurveE521(BasePointExt29))
    //fmt.Println("Point30 is",IsOnCurveE521(BasePointExt30))
    //fmt.Println("Point31 is",IsOnCurveE521(BasePointExt31))
    //fmt.Println("Point32 is",IsOnCurveE521(BasePointExt32))
    //fmt.Println("Point33 is",IsOnCurveE521(BasePointExt33))
    //fmt.Println("Point34 is",IsOnCurveE521(BasePointExt34))
    //fmt.Println("Point35 is",IsOnCurveE521(BasePointExt35))
    //fmt.Println("Point36 is",IsOnCurveE521(BasePointExt36))
    //fmt.Println("Point37 is",IsOnCurveE521(BasePointExt37))
    //fmt.Println("Point38 is",IsOnCurveE521(BasePointExt38))
    //fmt.Println("Point39 is",IsOnCurveE521(BasePointExt39))
    //fmt.Println("Point40 is",IsOnCurveE521(BasePointExt40))
    //fmt.Println("Point41 is",IsOnCurveE521(BasePointExt41))
    //fmt.Println("Point42 is",IsOnCurveE521(BasePointExt42))
    //fmt.Println("Point43 is",IsOnCurveE521(BasePointExt43))
    //fmt.Println("Point44 is",IsOnCurveE521(BasePointExt44))
    //fmt.Println("Point45 is",IsOnCurveE521(BasePointExt45))
    //fmt.Println("Point46 is",IsOnCurveE521(BasePointExt46))
    //fmt.Println("Point47 is",IsOnCurveE521(BasePointExt47))
    //fmt.Println("Point48 is",IsOnCurveE521(BasePointExt48))
    //fmt.Println("Point49 is",IsOnCurveE521(BasePointExt49))

    //elapsed := time.Since(start)
    //fmt.Println("")
    //fmt.Println("Computing and verifying 47 points took", elapsed)
    return PrecomputingMatrix
}
func FortyNinerE521 (Point TecExtendedCoordinates) TecExtendedCoordinates {
    Point03 := Triple(Point,CurveE521Prime)
    Point06 := Double(Point03,CurveE521Prime)
    Point12 := Double(Point06,CurveE521Prime)
    Point24 := Double(Point12,CurveE521Prime)
    Point48 := Double(Point24,CurveE521Prime)
    Point49 := AdditionV2(Point48,Point,CurveE521Prime,CurveE521Constant)

    return Point49
}

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
func DoubleWithZOne (Point TecExtendedCoordinates, prime *big.Int) TecExtendedCoordinates {
    var (
    	result TecExtendedCoordinates
    	A 	= new(big.Int)
	B 	= new(big.Int)
	v1	= new(big.Int)
    )

    //When Z is one, Point cant be Infinity-Point, therefore no check is needed.

    A.Exp(Point.EX, Two,prime)
    B.Exp(Point.EY, Two,prime)
    E := MulMod(Two,Point.ET,prime)
    G := AddMod(A,B,prime)
    F := SubMod(G, Two,prime)
    H := SubMod(A,B,prime)
    result.EX = MulMod(E,F,prime)
    result.EY = MulMod(G,H,prime)
    v1.Exp(G, Two,prime)
    v2 := MulMod(Two,G,prime)
    result.EZ = SubMod(v1,v2,prime)
    result.ET = MulMod(E,H,prime)

    return result
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
func Double (Point TecExtendedCoordinates, prime *big.Int) TecExtendedCoordinates{
    var (
	result TecExtendedCoordinates
	A 	= new(big.Int)
	B 	= new(big.Int)
	C 	= new(big.Int)
	v1	= new(big.Int)
    )

    if IsInfinityPoint(Point) == true {
	result = CurveE521InfinityPoint
    } else {
	A.Exp(Point.EX, Two,prime)
	B.Exp(Point.EY, Two,prime)
	C.Exp(Point.EZ, Two,prime)
	C = MulMod(C, Two,prime)
	v2 := AddMod(Point.EX,Point.EY,prime)
	v1.Exp(v2, Two,prime)
	G := AddMod(A,B,prime)
	E := SubMod(v1,G,prime)
	F := SubMod(G,C,prime)
	H := SubMod(A,B,prime)

	result.EX = MulMod(E,F,prime)
	result.EY = MulMod(G,H,prime)
	result.EZ = MulMod(F,G,prime)
	result.ET = MulMod(E,H,prime)
    }

    return result
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
func AdditionZ2OneV2 (Point1, Point2 TecExtendedCoordinates, prime, CurveConstant *big.Int) TecExtendedCoordinates {
    var result TecExtendedCoordinates

    //When Z2 is one, Point2 cant be an InfinityPoint, therefore only a check for Point1 is needed.
    //Also a test is made if Point1 and Point2 are inverse to one another. In this case the sum is the point at Infinity
    //Addition is defined as below when one point is InfinityPoint or when both points are inverse.
    //These are the special cases when added two inverse points or when adding with InfinityPoint

    if IsInfinityPoint(Point1) == true {
	result = Point2
    } else if IsInverseOnCurveE521(Point1,Point2) == true {
        result = CurveE521InfinityPoint
    } else {
	A := MulMod(Point1.EX,Point2.EX,prime)
	B := MulMod(Point1.EY,Point2.EY,prime)
	dC:= MulMod(Point2.ET,CurveConstant,prime)
	C := MulMod(Point1.ET,dC,prime)
	v1 := AddMod(Point1.EX,Point1.EY,prime)
	v2 := AddMod(Point2.EX,Point2.EY,prime)
	v3 := AddMod(A,B,prime)
	v4 := MulMod(v1,v2,prime)
	E := SubMod(v4,v3,prime)
	F := SubMod(Point1.EZ,C,prime)
	G := AddMod(Point1.EZ,C,prime)
	H := SubMod(B,A,prime)
	result.EX = MulMod(E,F,prime)
	result.EY = MulMod(G,H,prime)
	result.EZ = MulMod(F,G,prime)
	result.ET = MulMod(E,H,prime)
    }

    return result
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
func AdditionV2 (Point1, Point2 TecExtendedCoordinates, prime, CurveConstant *big.Int) TecExtendedCoordinates {
    var result TecExtendedCoordinates

    //Both Points are tested if they are InfinityPoints, or if the are inverse to one another.
    //Addition is defined as below when one point is InfinityPoint or when both points are inverse.
    //These are the special cases when added two inverse points or when adding with InfinityPoint

    if IsInfinityPoint(Point1) == true {
        result = Point2
	//fmt.Println("First point is zero")
    } else if IsInfinityPoint(Point2) == true {
	//fmt.Println("Second point is zero")
        result = Point1
    } else if IsInfinityPoint(Point1) == true && IsInfinityPoint(Point2) == true {
        result = CurveE521InfinityPoint
	//fmt.Println("Both points zero")
    } else if IsInverseOnCurveE521(Point1,Point2) == true {
	result = CurveE521InfinityPoint
    } else {
        //fmt.Println("We are here")
	A := MulMod(Point1.EX,Point2.EX,prime)
	B := MulMod(Point1.EY,Point2.EY,prime)
	dC:= MulMod(Point2.ET,CurveConstant,prime)
	C := MulMod(Point1.ET,dC,prime)
	D := MulMod(Point1.EZ,Point2.EZ,prime)
	v1:= AddMod(Point1.EX,Point1.EY,prime)
	v2:= AddMod(Point2.EX,Point2.EY,prime)
	v3:= AddMod(A,B,prime)
	v4:= MulMod(v1,v2,prime)
	E := SubMod(v4,v3,prime)
	F := SubMod(D,C,prime)
	G := AddMod(D,C,prime)
	H := SubMod(B,A,prime)
	result.EX = MulMod(E,F,prime)
	result.EY = MulMod(G,H,prime)
	result.EZ = MulMod(F,G,prime)
	result.ET = MulMod(E,H,prime)
    }

    return result
}

// Formulas used are these: http://hyperelliptic.org/EFD/g1p/data/twisted/extended/tripling/tpl-2015-c
// Source 2015 Chuengsatiansup
// Standard Elliptic Curve in Twisted Edward form: a * x^2 + y^2 = 1 +      d * x^2 * y^2
// Since E521 Equation is:			       x^2 + y^2 = 1 - 376014 * x^2 * y^2,
// It is assumed that a = 1, and XX is used instead of aXX
func Triple (Point TecExtendedCoordinates, prime *big.Int) TecExtendedCoordinates {
    var (
	result TecExtendedCoordinates
	YY	= new(big.Int)
	XX	= new(big.Int)
	ZZ	= new(big.Int)
    )

    if IsInfinityPoint(Point) == true {
	result = CurveE521InfinityPoint
    } else {
	YY.Exp(Point.EY, Two,prime)
	XX.Exp(Point.EX, Two,prime)
	ZZ.Exp(Point.EZ, Two,prime)
	Ap := AddMod(YY,XX,prime)
	twoZZ := MulMod(Two,ZZ,prime)
	Diff1 := SubMod(twoZZ,Ap,prime)
	B := MulMod(Two,Diff1,prime)
	xB := MulMod(XX,B,prime)
	yB := MulMod(YY,B,prime)
	Diff2 := SubMod(YY,XX,prime)
	AA := MulMod(Ap,Diff2,prime)
	F := SubMod(AA,yB,prime)
	G := AddMod(AA,xB,prime)
	Sum := AddMod(yB,AA,prime)
	xE := MulMod(Point.EX,Sum,prime)
	Diff3 := SubMod(xB,AA,prime)
	yH := MulMod(Point.EY,Diff3,prime)
	zF := MulMod(Point.EZ,F,prime)
	zG := MulMod(Point.EZ,G,prime)
	result.EX = MulMod(xE,zF,prime)
	result.EY = MulMod(zG,yH,prime)
	result.EZ = MulMod(zF,zG,prime)
	result.ET = MulMod(xE,yH,prime)
    }

    return result
}