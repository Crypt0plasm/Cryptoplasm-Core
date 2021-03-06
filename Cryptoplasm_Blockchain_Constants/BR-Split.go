package Cryptoplasm_Blockchain_Constants

import (
    "fmt"

    p "github.com/Crypt0plasm/Firefly-APD"
)
//
//	BR-Split.go						Block-Reward Split Computing Function
//================================================================================================
//************************************************************************************************
//================================================================================================
// 	Function List:
//
//	01 Comparison Functions operating on decimal type
//		01  - BaseStakingPercent			returns the Block Staking Reward % for a given Block-Height
//		02  - TreasurySplit				returns Multiple Split % for a given Block-Height
//		03  - ProcentSplitMatrix			returns all Split % for a given Block-Height
//		04  - ProcentSplitPrinter			prints information regarding the Block-Reward split
//================================================
//
// Func 01 - BaseStakingPercent
//
// BaseStakingPercent returns the Block Staking Reward % for a given Block-Height
// The Block-Height Type is a decimal.
func BaseStakingPercent(BlockHeightD *p.Decimal) *p.Decimal {
    var (
	Gray		= int64(91)
	PurpleInt,_ 	= Purple.Int64()
	IndigoInt,_ 	= Indigo.Int64()
	BlueInt,_ 	= Blue.Int64()
	GreenInt,_ 	= Green.Int64()
	YellowInt,_ 	= Yellow.Int64()
	OrangeInt,_ 	= Orange.Int64()

	From01		= int64(1)
	To01		= Gray - 1
	From02 		= Gray
	To02 		= 2 * Gray - 1
	From03		= 2 * Gray
	To03		= 3 * Gray - 1
	From04		= 3 * Gray
	To04		= 4 * Gray - 1
	From05		= 4 * Gray
	To05		= 5 * Gray - 1
	From06		= 5 * Gray
	To06		= 6 * Gray - 1
	From07		= 6 * Gray
	To07		= 7 * Gray - 1
	From08 		= PurpleInt
	To08 		= 2 * PurpleInt - 1
	From09		= 2 * PurpleInt
	To09		= 3 * PurpleInt - 1
	From10		= 3 * PurpleInt
	To10		= 4 * PurpleInt - 1
	From11		= 4 * PurpleInt
	To11		= 5 * PurpleInt - 1
	From12		= 5 * PurpleInt
	To12		= 6 * PurpleInt - 1
	From13		= 6 * PurpleInt
	To13		= 7 * PurpleInt - 1
	From14 		= IndigoInt
	To14 		= 2 * IndigoInt - 1
	From15		= 2 * IndigoInt
	To15		= 3 * IndigoInt - 1
	From16		= 3 * IndigoInt
	To16		= 4 * IndigoInt - 1
	From17		= 4 * IndigoInt
	To17		= 5 * IndigoInt - 1
	From18		= 5 * IndigoInt
	To18		= 6 * IndigoInt - 1
	From19		= 6 * IndigoInt
	To19		= 7 * IndigoInt - 1
	From20 		= BlueInt
	To20 		= 2 * BlueInt - 1
	From21		= 2 * BlueInt
	To21		= 3 * BlueInt - 1
	From22		= 3 * BlueInt
	To22		= 4 * BlueInt - 1
	From23		= 4 * BlueInt
	To23		= 5 * BlueInt - 1
	From24		= 5 * BlueInt
	To24		= 6 * BlueInt - 1
	From25		= 6 * BlueInt
	To25		= 7 * BlueInt - 1
	From26 		= GreenInt
	To26 		= 2 * GreenInt - 1
	From27		= 2 * GreenInt
	To27		= 3 * GreenInt - 1
	From28		= 3 * GreenInt
	To28		= 4 * GreenInt - 1
	From29		= 4 * GreenInt
	To29		= 5 * GreenInt - 1
	From30		= 5 * GreenInt
	To30		= 6 * GreenInt - 1
	From31		= 6 * GreenInt
	To31		= 7 * GreenInt - 1
	From32 		= YellowInt
	To32 		= 2 * YellowInt - 1
	From33		= 2 * YellowInt
	To33		= 3 * YellowInt - 1
	From34		= 3 * YellowInt
	To34		= 4 * YellowInt - 1
	From35		= 4 * YellowInt
	To35		= 5 * YellowInt - 1
	From36		= 5 * YellowInt
	To36		= 6 * YellowInt - 1
	From37		= 6 * YellowInt
	To37		= 7 * YellowInt - 1
	From38 		= OrangeInt
	To38 		= 2 * OrangeInt - 1
	From39		= 2 * OrangeInt
	To39		= 3 * OrangeInt - 1
	From40		= 3 * OrangeInt
	To40		= 4 * OrangeInt - 1
	From41		= 4 * OrangeInt
	To41		= 5 * OrangeInt - 1
	From42		= 5 * OrangeInt
	To42		= 6 * OrangeInt - 1

	No01 		= [...]int64{From01,To01}
	No02 		= [...]int64{From02,To02}
	No03 		= [...]int64{From03,To03}
	No04 		= [...]int64{From04,To04}
	No05 		= [...]int64{From05,To05}
	No06 		= [...]int64{From06,To06}
	No07 		= [...]int64{From07,To07}
	No08 		= [...]int64{From08,To08}
	No09 		= [...]int64{From09,To09}
	No10 		= [...]int64{From10,To10}
	No11 		= [...]int64{From11,To11}
	No12 		= [...]int64{From12,To12}
	No13 		= [...]int64{From13,To13}
	No14 		= [...]int64{From14,To14}
	No15 		= [...]int64{From15,To15}
	No16 		= [...]int64{From16,To16}
	No17 		= [...]int64{From17,To17}
	No18 		= [...]int64{From18,To18}
	No19 		= [...]int64{From19,To19}
	No20 		= [...]int64{From20,To20}
	No21 		= [...]int64{From21,To21}
	No22 		= [...]int64{From22,To22}
	No23 		= [...]int64{From23,To23}
	No24 		= [...]int64{From24,To24}
	No25 		= [...]int64{From25,To25}
	No26 		= [...]int64{From26,To26}
	No27 		= [...]int64{From27,To27}
	No28 		= [...]int64{From28,To28}
	No29 		= [...]int64{From29,To29}
	No30 		= [...]int64{From30,To30}
	No31 		= [...]int64{From31,To31}
	No32 		= [...]int64{From32,To32}
	No33 		= [...]int64{From33,To33}
	No34 		= [...]int64{From34,To34}
	No35 		= [...]int64{From35,To35}
	No36 		= [...]int64{From36,To36}
	No37 		= [...]int64{From37,To37}
	No38 		= [...]int64{From38,To38}
	No39 		= [...]int64{From39,To39}
	No40 		= [...]int64{From40,To40}
	No41 		= [...]int64{From41,To41}
	No42 		= [...]int64{From42,To42}
        
    	GrayArray	= [6][2]int64{No02,No03,No04,No05,No06,No07}
    	PurpleArray	= [6][2]int64{No08,No09,No10,No11,No12,No13}
    	IndigoArray	= [6][2]int64{No14,No15,No16,No17,No18,No19}
    	BlueArray	= [6][2]int64{No20,No21,No22,No23,No24,No25}
    	GreenArray	= [6][2]int64{No26,No27,No28,No29,No30,No31}
    	YellowArray	= [6][2]int64{No32,No33,No34,No35,No36,No37}
    	OrangeArray	= [6][2]int64{No38,No39,No40,No41,No42,No01}
    	STAr		= [7][6][2]int64{GrayArray,PurpleArray,IndigoArray,BlueArray,GreenArray,YellowArray,OrangeArray}
	BH 		= BlockHeightD
	BSP		= new(p.Decimal)
    )

    if DecimalLessThanOrEqual(p.NFI(STAr[6][5][0]),BH) == true && DecimalLessThanOrEqual(BH,p.NFI(STAr[6][5][1])) == true {
	BSP		= p.NFS("1.1")
    } else if DecimalLessThanOrEqual(p.NFI(STAr[0][0][0]),BH) == true && DecimalLessThanOrEqual(BH,p.NFI(STAr[0][0][1])) == true {
	BSP		= p.NFS("2.2")
    } else if DecimalLessThanOrEqual(p.NFI(STAr[0][1][0]),BH) == true && DecimalLessThanOrEqual(BH,p.NFI(STAr[0][1][1])) == true {
	BSP		= p.NFS("3.3")
    } else if DecimalLessThanOrEqual(p.NFI(STAr[0][2][0]),BH) == true && DecimalLessThanOrEqual(BH,p.NFI(STAr[0][2][1])) == true {
	BSP		= p.NFS("4.4")
    } else if DecimalLessThanOrEqual(p.NFI(STAr[0][3][0]),BH) == true && DecimalLessThanOrEqual(BH,p.NFI(STAr[0][3][1])) == true {
	BSP		= p.NFS("5.5")
    } else if DecimalLessThanOrEqual(p.NFI(STAr[0][4][0]),BH) == true && DecimalLessThanOrEqual(BH,p.NFI(STAr[0][4][1])) == true {
	BSP		= p.NFS("6.6")
    } else if DecimalLessThanOrEqual(p.NFI(STAr[0][5][0]),BH) == true && DecimalLessThanOrEqual(BH,p.NFI(STAr[0][5][1])) == true {
	BSP		= p.NFS("7.7")
    } else if DecimalLessThanOrEqual(p.NFI(STAr[1][0][0]),BH) == true && DecimalLessThanOrEqual(BH,p.NFI(STAr[1][0][1])) == true {
	BSP		= p.NFS("8.8")
    } else if DecimalLessThanOrEqual(p.NFI(STAr[1][1][0]),BH) == true && DecimalLessThanOrEqual(BH,p.NFI(STAr[1][1][1])) == true {
	BSP		= p.NFS("9.9")
    } else if DecimalLessThanOrEqual(p.NFI(STAr[1][2][0]),BH) == true && DecimalLessThanOrEqual(BH,p.NFI(STAr[1][2][1])) == true {
	BSP		= p.NFS("11")
    } else if DecimalLessThanOrEqual(p.NFI(STAr[1][3][0]),BH) == true && DecimalLessThanOrEqual(BH,p.NFI(STAr[1][3][1])) == true {
	BSP		= p.NFS("12.1")
    } else if DecimalLessThanOrEqual(p.NFI(STAr[1][4][0]),BH) == true && DecimalLessThanOrEqual(BH,p.NFI(STAr[1][4][1])) == true {
	BSP		= p.NFS("13.2")
    } else if DecimalLessThanOrEqual(p.NFI(STAr[1][5][0]),BH) == true && DecimalLessThanOrEqual(BH,p.NFI(STAr[1][5][1])) == true {
	BSP		= p.NFS("14.3")
    } else if DecimalLessThanOrEqual(p.NFI(STAr[2][0][0]),BH) == true && DecimalLessThanOrEqual(BH,p.NFI(STAr[2][0][1])) == true {
	BSP		= p.NFS("15.4")
    } else if DecimalLessThanOrEqual(p.NFI(STAr[2][1][0]),BH) == true && DecimalLessThanOrEqual(BH,p.NFI(STAr[2][1][1])) == true {
	BSP		= p.NFS("16.5")
    } else if DecimalLessThanOrEqual(p.NFI(STAr[2][2][0]),BH) == true && DecimalLessThanOrEqual(BH,p.NFI(STAr[2][2][1])) == true {
	BSP		= p.NFS("17.6")
    } else if DecimalLessThanOrEqual(p.NFI(STAr[2][3][0]),BH) == true && DecimalLessThanOrEqual(BH,p.NFI(STAr[2][3][1])) == true {
	BSP		= p.NFS("18.7")
    } else if DecimalLessThanOrEqual(p.NFI(STAr[2][4][0]),BH) == true && DecimalLessThanOrEqual(BH,p.NFI(STAr[2][4][1])) == true {
	BSP		= p.NFS("19.8")
    } else if DecimalLessThanOrEqual(p.NFI(STAr[2][5][0]),BH) == true && DecimalLessThanOrEqual(BH,p.NFI(STAr[2][5][1])) == true {
	BSP		= p.NFS("20.9")
    } else if DecimalLessThanOrEqual(p.NFI(STAr[3][0][0]),BH) == true && DecimalLessThanOrEqual(BH,p.NFI(STAr[3][0][1])) == true {
	BSP		= p.NFS("22")
    } else if DecimalLessThanOrEqual(p.NFI(STAr[3][1][0]),BH) == true && DecimalLessThanOrEqual(BH,p.NFI(STAr[3][1][1])) == true {
	BSP		= p.NFS("23.1")
    } else if DecimalLessThanOrEqual(p.NFI(STAr[3][2][0]),BH) == true && DecimalLessThanOrEqual(BH,p.NFI(STAr[3][2][1])) == true {
	BSP		= p.NFS("24.2")
    } else if DecimalLessThanOrEqual(p.NFI(STAr[3][3][0]),BH) == true && DecimalLessThanOrEqual(BH,p.NFI(STAr[3][3][1])) == true {
	BSP		= p.NFS("25.3")
    } else if DecimalLessThanOrEqual(p.NFI(STAr[3][4][0]),BH) == true && DecimalLessThanOrEqual(BH,p.NFI(STAr[3][4][1])) == true {
	BSP		= p.NFS("26.4")
    } else if DecimalLessThanOrEqual(p.NFI(STAr[3][5][0]),BH) == true && DecimalLessThanOrEqual(BH,p.NFI(STAr[3][5][1])) == true {
	BSP		= p.NFS("27.5")
    } else if DecimalLessThanOrEqual(p.NFI(STAr[4][0][0]),BH) == true && DecimalLessThanOrEqual(BH,p.NFI(STAr[4][0][1])) == true {
	BSP		= p.NFS("28.6")
    } else if DecimalLessThanOrEqual(p.NFI(STAr[4][1][0]),BH) == true && DecimalLessThanOrEqual(BH,p.NFI(STAr[4][1][1])) == true {
	BSP		= p.NFS("29.7")
    } else if DecimalLessThanOrEqual(p.NFI(STAr[4][2][0]),BH) == true && DecimalLessThanOrEqual(BH,p.NFI(STAr[4][2][1])) == true {
	BSP		= p.NFS("30.8")
    } else if DecimalLessThanOrEqual(p.NFI(STAr[4][3][0]),BH) == true && DecimalLessThanOrEqual(BH,p.NFI(STAr[4][3][1])) == true {
	BSP		= p.NFS("31.9")
    } else if DecimalLessThanOrEqual(p.NFI(STAr[4][4][0]),BH) == true && DecimalLessThanOrEqual(BH,p.NFI(STAr[4][4][1])) == true {
	BSP		= p.NFS("33")
    } else if DecimalLessThanOrEqual(p.NFI(STAr[4][5][0]),BH) == true && DecimalLessThanOrEqual(BH,p.NFI(STAr[4][5][1])) == true {
	BSP		= p.NFS("34.1")
    } else if DecimalLessThanOrEqual(p.NFI(STAr[5][0][0]),BH) == true && DecimalLessThanOrEqual(BH,p.NFI(STAr[5][0][1])) == true {
	BSP		= p.NFS("35.2")
    } else if DecimalLessThanOrEqual(p.NFI(STAr[5][1][0]),BH) == true && DecimalLessThanOrEqual(BH,p.NFI(STAr[5][1][1])) == true {
	BSP		= p.NFS("36.3")
    } else if DecimalLessThanOrEqual(p.NFI(STAr[5][2][0]),BH) == true && DecimalLessThanOrEqual(BH,p.NFI(STAr[5][2][1])) == true {
	BSP		= p.NFS("37.4")
    } else if DecimalLessThanOrEqual(p.NFI(STAr[5][3][0]),BH) == true && DecimalLessThanOrEqual(BH,p.NFI(STAr[5][3][1])) == true {
	BSP		= p.NFS("38.5")
    } else if DecimalLessThanOrEqual(p.NFI(STAr[5][4][0]),BH) == true && DecimalLessThanOrEqual(BH,p.NFI(STAr[5][4][1])) == true {
	BSP		= p.NFS("39.6")
    } else if DecimalLessThanOrEqual(p.NFI(STAr[5][5][0]),BH) == true && DecimalLessThanOrEqual(BH,p.NFI(STAr[5][5][1])) == true {
	BSP		= p.NFS("40.7")
    } else if DecimalLessThanOrEqual(p.NFI(STAr[6][0][0]),BH) == true && DecimalLessThanOrEqual(BH,p.NFI(STAr[6][0][1])) == true {
	BSP		= p.NFS("41.8")
    } else if DecimalLessThanOrEqual(p.NFI(STAr[6][1][0]),BH) == true && DecimalLessThanOrEqual(BH,p.NFI(STAr[6][1][1])) == true {
	BSP		= p.NFS("42.9")
    } else if DecimalLessThanOrEqual(p.NFI(STAr[6][2][0]),BH) == true && DecimalLessThanOrEqual(BH,p.NFI(STAr[6][2][1])) == true {
	BSP		= p.NFS("44")
    } else if DecimalLessThanOrEqual(p.NFI(STAr[6][3][0]),BH) == true && DecimalLessThanOrEqual(BH,p.NFI(STAr[6][3][1])) == true {
	BSP		= p.NFS("45.1")
    } else if DecimalLessThanOrEqual(p.NFI(STAr[6][4][0]),BH) == true && DecimalLessThanOrEqual(BH,p.NFI(STAr[6][4][1])) == true {
	BSP		= p.NFS("46.2")
    } else {
	BSP		= p.NFS("47.3")
    }

    return BSP
}
//================================================
//
// Func 02 - TreasurySplit
//
// TreasurySplit returns Multiple Split % for a given Block-Height.
// The Block-Height Type is decimal.
func TreasurySplit(BlockHeightD *p.Decimal) [4]*p.Decimal {
	//start := time.Now()
	var(
		TreasuryDecay = p.NFS("0.0006")
		TreasurySevenPart = p.NFS("0.6")
		TDUSTK = new(p.Decimal)
		TDUMNG = new(p.Decimal)
		TDUCQP = new(p.Decimal)
		CONTNP = new(p.Decimal)
	)

	BluePeriod := ADDs(DivInt(BlockHeightD,Blue),p.NFI(1))
	if DecimalEqual(BluePeriod,p.NFI(1)) == true {
		TDUSTK = p.NFS("0")
		TDUMNG = p.NFS("0")
		TDUCQP = p.NFS("0")
		CONTNP = TreasurySevenPart
	} else if DecimalLessThanOrEqual(BluePeriod,p.NFI(21)) == true {
		TDUSTK = MULxc(SUBs(BluePeriod,p.NFI(1)),p.NFI(2))
		TDUMNG = MULxc(SUBs(BluePeriod,p.NFI(1)),p.NFI(2))
		TDUCQP = MULxc(SUBs(BluePeriod,p.NFI(1)),p.NFI(3))
		CONTNP = SUBs(TreasurySevenPart,MULxc(SUBs(BluePeriod,p.NFI(1)),TreasuryDecay))
	} else if DecimalLessThanOrEqual(p.NFI(22),BluePeriod) == true && DecimalLessThanOrEqual(BluePeriod,p.NFI(990)) == true {
		TDUSTK = ADDs(MULxc(SUBs(BluePeriod,p.NFI(21)),p.NFS("2.5")),p.NFI(40))
		TDUMNG = ADDs(MULxc(SUBs(BluePeriod,p.NFI(21)),p.NFS("2.5")),p.NFI(40))
		TDUCQP = ADDs(MULxc(SUBs(BluePeriod,p.NFI(21)),p.NFI(2)),p.NFI(60))
		CONTNP = SUBs(TreasurySevenPart,MULxc(SUBs(BluePeriod,p.NFI(1)),TreasuryDecay))
	} else {
		TDUSTK = p.NFI(2465)
		TDUMNG = p.NFI(2465)
		TDUCQP = p.NFI(2000)
	}
	STK := MULxc(TDUSTK,TreasuryDecay)
	MNG := MULxc(TDUMNG,TreasuryDecay)
	CQP := MULxc(TDUCQP,TreasuryDecay)

	ResultMatrix := [...]*p.Decimal{CONTNP, STK, MNG, CQP}
	return ResultMatrix
	//elapsed := time.Since(start)
	//fmt.Println("Computing took", elapsed)
}
//================================================
//
// Func 03 - ProcentSplitMatrix
//
// ProcentSplitMatrix returns all Split % for a given Block-Height
// inside an array of arrays.
func ProcentSplitMatrix(BlockHeightD *p.Decimal) [4][3]*p.Decimal {
	var (
		StarProcentGOLD			= TruncPercent(p.NFS("0.1"))
		StarProcentSILVER		= TruncPercent(p.NFS("0.2"))
		StarProcentBRONZE		= TruncPercent(p.NFS("0.4"))

		BaseConquerorProcent 		= TruncPercent(p.NFS("0.8"))
		BaseStakingProcent 		= TruncPercent(BaseStakingPercent(BlockHeightD))
		TreasurySplitMatrix 		= TreasurySplit(BlockHeightD)
		BaseContinentProcent 		= TruncPercent(TreasurySplitMatrix[0])
		SecondaryStakingProcent 	= TruncPercent(TreasurySplitMatrix[1])
		SecondaryMiningProcent		= TruncPercent(TreasurySplitMatrix[2])
		SecondaryConquerorProcent 	= TruncPercent(TreasurySplitMatrix[3])

		TotalStakingProcent		= TruncPercent(ADDs(BaseStakingProcent,SecondaryStakingProcent))
		TotalConquerorProcent		= TruncPercent(ADDs(BaseConquerorProcent,SecondaryConquerorProcent))
		TreasuryProcent			= TruncPercent(MULxc(BaseContinentProcent,p.NFI(7)))
		StarProcent			= TruncPercent(SUMs(StarProcentGOLD,StarProcentSILVER,StarProcentBRONZE))

		TotalMiningProcent		= TruncPercent(DIFs(p.NFI(100),TotalStakingProcent,TotalConquerorProcent,TreasuryProcent,StarProcent))
		BaseMiningProcent		= TruncPercent(SUBs(TotalMiningProcent,SecondaryMiningProcent))
	)

	MiningProcentMatrix := [...]*p.Decimal{BaseMiningProcent,SecondaryMiningProcent,TotalMiningProcent}
	StakingProcentMatrix := [...]*p.Decimal{BaseStakingProcent,SecondaryStakingProcent,TotalStakingProcent}
	ConquerorProcentMatrix := [...]*p.Decimal{BaseConquerorProcent,SecondaryConquerorProcent,TotalConquerorProcent}
	ContinentProcentMatrix := [...]*p.Decimal{BaseContinentProcent,TreasuryProcent,p.NFI(777)}
	SplitMatrix := [4][3]*p.Decimal{MiningProcentMatrix,StakingProcentMatrix,ConquerorProcentMatrix,ContinentProcentMatrix}

	return SplitMatrix
}
//================================================
//
// Func 04 - ProcentSplitPrinter
//
// ProcentSplitPrinter prints information regarding the Block-Reward split
// for a given Block-Height
func ProcentSplitPrinter(BlockHeightD *p.Decimal) {
	M := ProcentSplitMatrix(BlockHeightD)

	fmt.Println("At BH",BHAmountConv2Print(BlockHeightD),"Miners get:")
	fmt.Println("A Base mining % of:",M[0][0],"%")
	fmt.Println("A Bonus mining % of:",M[0][1],"%")
	fmt.Println("A Total mining % of:",M[0][2],"%")
	fmt.Println("")
	fmt.Println("At BH",BHAmountConv2Print(BlockHeightD),"Stakers get:")
	fmt.Println("A Base staking % of:",M[1][0],"%")
	fmt.Println("A Bonus staking % of:",M[1][1],"%")
	fmt.Println("A Total staking % of:",M[1][2],"%")
	fmt.Println("")
	fmt.Println("At BH",BHAmountConv2Print(BlockHeightD),"Conqueror Pool gets:")
	fmt.Println("A Base % of:",M[2][0],"%")
	fmt.Println("A Bonus % of:",M[2][1],"%")
	fmt.Println("A Total % of:",M[2][2],"%")
	fmt.Println("")
	fmt.Println("Regardless of BH Star Owners get a fixed 0.7000 %")
	fmt.Println("0.1 % split between all GOLD Star Owners")
	fmt.Println("0.2 % split between all SILVER Star Owners")
	fmt.Println("0.4 % split between all BRONZE Star Owners")
	fmt.Println("")
	fmt.Println("Treasury starts with 0.6 % per Continent for a TOTAL of 4.2000 %")
	fmt.Println("But At BH",BHAmountConv2Print(BlockHeightD),"Treasury gets:")
	fmt.Println(M[3][0],"% per Continent")
	fmt.Println("for a total of",M[3][1],"%")

}