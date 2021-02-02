package constants

import (
	"fmt"
	"os"
	"time"

	"github.com/Cryptoplasm-Core/precision"
)

var c = CryptoplasmPrecisionContext

func ASApr(TotalDecimalPrecision uint32, angleAlfa, sideC, angleBeta *precision.Decimal) (*precision.Decimal, *precision.Decimal, *precision.Decimal) {
	var (
		angleA = angleAlfa
		angleB = angleBeta
		sdC    = sideC
		sdA    = new(precision.Decimal)
		sdB    = new(precision.Decimal)
		area   = new(precision.Decimal)
	)
	cc := c.WithPrecision(TotalDecimalPrecision)
	_, sdA, sdB, area = cc.ASA(angleA, sdC, angleB)
	//ASAcp doesnt return angleG as this isn't used in the computation
	return sdA, sdB, area
}

//Basic Operations done under CryptoplasmPrecisionContext without
//Condition and error reporting as supported by Firefly
//================================================
//
// Function 01 - ADDpr
//
// ADDpr adds two decimals within a custom Precision modified CryptoplasmPrecisionContext Context
func ADDpr(TotalDecimalPrecision uint32, member1, member2 *precision.Decimal) *precision.Decimal {
	var result = new(precision.Decimal)
	cc := c.WithPrecision(TotalDecimalPrecision)
	_, _ = cc.Add(result, member1, member2)
	return result
}

//================================================
//
// Function 01a - ADDcp
//
// ADDcp adds two decimals within CryptoplasmPrecisionContext Context
func ADDcp(member1, member2 *precision.Decimal) *precision.Decimal {
	var result = new(precision.Decimal)

	_, _ = c.Add(result, member1, member2)
	return result
}

