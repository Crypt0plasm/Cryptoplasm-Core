package Cryptoplasm_Blockchain_Constants

import (
	firefly "Cryptoplasm-Core/Packages/Firefly_Precision"
	"fmt"
	"math"
)

func FeePerByte(BlockHeight uint64) *firefly.Decimal {
	//uint64 means BlockHeight can go as high as 18.446.744.073.709.551.615
	//This BlockHeight would be reached before the End of the Universe,
	//if we are to consider one Block per minute;
	//Halfway before that would happen, converting it from uint64 to int64 would
	//overflow int64 and would break the code.
	//Also storing the BH in an uint64 would be enough until the End of the Universe.

	//Therefore BlockHeight should be Stored as firefly.Decimal type
	//with Negative = false, Finite Form = Finite, Exponent = 0
	//Such a type would be suffice permanently.

	//However implementing ERAs where one ERA represents 524.596.891 Blocks would mean
	//Using uint64 as BlockHeight Type would suffice until the end of the Universe if
	//after 524.596.891 Blocks, the BlockNumbers resets. ERAs are not implemented yet
	//in this code.

	var FpB = new(firefly.Decimal)
	BH := firefly.NFI(int64(BlockHeight))

	BHDigitsNumber := Count4Coma(BH)     //int64type
	THDigitsNumber := Count4Coma(FpBThr) //int64type
	//Max only works for floats
	//Precision can only be uint32
	Pr := uint32(math.Max(float64(BHDigitsNumber), float64(THDigitsNumber)))

	Difference := SUBpr(Pr, BH, FpBThr)
	if Difference.Negative == false {
		//This is the case is BH is equal or greater than FeePerByteBHThreshold
		//FeePerByteBHThreshold is BH 499.950.000
		//This BH is when the FpBMin will be attained, the minimum FeePerByte.
		FpB = FpBMin
	} else {
		FpBcP := CryptoplasmCurrencyPrecision //FpB computing Precision
		FpB = SUBpr(FpBcP, FpBMax, MULpr(FpBcP, BH, FpBInt))
	}
	return FpB
}

func FeeComputer(BlockHeight, TransactionSize, OutputNumber uint64) [2][3][3]*firefly.Decimal {
	FpB := FeePerByte(BlockHeight)
	ON := firefly.NFI(int64(OutputNumber))
	TS := firefly.NFI(int64(TransactionSize))

	ONd := Count4Coma(ON)
	TSd := Count4Coma(TS)
	Pr := uint32(math.Max(float64(TSd), float64(ONd))) + CryptoplasmCurrencyPrecision + 2

	TransxSizeFee := MULpr(Pr, FpB, TS)
	BaseOutputFee := MULpr(Pr, FpB, ON)

	NormalGasOutputFee := MULpr(Pr, BaseOutputFee, firefly.NFI(int64(FeeArray[0][2])))
	NormalGasTxSizeFee := MULpr(Pr, TransxSizeFee, firefly.NFI(int64(FeeArray[0][0])))
	NormalGasSummedFee := ADDpr(Pr, NormalGasOutputFee, NormalGasTxSizeFee)

	NormalPlsOutputFee := MULpr(Pr, BaseOutputFee, firefly.NFI(int64(FeeArray[1][2])))
	NormalPlsTxSizeFee := MULpr(Pr, TransxSizeFee, firefly.NFI(int64(FeeArray[1][0])))
	NormalPlsSummedFee := ADDpr(Pr, NormalPlsOutputFee, NormalPlsTxSizeFee)

	NormalMsmOutputFee := MULpr(Pr, BaseOutputFee, firefly.NFI(int64(FeeArray[2][2])))
	NormalMsmTxSizeFee := MULpr(Pr, TransxSizeFee, firefly.NFI(int64(FeeArray[2][0])))
	NormalMsmSummedFee := ADDpr(Pr, NormalMsmOutputFee, NormalMsmTxSizeFee)

	BlinkGasOutputFee := MULpr(Pr, BaseOutputFee, firefly.NFI(int64(FeeArray[0][3])))
	BlinkGasTxSizeFee := MULpr(Pr, TransxSizeFee, firefly.NFI(int64(FeeArray[0][1])))
	BlinkGasSummedFee := ADDpr(Pr, BlinkGasOutputFee, BlinkGasTxSizeFee)

	BlinkPlsOutputFee := MULpr(Pr, BaseOutputFee, firefly.NFI(int64(FeeArray[1][3])))
	BlinkPlsTxSizeFee := MULpr(Pr, TransxSizeFee, firefly.NFI(int64(FeeArray[1][1])))
	BlinkPlsSummedFee := ADDpr(Pr, BlinkPlsOutputFee, BlinkPlsTxSizeFee)

	BlinkMsmOutputFee := MULpr(Pr, BaseOutputFee, firefly.NFI(int64(FeeArray[2][3])))
	BlinkMsmTxSizeFee := MULpr(Pr, TransxSizeFee, firefly.NFI(int64(FeeArray[2][1])))
	BlinkMsmSummedFee := ADDpr(Pr, BlinkMsmOutputFee, BlinkMsmTxSizeFee)

	//Position 0 is the sum, 1 is size fee, 2 is the output fee
	NormalGAS := [...]*firefly.Decimal{NormalGasSummedFee, NormalGasTxSizeFee, NormalGasOutputFee}
	NormalPLS := [...]*firefly.Decimal{NormalPlsSummedFee, NormalPlsTxSizeFee, NormalPlsOutputFee}
	NormalMSM := [...]*firefly.Decimal{NormalMsmSummedFee, NormalMsmTxSizeFee, NormalMsmOutputFee}
	NormalFees := [3][3]*firefly.Decimal{NormalGAS, NormalPLS, NormalMSM}

	BlinkGAS := [...]*firefly.Decimal{BlinkGasSummedFee, BlinkGasTxSizeFee, BlinkGasOutputFee}
	BlinkPLS := [...]*firefly.Decimal{BlinkPlsSummedFee, BlinkPlsTxSizeFee, BlinkPlsOutputFee}
	BlinkMSM := [...]*firefly.Decimal{BlinkMsmSummedFee, BlinkMsmTxSizeFee, BlinkMsmOutputFee}
	BlinkFees := [3][3]*firefly.Decimal{BlinkGAS, BlinkPLS, BlinkMSM}

	ComputedFeeArray := [2][3][3]*firefly.Decimal{NormalFees, BlinkFees}
	return ComputedFeeArray
}

