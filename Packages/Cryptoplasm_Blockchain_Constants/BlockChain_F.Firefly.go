package Cryptoplasm_Blockchain_Constants

import (
	firefly "Cryptoplasm-Core/Packages/Firefly_Precision"
	"fmt"
	"os"
    	"strconv"
	"strings"
	"time"
)
var (
	c = CryptoplasmPrecisionContext
)
//
//	BlockChain_F.Firefly.go					Blockchain specific Firefly Functions
//
//================================================
// 	Function List:
//
//	01 Comparison Functions operating on decimal type
//		01  - DecimalEqual					x == y
//		02  - DecimalNotEqual				x != y
//		03  - DecimalLessThan				x < y
//		04  - DecimalLessThanOrEqual		x <= y
//		05  - DecimalGreaterThan			x > y
//		06  - DecimalGreaterThanOrEqual		x >= y
//	02 Addition Functions
//		01  - ADDel							adds 2 numbers with elastic precision
//		02  - ADDpr							adds 2 numbers within a specific precision context
//		03  - ADDcp							adds 2 numbers within CryptoplasmPrecisionContext
//		04  - SUMpr							adds multiple numbers within a specific precision context
//		05  - SUMcp							adds multiple numbers within CryptoplasmPrecisionContext
//	03 Subtraction Functions
//		01  - SUBel							subtracts 2 numbers with elastic precision
//		01  - SUBpr							subtracts 2 numbers within a specific precision context
//		02  - SUBcp							subtracts 2 numbers within CryptoplasmPrecisionContext
//		03  - DIFpr							subtracts multiple numbers within a specific precision context
//		04  - DIFcp							subtracts multiple numbers within CryptoplasmPrecisionContext
//	04 Multiplication Functions
//		01  - MULel							multiplies 2 numbers with elastic precision
//		01  - MULpr							multiplies 2 numbers within a specific precision context
//		02  - MULcp							multiplies 2 numbers within CryptoplasmPrecisionContext
//		03  - PRDpr							multiplies multiple numbers within a specific precision context
//		04  - PRDcp							multiplies multiple numbers within CryptoplasmPrecisionContext
//		05  - POWpr							computes x ** y within a specific precision context
//		06  - POWcp							computes x ** y within CryptoplasmPrecisionContext
//	05 Division Functions
//		01  - DIVpr							divides 2 numbers within a specific precision context
//		02  - DIVcp							divides 2 numbers within CryptoplasmPrecisionContext
//		03  - DivInt						returns x // y, uses elastic Precision (result is "integer")
//		04  - DivMod						returns x % y, uses elastic Precision (result is the rest)
//		05  - DigMax						max between the digit numbers of 2 Decimals
//  05a Mean Functions
//		01  - TwoMean						Returns the mean of two decimals
//	06 Truncate Functions
//		01  - TruncateCustom				Truncates using custom Precision (it must be know beforehand)
//		02  - TruncSeed						Truncates elastically to CryptoplasmSeedPrecision
//		03  - TruncToCurrency				Truncates elastically to CryptoplasmCurrencyPrecision
//	07 List Functions
//		01  - SumDL							Adds all the decimals in a slice of decimals
//		02  - LastDE						Returns the last element in a slice
//		03  - AppDec						Unites 2 slices made of decimals
//		04  - Reverse						Reverses a slice of decimals
//		05  - PrintDL						Prints the "decimals" from a slice of strings
//		06  - WriteList						Writes strings from a slice to an external file
//	08 Digit Manipulation Functions
//		01  - RemoveDecimals				removes the decimals of a number, uses floor function
//		02  - Count4Coma					Counts the number of digits before precision
//	09 Blockchain Specific Geometry Functions
//		01  - ASApr							Computes an ASA triangle with specified precision without returning the Gamma Angle
//	10 OverSend Functions
//		01  - OverSendLogBase				Returns the Logarithm Base used to computer the Overspend value for the given CP Amount
//		02  - OVSLogarithm					Computes the Logarithm in Base 777...777 for the given CP Amount
//		03  - CPAmount2StringDecomposer		Decomposes Integer Part of a cpAmount to a backwards slice of integers
//		04  - CPTxTaxV2						Computes the Transaction-Tax and its Per-Mille value, for the given CP Amount
//		05  - OverSendV2					Computes the Oversend value for the given CP Amount
//		06a - PseudoFiftyFiftyOverSendLong	Computes the pseudoFFOverSend
//		06b - PseudoFiftyFiftyOverSendShort	Computes the pseudoFFOverSend (OVerSend must be computed outside of function)
//		06c - TrueFiftyFiftyOverSendLong	Computes the FFOverSend
//		06d - TrueFiftyFiftyOverSendShort	Computes the FFOverSend (OVerSend must be computed outside of function)
//      07  - TxTaxPrinter					Computes and Prints all related TxTax information
//		07a - TxTaxDisplayOffset			Auxiliary TxTaxPrinter function
//		07b - TxTaxDisplayOffset			Auxiliary TxTaxPrinter function
//	11 Cryptoplasm Amount String Manipulation Function
//		01  - CPConvert2AU					Converts CP Amount to AtomicUnits (YoctoPlasms)
//		02  - YoctoPlasm2String				Converts YoctoPlasms into a slice os strings
//		03  - CPAmountConv2Print			Converts a CP Amount into a string that can be better used for display purposes
//
//================================================
//	01 Comparison Functions:
// 		The functions use the CryptoplasmPrecisionContext as base Context upon wherewith
// 		scaling happens automatically to accommodate their size
//================================================
//
// Function 01.01 - DecimalEqual
//
// DecimalEqual returns true if decimal x is equal to decimal y.
// The decimals must be equal to the last decimal, in this case their decimal length is similar
// They can be equal with dissimilar decimal length when one of them has extra zeros in the back
// Only works with valid Decimal type numbers.
func DecimalEqual(x, y *firefly.Decimal) bool {
    var Result bool
    digmax := DigMax(x,y)
    ComparisonContextPrecision := digmax + 1

    Difference := SUBpr(ComparisonContextPrecision, x, y)
    IsThreshold := Difference.IsZero()

    if IsThreshold == true {
        Result = true
    } else {
        Result = false
    }

    return Result
}
//================================================
//
// Function 01.02 - DecimalNotEqual
//
// DecimalNotEqual returns true if decimal x is not equal to decimal y.
// Only works with valid Decimal type numbers.
func DecimalNotEqual(x, y *firefly.Decimal) bool {
    var Result bool
    digmax := DigMax(x,y)
    ComparisonContextPrecision := digmax + 1

    Difference := SUBpr(ComparisonContextPrecision, x, y)
    IsThreshold := Difference.IsZero()

    if IsThreshold == true {
	Result = false
    } else {
	Result = true
    }

    return Result
}
//================================================
//
// Function 01.03 - DecimalLessThan
//
// DecimalLessThan returns true if decimal x is less than decimal y.
// Only works with valid Decimal type numbers.
// x equals y would return false as in this case x isnt less than y
func DecimalLessThan(x, y *firefly.Decimal) bool {
    var Result bool
    digmax := DigMax(x,y)
    ComparisonContextPrecision := digmax + 1

    Difference := SUBpr(ComparisonContextPrecision, x, y)
    //IsThreshold := Difference.IsZero()

    if Difference.Negative == true {
	Result = true
    } else {
	Result = false
    }

    return Result
}
//================================================
//
// Function 01.04 - DecimalLessThanOrEqual
//
// DecimalLessThanOrEqual returns true if decimal either
// x is less than decimal y, or if they are equal.
// Only works with valid Decimal type numbers.
func DecimalLessThanOrEqual(x, y *firefly.Decimal) bool {
    var Result bool
    digmax := DigMax(x,y)
    ComparisonContextPrecision := digmax + 1

    Difference := SUBpr(ComparisonContextPrecision, x, y)
    IsThreshold := Difference.IsZero()

    if Difference.Negative == true || IsThreshold == true{
	Result = true
    } else {
	Result = false
    }

    return Result
}
//================================================
//
// Function 01.05 - DecimalGreaterThan
//
// DecimalGreaterThan returns true if decimal x is greater than decimal y.
// Only works with valid Decimal type numbers.
// x equals y would return false as in this case x isn't less than y
func DecimalGreaterThan(x, y *firefly.Decimal) bool {
    var Result bool
    digmax := DigMax(x,y)
    ComparisonContextPrecision := digmax + 1

    Difference := SUBpr(ComparisonContextPrecision, y, x)
    //IsThreshold := Difference.IsZero()

    if Difference.Negative == true {
	Result = true
    } else {
	Result = false
    }

    return Result
}
//================================================
//
// Function 01.06 - DecimalGreaterThanOrEqual
//
// DecimalGreaterThanOrEqual returns true if decimal either
// x is greater than decimal y, or if they are equal.
// Only works with valid Decimal type numbers.
func DecimalGreaterThanOrEqual(x, y *firefly.Decimal) bool {
    var Result bool
    digmax := DigMax(x,y)
    ComparisonContextPrecision := uint32(digmax) + 1

    Difference := SUBpr(ComparisonContextPrecision, y, x)
    IsThreshold := Difference.IsZero()

    if Difference.Negative == true || IsThreshold == true{
	Result = true
    } else {
	Result = false
    }

    return Result
}
//================================================
//	02,03,04,05 Mathematical operator Functions:
// 		Mathematical operating functions used for computing
//		Addition, Subtraction, Div, Multiplication, etc
//		Basic Operations done under CryptoplasmPrecisionContext without
//		Condition and error reporting as supported by Firefly
//================================================
//
// Function 02.01 - ADDel
//
// ADDel adds two decimals within and elastically modified Precision CryptoplasmPrecisionContext Context
func ADDel(member1, member2 *firefly.Decimal) *firefly.Decimal {
	var result = new(firefly.Decimal)
	MaxDigits := DigMax(member1,member2)
	AdditionPrecision := MaxDigits + 1
	cc := c.WithPrecision(AdditionPrecision)
	_, _ = cc.Add(result, member1, member2)
	return result
}
//================================================
//
// Function 02.02 - ADDpr
//
// ADDpr adds two decimals within a custom Precision modified CryptoplasmPrecisionContext Context
func ADDpr(TotalDecimalPrecision uint32, member1, member2 *firefly.Decimal) *firefly.Decimal {
    var result = new(firefly.Decimal)
    cc := c.WithPrecision(TotalDecimalPrecision)
    _, _ = cc.Add(result, member1, member2)
    return result
}
//================================================
//
// Function 02.03 - ADDcp
//
// ADDcp adds two decimals within CryptoplasmPrecisionContext Context
func ADDcp(member1, member2 *firefly.Decimal) *firefly.Decimal {
    var result = new(firefly.Decimal)

    _, _ = c.Add(result, member1, member2)
    return result
}

