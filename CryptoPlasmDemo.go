package main

import (
    "flag"
    "fmt"
    "log"
    "os"
    "strings"
    "time"

    b "Cryptoplasm-Core/Cryptoplasm_Blockchain_Constants"
    el "Cryptoplasm-Core/Cryptoplasm_Elliptic"
    p "github.com/Crypt0plasm/Firefly-APD"
)

func main() {
	//start := time.Now()

	var (
	    	CmpDeciSeedH = `--cmp-seed-decimal	NO VALUE TO PASS;
Computes and Prints the Primary and Secondary Decimal Seeds.
The Decimal Seeds have a 40 decimal precision and are used to compute the BlockRewards.
`
	    	CmpTxTaxH = `--cmp-tx-tax=<cpAmount>
cpAmount is a String of Numbers that must be a valid cpAmount. It can have a decimal form.
Computes the Transaction-Tax; More information by running it.'
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
	    	CmpDiffedBrH = `--cmp-br-diff=<Block-Height><-><Block-Height>
Block-Height is a String of Numbers that must be a valid Block-Height. It must have a pos. integer form.
Computes and Prints the Sum of all the Block-Rewards between the two given Block-Heights. Separated by "-"
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
	    	PrtBrSplitH = `--prt-br-split=<Block-Height>
Block-Height is a String of Numbers that must be a valid Block-Height. It must have a pos. integer form.
Prints the Block-Reward Split, where one can see the Percents going to Miners, Stakers, Conqueror Pool,
Developer Funds, and Stars. These values fluctuate with Block-Height.
`
	    	GenerateKeysH = `--gen-keys		NO VALUE TO PASS
Generates a Pair of Private and Public Keys and its Corresponding Address. The keys are not saved.
They are only generated and displayed on screen.
`
	)

	const (
	    	CmpDeciSeed 	= "cmp-seed-decimal"		// Bool
	    	CmpTxTax 	= "cmp-tx-tax"			// String - cpAmount
	    	CmpTxFee    	= "cmp-tx-fee"			// String - BH
	    	CmpBr	    	= "cmp-br"			// String - BH
	    	CmpSummedBr 	= "cmp-br-sum"			// String - BH
	    	CmpDiffedBr	= "cmp-br-dif"			// String - BH
	    	ExpTotalsBr 	= "exp-br-all"			// Bool
	    	ExpSumBrCkp 	= "exp-br-sum-ckp"		// String - BH Interval
	    	PrtBrSplit	= "prt-br-split"		// String - BH
	    	PrtDeciSeed 	= "prt-seed-decimal"		// Bool
	    	PrtIntSeed  	= "prt-seed-integer"		// Bool
	    	PrtInterval 	= "prt-intervals"		// Bool
	    	SimulateFee 	= "sml-fee"			// String - BH
	    	GenerateKeys	= "gen-keys"			// Bool
	)

	FlagCmpDeciSeed 	:= flag.Bool	(CmpDeciSeed,	false,	CmpDeciSeedH)
	FlagCmpTxTax 		:= flag.String	(CmpTxTax,	"0", 	CmpTxTaxH)
	FlagCmpTxFee 		:= flag.String	(CmpTxFee,	"0",	CmpTxFeeH)
	FlagCmpBr			:= flag.String	(CmpBr,		"0",	CmpBrH)
	FlagCmpSummedBr		:= flag.String	(CmpSummedBr,	"0",	CmpSummedBrH)
	FlagCmpDiffedBr		:= flag.String	(CmpDiffedBr,	"0",	CmpDiffedBrH)
	FlagExpTotalsBr 	:= flag.Bool	(ExpTotalsBr,	false,	ExpTotalsBrH)
	FlagExpSumBrCkp 	:= flag.String	(ExpSumBrCkp,	"0",	ExpSumBrCkpH)
	FlagPrtBrSplit 		:= flag.String	(PrtBrSplit,	"0",	PrtBrSplitH)
	FlagPrtDeciSeed 	:= flag.Bool	(PrtDeciSeed,	false,	PrtDeciSeedH)
	FlagPrtIntSeed 		:= flag.Bool	(PrtIntSeed,	false,	PrtIntSeedH)
	FlagPrtInterval 	:= flag.Bool	(PrtInterval,	false,	PrtIntervalH)
	FlagSimulateFee 	:= flag.String	(SimulateFee,	"0",	SimulateFeeH)
	FlagGenerateKeys	:= flag.Bool	(GenerateKeys,	false,	GenerateKeysH)
	//
	flag.Parse()
	//
	//01)CmpDeciSeed	Bool
	//
	if *FlagCmpDeciSeed == true {
		fmt.Println("")
		fmt.Println("Computing and Printing Secondary and Primary Decimal Seeds")
		fmt.Println("")
		b.CryptoplasmDecimalSeedComputer()
	}
	//
	//02)CmpTxTaxMax	String - cpAmount
	//
	if *FlagCmpTxTax != "0" {
		cpAmount:= p.NFS(*FlagCmpTxTax)

		fmt.Println("")
		b.TxTaxPrinter(cpAmount)
	}
	//
	//03)CmpTxFee		String - BH
	//
	if *FlagCmpTxFee != "0" {
		start := time.Now()
		Fee := b.FeePerByte(*FlagCmpTxFee)
		fmt.Println("")
		fmt.Println("FeePerByte for BlockHeight", b.BHAmountConv2Print(p.NFS(*FlagCmpTxFee)), "is:")
		fmt.Println(b.KosonicDecimalConversion(Fee), "CP")
		elapsed := time.Since(start)
		fmt.Println("Computing the FeePerByte took", elapsed)
	}
	//
	//04)CmpBr		String - BH
	//
	if *FlagCmpBr != "0" {
		start := time.Now()
		BR := b.BlockRewardS(*FlagCmpBr)
		fmt.Println("")
		fmt.Println("BlockReward for Block-Height", *FlagCmpBr, "is:")
	    	elapsed := time.Since(start)
		fmt.Println(b.KosonicDecimalConversion(BR),"CP, computed in",elapsed)
	}
	//
	//05)CmpSummedBr	String - BH
	//
	if *FlagCmpSummedBr != "0" {
		fmt.Println("")
		Sum := b.BHRewardSeqSumS(*FlagCmpSummedBr)
		fmt.Println("")
		fmt.Println("Cryptoplasm Emission at Block-Height", *FlagCmpSummedBr,"is:")
		fmt.Println(b.KosonicDecimalConversion(Sum))
		fmt.Println("")
	}
	//
	//06)CmpDiffedBr	String - BH Interval
	//
	if *FlagCmpDiffedBr != "0" {
	    start := time.Now()
	    fmt.Println("")
	    BlockHeights := strings.SplitN(*FlagCmpDiffedBr,"-",2)
	    S1 := b.DIFxs(p.NFS(BlockHeights[0]), p.NFS("1")).String()
	    Sum1 := b.BHRewardSeqSumS(S1)
	    Sum2 := b.BHRewardSeqSumS(BlockHeights[1])
	    Sum := b.DIFxc(Sum2,Sum1)
	    elapsed := time.Since(start)
	    fmt.Println("")
	    fmt.Println("Cryptoplasm Emission from Block-Height",BlockHeights[0],"to Block-Height",BlockHeights[1],"is:")
	    fmt.Println(b.KosonicDecimalConversion(Sum),"CP, computed in",elapsed)
	    fmt.Println("")
	}

	//
	//07)ExpSumBrCkp	String - BH Interval
	//
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
	//
	//08)FlagPrtBrSplit	String - BH Interval
	//
	if *FlagPrtBrSplit != "0" {
		fmt.Println("")
		BHd := p.NFS(*FlagPrtBrSplit)
		b.ProcentSplitPrinter(BHd)
		fmt.Println("")
	}
	//
	//09)ExpTotalsBr	Bool
	//
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
	//
	//10)PrtDeciSeed	Bool
	//
	if *FlagPrtDeciSeed == true {
		fmt.Println("")
		fmt.Println("Printing Primary and Secondary Decimal Seeds:")
		fmt.Println("")
		b.CryptoplasmDecimalSeedPrinter()
	}
	//
	//11)PrtIntSeed		Bool
	//
	if *FlagPrtIntSeed == true {
		fmt.Println("")
		fmt.Println("Printing Integer Seeds:")
		fmt.Println("")
		b.CryptoplasmIntegerSeedPrinter()
	}
	//
	//12)PrtInterval	Bool
	//
	if *FlagPrtInterval == true {
		fmt.Println("")
		fmt.Println("Printing Intervals:")
		fmt.Println("")
		b.CryptoplasmIntervals()
	}
	//
	//13)FlagSimulateFee
	//
	if *FlagSimulateFee != "0" {
		var txsize, outputno uint32
		fmt.Println("")
		fmt.Println("How big is the transaction size in bytes ?")
		_, _ = fmt.Scanln(&txsize)
	    	fmt.Println("")
		fmt.Println("How many outputs are in the transaction ?")
		fmt.Println("A normal transaction has 2 outputs.")
		_, _ = fmt.Scanln(&outputno)
		start := time.Now()
		Fees := b.FeeComputer(*FlagSimulateFee, txsize, outputno)
		fmt.Println("For BlockHeight",b.BHAmountConv2Print(p.NFS(*FlagSimulateFee)),"following fees apply:")
		fmt.Println("===================================================================")
		fmt.Println("")
		fmt.Println("NORMAL Transactions")
		fmt.Println("")
		fmt.Println("GAS Costs for a Normal Transaction are:")
		fmt.Println(b.KosonicDecimalConversion(Fees[0][0][2]), "CP (UTXO based due to outputs number) + ", b.KosonicDecimalConversion(Fees[0][0][1]), "CP (due to size in bytes)")
		fmt.Println("TOTAL =", b.KosonicDecimalConversion(Fees[0][0][0]))
		fmt.Println("")
		fmt.Println("PLASMA Costs for a Normal Transaction are:")
		fmt.Println(b.KosonicDecimalConversion(Fees[0][1][2]), "CP (UTXO based due to outputs number) + ", b.KosonicDecimalConversion(Fees[0][1][1]), "CP (due to size in bytes)")
		fmt.Println("TOTAL =", b.KosonicDecimalConversion(Fees[0][1][0]))
		fmt.Println("")
		fmt.Println("MIASMA Costs for a Normal Transaction are:")
		fmt.Println(b.KosonicDecimalConversion(Fees[0][2][2]), "CP (UTXO based due to outputs number) + ", b.KosonicDecimalConversion(Fees[0][2][1]), "CP (due to size in bytes)")
		fmt.Println("TOTAL =", b.KosonicDecimalConversion(Fees[0][2][0]))
		fmt.Println("===================================================================")
		fmt.Println("")
		fmt.Println("BLINK Transactions")
		fmt.Println("")
		fmt.Println("GAS Costs for a Blink Transaction are:")
		fmt.Println(b.KosonicDecimalConversion(Fees[1][0][2]), "CP (UTXO based due to outputs number) + ", b.KosonicDecimalConversion(Fees[1][0][1]), "CP (due to size in bytes)")
		fmt.Println("TOTAL =", b.KosonicDecimalConversion(Fees[1][0][0]))
		fmt.Println("")
		fmt.Println("PLASMA Costs for a Blink Transaction are:")
		fmt.Println(b.KosonicDecimalConversion(Fees[1][1][2]), "CP (UTXO based due to outputs number) + ", b.KosonicDecimalConversion(Fees[1][1][1]), "CP (due to size in bytes)")
		fmt.Println("TOTAL =", b.KosonicDecimalConversion(Fees[1][1][0]))
		fmt.Println("")
		fmt.Println("MIASMA Costs for a Blink Transaction are:")
		fmt.Println(b.KosonicDecimalConversion(Fees[1][2][2]), "CP (UTXO based due to outputs number) + ", b.KosonicDecimalConversion(Fees[1][2][1]), "CP (due to size in bytes)")
		fmt.Println("TOTAL =", b.KosonicDecimalConversion(Fees[1][2][0]))
		elapsed := time.Since(start)
		fmt.Println("")
		fmt.Println("Computing all the fees took", elapsed)
	}
	//
	//
	//
    	if *FlagGenerateKeys == true {
	    Keys, Address := el.CryptoplasmCurve.CPFromRandomBits()
	    fmt.Println("Private Key is,",Keys.PrivateKey)
	    fmt.Println("Public  Key is,",Keys.PublicKey)
	    fmt.Println("Address Str is,",Address)
	    fmt.Println("")
	}

	fmt.Println("")
	fmt.Println("Use CryptoPlasmDemo.exe --help")
	fmt.Println("To see all other available options...")
	fmt.Println("===============================================================================================================")
}