//================================================
//
// Function 02 - SUMpr
//
// SUMpr adds multiple decimals within a custom Precision modified CryptoplasmPrecisionContext Context
func SUMpr(TotalDecimalPrecision uint32, first *precision.Decimal, rest ...*precision.Decimal) *precision.Decimal {
	var (
		sum     = new(precision.Decimal)
		restsum = precision.NFI(0)
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
func SUMcp(first *precision.Decimal, rest ...*precision.Decimal) *precision.Decimal {
	var (
		sum     = new(precision.Decimal)
		restsum = precision.NFI(0)
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
func SUBpr(TotalDecimalPrecision uint32, member1, member2 *precision.Decimal) *precision.Decimal {
	var result = new(precision.Decimal)
	cc := c.WithPrecision(TotalDecimalPrecision)
	_, _ = cc.Sub(result, member1, member2)
	return result
}

//================================================
//
// Function 03b - SUBcp
//
// SUBcp subtract two decimals within CryptoplasmPrecisionContext Context
func SUBcp(member1, member2 *precision.Decimal) *precision.Decimal {
	var result = new(precision.Decimal)

	_, _ = c.Sub(result, member1, member2)
	return result
}

//================================================
//
// Function 04a - DIFpr
//
// DIFpr subtract multiple decimals within a custom Precision modified CryptoplasmPrecisionContext Context
func DIFpr(TotalDecimalPrecision uint32, first *precision.Decimal, rest ...*precision.Decimal) *precision.Decimal {
	var (
		difference = new(precision.Decimal)
		restsum    = precision.NFI(0)
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
func DIFcp(first *precision.Decimal, rest ...*precision.Decimal) *precision.Decimal {
	var (
		difference = new(precision.Decimal)
		restsum    = precision.NFI(0)
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
func MULpr(TotalDecimalPrecision uint32, member1, member2 *precision.Decimal) *precision.Decimal {
	var result = new(precision.Decimal)
	cc := c.WithPrecision(TotalDecimalPrecision)
	_, _ = cc.Mul(result, member1, member2)
	return result
}

//================================================
//
// Function 05b - MULcp
//
// MULcp multiplies two decimals within CryptoplasmPrecisionContext Context
func MULcp(member1, member2 *precision.Decimal) *precision.Decimal {
	var result = new(precision.Decimal)

	_, _ = c.Mul(result, member1, member2)
	return result
}

//================================================
//
// Function 06a - PRDpr
//
// PRDpr multiplies multiple decimals within a custom Precision modified CryptoplasmPrecisionContext Context
func PRDpr(TotalDecimalPrecision uint32, first *precision.Decimal, rest ...*precision.Decimal) *precision.Decimal {
	var (
		product     = new(precision.Decimal)
		restproduct = precision.NFI(1)
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
func PRDcp(first *precision.Decimal, rest ...*precision.Decimal) *precision.Decimal {
	var (
		product     = new(precision.Decimal)
		restproduct = precision.NFI(1)
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
func POWpr(TotalDecimalPrecision uint32, member1, member2 *precision.Decimal) *precision.Decimal {
	var result = new(precision.Decimal)
	cc := c.WithPrecision(TotalDecimalPrecision)
	_, _ = cc.Pow(result, member1, member2)
	return result
}

//================================================
//
// Function 07b - POWcp
//
// POWcp multiplies two decimals within CryptoplasmPrecisionContext Context
func POWcp(member1, member2 *precision.Decimal) *precision.Decimal {
	var result = new(precision.Decimal)

	_, _ = c.Pow(result, member1, member2)
	return result
}

//================================================
//
// Function 08a - DIVpr
//
// DIVpr multiplies two decimals within a custom Precision modified CryptoplasmPrecisionContext Context
func DIVpr(TotalDecimalPrecision uint32, member1, member2 *precision.Decimal) *precision.Decimal {
	var result = new(precision.Decimal)
	cc := c.WithPrecision(TotalDecimalPrecision)
	_, _ = cc.Quo(result, member1, member2)
	return result
}

//================================================
//
// Function 08b - DIVcp
//
// DIVcp multiplies two decimals within CryptoplasmPrecisionContext Context
func DIVcp(member1, member2 *precision.Decimal) *precision.Decimal {
	var result = new(precision.Decimal)

	_, _ = c.Quo(result, member1, member2)
	return result
}

//================================================
//
// Function 09a - TruncateCustom
//
// TruncateCustom truncates the decimal to the precision number
func TruncateCustom(TotalDecimalPrecision uint32, number *precision.Decimal, DecimalPrecision uint32) *precision.Decimal {
	var result = new(precision.Decimal)
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
func TruncSeed(SeedNumber *precision.Decimal) *precision.Decimal {
	var result = new(precision.Decimal)

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
func TruncToCurrency(Amount2BecomeCP *precision.Decimal) *precision.Decimal {
	var result = new(precision.Decimal)

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
func SumDL(a []*precision.Decimal) *precision.Decimal {
	var sum = new(precision.Decimal)

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
func LastDE(a []*precision.Decimal) *precision.Decimal {
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
func AppDec(w1, w2 []*precision.Decimal) []*precision.Decimal {
	w3 := append(w1, w2...)
	return w3
}

//================================================
//
// Function 14 - Reverse
//
// Returns the Reverse of the Slice/Lists
func Reverse(a []*precision.Decimal) []*precision.Decimal {
	var Reversed = make([]*precision.Decimal, 0)
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
func RemoveDecimals(Number *precision.Decimal) *precision.Decimal {
	var Whole = new(precision.Decimal)
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
func Count4Coma(Number *precision.Decimal) int64 {
	Whole := RemoveDecimals(Number)
	Digits := Whole.NumDigits()
	return Digits
}

//================================================
//
// Function 17 - OverspendLog
//
// OverspendLog returns the logarithm base required for computing overspend
func OverspendLogBase(cpAmount *precision.Decimal) *precision.Decimal {
	var Base = new(precision.Decimal)

	DigitsNumber := Count4Coma(cpAmount)
	DND := precision.NFI(DigitsNumber) //making decimal the Digit number
	DigitsOperatingPrecision := Count4Coma(DND)

	Exponent := SUBpr(uint32(DigitsOperatingPrecision), DND, precision.NFI(2))
	e, _ := Exponent.Int64()
	for i := e; i >= 0; i-- {
		idec := precision.NFI(i)
		Value := MULpr(uint32(DigitsNumber), POWpr(uint32(DigitsNumber), precision.NFI(10), idec), precision.NFI(7))
		Base = ADDpr(uint32(DigitsNumber), Base, Value)
	}
	return Base
}

//================================================
//
// Function 17b - Logarithm
//
// Logaritm returns the logarithm from "number" in base "base"
func OVSLogarithm(base, number *precision.Decimal) *precision.Decimal {
	var (
		LogBase   = new(precision.Decimal)
		LogNumber = new(precision.Decimal)
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
// Function 18a - CPSend
//
// CPSend computes the AmountFee-FeeProMille, the AmountFee and
// how much the Recipient receives after the AmountFee is deducted from the AmountSent.
func CPSend(cpAmount *precision.Decimal) (*precision.Decimal, *precision.Decimal, *precision.Decimal) {

	//Setting IP(internal precision) and truncating cpAmount to 24 decimals
	NumberDigits := Count4Coma(cpAmount)
	//IP := 2 * CryptoplasmCurrencyPrecision + uint32(NumberDigits)
	IP := CryptoplasmCurrencyPrecision + uint32(NumberDigits) + 1
	tcpAmount := TruncToCurrency(cpAmount)

	//cpAmountTrunc := TruncCurrencyCustom(IP,cpAmount)

	OVSLogBase := OverspendLogBase(tcpAmount)
	FeeProMille := TruncToCurrency(OVSLogarithm(OVSLogBase, tcpAmount))
	AmountFee := TruncToCurrency(DIVpr(IP, MULpr(IP, tcpAmount, FeeProMille), precision.NFI(1000)))
	Recipient := TruncToCurrency(SUBpr(IP, tcpAmount, AmountFee))

	return FeeProMille, AmountFee, Recipient
}

//================================================
//
// Function 18b - CPOverSend
//
// CPOverSend computes the AmountFee-FeeProMille, the AmountFee and
// how much the Recipient receives after the AmountFee is deducted from the AmountSent.
func CPOverSend(cpAmount *precision.Decimal) *precision.Decimal {
	start := time.Now()
	var (
		PerfectOverSend = new(precision.Decimal)
		OverSend        = new(precision.Decimal)
		OSA             = new(precision.Decimal) //OverSendArgument
		OSIA            = new(precision.Decimal) //OverSendIterationArgument
		FPM             = new(precision.Decimal) //AmountFee-ProMille
		AF              = new(precision.Decimal) //AmountFee
		R               = new(precision.Decimal) //Recipient
		OsiaPrec        uint32                   //Number of Digits after comma OSIA must have to be computed
		Iteration       int                      //OverSpend Loop Iteration number
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
	OSA = DIVpr(IP, precision.NFI(1), POWpr(IP, precision.NFI(10), precision.NFI(int64(OSAT))))
	//From OSA the OSIA is derived which is the OverSendArgument being iterated, hence OverSendIterationArgument
	//OSIA is always truncated by the OsiaPrec, thus is grows as needed
	OSIA = TruncateCustom(IP, OSA, OsiaPrec)
	OsiaOffset := OverSpendDisplayFormat(OSIA)
	//Above, it is truncated at zero decimals, because it starts without decimals, and gains them while looping

	MaxNoOverSend, _, _ := precision.NewFromString("9.999999999999999999999999")
	//AFADP is used because in case the cpAmountTrunc is 0.xxx AFMP would be 24 and that wouldn't be enough
	//In this case 25 would be needed and that would be AFAP
	Difference := SUBpr(IP, tcpAmount, MaxNoOverSend)
	IsThreshold := Difference.IsZero()
	if IsThreshold == true || Difference.Negative == true {
		R = tcpAmount
		PerfectOverSend = tcpAmount
		//No extra CP is required as "OverSend" when transacted amount is below 10 CP
	} else if Difference.Negative == false {
		FPMo := TruncToCurrency(OVSLogarithm(OvSLogBase, tcpAmount))
		FPM = FPMo
		AFo := TruncToCurrency(DIVpr(IP, MULpr(IP, tcpAmount, FPM), precision.NFI(1000)))
		AF = AFo
		Ro := TruncToCurrency(SUBpr(IP, tcpAmount, AF))
		R = Ro

		Iteration = 1
		OverSend = TruncToCurrency(ADDpr(IP, tcpAmount, MULpr(IP, AFo, ADDpr(IP, precision.NFI(1), DIVpr(OsiaDivPrec, OSIA, precision.NFI(1000))))))
		LoopDirection := "up"
		Difference2 := SUBpr(IP, R, tcpAmount)
		IsThreshold2 := Difference2.IsZero()
		for IsThreshold2 == false {
			Iteration = Iteration + 1
			Difference3 := SUBpr(IP, R, tcpAmount)
			IsThreshold3 := Difference3.IsZero()
			if IsThreshold3 == true {
				//fmt.Println("")
				//fmt.Println("Computing done after",Iteration,"iterations")
				break
			}
			if Difference3.Negative == true {
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
					OSA = DIVpr(OsiaDivPrec, precision.NFI(1), POWpr(OsiaDivPrec, precision.NFI(10), precision.NFI(int64(OSAT))))
				}
				LoopDirection = "up"
				OvSLogBase = OverspendLogBase(OverSend)

				FPM = TruncToCurrency(OVSLogarithm(OvSLogBase, OverSend))
				AF = TruncToCurrency(DIVpr(IP, MULpr(IP, OverSend, FPM), precision.NFI(1000)))
				R = TruncToCurrency(SUBpr(IP, OverSend, AF))

				OSIA = TruncateCustom(OsiaDivPrec, ADDpr(OsiaDivPrec, OSIA, OSA), OsiaPrec)
				OsiaOffset = OverSpendDisplayFormat(OSIA)
				//Troubleshooting Comments can be commented away
				//fmt.Println("Iteration ",Iteration,"OSIA is",OSIA,"R is", R)
				fmt.Println("Computing Tx Tax, refining argument...", OsiaOffset, OSIA)
				OverSend = TruncToCurrency(ADDpr(IP, tcpAmount, MULpr(IP, AFo, ADDpr(IP, precision.NFI(1), DIVpr(OsiaDivPrec, OSIA, precision.NFI(1000))))))
			} else if Difference3.Negative == false {
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
					OSA = DIVpr(OsiaDivPrec, precision.NFI(1), POWpr(OsiaDivPrec, precision.NFI(10), precision.NFI(int64(OSAT))))
				}
				LoopDirection = "down"
				OvSLogBase = OverspendLogBase(OverSend)

				FPM = TruncToCurrency(OVSLogarithm(OvSLogBase, OverSend))
				AF = TruncToCurrency(DIVpr(IP, MULpr(IP, OverSend, FPM), precision.NFI(1000)))
				R = TruncToCurrency(SUBpr(IP, OverSend, AF))

				OSIA = TruncateCustom(OsiaDivPrec, SUBpr(OsiaDivPrec, OSIA, OSA), OsiaPrec)
				OsiaOffset = OverSpendDisplayFormat(OSIA)
				//Troubleshooting Comments can be commented away
				//fmt.Println("Iteration ",Iteration,"OSIA is",OSIA,"R is", R)
				fmt.Println("Computing Tx Tax, refining argument...", OsiaOffset, OSIA)
				OverSend = TruncToCurrency(ADDpr(IP, tcpAmount, MULpr(IP, AFo, ADDpr(IP, precision.NFI(1), DIVpr(OsiaDivPrec, OSIA, precision.NFI(1000))))))
			}
		}
	}
	//Refining OverSend to PerfectOverSend
	_, _, v3 := CPSend(OverSend)
	//Computes the AU, see bellow, only for observation purposes
	//Must be commented away
	//AtomicUnit := firefly.New(1,-24)
	Difference4 := SUBpr(IP, tcpAmount, v3)
	IsThreshold4 := Difference4.IsZero()

	if IsThreshold4 == true {
		PerfectOverSend = OverSend
		//Troubleshooting Comments can be commented away
		//fmt.Println("Computed OverSend resulted in the expected Receiver up to the last 24th decimal")
		//fmt.Println("")
		//fmt.Println("Computed OverSend was",OverSend)
		//fmt.Println("This is equal to PerfectOverSend")
		//fmt.Println("")
	} else {
		if Difference4.Negative == true {
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
		} else if Difference4.Negative == false {
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

func OverSpendDisplayFormat(Number *precision.Decimal) string {
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