//================================================
//
// Function 02.04 - SUMpr
//
// SUMpr adds multiple decimals within a custom Precision modified CryptoplasmPrecisionContext Context
func SUMpr(TotalDecimalPrecision uint32, first *firefly.Decimal, rest ...*firefly.Decimal) *firefly.Decimal {
    var (
	sum     = new(firefly.Decimal)
	restsum = firefly.NFI(0)
    )
    cc := c.WithPrecision(TotalDecimalPrecision)
    for _, item := range rest {
	_, _ = cc.Add(restsum, restsum, item)
    }
    _, _ = cc.Add(sum, first, restsum)
    return sum
}
//================================================
//
// Function 02.05 - SUMcp
//
// SUMcp adds multiple decimals within CryptoplasmPrecisionContext Context
func SUMcp(first *firefly.Decimal, rest ...*firefly.Decimal) *firefly.Decimal {
    var (
	sum     = new(firefly.Decimal)
	restsum = firefly.NFI(0)
    )

    for _, item := range rest {
	_, _ = c.Add(restsum, restsum, item)
    }
    _, _ = c.Add(sum, first, restsum)
    return sum
}
//================================================
//
// Function 03.01 - SUBel
//
// SUBel adds two decimals within and elastically modified Precision CryptoplasmPrecisionContext Context
func SUBel(member1, member2 *firefly.Decimal) *firefly.Decimal {
	var result = new(firefly.Decimal)
	MaxDigits := DigMax(member1,member2)
	AdditionPrecision := MaxDigits + 1
	cc := c.WithPrecision(AdditionPrecision)
	_, _ = cc.Sub(result, member1, member2)
	return result
}

//================================================
//
// Function 03.02 - SUBpr
//
// SUBpr subtract two decimals within a custom Precision modified CryptoplasmPrecisionContext Context
func SUBpr(TotalDecimalPrecision uint32, member1, member2 *firefly.Decimal) *firefly.Decimal {
    var result = new(firefly.Decimal)
    cc := c.WithPrecision(TotalDecimalPrecision)
    _, _ = cc.Sub(result, member1, member2)
    return result
}
//================================================
//
// Function 03.03 - SUBcp
//
// SUBcp subtract two decimals within CryptoplasmPrecisionContext Context
func SUBcp(member1, member2 *firefly.Decimal) *firefly.Decimal {
    var result = new(firefly.Decimal)

    _, _ = c.Sub(result, member1, member2)
    return result
}
//================================================
//
// Function 03.04 - DIFpr
//
// DIFpr subtract multiple decimals within a custom Precision modified CryptoplasmPrecisionContext Context
func DIFpr(TotalDecimalPrecision uint32, first *firefly.Decimal, rest ...*firefly.Decimal) *firefly.Decimal {
    var (
	difference = new(firefly.Decimal)
	restsum    = firefly.NFI(0)
    )
    cc := c.WithPrecision(TotalDecimalPrecision)
    for _, item := range rest {
	_, _ = cc.Add(restsum, restsum, item)
    }
    _, _ = cc.Sub(difference, first, restsum)
    return difference
}
//================================================
//
// Function 03.05 - DIFcp
//
// DIFcp subtract multiple decimals within CryptoplasmPrecisionContext Context
func DIFcp(first *firefly.Decimal, rest ...*firefly.Decimal) *firefly.Decimal {
    var (
	difference = new(firefly.Decimal)
	restsum    = firefly.NFI(0)
    )
    for _, item := range rest {
	_, _ = c.Add(restsum, restsum, item)
    }
    _, _ = c.Sub(difference, first, restsum)
    return difference
}
//================================================
//
// Function 03.01 - MULel
//
// MULel multiplies two decimals within and elastically modified Precision CryptoplasmPrecisionContext Context
func MULel(member1, member2 *firefly.Decimal) *firefly.Decimal {
	var result = new(firefly.Decimal)

	NumberDigitsMember1 := member1.NumDigits()
	NumberDigitsMember2 := member2.NumDigits()
	MultiplicationPrecision := NumberDigitsMember1 + NumberDigitsMember2

	//MaxDigits := DigMax(member1,member2)
	//MultiplicationPrecision := MaxDigits + 1
	cc := c.WithPrecision(uint32(MultiplicationPrecision))
	_, _ = cc.Mul(result, member1, member2)
	return result
}

