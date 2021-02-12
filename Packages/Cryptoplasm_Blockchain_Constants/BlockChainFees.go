package Cryptoplasm_Blockchain_Constants

import (
    firefly "Cryptoplasm-Core/Packages/Firefly_Precision"
    "fmt"
)
//================================================
//
// Function 01 - FeePerByte
//
// FeePerByte returns the Fee per Byte for the given BlockHeight
func FeePerByte(BlockHeightS string) *firefly.Decimal {
	var (
		FpB 	= new(firefly.Decimal)
	    	BH	= firefly.NFS(BlockHeightS)
	)

	if DecimalGreaterThanOrEqual(BH,FpBThr) == true {
		//This is the case is BH is equal or greater than FeePerByteBHThreshold
		//FeePerByteBHThreshold is BH 499.950.000
		//This BH is when the FpBMin will be attained, the minimum FeePerByte.
		FpB = FpBMin
	} else {
		FpBcP := CryptoplasmCurrencyPrecision + 1 //FpB computing Precision
		FpB = TruncToCurrency(SUBx(FpBcP, FpBMax, MULx(FpBcP, BH, FpBInt)))
	}
	return FpB
}
//================================================
//
// Function 02 - FeeComputer
//
// FeeComputer returns all the possible fees for a given Block Height
func FeeComputer(BlockHeight string, TransactionSize, OutputNumber uint32) [2][3][3]*firefly.Decimal {

	FpB := FeePerByte(BlockHeight)
	ON := int64(OutputNumber)
	ONd := firefly.NFI(ON)

	TS := int64(TransactionSize)
	TSd := firefly.NFI(TS)

	MaxDigitInt64 := MaxInt64(ON,TS)
	Pr := uint32(MaxDigitInt64) + CryptoplasmCurrencyPrecision + 2

	TransxSizeFee := MULx(Pr, FpB, TSd)
	BaseOutputFee := MULx(Pr, FpB, ONd)

	NormalGasOutputFee := MULx(Pr, BaseOutputFee, firefly.NFI(int64(FeeArray[0][2])))
	NormalGasTxSizeFee := MULx(Pr, TransxSizeFee, firefly.NFI(int64(FeeArray[0][0])))
	NormalGasSummedFee := ADDx(Pr, NormalGasOutputFee, NormalGasTxSizeFee)

	NormalPlsOutputFee := MULx(Pr, BaseOutputFee, firefly.NFI(int64(FeeArray[1][2])))
	NormalPlsTxSizeFee := MULx(Pr, TransxSizeFee, firefly.NFI(int64(FeeArray[1][0])))
	NormalPlsSummedFee := ADDx(Pr, NormalPlsOutputFee, NormalPlsTxSizeFee)

	NormalMsmOutputFee := MULx(Pr, BaseOutputFee, firefly.NFI(int64(FeeArray[2][2])))
	NormalMsmTxSizeFee := MULx(Pr, TransxSizeFee, firefly.NFI(int64(FeeArray[2][0])))
	NormalMsmSummedFee := ADDx(Pr, NormalMsmOutputFee, NormalMsmTxSizeFee)

	BlinkGasOutputFee := MULx(Pr, BaseOutputFee, firefly.NFI(int64(FeeArray[0][3])))
	BlinkGasTxSizeFee := MULx(Pr, TransxSizeFee, firefly.NFI(int64(FeeArray[0][1])))
	BlinkGasSummedFee := ADDx(Pr, BlinkGasOutputFee, BlinkGasTxSizeFee)

	BlinkPlsOutputFee := MULx(Pr, BaseOutputFee, firefly.NFI(int64(FeeArray[1][3])))
	BlinkPlsTxSizeFee := MULx(Pr, TransxSizeFee, firefly.NFI(int64(FeeArray[1][1])))
	BlinkPlsSummedFee := ADDx(Pr, BlinkPlsOutputFee, BlinkPlsTxSizeFee)

	BlinkMsmOutputFee := MULx(Pr, BaseOutputFee, firefly.NFI(int64(FeeArray[2][3])))
	BlinkMsmTxSizeFee := MULx(Pr, TransxSizeFee, firefly.NFI(int64(FeeArray[2][1])))
	BlinkMsmSummedFee := ADDx(Pr, BlinkMsmOutputFee, BlinkMsmTxSizeFee)

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
//================================================
//
// Function 03 - TxSimulator
//
// TxSimulator simulates a simple transaction
func TxSimulator(BlockHeightS string, TransactionSize, OutputNumber uint32, cpAmount *firefly.Decimal) {
	tcpAmount := TruncToCurrency(cpAmount)
	Fees := FeeComputer(BlockHeightS, TransactionSize, OutputNumber)
	AmountValue := AmountTier(tcpAmount)

	if AmountValue == 3 {
		fmt.Println("Above or equal 10")
		fmt.Println("")
		OverSend := OverSendV2(tcpAmount)
		NormalFee := Fees[0][0][0]
		BlinkFee := Fees[1][0][0]
		fmt.Println("")
		fmt.Println("For the Recipient to receive", tcpAmount, "CP")
		fmt.Println("It will take a send of", OverSend, "CP")
		fmt.Println("And either an extra", NormalFee, "CP as normal mining fee")
		fmt.Println("Or an extra", BlinkFee, "CP as mining fees for instant confirmation")
		fmt.Println("")
	} else if AmountValue == 2 {
		fmt.Println("Between 1 and 10")
		NormalFee := Fees[0][1][0]
		BlinkFee := Fees[1][1][0]
		fmt.Println("Recipient will receive", tcpAmount, "CP")
		fmt.Println("It will take only an extra", NormalFee, "CP as normal mining fee")
		fmt.Println("Or an extra", BlinkFee, "CP as mining fees for instant confirmation")
	} else if AmountValue == 1 {
		fmt.Println("Below 1")
		NormalFee := Fees[0][2][0]
		BlinkFee := Fees[1][2][0]
		fmt.Println("Recipient will receive", tcpAmount, "CP")
		fmt.Println("It will take only an extra", NormalFee, "CP as normal mining fee")
		fmt.Println("Or an extra", BlinkFee, "CP as mining fees for instant confirmation")
	}
}
//================================================
//
// Function 03 - AmountTier
//
// AmountTier returns the tier for a CryptoPlasm Amount.
// If CryptoPlasm Amount >= 10, its tier 3
// If CryptoPlasm Amount >= 1 and <10, its tier 2
// If CryptoPlasm Amount < 1, its tier 1
func AmountTier(cpAmount *firefly.Decimal) uint8 {
    var Result uint8
    Ten := firefly.NFI(10)
    One := firefly.NFI(1)
    tcpAmount := TruncToCurrency(cpAmount)

    if DecimalGreaterThanOrEqual(tcpAmount,Ten) == true {
	Result = 3
    } else if DecimalGreaterThanOrEqual(tcpAmount,One) == true && DecimalLessThan(tcpAmount,Ten) == true {
	Result = 2
    } else {
	Result = 1
    }

    return Result
}