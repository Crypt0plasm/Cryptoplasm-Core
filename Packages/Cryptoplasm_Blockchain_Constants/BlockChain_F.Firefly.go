package Cryptoplasm_Blockchain_Constants

import (
	firefly "Cryptoplasm-Core/Packages/Firefly_Precision"
	"fmt"
	"os"
    "strconv"
    "time"
)

var c = CryptoplasmPrecisionContext

//Comparison Functions operation on decimal type
//	1) DecimalEqual			x == y
//	2) DecimalNotEqual		x != y
//	3) DecimalLessThan		x < y
//	4) DecimalLessThanOrEqual	x <= y
//	5) DecimalGreaterThan		x > y
//	6) DecimalGreaterThanOrEqual	x >= y
// The functions use the CryptoplasmPrecisionContext as base Context upon which
// scaling happens to accommodate their size
//================================================
//
// Function 001 - DecimalEqual
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
// Function 002 - DecimalNotEqual
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
// Function 003 - DecimalLessThan
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
// Function 004 - DecimalLessThanOrEqual
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
// Function 005 - DecimalGreaterThan
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
// Function 006 - DecimalGreaterThanOrEqual
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
//
// Function 01 - DivInt
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
// Function 01 - DivMod
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
// Function 01 - DigMax
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

//Basic Operations done under CryptoplasmPrecisionContext without
//Condition and error reporting as supported by Firefly

//Basic Operations done under CryptoplasmPrecisionContext without
//Condition and error reporting as supported by Firefly
//================================================
//
// Function 01 - ADDpr
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
// Function 01a - ADDcp
//
// ADDcp adds two decimals within CryptoplasmPrecisionContext Context
func ADDcp(member1, member2 *firefly.Decimal) *firefly.Decimal {
	var result = new(firefly.Decimal)

	_, _ = c.Add(result, member1, member2)
	return result
}
//================================================
//
// Function 02 - SUMpr
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
// Function 02a - SUMcp
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
// Function 03 - SUBpr
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
// Function 03b - SUBcp
//
// SUBcp subtract two decimals within CryptoplasmPrecisionContext Context
func SUBcp(member1, member2 *firefly.Decimal) *firefly.Decimal {
	var result = new(firefly.Decimal)

	_, _ = c.Sub(result, member1, member2)
	return result
}
//================================================
//
// Function 04a - DIFpr
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
// Function 04b - DIFcp
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
// Function 05a - MULpr
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
// Function 05b - MULcp
//
// MULcp multiplies two decimals within CryptoplasmPrecisionContext Context
func MULcp(member1, member2 *firefly.Decimal) *firefly.Decimal {
	var result = new(firefly.Decimal)

	_, _ = c.Mul(result, member1, member2)
	return result
}
//================================================
//
// Function 06a - PRDpr
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
// Function 06b - PRDcp
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
// Function 07a - POWpr
//
// POWpr multiplies two decimals within a custom Precision modified CryptoplasmPrecisionContext Context
func POWpr(TotalDecimalPrecision uint32, member1, member2 *firefly.Decimal) *firefly.Decimal {
	var result = new(firefly.Decimal)
	cc := c.WithPrecision(TotalDecimalPrecision)
	_, _ = cc.Pow(result, member1, member2)
	return result
}
//================================================
//
// Function 07b - POWcp
//
// POWcp multiplies two decimals within CryptoplasmPrecisionContext Context
func POWcp(member1, member2 *firefly.Decimal) *firefly.Decimal {
	var result = new(firefly.Decimal)

	_, _ = c.Pow(result, member1, member2)
	return result
}
//================================================
//
// Function 08a - DIVpr
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
// Function 08b - DIVcp
//
// DIVcp multiplies two decimals within CryptoplasmPrecisionContext Context
func DIVcp(member1, member2 *firefly.Decimal) *firefly.Decimal {
	var result = new(firefly.Decimal)

	_, _ = c.Quo(result, member1, member2)
	return result
}
//================================================
//
// Function 09a - TruncateCustom
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
// Function 09b - TruncSeed
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
// Function 09c - TruncToCurrency
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
//
// Function 10 - PrintDL
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
// Function 11 - SumDL
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
// Function 12 - LastDE
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
// Function 13 - AppDec
//
// AppDec creates a new bigger slice from the 2 slices used as input
// slices must be made of decimals
func AppDec(w1, w2 []*firefly.Decimal) []*firefly.Decimal {
	w3 := append(w1, w2...)
	return w3
}
//================================================
//
// Function 14 - Reverse
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
// Function 15 - WriteList
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
//
// Function 16 - RemoveDecimals
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
// Function 16b - Count4Coma
//
// Count4Coma returns the number of digits before precision
func Count4Coma(Number *firefly.Decimal) int64 {
	Whole := RemoveDecimals(Number)
	Digits := Whole.NumDigits()
	return Digits
}

