package Cryptoplasm_Blockchain_Constants

import (
	p "Cryptoplasm-Core/Packages/Firefly_Precision"
	"fmt"
	"log"
	"os"
	"time"
)

//================================================
//
// Slices/Lists Functions
// These Functions are used to compute the BlockReward
// Original implementation was made using Python Lists,
// therefore a similar mechanic will be used in Go as well..
//
// Function 01 - CryptoplasmPrimaryGeometricListing
//
// Creates a list of geometric heights using the input parameters
func CryptoplasmPrimaryGeometricListing(a *p.Decimal, b int) []*p.Decimal {
	HeightFront := a
	AreaFront := a
	FirstElement := ADDcp(AreaFront, Seed1st)
	Carnatz := make([]*p.Decimal, 1)
	Carnatz[0] = FirstElement

	for i := 2; i < b; i++ {
		HeightFront = ADDcp(HeightFront, Seed2nd)
		AreaFront = ADDcp(HeightFront, Seed1st)
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
// Function 01b - CryptoplasmSecondaryGeometricListing
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
		Missing = DIFcp(e, Sum1, Sum2)
	} else {
		SliceBack2 = CryptoplasmPrimaryGeometricListing(p.NFI(0), 638-intc)
		Sum3 := SumDL(SliceBack2)
		Missing = ADDcp(MULcp(ci, d), DIFcp(e, Sum1, Sum3))
	}
	SliceFront = append(SliceFront, Missing)
	RevSlice := Reverse(SliceBack)
	ResultList := AppDec(SliceFront, RevSlice)
	return ResultList
}

