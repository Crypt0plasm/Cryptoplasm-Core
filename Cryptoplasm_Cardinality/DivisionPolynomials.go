package Cryptoplasm_Cardinality

import (
    aux "Cryptoplasm-Core/Auxiliary"
    "math/big"
    "strconv"
    "strings"
)
//======================================================================================================================
//
// PART I - Type Definitions
//
type Letter struct {
    Letter string
    Exponent *big.Int
}

type Coefficient struct {
    Numeral *big.Int
    ACoeff Letter
    BCoeff Letter
}

type YCoefficient struct {
    Numeral *big.Int
    NumeralExponent *big.Int
    YCoeff Letter
}

type Polynom struct {
    Degree uint64
    Rank []uint64
    Coefficients [][]Coefficient
}

type DivisionPolynom struct {
    NonY Polynom
    Y YCoefficient
}
//======================================================================================================================
//
// Part II - Printing Functions
//
// II.1 PrintDPolynom prints a Division-Polynom
//
func PrintDPolynom (P DivisionPolynom) string {
    var Result string
    P1 := PrintPolynom(P.NonY,"Short")
    P2 := PrintYCoefficient(P.Y)
    if P2 == "1" {
	Result = P1
    }else if P1 == "1" {
	Result = P2
    } else {
	Result = "(" + P1 + ")*" + P2
    }
    return Result
}
//
// II.2 PrintPolynom prints a simple AB-Polynom; The AB-Polynom is the nonY Part of the Division Polynom
//
func PrintPolynom (P Polynom, Mode string) string {
    var (
	Result string
	StringSlice1 = make([]string, len(P.Rank))
	ReverseSlice = make([]string, len(P.Rank))
	X string
    )

    if Mode == "Long" {
	for i:=0; i<len(P.Rank); i++ {
	    X = "X^"+strconv.Itoa(int(P.Rank[i]))     //strconv.Itoa(int(P.Rank[i]))
	    StringSlice1[i] = PrintCoefficients(P.Coefficients[i]) + X
	}
	ReverseSlice = aux.ReverseStringSlice(StringSlice1)
	for i:=0; i<len(ReverseSlice); i++ {
	    if i == 0 {
		Result = ReverseSlice[i]
	    } else if len(ReverseSlice) > 1 && ReverseSlice[i][:1] == "-" {
		Result = Result + ReverseSlice[i]
	    } else {
		Result = Result + "+" + ReverseSlice[i]
	    }
	}
    } else if Mode == "Short" {
	for i:=0; i<len(P.Rank); i++ {
	    if P.Rank[i] == 0 {
		X = ""
	    } else if P.Rank[i] == 1 {
		X = "X"
	    } else {
		X = "X^"+strconv.Itoa(int(P.Rank[i]))
	    }
	    if PrintCoefficients(P.Coefficients[i]) == "0" {
		StringSlice1[i] = "."
	    } else {
		StringSlice1[i] = PrintCoefficients(P.Coefficients[i]) + X
	    }
	}
	ReverseSlice = aux.ReverseStringSlice(StringSlice1)
	for i:=0; i<len(ReverseSlice); i++ {
	    if i == 0 {
		Result = ReverseSlice[i]
	    }else if len(ReverseSlice) > 1 && ReverseSlice[i][:1] == "-" {
		Result = Result + ReverseSlice[i]
	    } else {
		Result = Result + "+" + ReverseSlice[i]
	    }
	}
    }
    FinalResult := strings.ReplaceAll(Result,"+.","")
    return FinalResult
}
//
// II.3 PrintCoefficients prints a Chain of Coefficients
//
func PrintCoefficients (C []Coefficient) string {
    var (
	Result string
	StringSlice = make([]string, len(C))
	Size bool
    )
    for i:=0; i<len(C); i++ {
	StringSlice[i] = PrintCoefficient(C[i])
    }

    m := make(map[string]bool)
    for i:=0; i<len(StringSlice); i++ {
        if StringSlice[i] != "0" {
            m[StringSlice[i]] = true
	}
    }

    for i:=0; i<len(StringSlice); i++ {
        if m[StringSlice[i]] == true {
	    if i == 0 {
		Result = StringSlice[i]
	    } else {
	        if Result == "" {
	            Result = StringSlice[i]
		} else {
		    if StringSlice[i][:1] == "-" {
			Result = Result + StringSlice[i]
		    } else {
			Result = Result + "+" + StringSlice[i]
		    }

		    Size = true
		}
	    }
	} else if m[StringSlice[i]] == false && len(StringSlice) == 1 {
	    Result = "0"
	}
    }

    if Size == true {
	Result = "(" + Result + ")"
    }
    return Result
}
//
// II.4 PrintYCoefficient prints the YCoefficient, the Y Polynom Part of the DivisionPolynom.
//
func PrintYCoefficient (C YCoefficient) string {
    var (
        Y string
        Number string
	Result string
    )

    if C.NumeralExponent.Cmp(DPZero) == 0 {
	Number = "1"
    } else if C.NumeralExponent.Cmp(DPOne) == 0 {
	Number = C.Numeral.Text(10)
    } else {
	Number = C.Numeral.Text(10)+"^"+C.NumeralExponent.Text(10)
    }

    if C.YCoeff.Exponent.Cmp(DPZero) == 0 {
	Y = "1"
    } else if C.YCoeff.Exponent.Cmp(DPOne) == 0 {
	Y = "Y"
    } else {
	Y = "Y^"+C.YCoeff.Exponent.Text(10)
    }

    if Number == "1" && Y == "1" {
	Result = "1"
    } else {
	Result = Number + Y
    }
    return Result
}
//
// II.5 PrintCoefficient prints a single Coefficient.
//
func PrintCoefficient (C Coefficient) string {
    var (
	Result string
	NS, AS, BS string
    )
    S1 := C.Numeral.Text(10)
    S2 := C.ACoeff.Exponent.Text(10)
    S3 := C.BCoeff.Exponent.Text(10)
    if S1 == "0" && S2 == "0" && S3 == "0" {
    	Result = "0"
    } else if S1 == "0" {
	Result = "0"
    } else if S1 == "1" {
	NS = "1"
	if S2 == "0" {
	    AS = ""
	    if S3 == "0" {
		BS = ""
		Result = NS + AS + BS
	    } else if S3 == "1" {
		BS = "B"
		Result = NS + AS + BS
	    } else {
		BS = "B^" + S3
		Result = NS + AS + BS
	    }
	} else if S2 == "1" {
	    AS = "A"
	    if S3 == "0" {
		BS = ""
		Result = NS + AS + BS
	    } else if S3 == "1" {
		BS = "B"
		Result = NS + AS + BS
	    } else {
		BS = "B^" + S3
		Result = NS + AS + BS
	    }
	} else {
	    AS = "A^" + S2
	    if S3 == "0" {
		BS = ""
		Result = NS + AS + BS
	    } else if S3 == "1" {
		BS = "B"
		Result = NS + AS + BS
	    } else {
		BS = "B^" + S3
		Result = NS + AS + BS
	    }
	}
    } else {
	NS = S1
	if S2 == "0" {
	    AS = ""
	    if S3 == "0" {
		BS = ""
		Result = NS + AS + BS
	    } else if S3 == "1" {
		BS = "B"
		Result = NS + AS + BS
	    } else {
		BS = "B^" + S3
		Result = NS + AS + BS
	    }
	} else if S2 == "1" {
	    AS = "A"
	    if S3 == "0" {
		BS = ""
		Result = NS + AS + BS
	    } else if S3 == "1" {
		BS = "B"
		Result = NS + AS + BS
	    } else {
		BS = "B^" + S3
		Result = NS + AS + BS
	    }
	} else {
	    AS = "A^" + S2
	    if S3 == "0" {
		BS = ""
		Result = NS + AS + BS
	    } else if S3 == "1" {
		BS = "B"
		Result = NS + AS + BS
	    } else {
		BS = "B^" + S3
		Result = NS + AS + BS
	    }
	}
    }
    return Result
}
//======================================================================================================================
//
// Part III - Auxiliary Function
//
// III.1 - Function RemoveCoefficient remove a Coefficient from the specified Position
//
func RemoveCoefficient (Chain []Coefficient, Position int) []Coefficient {
    copy(Chain[Position:], Chain[Position+1:])
    Chain = Chain[:len(Chain)-1]
    return Chain
}
//
// III.2 - Function AppendCoefficient appends a Coefficient to a Coefficient-Chain.
//
func AppendCoefficient (Chain []Coefficient, ToAppend Coefficient) (Result []Coefficient) {
    Result = append(Chain,ToAppend)
    return Result
}
//
// III.2 - Function AppendChain appends a Coefficient-Chain to a Coefficient-Chain of Chains.
//
func AppendChain(Chain [][]Coefficient, ToAppend []Coefficient) (Result [][]Coefficient) {
    Result = append(Chain,ToAppend)
    return Result
}
//
// III.3 - Function AddCoefficientChains adds two Coefficients Chains into one, by appending them together
//
func AddCoefficientChains (C1, C2 []Coefficient) (Result []Coefficient) {
    IntermediaryResult := append(C1,C2...)
    Result = ReduceCoefficients(IntermediaryResult)
    return Result
}
//
// III.4 - Function AddCoefficientChains adds two Coefficients Chains into one, by appending them together
//
func AppendDPolynom(DP []DivisionPolynom, ToAppend DivisionPolynom) (Result []DivisionPolynom) {
    Result = append(DP,ToAppend)
    return Result
}
//
// III.4 - Function ReduceCoefficients reduces multiple similar Coefficient values, if they exist
// (values that have the same AB exponent), to a single value.
//
func ReduceCoefficients (ChainLink []Coefficient) []Coefficient {
    var (
	FinalChain []Coefficient
    )

    if len(ChainLink) == 1 {
	FinalChain = ChainLink
    } else {
	Chain := ChainLink
	for ok := true; ok; ok = !(len(Chain) == 0) {
	    ChainLength := len(Chain)
	    BoolChain := make([]bool, ChainLength)
	    BoolChain[0] = true
	    for i:=0; i<len(BoolChain)-1; i++ {
		Cmp1 := Chain[0].ACoeff.Exponent.Cmp(Chain[i+1].ACoeff.Exponent)
		Cmp2 := Chain[0].BCoeff.Exponent.Cmp(Chain[i+1].BCoeff.Exponent)
		if Cmp1 == 0 && Cmp2 == 0 {
		    BoolChain[i+1] = true
		}
	    }
	    var (
		Sum = big.NewInt(0)
		Element Coefficient
		RemovalCount = 0
	    )
	    for j:=0; j<len(BoolChain); j++ {
		if BoolChain[j] == true {
		    Element = Chain[j-RemovalCount]
		    Sum.Add(Sum,Chain[j-RemovalCount].Numeral)
		    Chain = RemoveCoefficient(Chain,j-RemovalCount)
		    RemovalCount = RemovalCount + 1
		}
	    }
	    Element.Numeral = Sum
	    FinalChain = AppendCoefficient(FinalChain,Element)
	}
    }
    return FinalChain
}
//======================================================================================================================
//
// Part IV - Coefficient and Polynom Scalar Multiplication Function
//
// IV.1 - Function ScalarMulCoefficient multiplies a Coefficient by a Scalar, which is a big.Int number.
//
func ScalarMulCoefficient (C Coefficient,S *big.Int) Coefficient {
    var (
	Product Coefficient
	MulValue = new(big.Int)
    )
    if C.Numeral.Cmp(DPZero) == 0 {
	MulValue = DPZero
    } else {
	MulValue.Mul(C.Numeral,S)
    }
    Product.Numeral = MulValue
    Product.ACoeff = C.ACoeff
    Product.BCoeff = C.BCoeff
    return Product
}
//
// IV.2 - Function ScalarMulCoefficients multiplies a Coefficient-Chain by a Scalar, which is a big.Int number.
//
func ScalarMulCoefficients (C []Coefficient,S *big.Int) []Coefficient {
    var (
	Multiplication []Coefficient
    )
    for i:=0; i<len(C); i++{
	MultipliedCoeff := ScalarMulCoefficient(C[i],S)
	Multiplication = AppendCoefficient(Multiplication,MultipliedCoeff)
    }
    return Multiplication
}
//
// IV.3 - Function ScalarPolynomMul multiplies a Polynom by a Scalar, which is a big.Int number.
//
func ScalarPolynomMul (C Polynom, S *big.Int) Polynom {
    var (
	MultipliedP Polynom
	MultipliedChain [][]Coefficient
    )
    for i:=0; i<len(C.Rank); i++{
	NegatedCoeffChain := ScalarMulCoefficients(C.Coefficients[i],S)
	MultipliedChain = AppendChain(MultipliedChain,NegatedCoeffChain)
    }
    MultipliedP.Degree = C.Degree
    MultipliedP.Rank = C.Rank
    MultipliedP.Coefficients = MultipliedChain
    return MultipliedP
}
//
// IV.4 - Function NegatePolynom multiplies a Polynom by a -1 Scalar, which in effect negates it.
//
func NegatePolynom (C Polynom) Polynom {
    Multiplier := DPmOne
    Result := ScalarPolynomMul(C,Multiplier)
    return Result
}
//======================================================================================================================
//
// Part V - Coefficient Operation Functions, needed for Polynom Operation Function.
//
// V.1 - Function MulYCoefficient multiplies two YCoefficient.
//
func MulYCoefficient (C1, C2 YCoefficient)  YCoefficient {
    var(
        Product YCoefficient
        MulExp = new(big.Int)
	C1N = new(big.Int)
	C2N = new(big.Int)
	C1Y = new(big.Int)
	C2Y = new(big.Int)
	ExpY = new(big.Int)
    )
    //Constructing Numeral - Numeral remains2
    MulNum := Two

    //Constructing Exponent
    C1N = C1.NumeralExponent
    C2N = C2.NumeralExponent
    MulExp.Add(C1N,C2N)

    //Constructing Letter Y
    C1Y = C1.YCoeff.Exponent
    C2Y = C2.YCoeff.Exponent
    ExpY.Add(C1Y,C2Y)
    MulLetter := Letter{"Y",ExpY}

    //Constructing Product
    Product.Numeral = MulNum
    Product.NumeralExponent = MulExp
    Product.YCoeff = MulLetter
    return Product
}
//
// V.2 - Function MulCoefficient multiplies two Coefficient.
//
func MulCoefficient (C1, C2 Coefficient)  Coefficient {
    var (
	Product Coefficient		//Result ABStringCoeff doesnt work,
	ProductA Letter
	ProductB Letter
	NumberValue = new(big.Int)
	ExpA = new(big.Int)
	ExpB = new(big.Int)
	C1N = new(big.Int)
	C2N = new(big.Int)
	C1A = new(big.Int)
	C2A = new(big.Int)
	C1B = new(big.Int)
	C2B = new(big.Int)
    )

    //Multiplying their NumberValue
    C1N = C1.Numeral
    C2N = C2.Numeral
    NumberValue.Mul(C1N,C2N)

    //Constructing Letter A
    C1A = C1.ACoeff.Exponent
    C2A = C2.ACoeff.Exponent
    ExpA.Add(C1A,C2A)
    ProductA.Letter = "A"
    ProductA.Exponent = ExpA

    //Constructing Letter B
    C1B = C1.BCoeff.Exponent
    C2B = C2.BCoeff.Exponent
    ExpB.Add(C1B,C2B)
    ProductB.Letter = "B"
    ProductB.Exponent = ExpB

    //Constructing the Product
    Product.Numeral = NumberValue
    Product.ACoeff = ProductA
    Product.BCoeff = ProductB

    return Product
}
//
// V.3 - Function MulCoefficients multiplies two Coefficient-Chains.
//
func MulCoefficients (C1, C2 []Coefficient)  []Coefficient {
    var(
        Product = make([]Coefficient, len(C1)*len(C2))
	Count uint64
    )
    for i:=0; i<len(C1); i++ {
	for k := 0; k < len(C2); k++ {
	    if i == 0 && k == 0 {
		Count = 0
	    } else {
		Count = Count + 1
	    }
	    Multiplier := MulCoefficient(C1[i],C2[k])
	    Product[Count] = Multiplier
	}
    }
    FinalProduct := ReduceCoefficients(Product)
    return FinalProduct
}
//======================================================================================================================
//
// Part VI - Polynom Operation Functions
//
// VI.1 - Function UniqueSlice creates a Unique uint64 Slice from two uint64 Slices that might have similar elements.
// It is used to unify Rank Slices from two Polynom.
//
func UniqueSlice(Slice1,Slice2 []uint64) (Unique []uint64) {
    Unique = Slice1
    for i:=0; i<len(Slice2); i++ {
	var Check = 0
	for j := 0; j < len(Slice1); j++ {
	    if Slice2[i] == (Slice1[j]) {
	        Check = Check + 0
	    } else {
	        Check = Check + 1
	    }
	    if Check == len(Slice1) {
		Unique = append(Unique,Slice2[i])
	    }
	}
    }
    return Unique
}
//
// VI.2 - Function ModifyRank increases the uint64 slice elements by a given Value
// It is used in Polynom multiplication.
//
func ModifyRank (Rank []uint64, Value uint64) []uint64 {
    var Result = make([]uint64, len(Rank))
    for i:=0; i<len(Rank); i++ {
	Result[i] = Rank[i] + Value
    }
    return Result
}
//
// VI.3 - Function ComposePolynom creates a Polynom from its composing Elements
// It is used in Polynom multiplication.
//
func ComposePolynom(Degree uint64, Rank []uint64, Chain [][]Coefficient) Polynom {
    var ComposedPolynom Polynom
    ComposedPolynom.Degree = Degree
    ComposedPolynom.Rank = Rank
    ComposedPolynom.Coefficients = Chain
    return ComposedPolynom
}
//
// VI.4 - Function PolynomAdd adds two Polynom together
// It is also a part of Polynom multiplication.
//
func PolynomAdd (P1, P2 Polynom) Polynom {
    var (
        CChain []Coefficient
	SummedDegree uint64
	SummedRank []uint64
	SummedCoefficients [][]Coefficient
	SummedPolynom Polynom

    )

    //Part I - Setting Degree
    if P1.Degree == P2.Degree {
        SummedDegree = P1.Degree
    } else {
        if P1.Degree < P2.Degree {
            SummedDegree = P2.Degree
	} else {
	    SummedDegree = P1.Degree
	}
    }
    //Part II - Setting Rank
    SummedRank = UniqueSlice(P1.Rank,P2.Rank)

    //Part III - Setting Coefficients
    //Iterating over the 1st Polynom Coefficient-Chains
    for i:=0; i<len(P1.Rank); i++ {
	var Count = uint64(0)
	for j := 0; j<len(P2.Rank); j++ {
	    if P1.Rank[i] == P2.Rank[j] {
	        CChain = AddCoefficientChains(P1.Coefficients[i],P2.Coefficients[j])
	        SummedCoefficients = AppendChain(SummedCoefficients,CChain)
	        break
	    } else {
		Count = Count + 1
	    }
	}
	if Count == uint64(len(P2.Rank)) {
	    CChain = P1.Coefficients[i]
	    SummedCoefficients = AppendChain(SummedCoefficients,CChain)
	}
    }
    //Iterating over the 2nd Polynom Coefficient-Chains
    for k:=0; k<len(P2.Rank); k++ {
	var Count = uint64(0)
	for l := 0; l<len(P1.Rank); l++ {
	    if P2.Rank[k] == P1.Rank[l] {
		break
	    } else {
		Count = Count + 1
	    }
	}
	if Count == uint64(len(P1.Rank)) {
	    CChain = P2.Coefficients[k]
	    SummedCoefficients = AppendChain(SummedCoefficients,CChain)
	}
    }

    SummedPolynom.Degree = SummedDegree
    SummedPolynom.Rank = SummedRank
    SummedPolynom.Coefficients = SummedCoefficients
    return SummedPolynom
}
//
// VI.5 - Function PolynomSub subtracts a Polynom from another Polynom.
// It is used in computing the DivisionPolynom
//
func PolynomSub (P1, P2 Polynom) Polynom {
    P3 := NegatePolynom(P2)
    Result := PolynomAdd(P1,P3)
    return Result
}
//
// VI.6 - Function PolynomMul multiplies two Polynom together
// It is used in computing the DivisionPolynom
//
func PolynomMul (P1, P2 Polynom) Polynom {
    var (
        PolynomChain = make([]Polynom, len(P1.Rank))
        Ranks [][]uint64
        Product Polynom
    )

    for i:=0; i<len(P1.Rank); i++ {
        var (
	    //ID = new(big.Int)
	    CChain []Coefficient
            IntermediaryDegree uint64
	    IntermediaryRank []uint64
            IntermediaryCoefficients [][]Coefficient
	    Intermediary Polynom
        )
	//Part I - Setting Degree of the Intermediary Polynom
	IntermediaryDegree = P1.Rank[i] + P2.Degree

	//Part II - Setting Rank of the Intermediary Polynom
	IntermediaryRank = ModifyRank(P2.Rank,P1.Rank[i])
	CopiedIR := make([]uint64, len(IntermediaryRank))
	copy(CopiedIR,IntermediaryRank)
	Ranks = append(Ranks,CopiedIR)

	//Part III - Setting Coefficients of the Intermediary Polynom
	for j:=0; j<len(P2.Rank); j++ {
	    CChain = MulCoefficients(P1.Coefficients[i],P2.Coefficients[j])
	    IntermediaryCoefficients = AppendChain(IntermediaryCoefficients,CChain)
	}

	//Creating the Intermediary Polynom
	Intermediary = ComposePolynom(IntermediaryDegree,Ranks[i],IntermediaryCoefficients)

	//Making the PolynomChain
	PolynomChain[i] = Intermediary
    }

    //Setting Product
    if len(PolynomChain) == 1 {
        Product = PolynomChain[0]
    } else if len(PolynomChain) == 2 {
        Product = PolynomAdd(PolynomChain[0],PolynomChain[1])
    } else if len(PolynomChain) > 2 {
        Product = PolynomAdd(PolynomChain[0],PolynomChain[1])
	for i:=2; i<len(P1.Rank); i++ {
	    Product = PolynomAdd(Product,PolynomChain[i])
	}
    }
    return Product
}
//
// VI.7 - Function PolynomSqr multiplies a Polynom with itself.
// It is used in computing the DivisionPolynom
//
func PolynomSqr (P Polynom) Polynom {
    var Sqr Polynom
    Sqr = PolynomMul(P,P)
    return Sqr
}
//
// VI.8 - Function PolynomCube multiplies a Polynom with itself two times.
// It is used in computing the DivisionPolynom
//
func PolynomCube (P Polynom) Polynom {
    var Cube Polynom
    Sqr := PolynomSqr(P)
    Cube = PolynomMul(Sqr,P)
    return Cube
}
//======================================================================================================================
//
// Part VII - DivisionPolynom Operation Functions
//
// VII.1 - Function DPolynomMul multiplies two DivisionPolynom together.
// It is used in computing the DivisionPolynom
//
func DPolynomMul (P1, P2 DivisionPolynom) DivisionPolynom {
    var Result DivisionPolynom
    MulPoly := PolynomMul(P1.NonY,P2.NonY)
    MulYPoly := MulYCoefficient(P1.Y,P2.Y)

    Result.NonY = MulPoly
    Result.Y = MulYPoly
    return Result
}
//
// VII.2 - Function DPolynomSqr multiplies a DivisionPolynom with itself.
// It is used in computing the DivisionPolynom
//
func DPolynomSqr (P DivisionPolynom) DivisionPolynom {
    var Sqr DivisionPolynom
    Sqr = DPolynomMul(P,P)
    return Sqr
}
//
// VII.2 - Function DPolynomCube multiplies a DivisionPolynom with itself two times.
// It is used in computing the DivisionPolynom
//
func DPolynomCube (P DivisionPolynom) DivisionPolynom {
    var Cube DivisionPolynom
    Sqr := DPolynomSqr(P)
    Cube = DPolynomMul(Sqr,P)
    return Cube
}
//
//
//
func IsDivisibleByTwo (Number *big.Int) (bool, *big.Int) {
    var (
	Quotient = new(big.Int)
	Remainder = new(big.Int)
	Result bool
    )

    Quotient.QuoRem(Number,Two,Remainder)
    Cmp := Remainder.Cmp(DPZero)
    if Cmp == 0 {
	Result = true
    } else {
	Result = false
    }

    return Result, Quotient
}

func ConvertYtoAB (C YCoefficient) Polynom {
    var (
	Start 	= big.NewInt(0)
    	Scalar = new(big.Int)
    	Result Polynom
    )
    Scalar.Exp(C.Numeral,C.NumeralExponent,nil)
    Truth,Quotient := IsDivisibleByTwo(C.YCoeff.Exponent)
    if Truth == true {
	for i:=new(big.Int).Set(Start); i.Cmp(Quotient) == -1; i.Add(i,DPOne) {
	    if i.Cmp(DPZero) == 0 {
		//fmt.Println("i is",i)
		Result = YSquared
	    }else  {
	        //fmt.Println("i is",i)
	        Result = PolynomMul(Result,YSquared)
	    }
	}
    }
    FinalResult := ScalarPolynomMul(Result,Scalar)
    return FinalResult
}

func ReduceDivisionPolynom (P DivisionPolynom) Polynom {
    var Result Polynom
    if PrintYCoefficient(P.Y) == "1" {
        Result = P.NonY
    } else {
        Multiplier := ConvertYtoAB(P.Y)
        Result = PolynomMul(P.NonY,Multiplier)
    }
    return Result
}