package Cryptoplasm_Blockchain_Constants

import (
    "fmt"
    "log"
    "os"
    "time"

    p "github.com/Crypt0plasm/Firefly-APD"
)
//
//	BR_Compute.go						Block Reward Compute Functions
//
//================================================
// 	Function List:
//
//	01 Listing Functions used to compute the Geometric Heights
//		01 - CryptoplasmPrimaryGeometricListing		Creates list if geometric heights
//		02 - CryptoplasmSecondaryGeometricListing	Creates list if geometric heights
//	02 Cryptoplasm DNA Function that is used to compute the Block-Rewards
//		01 - CryptoplasmGeometricKamelSequence		Creates the Kamel DNA
//	03 Block-Reward computing Functions
//		01a - BlockRewardS				Computes BR from BH as string
//		01b - BlockRewardD				Computes BR from BH as decimal
//		02  - ConvGH					Computes the BR from the Geometric Height
//		03  - BlockGeometricHeight			Computes the Geometric Height from decimal BH
//	04 Block-Reward Exporting Function used to export externally the Block-Rewards
//		01  - ExportBr					Exports all Block Rewards to a file
//	05 Block-Reward Summing functions
//		01a - BHRewardSeqSumS				Computes BR Sum from string BH sequentially
//		01b - BHRewardSeqSumD				Computes BR Sum from decimal BH sequentially
//		02a - BHRewardSeqSumCheckpointS			Computes BR Sum at specific Checkpoints, Checkpoint string
//		02b - BHRewardSeqSumCheckpointD			Computes BR Sum at specific Checkpoints, Checkpoint decimal
//		03  - BHRewardAdder				The Sequential BR sum computer
//		04a - BHRewardIntSumS				Computes BR Sum from string BH intermittently. Slow !!!
//		04b - BHRewardIntSumD				Computes BR Sum from decimal BH intermittently. Slow !!!
//
//================================================
//
// Slices/Lists Functions
// These Functions are used to compute the BlockReward
// Original implementation was made using Python Lists,
// therefore a similar mechanic will be used in Go as well..
//
// Function 01.01 - CryptoplasmPrimaryGeometricListing
//
// Creates a list of geometric heights using the input parameters
func CryptoplasmPrimaryGeometricListing(a *p.Decimal, b int) []*p.Decimal {
	HeightFront := a
	AreaFront := a
	FirstElement := ADDs(AreaFront, Seed1st)
	Carnatz := make([]*p.Decimal, 1)
	Carnatz[0] = FirstElement

	for i := 2; i < b; i++ {
		HeightFront = ADDs(HeightFront, Seed2nd)
		AreaFront = ADDs(HeightFront, Seed1st)
		Carnatz = append(Carnatz, AreaFront)
	}
	return Carnatz
}

