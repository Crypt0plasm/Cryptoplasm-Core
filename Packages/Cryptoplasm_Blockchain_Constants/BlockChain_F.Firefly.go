package Cryptoplasm_Blockchain_Constants

import (
	firefly "Cryptoplasm-Core/Packages/Firefly_Precision"
	"fmt"
	"os"
	"time"
)

var c							= CryptoplasmPrecisionContext

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

//Basic Operations done under CryptoplasmPrecisionContext without
//Condition and error reporting as supported by Firefly
//================================================
//
// Function 01 - ADDpr
//
// ADDpr adds two decimals within a custom Precision modified CryptoplasmPrecisionContext Context
func ADDpr(TotalDecimalPrecision uint32, member1, member2 *firefly.Decimal) *firefly.Decimal {
	var result			= new(firefly.Decimal)
	cc := c.WithPrecision(TotalDecimalPrecision)
	_, _ = cc.Add(result,member1,member2)
	return result
}
//================================================
//
// Function 01a - ADDcp
//
// ADDcp adds two decimals within CryptoplasmPrecisionContext Context
func ADDcp(member1, member2 *firefly.Decimal) *firefly.Decimal {
	var result			= new(firefly.Decimal)

	_, _ = c.Add(result,member1,member2)
	return result
}
//================================================
//
// Function 02 - SUMpr
//
// SUMpr adds multiple decimals within a custom Precision modified CryptoplasmPrecisionContext Context
func SUMpr(TotalDecimalPrecision uint32, first *firefly.Decimal, rest ...*firefly.Decimal) *firefly.Decimal {
	var (
		sum 			= new(firefly.Decimal)
		restsum 		= firefly.NFI(0)
	)
	cc := c.WithPrecision(TotalDecimalPrecision)
	for _, item := range rest {
		_, _ = cc.Add(restsum,restsum,item)
	}
	_, _ = cc.Add(sum,first, restsum)
	return sum
}
//================================================
//
// Function 02a - SUMcp
//
// SUMcp adds multiple decimals within CryptoplasmPrecisionContext Context
func SUMcp(first *firefly.Decimal, rest ...*firefly.Decimal) *firefly.Decimal {
	var (
		sum 			= new(firefly.Decimal)
		restsum 		= firefly.NFI(0)
	)

	for _, item := range rest {
		_, _ = c.Add(restsum,restsum,item)
	}
	_, _ = c.Add(sum,first, restsum)
	return sum
}
//================================================
//
// Function 02 - SUBpr
//
// SUBpr subtract two decimals within a custom Precision modified CryptoplasmPrecisionContext Context
func SUBpr(TotalDecimalPrecision uint32, member1, member2 *firefly.Decimal) *firefly.Decimal {
	var result			= new(firefly.Decimal)
	cc := c.WithPrecision(TotalDecimalPrecision)
	_, _ = cc.Sub(result,member1,member2)
	return result
}
//================================================
//
// Function 02b - SUBcp
//
// SUBcp subtract two decimals within CryptoplasmPrecisionContext Context
func SUBcp(member1, member2 *firefly.Decimal) *firefly.Decimal {
	var result			= new(firefly.Decimal)

	_, _ = c.Sub(result,member1,member2)
	return result
}
//================================================
//
// Function 03a - DIFpr
//
// DIFpr subtract multiple decimals within a custom Precision modified CryptoplasmPrecisionContext Context
func DIFpr(TotalDecimalPrecision uint32, first *firefly.Decimal, rest ...*firefly.Decimal) *firefly.Decimal {
	var (
		difference 		= new(firefly.Decimal)
		restsum 		= firefly.NFI(0)
	)
	cc := c.WithPrecision(TotalDecimalPrecision)
	for _, item := range rest {
		_, _ = cc.Add(restsum, restsum,item)
	}
	_, _ = cc.Sub(difference,first, restsum)
	return difference
}
//================================================
//
// Function 03b - DIFcp
//
// DIFcp subtract multiple decimals within CryptoplasmPrecisionContext Context
func DIFcp(first *firefly.Decimal, rest ...*firefly.Decimal) *firefly.Decimal {
	var (
		difference 		= new(firefly.Decimal)
		restsum 		= firefly.NFI(0)
	)
	for _, item := range rest {
		_, _ = c.Add(restsum, restsum,item)
	}
	_, _ = c.Sub(difference,first, restsum)
	return difference
}
//================================================
//
// Function 04 - MULpr
//
// MULpr multiplies two decimals within a custom Precision modified CryptoplasmPrecisionContext Context
func MULpr(TotalDecimalPrecision uint32, member1, member2 *firefly.Decimal) *firefly.Decimal {
	var result			= new(firefly.Decimal)
	cc := c.WithPrecision(TotalDecimalPrecision)
	_, _ = cc.Mul(result,member1,member2)
	return result
}
//================================================
//
// Function 04b - MULcp
//
// MULcp multiplies two decimals within CryptoplasmPrecisionContext Context
func MULcp(member1, member2 *firefly.Decimal) *firefly.Decimal {
	var result			= new(firefly.Decimal)

	_, _ = c.Mul(result,member1,member2)
	return result
}
//================================================
//
// Function 05b - PRDpr
//
// PRDpr multiplies multiple decimals within a custom Precision modified CryptoplasmPrecisionContext Context
func PRDpr(TotalDecimalPrecision uint32, first *firefly.Decimal, rest ...*firefly.Decimal) *firefly.Decimal {
	var (
		product 		= new(firefly.Decimal)
		restproduct 	= firefly.NFI(1)
	)
	cc := c.WithPrecision(TotalDecimalPrecision)
	for _, item := range rest {
		_, _ = cc.Mul(restproduct,restproduct,item)
	}
	_, _ = cc.Mul(product,first,restproduct)

	return product
}
//================================================
//
// Function 05b - PRDcp
//
// PRDcp multiplies multiple decimals within CryptoplasmPrecisionContext Context
func PRDcp(first *firefly.Decimal, rest ...*firefly.Decimal) *firefly.Decimal {
	var (
		product 		= new(firefly.Decimal)
		restproduct 	= firefly.NFI(1)
	)

	for _, item := range rest {
		_, _ = c.Mul(restproduct,restproduct,item)
	}
	_, _ = c.Mul(product,first,restproduct)

	return product
}
//================================================
//
// Function 06a - POWpr
//
// POWpr multiplies two decimals within a custom Precision modified CryptoplasmPrecisionContext Context
func POWpr(TotalDecimalPrecision uint32, member1, member2 *firefly.Decimal) *firefly.Decimal {
	var result			= new(firefly.Decimal)
	cc := c.WithPrecision(TotalDecimalPrecision)
	_, _ = cc.Pow(result,member1,member2)
	return result
}
//================================================
//
// Function 06b - POWcp
//
// POWcp multiplies two decimals within CryptoplasmPrecisionContext Context
func POWcp(member1, member2 *firefly.Decimal) *firefly.Decimal {
	var result			= new(firefly.Decimal)

	_, _ = c.Pow(result,member1,member2)
	return result
}
//================================================
//
// Function 07a - DIVpr
//
// DIVpr multiplies two decimals within a custom Precision modified CryptoplasmPrecisionContext Context
func DIVpr(TotalDecimalPrecision uint32, member1, member2 *firefly.Decimal) *firefly.Decimal {
	var result			= new(firefly.Decimal)
	cc := c.WithPrecision(TotalDecimalPrecision)
	_, _ = cc.Quo(result,member1,member2)
	return result
}
//================================================
//
// Function 07b - DIVcp
//
// DIVcp multiplies two decimals within CryptoplasmPrecisionContext Context
func DIVcp(member1, member2 *firefly.Decimal) *firefly.Decimal {
	var result			= new(firefly.Decimal)

	_, _ = c.Quo(result,member1,member2)
	return result
}
//================================================
//
// Function 07b - GetDigits
//
// GetDigits returns the number of digits before the coma for a decimal
func GetDigits(Number *firefly.Decimal) int64 {
	Digits := Number.NumDigits()
	return Digits
}
//================================================
//
// Function 08 - Truncate
//
// Truncate truncates the decimal to the precision number
func Truncate(number *firefly.Decimal, DecimalPrecision uint32) *firefly.Decimal {
	var result			= new(firefly.Decimal)
	ConvertedPrecision := int32(DecimalPrecision)
	ConvertedPrecision = 0 - ConvertedPrecision
	_, _ = c.Quantize(result,number,ConvertedPrecision)
	return result
}
//================================================
//
// Function 08 - TruncateCustom
//
// TruncateCustom truncates the decimal to the precision number
func TruncateCustom(TotalDecimalPrecision uint32, number *firefly.Decimal, DecimalPrecision uint32) *firefly.Decimal {
	var result			= new(firefly.Decimal)
	ConvertedPrecision := int32(DecimalPrecision)
	ConvertedPrecision = 0 - ConvertedPrecision
	cc := c.WithPrecision(TotalDecimalPrecision)
	_, _ = cc.Quantize(result,number,ConvertedPrecision)
	return result
}
//================================================
//
// Function 08a - TruncSeed
//
// TruncSeed truncates the decimal to CryptoplasmSeedPrecision
func TruncSeed(number *firefly.Decimal) *firefly.Decimal {
	var result			= new(firefly.Decimal)
	CSP := 0 - int32(CryptoplasmSeedPrecision)
	_, _ = c.Quantize(result,number,CSP)
	return result
}
//================================================
//
// Function 08b - TruncCurrency
//
// TruncCurrency truncates the decimal to CryptoplasmCurrencyPrecision
func TruncCurrency(number *firefly.Decimal) *firefly.Decimal {
	var result			= new(firefly.Decimal)
	CCP := 0 - int32(CryptoplasmCurrencyPrecision)
	_, _ = c.Quantize(result,number,CCP)
	return result
}
//
//================================================
//
// Function 08c - TruncCurrencyCustom
//
// TruncCurrencyCustom truncates the decimal to CryptoplasmCurrencyPrecision
// within a custom length Context. Used when the base CryptoplasmContextPrecision
// isn't enough, used for numbers that dont fit within it.
func TruncCurrencyCustom(TotalDecimalPrecision uint32, number *firefly.Decimal) *firefly.Decimal {
	var result			= new(firefly.Decimal)
	cc := c.WithPrecision(TotalDecimalPrecision)
	CCP := 0 - int32(CryptoplasmCurrencyPrecision)
	_, _ = cc.Quantize(result,number,CCP)
	return result
}
//
//================================================
//
// Function 09 - PrintDL
//
// PrintDL short for PrintDecimalList, prints the decimals
// within the given list/slice
func PrintDL (a []string) {
	for i := 0; i < len(a); i++ {
		fmt.Println("Element is,", a[i])
	}
}
//================================================
//
// Function 03 - SumDL
//
// SumDL short for SumDecimalList, return the sum of
// the decimals within the list/slice
func SumDL (a []*firefly.Decimal) *firefly.Decimal {
	var sum = new(firefly.Decimal)

	for i := 0; i < len(a); i++ {
		sum = ADDcp(sum,a[i])
	}
	return sum
}

