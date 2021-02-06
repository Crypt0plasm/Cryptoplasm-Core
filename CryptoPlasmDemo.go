package main

import (
	b "Cryptoplasm-Core/Packages/Cryptoplasm_Blockchain_Constants"
	p "Cryptoplasm-Core/Packages/Firefly_Precision"
	"flag"
	"fmt"
    	"log"
    	"os"
    	"strings"
    	"time"
)

func main() {
	//start := time.Now()

	var (
	    	CmpDeciSeedH = `--cmp-seed-decimal	NO VALUE TO PASS;
Computes and Prints the Primary and Secondary Decimal Seeds.
The Decimal Seeds have a 40 decimal precision and are used to compute the BlockRewards.
`
	    	CmpTxTaxMinH = `--cmp-tx-tax-min=<cpAmount>
cpAmount is a string of numbers that must be a valid cpAmount. It can have a decimal form.
Computes how much the Recipient receives when the Sender pays the minimum Transaction-Tax
The Tx-Tax is deducted from Recipient's amount. Therefore the Recipient "pays" the Tx-Tax'
`
	    	CmpTxTaxMaxH = `--cmp-tx-tax-max=<cpAmount>
cpAmount is a String of Numbers that must be a valid cpAmount. It can have a decimal form.
Computes how much the Sender has to pay when he/she pays the maximum Transaction Tax
The Tx-Tax is added to the amount to be sent. Therefore the Sender "pays" the Tx-Tax'
`
	    	CmpTxFeeH = `--cmp-tx-fee=<Block-Height>
Block-Height is a String of Numbers that must be a valid Block-Height. It must have a pos. integer form.
Computes and Prints the Transaction-Fee for the given Block-Height. It represents a FeePerByte cpAmount.
`
	    	CmpBrH = `--cmp-br=<Block-Height>
Block-Height is a String of Numbers that must be a valid Block-Height. It must have a pos. integer form.
Computes and Prints the Block-Reward for the given Block-Height. It represent a cpAmount
`
	    	CmpSummedBrH = `--cmp-br-sum=<Block-Height>
Block-Height is a String of Numbers that must be a valid Block-Height. It must have a pos. integer form.
Computes and Prints the Sum of all the Block-Rewards emitted up to and including the given Block-Height.
`
	    	ExpTotalsBrH = `--exp-br-all		NO VALUE TO PASS
Computes and Exports to an external file, all the Block-Rewards for whole White Interval.
The White Interval is composed of 524.596.891 Blocks. Therefore a text file with 524.596.891 lines will
be created. A maximum of 17 GigaBytes are required for the output file. Duration approximately 1 hour.
`
	    	ExpSumBrCkpH = `--exp-br-sum-ckp=<Block-Number>
Block-Number is a String of Numbers representing a number of Blocks. It must have a pos. integer form.
Computes and Exports the Sum of BlockRewards at each Block_Height interval specified by the "Block-Number".
Considering 1 Minute Block-Time, choosing for example 525600 would compute and export the Sum of BRs that
would be emitted in total, in 1 year intervals. Useful if one wants to plot the resulted numbers to create
an "Emission Graph" which displays "Total Emission" in year intervals.
The Computation runs for the Whole White Period - 524.596.891 Blocks, until tail emission.
`
	    	PrtDeciSeedH = `--prt-seed-decimal	NO VALUE TO PASS
Prints the Primary and Secondary Decimal Seeds, without computing them, as they are stored in the code.
The Decimal Seeds have a 40 decimal precision and are used to compute the BlockRewards.
`
	    	PrtIntSeedH = `--prt-seed-integer	NO VALUE TO PASS
Prints the Integer Seeds, as they are stored in the code.
The Integer Seeds, same as the Decimal Seeds, are used to compute the BlockRewards.
`
	    	PrtIntervalH = `--prt-intervals		NO VALUE TO PASS
Prints the Cryptoplasm "Rainbow" Intervals - Purple, Indigo, Blue, Green, Yellow, Orange, Red and White
Displays their lengths in Blocks and their time in years, days, hours and minutes considering 1min/Block
`
	    	SimulateFeeH = `--sml-fee=<Block-Height>
Block-Height is a String of Numbers that must be a valid Block-Height. It must have a pos. integer form.
Computes and Prints all Transaction-Fee types for the given Block-Height. Normal and Blink Tx-Fees,
GAS(>10_CP), PLASMA(>1_CP && <10_CP) and MIASMA (<1_CP) Transaction-Fees.
`
	)

	const (
		CmpDeciSeed = "cmp-seed-decimal"	// Bool
	    	CmpTxTaxMin = "cmp-tx-tax-min"		// String - cpAmount
	    	CmpTxTaxMax = "cmp-tx-tax-max"		// String - cpAmount
	    	CmpTxFee    = "cmp-tx-fee"		// String - BH
	    	CmpBr	    = "cmp-br"			// String - BH
	    	CmpSummedBr = "cmp-br-sum"		// String - BH
	    	ExpTotalsBr = "exp-br-all"		// Bool
	    	ExpSumBrCkp = "exp-br-sum-ckp"		// String - BH Interval
	    	PrtDeciSeed = "prt-seed-decimal"	// Bool
	    	PrtIntSeed  = "prt-seed-integer"	// Bool
	    	PrtInterval = "prt-intervals"		// Bool
	    	SimulateFee = "sml-fee"			// String - BH
	)

	FlagCmpDeciSeed := flag.Bool	(CmpDeciSeed,	false,	CmpDeciSeedH)
    	FlagCmpTxTaxMin := flag.String	(CmpTxTaxMin,	"0", 	CmpTxTaxMinH)
    	FlagCmpTxTaxMax := flag.String	(CmpTxTaxMax,	"0", 	CmpTxTaxMaxH)
	FlagCmpTxFee 	:= flag.String	(CmpTxFee,	"0",	CmpTxFeeH)
	FlagCmpBr	:= flag.String	(CmpBr,		"0",	CmpBrH)
	FlagCmpSummedBr	:= flag.String	(CmpSummedBr,	"0",	CmpSummedBrH)
	FlagExpSumBrCkp := flag.String	(ExpSumBrCkp,	"0",	ExpSumBrCkpH)
    	FlagExpTotalsBr := flag.Bool	(ExpTotalsBr,	false,	ExpTotalsBrH)
	FlagPrtDeciSeed := flag.Bool	(PrtDeciSeed,	false,	PrtDeciSeedH)
    	FlagPrtIntSeed 	:= flag.Bool	(PrtIntSeed,	false,	PrtIntSeedH)
    	FlagPrtInterval := flag.Bool	(PrtInterval,	false,	PrtIntervalH)
    	FlagSimulateFee := flag.String	(SimulateFee,	"0",	SimulateFeeH)

	flag.Parse()

    	//01)CmpDeciSeed	Bool
    	if *FlagCmpDeciSeed == true {
		fmt.Println("")
		fmt.Println("Computing and Printing Secondary and Primary Decimal Seeds")
		fmt.Println("")
		b.CryptoplasmDecimalSeedComputer()
    	}
    	//02)CmpTxTaxMin	String - cpAmount
    	if *FlagCmpTxTaxMin != "0" {
		cpAmount:= p.NFS(*FlagCmpTxTaxMin)
		tcpAmount := b.TruncToCurrency(cpAmount)
		FPM, AF, R := b.CPTxTax(cpAmount)
		fmt.Println("")
		fmt.Println("Sending", b.CPAmountConv2Print(tcpAmount), "CP, will cost an Amount-Fee of", FPM, "promille")
		fmt.Println("The Tx-Tax amounts to", b.CPAmountConv2Print(AF), "CP")
		fmt.Println("You send", b.CPAmountConv2Print(tcpAmount), "CP, the Recipient gets", b.CPAmountConv2Print(R), "CP")
    	}
    	//03)CmpTxTaxMax	String - cpAmount
    	if *FlagCmpTxTaxMax != "0" {
		cpAmount:= p.NFS(*FlagCmpTxTaxMax)
		tcpAmount := b.TruncToCurrency(cpAmount)
		PerfectOverSend := b.CPOverSend(cpAmount)

		FPMx, AFx, Rx := b.CPTxTax(PerfectOverSend)
		x := len(b.CPAmountConv2Print(tcpAmount)) - len(b.CPAmountConv2Print(AFx))
		DisplayOffset := strings.Repeat(" ",x-1)
		fmt.Println("")
		fmt.Println("The Tx-Tax for:",b.CPAmountConv2Print(tcpAmount), "CP")
		fmt.Println("is equal to ...",DisplayOffset,b.CPAmountConv2Print(AFx), "CP")
		fmt.Println("and represents:", FPMx, "promille")
		fmt.Println("")
		fmt.Println("For the Recipient to get", b.CPAmountConv2Print(tcpAmount), "CP")
		fmt.Println("A send of a total of....", b.CPAmountConv2Print(PerfectOverSend), "CP, is required")
		fmt.Println("")
		fmt.Println("VERIFICATION:")
		fmt.Println("The inputted Amount of CP...........:",b.CPAmountConv2Print(tcpAmount), "must be equal to")
		fmt.Println("The amount the Recipient gets.......:",b.CPAmountConv2Print(Rx))
		fmt.Println("Which was computed using the resulted PerfectOverSend.")
		fmt.Println("PerfectOverSend is the big amount of:",b.CPAmountConv2Print(PerfectOverSend))
    	}
    	//04)CmpTxFee		String - BH
    	if *FlagCmpTxFee != "0" {
		start := time.Now()
		Fee := b.FeePerByte(*FlagCmpTxFee)
		fmt.Println("")
		fmt.Println("FeePerByte for BlockHeight", *FlagCmpTxFee, "is:")
		fmt.Println(b.CPAmountConv2Print(Fee), "CP")
		elapsed := time.Since(start)
		fmt.Println("Computing the FeePerByte took", elapsed)
    	}
    	//05)CmpBr		String - BH
    	if *FlagCmpBr != "0" {
		start := time.Now()
		BR := b.BlockRewardS(*FlagCmpBr)
		fmt.Println("")
		fmt.Println("BlockReward for Block-Height", *FlagCmpBr, "is:")
	    	elapsed := time.Since(start)
		fmt.Println(b.CPAmountConv2Print(BR),"CP, computed in",elapsed)
    	}
    	//06)CmpSummedBr	String - BH
    	if *FlagCmpSummedBr != "0" {
	    	fmt.Println("")
	    	Sum := b.BHRewardSeqSumS(*FlagCmpSummedBr)
	    	fmt.Println("")
	    	fmt.Println("Cryptoplasm Emission at Block-Height", *FlagCmpSummedBr,"is:")
	    	fmt.Println(b.CPAmountConv2Print(Sum))
	    	fmt.Println("")
	}
	//07)ExpSumBrCkp	String - BH Interval
    	if *FlagExpSumBrCkp != "0" {
	    var answer, answer2, Name string
	    fmt.Println("")
	    fmt.Println("Are you sure you want to compute and export Block-Rewards Sum")
	    fmt.Println("for 524.596.891 Blocks (998 years) in ",*FlagExpSumBrCkp,"Block Intervals")
	    fmt.Println("Setting the Intervals lower than 525600(1 year), may extend the computation time...")
	    fmt.Println("Answer with yes to continue")
	    _, _ = fmt.Scanln(&answer)
	    if answer == "yes" {
		fmt.Println("Enter a name for the export file, example: SummedIntervals.txt")
		_, _ = fmt.Scanln(&Name)
		fmt.Println("")
		fmt.Println("Do you wish to remove the decimals from the computed sums ?")
		fmt.Println("This enables plotting in Excel (excel doesnt support ver big or very precise numbers..)")
		fmt.Println("Answer with yes to confirm removing decimals")
		_, _ = fmt.Scanln(&answer2)
		SumSlice := b.BHRewardSeqSumCheckpointS(*FlagExpSumBrCkp)
		if answer == "yes" {
		    fmt.Println("")
		    fmt.Println("Removing decimals...")
		    for i := 0; i < len(SumSlice); i++ {
			SumSlice[i] = b.RemoveDecimals(SumSlice[i])
			fmt.Println(SumSlice[i])
		    }
		    fmt.Println("Done removing decimals...")
		}
		//Preparing Output file:
		OutputFile, err := os.Create(Name)
		if err != nil {
		    log.Fatal(err)
		}
		defer OutputFile.Close()
		//Writing to file
		fmt.Println("Writing to file ...")
		for i := 0; i < len(SumSlice); i++ {
		    _, _ = fmt.Fprintln(OutputFile, SumSlice[i])
		}
		fmt.Println("Done Writing to file ...")

	    }
    	}
	//08)ExpTotalsBr	Bool
    	if *FlagExpTotalsBr == true {
	    	var answer, filename string
		fmt.Println("")
		fmt.Println("Are you sure you want to export all 524.596.891 BlockRewards ?")
		fmt.Println("This might take a couple of hours on slower systems....")
		fmt.Println("Answer with yes to continue")
		_, _ = fmt.Scanln(&answer)
		if answer == "yes" {
	    		fmt.Println("Enter a name for the export file, example: BlockRewards.txt")
	    		_, _ = fmt.Scanln(&filename)
	    		b.ExportBr(filename)
		}
    	}
    	//09)PrtDeciSeed	Bool
    	if *FlagPrtDeciSeed == true {
		fmt.Println("")
		fmt.Println("Printing Primary and Secondary Decimal Seeds:")
		fmt.Println("")
		b.CryptoplasmDecimalSeedPrinter()
    	}
    	//10)PrtIntSeed		Bool
	if *FlagPrtIntSeed == true {
	    	fmt.Println("")
		fmt.Println("Printing Integer Seeds:")
	    	fmt.Println("")
		b.CryptoplasmIntegerSeedPrinter()
	}
	//11)PrtInterval	Bool
	if *FlagPrtInterval == true {
	    	fmt.Println("")
		fmt.Println("Printing Intervals:")
	    	fmt.Println("")
		b.CryptoplasmIntervals()
	}
	//12)FlagSimulateFee
	if *FlagSimulateFee != "0" {
		var txsize, outputno uint64
		fmt.Println("")
		fmt.Println("How big is the transaction size in bytes ?")
		_, _ = fmt.Scanln(&txsize)
	    	fmt.Println("")
		fmt.Println("How many outputs are in the transaction ?")
		fmt.Println("A normal transaction has 2 outputs.")
		_, _ = fmt.Scanln(&outputno)
		start := time.Now()
		Fees := b.FeeComputer(*FlagSimulateFee, txsize, outputno)
		fmt.Println("===================================================================")
		fmt.Println("")
		fmt.Println("NORMAL Transactions")
		fmt.Println("")
		fmt.Println("GAS Costs for a Normal Transaction are:")
		fmt.Println(b.CPAmountConv2Print(Fees[0][0][2]), "CP (UTXO based due to outputs number) + ", b.CPAmountConv2Print(Fees[0][0][1]), "CP (due to size in bytes)")
		fmt.Println("TOTAL =", b.CPAmountConv2Print(Fees[0][0][0]))
		fmt.Println("")
		fmt.Println("PLASMA Costs for a Normal Transaction are:")
		fmt.Println(b.CPAmountConv2Print(Fees[0][1][2]), "CP (UTXO based due to outputs number) + ", b.CPAmountConv2Print(Fees[0][1][1]), "CP (due to size in bytes)")
		fmt.Println("TOTAL =", b.CPAmountConv2Print(Fees[0][1][0]))
		fmt.Println("")
		fmt.Println("MIASMA Costs for a Normal Transaction are:")
		fmt.Println(b.CPAmountConv2Print(Fees[0][2][2]), "CP (UTXO based due to outputs number) + ", b.CPAmountConv2Print(Fees[0][2][1]), "CP (due to size in bytes)")
		fmt.Println("TOTAL =", b.CPAmountConv2Print(Fees[0][2][0]))
		fmt.Println("===================================================================")
		fmt.Println("")
		fmt.Println("BLINK Transactions")
		fmt.Println("")
		fmt.Println("GAS Costs for a Blink Transaction are:")
		fmt.Println(b.CPAmountConv2Print(Fees[1][0][2]), "CP (UTXO based due to outputs number) + ", b.CPAmountConv2Print(Fees[1][0][1]), "CP (due to size in bytes)")
		fmt.Println("TOTAL =", b.CPAmountConv2Print(Fees[1][0][0]))
		fmt.Println("")
		fmt.Println("PLASMA Costs for a Blink Transaction are:")
		fmt.Println(b.CPAmountConv2Print(Fees[1][1][2]), "CP (UTXO based due to outputs number) + ", b.CPAmountConv2Print(Fees[1][1][1]), "CP (due to size in bytes)")
		fmt.Println("TOTAL =", b.CPAmountConv2Print(Fees[1][1][0]))
		fmt.Println("")
		fmt.Println("MIASMA Costs for a Blink Transaction are:")
		fmt.Println(b.CPAmountConv2Print(Fees[1][2][2]), "CP (UTXO based due to outputs number) + ", b.CPAmountConv2Print(Fees[1][2][1]), "CP (due to size in bytes)")
		fmt.Println("TOTAL =", b.CPAmountConv2Print(Fees[1][2][0]))
		elapsed := time.Since(start)
		fmt.Println("")
		fmt.Println("Computing all the fees took", elapsed)
	}

	fmt.Println("")
	fmt.Println("Use CryptoPlasmDemo.exe --help")
	fmt.Println("To see all other available options...")
	fmt.Println("===============================================================================================================")
}