func TxSimulator(BlockHeight, TransactionSize, OutputNumber uint64, cpAmount *firefly.Decimal) {
	Ten := firefly.NFI(10)
	One := firefly.NFI(1)

	NumberDigits := Count4Coma(cpAmount)
	IP := CryptoplasmCurrencyPrecision + uint32(NumberDigits) + 1
	tcpAmount := TruncToCurrency(cpAmount)
	Fees := FeeComputer(BlockHeight, TransactionSize, OutputNumber)

	Difference1 := SUBpr(IP, tcpAmount, Ten)
	IsThreshold1 := Difference1.IsZero()
	Difference2 := SUBpr(IP, tcpAmount, One)
	IsThreshold2 := Difference2.IsZero()

	if IsThreshold1 == true || Difference1.Negative == false {
		fmt.Println("Above or equal 10")
		fmt.Println("")
		OverSend := CPOverSend(tcpAmount)
		NormalFee := Fees[0][0][0]
		BlinkFee := Fees[1][0][0]
		fmt.Println("")
		fmt.Println("For the Recipient to receive", tcpAmount, "CP")
		fmt.Println("It will take a send of", OverSend, "CP")
		fmt.Println("And either an extra", NormalFee, "CP as normal mining fee")
		fmt.Println("Or an extra", BlinkFee, "CP as mining fees for instant confirmation")
		fmt.Println("")
	} else if Difference1.Negative == true && (IsThreshold2 == true || Difference2.Negative == false) {
		fmt.Println("Between 1 and 10")
		NormalFee := Fees[0][1][0]
		BlinkFee := Fees[1][1][0]
		fmt.Println("Recipient will receive", tcpAmount, "CP")
		fmt.Println("It will take only an extra", NormalFee, "CP as normal mining fee")
		fmt.Println("Or an extra", BlinkFee, "CP as mining fees for instant confirmation")
	} else {
		fmt.Println("Below 1")
		NormalFee := Fees[0][2][0]
		BlinkFee := Fees[1][2][0]
		fmt.Println("Recipient will receive", tcpAmount, "CP")
		fmt.Println("It will take only an extra", NormalFee, "CP as normal mining fee")
		fmt.Println("Or an extra", BlinkFee, "CP as mining fees for instant confirmation")
	}
}