//================================================
//
// Function 04.01 - MULpr
//
// MULpr multiplies two decimals within a custom Precision modified CryptoplasmPrecisionContext Context
func MULpr(TotalDecimalPrecision uint32, member1, member2 *firefly.Decimal) *firefly.Decimal {
    var result = new(firefly.Decimal)
    cc := c.WithPrecision(TotalDecimalPrecision)
    _, _ = cc.Mul(result, member1, member2)
    return result
}
//================================================
//
// Function 04.02 - MULcp
//
// MULcp multiplies two decimals within CryptoplasmPrecisionContext Context
func MULcp(member1, member2 *firefly.Decimal) *firefly.Decimal {
    var result = new(firefly.Decimal)

    _, _ = c.Mul(result, member1, member2)
    return result
}
//================================================
//
// Function 04.03 - PRDpr
//
// PRDpr multiplies multiple decimals within a custom Precision modified CryptoplasmPrecisionContext Context
func PRDpr(TotalDecimalPrecision uint32, first *firefly.Decimal, rest ...*firefly.Decimal) *firefly.Decimal {
    var (
	product     = new(firefly.Decimal)
	restproduct = firefly.NFI(1)
    )
    cc := c.WithPrecision(TotalDecimalPrecision)
    for _, item := range rest {
	_, _ = cc.Mul(restproduct, restproduct, item)
    }
    _, _ = cc.Mul(product, first, restproduct)

    return product
}
//================================================
//
// Function 04.04 - PRDcp
//
// PRDcp multiplies multiple decimals within CryptoplasmPrecisionContext Context
func PRDcp(first *firefly.Decimal, rest ...*firefly.Decimal) *firefly.Decimal {
    var (
	product     = new(firefly.Decimal)
	restproduct = firefly.NFI(1)
    )

    for _, item := range rest {
	_, _ = c.Mul(restproduct, restproduct, item)
    }
    _, _ = c.Mul(product, first, restproduct)

    return product
}
//================================================
//
// Function 04.05 - POWpr
//
// POWpr multiplies two decimals within a custom Precision modified CryptoplasmPrecisionContext Context
func POWpr(TotalDecimalPrecision uint32, x, y *firefly.Decimal) *firefly.Decimal {
    var result = new(firefly.Decimal)
    cc := c.WithPrecision(TotalDecimalPrecision)
    _, _ = cc.Pow(result, x, y)
    return result
}
//================================================
//
// Function 04.06 - POWcp
//
// POWcp multiplies two decimals within CryptoplasmPrecisionContext Context
func POWcp(x, y *firefly.Decimal) *firefly.Decimal {
    var result = new(firefly.Decimal)

    _, _ = c.Pow(result, x, y)
    return result
}
//================================================
//
// Function 05.01 - DIVpr
//
// DIVpr multiplies two decimals within a custom Precision modified CryptoplasmPrecisionContext Context
func DIVpr(TotalDecimalPrecision uint32, member1, member2 *firefly.Decimal) *firefly.Decimal {
    var result = new(firefly.Decimal)
    cc := c.WithPrecision(TotalDecimalPrecision)
    _, _ = cc.Quo(result, member1, member2)
    return result
}
//================================================
//
// Function 05.02 - DIVcp
//
// DIVcp multiplies two decimals within CryptoplasmPrecisionContext Context
func DIVcp(member1, member2 *firefly.Decimal) *firefly.Decimal {
    var result = new(firefly.Decimal)

    _, _ = c.Quo(result, member1, member2)
    return result
}
//================================================
//
// Function 05.03 - DivInt
//
// DivInt returns the integer part of x divided by y
// It is equal to x // y
// Returned Value is also of decimal Type
func DivInt (x, y *firefly.Decimal) *firefly.Decimal {
    var result = new(firefly.Decimal)
    digmax := DigMax(x,y)
    DCP := digmax + 1		//DivisionContextPrecision
    cc := c.WithPrecision(DCP)
    _,_ = cc.QuoInteger(result,x,y)
    return result
}
//================================================
//
// Function 05.04 - DivMod
//
// DivMod returns the remainder from the division of x to y
// It is equal to x % y
// Returned Value is also of decimal Type
func DivMod (x, y *firefly.Decimal) *firefly.Decimal {
    var result = new(firefly.Decimal)
    digmax := DigMax(x,y)
    DCP := digmax + 1		//DivisionContextPrecision
    divresult := TruncateCustom(DCP,DivInt(x,y),0)
    result = SUBpr(DCP,x,MULpr(DCP,y,divresult))
    return result
}
//================================================
//
// Function 05.05 - DigMax
//
// DigMax returns maximum between the number of digits of two Decimals
// Used for autoscaling precision for operations over Decimals
func DigMax (x, y *firefly.Decimal) uint32 {
    var digmax int64

    xdig := x.NumDigits()
    ydig := y.NumDigits()
    digdiff := xdig - ydig
    if digdiff <= 0 {
	digmax = ydig
    } else if digdiff >= 0{
	digmax = xdig
    }
    return uint32(digmax)
}
//================================================
//	05a Mean Functions:
// 		Different types of means used for computing purposes
//		In specific ways
//================================================
func TwoMean(number1, number2 *firefly.Decimal) *firefly.Decimal {
	var result = new(firefly.Decimal)
	MaxDigits := DigMax(number1,number1)
	DivisionPrecision := MaxDigits + 1
	result = DIVpr(DivisionPrecision,ADDel(number1,number2),firefly.NFI(2))
	return result
}
//================================================
//	06 Truncate Functions:
// 		Functions used to Truncate Decimals to specific precision
//		In specific ways
//================================================
//
// Function 06.01 - TruncateCustom
//
// TruncateCustom truncates the decimal to the specified precision number
func TruncateCustom(TotalDecimalPrecision uint32, number *firefly.Decimal, DecimalPrecision uint32) *firefly.Decimal {
	var result = new(firefly.Decimal)
	ConvertedPrecision := int32(DecimalPrecision)
	ConvertedPrecision = 0 - ConvertedPrecision
	cc := c.WithPrecision(TotalDecimalPrecision)
	_, _ = cc.Quantize(result, number, ConvertedPrecision)
	return result
}
//================================================
//
// Function 06.02 - TruncSeed
//
// TruncSeed truncates the decimal to CryptoplasmSeedPrecision
func TruncSeed(SeedNumber *firefly.Decimal) *firefly.Decimal {
	var result = new(firefly.Decimal)

	NumberDigits := Count4Coma(SeedNumber)
	TruncatingContextPrecision := uint32(NumberDigits) + CryptoplasmSeedPrecision
	cc := c.WithPrecision(TruncatingContextPrecision)

	CSP := 0 - int32(CryptoplasmSeedPrecision)
	_, _ = cc.Quantize(result, SeedNumber, CSP)
	return result
}
//================================================
//
// Function 06.03 - TruncToCurrency
//
// TruncToCurrency truncates the decimal to CryptoplasmCurrencyPrecision
// It is Context Precision Independent
func TruncToCurrency(Amount2BecomeCP *firefly.Decimal) *firefly.Decimal {
	var result = new(firefly.Decimal)

	NumberDigits := Count4Coma(Amount2BecomeCP)
	TruncatingContextPrecision := uint32(NumberDigits) + CryptoplasmCurrencyPrecision
	cc := c.WithPrecision(TruncatingContextPrecision)

	CCP := 0 - int32(CryptoplasmCurrencyPrecision)
	_, _ = cc.Quantize(result, Amount2BecomeCP, CCP)
	return result
}
//================================================
//	07 List Function:
// 		Functions that operate on slices of decimals
//		for different Purposes,
//		basically "emulating" Python List capability.
//================================================
//
// Function 07.01 - SumDL
//
// SumDL short for SumDecimalList, return the sum of
// the decimals within the list/slice
func SumDL(a []*firefly.Decimal) *firefly.Decimal {
	var sum = new(firefly.Decimal)

	for i := 0; i < len(a); i++ {
		sum = ADDcp(sum, a[i])
	}
	return sum
}
//================================================
//
// Function 07.02 - LastDE
//
// LastDE short for LastDecimalElement, returns the last element
// in the slice. Equivalent to pythons -1 index
func LastDE(a []*firefly.Decimal) *firefly.Decimal {
	Length := len(a)
	LastElementIndex := Length - 1
	LastElement := a[LastElementIndex]
	return LastElement
}
//================================================
//
// Function 07.03 - AppDec
//
// AppDec creates a new bigger slice from the 2 slices used as input
// slices must be made of decimals
func AppDec(w1, w2 []*firefly.Decimal) []*firefly.Decimal {
	w3 := append(w1, w2...)
	return w3
}
//================================================
//
// Function 07.04 - Reverse
//
// Returns the Reverse of the Slice/Lists
func Reverse(a []*firefly.Decimal) []*firefly.Decimal {
	var Reversed = make([]*firefly.Decimal, 0)
	Length := len(a)
	LastElementIndex := Length - 1
	for i := LastElementIndex; i >= 0; i-- {
		Reversed = append(Reversed, a[i])
	}
	return Reversed
}
//================================================
//
// Function 07.05 - PrintDL
//
// PrintDL short for PrintDecimalList, prints the decimals
// within the given list/slice
func PrintDL(a []string) {
    for i := 0; i < len(a); i++ {
	fmt.Println("Element is,", a[i])
    }
}
//================================================
//
// Function 07.06 - WriteList
//
// WriteList writes the strings from the slice to an external file
// as Name can be used "File.txt" as the output file.
func WriteList(Name string, List []string) {
	f, err := os.Create(Name)

	if err != nil {
		fmt.Println(err)
		_ = f.Close()
		return
	}

	for _, v := range List {
		_, _ = fmt.Fprintln(f, v)
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("file written successfully")
	return
}
//================================================
//	08 Digit Manipulations Function:
// 		Operations done on the number of Digits of a decimal
//		Used for setting elastic precision on different other functions
//================================================
//
// Function 08.01 - RemoveDecimals
//
// RemoveDecimals removes the decimals and leaves the resulted number
// without them
func RemoveDecimals(Number *firefly.Decimal) *firefly.Decimal {
	var Whole = new(firefly.Decimal)
	NumberDigits := Number.NumDigits()
	cc := c.WithPrecision(uint32(NumberDigits))
	_, _ = cc.Floor(Whole, Number)
	return Whole
}
//================================================
//
// Function 08.02 - Count4Coma
//
// Count4Coma returns the number of digits before precision
func Count4Coma(Number *firefly.Decimal) int64 {
	Whole := RemoveDecimals(Number)
	Digits := Whole.NumDigits()
	return Digits
}
//================================================
//	09 Blockchain Specific Geometry Functions
//		For now only a custom ASA triangle function is needed
//================================================
//
// Function 09.01 - ASApr
//
// ASApr computes and ASA triangle with a custom precision
// as used in the Seed computer. Therefore it doesnt return the AngleG
func ASApr(TotalDecimalPrecision uint32, angleAlfa, sideC, angleBeta *firefly.Decimal) (*firefly.Decimal, *firefly.Decimal, *firefly.Decimal) {
    var (
	angleA = angleAlfa
	angleB = angleBeta
	sdC    = sideC
	sdA    = new(firefly.Decimal)
	sdB    = new(firefly.Decimal)
	area   = new(firefly.Decimal)
    )
    cc := c.WithPrecision(TotalDecimalPrecision)
    _, sdA, sdB, area = cc.ASA(angleA, sdC, angleB)
    //ASAcp doesnt return angleG as this isn't used in the computation
    return sdA, sdB, area
}
//================================================
//	10 Transaction Tax and OverSend related Functions:
// 		Functions that are used for computing
//		OverSend and its related Transaction Tax are here
//================================================
//
// Function 10.01 - OverSendLogBase
//
// OverSendLogBase returns the logarithm base required for computing overspend
func OverSendLogBase(cpAmount *firefly.Decimal) *firefly.Decimal {
	var Base = new(firefly.Decimal)

	DigitsNumber := Count4Coma(cpAmount)
	DND := firefly.NFI(DigitsNumber) //making decimal the Digit number
	DigitsOperatingPrecision := Count4Coma(DND)

	Exponent := SUBpr(uint32(DigitsOperatingPrecision), DND, firefly.NFI(2))
	e := firefly.INT64(Exponent)
	for i := e; i >= 0; i-- {
		idec := firefly.NFI(i)
		Value := MULpr(uint32(DigitsNumber), POWpr(uint32(DigitsNumber), firefly.NFI(10), idec), firefly.NFI(7))
		Base = ADDpr(uint32(DigitsNumber), Base, Value)
	}
	return Base
}
//================================================
//
// Function 10.02 - OVSLogarithm
//
// OVSLogarithm returns the logarithm from "number" in base "base".
func OVSLogarithm(base, number *firefly.Decimal) *firefly.Decimal {
	var (
		LogBase   = new(firefly.Decimal)
		LogNumber = new(firefly.Decimal)
	)
	//For LogBase and LogNumber Context precision
	//2+24 Context precision is enough, for base and number below e^100
	//if such were the case, a 3+24 (CryptoplasmCurrencyPrecision)
	//precision would be required. However e^100 has an Integer of 44 digits, namely
	//26.881.171.418.161.354.484.126.255.515.800.135.873.611.118
	//So one would have to need to compute the OverSend for a CP amount
	//bigger than this number to have the need to use a 3+24 context precision,
	//for computing the first logarithm below. Therefore 2+24 context precision
	//for both logarithms should be enough
	//27 Context Precision would be enough to compute the needed logarithm
	//for many more coins that could ever be minted until the End of the Universe.
	//if the Cryptoplasm emission would be repeated for every subsequent 524.596.891 Blocks (1 ERA)
	//1(ERA ~ 107 to 110 Trillion CP).

	//As the resulted LNs have the same number of digits for their integer part
	//a context Precision of 1+24 would always be enough, as the division would always
	//look like 1,.....

	//+3 is used, so such a high amount of coins to compute the OverSend for will also
	//work, and, as has been tested, indeed the code allows it to work.

	NumberDigits := number.NumDigits()
	IP := 2*CryptoplasmCurrencyPrecision + uint32(NumberDigits)
	cc := c.WithPrecision(IP)
	_, _ = cc.Ln(LogBase, base)
	_, _ = cc.Ln(LogNumber, number)
	CustomLog := DIVpr(IP, LogNumber, LogBase)
	return CustomLog
}
//================================================
//
// Function 10.03 - CPAmount2StringDecomposer
//
// CPAmount2StringDecomposer creates a string of uint8 from the integer digits
// of a cpAmount, to be used for calculating the TxTax
func CPAmount2StringDecomposer(cpAmount *firefly.Decimal) []uint8 {
	var NumberSlice []uint8
	tcpAmount := TruncToCurrency(cpAmount)
	tcpAmountInteger := RemoveDecimals(tcpAmount)
	NumberDigits := tcpAmountInteger.NumDigits()
	NumberToManipulate := tcpAmountInteger
	for i := 0; i < int(NumberDigits); i++ {
		Div10 := DIVpr(uint32(NumberDigits),NumberToManipulate,firefly.NFI(10))
		RemoveLast := RemoveDecimals(Div10)
		Multiplication := MULel(RemoveLast,firefly.NFI(10))
		LastDigit := SUBel(NumberToManipulate,Multiplication)
		LastDigitINT64 := firefly.INT64(LastDigit)
		LastDigitSmall := uint8(LastDigitINT64)
		NumberToManipulate = RemoveLast
		NumberSlice = append(NumberSlice,LastDigitSmall)

	}
	return NumberSlice
}
//================================================
//
// Function 10.04 - CPTxTaxV2
//
// CPTxTaxV2 computes the modified TxTax.
func CPTxTaxV2(cpAmount *firefly.Decimal) (*firefly.Decimal, *firefly.Decimal, *firefly.Decimal) {
	var (
		TxTax = new(firefly.Decimal)
		MultipliedTax = new(firefly.Decimal)
	)
	tcpAmount := TruncToCurrency(cpAmount)
	BaseDigitNumberMain := tcpAmount.NumDigits()
	MainDivPrecision := 2*CryptoplasmCurrencyPrecision + uint32(BaseDigitNumberMain)
	//fmt.Println("Computing TxTax")
	if DecimalGreaterThanOrEqual(tcpAmount,firefly.NFI(10)) == true {
		AmountNumberSlice := CPAmount2StringDecomposer(cpAmount)
		for i := 1; i < len(AmountNumberSlice); i++ {
			BaseStringPoint := "START Computing TxTax"
			StringPoint := strings.Repeat(".",i)
			StringToPrint := BaseStringPoint + StringPoint
			fmt.Print("\r",StringToPrint)

			Number:=POWpr(uint32(i)+1,firefly.NFI(10),firefly.NFI(int64(i)))
			LogBase := OverSendLogBase(Number)
			ProMille := OVSLogarithm(LogBase,Number)
			BaseDigitNumber := Number.NumDigits()
			DivPrecision := 2*CryptoplasmCurrencyPrecision + uint32(BaseDigitNumber)
			Tax := DIVpr(DivPrecision,MULel(Number,ProMille),firefly.NFI(1000))
			if AmountNumberSlice[i] == 0 {
				MultipliedTax = firefly.NFI(0)
			} else {
				MultipliedTax = MULel(Tax,firefly.NFI(int64(AmountNumberSlice[i])))
			}
			TxTax = ADDel(TxTax,MultipliedTax)
		}
	}
	fmt.Println("DONE")
	TxTax = TruncToCurrency(TxTax)
	Recipient := SUBel(tcpAmount,TxTax)
	CumulativeFeeProMille := TruncToCurrency(DIVpr(MainDivPrecision,MULel(TxTax,firefly.NFI(1000)),tcpAmount))
	//fmt.Print("===")
	return CumulativeFeeProMille, TxTax, Recipient
}
//================================================
//
// Function 10.05 - OverSendV2
//
// OverSendV2 computes the modified TxTax.
func OverSendV2(cpAmount *firefly.Decimal) *firefly.Decimal {
	fmt.Println("START Computing OverSend...")
	tcpAmount := TruncToCurrency(cpAmount)
	_,TxTaxAmount,_ := CPTxTaxV2(tcpAmount)
	OverSend := ADDel(tcpAmount,TxTaxAmount)
	Iteration := 0
	for DecimalEqual(TxTaxAmount,firefly.NFI(0)) == false {
		Iteration = Iteration + 1
		_,TxTaxAmount,_ = CPTxTaxV2(TxTaxAmount)
		OverSend = ADDel(OverSend,TxTaxAmount)
	}
	//fmt.Println("Last TxTax computation:")
	_,LastTxTax,_ := CPTxTaxV2(OverSend)
	LastOverSend := ADDel(tcpAmount,LastTxTax)
	fmt.Println("DONE Computing OverSend, after ",Iteration," iteration(s) ...")
	return LastOverSend
}
//================================================
//
// Function 10.06a - PseudoFiftyFiftyOverSendLong
//
// PseudoFiftyFiftyOverSendLong computes a PseudoFiftyFiftyOverSend
// used to later compute the TrueFiftyFiftyOverSend.
// It is called "long", because it includes the OverSend Computation.
func PseudoFiftyFiftyOverSendLong (cpAmount *firefly.Decimal) *firefly.Decimal {
	var PseudoOverSend = new(firefly.Decimal)
	PerfectOverSend := OverSendV2(cpAmount)
	PseudoOverSend = PseudoFiftyFiftyOverSendShort(cpAmount,PerfectOverSend)
	return PseudoOverSend
}
//================================================
//
// Function 10.06b - PseudoFiftyFiftyOverSendShort
//
// PseudoFiftyFiftyOverSendShort computes a PseudoFiftyFiftyOverSend
// used to later compute the TrueFiftyFiftyOverSend.
// It is called "short", because it relies on the OverSend value.
func PseudoFiftyFiftyOverSendShort(cpAmount, PerfectOverSend *firefly.Decimal) *firefly.Decimal {
	var PseudoFFOverSend = new(firefly.Decimal)
	MaxDigits := DigMax(cpAmount,PerfectOverSend)
	DivisionPrecision := MaxDigits + 1

	tcpAmount := TruncToCurrency(cpAmount)
	_, TxTaxMin, _ := CPTxTaxV2(tcpAmount)
	_, TxTaxMax, _ := CPTxTaxV2(PerfectOverSend)

	MeanTx := TwoMean(TxTaxMin,TxTaxMax)
	HalfMeanTx := DIVpr(DivisionPrecision,MeanTx,firefly.NFI(2))
	PseudoFFOverSend = ADDel(cpAmount,HalfMeanTx)
	return PseudoFFOverSend
}
//================================================
//
// Function 10.06c - TrueFiftyFiftyOverSendLong
//
// TrueFiftyFiftyOverSendLong computes the TrueFiftyFiftyOverSend
// The TrueFiftyFiftyOverSend is the OverSend value for which both the
// Sender and the Receiver pay the same amount of Transaction Tax.
// It is called "long", because it includes the OverSend Computation.
func TrueFiftyFiftyOverSendLong (cpAmount *firefly.Decimal) *firefly.Decimal {
	var TrueFiftyFiftyOverSend = new(firefly.Decimal)
	PerfectOverSend := OverSendV2(cpAmount)
	TrueFiftyFiftyOverSend = TrueFiftyFiftyOverSendShort(cpAmount,PerfectOverSend)
	//x1,x2 := TrueFiftyFiftyOverSendShort(cpAmount,PerfectOverSend)
	return TrueFiftyFiftyOverSend
	//return x1,x2
}
//================================================
//
// Function 10.06d - TrueFiftyFiftyOverSendShort
//
// TrueFiftyFiftyOverSendShort computes the TrueFiftyFiftyOverSend
// The TrueFiftyFiftyOverSend is the OverSend value for which both the
// Sender and the Receiver pay the same amount of Transaction Tax.
// It is called "short", because it relies on the OverSend value.
func TrueFiftyFiftyOverSendShort(cpAmount, PerfectOverSend *firefly.Decimal) *firefly.Decimal {
	//start := time.Now()
	var (
		TrueFiftyFifty 			= new(firefly.Decimal)
		MinimumDifference 		= new(firefly.Decimal)
		YoctoPlasm				= firefly.NFS("0.000000000000000000000001")
		Iteration       		int                     	//OverSpend Loop Iteration number
	)

	//Setting IP(internal precision) and truncating cpAmount to 24 decimals
	NumberDigits := Count4Coma(PerfectOverSend)
	IP := CryptoplasmCurrencyPrecision + uint32(NumberDigits) + 1
	tcpAmount := TruncToCurrency(cpAmount)

	OverSendFF := PseudoFiftyFiftyOverSendShort(cpAmount,PerfectOverSend)
	_, _, RecipientFF := CPTxTaxV2(OverSendFF)

	TxTaxS := TruncToCurrency(SUBpr(IP,OverSendFF,tcpAmount))
	TxTaxR := TruncToCurrency(SUBpr(IP,tcpAmount,RecipientFF))

	FluctuatingOverSend := SUMpr(IP,TxTaxR,DIVpr(IP,SUBpr(IP,TxTaxS,TxTaxR),firefly.NFI(2)),tcpAmount)
	for DecimalNotEqual(TxTaxS,TxTaxR) == true {
		Iteration = Iteration + 1
		if DecimalGreaterThan(TxTaxS,TxTaxR) == true {
			MinimumDifference = SUBpr(IP,TxTaxS,TxTaxR)
		} else {
			MinimumDifference = SUBpr(IP,TxTaxR,TxTaxS)
		}
		if DecimalEqual(TxTaxS,TxTaxR) == true || DecimalEqual(MinimumDifference,YoctoPlasm) == true{
			break
		}
		if DecimalGreaterThan(TxTaxS,TxTaxR) == true {
			//Loop Down
			fmt.Println("Computing FiftyFiftyOverSend, refining difference...")
			_, _, RFiftyFifty := CPTxTaxV2(FluctuatingOverSend)
			TxTaxS = TruncToCurrency(SUBpr(IP,FluctuatingOverSend,tcpAmount))
			TxTaxR = TruncToCurrency(SUBpr(IP,tcpAmount,RFiftyFifty))
			FluctuatingOverSend = SUMpr(IP,TxTaxR,DIVpr(IP,SUBpr(IP,TxTaxS,TxTaxR),firefly.NFI(2)),tcpAmount)
			//fmt.Println("FlucOvs is",FluctuatingOverSend,"TxTaxS is",TxTaxS,"TxTaxSR is",TxTaxR)

		} else if DecimalLessThan(TxTaxS,TxTaxR) == true {
			//Loop Up
			fmt.Println("Computing FiftyFiftyOverSend, refining difference...")
			_, _, RFiftyFifty := CPTxTaxV2(FluctuatingOverSend)

			TxTaxS = TruncToCurrency(SUBpr(IP,FluctuatingOverSend,tcpAmount))
			TxTaxR = TruncToCurrency(SUBpr(IP,tcpAmount,RFiftyFifty))
			FluctuatingOverSend = SUMpr(IP,TxTaxR,DIVpr(IP,SUBpr(IP,TxTaxS,TxTaxR),firefly.NFI(2)),tcpAmount)
			//fmt.Println("FlucOvs is",FluctuatingOverSend,"TxTaxS is",TxTaxS,"TxTaxSR is",TxTaxR)
		}
	}
	TrueFiftyFifty = FluctuatingOverSend
	_, _, RecipientFiftyFifty := CPTxTaxV2(TrueFiftyFifty)

	TxTaxS = SUBpr(IP,TrueFiftyFifty,tcpAmount)
	TxTaxR = SUBpr(IP,tcpAmount,RecipientFiftyFifty)
	//elapsed := time.Since(start)
	//fmt.Println("Computing FiftyFiftyOverSend took", elapsed, "with", Iteration, "Iterations")
	//fmt.Println("")
	return TrueFiftyFifty
}
//================================================
//
// Function 10.07 - TxTaxPrinter
//
// TxTaxPrinter computes and prints all OverSend related information
// for the given cpAmount
func TxTaxPrinter(cpAmount *firefly.Decimal) {
	start := time.Now()
	var (
		BarNumberOffset			int
		Len00					string
	)
	tcpAmount := TruncToCurrency(cpAmount)
	//Computing Oversend and FiftyFiftyOverSend
	Zero := "0,[000|000|000|000][000|000|000|000]"

	fmt.Println("\tPART 1 START - tcpAmount known, computing TxTax based on tcpAmount:")
	FeeProMilleMin, TxTaxMin, 	RecipientMin 	:= CPTxTaxV2(tcpAmount)
	fmt.Println("\tPART 1 FINISH - Done Computing TxTax based on tcpAmount")
	fmt.Println("")

	fmt.Println("\tPART 2 START - Computing OverSend, based on tcpAmount:")
	OverSend := OverSendV2(tcpAmount)
	FeeProMilleMax, TxTaxMax, 	RecipientMax 	:= CPTxTaxV2(OverSend)
	fmt.Println("\tPART 2 FINISH - Done Computing TxTax based on OverSend")
	fmt.Println("")

	fmt.Println("\tPART 3 - Computing FiftyFiftyOverSend, based on tcpAmount amd OverSend")
	FiftyFiftyOverSend := TrueFiftyFiftyOverSendShort(cpAmount,OverSend)
	FeeProMilleFF, 	TxTaxFF, 	RecipientFF 	:= CPTxTaxV2(FiftyFiftyOverSend)
	fmt.Println("\tPART 3 FINISH - Done Computing TxTax based on FiftyFiftyOverSend")
	fmt.Println("")

	SenderTxTax := SUBel(FiftyFiftyOverSend,tcpAmount)
	RecipientTxTax := SUBel(tcpAmount,RecipientFF)

	OverSendAmountLength 	:= len(CPAmountConv2Print(OverSend))
	TargetAmountLength 		:= len(CPAmountConv2Print(tcpAmount))

	if OverSendAmountLength > TargetAmountLength {
		BarNumberOffset = len(CPAmountConv2Print(OverSend))
		Len00 = TxTaxDisplayOffset(OverSend,tcpAmount)
	} else if OverSendAmountLength == TargetAmountLength {
		BarNumberOffset = len(CPAmountConv2Print(tcpAmount))
	}

	BarLength 		:= 33 + 1 + BarNumberOffset + 3
	BarLengthString := strings.Repeat("=",BarLength)
	Len01 			:= TxTaxDisplayOffset(OverSend,TxTaxMin)
	Len02 			:= TxTaxDisplayOffset(OverSend,FeeProMilleMin)
	Len04 			:= TxTaxDisplayOffsetS(OverSend,Zero)
	Len01b			:= TxTaxDisplayOffset(OverSend,TxTaxMax)
	Len02b			:= TxTaxDisplayOffset(OverSend,FeeProMilleMax)
	Len01c			:= TxTaxDisplayOffset(OverSend,TxTaxFF)
	Len02c			:= TxTaxDisplayOffset(OverSend,FeeProMilleFF)
	Len04c			:= TxTaxDisplayOffset(OverSend,SenderTxTax)
	Len04d			:= TxTaxDisplayOffset(OverSend,RecipientTxTax)

	OptionNumber1 := `Option Number 1, Sender sends the Target-Amount:
When the Tx-Tax is deducted from this amount, the Recipient gets the Target-Amount,
minus the Tx-Tax. Therefore it is said that the Recipient lost/"payed" the Tx-Tax,
while the Sender has spent no extra money on the Tx-Tax.
`
	OptionNumber2 := `Option Number 2, Sender sends the OverSend-Amount:
When the Tx-Tax is deducted from this amount, the Recipient receives the Target-Amount
Therefore it is said that the Sender "payed" the Tx-Tax.
While the Recipient has lost no money because he received the intended Target-Amount.
`
	OptionNumber3 := `Option Number 3, Sender sends the FiftyFiftyOverSend-Amount: 
Sender and Receiver "split" the Tx-Tax. The Tx-Tax deducted from this amount makes 
it so, that when comparing to the Target-Amount, both Sender & Receiver have "lost" 
the exact amount of money. The Tx-Tax Split lost/"paid" by both the Sender and the 
Receiver is either perfectly equal, or differs by 1 YoctoPlasm. The difference 
happens when the Tx-Tax's last decimal is an uneven number.
`

	fmt.Println("")
	fmt.Print("\t",OptionNumber1)
	fmt.Println("")
	fmt.Print("\t",OptionNumber2)
	fmt.Println("")
	fmt.Print("\t",OptionNumber3)
	fmt.Println("")
	if OverSendAmountLength > TargetAmountLength {
		fmt.Println("       Target Cryptoplasm Amount:",Len00,CPAmountConv2Print(tcpAmount),"CP")
	} else if OverSendAmountLength == TargetAmountLength {
		fmt.Println("       Target Cryptoplasm Amount:",CPAmountConv2Print(tcpAmount),"CP")
	}
	fmt.Println(BarLengthString)
	//==================================================================================================================
	fmt.Println("=============Option1=============")				//TARGET AMOUNT
	// First Line
	if OverSendAmountLength > TargetAmountLength {
		fmt.Println("     Sender sends[Target Amount]:",Len00,CPAmountConv2Print(tcpAmount),"CP")
	} else if OverSendAmountLength == TargetAmountLength {
		fmt.Println("     Sender sends[Target Amount]:",CPAmountConv2Print(tcpAmount),"CP")
	}
	//Second and third Line
	fmt.Println("       the Tx-Tax[Target Amount]:",Len01,CPAmountConv2Print(TxTaxMin),"CP")
	fmt.Println("                  and represents:",Len02,FeeProMilleMin,"promille")
	//Fourth Line
	if len(CPAmountConv2Print(tcpAmount)) == len(CPAmountConv2Print(RecipientMin)) {
		if OverSendAmountLength > TargetAmountLength {
			fmt.Println("              Recipient receives:",Len00,CPAmountConv2Print(RecipientMin),"CP")
		} else if OverSendAmountLength == TargetAmountLength {
			fmt.Println("              Recipient receives:",CPAmountConv2Print(RecipientMin),"CP")
		}
	} else {
		Len03 := TxTaxDisplayOffset(tcpAmount,RecipientMin)
		fmt.Println("              Recipient receives:",Len03,CPAmountConv2Print(RecipientMin),"CP")
	}
	fmt.Println("=================================")
	//Fifth and Sixth Line
	fmt.Println("    Sender has payed a Tx-Tax of:",Len04,Zero,"CP")
	fmt.Println(" Recipient has payed a Tx-Tax of:",Len01,CPAmountConv2Print(TxTaxMin),"CP")
	fmt.Println(BarLengthString)
	//==================================================================================================================
	fmt.Println("=============Option2=============")				//OVERSEND
	//First, Second and Third Line
	fmt.Println("          Sender sends[OverSend]:",CPAmountConv2Print(OverSend),"CP")
	fmt.Println("            the Tx-Tax[OverSend]:",Len01b,CPAmountConv2Print(TxTaxMax),"CP")
	fmt.Println("                  and represents:",Len02b,FeeProMilleMax,"promille")
	//Fourth Line
	if len(CPAmountConv2Print(OverSend)) == len(CPAmountConv2Print(RecipientMax)) {
		if OverSendAmountLength > TargetAmountLength {
			fmt.Println("              Recipient receives:",Len00,CPAmountConv2Print(RecipientMax),"CP")
		} else if OverSendAmountLength == TargetAmountLength {
			fmt.Println("              Recipient receives:",CPAmountConv2Print(RecipientMax),"CP")
		}
	} else {
		Len03b := TxTaxDisplayOffset(OverSend,RecipientMax)
		fmt.Println("              Recipient receives:",Len03b,CPAmountConv2Print(RecipientMax),"CP")
	}
	fmt.Println("=================================")
	//Fifth and Sixth Line
	fmt.Println("    Sender has payed a Tx-Tax of:",Len01b,CPAmountConv2Print(TxTaxMax),"CP")
	fmt.Println(" Recipient has payed a Tx-Tax of:",Len04,Zero,"CP")
	fmt.Println(BarLengthString)
	//==================================================================================================================
	fmt.Println("=============Option3=============")				//FIFTYFIFTYOVERSEND
	//First, Second and Third Line
	fmt.Println("Sender sends[FiftyFiftyOverSend]:",CPAmountConv2Print(FiftyFiftyOverSend),"CP")
	fmt.Println("  the Tx-Tax[FiftyFiftyOverSend]:",Len01c,CPAmountConv2Print(TxTaxFF),"CP")
	fmt.Println("                  and represents:",Len02c,FeeProMilleFF,"promille")
	//Fourth Line
	if len(CPAmountConv2Print(FiftyFiftyOverSend)) == len(CPAmountConv2Print(RecipientFF)) {
		if OverSendAmountLength > TargetAmountLength {
			fmt.Println("              Recipient receives:",Len00,CPAmountConv2Print(RecipientFF),"CP")
		} else if OverSendAmountLength == TargetAmountLength {
			fmt.Println("              Recipient receives:",CPAmountConv2Print(RecipientFF),"CP")
		}
	} else {
		Len03c := TxTaxDisplayOffset(FiftyFiftyOverSend,RecipientFF)
		fmt.Println("              Recipient receives:",Len03c,CPAmountConv2Print(RecipientFF),"CP")
	}
	fmt.Println("=================================")
	//Fifth and Sixth Line
	fmt.Println("    Sender has payed a Tx-Tax of:",Len04c,CPAmountConv2Print(SenderTxTax),"CP")
	fmt.Println(" Recipient has payed a Tx-Tax of:",Len04d,CPAmountConv2Print(RecipientTxTax),"CP")
	fmt.Println(BarLengthString)
	elapsed := time.Since(start)
	fmt.Println("Computing the TxTaxPrinter took", elapsed)
	fmt.Println("")
	}

//================================================
//
// Function 10.07a - TxTaxDisplayOffset
//
// TxTaxDisplayOffset in an auxiliary TxTaxPrinter function
func TxTaxDisplayOffset (cpAmount1, cpAmount2 *firefly.Decimal) string {
	var Result string
	Length := len(CPAmountConv2Print(cpAmount1)) - len(CPAmountConv2Print(cpAmount2))
	Result = strings.Repeat(" ",Length-1)
	return Result
}
//
// Function 10.07b - TxTaxDisplayOffsetS
//
// TxTaxDisplayOffsetS in an auxiliary TxTaxPrinter function
func TxTaxDisplayOffsetS (cpAmount1 *firefly.Decimal, AmountAsString string) string {
	var Result string
	Length := len(CPAmountConv2Print(cpAmount1)) - len(AmountAsString)
	Result = strings.Repeat(" ",Length-1)
	return Result
}
//================================================
//	11 CP Amount String Manipulation Function:
// 		Functions that manipulate CP Amounts
//		formatting them for displaying purposes.
//================================================
//
// Function 11.01 - CPConvert2AU
//
// CPConvert2AU converts a CryptoPlasm amount into Atomic Units
func CPConvert2AU(cpAmount *firefly.Decimal) *firefly.Decimal {
    tcpAmount := TruncToCurrency(cpAmount)
    NumberDigits := Count4Coma(cpAmount)
    IP := uint32(NumberDigits) + CryptoplasmCurrencyPrecision
    AU := MULpr(IP,tcpAmount,AUs)

    return AU
}
//================================================
//
// Function 11.02 - YoctoPlasm2String
//
// YoctoPlasm2String converts a CryptoPlasm AUs (YoctoPlasms)
// into a slice of strings
func YoctoPlasm2String(Number *firefly.Decimal) []string {
    var SliceStr []string
    Ten := firefly.NFI(10)
    AuDigits := Number.NumDigits()
    Exp := AuDigits - 1
    IP 	:= uint32(AuDigits)
    //Exp := firefly.NFI(NumberDigitsAU - 1)
    ToSequence := Number
    for i := Exp; i >= 0; i-- {
        idec := firefly.NFI(i)
        Power := POWpr(IP,Ten,idec)
        Division := DIVpr(IP,ToSequence,Power)

        DigitIs := TruncateCustom(IP,Division,0)
        DI := firefly.INT64(DigitIs)
        DigitIsString := strconv.Itoa(int(DI))
        SliceStr = append(SliceStr,DigitIsString)

        Rest := SUBpr(IP,Division,DigitIs)
        SmallAU := MULpr(IP,Rest,Power)
	ToSequence = SmallAU
    }
    return SliceStr
}
//================================================
//
// Function 11.03 - CPAmountConv2Print
//
// CPAmountConv2Print converts CryptoPlasm amount into a string
// to be used for printing purposes.
// For now only a "." as decimal character is implemented.
// Different Schemas can be added (for instance using coma<,> as decimal separator
// instead of point<.>; Or using points for thousand separator,
// or even separating at 2 position for Lakhs and Crores.
func CPAmountConv2Print (cpAmount *firefly.Decimal) string {
    var (
    	StringResult 		string
		ComaPosition 		int32
		PointPosition 		int32
		DigitTier		int32
		NewSlicePositions 	int64
    )
    AU := CPConvert2AU(cpAmount)
    SliceStr := YoctoPlasm2String(AU)
    NumberDigits := Count4Coma(AU)

    //Computing the Coma position
    if NumberDigits <= 25 {
	ComaPosition = 1
    } else {
	ComaPosition = int32(NumberDigits) - 24
    }
    //Adding the Coma aka Decimal Separator
    SliceStr = append(SliceStr,"")
    copy(SliceStr[ComaPosition+1:],SliceStr[ComaPosition:])
    SliceStr[ComaPosition] = ","

    if NumberDigits >= 28 {
	//Computing the 1000 Separator positions
	Difference := NumberDigits - 25
	if Difference  % 3 == 0 {
	    DigitTier = 1
	} else if Difference  % 3 == 1 {
	    DigitTier = 2
	} else if Difference  % 3 == 2{
	    DigitTier = 3
	}
	TSNumber := (NumberDigits - 25 ) / 3

	//Adding the 1000 Separator as points
	for i := 1; i<=int(TSNumber); i++ {
	    PointPosition = (int32(i)-1) * 4 + DigitTier
	    SliceStr = append(SliceStr,"")
	    copy(SliceStr[PointPosition+1:],SliceStr[PointPosition:])
	    SliceStr[PointPosition] = "."
	}
	NewSlicePositions = NumberDigits + TSNumber + 1
    } else if NumberDigits < 28 && NumberDigits > 24 {
	NewSlicePositions = NumberDigits + 1
    } else {
	NewSlicePositions = 26
    }

    for i:=0; i<=9; i++{
        var x, y, z, s int
        var Insert string
        InsertFront := "["
        InsertMiddle := "|"
        InsertEnd := "]"

        s = 4
        if i == 0 {
            x = -1
            y = 0
            z = 0
	    Insert = InsertEnd
	} else if  i != 0 && i <= 3 {
	    x = s * i
	    y = x
	    z = x + 1
	    Insert = InsertMiddle
	} else if i == 4 {
	    x = s * i
	    y = x
	    z = x + 1
	    Insert = InsertFront
	} else if i == 5 {
	    x = s * (i - 1) + 1
	    y = x
	    z = x + 1
	    Insert = InsertEnd
	} else if i > 5 && i <= 8 {
	    x = s * (i - 1) + 1
	    y = x
	    z = x + 1
	    Insert = InsertMiddle
	} else {
	    x = s * (i - 1) + 1
	    y = x
	    z = x + 1
	    Insert = InsertFront
	}
	//fmt.Println("SliceStr este inainte",SliceStr,i)
	SliceStr = append(SliceStr,"")
	//fmt.Println("SliceStr este dupa",SliceStr,i)
	copy(SliceStr[int(NewSlicePositions)-x:],SliceStr[int(NewSlicePositions)-z:])
	SliceStr[int(NewSlicePositions)-y] = Insert
	NewSlicePositions = NewSlicePositions + 1
    }

    //Removing "0," from the SliceString, displaying only Decimals
    if len(SliceStr) == 36 && SliceStr[0] == "0" {
	SliceStr = SliceStr[2:]
    }

    for i := 0; i < len(SliceStr); i++ {
        StringResult = StringResult + SliceStr[i]
    }
    return StringResult
}