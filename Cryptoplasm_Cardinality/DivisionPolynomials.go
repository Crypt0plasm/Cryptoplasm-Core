package Cryptoplasm_Cardinality

import (
    aux "Cryptoplasm-Core/Auxiliary"
    "math/big"
    "strconv"
    "strings"
)

type Letter struct {
    Letter string
    Exponent *big.Int
}

type Coefficient struct {
    Numeral *big.Int
    ACoeff Letter
    BCoeff Letter
}

type Polynom struct {
    Degree uint64
    Rank []uint64
    Coefficients [][]Coefficient
}

var (
    Zero = big.NewInt(0)
    One  = big.NewInt(1)
    Two  = big.NewInt(2)

    A0 = Letter{
	Letter: "A",
	Exponent: Zero,
    }
    B0 = Letter{
	Letter: "B",
	Exponent: Zero,
    }
    A1 = Letter{
	Letter: "A",
	Exponent: One,
    }
    B1 = Letter{
	Letter: "B",
	Exponent: One,
    }
    YSqX0 = Coefficient{One,A0,B1}
    YSqX1 = Coefficient{One,A1,B0}
    YSqX2 = Coefficient{Zero,A0,B0}
    YSqX3 = Coefficient{One,A0,B0}

    YSquared = Polynom {
        3,
	[]uint64{0, 1, 2, 3},
	[][]Coefficient{{YSqX0},{YSqX1},{YSqX2},{YSqX3}},
    }
)

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
	    } else {
		Result = Result + "+" + ReverseSlice[i]
	    }
	}
    }
    FinalResult := strings.ReplaceAll(Result,"+.","")
    return FinalResult
}

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
		    Result = Result + "+" + StringSlice[i]
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
//
// Auxiliary Function
//
//Adding Coefficients Chains simply creates a single Chain from the initial Chains and Reduces the resulting Slice
func RemoveCoefficient (Chain []Coefficient, Position int) []Coefficient {
    copy(Chain[Position:], Chain[Position+1:])
    Chain = Chain[:len(Chain)-1]
    return Chain
}

func AppendCoefficient (Chain []Coefficient, ToAppend Coefficient) (Result []Coefficient) {
    Result = append(Chain,ToAppend)
    return Result
}

func AppendChain(Chain [][]Coefficient, ToAppend []Coefficient) (Result [][]Coefficient) {
    Result = append(Chain,ToAppend)
    return Result
}

func AppendPolynom(PolynomChain []Polynom, ToAppend Polynom) (Result []Polynom) {
    Result = append(PolynomChain,ToAppend)
    return Result
}

func AddCoefficientChains (C1, C2 []Coefficient) (Result []Coefficient) {
    IntermediaryResult := append(C1,C2...)
    Result = ReduceCoefficients(IntermediaryResult)
    return Result
}
//
//Adding Coefficients simply puts single Coefficients in a Coefficient Slice and Reduces the resulting slice
func AddCoefficients (Coefficients ...Coefficient) (Result []Coefficient) {
    var IntermediaryResult = make([]Coefficient, len(Coefficients))

    for i, SingleCoefficient := range Coefficients {
	IntermediaryResult[i] = SingleCoefficient
    }
    Result = ReduceCoefficients(IntermediaryResult)
    return Result
}

//Reduce Similar values (values that have the same AB exponent) to a single value
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

//Polynom Functions
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

func ComposePolynom(Degree uint64, Rank []uint64, Chain [][]Coefficient) Polynom {
    var ComposedPolynom Polynom
    ComposedPolynom.Degree = Degree
    ComposedPolynom.Rank = Rank
    ComposedPolynom.Coefficients = Chain
    return ComposedPolynom
}

func ModifyRank (Rank []uint64, Value uint64) []uint64 {
    var Result = make([]uint64, len(Rank))
    for i:=0; i<len(Rank); i++ {
	Result[i] = Rank[i] + Value
    }
    return Result
}

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

func PolynomSqr (P Polynom) Polynom {
    var Sqr Polynom
    Sqr = PolynomMul(P,P)
    return Sqr
}

func PolynomCube (P Polynom) Polynom {
    var Cube Polynom
    Sqr := PolynomSqr(P)
    Cube = PolynomMul(Sqr,P)
    return Cube
}