//================================================
//
// Function 04 - LastDE
//
// LastDE short for LastDecimalElement, returns the last element
// in the slice. Equivalent to pythons -1 index
func LastDE (a []*firefly.Decimal) *firefly.Decimal {
	Length := len(a)
	LastElementIndex := Length - 1
	LastElement := a[LastElementIndex]
	return LastElement
}
//================================================
//
// Function 05 - AppDec
//
// AppDec creates a new bigger slice from the 2 slices used as input
// slices must be made of decimals
func AppDec (w1, w2 []*firefly.Decimal) []*firefly.Decimal {
	w3 := append(w1,w2...)
	return w3
}
//================================================
//
// Function 06 - Reverse
//
// Returns the Reverse of the Slice/Lists
func Reverse (a []*firefly.Decimal) []*firefly.Decimal {
	var Reversed 		= make([]*firefly.Decimal,0)
	Length := len(a)
	LastElementIndex := Length - 1
	for i := LastElementIndex; i >= 0; i-- {
		Reversed = append(Reversed,a[i])
	}
	return Reversed
}
//================================================
//
// Function 07 - WriteList
//
// WriteList writes the strings from the slice to an external file
// as Name can be used "File.txt" as the output file.
func WriteList(Name string, List[]string) {
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
// Function 07 - RemoveDecimals
//
// RemoveDecimals removes the decimals and leaves the resulted number
// without them
func RemoveDecimals(Number *firefly.Decimal) *firefly.Decimal {
	var Whole			= new(firefly.Decimal)
	NumberDigits := Number.NumDigits()
	cc := c.WithPrecision(uint32(NumberDigits))
	_,_ = cc.Floor(Whole,Number)
	return Whole
}
//================================================
//
// Function 07 - Count4Coma
//
// Count4Coma returns the number of digits before precision
func Count4Coma(Number *firefly.Decimal) int64 {
	Whole := RemoveDecimals(Number)
	Digits := Whole.NumDigits()
	return Digits
}
//================================================
//
// Function 08 - OverspendLog
//
// OverspendLog returns the logarithm base required for computing overspend
func OverspendLogBase(cpAmount *firefly.Decimal) *firefly.Decimal {
	var Base = new(firefly.Decimal)

	DigitsNumber := Count4Coma(cpAmount)
	DND := firefly.NFI(DigitsNumber)	//making decimal the Digit number
	DigitsOperatingPrecision := Count4Coma(DND)

	Exponent := SUBpr(uint32(DigitsOperatingPrecision),DND,firefly.NFI(2))
	e,_ := Exponent.Int64()
	for i := e; i >= 0; i-- {
		idec := firefly.NFI(i)
		Value := MULpr(uint32(DigitsNumber),POWpr(uint32(DigitsNumber),firefly.NFI(10),idec),firefly.NFI(7))
		Base = ADDpr(uint32(DigitsNumber),Base,Value)
	}
	return Base
}
//================================================
//
// Function 09 - Logarithm
//
// Logaritm returns the logarithm from "number" in base "base"
func OVSLogarithm(base,number *firefly.Decimal) *firefly.Decimal {
	var (
		LogBase 		= new(firefly.Decimal)
		LogNumber 		= new(firefly.Decimal)
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
	IP := 2 * CryptoplasmCurrencyPrecision + uint32(NumberDigits)
	cc := c.WithPrecision(IP)
	_,_ = cc.Ln(LogBase,base)
	_,_ = cc.Ln(LogNumber,number)
	CustomLog := DIVpr(IP,LogNumber,LogBase)
	return CustomLog
}
//================================================
//
// Function 10 - CPSend
//
// CPSend computes the AmountFee-FeeProMille, the AmountFee and
// how much the Recipient receives after the AmountFee is deducted from the AmountSent.
func CPSend(cpAmount *firefly.Decimal) (*firefly.Decimal,*firefly.Decimal,*firefly.Decimal) {

	//Setting IP(internal precision) and truncating cpAmount to 24 decimals
	NumberDigits := Count4Coma(cpAmount)
	IP := 2 * CryptoplasmCurrencyPrecision + uint32(NumberDigits)
	//cpAmountTrunc := TruncCurrencyCustom(IP,cpAmount)

	OVSLogBase 	:= OverspendLogBase(cpAmount)
	FeeProMille := TruncCurrencyCustom(IP,OVSLogarithm(OVSLogBase,cpAmount))
	AmountFee 	:= TruncCurrencyCustom(IP,DIVpr(IP,MULpr(IP,cpAmount,FeeProMille),firefly.NFI(1000)))
	Recipient 	:= TruncCurrencyCustom(IP,SUBpr(IP,cpAmount,AmountFee))

	return FeeProMille, AmountFee, Recipient
}
//================================================
//func CPOverSendWithDigits(cpAmount *firefly.Decimal) *firefly.Decimal {
//	Digits := GetDigits(cpAmount)
//	Result := CPOverSend(cpAmount,Digits)
//	return Result
//}
//
// Function 11 - CPOverSend
//
// CPOverSend computes the AmountFee-FeeProMille, the AmountFee and
// how much the Recipient receives after the AmountFee is deducted from the AmountSent.
func CPOverSend(cpAmount *firefly.Decimal) *firefly.Decimal {
	start := time.Now()
	var (
		PerfectOverSend		= new(firefly.Decimal)
		OverSend			= new(firefly.Decimal)
		OSA					= new(firefly.Decimal)	//OverSendArgument
		OSIA				= new(firefly.Decimal)	//OverSendIterationArgument
		FPM					= new(firefly.Decimal)	//AmountFee-ProMille
		AF					= new(firefly.Decimal)	//AmountFee
		R					= new(firefly.Decimal)	//Recipient
		OsiaPrec 			uint32					//Number of Digits after comma OSIA must have to be computed
		Iteration			int						//OverSpend Loop Iteration number
	)

	//Setting IP(internal precision) and truncating cpAmount to 24 decimals
	NumberDigits := Count4Coma(cpAmount)
	IP := CryptoplasmCurrencyPrecision + uint32(NumberDigits) + 1
	cpAmountTrunc := TruncCurrencyCustom(IP,cpAmount)

	//OsiaDivPrec might need more Precision than this, so it gets its own precision
	//OsiaPrec starts with 0, because OSAT is negative
	//OsiaPrec is equal to the positive OSAT, as OSAT increases, so does OsiaPrec
	//OsiaDivPrec starts with 3
	OsiaDivPrec := uint32(3)

	//The Logarithm Base (777...777) that is used for computing OverSend given amount cpAmountTrunc
	OvSLogBase := OverspendLogBase(cpAmountTrunc)


	OSAT := -2										//int type OverSend-Argument-Tier, starts at -2
	//OSAT at -2 needs initially addition precision of 3 (-3 would need 4)
	//Once OSAT is set, or newly set, OverSendArgument (OSA) is computed from it:
	OSA = DIVpr(IP,firefly.NFI(1),POWpr(IP,firefly.NFI(10),firefly.NFI(int64(OSAT))))
	//From OSA the OSIA is derived which is the OverSendArgument being iterated, hence OverSendIterationArgument
	//OSIA is always truncated by the OsiaPrec, thus is grows as needed
	OSIA = TruncateCustom(IP,OSA,OsiaPrec)
	//Above, it is truncated at zero decimals, because it starts without decimals, and gains them while looping



	MaxNoOverSend,_,_ := firefly.NewFromString("9.999999999999999999999999")
	//AFADP is used because in case the cpAmountTrunc is 0.xxx AFMP would be 24 and that wouldn't be enough
	//In this case 25 would be needed and that would be AFAP
	Difference := SUBpr(IP,cpAmountTrunc,MaxNoOverSend)
	IsThreshold := Difference.IsZero()
	if IsThreshold == true || Difference.Negative == true {
		R = cpAmountTrunc
		PerfectOverSend = cpAmountTrunc
		//No extra CP is required as "OverSend" when transacted amount is below 10 CP
	} else if Difference.Negative == false {
		FPMo 	:= TruncCurrencyCustom(IP,OVSLogarithm(OvSLogBase,cpAmountTrunc))								 	//fits in 24+1 context as seen from OVSLogarithm function
		FPM = FPMo
		AFo 	:= TruncCurrencyCustom(IP,DIVpr(IP,MULpr(IP,cpAmountTrunc,FPM),firefly.NFI(1000))) 	//AFADP must be used for Multiplication: 999999.(24decimals) * 1.(24decimals) = could end up (7digits).(24decimals)
		AF 	= AFo
		Ro 		:= TruncCurrencyCustom(IP,SUBpr(IP,cpAmountTrunc,AF))
		R 	= Ro

		Iteration = 1
		OverSend = TruncCurrencyCustom(IP,ADDpr(IP,cpAmountTrunc,MULpr(IP,AFo,ADDpr(IP,firefly.NFI(1),DIVpr(OsiaDivPrec,OSIA,firefly.NFI(1000))))))
		LoopDirection := "up"
		Difference2 := SUBpr(IP,R,cpAmountTrunc)
		IsThreshold2 := Difference2.IsZero()
		for IsThreshold2 == false {
			Iteration = Iteration + 1
			//Troubleshooting Comments can be commented away
			if Iteration == 1000 {
				break
			}
			Difference3 := SUBpr(IP,R,cpAmountTrunc)
			IsThreshold3 := Difference3.IsZero()
			if IsThreshold3 == true {
				//fmt.Println("")
				//fmt.Println("Computing done after",Iteration,"iterations")
				break
			}
			if Difference3.Negative == true {
				if LoopDirection == "down" {
					OSIA = TruncateCustom(OsiaDivPrec,ADDpr(OsiaDivPrec,OSIA,OSA),OsiaPrec)
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
					OSA = DIVpr(OsiaDivPrec,firefly.NFI(1),POWpr(OsiaDivPrec,firefly.NFI(10),firefly.NFI(int64(OSAT))))
				}
				LoopDirection = "up"
				OvSLogBase = OverspendLogBase(OverSend)

				FPM = TruncCurrencyCustom(IP,OVSLogarithm(OvSLogBase,OverSend))
				AF 	= TruncCurrencyCustom(IP,DIVpr(IP,MULpr(IP,OverSend,FPM),firefly.NFI(1000)))
				R 	= TruncCurrencyCustom(IP,SUBpr(IP,OverSend,AF))

				OSIA = TruncateCustom(OsiaDivPrec,ADDpr(OsiaDivPrec,OSIA,OSA),OsiaPrec)
				//Troubleshooting Comments can be commented away
				//fmt.Println("Iteration ",Iteration,"OSIA is",OSIA,"R is", R)
				fmt.Println("Iterating OverSpend computing,..",Iteration)
				OverSend = TruncCurrencyCustom(IP,ADDpr(IP,cpAmountTrunc,MULpr(IP,AFo,ADDpr(IP,firefly.NFI(1),DIVpr(OsiaDivPrec,OSIA,firefly.NFI(1000))))))
			} else if Difference3.Negative == false {
				if LoopDirection == "up" {
					OSIA = TruncateCustom(OsiaDivPrec,SUBpr(OsiaDivPrec,OSIA,OSA),OsiaPrec)
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
					OSA = DIVpr(OsiaDivPrec,firefly.NFI(1),POWpr(OsiaDivPrec,firefly.NFI(10),firefly.NFI(int64(OSAT))))
				}
				LoopDirection = "down"
				OvSLogBase = OverspendLogBase(OverSend)

				FPM = TruncCurrencyCustom(IP,OVSLogarithm(OvSLogBase,OverSend))
				AF 	= TruncCurrencyCustom(IP,DIVpr(IP,MULpr(IP,OverSend,FPM),firefly.NFI(1000)))
				R 	= TruncCurrencyCustom(IP,SUBpr(IP,OverSend,AF))

				OSIA = TruncateCustom(OsiaDivPrec,SUBpr(OsiaDivPrec,OSIA,OSA),OsiaPrec)
				//Troubleshooting Comments can be commented away
				//fmt.Println("Iteration ",Iteration,"OSIA is",OSIA,"R is", R)
				fmt.Println("Iterating OverSpend computing,..",Iteration)
				OverSend = TruncCurrencyCustom(IP,ADDpr(IP,cpAmountTrunc,MULpr(IP,AFo,ADDpr(IP,firefly.NFI(1),DIVpr(OsiaDivPrec,OSIA,firefly.NFI(1000))))))
			}
		}
	}
	//Refining OverSend to PerfectOverSend
	_,_,v3 := CPSend(OverSend)
	//Computes the AU, see bellow, only for observation purposes
	//Must be commented away
	//AtomicUnit := firefly.New(1,-24)
	Difference4 := SUBpr(IP,cpAmountTrunc,v3)
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
			U := DIFpr(IP,v3,cpAmountTrunc)
			//The Number of Atomic Units by which the computed OverSend is off
			//Is only used in the comments for observation purpose
			//AU := DIVpr(IP,U,AtomicUnit)
			PerfectOverSend = DIFpr(IP,OverSend,U)
			//Troubleshooting Comments can be commented away
			//fmt.Println("Computed OverSend was off by ", AU,"AU(s) more")
			//fmt.Println("")
			//fmt.Println("Computed OverSend was      ",OverSend)
			//fmt.Println("Computed PerfectOverSend is",PerfectOverSend)
			//fmt.Println("")
		} else if Difference4.Negative == false {
			U := DIFpr(IP,cpAmountTrunc,v3)
			//The Number of Atomic Units by which the computed OverSend is off
			//Is only used in the comments for observation purpose
			//AU := DIVpr(IP,U,AtomicUnit)
			PerfectOverSend = ADDpr(IP,OverSend,U)
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
	//fmt.Println("OverSend for:", cpAmountTrunc, "CP")
	//fmt.Println("Is equal to:.", PerfectOverSend, "CP")
	//fmt.Println("")
	//fmt.Println("VERIFICATION:")
	//FPMx, AFx, Rx := CPSend(PerfectOverSend)
	//fmt.Println("Sending",PerfectOverSend,"CP, will cost an Amount-Fee of",FPMx,"promille")
	//fmt.Println("The Amount-Fee will therefore amount to",AFx,"CP")
	//fmt.Println("Sending",PerfectOverSend,"CP, the Recipient gets",Rx,"CP")

	elapsed := time.Since(start)
	fmt.Println("Computing OverSpend took", elapsed,"and",Iteration,"Iterations had to be performed")
	return PerfectOverSend
}