//================================================
//
// Slices/Lists Functions
// These Functions are used to compute the BlockReward
// Original implementation was made using Python Lists,
// therefore a similar mechanic will be used in Go as well..
//
// Function 01.02 - CryptoplasmSecondaryGeometricListing
//
// The Magic that is part of computing the BlockRewards
func CryptoplasmSecondaryGeometricListing(x1 int, x2, x3 *p.Decimal) []*p.Decimal {
	var (
		intc       = x1
		d          = x2
		e          = x3
		Missing    = new(p.Decimal)
		SliceFront = make([]*p.Decimal, 1)
		SliceBack  = make([]*p.Decimal, 1)
		SliceBack2 = make([]*p.Decimal, 1)
	)
	ci := p.NFI(int64(intc))
	SliceFront = CryptoplasmPrimaryGeometricListing(p.NFI(0), intc)
	SliceBack = CryptoplasmPrimaryGeometricListing(d, 638-intc)
	Sum1 := SumDL(SliceFront)
	Sum2 := SumDL(SliceBack)
	if d.Negative == false {
		Missing = DIFs(e, Sum1, Sum2)
	} else {
		SliceBack2 = CryptoplasmPrimaryGeometricListing(p.NFI(0), 638-intc)
		Sum3 := SumDL(SliceBack2)
		Missing = ADDs(MULs(ci, d), DIFs(e, Sum1, Sum3))
	}
	SliceFront = append(SliceFront, Missing)
	RevSlice := Reverse(SliceBack)
	ResultList := AppDec(SliceFront, RevSlice)
	return ResultList
}
//================================================================================================
//************************************************************************************************
//================================================================================================
//
// Func 02.01 - CryptoplasmGeometricKamelSequence
//
// Creates the Geometric Kamel Sequences upon which
// the BlockRewards are computed. Result is a slice
// with 823543 string elements.
func CryptoplasmGeometricKamelSequence() []string {
	var (
		PurpleUpPeak   = []string{"A"}
		PurpleDownPeak = []string{"B"}
		IndigoUpPeak   = []string{"C"}
		IndigoDownPeak = []string{"D"}
		BlueUpPeak     = []string{"E"}
		BlueDownPeak   = []string{"F"}
		GreenUpPeak    = []string{"G"}
		GreenDownPeak  = []string{"H"}
		YellowUpPeak   = []string{"I"}
		YellowDownPeak = []string{"J"}
		OrangeUpPeak   = []string{"K"}
		OrangeDownPeak = []string{"L"}
		RedUpPeak      = []string{"M"}
		RedDownPeak    = []string{"N"}
		WhitePeak      = []string{"O"}
	)

	PurpleUp1 	:= PurpleUpPeak
	PurpleDown1 	:= PurpleDownPeak
	//=======================================================
	PurpleUp2 	:= append(PurpleUp1, PurpleUp1...)
	PurpleUp3 	:= append(PurpleUp2, PurpleUp1...)
	PurpleUp4 	:= append(PurpleUp3, PurpleUp1...)
	PurpleUp5 	:= append(PurpleUp4, PurpleUp1...)
	PurpleUp6 	:= append(PurpleUp5, PurpleUp1...)
	PurpleDown2 	:= append(PurpleDown1, PurpleDown1...)
	PurpleDown3 	:= append(PurpleDown2, PurpleDown1...)
	PurpleDown4 	:= append(PurpleDown3, PurpleDown1...)
	PurpleDown5 	:= append(PurpleDown4, PurpleDown1...)
	PurpleDown6 	:= append(PurpleDown5, PurpleDown1...)

	IndigoUp1 	:= append(PurpleUp6, IndigoUpPeak...)
	IndigoDown1 	:= append(IndigoDownPeak, PurpleDown6...)
	//=======================================================
	IndigoUp2 	:= append(IndigoUp1, IndigoUp1...)
	IndigoUp3 	:= append(IndigoUp2, IndigoUp1...)
	IndigoUp4 	:= append(IndigoUp3, IndigoUp1...)
	IndigoUp5 	:= append(IndigoUp4, IndigoUp1...)
	IndigoUp6 	:= append(IndigoUp5, IndigoUp1...)
	IndigoDown2 	:= append(IndigoDown1, IndigoDown1...)
	IndigoDown3 	:= append(IndigoDown2, IndigoDown1...)
	IndigoDown4 	:= append(IndigoDown3, IndigoDown1...)
	IndigoDown5 	:= append(IndigoDown4, IndigoDown1...)
	IndigoDown6 	:= append(IndigoDown5, IndigoDown1...)

	BlueUp1		:= append(IndigoUp6, PurpleUp4...)
	BlueUp1 	 = append(BlueUp1, BlueUpPeak...)
	BlueUp1 	 = append(BlueUp1, PurpleDown2...)
	BlueDown1	:= append(PurpleUp4, BlueDownPeak...)
	BlueDown1 	 = append(BlueDown1, PurpleDown2...)
	BlueDown1 	 = append(BlueDown1, IndigoDown6...)
	//=======================================================
	BlueUp2 	:= append(BlueUp1, BlueUp1...)
	BlueUp3 	:= append(BlueUp2, BlueUp1...)
	BlueUp4 	:= append(BlueUp3, BlueUp1...)
	BlueUp5 	:= append(BlueUp4, BlueUp1...)
	BlueUp6 	:= append(BlueUp5, BlueUp1...)
	BlueDown2 	:= append(BlueDown1, BlueDown1...)
	BlueDown3 	:= append(BlueDown2, BlueDown1...)
	BlueDown4 	:= append(BlueDown3, BlueDown1...)
	BlueDown5 	:= append(BlueDown4, BlueDown1...)
	BlueDown6 	:= append(BlueDown5, BlueDown1...)

	GreenUp1 	:= append(BlueUp6, IndigoUp3...)
	GreenUp1         = append(GreenUp1, PurpleUp1...)
	GreenUp1         = append(GreenUp1, GreenUpPeak...)
	GreenUp1         = append(GreenUp1, PurpleDown5...)
	GreenUp1         = append(GreenUp1, IndigoDown3...)
	GreenDown1 	:= append(IndigoUp4, PurpleUp2...)
	GreenDown1       = append(GreenDown1, GreenDownPeak...)
	GreenDown1 	 = append(GreenDown1, PurpleDown4...)
	GreenDown1 	 = append(GreenDown1, IndigoDown2...)
	GreenDown1 	 = append(GreenDown1, BlueDown6...)
	//=======================================================
	GreenUp2 	:= append(GreenUp1, GreenUp1...)
	GreenUp3 	:= append(GreenUp2, GreenUp1...)
	GreenUp4 	:= append(GreenUp3, GreenUp1...)
	GreenUp5 	:= append(GreenUp4, GreenUp1...)
	GreenUp6 	:= append(GreenUp5, GreenUp1...)
	GreenDown2 	:= append(GreenDown1, GreenDown1...)
	GreenDown3 	:= append(GreenDown2, GreenDown1...)
	GreenDown4 	:= append(GreenDown3, GreenDown1...)
	GreenDown5 	:= append(GreenDown4, GreenDown1...)
	GreenDown6 	:= append(GreenDown5, GreenDown1...)

	YellowUp1 	:= append(GreenUp6, BlueUp2...)
	YellowUp1 	 = append(YellowUp1, IndigoUp1...)
	YellowUp1 	 = append(YellowUp1, PurpleUp2...)
	YellowUp1 	 = append(YellowUp1, YellowUpPeak...)
	YellowUp1 	 = append(YellowUp1, PurpleDown4...)
	YellowUp1 	 = append(YellowUp1, IndigoDown5...)
	YellowUp1 	 = append(YellowUp1, BlueDown4...)
	YellowDown1 	:= append(BlueUp4, IndigoUp3...)
	YellowDown1 	 = append(YellowDown1, PurpleUp2...)
	YellowDown1 	 = append(YellowDown1, YellowDownPeak...)
	YellowDown1 	 = append(YellowDown1, PurpleDown4...)
	YellowDown1 	 = append(YellowDown1, IndigoDown3...)
	YellowDown1 	 = append(YellowDown1, BlueDown2...)
	YellowDown1 	 = append(YellowDown1, GreenDown6...)
	//=======================================================
	YellowUp2 	:= append(YellowUp1, YellowUp1...)
	YellowUp3 	:= append(YellowUp2, YellowUp1...)
	YellowUp4 	:= append(YellowUp3, YellowUp1...)
	YellowUp5 	:= append(YellowUp4, YellowUp1...)
	YellowUp6 	:= append(YellowUp5, YellowUp1...)
	YellowDown2 	:= append(YellowDown1, YellowDown1...)
	YellowDown3 	:= append(YellowDown2, YellowDown1...)
	YellowDown4 	:= append(YellowDown3, YellowDown1...)
	YellowDown5 	:= append(YellowDown4, YellowDown1...)
	YellowDown6 	:= append(YellowDown5, YellowDown1...)

	OrangeUp1 	:= append(YellowUp6, BlueUp6...)
	OrangeUp1 	 = append(OrangeUp1, IndigoUp2...)
	OrangeUp1 	 = append(OrangeUp1, PurpleUp1...)
	OrangeUp1 	 = append(OrangeUp1, OrangeUpPeak...)
	OrangeUp1 	 = append(OrangeUp1, PurpleDown5...)
	OrangeUp1 	 = append(OrangeUp1, IndigoDown4...)
	OrangeUp1 	 = append(OrangeUp1, GreenDown6...)
	OrangeDown1 	:= append(GreenUp4, BlueUp4...) //Glitch in the Matrix
	OrangeDown1 	 = append(OrangeDown1, BlueUp1...)
	OrangeDown1 	 = append(OrangeDown1, IndigoUp4...)
	OrangeDown1 	 = append(OrangeDown1, PurpleUp3...)
	OrangeDown1 	 = append(OrangeDown1, OrangeDownPeak...)
	OrangeDown1 	 = append(OrangeDown1, PurpleDown3...)
	OrangeDown1 	 = append(OrangeDown1, IndigoDown2...)
	OrangeDown1 	 = append(OrangeDown1, BlueDown1...)
	OrangeDown1 	 = append(OrangeDown1, GreenDown2...)
	OrangeDown1 	 = append(OrangeDown1, YellowDown6...)
	//=======================================================
	OrangeUp2 	:= append(OrangeUp1, OrangeUp1...)
	OrangeUp3 	:= append(OrangeUp2, OrangeUp1...)
	OrangeUp4 	:= append(OrangeUp3, OrangeUp1...)
	OrangeUp5 	:= append(OrangeUp4, OrangeUp1...)
	OrangeDown2 	:= append(OrangeDown1, OrangeDown1...)
	OrangeDown3 	:= append(OrangeDown2, OrangeDown1...)
	OrangeDown4 	:= append(OrangeDown3, OrangeDown1...)
	OrangeDown5 	:= append(OrangeDown4, OrangeDown1...)
	OrangeDown6 	:= append(OrangeDown5, OrangeDown1...)

	RedUp 		:= append(OrangeUp5, YellowUp4...)
	RedUp 		 = append(RedUp, GreenUp4...)
	RedUp 		 = append(RedUp, BlueUp6...)
	RedUp 		 = append(RedUp, IndigoUp2...)
	RedUp 		 = append(RedUp, PurpleUp2...)
	RedUp 		 = append(RedUp, RedUpPeak...)
	RedUp 		 = append(RedUp, PurpleDown4...)
	RedUp 		 = append(RedUp, IndigoDown4...)
	RedUp 		 = append(RedUp, GreenDown2...)
	RedUp 		 = append(RedUp, YellowDown2...)
	RedUp 		 = append(RedUp, OrangeDown1...)
	RedDown1 	:= append(YellowUp5, GreenUp4...)
	RedDown1 	 = append(RedDown1, BlueUp3...)
	RedDown1 	 = append(RedDown1, IndigoUp6...)
	RedDown1 	 = append(RedDown1, PurpleUp2...)
	RedDown1 	 = append(RedDown1, RedDownPeak...)
	RedDown1 	 = append(RedDown1, PurpleDown4...)
	RedDown1 	 = append(RedDown1, BlueDown3...)
	RedDown1 	 = append(RedDown1, GreenDown2...)
	RedDown1 	 = append(RedDown1, YellowDown1...)
	RedDown1 	 = append(RedDown1, OrangeDown6...)

	RedDown2 	:= append(RedDown1, RedDown1...)
	RedDown3 	:= append(RedDown2, RedDown1...)
	RedDown4 	:= append(RedDown3, RedDown1...)
	RedDown5 	:= append(RedDown4, RedDown1...)

	Kamel 		:= append(RedUp, OrangeUp2...)
	Kamel 		 = append(Kamel, YellowUp3...)
	Kamel 		 = append(Kamel, GreenUp1...)
	Kamel 		 = append(Kamel, BlueUp2...)
	Kamel 		 = append(Kamel, IndigoUp1...)
	Kamel 		 = append(Kamel, PurpleUp3...)
	Kamel 		 = append(Kamel, WhitePeak...)
	Kamel 		 = append(Kamel, PurpleDown3...)
	Kamel 		 = append(Kamel, IndigoDown5...)
	Kamel 		 = append(Kamel, BlueDown4...)
	Kamel 		 = append(Kamel, GreenDown5...)
	Kamel 		 = append(Kamel, YellowDown3...)
	Kamel 		 = append(Kamel, OrangeDown4...)
	Kamel 		 = append(Kamel, RedDown5...)

	return Kamel
}
//================================================================================================
//************************************************************************************************
//================================================================================================
//
// Func 03.01a - BlockRewardS
//
// BlockRewardS returns the Block Reward for a given Block-Height
// The Block-Height Type is string.
func BlockRewardS(BlockHeightS string) *p.Decimal {
    //start := time.Now()
    BHd 	:= p.NFS(BlockHeightS)
    BR 		:= BlockRewardD(BHd)
    return BR
    //elapsed := time.Since(start)
    //fmt.Println("Computing took", elapsed)
}
//================================================
//
// Func 03.01b - BlockRewardD
//
// BlockRewardD returns the Block Reward for a given Block Height
func BlockRewardD(BlockHeightD *p.Decimal) *p.Decimal {
	//start := time.Now()
	GH := BlockGeometricHeight(BlockHeightD)
	BR := ConvGH(GH)
	//elapsed := time.Since(start)
	//fmt.Println("Computing took", elapsed)
	return BR
}
//================================================
//
// Func 03.02 - ConvGH
//
// ConvGH returns the Block Reward for a Geometric Height
func ConvGH(GeometricHeight *p.Decimal) *p.Decimal {
	GH := GeometricHeight
	GHDigits := Count4Coma(GH)
	MP := uint32(GHDigits) + CryptoplasmSeedPrecision

	BR := ADDx(MP, StartBRd, MULx(MP, GH, Seed3rd))
	TruncatedBR := TruncToCurrency(BR)

	//Increasing highest Block for the whole BR sum to be exactly
	//TheoreticalEraEmission - PreMinedAmount
	Exception := SUBs(TruncatedBR, Seed5th)
	IsHighestBlock := Exception.IsZero()
	if IsHighestBlock == true {
		TruncatedBR = ADDs(DIFs(TotalEmission, PreMine, Seed4th), TruncatedBR)
	}
	return TruncatedBR
}
//================================================
//
// Func 03.03 - BlockGeometricHeight
//
// BlockGeometricHeight returns the geometric height
// for a given BlockHeight on the Kamel Graph
// Used to calculate the BlockRewards
func BlockGeometricHeight(BlockHeight *p.Decimal) *p.Decimal {
	//Using uint64 as BlockHeight type makes this following code work
	//until BH 18.446.744.073.709.551.615. A higher value will overflow
	//the variable and cause the code to not work
	//Considering 1 minute per block, this BH would be attained in 35.1 trillion years
	//This wouldn't last until the end of the universe.
	//I still have time to modify the code until then.
	//Implementing ERAs will make this code last until the end of the Universe
	//1 Era equals 524.596.891 Blocks. Next Block would be Era2.Block1. Rewards will repeat.
	var (
		GH             	= new(p.Decimal)
		CryptoplasmDNA 	= CryptoplasmGeometricKamelSequence()
		Purples		= DivInt(BlockHeight,Purple)
	    	Rest		= DivMod(BlockHeight,Purple)
	    	RestInt 	= p.INT64(Rest)
	)
	if DecimalEqual(Purples,p.NFI(0)) == true {
		GH = Ax[RestInt-1]
	} else if DecimalNotEqual(Purples,p.NFI(0)) == true && DecimalLessThanOrEqual(Purples,p.NFI(int64(len(CryptoplasmDNA)))) {
		for i := 0; i < int(p.INT64(Purples)); i++ {
			element := CryptoplasmDNA[i]
			if element == "A" {
				GH = ADDs(GH, LastDE(Ax))
			} else if element == "B" {
				GH = ADDs(GH, LastDE(Bx))
			} else if element == "C" {
				GH = ADDs(GH, LastDE(Cx))
			} else if element == "D" {
				GH = ADDs(GH, LastDE(Dx))
			} else if element == "E" {
				GH = ADDs(GH, LastDE(Ex))
			} else if element == "F" {
				GH = ADDs(GH, LastDE(Fx))
			} else if element == "G" {
				GH = ADDs(GH, LastDE(Gx))
			} else if element == "H" {
				GH = ADDs(GH, LastDE(Hx))
			} else if element == "I" {
				GH = ADDs(GH, LastDE(Ix))
			} else if element == "J" {
				GH = ADDs(GH, LastDE(Jx))
			} else if element == "K" {
				GH = ADDs(GH, LastDE(Kx))
			} else if element == "L" {
				GH = ADDs(GH, LastDE(Lx))
			} else if element == "M" {
				GH = ADDs(GH, LastDE(Mx))
			} else if element == "N" {
				GH = ADDs(GH, LastDE(Nx))
			} else {
				GH = ADDs(GH, LastDE(Ox))
			}
		}
		GH = SUBs(GH, MULs(Purples, Seed1st))
		if DecimalEqual(Rest,p.NFI(0)) == true {
			GH = ADDs(GH, Seed1st)
		} else if DecimalLessThan(Purples,p.NFI(int64(len(CryptoplasmDNA)))) {
			element := CryptoplasmDNA[p.INT64(Purples)]
			if element == "A" {
				GH = ADDs(GH, Ax[RestInt-1])
			} else if element == "B" {
				GH = ADDs(GH, Bx[RestInt-1])
			} else if element == "C" {
				GH = ADDs(GH, Cx[RestInt-1])
			} else if element == "D" {
				GH = ADDs(GH, Dx[RestInt-1])
			} else if element == "E" {
				GH = ADDs(GH, Ex[RestInt-1])
			} else if element == "F" {
				GH = ADDs(GH, Fx[RestInt-1])
			} else if element == "G" {
				GH = ADDs(GH, Gx[RestInt-1])
			} else if element == "H" {
				GH = ADDs(GH, Hx[RestInt-1])
			} else if element == "I" {
				GH = ADDs(GH, Ix[RestInt-1])
			} else if element == "J" {
				GH = ADDs(GH, Jx[RestInt-1])
			} else if element == "K" {
				GH = ADDs(GH, Kx[RestInt-1])
			} else if element == "L" {
				GH = ADDs(GH, Lx[RestInt-1])
			} else if element == "M" {
				GH = ADDs(GH, Mx[RestInt-1])
			} else if element == "N" {
				GH = ADDs(GH, Nx[RestInt-1])
			} else {
				GH = ADDs(GH, Ox[RestInt-1])
			}
		}
	} else {
		GH = p.NFI(0)
	}
	//fmt.Println("BlockHeight geometric,", GH)
	//Geometric Height is a SeedNumber, so it has Seed Decimal Precision
	GH = TruncSeed(GH)
	return GH
}
//================================================================================================
//************************************************************************************************
//================================================================================================
//
// Func 04.01 - ExportBr
//
// ExportBR exports the whole 524.596.891 BlockRewards
// to an output file.
// as Name something like "File.txt" can be used.
func ExportBr(Name string) {
	start := time.Now()
	var (
		GH             = new(p.Decimal)
		BR             = new(p.Decimal)
		B              = new(p.Decimal)
		S              = new(p.Decimal)
		CryptoplasmDNA = CryptoplasmGeometricKamelSequence()
	)

	//UsedName := fmt.Sprintf("BlockRewardsGolang_%d.txt", CryptoplasmContextPrecision)
	fmt.Println("")
	fmt.Println("Creating ", Name, "..., exporting started...,")
	fmt.Println("Please be patient as the first Green Interval is computed...")
	fmt.Println("")
	OutputFile, err := os.Create(Name)
	if err != nil {
		log.Fatal(err)
	}
	defer OutputFile.Close()

	for i := 0; i < len(CryptoplasmDNA); i++ {
		if (i+1)%343 == 0 {
			Periods := (i + 1) / 343
			intermediary := time.Since(start)
			BlocksNo := int(p.INT64(Green)) * Periods
			fmt.Println("Written", Periods, "/2041 Green Periods, ", BHAmountConv2Print(p.NFI(int64(BlocksNo))), "Blocks, Elapsed Time ", intermediary)
			//defer timeTrack(time.Now(), "Written")
		}
		element := CryptoplasmDNA[i]
		if element == "A" {
			for j := 0; j < 637; j++ {
				GH = ADDs(B, Ax[j])
				BR = ConvGH(GH)
				_, _ = fmt.Fprintln(OutputFile, BR)
				S = ADDs(S, BR)
				if j == 636 {
					B = SUBs(GH, Seed1st)
				}
			}
		} else if element == "B" {
			for j := 0; j < 637; j++ {
				GH = ADDs(B, Bx[j])
				BR = ConvGH(GH)
				_, _ = fmt.Fprintln(OutputFile, BR)
				S = ADDs(S, BR)
				if j == 636 {
					B = SUBs(GH, Seed1st)
				}
			}
		} else if element == "C" {
			for j := 0; j < 637; j++ {
				GH = ADDs(B, Cx[j])
				BR = ConvGH(GH)
				_, _ = fmt.Fprintln(OutputFile, BR)
				S = ADDs(S, BR)
				if j == 636 {
					B = SUBs(GH, Seed1st)
				}
			}
		} else if element == "D" {
			for j := 0; j < 637; j++ {
				GH = ADDs(B, Dx[j])
				BR = ConvGH(GH)
				_, _ = fmt.Fprintln(OutputFile, BR)
				S = ADDs(S, BR)
				if j == 636 {
					B = SUBs(GH, Seed1st)
				}
			}
		} else if element == "E" {
			for j := 0; j < 637; j++ {
				GH = ADDs(B, Ex[j])
				BR = ConvGH(GH)
				_, _ = fmt.Fprintln(OutputFile, BR)
				S = ADDs(S, BR)
				if j == 636 {
					B = SUBs(GH, Seed1st)
				}
			}
		} else if element == "F" {
			for j := 0; j < 637; j++ {
				GH = ADDs(B, Fx[j])
				BR = ConvGH(GH)
				_, _ = fmt.Fprintln(OutputFile, BR)
				S = ADDs(S, BR)
				if j == 636 {
					B = SUBs(GH, Seed1st)
				}
			}
		} else if element == "G" {
			for j := 0; j < 637; j++ {
				GH = ADDs(B, Gx[j])
				BR = ConvGH(GH)
				_, _ = fmt.Fprintln(OutputFile, BR)
				S = ADDs(S, BR)
				if j == 636 {
					B = SUBs(GH, Seed1st)
				}
			}
		} else if element == "H" {
			for j := 0; j < 637; j++ {
				GH = ADDs(B, Hx[j])
				BR = ConvGH(GH)
				_, _ = fmt.Fprintln(OutputFile, BR)
				S = ADDs(S, BR)
				if j == 636 {
					B = SUBs(GH, Seed1st)
				}
			}
		} else if element == "I" {
			for j := 0; j < 637; j++ {
				GH = ADDs(B, Ix[j])
				BR = ConvGH(GH)
				_, _ = fmt.Fprintln(OutputFile, BR)
				S = ADDs(S, BR)
				if j == 636 {
					B = SUBs(GH, Seed1st)
				}
			}
		} else if element == "J" {
			for j := 0; j < 637; j++ {
				GH = ADDs(B, Jx[j])
				BR = ConvGH(GH)
				_, _ = fmt.Fprintln(OutputFile, BR)
				S = ADDs(S, BR)
				if j == 636 {
					B = SUBs(GH, Seed1st)
				}
			}
		} else if element == "K" {
			for j := 0; j < 637; j++ {
				GH = ADDs(B, Kx[j])
				BR = ConvGH(GH)
				_, _ = fmt.Fprintln(OutputFile, BR)
				S = ADDs(S, BR)
				if j == 636 {
					B = SUBs(GH, Seed1st)
				}
			}
		} else if element == "L" {
			for j := 0; j < 637; j++ {
				GH = ADDs(B, Lx[j])
				BR = ConvGH(GH)
				_, _ = fmt.Fprintln(OutputFile, BR)
				S = ADDs(S, BR)
				if j == 636 {
					B = SUBs(GH, Seed1st)
				}
			}
		} else if element == "M" {
			for j := 0; j < 637; j++ {
				GH = ADDs(B, Mx[j])
				BR = ConvGH(GH)
				_, _ = fmt.Fprintln(OutputFile, BR)
				S = ADDs(S, BR)
				if j == 636 {
					B = SUBs(GH, Seed1st)
				}
			}
		} else if element == "N" {
			for j := 0; j < 637; j++ {
				GH = ADDs(B, Nx[j])
				BR = ConvGH(GH)
				_, _ = fmt.Fprintln(OutputFile, BR)
				S = ADDs(S, BR)
				if j == 636 {
					B = SUBs(GH, Seed1st)
				}
			}
		} else if element == "O" {
			for j := 0; j < 637; j++ {
				GH = ADDs(B, Ox[j])
				BR = ConvGH(GH)
				_, _ = fmt.Fprintln(OutputFile, BR)
				S = ADDs(S, BR)
				if j == 636 {
					B = SUBs(GH, Seed1st)
				}
			}
		}
	}
	fmt.Println("Sum or printed BlockRewards is: ", S)
	fmt.Println("Exporting done !")
	fmt.Println("=========================================================")
	elapsed := time.Since(start)
	fmt.Println("Computing took", elapsed)
}
//================================================================================================
//************************************************************************************************
//================================================================================================
//
// Func 05.01a - BHRewardSeqSumS
//
// BHRewardSeqSumS returns the Block-Reward Sum for the given BlockHeight;
// when BlockHeight as input is a String. Method used is sequentially.
// The recommended method to compute the Reward Sum for a given Block Height
func BHRewardSeqSumS(BlockHeightS string) *p.Decimal {
    //start := time.Now()
    BHd 	:= p.NFS(BlockHeightS)
    CPd		:= BHd
    SumSlice	:= BHRewardAdder(BHd,CPd)
    Sum := SumSlice[0]
    return Sum
    //elapsed := time.Since(start)
    //fmt.Println("Computing took", elapsed)
}
//================================================
//
// Func 05.01b - BHRewardSeqSumD
//
// BHRewardSeqSumD returns the Block-Reward Sum for the given BlockHeight;
// when BlockHeight as input is a Decimal. Method used is sequentially.
// The recommended method to compute the Reward Sum for a given Block Height
func BHRewardSeqSumD(BlockHeightD *p.Decimal) *p.Decimal {
    //start := time.Now()
    BHd 	:= BlockHeightD
    CPd		:= BlockHeightD
    SumSlice	:= BHRewardAdder(BHd,CPd)
    Sum := SumSlice[0]
    return Sum
    //elapsed := time.Since(start)
    //fmt.Println("Computing took", elapsed)
}
//================================================
//
// Func 05.02a - BHRewardSeqSumCheckpointS
//
// BHRewardSeqSumCheckpointS returns multiple BlockReward Sums, computed each
// inputted  BlockHeightCheckpoint - for the whole Emission Period (524.596.891 Blocks);
// when BlockHeightCheckpoint as input is a String. Method used is sequentially.
func BHRewardSeqSumCheckpointS(BlockHeightCheckpointS string) []*p.Decimal {
    //start := time.Now()
    BHd 	:= p.NFI(524596891)
    CPd 	:= p.NFS(BlockHeightCheckpointS)
    SumSlice	:= BHRewardAdder(BHd,CPd)
    // Either Print oder List String list
    return SumSlice
    //elapsed := time.Since(start)
    //fmt.Println("Computing took", elapsed)
}
//================================================
//
// Func 05.02b - BHRewardSeqSumCheckpointD
//
// BHRewardSeqSumCheckpointD returns multiple BlockReward Sums, computed each
// inputted  BlockHeightCheckpoint - for the whole Emission Period (524.596.891 Blocks);
// when BlockHeightCheckpoint as input is a Decimal. Method used is sequentially.
func BHRewardSeqSumCheckpointD(BlockHeightCheckpointD *p.Decimal) []*p.Decimal {
    //start := time.Now()
    BHd 	:= p.NFI(524596891)
    CPd 	:= BlockHeightCheckpointD
    SumSlice	:= BHRewardAdder(BHd,CPd)
    // Either Print oder List String list
    return SumSlice
    //elapsed := time.Since(start)
    //fmt.Println("Computing took", elapsed)
}
//================================================
//
// Func 05.03 - BHRewardAdder
//
// BHRewardAdder returns the Reward Sum for the given BlockHeight
// when BlockHeight as input is of Decimal type. Method used is sequentially.
// The recommended method to compute the Reward Sum for a given Block Height
func BHRewardAdder(BlockHeightD, CheckPointD *p.Decimal) []*p.Decimal{
    start := time.Now()
    var (
	GH             		= new(p.Decimal)
	BR             		= new(p.Decimal)
	B              		= new(p.Decimal)
	S              		= new(p.Decimal)
	PurplesIt		= new(p.Decimal)
	SP			uint32
	Ipr			int64
	CryptoplasmDNA 		= CryptoplasmGeometricKamelSequence()
	GPeriods			= DivInt(BlockHeightD,Green)
	Purples			= DivInt(BlockHeightD,Purple)
	Rest			= DivMod(BlockHeightD,Purple)
	RestInt 		= p.INT64(Rest)
	CurrentBlock		= new(p.Decimal)
	SliceSums		[]*p.Decimal
    )

    BlockHeightDigits := Count4Coma(BlockHeightD)
    BHp := CryptoplasmCurrencyPrecision + uint32(BlockHeightDigits) + 1

    Ipr = Count4Coma(Purples)		//i-precision
    i := p.NFI(0)
    SP = 4 + CryptoplasmCurrencyPrecision

    if DecimalEqual(Purples,p.NFI(823543)) == true {
	PurplesIt = SUBx(BHp,Purples,p.NFI(1))
    } else {
        PurplesIt = Purples
    }

    for DecimalLessThanOrEqual(i,PurplesIt) == true {

	if DecimalEqual(DivMod(ADDx(uint32(Ipr),i,p.NFI(1)),p.NFI(343)),p.NFI(0)) == true {
	    Periods := DIVs(ADDx(uint32(Ipr),i,p.NFI(1)),p.NFI(343))
	    intermediary := time.Since(start)

	    //BlocksNoPr := uint32(Count4Coma(Periods)) + 6
	    BlocksNo := MULs(Green,Periods)

	    fmt.Println("Computed",Periods,"/",GPeriods,"Green Periods,",BHAmountConv2Print(BlocksNo),"Blocks, Elapsed Time ", intermediary)
	    //defer timeTrack(time.Now(), "Written")
	}
	element := CryptoplasmDNA[p.INT64(i)]
	if element == "A" {
	    if DecimalEqual(i,Purples) == true {
		for j := 0; j < int(RestInt); j++ {
		    GH = ADDs(B, Ax[j])
		    BR = ConvGH(GH)

		    S = ADDx(SP,S,BR)
		    CurrentBlock = ADDx(BHp,CurrentBlock,p.NFI(1))
		    if DecimalEqual(DivMod(CurrentBlock,CheckPointD),p.NFI(0)) == true {
			SliceSums = append(SliceSums,S)
		    }
		    SumNumberDigits := Count4Coma(S)
		    SP = CryptoplasmCurrencyPrecision + uint32(SumNumberDigits) + 2

		    if j == 636 {
			B = SUBs(GH, Seed1st)
		    }
		}
	    } else {
		for j := 0; j < 637; j++ {
		    GH = ADDs(B, Ax[j])
		    BR = ConvGH(GH)

		    S = ADDx(SP,S,BR)
		    CurrentBlock = ADDx(BHp,CurrentBlock,p.NFI(1))
		    if DecimalEqual(DivMod(CurrentBlock,CheckPointD),p.NFI(0)) == true {
			SliceSums = append(SliceSums,S)
		    }
		    SumNumberDigits := Count4Coma(S)
		    SP = CryptoplasmCurrencyPrecision + uint32(SumNumberDigits) + 2

		    if j == 636 {
			B = SUBs(GH, Seed1st)
		    }
		}
	    }
	} else if element == "B" {
	    if DecimalEqual(i,Purples) == true {
		for j := 0; j < int(RestInt); j++ {
		    GH = ADDs(B, Bx[j])
		    BR = ConvGH(GH)

		    S = ADDx(SP,S,BR)
		    CurrentBlock = ADDx(BHp,CurrentBlock,p.NFI(1))
		    if DecimalEqual(DivMod(CurrentBlock,CheckPointD),p.NFI(0)) == true {
			SliceSums = append(SliceSums,S)
		    }
		    SumNumberDigits := Count4Coma(S)
		    SP = CryptoplasmCurrencyPrecision + uint32(SumNumberDigits) + 2

		    if j == 636 {
			B = SUBs(GH, Seed1st)
		    }
		}
	    } else {
		for j := 0; j < 637; j++ {
		    GH = ADDs(B, Bx[j])
		    BR = ConvGH(GH)

		    S = ADDx(SP,S,BR)
		    CurrentBlock = ADDx(BHp,CurrentBlock,p.NFI(1))
		    if DecimalEqual(DivMod(CurrentBlock,CheckPointD),p.NFI(0)) == true {
			SliceSums = append(SliceSums,S)
		    }
		    SumNumberDigits := Count4Coma(S)
		    SP = CryptoplasmCurrencyPrecision + uint32(SumNumberDigits) + 2

		    if j == 636 {
			B = SUBs(GH, Seed1st)
		    }
		}
	    }
	} else if element == "C" {
	    if DecimalEqual(i,Purples) == true {
		for j := 0; j < int(RestInt); j++ {
		    GH = ADDs(B, Cx[j])
		    BR = ConvGH(GH)

		    S = ADDx(SP,S,BR)
		    CurrentBlock = ADDx(BHp,CurrentBlock,p.NFI(1))
		    if DecimalEqual(DivMod(CurrentBlock,CheckPointD),p.NFI(0)) == true {
			SliceSums = append(SliceSums,S)
		    }
		    SumNumberDigits := Count4Coma(S)
		    SP = CryptoplasmCurrencyPrecision + uint32(SumNumberDigits) + 2

		    if j == 636 {
			B = SUBs(GH, Seed1st)
		    }
		}
	    } else {
		for j := 0; j < 637; j++ {
		    GH = ADDs(B, Cx[j])
		    BR = ConvGH(GH)

		    S = ADDx(SP,S,BR)
		    CurrentBlock = ADDx(BHp,CurrentBlock,p.NFI(1))
		    if DecimalEqual(DivMod(CurrentBlock,CheckPointD),p.NFI(0)) == true {
			SliceSums = append(SliceSums,S)
		    }
		    SumNumberDigits := Count4Coma(S)
		    SP = CryptoplasmCurrencyPrecision + uint32(SumNumberDigits) + 2

		    if j == 636 {
			B = SUBs(GH, Seed1st)
		    }
		}
	    }
	} else if element == "D" {
	    if DecimalEqual(i,Purples) == true {
		for j := 0; j < int(RestInt); j++ {
		    GH = ADDs(B, Dx[j])
		    BR = ConvGH(GH)

		    S = ADDx(SP,S,BR)
		    CurrentBlock = ADDx(BHp,CurrentBlock,p.NFI(1))
		    if DecimalEqual(DivMod(CurrentBlock,CheckPointD),p.NFI(0)) == true {
			SliceSums = append(SliceSums,S)
		    }
		    SumNumberDigits := Count4Coma(S)
		    SP = CryptoplasmCurrencyPrecision + uint32(SumNumberDigits) + 2

		    if j == 636 {
			B = SUBs(GH, Seed1st)
		    }
		}
	    } else {
		for j := 0; j < 637; j++ {
		    GH = ADDs(B, Dx[j])
		    BR = ConvGH(GH)

		    S = ADDx(SP,S,BR)
		    CurrentBlock = ADDx(BHp,CurrentBlock,p.NFI(1))
		    if DecimalEqual(DivMod(CurrentBlock,CheckPointD),p.NFI(0)) == true {
			SliceSums = append(SliceSums,S)
		    }
		    SumNumberDigits := Count4Coma(S)
		    SP = CryptoplasmCurrencyPrecision + uint32(SumNumberDigits) + 2

		    if j == 636 {
			B = SUBs(GH, Seed1st)
		    }
		}
	    }
	} else if element == "E" {
	    if DecimalEqual(i,Purples) == true {
		for j := 0; j < int(RestInt); j++ {
		    GH = ADDs(B, Ex[j])
		    BR = ConvGH(GH)

		    S = ADDx(SP,S,BR)
		    CurrentBlock = ADDx(BHp,CurrentBlock,p.NFI(1))
		    if DecimalEqual(DivMod(CurrentBlock,CheckPointD),p.NFI(0)) == true {
			SliceSums = append(SliceSums,S)
		    }
		    SumNumberDigits := Count4Coma(S)
		    SP = CryptoplasmCurrencyPrecision + uint32(SumNumberDigits) + 2

		    if j == 636 {
			B = SUBs(GH, Seed1st)
		    }
		}
	    } else {
		for j := 0; j < 637; j++ {
		    GH = ADDs(B, Ex[j])
		    BR = ConvGH(GH)

		    S = ADDx(SP,S,BR)
		    CurrentBlock = ADDx(BHp,CurrentBlock,p.NFI(1))
		    if DecimalEqual(DivMod(CurrentBlock,CheckPointD),p.NFI(0)) == true {
			SliceSums = append(SliceSums,S)
		    }
		    SumNumberDigits := Count4Coma(S)
		    SP = CryptoplasmCurrencyPrecision + uint32(SumNumberDigits) + 2

		    if j == 636 {
			B = SUBs(GH, Seed1st)
		    }
		}
	    }
	} else if element == "F" {
	    if DecimalEqual(i,Purples) == true {
		for j := 0; j < int(RestInt); j++ {
		    GH = ADDs(B, Fx[j])
		    BR = ConvGH(GH)

		    S = ADDx(SP,S,BR)
		    CurrentBlock = ADDx(BHp,CurrentBlock,p.NFI(1))
		    if DecimalEqual(DivMod(CurrentBlock,CheckPointD),p.NFI(0)) == true {
			SliceSums = append(SliceSums,S)
		    }
		    SumNumberDigits := Count4Coma(S)
		    SP = CryptoplasmCurrencyPrecision + uint32(SumNumberDigits) + 2

		    if j == 636 {
			B = SUBs(GH, Seed1st)
		    }
		}
	    } else {
		for j := 0; j < 637; j++ {
		    GH = ADDs(B, Fx[j])
		    BR = ConvGH(GH)

		    S = ADDx(SP,S,BR)
		    CurrentBlock = ADDx(BHp,CurrentBlock,p.NFI(1))
		    if DecimalEqual(DivMod(CurrentBlock,CheckPointD),p.NFI(0)) == true {
			SliceSums = append(SliceSums,S)
		    }
		    SumNumberDigits := Count4Coma(S)
		    SP = CryptoplasmCurrencyPrecision + uint32(SumNumberDigits) + 2

		    if j == 636 {
			B = SUBs(GH, Seed1st)
		    }
		}
	    }
	} else if element == "G" {
	    if DecimalEqual(i,Purples) == true {
		for j := 0; j < int(RestInt); j++ {
		    GH = ADDs(B, Gx[j])
		    BR = ConvGH(GH)

		    S = ADDx(SP,S,BR)
		    CurrentBlock = ADDx(BHp,CurrentBlock,p.NFI(1))
		    if DecimalEqual(DivMod(CurrentBlock,CheckPointD),p.NFI(0)) == true {
			SliceSums = append(SliceSums,S)
		    }
		    SumNumberDigits := Count4Coma(S)
		    SP = CryptoplasmCurrencyPrecision + uint32(SumNumberDigits) + 2

		    if j == 636 {
			B = SUBs(GH, Seed1st)
		    }
		}
	    } else {
		for j := 0; j < 637; j++ {
		    GH = ADDs(B, Gx[j])
		    BR = ConvGH(GH)

		    S = ADDx(SP,S,BR)
		    CurrentBlock = ADDx(BHp,CurrentBlock,p.NFI(1))
		    if DecimalEqual(DivMod(CurrentBlock,CheckPointD),p.NFI(0)) == true {
			SliceSums = append(SliceSums,S)
		    }
		    SumNumberDigits := Count4Coma(S)
		    SP = CryptoplasmCurrencyPrecision + uint32(SumNumberDigits) + 2

		    if j == 636 {
			B = SUBs(GH, Seed1st)
		    }
		}
	    }
	} else if element == "H" {
	    if DecimalEqual(i,Purples) == true {
		for j := 0; j < int(RestInt); j++ {
		    GH = ADDs(B, Hx[j])
		    BR = ConvGH(GH)

		    S = ADDx(SP,S,BR)
		    CurrentBlock = ADDx(BHp,CurrentBlock,p.NFI(1))
		    if DecimalEqual(DivMod(CurrentBlock,CheckPointD),p.NFI(0)) == true {
			SliceSums = append(SliceSums,S)
		    }
		    SumNumberDigits := Count4Coma(S)
		    SP = CryptoplasmCurrencyPrecision + uint32(SumNumberDigits) + 2

		    if j == 636 {
			B = SUBs(GH, Seed1st)
		    }
		}
	    } else {
		for j := 0; j < 637; j++ {
		    GH = ADDs(B, Hx[j])
		    BR = ConvGH(GH)

		    S = ADDx(SP,S,BR)
		    CurrentBlock = ADDx(BHp,CurrentBlock,p.NFI(1))
		    if DecimalEqual(DivMod(CurrentBlock,CheckPointD),p.NFI(0)) == true {
			SliceSums = append(SliceSums,S)
		    }
		    SumNumberDigits := Count4Coma(S)
		    SP = CryptoplasmCurrencyPrecision + uint32(SumNumberDigits) + 2

		    if j == 636 {
			B = SUBs(GH, Seed1st)
		    }
		}
	    }
	} else if element == "I" {
	    if DecimalEqual(i,Purples) == true {
		for j := 0; j < int(RestInt); j++ {
		    GH = ADDs(B, Ix[j])
		    BR = ConvGH(GH)

		    S = ADDx(SP,S,BR)
		    CurrentBlock = ADDx(BHp,CurrentBlock,p.NFI(1))
		    if DecimalEqual(DivMod(CurrentBlock,CheckPointD),p.NFI(0)) == true {
			SliceSums = append(SliceSums,S)
		    }
		    SumNumberDigits := Count4Coma(S)
		    SP = CryptoplasmCurrencyPrecision + uint32(SumNumberDigits) + 2

		    if j == 636 {
			B = SUBs(GH, Seed1st)
		    }
		}
	    } else {
		for j := 0; j < 637; j++ {
		    GH = ADDs(B, Ix[j])
		    BR = ConvGH(GH)

		    S = ADDx(SP,S,BR)
		    CurrentBlock = ADDx(BHp,CurrentBlock,p.NFI(1))
		    if DecimalEqual(DivMod(CurrentBlock,CheckPointD),p.NFI(0)) == true {
			SliceSums = append(SliceSums,S)
		    }
		    SumNumberDigits := Count4Coma(S)
		    SP = CryptoplasmCurrencyPrecision + uint32(SumNumberDigits) + 2

		    if j == 636 {
			B = SUBs(GH, Seed1st)
		    }
		}
	    }
	} else if element == "J" {
	    if DecimalEqual(i,Purples) == true {
		for j := 0; j < int(RestInt); j++ {
		    GH = ADDs(B, Jx[j])
		    BR = ConvGH(GH)

		    S = ADDx(SP,S,BR)
		    CurrentBlock = ADDx(BHp,CurrentBlock,p.NFI(1))
		    if DecimalEqual(DivMod(CurrentBlock,CheckPointD),p.NFI(0)) == true {
			SliceSums = append(SliceSums,S)
		    }
		    SumNumberDigits := Count4Coma(S)
		    SP = CryptoplasmCurrencyPrecision + uint32(SumNumberDigits) + 2

		    if j == 636 {
			B = SUBs(GH, Seed1st)
		    }
		}
	    } else {
		for j := 0; j < 637; j++ {
		    GH = ADDs(B, Jx[j])
		    BR = ConvGH(GH)

		    S = ADDx(SP,S,BR)
		    CurrentBlock = ADDx(BHp,CurrentBlock,p.NFI(1))
		    if DecimalEqual(DivMod(CurrentBlock,CheckPointD),p.NFI(0)) == true {
			SliceSums = append(SliceSums,S)
		    }
		    SumNumberDigits := Count4Coma(S)
		    SP = CryptoplasmCurrencyPrecision + uint32(SumNumberDigits) + 2

		    if j == 636 {
			B = SUBs(GH, Seed1st)
		    }
		}
	    }
	} else if element == "K" {
	    if DecimalEqual(i,Purples) == true {
		for j := 0; j < int(RestInt); j++ {
		    GH = ADDs(B, Kx[j])
		    BR = ConvGH(GH)

		    S = ADDx(SP,S,BR)
		    CurrentBlock = ADDx(BHp,CurrentBlock,p.NFI(1))
		    if DecimalEqual(DivMod(CurrentBlock,CheckPointD),p.NFI(0)) == true {
			SliceSums = append(SliceSums,S)
		    }
		    SumNumberDigits := Count4Coma(S)
		    SP = CryptoplasmCurrencyPrecision + uint32(SumNumberDigits) + 2

		    if j == 636 {
			B = SUBs(GH, Seed1st)
		    }
		}
	    } else {
		for j := 0; j < 637; j++ {
		    GH = ADDs(B, Kx[j])
		    BR = ConvGH(GH)

		    S = ADDx(SP,S,BR)
		    CurrentBlock = ADDx(BHp,CurrentBlock,p.NFI(1))
		    if DecimalEqual(DivMod(CurrentBlock,CheckPointD),p.NFI(0)) == true {
			SliceSums = append(SliceSums,S)
		    }
		    SumNumberDigits := Count4Coma(S)
		    SP = CryptoplasmCurrencyPrecision + uint32(SumNumberDigits) + 2

		    if j == 636 {
			B = SUBs(GH, Seed1st)
		    }
		}
	    }
	} else if element == "L" {
	    if DecimalEqual(i,Purples) == true {
		for j := 0; j < int(RestInt); j++ {
		    GH = ADDs(B, Lx[j])
		    BR = ConvGH(GH)

		    S = ADDx(SP,S,BR)
		    CurrentBlock = ADDx(BHp,CurrentBlock,p.NFI(1))
		    if DecimalEqual(DivMod(CurrentBlock,CheckPointD),p.NFI(0)) == true {
			SliceSums = append(SliceSums,S)
		    }
		    SumNumberDigits := Count4Coma(S)
		    SP = CryptoplasmCurrencyPrecision + uint32(SumNumberDigits) + 2

		    if j == 636 {
			B = SUBs(GH, Seed1st)
		    }
		}
	    } else {
		for j := 0; j < 637; j++ {
		    GH = ADDs(B, Lx[j])
		    BR = ConvGH(GH)

		    S = ADDx(SP,S,BR)
		    CurrentBlock = ADDx(BHp,CurrentBlock,p.NFI(1))
		    if DecimalEqual(DivMod(CurrentBlock,CheckPointD),p.NFI(0)) == true {
			SliceSums = append(SliceSums,S)
		    }
		    SumNumberDigits := Count4Coma(S)
		    SP = CryptoplasmCurrencyPrecision + uint32(SumNumberDigits) + 2

		    if j == 636 {
			B = SUBs(GH, Seed1st)
		    }
		}
	    }
	} else if element == "M" {
	    if DecimalEqual(i,Purples) == true {
		for j := 0; j < int(RestInt); j++ {
		    GH = ADDs(B, Mx[j])
		    BR = ConvGH(GH)

		    S = ADDx(SP,S,BR)
		    CurrentBlock = ADDx(BHp,CurrentBlock,p.NFI(1))
		    if DecimalEqual(DivMod(CurrentBlock,CheckPointD),p.NFI(0)) == true {
			SliceSums = append(SliceSums,S)
		    }
		    SumNumberDigits := Count4Coma(S)
		    SP = CryptoplasmCurrencyPrecision + uint32(SumNumberDigits) + 2

		    if j == 636 {
			B = SUBs(GH, Seed1st)
		    }
		}
	    } else {
		for j := 0; j < 637; j++ {
		    GH = ADDs(B, Mx[j])
		    BR = ConvGH(GH)

		    S = ADDx(SP,S,BR)
		    CurrentBlock = ADDx(BHp,CurrentBlock,p.NFI(1))
		    if DecimalEqual(DivMod(CurrentBlock,CheckPointD),p.NFI(0)) == true {
			SliceSums = append(SliceSums,S)
		    }
		    SumNumberDigits := Count4Coma(S)
		    SP = CryptoplasmCurrencyPrecision + uint32(SumNumberDigits) + 2

		    if j == 636 {
			B = SUBs(GH, Seed1st)
		    }
		}
	    }
	} else if element == "N" {
	    if DecimalEqual(i,Purples) == true {
		for j := 0; j < int(RestInt); j++ {
		    GH = ADDs(B, Nx[j])
		    BR = ConvGH(GH)

		    S = ADDx(SP,S,BR)
		    CurrentBlock = ADDx(BHp,CurrentBlock,p.NFI(1))
		    if DecimalEqual(DivMod(CurrentBlock,CheckPointD),p.NFI(0)) == true {
			SliceSums = append(SliceSums,S)
		    }
		    SumNumberDigits := Count4Coma(S)
		    SP = CryptoplasmCurrencyPrecision + uint32(SumNumberDigits) + 2

		    if j == 636 {
			B = SUBs(GH, Seed1st)
		    }
		}
	    } else {
		for j := 0; j < 637; j++ {
		    GH = ADDs(B, Nx[j])
		    BR = ConvGH(GH)

		    S = ADDx(SP,S,BR)
		    CurrentBlock = ADDx(BHp,CurrentBlock,p.NFI(1))
		    if DecimalEqual(DivMod(CurrentBlock,CheckPointD),p.NFI(0)) == true {
			SliceSums = append(SliceSums,S)
		    }
		    SumNumberDigits := Count4Coma(S)
		    SP = CryptoplasmCurrencyPrecision + uint32(SumNumberDigits) + 2

		    if j == 636 {
			B = SUBs(GH, Seed1st)
		    }
		}
	    }
	} else if element == "O" {
	    if DecimalEqual(i,Purples) == true {
		for j := 0; j < int(RestInt); j++ {
		    GH = ADDs(B, Ox[j])
		    BR = ConvGH(GH)

		    S = ADDx(SP,S,BR)
		    CurrentBlock = ADDx(BHp,CurrentBlock,p.NFI(1))
		    if DecimalEqual(DivMod(CurrentBlock,CheckPointD),p.NFI(0)) == true {
			SliceSums = append(SliceSums,S)
		    }
		    SumNumberDigits := Count4Coma(S)
		    SP = CryptoplasmCurrencyPrecision + uint32(SumNumberDigits) + 2

		    if j == 636 {
			B = SUBs(GH, Seed1st)
		    }
		}
	    } else {
		for j := 0; j < 637; j++ {
		    GH = ADDs(B, Ox[j])
		    BR = ConvGH(GH)

		    S = ADDx(SP,S,BR)
		    CurrentBlock = ADDx(BHp,CurrentBlock,p.NFI(1))
		    if DecimalEqual(DivMod(CurrentBlock,CheckPointD),p.NFI(0)) == true {
			SliceSums = append(SliceSums,S)
		    }
		    SumNumberDigits := Count4Coma(S)
		    SP = CryptoplasmCurrencyPrecision + uint32(SumNumberDigits) + 2

		    if j == 636 {
			B = SUBs(GH, Seed1st)
		    }
		}
	    }
	}

	//Incrementing i
	i = ADDx(uint32(Ipr),i,p.NFI(1))
    }
    SliceSums = append(SliceSums,S)
    //elapsed := time.Since(start)
    //fmt.Println("Computing took", elapsed)
    //fmt.Println("Computing Sum sequentially for BH",BlockHeightD,"took", elapsed, "and is ",S)
    //if DecimalNotEqual(BlockHeightD,CheckPointD) == true && DecimalLessThan(CheckPointD,BlockHeightD) == true {
    //	fmt.Println("Block-Reward-Sums for every",CheckPointD, "Blocks have been written to ",Name)
    //} else {
    //    fmt.Println("No CheckPoint Export file has been created")
    //}
    return SliceSums
}
//================================================
//
// Func 05.04a - BHRewardIntSumS
//
// BHRewardIntSumS returns the Reward Sum for the given BlockHeight
// when BlockHeight as input is a String. Method used is intermittently.
// This method isn't recommended to compute the Reward Sum for a given Block Height
// As it is several orders of magnitude slower than the sequential method
// The logic behind this method can be used when a single Block-Reward must be computed.
func BHRewardIntSumS(BlockHeightS string) *p.Decimal {
    //start := time.Now()
    BHd 	:= p.NFS(BlockHeightS)
    Sum		:= BHRewardIntSumD(BHd)
    return Sum
    //elapsed := time.Since(start)
    //fmt.Println("Computing took", elapsed)
}
//================================================
//
// Func 05.04b - BHRewardIntSumD
//
// BHRewardIntSumD returns the Reward Sum for the given BlockHeight
// when BlockHeight as input is a Decimal. Method used is intermittently.
// This method isn't recommended to compute the Reward Sum for a given Block Height
// As it is several orders of magnitude slower than the sequential method
// The logic behind this method can be used when a single Block-Reward must be computed.
func BHRewardIntSumD(BlockHeightD *p.Decimal) *p.Decimal {
    start := time.Now()
    var BrSum = new(p.Decimal)

    BHDigits := Count4Coma(BlockHeightD)	//Count4Coma can be used because BlockHeightD has no decimals
    i := p.NFI(1)
    BrSum = BlockRewardD(i)

    SumNumberDigits := Count4Coma(BrSum)
    SP := CryptoplasmCurrencyPrecision + uint32(SumNumberDigits) + 1
    i = ADDx(uint32(BHDigits),i,p.NFI(1))

    for DecimalLessThanOrEqual(i,BlockHeightD) == true {
        fmt.Println("Adding BlockReward at BH,",BHAmountConv2Print(i),"...")
        BR2add := BlockRewardD(i)
	BrSum = ADDx(SP,BrSum,BR2add)
	SumNumberDigits = Count4Coma(BrSum)
	SP = CryptoplasmCurrencyPrecision + uint32(SumNumberDigits) + 1
	i = ADDx(uint32(BHDigits),i,p.NFI(1))
    }
    elapsed := time.Since(start)
    fmt.Println("Computing Sum intermittently for BH",BHAmountConv2Print(BlockHeightD),"took", elapsed, "and is ",BrSum)
    return BrSum
}