//================================================
//
// Function 17 - ASApr
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
//
// Function 18 - OverspendLog
//
// OverspendLog returns the logarithm base required for computing overspend
func OverspendLogBase(cpAmount *firefly.Decimal) *firefly.Decimal {
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
// Function 18b - Logarithm
//
// Logaritm returns the logarithm from "number" in base "base"
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
// Function 19a - CPSend
//
// CPSend computes the AmountFee-FeeProMille, the AmountFee and
// how much the Recipient receives after the AmountFee is deducted from the AmountSent.
func CPSend(cpAmount *firefly.Decimal) (*firefly.Decimal, *firefly.Decimal, *firefly.Decimal) {

	//Setting IP(internal precision) and truncating cpAmount to 24 decimals
	NumberDigits := Count4Coma(cpAmount)
	//IP := 2 * CryptoplasmCurrencyPrecision + uint32(NumberDigits)
	IP := CryptoplasmCurrencyPrecision + uint32(NumberDigits) + 1
	tcpAmount := TruncToCurrency(cpAmount)

	//cpAmountTrunc := TruncCurrencyCustom(IP,cpAmount)

	OVSLogBase := OverspendLogBase(tcpAmount)
	FeeProMille := TruncToCurrency(OVSLogarithm(OVSLogBase, tcpAmount))
	AmountFee := TruncToCurrency(DIVpr(IP, MULpr(IP, tcpAmount, FeeProMille), firefly.NFI(1000)))
	Recipient := TruncToCurrency(SUBpr(IP, tcpAmount, AmountFee))

	return FeeProMille, AmountFee, Recipient
}
//================================================
//
// Function 19b - CPOverSend
//
// CPOverSend computes the AmountFee-FeeProMille, the AmountFee and
// how much the Recipient receives after the AmountFee is deducted from the AmountSent.
func CPOverSend(cpAmount *firefly.Decimal) *firefly.Decimal {
	start := time.Now()
	var (
		PerfectOverSend = new(firefly.Decimal)
		OverSend        = new(firefly.Decimal)
		OSA             = new(firefly.Decimal) //OverSendArgument
		OSIA            = new(firefly.Decimal) //OverSendIterationArgument
		FPM             = new(firefly.Decimal) //AmountFee-ProMille
		AF              = new(firefly.Decimal) //AmountFee
		R               = new(firefly.Decimal) //Recipient
		OsiaPrec        uint32                 //Number of Digits after comma OSIA must have to be computed
		Iteration       int                    //OverSpend Loop Iteration number
	)

	//Setting IP(internal precision) and truncating cpAmount to 24 decimals
	NumberDigits := Count4Coma(cpAmount)
	IP := CryptoplasmCurrencyPrecision + uint32(NumberDigits) + 1
	tcpAmount := TruncToCurrency(cpAmount)

	//OsiaDivPrec might need more Precision than this, so it gets its own precision
	//OsiaPrec starts with 0, because OSAT is negative
	//OsiaPrec is equal to the positive OSAT, as OSAT increases, so does OsiaPrec
	//OsiaDivPrec starts with 3
	OsiaDivPrec := uint32(3)

	//The Logarithm Base (777...777) that is used for computing OverSend given amount cpAmountTrunc
	OvSLogBase := OverspendLogBase(tcpAmount)

	OSAT := -2 //int type OverSend-Argument-Tier, starts at -2
	//OSAT at -2 needs initially addition precision of 3 (-3 would need 4)
	//Once OSAT is set, or newly set, OverSendArgument (OSA) is computed from it:
	OSA = DIVpr(IP, firefly.NFI(1), POWpr(IP, firefly.NFI(10), firefly.NFI(int64(OSAT))))
	//From OSA the OSIA is derived which is the OverSendArgument being iterated, hence OverSendIterationArgument
	//OSIA is always truncated by the OsiaPrec, thus is grows as needed
	OSIA = TruncateCustom(IP, OSA, OsiaPrec)
    	OsiaOffset := OverSpendDisplayFormat(OSIA)
	//Above, it is truncated at zero decimals, because it starts without decimals, and gains them while looping

	MaxNoOverSend, _, _ := firefly.NewFromString("9.999999999999999999999999")
	//CompareResult1 := DecimalLessThanOrEqual(tcpAmount,MaxNoOverSend)
	if DecimalLessThanOrEqual(tcpAmount,MaxNoOverSend) == true {
		R = tcpAmount
		PerfectOverSend = tcpAmount
		//No extra CP is required as "OverSend" when transacted amount is below 10 CP
	} else if DecimalLessThanOrEqual(tcpAmount,MaxNoOverSend) == false {
		FPMo := TruncToCurrency(OVSLogarithm(OvSLogBase, tcpAmount))
		FPM = FPMo
		AFo := TruncToCurrency(DIVpr(IP, MULpr(IP, tcpAmount, FPM), firefly.NFI(1000)))
		AF = AFo
		Ro := TruncToCurrency(SUBpr(IP, tcpAmount, AF))
		R = Ro

		Iteration = 1
		OverSend = TruncToCurrency(ADDpr(IP, tcpAmount, MULpr(IP, AFo, ADDpr(IP, firefly.NFI(1), DIVpr(OsiaDivPrec, OSIA, firefly.NFI(1000))))))
		LoopDirection := "up"

		for DecimalNotEqual(R,tcpAmount) == true {
			Iteration = Iteration + 1
			if DecimalEqual(R,tcpAmount) == true {
				//fmt.Println("")
				//fmt.Println("Computing done after",Iteration,"iterations")
				break
			}
			if DecimalLessThan(R,tcpAmount) == true {
				if LoopDirection == "down" {
					OSIA = TruncateCustom(OsiaDivPrec, ADDpr(OsiaDivPrec, OSIA, OSA), OsiaPrec)
				    	OsiaOffset = OverSpendDisplayFormat(OSIA)
					//
					OSAT = OSAT + 1
					//Setting Precisions derived from OSAT, OsiaPrec and OsiaDivPrec
					if OSAT > 0 {
						OsiaPrec = uint32(OSAT)
					}
					if OsiaPrec > 0 {
						OsiaDivPrec = OsiaPrec + 3
					}
					//
					OSA = DIVpr(OsiaDivPrec, firefly.NFI(1), POWpr(OsiaDivPrec, firefly.NFI(10), firefly.NFI(int64(OSAT))))
				}
				LoopDirection = "up"
				OvSLogBase = OverspendLogBase(OverSend)

				FPM = TruncToCurrency(OVSLogarithm(OvSLogBase, OverSend))
				AF = TruncToCurrency(DIVpr(IP, MULpr(IP, OverSend, FPM), firefly.NFI(1000)))
				R = TruncToCurrency(SUBpr(IP, OverSend, AF))

				OSIA = TruncateCustom(OsiaDivPrec, ADDpr(OsiaDivPrec, OSIA, OSA), OsiaPrec)
			    	OsiaOffset = OverSpendDisplayFormat(OSIA)
				//Troubleshooting Comments can be commented away
				//fmt.Println("Iteration ",Iteration,"OSIA is",OSIA,"R is", R)
			    	fmt.Println("Computing Tx Tax, refining argument...",OsiaOffset,OSIA)
				OverSend = TruncToCurrency(ADDpr(IP, tcpAmount, MULpr(IP, AFo, ADDpr(IP, firefly.NFI(1), DIVpr(OsiaDivPrec, OSIA, firefly.NFI(1000))))))
			} else if DecimalGreaterThan(R,tcpAmount) == true {
				if LoopDirection == "up" {
					OSIA = TruncateCustom(OsiaDivPrec, SUBpr(OsiaDivPrec, OSIA, OSA), OsiaPrec)
				    	OsiaOffset = OverSpendDisplayFormat(OSIA)
					//
					OSAT = OSAT + 1
					//Setting Precisions derived from OSAT, OsiaPrec and OsiaDivPrec
					if OSAT > 0 {
						OsiaPrec = uint32(OSAT)
					}
					if OsiaPrec > 0 {
						OsiaDivPrec = OsiaPrec + 3
					}
					//
					OSA = DIVpr(OsiaDivPrec, firefly.NFI(1), POWpr(OsiaDivPrec, firefly.NFI(10), firefly.NFI(int64(OSAT))))
				}
				LoopDirection = "down"
				OvSLogBase = OverspendLogBase(OverSend)

				FPM = TruncToCurrency(OVSLogarithm(OvSLogBase, OverSend))
				AF = TruncToCurrency(DIVpr(IP, MULpr(IP, OverSend, FPM), firefly.NFI(1000)))
				R = TruncToCurrency(SUBpr(IP, OverSend, AF))

				OSIA = TruncateCustom(OsiaDivPrec, SUBpr(OsiaDivPrec, OSIA, OSA), OsiaPrec)
			    	OsiaOffset = OverSpendDisplayFormat(OSIA)
				//Troubleshooting Comments can be commented away
				//fmt.Println("Iteration ",Iteration,"OSIA is",OSIA,"R is", R)
				fmt.Println("Computing Tx Tax, refining argument...",OsiaOffset,OSIA)
				OverSend = TruncToCurrency(ADDpr(IP, tcpAmount, MULpr(IP, AFo, ADDpr(IP, firefly.NFI(1), DIVpr(OsiaDivPrec, OSIA, firefly.NFI(1000))))))
			}
		}
	}
	//Refining OverSend to PerfectOverSend
	_, _, v3 := CPSend(OverSend)
	//Computes the AU, see bellow, only for observation purposes
	//Must be commented away
	//AtomicUnit := firefly.New(1,-24)
	if DecimalEqual(tcpAmount,v3) == true {
		PerfectOverSend = OverSend
		//Troubleshooting Comments can be commented away
		//fmt.Println("Computed OverSend resulted in the expected Receiver up to the last 24th decimal")
		//fmt.Println("")
		//fmt.Println("Computed OverSend was",OverSend)
		//fmt.Println("This is equal to PerfectOverSend")
		//fmt.Println("")
	} else {
		if DecimalLessThan(tcpAmount,v3) == true {
			U := DIFpr(IP, v3, tcpAmount)
			//The Number of Atomic Units by which the computed OverSend is off
			//Is only used in the comments for observation purpose
			//AU := DIVpr(IP,U,AtomicUnit)
			PerfectOverSend = DIFpr(IP, OverSend, U)
			//Troubleshooting Comments can be commented away
			//fmt.Println("Computed OverSend was off by ", AU,"AU(s) more")
			//fmt.Println("")
			//fmt.Println("Computed OverSend was      ",OverSend)
			//fmt.Println("Computed PerfectOverSend is",PerfectOverSend)
			//fmt.Println("")
		} else if DecimalGreaterThan(tcpAmount,v3) == true {
			U := DIFpr(IP, tcpAmount, v3)
			//The Number of Atomic Units by which the computed OverSend is off
			//Is only used in the comments for observation purpose
			//AU := DIVpr(IP,U,AtomicUnit)
			PerfectOverSend = ADDpr(IP, OverSend, U)
			//Troubleshooting Comments can be commented away
			//fmt.Println("Computed OverSend was off by ", AU,"AU(s) less")
			//fmt.Println("")
			//fmt.Println("Computed OverSend was      ",OverSend)
			//fmt.Println("Computed PerfectOverSend is",PerfectOverSend)
			//fmt.Println("")
		}
	}
	//Verification
	//fmt.Println("")
	//fmt.Println("OverSend for:", tcpAmount, "CP")
	//fmt.Println("Is equal to:.", PerfectOverSend, "CP")
	//fmt.Println("")
	//fmt.Println("VERIFICATION:")
	//FPMx, AFx, Rx := CPSend(PerfectOverSend)
	//fmt.Println("Sending",PerfectOverSend,"CP, will cost an Amount-Fee of",FPMx,"promille")
	//fmt.Println("The Amount-Fee will therefore amount to",AFx,"CP")
	//fmt.Println("Sending",PerfectOverSend,"CP, the Recipient gets",Rx,"CP")

	elapsed := time.Since(start)
	fmt.Println("")
	fmt.Println("Computing OverSend took", elapsed, "with", Iteration, "Iterations")
	return PerfectOverSend
}
//================================================
//
// Function 19b - OverSpendDisplayFormat
//
// OverSpendDisplayFormat computes the length of a string needed
// in order to properly display when OverSend is running
func OverSpendDisplayFormat(Number *firefly.Decimal) string {
    var Result string
    NumberDigits := Count4Coma(Number)
    Negative := Number.Negative
    if Negative == false {
        if NumberDigits == 3 {
            Result = "  "
	} else if NumberDigits == 2 {
	    Result = "   "
	} else {
	    Result = "    "
	}
    } else {
	if NumberDigits == 3 {
	    Result = " "
	} else if NumberDigits == 2 {
	    Result = "  "
	} else {
	    Result = "   "
	}
    }
    return Result
}

//================================================
//
// Function 20 - CPConvert2AU
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
// Function 20b - YoctoPlasm2String
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
// Function 20c - CPAmountConv2Print
//
// CPAmountConv2Print converts CryptoPlasm amount into a string
// to be used for printing purposes.
// For now only a "." as decimal character is implemented.
// Different Schemas can be added (for instance using coma<,> as decimal separator
// instead of point<.>; Or using points for thousand separator,
// or even separating at 2 position for Lakhs and Crores.
func CPAmountConv2Print (cpAmount *firefly.Decimal) string {
    var (
    	StringResult string
	ComaPosition int32
    )
    AU := CPConvert2AU(cpAmount)
    SliceStr := YoctoPlasm2String(AU)
    NumberDigits := Count4Coma(AU)

    if NumberDigits <= 25 {
	ComaPosition = 1
    } else {
	ComaPosition = int32(NumberDigits) - 24
    }

    SliceStr = append(SliceStr,"")
    copy(SliceStr[ComaPosition+1:],SliceStr[ComaPosition:])
    SliceStr[ComaPosition] = ","

    for i := 0; i < len(SliceStr); i++ {
        StringResult = StringResult + SliceStr[i]
    }
    return StringResult
}