package Cryptoplasm_Blockchain_Constants

import (
    "fmt"

    p "github.com/Crypt0plasm/Firefly-APD"
)
//
//	BlockChainFees.go				Blockchain specific p Functions
//================================================================================================
//************************************************************************************************
//================================================================================================
// 	Function List:
//
//	01 Comparison Functions operating on decimal type
//		01  - FeePerByte			Returns the FeePerByte for the given BH
//		02  - FeeComputer			Returns all the possible fees for a given Block Height
//		03  - AmountTier 			Returns the tier for a CryptoPlasm Amount.
//	02 TaxSimulation Functions
//		01  - TxSimulator			Simulates a simple transaction
//================================================================================================
//************************************************************************************************
//================================================================================================
//
// Function 01.01 - FeePerByte
//
// FeePerByte returns the Fee per Byte for the given BlockHeight
func FeePerByte(BlockHeightS string) *p.Decimal {
	var (
		FpB 	= new(p.Decimal)
	    	BH	= p.NFS(BlockHeightS)
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
// Function 01.02 - FeeComputer
//
// FeeComputer returns all the possible fees for a given Block Height
func FeeComputer(BlockHeight string, TransactionSize, OutputNumber uint32) [2][3][3]*p.Decimal {

	FpB := FeePerByte(BlockHeight)
	ON := int64(OutputNumber)
	ONd := p.NFI(ON)

	TS := int64(TransactionSize)
	TSd := p.NFI(TS)

	MaxDigitInt64 := MaxInt64(ON,TS)
	Pr := uint32(MaxDigitInt64) + CryptoplasmCurrencyPrecision + 2

	TransxSizeFee := MULx(Pr, FpB, TSd)
	BaseOutputFee := MULx(Pr, FpB, ONd)

	NormalGasOutputFee := MULx(Pr, BaseOutputFee, p.NFI(int64(FeeArray[0][2])))
	NormalGasTxSizeFee := MULx(Pr, TransxSizeFee, p.NFI(int64(FeeArray[0][0])))
	NormalGasSummedFee := ADDx(Pr, NormalGasOutputFee, NormalGasTxSizeFee)

	NormalPlsOutputFee := MULx(Pr, BaseOutputFee, p.NFI(int64(FeeArray[1][2])))
	NormalPlsTxSizeFee := MULx(Pr, TransxSizeFee, p.NFI(int64(FeeArray[1][0])))
	NormalPlsSummedFee := ADDx(Pr, NormalPlsOutputFee, NormalPlsTxSizeFee)

	NormalMsmOutputFee := MULx(Pr, BaseOutputFee, p.NFI(int64(FeeArray[2][2])))
	NormalMsmTxSizeFee := MULx(Pr, TransxSizeFee, p.NFI(int64(FeeArray[2][0])))
	NormalMsmSummedFee := ADDx(Pr, NormalMsmOutputFee, NormalMsmTxSizeFee)

	BlinkGasOutputFee := MULx(Pr, BaseOutputFee, p.NFI(int64(FeeArray[0][3])))
	BlinkGasTxSizeFee := MULx(Pr, TransxSizeFee, p.NFI(int64(FeeArray[0][1])))
	BlinkGasSummedFee := ADDx(Pr, BlinkGasOutputFee, BlinkGasTxSizeFee)

	BlinkPlsOutputFee := MULx(Pr, BaseOutputFee, p.NFI(int64(FeeArray[1][3])))
	BlinkPlsTxSizeFee := MULx(Pr, TransxSizeFee, p.NFI(int64(FeeArray[1][1])))
	BlinkPlsSummedFee := ADDx(Pr, BlinkPlsOutputFee, BlinkPlsTxSizeFee)

	BlinkMsmOutputFee := MULx(Pr, BaseOutputFee, p.NFI(int64(FeeArray[2][3])))
	BlinkMsmTxSizeFee := MULx(Pr, TransxSizeFee, p.NFI(int64(FeeArray[2][1])))
	BlinkMsmSummedFee := ADDx(Pr, BlinkMsmOutputFee, BlinkMsmTxSizeFee)

	//Position 0 is the sum, 1 is size fee, 2 is the output fee
	NormalGAS := [...]*p.Decimal{NormalGasSummedFee, NormalGasTxSizeFee, NormalGasOutputFee}
	NormalPLS := [...]*p.Decimal{NormalPlsSummedFee, NormalPlsTxSizeFee, NormalPlsOutputFee}
	NormalMSM := [...]*p.Decimal{NormalMsmSummedFee, NormalMsmTxSizeFee, NormalMsmOutputFee}
	NormalFees := [3][3]*p.Decimal{NormalGAS, NormalPLS, NormalMSM}

	BlinkGAS := [...]*p.Decimal{BlinkGasSummedFee, BlinkGasTxSizeFee, BlinkGasOutputFee}
	BlinkPLS := [...]*p.Decimal{BlinkPlsSummedFee, BlinkPlsTxSizeFee, BlinkPlsOutputFee}
	BlinkMSM := [...]*p.Decimal{BlinkMsmSummedFee, BlinkMsmTxSizeFee, BlinkMsmOutputFee}
	BlinkFees := [3][3]*p.Decimal{BlinkGAS, BlinkPLS, BlinkMSM}

	ComputedFeeArray := [2][3][3]*p.Decimal{NormalFees, BlinkFees}
	return ComputedFeeArray
}
//================================================
//
// Function 01.03 - AmountTier
//
// AmountTier returns the tier for a CryptoPlasm Amount.
// If CryptoPlasm Amount >= 10, its tier 3
// If CryptoPlasm Amount >= 1 and <10, its tier 2
// If CryptoPlasm Amount < 1, its tier 1
func AmountTier(cpAmount *p.Decimal) uint8 {
    var Result uint8
    Ten := p.NFI(10)
    One := p.NFI(1)
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
//================================================================================================
//
// Function 02.01 - TxSimulator
//
// TxSimulator simulates a simple transaction
func TxSimulator(BlockHeightS string, TransactionSize, OutputNumber uint32, cpAmount *p.Decimal) {
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