//================================================
//
// Func 02 - CryptoplasmGeometricKamelSequence
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

	PurpleUp1 := PurpleUpPeak
	PurpleDown1 := PurpleDownPeak
	//=======================================================
	PurpleUp2 := append(PurpleUp1, PurpleUp1...)
	PurpleUp3 := append(PurpleUp2, PurpleUp1...)
	PurpleUp4 := append(PurpleUp3, PurpleUp1...)
	PurpleUp5 := append(PurpleUp4, PurpleUp1...)
	PurpleUp6 := append(PurpleUp5, PurpleUp1...)
	PurpleDown2 := append(PurpleDown1, PurpleDown1...)
	PurpleDown3 := append(PurpleDown2, PurpleDown1...)
	PurpleDown4 := append(PurpleDown3, PurpleDown1...)
	PurpleDown5 := append(PurpleDown4, PurpleDown1...)
	PurpleDown6 := append(PurpleDown5, PurpleDown1...)

	IndigoUp1 := append(PurpleUp6, IndigoUpPeak...)
	IndigoDown1 := append(IndigoDownPeak, PurpleDown6...)
	//=======================================================
	IndigoUp2 := append(IndigoUp1, IndigoUp1...)
	IndigoUp3 := append(IndigoUp2, IndigoUp1...)
	IndigoUp4 := append(IndigoUp3, IndigoUp1...)
	IndigoUp5 := append(IndigoUp4, IndigoUp1...)
	IndigoUp6 := append(IndigoUp5, IndigoUp1...)
	IndigoDown2 := append(IndigoDown1, IndigoDown1...)
	IndigoDown3 := append(IndigoDown2, IndigoDown1...)
	IndigoDown4 := append(IndigoDown3, IndigoDown1...)
	IndigoDown5 := append(IndigoDown4, IndigoDown1...)
	IndigoDown6 := append(IndigoDown5, IndigoDown1...)

	BlueUp1 := append(IndigoUp6, PurpleUp4...)
	BlueUp1 = append(BlueUp1, BlueUpPeak...)
	BlueUp1 = append(BlueUp1, PurpleDown2...)
	BlueDown1 := append(PurpleUp4, BlueDownPeak...)
	BlueDown1 = append(BlueDown1, PurpleDown2...)
	BlueDown1 = append(BlueDown1, IndigoDown6...)
	//=======================================================
	BlueUp2 := append(BlueUp1, BlueUp1...)
	BlueUp3 := append(BlueUp2, BlueUp1...)
	BlueUp4 := append(BlueUp3, BlueUp1...)
	BlueUp5 := append(BlueUp4, BlueUp1...)
	BlueUp6 := append(BlueUp5, BlueUp1...)
	BlueDown2 := append(BlueDown1, BlueDown1...)
	BlueDown3 := append(BlueDown2, BlueDown1...)
	BlueDown4 := append(BlueDown3, BlueDown1...)
	BlueDown5 := append(BlueDown4, BlueDown1...)
	BlueDown6 := append(BlueDown5, BlueDown1...)

	GreenUp1 := append(BlueUp6, IndigoUp3...)
	GreenUp1 = append(GreenUp1, PurpleUp1...)
	GreenUp1 = append(GreenUp1, GreenUpPeak...)
	GreenUp1 = append(GreenUp1, PurpleDown5...)
	GreenUp1 = append(GreenUp1, IndigoDown3...)
	GreenDown1 := append(IndigoUp4, PurpleUp2...)
	GreenDown1 = append(GreenDown1, GreenDownPeak...)
	GreenDown1 = append(GreenDown1, PurpleDown4...)
	GreenDown1 = append(GreenDown1, IndigoDown2...)
	GreenDown1 = append(GreenDown1, BlueDown6...)
	//=======================================================
	GreenUp2 := append(GreenUp1, GreenUp1...)
	GreenUp3 := append(GreenUp2, GreenUp1...)
	GreenUp4 := append(GreenUp3, GreenUp1...)
	GreenUp5 := append(GreenUp4, GreenUp1...)
	GreenUp6 := append(GreenUp5, GreenUp1...)
	GreenDown2 := append(GreenDown1, GreenDown1...)
	GreenDown3 := append(GreenDown2, GreenDown1...)
	GreenDown4 := append(GreenDown3, GreenDown1...)
	GreenDown5 := append(GreenDown4, GreenDown1...)
	GreenDown6 := append(GreenDown5, GreenDown1...)

	YellowUp1 := append(GreenUp6, BlueUp2...)
	YellowUp1 = append(YellowUp1, IndigoUp1...)
	YellowUp1 = append(YellowUp1, PurpleUp2...)
	YellowUp1 = append(YellowUp1, YellowUpPeak...)
	YellowUp1 = append(YellowUp1, PurpleDown4...)
	YellowUp1 = append(YellowUp1, IndigoDown5...)
	YellowUp1 = append(YellowUp1, BlueDown4...)
	YellowDown1 := append(BlueUp4, IndigoUp3...)
	YellowDown1 = append(YellowDown1, PurpleUp2...)
	YellowDown1 = append(YellowDown1, YellowDownPeak...)
	YellowDown1 = append(YellowDown1, PurpleDown4...)
	YellowDown1 = append(YellowDown1, IndigoDown3...)
	YellowDown1 = append(YellowDown1, BlueDown2...)
	YellowDown1 = append(YellowDown1, GreenDown6...)
	//=======================================================
	YellowUp2 := append(YellowUp1, YellowUp1...)
	YellowUp3 := append(YellowUp2, YellowUp1...)
	YellowUp4 := append(YellowUp3, YellowUp1...)
	YellowUp5 := append(YellowUp4, YellowUp1...)
	YellowUp6 := append(YellowUp5, YellowUp1...)
	YellowDown2 := append(YellowDown1, YellowDown1...)
	YellowDown3 := append(YellowDown2, YellowDown1...)
	YellowDown4 := append(YellowDown3, YellowDown1...)
	YellowDown5 := append(YellowDown4, YellowDown1...)
	YellowDown6 := append(YellowDown5, YellowDown1...)

	OrangeUp1 := append(YellowUp6, BlueUp6...)
	OrangeUp1 = append(OrangeUp1, IndigoUp2...)
	OrangeUp1 = append(OrangeUp1, PurpleUp1...)
	OrangeUp1 = append(OrangeUp1, OrangeUpPeak...)
	OrangeUp1 = append(OrangeUp1, PurpleDown5...)
	OrangeUp1 = append(OrangeUp1, IndigoDown4...)
	OrangeUp1 = append(OrangeUp1, GreenDown6...)
	OrangeDown1 := append(GreenUp4, BlueUp4...) //Glitch in the Matrix
	OrangeDown1 = append(OrangeDown1, BlueUp1...)
	OrangeDown1 = append(OrangeDown1, IndigoUp4...)
	OrangeDown1 = append(OrangeDown1, PurpleUp3...)
	OrangeDown1 = append(OrangeDown1, OrangeDownPeak...)
	OrangeDown1 = append(OrangeDown1, PurpleDown3...)
	OrangeDown1 = append(OrangeDown1, IndigoDown2...)
	OrangeDown1 = append(OrangeDown1, BlueDown1...)
	OrangeDown1 = append(OrangeDown1, GreenDown2...)
	OrangeDown1 = append(OrangeDown1, YellowDown6...)
	//=======================================================
	OrangeUp2 := append(OrangeUp1, OrangeUp1...)
	OrangeUp3 := append(OrangeUp2, OrangeUp1...)
	OrangeUp4 := append(OrangeUp3, OrangeUp1...)
	OrangeUp5 := append(OrangeUp4, OrangeUp1...)
	//OrangeUp6 	:= append(OrangeUp5, OrangeUp1...)
	OrangeDown2 := append(OrangeDown1, OrangeDown1...)
	OrangeDown3 := append(OrangeDown2, OrangeDown1...)
	OrangeDown4 := append(OrangeDown3, OrangeDown1...)
	OrangeDown5 := append(OrangeDown4, OrangeDown1...)
	OrangeDown6 := append(OrangeDown5, OrangeDown1...)

	RedUp := append(OrangeUp5, YellowUp4...)
	RedUp = append(RedUp, GreenUp4...)
	RedUp = append(RedUp, BlueUp6...)
	RedUp = append(RedUp, IndigoUp2...)
	RedUp = append(RedUp, PurpleUp2...)
	RedUp = append(RedUp, RedUpPeak...)
	RedUp = append(RedUp, PurpleDown4...)
	RedUp = append(RedUp, IndigoDown4...)
	RedUp = append(RedUp, GreenDown2...)
	RedUp = append(RedUp, YellowDown2...)
	RedUp = append(RedUp, OrangeDown1...)
	RedDown1 := append(YellowUp5, GreenUp4...)
	RedDown1 = append(RedDown1, BlueUp3...)
	RedDown1 = append(RedDown1, IndigoUp6...)
	RedDown1 = append(RedDown1, PurpleUp2...)
	RedDown1 = append(RedDown1, RedDownPeak...)
	RedDown1 = append(RedDown1, PurpleDown4...)
	RedDown1 = append(RedDown1, BlueDown3...)
	RedDown1 = append(RedDown1, GreenDown2...)
	RedDown1 = append(RedDown1, YellowDown1...)
	RedDown1 = append(RedDown1, OrangeDown6...)

	RedDown2 := append(RedDown1, RedDown1...)
	RedDown3 := append(RedDown2, RedDown1...)
	RedDown4 := append(RedDown3, RedDown1...)
	RedDown5 := append(RedDown4, RedDown1...)

	Kamel := append(RedUp, OrangeUp2...)
	Kamel = append(Kamel, YellowUp3...)
	Kamel = append(Kamel, GreenUp1...)
	Kamel = append(Kamel, BlueUp2...)
	Kamel = append(Kamel, IndigoUp1...)
	Kamel = append(Kamel, PurpleUp3...)
	Kamel = append(Kamel, WhitePeak...)
	Kamel = append(Kamel, PurpleDown3...)
	Kamel = append(Kamel, IndigoDown5...)
	Kamel = append(Kamel, BlueDown4...)
	Kamel = append(Kamel, GreenDown5...)
	Kamel = append(Kamel, YellowDown3...)
	Kamel = append(Kamel, OrangeDown4...)
	Kamel = append(Kamel, RedDown5...)

	return Kamel
}
//================================================
//
// Func 03 - BlockReward
//
// BlockReward returns the Block Reward for a given Block Height
func BlockReward(BlockHeight uint64) *p.Decimal {
	//start := time.Now()
	GH := BlockGeometricHeight(BlockHeight)
	BR := ConvGH(GH)
	//elapsed := time.Since(start)
	//fmt.Println("Computing took", elapsed)
	return BR
}
//================================================
//
// Func 03b - ConvGH
//
// ConvGH returns the Block Reward for a Geometric Height
func ConvGH(GeometricHeight *p.Decimal) *p.Decimal {
	GH := GeometricHeight
	GHDigits := Count4Coma(GH)
	MP := uint32(GHDigits) + CryptoplasmSeedPrecision

	BR := ADDpr(MP, StartBRd, MULpr(MP, GH, Seed3rd))
	TruncatedBR := TruncToCurrency(BR)

	//Increasing highest Block for the whole BR sum to be exactly
	//TheoreticalEraEmission - PreMinedAmount
	Exception := SUBcp(TruncatedBR, Seed5th)
	IsHighestBlock := Exception.IsZero()
	if IsHighestBlock == true {
		TruncatedBR = ADDcp(DIFcp(TotalEmission, PreMine, Seed4th), TruncatedBR)
	}
	return TruncatedBR
}

