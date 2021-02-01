package main

import (
	b "Cryptoplasm-Core/Packages/Cryptoplasm_Blockchain_Constants"
	firefly "Cryptoplasm-Core/Packages/Firefly_Precision"
	"flag"
	"fmt"
	"time"
)

func main() {
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
		cusHelp = `Computes how much the Recipient receives if the Sender pays,
the "min Transaction Fee", that is, when the fee is subtracted from the transacted amount.
As such it computes the Transaction-Tax for a given amount CP Amount.
Any numeric value with coma(point) can be used. Software truncates it to 24 Cryptoplasm decimal precision`
		cosHelp = `Computes the "Max Transaction Tax"", and how much more one has to send,
for the Recipient to receive the inputted amount. The Amount with Tax is named OverSend.
Any numeric value with coma(point) can be used. Software truncates it to 24 Cryptoplasm decimal precision`
		pfeHelp = `Prints the Fee per Byte for the given BlockHeight`
		sfeHelp = `Simulates all Fee types for the given BlockHeight
Transaction size and Output Number are required as well`
		stxHelp = `Simulates a transaction for the given BlockHeight`
	)

	const (
		VER = "version"
		CDS = "compute-decimal-seeds"
		PSD = "print-decimal-seeds"
		PSI = "print-integer-seeds"
		PIN = "print-intervals"
		PBR = "print-br"
		EBR = "export-br"
		CUS = "compute-min-tx-tax"
		COS = "compute-max-tx-tax"
		PFE = "print-fee"
		SFE = "simulate-fee"
		STX = "simulate-tx"
	)

	flagVer := flag.Bool(VER, false, version)
	flagCDS := flag.Bool(CDS, false, cdsHelp)
	flagPSD := flag.Bool(PSD, false, psdHelp)
	flagPSI := flag.Bool(PSI, false, psiHelp)
	flagPIN := flag.Bool(PIN, false, pinHelp)
	flagPBR := flag.Uint64(PBR, 0, pbrHelp)
	flagEBR := flag.Bool(EBR, false, ebrHelp)
	flagCUS := flag.String(CUS, "0", cusHelp)
	flagCOS := flag.String(COS, "0", cosHelp)
	flagPFE := flag.Uint64(PFE, 0, pfeHelp)
	flagSFE := flag.Uint64(SFE, 0, sfeHelp)
	flagSTX := flag.Uint64(STX, 0, stxHelp)

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
		var answer, filename string
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
		fmt.Println("BlockReward for BlockHeight", *flagPBR, "is ", BR, "Cryptoplasm")
		elapsed := time.Since(start)
		fmt.Println("Computing the BR took", elapsed)
	}
	if *flagCUS != "0" {
		cpAmount, _, _ := firefly.NewFromString(*flagCUS)
		tcpAmount := b.TruncToCurrency(cpAmount)
		FPM, AF, R := b.CPSend(cpAmount)
		fmt.Println("Sending", tcpAmount, "CP, will cost an Amount-Fee of", FPM, "promille")
		fmt.Println("The Amount-Fee amounts to", AF, "CP")
		fmt.Println("You send", tcpAmount, "CP, the Recipient gets", R, "CP")
	}
	if *flagCOS != "0" {
		cpAmount, _, _ := firefly.NewFromString(*flagCOS)
		tcpAmount := b.TruncToCurrency(cpAmount)
		PerfectOverSend := b.CPOverSend(cpAmount)

	    	FPMx, AFx, Rx := b.CPSend(PerfectOverSend)
	    	fmt.Println("")
	    	fmt.Println("The Tx-Tax for:", tcpAmount, "CP")
	    	fmt.Println("is equal to ...", AFx, "CP")
	    	fmt.Println("and represents:", FPMx, "promille")
	    	fmt.Println("")
	    	fmt.Println("For the Recipient to get", tcpAmount, "CP")
	    	fmt.Println("A send of a total of....", PerfectOverSend, "CP, is required")
	    	fmt.Println("")
		fmt.Println("VERIFICATION:")
		fmt.Println("The inputted Amount of CP...........:",tcpAmount, "must be equal to")
	    	fmt.Println("The amount the Recipient gets.......:",Rx)
		fmt.Println("Which was computed using the resulted PerfectOverSend.")
		fmt.Println("PerfectOverSend is the big amount of:",PerfectOverSend)
	}
	if *flagPFE != 0 {
		start := time.Now()
		Fee := b.FeePerByte(*flagPFE)
		fmt.Println("")
		fmt.Println("FeePerByte for BlockHeight", *flagPFE, "is:")
		fmt.Println(Fee, "CP")
		elapsed := time.Since(start)
		fmt.Println("Computing the FeePerByte took", elapsed)
	}
	if *flagSFE != 0 {
		var txsize, outputno uint64
		fmt.Println("")
		fmt.Println("How big is the transaction size in bytes ?")
		_, _ = fmt.Scanln(&txsize)
		fmt.Println("How many outputs are in the transaction ?")
		fmt.Println("A normal transaction has 2 outputs.")
		_, _ = fmt.Scanln(&outputno)
		start := time.Now()
		Fees := b.FeeComputer(*flagSFE, txsize, outputno)
		fmt.Println("===================================================================")
		fmt.Println("")
		fmt.Println("NORMAL Transactions")
		fmt.Println("")
		fmt.Println("GAS Costs for a Normal Transaction are:")
		fmt.Println(Fees[0][0][2], "CP (utxo based due to outputs number) + ", Fees[0][0][1], "CP (due to size in bytes)")
		fmt.Println("TOTAL =", Fees[0][0][0])
		fmt.Println("")
		fmt.Println("PLASMA Costs for a Normal Transaction are:")
		fmt.Println(Fees[0][1][2], "CP (utxo based due to outputs number) + ", Fees[0][1][1], "CP (due to size in bytes)")
		fmt.Println("TOTAL =", Fees[0][1][0])
		fmt.Println("")
		fmt.Println("MIASMA Costs for a Normal Transaction are:")
		fmt.Println(Fees[0][2][2], "CP (utxo based due to outputs number) + ", Fees[0][2][1], "CP (due to size in bytes)")
		fmt.Println("TOTAL =", Fees[0][2][0])
		fmt.Println("===================================================================")
		fmt.Println("")
		fmt.Println("BLINK Transactions")
		fmt.Println("")
		fmt.Println("GAS Costs for a Blink Transaction are:")
		fmt.Println(Fees[1][0][2], "CP (utxo based due to outputs number) + ", Fees[1][0][1], "CP (due to size in bytes)")
		fmt.Println("TOTAL =", Fees[1][0][0])
		fmt.Println("")
		fmt.Println("PLASMA Costs for a Blink Transaction are:")
		fmt.Println(Fees[1][1][2], "CP (utxo based due to outputs number) + ", Fees[1][1][1], "CP (due to size in bytes)")
		fmt.Println("TOTAL =", Fees[1][1][0])
		fmt.Println("")
		fmt.Println("MIASMA Costs for a Blink Transaction are:")
		fmt.Println(Fees[1][2][2], "CP (utxo based due to outputs number) + ", Fees[1][2][1], "CP (due to size in bytes)")
		fmt.Println("TOTAL =", Fees[1][2][0])
		elapsed := time.Since(start)
		fmt.Println("")
		fmt.Println("Computing all the fees took", elapsed)
	}
	if *flagSTX != 0 {
		var txsize, outputno uint64
		var Amount string
		//
		fmt.Println("")
		fmt.Println("How much CP do you want to send ?")
		_, _ = fmt.Scanln(&Amount)
		cpAmount, _, _ := firefly.NewFromString(Amount)
		tcpAmount := b.TruncToCurrency(cpAmount)
		//
		fmt.Println("")
		fmt.Println("How big is the transaction size in bytes ?")
		_, _ = fmt.Scanln(&txsize)
		//
		fmt.Println("")
		fmt.Println("How many outputs are in the transaction ?")
		fmt.Println("A normal transaction has 2 outputs.")
		_, _ = fmt.Scanln(&outputno)
		//
		start := time.Now()
		b.TxSimulator(*flagSTX, txsize, outputno, tcpAmount)
		elapsed := time.Since(start)
		fmt.Println("Displaying Fees took", elapsed)
	}

	fmt.Println("")
	fmt.Println("Use CryptoPlasmDemo.exe --help")
	fmt.Println("To see all other available options...")
	fmt.Println("===============================================================================================================")

}
