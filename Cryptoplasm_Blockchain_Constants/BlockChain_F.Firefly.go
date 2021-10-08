package Cryptoplasm_Blockchain_Constants

import (
    "fmt"
    "os"
    "strconv"
    "strings"
    "time"

    p "github.com/Crypt0plasm/Firefly-APD"
)
var (
	c = CryptoplasmPrecisionContext
)
//
//	BlockChain_F.Firefly.go				Blockchain specific Firefly Functions
//================================================================================================
//************************************************************************************************
//================================================================================================
// 	Function List:
//
//	01 Comparison Functions operating on decimal type
//		00  - SummedMaxLengthPlusOne		SummedMaxLength returns the sum of the maximums length of digits (b4 and after coma)
//		00a - MaxInt32				Returns the maximum between two int32 numbers
//		00b - MaxInt64				Returns the maximum between two int64 numbers
//		01  - DecimalEqual			x == y
//		02  - DecimalNotEqual			x != y
//		03  - DecimalLessThan			x < y
//		04  - DecimalLessThanOrEqual		x <= y
//		05  - DecimalGreaterThan		x > y
//		06  - DecimalGreaterThanOrEqual		x >= y
//	02 Addition Functions
//		01  - ADDx				Adds 2 numbers with custom total precision
//		02  - ADDs				Adds 2 numbers within CryptoplasmPrecisionContext (70 total precision)
//		03  - ADD				Adds 2 numbers with custom decimal precision and elastic integer precision
//		03a - ADDxs				Adds 2 numbers with 70 decimal precision and elastic integer precision
//		03b - ADDxc				Adds 2 numbers with 100 decimal precision and elastic integer precision
//		04  - SUMx				Adds multiple numbers with custom total precision
//		05  - SUMs				Adds multiple numbers within CryptoplasmPrecisionContext (70 total precision)
//		06  - SUM				Adds multiple numbers with custom decimal precision and elastic integer precision
//		06a - SUMxs				Adds multiple numbers with 70 decimal precision and elastic integer precision
//		06b - SUMxc				Adds multiple numbers with 100 decimal precision and elastic integer precision
//	03 Subtraction Functions
//		01  - SUBx				Subtracts 2 numbers with custom total precision
//		02  - SUBs				Subtracts 2 numbers within CryptoplasmPrecisionContext (70 total precision)
//		03  - SUB				Subtracts 2 numbers with custom decimal precision and elastic integer precision
//		03a - SUBxs				Subtracts 2 numbers with 70 decimal precision and elastic integer precision
//		03b - SUBxc				Subtracts 2 numbers with 100 decimal precision and elastic integer precision
//		04  - DIFx				Subtracts multiple numbers with custom total precision
//		05  - DIFs				Subtracts multiple numbers within CryptoplasmPrecisionContext (70 total precision)
//		06  - DIF				Subtracts multiple numbers with custom decimal precision and elastic integer precision
//		06a - DIFxs				Subtracts multiple numbers with 70 decimal precision and elastic integer precision
//		06b - DIFxc				Subtracts multiple numbers with 100 decimal precision and elastic integer precision
//	04 Multiplication Functions
//		01  - MULx				Multiplies 2 numbers with custom total precision
//		02  - MULs				Multiplies 2 numbers within CryptoplasmPrecisionContext (70 total precision)
//		03  - MULxc				Multiplies 2 numbers with elastic integer precision and 100 max decimal precision
//		04  - PRDx				Multiplies multiple numbers within a specific precision context
//		05  - PRDs				Multiplies multiple numbers within CryptoplasmPrecisionContext
//		06  - PRDxc				Multiplies multiple numbers with elastic integer precision and 100 max decimal precision
//		07  - POWx				Computes x ** y within a specific precision context
//		08  - POWs				Computes x ** y within CryptoplasmPrecisionContext
//		09  - POWxc				Computes x ** y with elastic integer precision and 100 max decimal precision
//	05 Division Functions
//		01  - DIVx				Divides 2 numbers within a specific precision context
//		02  - DIVs				Divides 2 numbers within CryptoplasmPrecisionContext
//		03  - DIVxc				Divides 2 numbers with elastic integer precision and 100/101 max decimal precision
//		04  - DivInt				Returns x // y, uses elastic Precision (result is "integer")
//		05  - DivMod				Returns x % y, uses elastic Precision (result is the rest)
//  05a Mean Functions
//		01  - TwoMean				Returns the mean of two decimals
//	06 Truncate Functions
//		01  - TruncateCustom			Truncates using custom Precision (it must be know beforehand)
//		02  - TruncSeed				Truncates elastically to CryptoplasmSeedPrecision
//		03  - TruncToCurrency			Truncates elastically to CryptoplasmCurrencyPrecision
//		04  - TruncPercent			Truncates elastically to CryptoplasmPercentPrecision
//	07 List Functions
//		01  - SumDL				Adds all the decimals in a slice of decimals
//		02  - LastDE				Returns the last element in a slice
//		03  - AppDec				Unites 2 slices made of decimals
//		04  - Reverse				Reverses a slice of decimals
//		05  - PrintDL				Prints the "decimals" from a slice of strings
//		06  - WriteList				Writes strings from a slice to an external file
//	08 Digit Manipulation Functions
//		01  - RemoveDecimals			Removes the decimals of a number, uses floor function
//		02  - Count4Coma			Counts the number of digits before precision
//	09 Blockchain Specific Geometry Functions
//		01  - ASApr				Computes an ASA triangle with specified precision without returning the Gamma Angle
//	10 OverSend Functions
//		01  - OverSendLogBase			Returns the Logarithm Base used to computer the Overspend value for the given CP Amount
//		02  - OVSLogarithm			Computes the Logarithm in Base 777...777 for the given CP Amount
//		03  - CPAmount2StringDecomposer		Decomposes Integer Part of a cpAmount to a backwards slice of integers
//		04  - CPTxTaxV2				Computes the Transaction-Tax and its Per-Mille value, for the given CP Amount
//		05  - OverSendV2			Computes the Oversend value for the given CP Amount
//		06a - PseudoFiftyFiftyOverSendLong	Computes the pseudoFFOverSend
//		06b - PseudoFiftyFiftyOverSendShort	Computes the pseudoFFOverSend (OVerSend must be computed outside of function)
//		06c - TrueFiftyFiftyOverSendLong	Computes the FFOverSend
//		06d - TrueFiftyFiftyOverSendShort	Computes the FFOverSend (OVerSend must be computed outside of function)
//      	07  - TxTaxPrinter			Computes and Prints all related TxTax information
//		07a - TxTaxDisplayOffset		Auxiliary TxTaxPrinter function
//		07b - TxTaxDisplayOffset		Auxiliary TxTaxPrinter function
//	11 Cryptoplasm Amount String Manipulation Function
//		01  - CPConvert2AU			Converts CP Amount to AtomicUnits (YoctoPlasms)
//		02  - YoctoPlasm2String			Converts YoctoPlasms into a slice os strings
//		03  - CPAmountConv2Print		Converts a CP Amount into a string that can be better used for display purposes
//
//================================================================================================
//************************************************************************************************
//================================================================================================
//
// Function 01.00 - SummedMaxLength
//
// SummedMaxLength returns the sum of the maximums length of digits.
// before and after the coma for two decimals.
// Used in comparison operation, and to elastically increase precision based on integer part size of the decimals
// Even thought it on itself is enough to secure total operation precision, it is used as extra buffer when computing
// total operation precision for ADD SUB MUL and DIV functions. (because an additional DecimalPrecision is added on top of it)
func SummedMaxLengthPlusOne (x, y *p.Decimal) uint32 {
	var SML uint32
	IntegerDigitsMember1 := Count4Coma(x)		//int64
	IntegerDigitsMember2 := Count4Coma(y)		//int64
	DecimalDigitsMember1 := 0 - x.Exponent		//int32
	DecimalDigitsMember2 := 0 - y.Exponent		//int32

	MaxIntegerDigitsInt64 := MaxInt64(IntegerDigitsMember1,IntegerDigitsMember2)			//int64
	MaxDecimalDigitsInt32 := MaxInt32(DecimalDigitsMember1,DecimalDigitsMember2)			//int32

	// Observation1
	// Converting Int64 to Int32, going down from 9.223.372.036.854.775.807 to 2.147.483.647
	// As 2.147.483.647 are already a huge number, no check is implemented here.
	MaxIntegerDigitsInt32 := int32(MaxIntegerDigitsInt64)

	// Observation 2
	// SML is of uint32 type, this means this function works reliably for numbers with up to
	// 2.147.483.647 digits each.
	// Two times this is 4.294.967.294. Adding another 1 equals 4.294.967.295.
	// This is the maximum number representable by uint32.
	//
	// So the maximum length x and y can have is 2.147.483.647 digits before and after the coma.
	SML = uint32(MaxDecimalDigitsInt32) + uint32(MaxIntegerDigitsInt32) + 1
	return SML
}
//================================================================================================
//	01 Comparison Functions between integers:
//================================================
//
// Function 01.00a - MaxInt32
//
// MaxInt32 returns the maximum between two int32 numbers
func MaxInt32(x, y int32) int32 {
	var max int32
	digdiff := x - y
	if digdiff <= 0 {
		max = y
	} else if digdiff >= 0{
		max = x
	}
	return max
}
//================================================
//
// Function 01.00b - MaxInt64
//
// MaxInt64 returns the maximum between two int64 numbers
func MaxInt64(x, y int64) int64 {
	var max int64
	digdiff := x - y
	if digdiff <= 0 {
		max = y
	} else if digdiff >= 0{
		max = x
	}
	return max
}
//================================================================================================
//	01 Comparison Functions between decimals:
//	The functions use the SummedMaxLengthPlusOne function to set the ComparisonContextPrecision
//================================================================================================
//
// Function 01.01 - DecimalEqual
//
// DecimalEqual returns true if decimal x is equal to decimal y.
func DecimalEqual(x, y *p.Decimal) bool {
    var Result bool
    ComparisonContextPrecision := SummedMaxLengthPlusOne(x,y)

    Difference := SUBx(ComparisonContextPrecision, x, y)
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
func DecimalNotEqual(x, y *p.Decimal) bool {
    var Result bool
	ComparisonContextPrecision := SummedMaxLengthPlusOne(x,y)

    Difference := SUBx(ComparisonContextPrecision, x, y)
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
func DecimalLessThan(x, y *p.Decimal) bool {
    var Result bool
	ComparisonContextPrecision := SummedMaxLengthPlusOne(x,y)

    Difference := SUBx(ComparisonContextPrecision, x, y)
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
func DecimalLessThanOrEqual(x, y *p.Decimal) bool {
    var Result bool
	ComparisonContextPrecision := SummedMaxLengthPlusOne(x,y)

    Difference := SUBx(ComparisonContextPrecision, x, y)
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
func DecimalGreaterThan(x, y *p.Decimal) bool {
    var Result bool
	ComparisonContextPrecision := SummedMaxLengthPlusOne(x,y)

    Difference := SUBx(ComparisonContextPrecision, y, x)
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
func DecimalGreaterThanOrEqual(x, y *p.Decimal) bool {
    var Result bool
	ComparisonContextPrecision := SummedMaxLengthPlusOne(x,y)

    Difference := SUBx(ComparisonContextPrecision, y, x)
    IsThreshold := Difference.IsZero()

    if Difference.Negative == true || IsThreshold == true{
	Result = true
    } else {
	Result = false
    }

    return Result
}
//================================================================================================
//	02,03,04,05 Mathematical operator Functions:
// 		Mathematical operating functions used for computing
//		Addition, Subtraction, Div, Multiplication, etc
//		Basic Operations done under CryptoplasmPrecisionContext without
//		Condition and error reporting as supported by p
//================================================================================================
//************************************************************************************************
//================================================================================================
//
// Function 02.01 - ADDx
//
// ADDx adds two decimals within a custom Precision modified CryptoplasmPrecisionContext Context
func ADDx(TotalDecimalPrecision uint32, member1, member2 *p.Decimal) *p.Decimal {
	var result = new(p.Decimal)
	cc := c.WithPrecision(TotalDecimalPrecision)
	_, _ = cc.Add(result, member1, member2)
	return result
}
//================================================
//
// Function 02.02 - ADDs
//
// ADDs adds two decimals within CryptoplasmPrecisionContext Context
func ADDs(member1, member2 *p.Decimal) *p.Decimal {
	var result = new(p.Decimal)
	_, _ = c.Add(result, member1, member2)
	return result
}
//================================================
//
// Function 02.03 - ADD
//
// ADD adds two decimals within custom Precision modified CryptoplasmPrecisionContext Context
// The Precision has "DecimalPrecision" decimal Precision plus elastic integer Precision.
// The Precision scales with the number size, but is limited to "DecimalPrecision" decimals.
func ADD(DecimalPrecision uint32, member1, member2 *p.Decimal) *p.Decimal {
	var result = new(p.Decimal)
	DNBDP := SummedMaxLengthPlusOne(member1,member2) 	//DigitNumberBasedDecimalPrecision
	//Observation
	// As "SummedMaxLengthPlusOne" returns a uint32 variable (maximum of 4.294.967.295)
	// TotalDecimalPrecision will overflow uint32 if adding the "DecimalPrecision" on top of DNBDP because
	// it (TotalDecimalPrecision) would get bigger than 4.294.967.295.
	// However, this isn't expected to happen, which is why no check or error detection is implemented.
	TotalDecimalPrecision := DNBDP + DecimalPrecision

	cc := c.WithPrecision(TotalDecimalPrecision)
	_, _ = cc.Add(result, member1, member2)
	return result
}
//================================================
//
// Function 02.03a - ADDxs
//
// ADDxs adds two decimals within custom Precision modified CryptoplasmPrecisionContext Context
// The Precision has 70 decimal Precision plus elastic integer Precision.
// The Precision scales with the number size, but is limited to 70 decimals.
func ADDxs(member1, member2 *p.Decimal) *p.Decimal {
	return ADD(CryptoplasmStdMathPrecision, member1, member2)
}
//================================================
//
// Function 02.03b - ADDxc
//
// ADDxc adds two decimals within custom Precision modified CryptoplasmPrecisionContext Context
// The Precision has 100 decimal Precision plus elastic integer Precision.
// The Precision scales with the number size, but is limited to 100 decimals.
func ADDxc(member1, member2 *p.Decimal) *p.Decimal {
	return ADD(CryptoplasmMaxMathPrecision, member1, member2)
}
//================================================
//
// Function 02.04 - SUMx
//
// SUMx adds multiple decimals within a custom Precision modified CryptoplasmPrecisionContext Context
func SUMx(TotalDecimalPrecision uint32, first *p.Decimal, rest ...*p.Decimal) *p.Decimal {
    var (
		sum     = new(p.Decimal)
		restsum = p.NFI(0)
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
// Function 02.05 - SUMs
//
// SUMs adds multiple decimals within CryptoplasmPrecisionContext Context
func SUMs(first *p.Decimal, rest ...*p.Decimal) *p.Decimal {
    var (
	sum     = new(p.Decimal)
	restsum = p.NFI(0)
    )

    for _, item := range rest {
	_, _ = c.Add(restsum, restsum, item)
    }
    _, _ = c.Add(sum, first, restsum)
    return sum
}
//================================================
//
// Function 02.06 - SUM
//
// SUM sums multiple decimals within custom Precision modified CryptoplasmPrecisionContext Context
// The Precision has "DecimalPrecision" decimal Precision plus elastic integer Precision.
// The Precision scales with the number size, but is limited to "DecimalPrecision" decimals.
func SUM(DecimalPrecision uint32, first *p.Decimal, rest ...*p.Decimal) *p.Decimal {
	var (
		sum     = new(p.Decimal)
		restsum = p.NFI(0)
	)
	for _, item := range rest {
		restsum = ADD(DecimalPrecision,restsum,item)
	}
	sum = ADD(DecimalPrecision,first,restsum)
	return sum
}
//================================================
//
// Function 02.06a - SUMxs
//
// SUMxs adds two decimals within custom Precision modified CryptoplasmPrecisionContext Context
// The Precision has 70 decimal Precision plus elastic integer Precision.
// The Precision scales with the number size, but is limited to 70 decimals.
func SUMxs(first *p.Decimal, rest ...*p.Decimal) *p.Decimal {
	return SUM(CryptoplasmStdMathPrecision, first, rest...)
}
//================================================
//
// Function 02.06b - SUMxc
//
// SUMxc adds two decimals within custom Precision modified CryptoplasmPrecisionContext Context
// The Precision has 100 decimal Precision plus elastic integer Precision.
// The Precision scales with the number size, but is limited to 100 decimals.
func SUMxc(first *p.Decimal, rest ...*p.Decimal) *p.Decimal {
	return SUM(CryptoplasmMaxMathPrecision, first, rest...)
}
//================================================================================================
//************************************************************************************************
//================================================================================================
//
// Function 03.01 - SUBx
//
// SUBx subtract two decimals within a custom Precision modified CryptoplasmPrecisionContext Context
func SUBx(TotalDecimalPrecision uint32, member1, member2 *p.Decimal) *p.Decimal {
	var result = new(p.Decimal)
	cc := c.WithPrecision(TotalDecimalPrecision)
	_, _ = cc.Sub(result, member1, member2)
	return result
}
//================================================
//
// Function 03.02 - SUBs
//
// SUBs subtract two decimals within CryptoplasmPrecisionContext Context
func SUBs(member1, member2 *p.Decimal) *p.Decimal {
	var result = new(p.Decimal)
	_, _ = c.Sub(result, member1, member2)
	return result
}
//================================================
//
// Function 03.03 - SUB
//
// SUB subtracts two decimals within custom Precision modified CryptoplasmPrecisionContext Context
// The Precision has "DecimalPrecision" decimal Precision plus elastic integer Precision.
// The Precision scales with the number size, but is limited to "DecimalPrecision" decimals.

func SUB(DecimalPrecision uint32, member1, member2 *p.Decimal) *p.Decimal {
	var result = new(p.Decimal)
	DNBDP := SummedMaxLengthPlusOne(member1,member2)	//DigitNumberBasedDecimalPrecision
	//Observation
	// As "SummedMaxLengthPlusOne" returns a uint32 variable (maximum of 4.294.967.295)
	// TotalDecimalPrecision will overflow uint32 if adding the "DecimalPrecision" on top of DNBDP because
	// it (TotalDecimalPrecision) would get bigger than 4.294.967.295.
	// However, this isn't expected to happen, which is why no check or error detection is implemented.
	TotalDecimalPrecision := DNBDP + DecimalPrecision

	cc := c.WithPrecision(TotalDecimalPrecision)
	_, _ = cc.Sub(result, member1, member2)
	return result
}
//================================================
//
// Function 03.03a - SUBxs
//
// SUBxs subtracts two decimals within custom Precision modified CryptoplasmPrecisionContext Context
// The Precision has 70 decimal Precision plus elastic integer Precision.
// The Precision scales with the number size, but is limited to 70 decimals.
func SUBxs(member1, member2 *p.Decimal) *p.Decimal {
	return SUB(CryptoplasmStdMathPrecision, member1, member2)
}
//================================================
//
// Function 03.03b - SUBxc
//
// SUBxc subtracts two decimals within custom Precision modified CryptoplasmPrecisionContext Context
// The Precision has 100 decimal Precision plus elastic integer Precision.
// The Precision scales with the number size, but is limited to 100 decimals.
func SUBxc(member1, member2 *p.Decimal) *p.Decimal {
	return SUB(CryptoplasmMaxMathPrecision, member1, member2)
}
//================================================
//
// Function 03.04 - DIFx
//
// DIFx subtracts multiple decimals within a custom Precision modified CryptoplasmPrecisionContext Context
func DIFx(TotalDecimalPrecision uint32, first *p.Decimal, rest ...*p.Decimal) *p.Decimal {
	var (
		sum     = new(p.Decimal)
		restsum = p.NFI(0)
	)
	cc := c.WithPrecision(TotalDecimalPrecision)
	for _, item := range rest {
		_, _ = cc.Add(restsum, restsum, item)
	}
	_, _ = cc.Sub(sum, first, restsum)
	return sum
}
//================================================
//
// Function 03.05 - DIFs
//
// DIFs subtracts multiple decimals within CryptoplasmPrecisionContext Context
func DIFs(first *p.Decimal, rest ...*p.Decimal) *p.Decimal {
	var (
		sum     = new(p.Decimal)
		restsum = p.NFI(0)
	)

	for _, item := range rest {
		_, _ = c.Add(restsum, restsum, item)
	}
	_, _ = c.Sub(sum, first, restsum)
	return sum
}
//================================================
//
// Function 03.06 - DIF
//
// DIF subtracts multiple decimals within custom Precision modified CryptoplasmPrecisionContext Context
// The Precision has "DecimalPrecision" decimal Precision plus elastic integer Precision.
// The Precision scales with the number size, but is limited to "DecimalPrecision" decimals.
func DIF(DecimalPrecision uint32, first *p.Decimal, rest ...*p.Decimal) *p.Decimal {
	var (
		sum     = new(p.Decimal)
		restsum = p.NFI(0)
	)
	for _, item := range rest {
		restsum = ADD(DecimalPrecision,restsum,item)
	}
	sum = SUB(DecimalPrecision,first,restsum)
	return sum
}
//================================================
//
// Function 03.06a - DIFxs
//
// DIFxs subtracts two decimals within custom Precision modified CryptoplasmPrecisionContext Context
// The Precision has 70 decimal Precision plus elastic integer Precision.
// The Precision scales with the number size, but is limited to 70 decimals.
func DIFxs(first *p.Decimal, rest ...*p.Decimal) *p.Decimal {
	return DIF(CryptoplasmStdMathPrecision, first, rest...)
}
//================================================
//
// Function 03.06b - DIFxc
//
// DIFxc subtracts two decimals within custom Precision modified CryptoplasmPrecisionContext Context
// The Precision has 100 decimal Precision plus elastic integer Precision.
// The Precision scales with the number size, but is limited to 100 decimals.
func DIFxc(first *p.Decimal, rest ...*p.Decimal) *p.Decimal {
	return DIF(CryptoplasmMaxMathPrecision, first, rest...)
}
//================================================================================================
//************************************************************************************************
//================================================================================================
//
// Function 04.01 - MULx
//
// MULx multiplies two decimals within a custom Precision modified CryptoplasmPrecisionContext Context
func MULx(TotalDecimalPrecision uint32, member1, member2 *p.Decimal) *p.Decimal {
	var result = new(p.Decimal)
	cc := c.WithPrecision(TotalDecimalPrecision)
	_, _ = cc.Mul(result, member1, member2)
	return result
}
//================================================
//
// Function 04.02 - MULs
//
// MULs multiplies two decimals within CryptoplasmPrecisionContext Context
func MULs(member1, member2 *p.Decimal) *p.Decimal {
	var result = new(p.Decimal)
	_, _ = c.Mul(result, member1, member2)
	return result
}
//================================================
//
// Function 04.03 - MULxc
//
// MULxc multiplies two decimals within an elastically modified Precision CryptoplasmPrecisionContext Context
// The elastic Precision's decimal limit is set to 100, while the integer precision scales without any "limits".
// Any limits means only a theoretical hard limit of 4.294.967.195 digits, 100 units less than uint32.
// This is however expected never to be achieved.
func MULxc(member1, member2 *p.Decimal) *p.Decimal {
	var (
		result = new(p.Decimal)
		DecimalPrecision uint32
	)

	IntegerDigitsMember1 := Count4Coma(member1)				//int64
	IntegerDigitsMember2 := Count4Coma(member2)				//int64
	DecimalDigitsMember1 := 0 - member1.Exponent				//int32
	DecimalDigitsMember2 := 0 - member2.Exponent				//int32
	IntegerSumInt64 := IntegerDigitsMember1 + IntegerDigitsMember2		//int64 9.223.372.036.854.775.807
	DecimalSumInt32 := DecimalDigitsMember1 + DecimalDigitsMember2		//int32 2.147.483.647
	IntegerSumUint32 := uint32(IntegerSumInt64)				// from 9.223.372.036.854.775.807 to max 4.294.967.295
	DecimalSumUint32 := uint32(DecimalSumInt32)				// from 2.147.483.647 to max 4.294.967.295

	//Max IntegerSum can be 4.294.967.295
	//Max DecimalSum is limited to 100.
	//As these are added to give the total precision, Max IntegerSum can be as high as 4.294.967.195

	if DecimalSumUint32 < CryptoplasmMaxMathPrecision {
		DecimalPrecision = DecimalSumUint32
	} else {
		DecimalPrecision = CryptoplasmMaxMathPrecision
	}
	MultiplicationPrecision := IntegerSumUint32 + DecimalPrecision

	cc := c.WithPrecision(MultiplicationPrecision)
	_, _ = cc.Mul(result, member1, member2)
	return result
}
//================================================
//
// Function 04.04 - PRDx
//
// PRDx multiplies multiple decimals within a custom Precision modified CryptoplasmPrecisionContext Context
func PRDx(TotalDecimalPrecision uint32, first *p.Decimal, rest ...*p.Decimal) *p.Decimal {
    var (
	product     = new(p.Decimal)
	restproduct = p.NFI(1)
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
// Function 04.05 - PRDs
//
// PRDs multiplies multiple decimals within CryptoplasmPrecisionContext Context
func PRDs(first *p.Decimal, rest ...*p.Decimal) *p.Decimal {
    var (
	product     = new(p.Decimal)
	restproduct = p.NFI(1)
    )

    for _, item := range rest {
	_, _ = c.Mul(restproduct, restproduct, item)
    }
    _, _ = c.Mul(product, first, restproduct)

    return product
}
//================================================
//
// Function 04.06 - PRDxc
//
// PRDxc multiplies two decimals within an elastically modified Precision CryptoplasmPrecisionContext Context
// The elastic Precision's decimal limit is set to 100, while the integer precision scales without any "limits".
// Any limits means only a theoretical hard limit of 4.294.967.195 digits, 100 units less than uint32.
// This is however expected never to happen.
func PRDxc(first *p.Decimal, rest ...*p.Decimal) *p.Decimal {
	var (
		product     = new(p.Decimal)
		restproduct = p.NFI(1)
	)

	for _, item := range rest {
		restproduct = MULxc(restproduct,item)
	}
	product = MULxc(first,restproduct)
	_, _ = c.Mul(product, first, restproduct)

	return product
}

//================================================
//
// Function 04.07 - POWx
//
// POWx computes x ** y within a custom Precision modified CryptoplasmPrecisionContext Context
func POWx(TotalDecimalPrecision uint32, member1, member2 *p.Decimal) *p.Decimal {
    var result = new(p.Decimal)
    cc := c.WithPrecision(TotalDecimalPrecision)
    _, _ = cc.Pow(result, member1, member2)
    return result
}
//================================================
//
// Function 04.08 - POWs
//
// POWs computes x ** y within CryptoplasmPrecisionContext Context
func POWs(member1, member2 *p.Decimal) *p.Decimal {
    var result = new(p.Decimal)

    _, _ = c.Pow(result, member1, member2)
    return result
}
//================================================
//
// Function 04.06 - POWxc
//
// POWxc computes x ** y within an elastically modified Precision CryptoplasmPrecisionContext Context
// The elastic Precision's decimal limit is set to 100, while the integer precision scales without any "limits".
// Any limits means only a theoretical hard limit of 4.294.967.195 digits, 100 units less than uint32.
// This is however expected never to happen.
func POWxc(member1, member2 *p.Decimal) *p.Decimal {
	var result     = new(p.Decimal)

	IntegerDigitsMember1 	:= Count4Coma(member1)							//int64
	IntegerMember2 			:= p.INT64(RemoveDecimals(member2))		//int64
	IntegerPrecision 		:= IntegerDigitsMember1 * (IntegerMember2 + 1)	//int64
	//Observation, if above values are greater than uint32, the conversion below ov overflows.
	//This is however expected never to happen.
	TotalPowerPrecision		:= uint32(IntegerPrecision) + CryptoplasmMaxMathPrecision

	cc := c.WithPrecision(TotalPowerPrecision)
	_, _ = cc.Pow(result, member1, member2)

	return result
}
//================================================
//
// Function 05.01 - DIVx
//
// DIVx divides two decimals within a custom Precision modified CryptoplasmPrecisionContext Context
func DIVx(TotalDecimalPrecision uint32, member1, member2 *p.Decimal) *p.Decimal {
    var result = new(p.Decimal)
    cc := c.WithPrecision(TotalDecimalPrecision)
    _, _ = cc.Quo(result, member1, member2)
    return result
}
//================================================
//
// Function 05.02 - DIVs
//
// DIVs divides two decimals within CryptoplasmPrecisionContext Context
func DIVs(member1, member2 *p.Decimal) *p.Decimal {
    var result = new(p.Decimal)
    _, _ = c.Quo(result, member1, member2)
    return result
}
//================================================
//
// Function 05.02 - DIVxc
//
// DIVxc divides 2 numbers with elastic integer precision and 100 max decimal precision
func DIVxc(member1, member2 *p.Decimal) *p.Decimal {
	var (
		result = new(p.Decimal)
		IntegerPrecision uint32
	)

	IntegerDigitsMember1 := Count4Coma(member1)			//int64		//Number of integer digits
	IntegerDigitsMember2 := Count4Coma(member2)			//int64		//Number of integer digits

	DecimalDigitsMember1 := 0 - member1.Exponent		//int32		//Number of decimals digits
	DecimalDigitsMember2 := 0 - member2.Exponent		//int32		//Number of decimals digits

	NumberDigitsMember1 := member1.NumDigits()			//Total Number of digits
	NumberDigitsMember2 := member2.NumDigits()			//Total Number of digits

	IntegerMember1 := RemoveDecimals(member1)			//Integer Value without decimals
	IntegerMember2 := RemoveDecimals(member2)			//Integer Value without decimals

	if DecimalGreaterThan(IntegerMember1,p.NFI(0)) == true  && DecimalGreaterThan(IntegerMember2,p.NFI(0)) {
		//Case 1 Integer Part is similar
		//fmt.Println("Case1")
		if DecimalEqual(IntegerMember1,IntegerMember2) == true {
			//fmt.Println("Case1.1")
			if DecimalGreaterThanOrEqual(member1,member2) == true {
				//fmt.Println("Case1.1.1")
				IntegerPrecision = 1
			} else {
				//fmt.Println("Case1.1.2")
				IntegerPrecision = 0
			}
		} else if DecimalGreaterThan(IntegerMember1,IntegerMember2) == true {
			//fmt.Println("Case1.2")
			if IntegerDigitsMember1 == IntegerDigitsMember2 {
				//fmt.Println("Case1.2.1")
				IntegerPrecision = 1
			} else if IntegerDigitsMember1 > IntegerDigitsMember2 {
				//fmt.Println("Case1.2.2")
				IntegerPrecision = uint32(IntegerDigitsMember1) - uint32(IntegerDigitsMember2) + 1
				//fmt.Println("IntegerPrecision is",IntegerPrecision)
			}
		} else  {
			//fmt.Println("Case1.3")
			IntegerPrecision = 0
		}
	} else if DecimalGreaterThan(IntegerMember1,p.NFI(0)) == true && DecimalEqual(IntegerMember2,p.NFI(0)) {
		//Case 2 Integer Part of member2 is zero
		fmt.Println("Case2")
		if int32(NumberDigitsMember2) == DecimalDigitsMember2 {
			//fmt.Println("Case2.1")
			IntegerPrecision = uint32(IntegerDigitsMember1) + 1
		} else {
			//fmt.Println("Case2.2")
			Zeros := DecimalDigitsMember2 - int32(NumberDigitsMember2)
			IntegerPrecision = uint32(IntegerDigitsMember1) + 1 + uint32(Zeros)
		}
	} else if  DecimalGreaterThan(IntegerMember2,p.NFI(0)) == true && DecimalEqual(IntegerMember1,p.NFI(0)) {
		//Case 3 Integer Part of member1 is zero
		//fmt.Println("Case3")
		IntegerPrecision = 0
	} else if DecimalEqual(IntegerMember1,p.NFI(0)) && DecimalEqual(IntegerMember2,p.NFI(0)) {
		//Case 4 both Integer Parts are zero
		//fmt.Println("Case4")
		Zeros1 := DecimalDigitsMember1 - int32(NumberDigitsMember1)
		Zeros2 := DecimalDigitsMember2 - int32(NumberDigitsMember2)
		if Zeros1 < Zeros2 {
			//fmt.Println("Case4.1")
			IntegerPrecision = uint32(Zeros2 - Zeros1) + 1
		} else if Zeros1 > Zeros2 {
			//fmt.Println("Case4.2")
			IntegerPrecision = 0
		} else if Zeros1 == Zeros2 {
			//fmt.Println("Case4.3")
			if DecimalLessThan(member1,member2) == true {
				//fmt.Println("Case4.3.1")
				IntegerPrecision = 0
			} else {
				//fmt.Println("Case4.3.2")
				IntegerPrecision = 1
			}
		}
	}

	TotalDivisionPrecision := IntegerPrecision + CryptoplasmMaxMathPrecision
	result = DIVx(TotalDivisionPrecision,member1,member2)
	return result
}

//================================================
//
// Function 05.04 - DivInt
//
// DivInt returns the integer part of x divided by y
// It is equal to x // y
// Returned Value is also of decimal Type
func DivInt (member1, member2 *p.Decimal) *p.Decimal {
    var result = new(p.Decimal)
	DCP := SummedMaxLengthPlusOne(member1,member2)	//DivisionContextPrecision
    cc := c.WithPrecision(DCP)
    _,_ = cc.QuoInteger(result,member1,member2)
    return result
}
//================================================
//
// Function 05.05 - DivMod
//
// DivMod returns the remainder from the division of x to y
// It is equal to x % y
// Returned Value is also of decimal Type
func DivMod (member1, member2 *p.Decimal) *p.Decimal {
    var result = new(p.Decimal)
	DCP := SummedMaxLengthPlusOne(member1,member2)	//DivisionContextPrecision
	divresult := TruncateCustom(DivInt(member1,member2),0)
    result = SUBx(DCP,member1,MULx(DCP,member2,divresult))
    return result
}
//================================================
//	05a Mean Functions:
// 		Different types of means used for computing purposes
//		In specific ways
//================================================
//
// Function 05a.01 - TwoMean
//
// TwoMean returns the mean of two decimals
func TwoMean(member1, member2 *p.Decimal) *p.Decimal {
	var result = new(p.Decimal)
	DCP := SummedMaxLengthPlusOne(member1,member2)	//DivisionContextPrecision
	result = DIVx(DCP,ADDxc(member2,member2),p.NFI(2))
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
func TruncateCustom(Number *p.Decimal, DecimalPrecision uint32) *p.Decimal {
	var result = new(p.Decimal)

	NumberDigits := Count4Coma(Number)
	TruncatingContextPrecision := uint32(NumberDigits) + DecimalPrecision
	cc := c.WithPrecision(TruncatingContextPrecision)

	CSP := 0 - int32(DecimalPrecision)
	_, _ = cc.Quantize(result, Number, CSP)
	return result
}

//================================================
//
// Function 06.02 - TruncSeed
//
// TruncSeed truncates the decimal to CryptoplasmSeedPrecision
func TruncSeed(SeedNumber *p.Decimal) *p.Decimal {
	return TruncateCustom(SeedNumber,CryptoplasmSeedPrecision)
}
//================================================
//
// Function 06.03 - TruncToCurrency
//
// TruncToCurrency truncates the decimal to CryptoplasmCurrencyPrecision
// It is Context Precision Independent
func TruncToCurrency(Amount2BecomeCP *p.Decimal) *p.Decimal {
	return TruncateCustom(Amount2BecomeCP,CryptoplasmCurrencyPrecision)
}
//================================================
//
// Function 06.03 - TruncPercent
//
// TruncPercent truncates the decimal to CryptoplasmCurrencyPrecision
// It is Context Precision Independent
func TruncPercent(Amount2BecomeCP *p.Decimal) *p.Decimal {
	return TruncateCustom(Amount2BecomeCP,CryptoplasmPercentPrecision)
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
func SumDL(a []*p.Decimal) *p.Decimal {
	var sum = new(p.Decimal)

	for i := 0; i < len(a); i++ {
		sum = ADDs(sum, a[i])
	}
	return sum
}
//================================================
//
// Function 07.02 - LastDE
//
// LastDE short for LastDecimalElement, returns the last element
// in the slice. Equivalent to pythons -1 index
func LastDE(a []*p.Decimal) *p.Decimal {
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
func AppDec(w1, w2 []*p.Decimal) []*p.Decimal {
	w3 := append(w1, w2...)
	return w3
}
//================================================
//
// Function 07.04 - Reverse
//
// Returns the Reverse of the Slice/Lists
func Reverse(a []*p.Decimal) []*p.Decimal {
	var Reversed = make([]*p.Decimal, 0)
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
func RemoveDecimals(Number *p.Decimal) *p.Decimal {
	var Whole = new(p.Decimal)
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
func Count4Coma(Number *p.Decimal) int64 {
	Whole := RemoveDecimals(Number)
	Int64Digits := Whole.NumDigits()	//int64, up to 9223372036854775807
	return Int64Digits
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
func ASApr(TotalDecimalPrecision uint32, angleAlfa, sideC, angleBeta *p.Decimal) (*p.Decimal, *p.Decimal, *p.Decimal) {
    var (
	angleA = angleAlfa
	angleB = angleBeta
	sdC    = sideC
	sdA    = new(p.Decimal)
	sdB    = new(p.Decimal)
	area   = new(p.Decimal)
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
func OverSendLogBase(cpAmount *p.Decimal) *p.Decimal {
	var Base = new(p.Decimal)

	DigitsNumber := Count4Coma(cpAmount)
	DND := p.NFI(DigitsNumber) //making decimal the Digit number
	DigitsOperatingPrecision := Count4Coma(DND)

	Exponent := SUBx(uint32(DigitsOperatingPrecision), DND, p.NFI(2))
	e := p.INT64(Exponent)
	for i := e; i >= 0; i-- {
		idec := p.NFI(i)
		Value := MULx(uint32(DigitsNumber), POWx(uint32(DigitsNumber), p.NFI(10), idec), p.NFI(7))
		Base = ADDx(uint32(DigitsNumber), Base, Value)
	}
	return Base
}
//================================================
//
// Function 10.02 - OVSLogarithm
//
// OVSLogarithm returns the logarithm from "number" in base "base".
func OVSLogarithm(base, number *p.Decimal) *p.Decimal {
	var (
		LogBase   = new(p.Decimal)
		LogNumber = new(p.Decimal)
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
	CustomLog := DIVx(IP, LogNumber, LogBase)
	return CustomLog
}
//================================================
//
// Function 10.03 - CPAmount2StringDecomposer
//
// CPAmount2StringDecomposer creates a string of uint8 from the integer digits
// of a cpAmount, to be used for calculating the TxTax
func CPAmount2StringDecomposer(cpAmount *p.Decimal) []uint8 {
	var NumberSlice []uint8
	tcpAmount := TruncToCurrency(cpAmount)
	tcpAmountInteger := RemoveDecimals(tcpAmount)
	NumberDigits := tcpAmountInteger.NumDigits()
	NumberToManipulate := tcpAmountInteger
	for i := 0; i < int(NumberDigits); i++ {
		Div10 := DIVx(uint32(NumberDigits),NumberToManipulate,p.NFI(10))
		RemoveLast := RemoveDecimals(Div10)
		Multiplication := MULxc(RemoveLast,p.NFI(10))
		LastDigit := SUBxs(NumberToManipulate,Multiplication)
		LastDigitINT64 := p.INT64(LastDigit)
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
func CPTxTaxV2(cpAmount *p.Decimal) (*p.Decimal, *p.Decimal, *p.Decimal) {
	var (
		TxTax = new(p.Decimal)
		MultipliedTax = new(p.Decimal)
	)
	tcpAmount := TruncToCurrency(cpAmount)
	BaseDigitNumberMain := tcpAmount.NumDigits()
	MainDivPrecision := 2*CryptoplasmCurrencyPrecision + uint32(BaseDigitNumberMain)
	//fmt.Println("Computing TxTax")
	if DecimalGreaterThanOrEqual(tcpAmount,p.NFI(10)) == true {
		AmountNumberSlice := CPAmount2StringDecomposer(cpAmount)
		for i := 1; i < len(AmountNumberSlice); i++ {
			BaseStringPoint := "START Computing TxTax"
			StringPoint := strings.Repeat(".",i)
			StringToPrint := BaseStringPoint + StringPoint
			fmt.Print("\r",StringToPrint)

			Number:=POWx(uint32(i)+1,p.NFI(10),p.NFI(int64(i)))
			LogBase := OverSendLogBase(Number)
			ProMille := OVSLogarithm(LogBase,Number)
			BaseDigitNumber := Number.NumDigits()
			DivPrecision := 2*CryptoplasmCurrencyPrecision + uint32(BaseDigitNumber)
			Tax := DIVx(DivPrecision,MULxc(Number,ProMille),p.NFI(1000))
			if AmountNumberSlice[i] == 0 {
				MultipliedTax = p.NFI(0)
			} else {
				MultipliedTax = MULxc(Tax,p.NFI(int64(AmountNumberSlice[i])))
			}
			TxTax = ADDxs(TxTax,MultipliedTax)
		}
	}
	fmt.Println("DONE")
	TxTax = TruncToCurrency(TxTax)
	Recipient := SUBxs(tcpAmount,TxTax)
	CumulativeFeeProMille := TruncToCurrency(DIVx(MainDivPrecision,MULxc(TxTax,p.NFI(1000)),tcpAmount))
	//fmt.Print("===")
	return CumulativeFeeProMille, TxTax, Recipient
}
//================================================
//
// Function 10.05 - OverSendV2
//
// OverSendV2 computes the modified TxTax.
func OverSendV2(cpAmount *p.Decimal) *p.Decimal {
	fmt.Println("START Computing OverSend...")
	tcpAmount := TruncToCurrency(cpAmount)
	_,TxTaxAmount,_ := CPTxTaxV2(tcpAmount)
	OverSend := ADDxs(tcpAmount,TxTaxAmount)
	Iteration := 0
	for DecimalEqual(TxTaxAmount,p.NFI(0)) == false {
		Iteration = Iteration + 1
		_,TxTaxAmount,_ = CPTxTaxV2(TxTaxAmount)
		OverSend = ADDxs(OverSend,TxTaxAmount)
	}
	//fmt.Println("Last TxTax computation:")
	_,LastTxTax,_ := CPTxTaxV2(OverSend)
	LastOverSend := ADDxs(tcpAmount,LastTxTax)
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
func PseudoFiftyFiftyOverSendLong (cpAmount *p.Decimal) *p.Decimal {
	var PseudoOverSend = new(p.Decimal)
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
func PseudoFiftyFiftyOverSendShort(cpAmount, PerfectOverSend *p.Decimal) *p.Decimal {
	var PseudoFFOverSend = new(p.Decimal)
	DCP := SummedMaxLengthPlusOne(cpAmount,PerfectOverSend)	//DivisionContextPrecision

	tcpAmount := TruncToCurrency(cpAmount)
	_, TxTaxMin, _ := CPTxTaxV2(tcpAmount)
	_, TxTaxMax, _ := CPTxTaxV2(PerfectOverSend)

	MeanTx := TwoMean(TxTaxMin,TxTaxMax)
	HalfMeanTx := DIVx(DCP,MeanTx,p.NFI(2))
	PseudoFFOverSend = ADDxs(cpAmount,HalfMeanTx)
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
func TrueFiftyFiftyOverSendLong (cpAmount *p.Decimal) *p.Decimal {
	var TrueFiftyFiftyOverSend = new(p.Decimal)
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
func TrueFiftyFiftyOverSendShort(cpAmount, PerfectOverSend *p.Decimal) *p.Decimal {
	//start := time.Now()
	var (
		TrueFiftyFifty 			= new(p.Decimal)
		MinimumDifference 		= new(p.Decimal)
		YoctoPlasm				= p.NFS("0.000000000000000000000001")
		Iteration       		int                     	//OverSpend Loop Iteration number
	)

	//Setting IP(internal precision) and truncating cpAmount to 24 decimals
	NumberDigits := Count4Coma(PerfectOverSend)
	IP := CryptoplasmCurrencyPrecision + uint32(NumberDigits) + 1
	tcpAmount := TruncToCurrency(cpAmount)

	OverSendFF := PseudoFiftyFiftyOverSendShort(cpAmount,PerfectOverSend)
	_, _, RecipientFF := CPTxTaxV2(OverSendFF)

	TxTaxS := TruncToCurrency(SUBx(IP,OverSendFF,tcpAmount))
	TxTaxR := TruncToCurrency(SUBx(IP,tcpAmount,RecipientFF))

	FluctuatingOverSend := SUMx(IP,TxTaxR,DIVx(IP,SUBx(IP,TxTaxS,TxTaxR),p.NFI(2)),tcpAmount)
	for DecimalNotEqual(TxTaxS,TxTaxR) == true {
		Iteration = Iteration + 1
		if DecimalGreaterThan(TxTaxS,TxTaxR) == true {
			MinimumDifference = SUBx(IP,TxTaxS,TxTaxR)
		} else {
			MinimumDifference = SUBx(IP,TxTaxR,TxTaxS)
		}
		if DecimalEqual(TxTaxS,TxTaxR) == true || DecimalEqual(MinimumDifference,YoctoPlasm) == true{
			break
		}
		if DecimalGreaterThan(TxTaxS,TxTaxR) == true {
			//Loop Down
			fmt.Println("Computing FiftyFiftyOverSend, refining difference...")
			_, _, RFiftyFifty := CPTxTaxV2(FluctuatingOverSend)
			TxTaxS = TruncToCurrency(SUBx(IP,FluctuatingOverSend,tcpAmount))
			TxTaxR = TruncToCurrency(SUBx(IP,tcpAmount,RFiftyFifty))
			FluctuatingOverSend = SUMx(IP,TxTaxR,DIVx(IP,SUBx(IP,TxTaxS,TxTaxR),p.NFI(2)),tcpAmount)
			//fmt.Println("FlucOvs is",FluctuatingOverSend,"TxTaxS is",TxTaxS,"TxTaxSR is",TxTaxR)

		} else if DecimalLessThan(TxTaxS,TxTaxR) == true {
			//Loop Up
			fmt.Println("Computing FiftyFiftyOverSend, refining difference...")
			_, _, RFiftyFifty := CPTxTaxV2(FluctuatingOverSend)

			TxTaxS = TruncToCurrency(SUBx(IP,FluctuatingOverSend,tcpAmount))
			TxTaxR = TruncToCurrency(SUBx(IP,tcpAmount,RFiftyFifty))
			FluctuatingOverSend = SUMx(IP,TxTaxR,DIVx(IP,SUBx(IP,TxTaxS,TxTaxR),p.NFI(2)),tcpAmount)
			//fmt.Println("FlucOvs is",FluctuatingOverSend,"TxTaxS is",TxTaxS,"TxTaxSR is",TxTaxR)
		}
	}
	TrueFiftyFifty = FluctuatingOverSend
	_, _, RecipientFiftyFifty := CPTxTaxV2(TrueFiftyFifty)

	TxTaxS = SUBx(IP,TrueFiftyFifty,tcpAmount)
	TxTaxR = SUBx(IP,tcpAmount,RecipientFiftyFifty)
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
func TxTaxPrinter(cpAmount *p.Decimal) {
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

	SenderTxTax := SUBxs(FiftyFiftyOverSend,tcpAmount)
	RecipientTxTax := SUBxs(tcpAmount,RecipientFF)

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
func TxTaxDisplayOffset (cpAmount1, cpAmount2 *p.Decimal) string {
	var Result string
	Length := len(CPAmountConv2Print(cpAmount1)) - len(CPAmountConv2Print(cpAmount2))
	Result = strings.Repeat(" ",Length-1)
	return Result
}
//
// Function 10.07b - TxTaxDisplayOffsetS
//
// TxTaxDisplayOffsetS in an auxiliary TxTaxPrinter function
func TxTaxDisplayOffsetS (cpAmount1 *p.Decimal, AmountAsString string) string {
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
func CPConvert2AU(cpAmount *p.Decimal) *p.Decimal {
    tcpAmount := TruncToCurrency(cpAmount)
    NumberDigits := Count4Coma(cpAmount)
    IP := uint32(NumberDigits) + CryptoplasmCurrencyPrecision
    AU := MULx(IP,tcpAmount,AUs)

    return AU
}
//================================================
//
// Function 11.02 - YoctoPlasm2String
//
// YoctoPlasm2String converts a CryptoPlasm AUs (YoctoPlasms)
// into a slice of strings
func YoctoPlasm2String(Number *p.Decimal) []string {
    var SliceStr []string
    Ten := p.NFI(10)
    AuDigits := Number.NumDigits()
    Exp := AuDigits - 1
    IP 	:= uint32(AuDigits)
    //Exp := p.NFI(NumberDigitsAU - 1)
    ToSequence := Number
    for i := Exp; i >= 0; i-- {
        idec := p.NFI(i)
        Power := POWx(IP,Ten,idec)
        Division := DIVx(IP,ToSequence,Power)
		DigitIs := TruncateCustom(Division,0)
        DI := p.INT64(DigitIs)
        DigitIsString := strconv.Itoa(int(DI))
        SliceStr = append(SliceStr,DigitIsString)

        Rest := SUBx(IP,Division,DigitIs)
        SmallAU := MULx(IP,Rest,Power)
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
func CPAmountConv2Print (cpAmount *p.Decimal) string {
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
//================================================
//
// Function 11.03 - BHAmountConv2Print
//
// BHAmountConv2Print converts the BlockHeight decimal into a string
// to be used for printing purposes. A "." is inserted ever 1000.
// For now only a "." as decimal character is implemented.
// Different Schemas can be added (for instance using coma<,> as decimal separator
// instead of point<.>; Or using points for thousand separator,
// or even separating at 2 position for Lakhs and Crores.
func BHAmountConv2Print (BlockHeight *p.Decimal) string {
    var (
    	StringResult 		string
	DigitTier		int32
	PointPosition 		int32
    )

    NumberDigits := Count4Coma(BlockHeight)
    SliceStr := YoctoPlasm2String(BlockHeight)
    TSNumber := NumberDigits / 3

    //Computing the 1000 Separator positions
    if NumberDigits  % 3 == 0 {
        DigitTier = 0
    } else if NumberDigits  % 3 == 1 {
        DigitTier = 1
    } else if NumberDigits  % 3 == 2 {
        DigitTier = 2
    }

    //Adding the 1000 Separator as points
    for i := 1; i<=int(TSNumber); i++ {
        PointPosition = (int32(i)-1) * 4 + DigitTier
        SliceStr = append(SliceStr,"")
        copy(SliceStr[PointPosition+1:],SliceStr[PointPosition:])
        SliceStr[PointPosition] = "."
    }

    //Remove the first element from the Slice if it is a "."
    if SliceStr[0] == "." {
        SliceStr = append(SliceStr[:0], SliceStr[1:]...)
    }

    //Adding Brackets at the beginning and end of the Slice of strings.
    ElementToAppendStart 	:= "["
    ElementToAppendEnd 		:= "]"

    //Appending on First Position
    SliceStr = append(SliceStr,"")
    copy(SliceStr[1:],SliceStr[0:])
    SliceStr[0] = ElementToAppendStart

    //Appending on Last Position
    SliceStr = append(SliceStr,ElementToAppendEnd)

    //Converting the Slice of Strings to a String as final step
    for i := 0; i < len(SliceStr); i++ {
	StringResult = StringResult + SliceStr[i]
    }

    return StringResult
}