//================================================
//
// Func 03c - BlockGeometricHeight
//
// BlockGeometricHeight returns the geometric height
// for a given BlockHeight on the Kamel Graph
// Used to calculate the BlockRewards
func BlockGeometricHeight(BlockHeight uint64) *p.Decimal {
	//Using uint64 as BlockHeight type makes this following code work
	//until BH 18.446.744.073.709.551.615. A higher value will overflow
	//the variable and cause the code to not work
	//Considering 1 minute per block, this BH would be attained in 35.1 trillion years
	//This wouldn't last until the end of the universe.
	//I still have time to modify the code until then.
	//Implementing ERAs will make this code last until the end of the Universe
	//1 Era equals 524.596.891 Blocks. Next Block would be Era2.Block1. Rewards will repeat.
	var (
		GH             = new(p.Decimal)
		CryptoplasmDNA = CryptoplasmGeometricKamelSequence()
		Purples        = BlockHeight / 637
		Rest           = BlockHeight % 637
	)
	if Purples == 0 {
		GH = Ax[Rest-1]
	} else if Purples != 0 && Purples <= uint64(len(CryptoplasmDNA)) {
		for i := 0; i < int(Purples); i++ {
			element := CryptoplasmDNA[i]
			if element == "A" {
				GH = ADDcp(GH, LastDE(Ax))
			} else if element == "B" {
				GH = ADDcp(GH, LastDE(Bx))
			} else if element == "C" {
				GH = ADDcp(GH, LastDE(Cx))
			} else if element == "D" {
				GH = ADDcp(GH, LastDE(Dx))
			} else if element == "E" {
				GH = ADDcp(GH, LastDE(Ex))
			} else if element == "F" {
				GH = ADDcp(GH, LastDE(Fx))
			} else if element == "G" {
				GH = ADDcp(GH, LastDE(Gx))
			} else if element == "H" {
				GH = ADDcp(GH, LastDE(Hx))
			} else if element == "I" {
				GH = ADDcp(GH, LastDE(Ix))
			} else if element == "J" {
				GH = ADDcp(GH, LastDE(Jx))
			} else if element == "K" {
				GH = ADDcp(GH, LastDE(Kx))
			} else if element == "L" {
				GH = ADDcp(GH, LastDE(Lx))
			} else if element == "M" {
				GH = ADDcp(GH, LastDE(Mx))
			} else if element == "N" {
				GH = ADDcp(GH, LastDE(Nx))
			} else {
				GH = ADDcp(GH, LastDE(Ox))
			}
		}
		GH = SUBcp(GH, MULcp(p.NFI(int64(Purples)), Seed1st))
		if Rest == 0 {
			GH = ADDcp(GH, Seed1st)
		} else if Purples < uint64(len(CryptoplasmDNA)) {
			element := CryptoplasmDNA[Purples]
			if element == "A" {
				GH = ADDcp(GH, Ax[Rest-1])
			} else if element == "B" {
				GH = ADDcp(GH, Bx[Rest-1])
			} else if element == "C" {
				GH = ADDcp(GH, Cx[Rest-1])
			} else if element == "D" {
				GH = ADDcp(GH, Dx[Rest-1])
			} else if element == "E" {
				GH = ADDcp(GH, Ex[Rest-1])
			} else if element == "F" {
				GH = ADDcp(GH, Fx[Rest-1])
			} else if element == "G" {
				GH = ADDcp(GH, Gx[Rest-1])
			} else if element == "H" {
				GH = ADDcp(GH, Hx[Rest-1])
			} else if element == "I" {
				GH = ADDcp(GH, Ix[Rest-1])
			} else if element == "J" {
				GH = ADDcp(GH, Jx[Rest-1])
			} else if element == "K" {
				GH = ADDcp(GH, Kx[Rest-1])
			} else if element == "L" {
				GH = ADDcp(GH, Lx[Rest-1])
			} else if element == "M" {
				GH = ADDcp(GH, Mx[Rest-1])
			} else if element == "N" {
				GH = ADDcp(GH, Nx[Rest-1])
			} else {
				GH = ADDcp(GH, Ox[Rest-1])
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

//================================================
//
// Func 4 - ExportBr
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
	fmt.Println("Please be patient as the first Blue Interval is computed...")
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
			BlocksNo := 31213 * Periods
			fmt.Println("Written", Periods, "/2041 Blue Periods, ", BlocksNo, "Blocks, Elapsed Time ", intermediary)
			//defer timeTrack(time.Now(), "Written")
		}
		element := CryptoplasmDNA[i]
		if element == "A" {
			for j := 0; j < 637; j++ {
				GH = ADDcp(B, Ax[j])
				BR = ConvGH(GH)
				_, _ = fmt.Fprintln(OutputFile, BR)
				S = ADDcp(S, BR)
				if j == 636 {
					B = SUBcp(GH, Seed1st)
				}
			}
		} else if element == "B" {
			for j := 0; j < 637; j++ {
				GH = ADDcp(B, Bx[j])
				BR = ConvGH(GH)
				_, _ = fmt.Fprintln(OutputFile, BR)
				S = ADDcp(S, BR)
				if j == 636 {
					B = SUBcp(GH, Seed1st)
				}
			}
		} else if element == "C" {
			for j := 0; j < 637; j++ {
				GH = ADDcp(B, Cx[j])
				BR = ConvGH(GH)
				_, _ = fmt.Fprintln(OutputFile, BR)
				S = ADDcp(S, BR)
				if j == 636 {
					B = SUBcp(GH, Seed1st)
				}
			}
		} else if element == "D" {
			for j := 0; j < 637; j++ {
				GH = ADDcp(B, Dx[j])
				BR = ConvGH(GH)
				_, _ = fmt.Fprintln(OutputFile, BR)
				S = ADDcp(S, BR)
				if j == 636 {
					B = SUBcp(GH, Seed1st)
				}
			}
		} else if element == "E" {
			for j := 0; j < 637; j++ {
				GH = ADDcp(B, Ex[j])
				BR = ConvGH(GH)
				_, _ = fmt.Fprintln(OutputFile, BR)
				S = ADDcp(S, BR)
				if j == 636 {
					B = SUBcp(GH, Seed1st)
				}
			}
		} else if element == "F" {
			for j := 0; j < 637; j++ {
				GH = ADDcp(B, Fx[j])
				BR = ConvGH(GH)
				_, _ = fmt.Fprintln(OutputFile, BR)
				S = ADDcp(S, BR)
				if j == 636 {
					B = SUBcp(GH, Seed1st)
				}
			}
		} else if element == "G" {
			for j := 0; j < 637; j++ {
				GH = ADDcp(B, Gx[j])
				BR = ConvGH(GH)
				_, _ = fmt.Fprintln(OutputFile, BR)
				S = ADDcp(S, BR)
				if j == 636 {
					B = SUBcp(GH, Seed1st)
				}
			}
		} else if element == "H" {
			for j := 0; j < 637; j++ {
				GH = ADDcp(B, Hx[j])
				BR = ConvGH(GH)
				_, _ = fmt.Fprintln(OutputFile, BR)
				S = ADDcp(S, BR)
				if j == 636 {
					B = SUBcp(GH, Seed1st)
				}
			}
		} else if element == "I" {
			for j := 0; j < 637; j++ {
				GH = ADDcp(B, Ix[j])
				BR = ConvGH(GH)
				_, _ = fmt.Fprintln(OutputFile, BR)
				S = ADDcp(S, BR)
				if j == 636 {
					B = SUBcp(GH, Seed1st)
				}
			}
		} else if element == "J" {
			for j := 0; j < 637; j++ {
				GH = ADDcp(B, Jx[j])
				BR = ConvGH(GH)
				_, _ = fmt.Fprintln(OutputFile, BR)
				S = ADDcp(S, BR)
				if j == 636 {
					B = SUBcp(GH, Seed1st)
				}
			}
		} else if element == "K" {
			for j := 0; j < 637; j++ {
				GH = ADDcp(B, Kx[j])
				BR = ConvGH(GH)
				_, _ = fmt.Fprintln(OutputFile, BR)
				S = ADDcp(S, BR)
				if j == 636 {
					B = SUBcp(GH, Seed1st)
				}
			}
		} else if element == "L" {
			for j := 0; j < 637; j++ {
				GH = ADDcp(B, Lx[j])
				BR = ConvGH(GH)
				_, _ = fmt.Fprintln(OutputFile, BR)
				S = ADDcp(S, BR)
				if j == 636 {
					B = SUBcp(GH, Seed1st)
				}
			}
		} else if element == "M" {
			for j := 0; j < 637; j++ {
				GH = ADDcp(B, Mx[j])
				BR = ConvGH(GH)
				_, _ = fmt.Fprintln(OutputFile, BR)
				S = ADDcp(S, BR)
				if j == 636 {
					B = SUBcp(GH, Seed1st)
				}
			}
		} else if element == "N" {
			for j := 0; j < 637; j++ {
				GH = ADDcp(B, Nx[j])
				BR = ConvGH(GH)
				_, _ = fmt.Fprintln(OutputFile, BR)
				S = ADDcp(S, BR)
				if j == 636 {
					B = SUBcp(GH, Seed1st)
				}
			}
		} else if element == "O" {
			for j := 0; j < 637; j++ {
				GH = ADDcp(B, Ox[j])
				BR = ConvGH(GH)
				_, _ = fmt.Fprintln(OutputFile, BR)
				S = ADDcp(S, BR)
				if j == 636 {
					B = SUBcp(GH, Seed1st)
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
