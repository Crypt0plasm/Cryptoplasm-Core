package main

import (
	b "Cryptoplasm-Core/Packages/Cryptoplasm_Blockchain_Constants"
	firefly "Cryptoplasm-Core/Packages/Firefly_Precision"
	"flag"
	"fmt"
	"time"
)

func main () {
	//start := time.Now()

	var (
		version = `1.0.0`
		cdsHelp = `Computes and prints (as they are computed) the Primary and Secondary 
Decimal Seeds which are used to compute the BlockReward(s).`
		psdHelp = `Prints directly (without computing) the Primary and Secondary
Decimal Seeds which are used to compute the BlockReward(s).`
		psiHelp = `Prints the Integer Seeds used to compute the BlockReward(s).`
		pinHelp = `Prints CryptoPlasm Intervals. Displays interval sizes and time per interval.`
		pbrHelp = `Prints the Block Reward for a given BlockHeight.
Height must be a positive integer.`
		ebrHelp = `Exports all 524.596.891 BlockRewards`
		cusHelp = `Computes the AmountFee-ProMille and AmountFee for a given CP Amount.
Also computes how much CP the Recipient receives after the AmountFee is deducted.
The Recipient amount represents the "Underspend" amount.
Enter a Cryptoplasm Amount using this format: 123412.2134412412`
		cosHelp = `Computes the "OverSend" Amount one needs to send in order for the Recipient
to receive the inputted Amount. The Amount to be send represent the OverSend Amount
Enter a Cryptoplasm Amount using this format: 3312353.7653423432
Maximum number for which OverSpend must have in total under 49 precision (int+dec)`
		pfeHelp = `Prints the Fee per Byte for the given BlockHeight`
		sfeHelp = `Simulates all Fee types for the given BlockHeight
Transaction size and Output Number are required as well`
	)

	const (
		VER = "version"
		CDS = "compute-decimal-seeds"
		PSD = "print-decimal-seeds"
		PSI = "print-integer-seeds"
		PIN = "print-intervals"
		PBR = "print-br"
		EBR = "export-br"
		CUS = "compute-underspend"
		COS = "compute-overspend"
		PFE = "print-fee"
		SFE = "simulate-fee"
		STX = "simulate-tx"
	)

	flagVer := flag.Bool(VER, false, version)
	flagCDS := flag.Bool(CDS, false, cdsHelp)
	flagPSD := flag.Bool(PSD, false, psdHelp)
	flagPSI := flag.Bool(PSI, false, psiHelp)
	flagPIN := flag.Bool(PIN, false, pinHelp)
	flagPBR := flag.Uint64(PBR,0,pbrHelp)
	flagEBR := flag.Bool(EBR, false, ebrHelp)
	flagCUS := flag.String(CUS, "0", cusHelp)
	flagCOS := flag.String(COS, "0", cosHelp)
	flagPFE := flag.Uint64(PFE,0,pfeHelp)
	flagSFE := flag.Uint64(SFE,0,sfeHelp)


	flag.Parse()

	if *flagVer == true {
		fmt.Println("CryptoPlasm Demo version:", version)
	}
	if *flagCDS == true {
		fmt.Println("Computing and Printing Secondary and Primary Decimal Seeds")
		b.CryptoplasmDecimalSeedComputer()
	}
	if *flagPSD == true {
		fmt.Println("Printing Primary and Secondary Decimal Seeds")
		b.CryptoplasmDecimalSeedPrinter()
	}
	if *flagPSI == true {
		fmt.Println("Printing Integer Seeds")
		b.CryptoplasmIntegerSeedPrinter()
	}
	if *flagPIN == true {
		fmt.Println("Printing Intervals")
		b.CryptoplasmIntervals()
	}
	if *flagEBR == true {
		fmt.Println("Exporting BlockRewards:")
		var answer,filename string
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
	if *flagPBR != 0 {
		start := time.Now()
		BR := b.BlockReward(*flagPBR)
		fmt.Println("")
		fmt.Println("BlockReward for BlockHeight",*flagPBR,"is ",BR,"Cryptoplasm")
		elapsed := time.Since(start)
		fmt.Println("Computing the BR took", elapsed)
	}
	if *flagCUS != "0" {
		Amount,_,_ := firefly.NewFromString(*flagCUS)
		truncAmount := b.TruncCurrency(Amount)
		FPM, AF, R := b.CPSend(Amount)
		fmt.Println("Sending",truncAmount,"CP, will cost an Amount-Fee of",FPM,"promille")
		fmt.Println("The Amount-Fee amounts to",AF,"CP")
		fmt.Println("You send",truncAmount,"CP, the Recipient gets",R,"CP")
	}
	if *flagCOS != "0" {
		Amount,_,_ := firefly.NewFromString(*flagCOS)
		truncAmount := b.TruncCurrency(Amount)
		PerfectOverSend:=b.CPOverSend(truncAmount)

		fmt.Println("")
		fmt.Println("OverSpend for:", truncAmount, "CP")
		fmt.Println("Is equal to:..", PerfectOverSend, "CP")
		fmt.Println("")
		fmt.Println("VERIFICATION:")
		FPMx, AFx, Rx := b.CPSend(PerfectOverSend)
		fmt.Println("Sending",PerfectOverSend,"CP, will cost an Amount-Fee of",FPMx,"promille")
		fmt.Println("The Amount-Fee will therefore amount to",AFx,"CP")
		fmt.Println("Sending",PerfectOverSend,"CP, the Recipient gets",Rx,"CP")
	}
	if *flagPFE != 0 {
		start := time.Now()
		Fee := b.FeePerByte(*flagPFE)
		fmt.Println("")
		fmt.Println("FeePerByte for BlockHeight",*flagPFE,"is:")
		fmt.Println(Fee,"CP")
		elapsed := time.Since(start)
		fmt.Println("Computing the FeePerByte took", elapsed)
	}
	if *flagSFE != 0 {
		var txsize,outputno uint64
		fmt.Println("")
		fmt.Println("How big is the transaction size in bytes ?")
		_, _ = fmt.Scanln(&txsize)
		fmt.Println("How many outputs are in the transaction ?")
		fmt.Println("A normal transaction has 2 outputs.")
		_, _ = fmt.Scanln(&outputno)
		start := time.Now()
		Fees := b.FeeComputer(*flagSFE,txsize,outputno)
		fmt.Println("===================================================================")
		fmt.Println("")
		fmt.Println("NORMAL Transactions")
		fmt.Println("")
		fmt.Println("GAS Costs for a Normal Transaction are:")
		fmt.Println(Fees[0][0][2],"CP (utxo based due to outputs number) + ",Fees[0][0][1],"CP (due to size in bytes)")
		fmt.Println("TOTAL =",Fees[0][0][0])
		fmt.Println("")
		fmt.Println("PLASMA Costs for a Normal Transaction are:")
		fmt.Println(Fees[0][1][2],"CP (utxo based due to outputs number) + ",Fees[0][1][1],"CP (due to size in bytes)")
		fmt.Println("TOTAL =",Fees[0][1][0])
		fmt.Println("")
		fmt.Println("MIASMA Costs for a Normal Transaction are:")
		fmt.Println(Fees[0][2][2],"CP (utxo based due to outputs number) + ",Fees[0][2][1],"CP (due to size in bytes)")
		fmt.Println("TOTAL =",Fees[0][2][0])
		fmt.Println("===================================================================")
		fmt.Println("")
		fmt.Println("BLINK Transactions")
		fmt.Println("")
		fmt.Println("GAS Costs for a Blink Transaction are:")
		fmt.Println(Fees[1][0][2],"CP (utxo based due to outputs number) + ",Fees[1][0][1],"CP (due to size in bytes)")
		fmt.Println("TOTAL =",Fees[1][0][0])
		fmt.Println("")
		fmt.Println("PLASMA Costs for a Blink Transaction are:")
		fmt.Println(Fees[1][1][2],"CP (utxo based due to outputs number) + ",Fees[1][1][1],"CP (due to size in bytes)")
		fmt.Println("TOTAL =",Fees[1][1][0])
		fmt.Println("")
		fmt.Println("MIASMA Costs for a Blink Transaction are:")
		fmt.Println(Fees[1][2][2],"CP (utxo based due to outputs number) + ",Fees[1][2][1],"CP (due to size in bytes)")
		fmt.Println("TOTAL =",Fees[1][2][0])
		elapsed := time.Since(start)
		fmt.Println("")
		fmt.Println("Computing all the fees took", elapsed)
	}

	fmt.Println("")
	fmt.Println("Use CryptoPlasmDemo.exe --help")
	fmt.Println("To see all other available options...")
	fmt.Println("===============================================================================================================")

}