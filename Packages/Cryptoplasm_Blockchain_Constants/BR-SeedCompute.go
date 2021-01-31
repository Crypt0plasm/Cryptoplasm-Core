package Cryptoplasm_Blockchain_Constants

import (
	p "Cryptoplasm-Core/Packages/Firefly_Precision"
	"fmt"
	"time"
)

func CryptoplasmDecimalSeedComputer () {
	var(
		//Minimum math precision that obtains the same 40 decimal precision
		//obtained with the higher much more computing intensive 250 math precision
		Pp  uint32		= 11		//Purple 		Precision
		Ip  uint32		= 13		//Indigo 		Precision
		Bp  uint32		= 14		//Blue 			Precision
		Gp  uint32		= 16		//Green 		Precision
		YUp uint32		= 17		//Yellow Up 	Precision
		YDp uint32		= 16		//Yellow Down 	Precision
		Op  uint32		= 18		//Orange		Precision
		RUp uint32		= 18		//Red Up		Precision
		RDp uint32		= 20		//Red Down		Precision
		Whp uint32		= 20		//White			Precision

		CSP				= CryptoplasmSeedPrecision
	)

	//===============
	start := time.Now()
	fmt.Println("========================================================")
	fmt.Println("Computing and Printing Secondary Decimal Seeds")
	//
	//
	//Calculating PurpleUp
	//
	y, z, w := ASApr(CSP+Pp,p.NFI(90), Purple, p.NFI(55))
	PurpleBaseUp := w
	PurpleSaveUp := y
	PurpleUpHeight := z

	y, z, w = ASApr(CSP+Pp,p.NFI(111), y, p.NFI(1))
	PurpleUp := w
	PurpleTotalUp := ADDpr(CSP+Pp,PurpleBaseUp, PurpleUp)

	PurpleUpHeightSeed := TruncSeed(PurpleUpHeight)
	fmt.Println("PurpleUpHeightSeed is:         ", PurpleUpHeightSeed)

	PurpleUpAreaSeed := TruncSeed(PurpleTotalUp)
	fmt.Println("PurpleUpAreaSeed is:        ", PurpleUpAreaSeed)
	//
	//
	//Calculating PurpleDown
	//
	y, z, w = ASApr(CSP+Pp,p.NFI(49), Purple, p.NFI(90))
	PurpleBaseDown := w
	PurpleSaveDown := z
	PurpleDownHeight := y

	y, z, w = ASApr(CSP+Pp,p.NFI(7), z, p.NFI(105))
	PurpleDown := w
	PurpleTotalDown := ADDpr(CSP+Pp,PurpleBaseDown, PurpleDown)

	PurpleDownHeightSeedC := TruncSeed(PurpleDownHeight)
	fmt.Println("PurpleDownHeightSeed is:       ", PurpleDownHeightSeedC)

	PurpleDownAreaSeedC := TruncSeed(PurpleTotalDown)
	fmt.Println("PurpleDownAreaSeed is:      ", PurpleDownAreaSeedC)
	fmt.Println("========================================================")
	//
	//
	//Calculating IndigoUp
	//
	//fmt.Println("Purple este",Purple)
	y, z, w = ASApr(CSP+Ip,p.NFI(90), Indigo, p.NFI(53))
	IndigoBaseUp := w
	IndigoSaveUp := y
	IndigoUpHeight := z

	y, z, w = ASApr(CSP+Ip,p.NFI(102), y, p.NFI(2))
	IndigoUp := w
	IndigoPurpleSideLeft := SUBpr(CSP+Ip,y, MULpr(CSP+Ip,PurpleSaveUp, p.NFI(6)))
	IndigoSave01 := z

	y, z, w = ASApr(CSP+Ip,p.NFI(104), IndigoPurpleSideLeft, p.NFI(1))
	IndigoUp = ADDpr(CSP+Ip,IndigoUp, w)
	IndigoSave02 := z
	IndigoPurpleSideRight := ADDpr(CSP+Ip,IndigoSave01, IndigoSave02)

	y, z, w = ASApr(CSP+Ip,p.NFI(105), IndigoPurpleSideRight, p.NFI(7))
	IndigoUp = ADDpr(CSP+Ip,IndigoUp, w)
	IndigoUp = ADDpr(CSP+Ip,IndigoUp, MULpr(CSP+Ip,PurpleUp, p.NFI(6)))
	IndigoTotalUp := ADDpr(CSP+Ip,IndigoBaseUp, IndigoUp)

	IndigoUpHeightPeak := SUBpr(CSP+Ip,IndigoUpHeight, MULpr(CSP+Ip,PurpleUpHeight, p.NFI(6)))
	IndigoUpHeightPeakSeedC := TruncSeed(IndigoUpHeightPeak)
	fmt.Println("IndigoUpHeightPeakSeed is:     ", IndigoUpHeightPeakSeedC)
	//fmt.Println("Purple este",Purple)
	IUA := ADDpr(CSP+Ip,MULpr(CSP+Ip,p.NFI(6), PurpleTotalUp), PRDpr(CSP+Ip,Purple,p.NFI(21),PurpleUpHeight))

	IndigoUpAreaPeak := SUBpr(CSP+Ip,IndigoTotalUp, IUA)
	IndigoUpAreaPeakSeedC := TruncSeed(IndigoUpAreaPeak)
	fmt.Println("IndigoUpAreaPeakSeed is:    ", IndigoUpAreaPeakSeedC)
	//
	//
	//Calculating IndigoDown
	//
	y, z, w = ASApr(CSP+Ip,p.NFI(42), Indigo, p.NFI(90))
	IndigoBaseDown := w
	IndigoSaveDown := z
	IndigoDownHeight := y

	y, z, w = ASApr(CSP+Ip,p.NFI(7), z, p.NFI(97))
	IndigoDown := w
	IndigoSave03 := z

	y, z, w = ASApr(CSP+Ip,p.NFI(104), y, p.NFI(1))
	IndigoDown = ADDpr(CSP+Ip,IndigoDown, w)
	IndigoSave04 := z
	IndigoPurpleSideRight = SUBpr(CSP+Ip,ADDpr(CSP+Ip,IndigoSave04, IndigoSave03), MULpr(CSP+Ip,p.NFI(6), PurpleSaveDown))

	y, z, w = ASApr(CSP+Ip,p.NFI(7), IndigoPurpleSideRight, p.NFI(105))
	IndigoDown = ADDpr(CSP+Ip,IndigoDown, w)
	IndigoDown = ADDpr(CSP+Ip,IndigoDown, MULpr(CSP+Ip,p.NFI(6), PurpleDown))
	IndigoTotalDown := ADDpr(CSP+Ip,IndigoBaseDown, IndigoDown)

	IndigoDownHeightPeak := SUBpr(CSP+Ip,MULpr(CSP+Ip,p.NFI(6), PurpleDownHeight), IndigoDownHeight)
	IndigoDownHeightPeakSeedC := TruncSeed(IndigoDownHeightPeak)
	fmt.Println("IndigoDownHeightPeakSeed is:   ", IndigoDownHeightPeakSeedC)

	IDA1 := MULpr(CSP+Ip,PurpleDownHeight, p.NFI(21))
	IDA2 := SUBpr(CSP+Ip,IDA1, IndigoDownHeightPeak)
	IDA3 := MULpr(CSP+Ip,Purple, IDA2)
	IDA4 := MULpr(CSP+Ip,PurpleTotalDown, p.NFI(6))
	IDA5 := ADDpr(CSP+Ip,IDA3, IDA4)
	IndigoDownAreaPeak := SUBpr(CSP+Ip,IndigoTotalDown, IDA5)
	IndigoDownAreaPeakSeedC := TruncSeed(IndigoDownAreaPeak)

	fmt.Println("IndigoDownAreaPeakSeed is:  ", IndigoDownAreaPeakSeedC)
	fmt.Println("========================================================")
	//
	//
	//Calculating BlueUp
	//
	y,z,w = ASApr(CSP+Bp,p.NFI(90),Blue,p.NFI(50))
	BlueBaseUp := w
	BlueSaveUp := y
	BlueUpHeight := z

	y,z,w = ASApr(CSP+Bp,p.NFI(92),y,p.NFI(3))
	BlueUp := w
	BlueIndigoSideLeft := SUBpr(CSP+Bp,y,MULpr(CSP+Bp,p.NFI(6),IndigoSaveUp))
	BlueSave01 := z

	y,z,w = ASApr(CSP+Bp,p.NFI(95),BlueIndigoSideLeft,p.NFI(2))
	BlueUp = ADDpr(CSP+Bp,BlueUp,w)
	BlueSave02 := z
	BlueSave03 := y
	BlueIndigoSideRight := ADDpr(CSP+Bp,BlueSave01,BlueSave02)

	y,z,w = ASApr(CSP+Bp,p.NFI(97),BlueIndigoSideRight,p.NFI(7))
	BlueUp = ADDpr(CSP+Bp,BlueUp,w)
	BlueSave04 := z
	BlueSave05 := y
	IndigoPurpleSideLeft = SUBpr(CSP+Bp,ADDpr(CSP+Bp,BlueSave03,BlueSave04),MULpr(CSP+Bp,PurpleSaveUp,p.NFI(4)))

	y,z,w = ASApr(CSP+Bp,p.NFI(104),IndigoPurpleSideLeft,p.NFI(1))
	BlueUp = ADDpr(CSP+Bp,BlueUp,w)
	BlueSave06 := z
	IndigoPurpleSideRight = SUBpr(CSP+Bp,ADDpr(CSP+Bp,BlueSave05,BlueSave06),MULpr(CSP+Bp,PurpleSaveDown,p.NFI(2)))

	y,z,w = ASApr(CSP+Bp,p.NFI(105),IndigoPurpleSideRight,p.NFI(7))
	BlueUp = ADDpr(CSP+Bp,BlueUp,w)
	BU1 	:= MULpr(CSP+Bp,p.NFI(6),IndigoUp)
	BU2 	:= MULpr(CSP+Bp,p.NFI(4),PurpleUp)
	BU3 	:= MULpr(CSP+Bp,p.NFI(2),PurpleDown)
	BlueUp = SUMpr(CSP+Bp,BlueUp,BU1,BU2,BU3)
	BlueTotalUp := ADDpr(CSP+Bp,BlueBaseUp,BlueUp)

	BUH11 	:= MULpr(CSP+Bp,p.NFI(6),IndigoUpHeight)
	BUH1 	:= SUBpr(CSP+Bp,BlueUpHeight, BUH11)
	BUH21 	:= MULpr(CSP+Bp,p.NFI(2),PurpleDownHeight)
	BUH22 	:= ADDpr(CSP+Bp,BUH21,BUH1)
	BUH23	:= MULpr(CSP+Bp,p.NFI(4),PurpleUpHeight)
	BUH2 	:= SUBpr(CSP+Bp,BUH23,BUH22)
	BlueUpHeightPeakSeedC := TruncSeed(BUH2)
	fmt.Println("BlueUpHeightPeakSeed is:       ", BlueUpHeightPeakSeedC)


	BUA11 	:= PRDpr(CSP+Bp,Indigo,IndigoUpHeight,p.NFI(21))
	BUA12 	:= MULpr(CSP+Bp,IndigoTotalUp,p.NFI(6))
	BUA13	:= ADDpr(CSP+Bp,BUA11,BUA12)
	BUA1	:= SUBpr(CSP+Bp,BlueTotalUp,BUA13)

	BUA21	:= MULpr(CSP+Bp,BUH1,p.NFI(3))
	BUA22	:= MULpr(CSP+Bp,PurpleDownHeight,p.NFI(3))
	BUA23	:= MULpr(CSP+Bp,PurpleUpHeight,p.NFI(6))
	BUA24	:= SUMpr(CSP+Bp,BUA21,BUA22,BUA23)
	BUA25	:= MULpr(CSP+Bp,Purple,BUA24)
	BUA26	:= MULpr(CSP+Bp,PurpleTotalDown,p.NFI(2))
	BUA27	:= MULpr(CSP+Bp,PurpleTotalUp,p.NFI(4))
	BUA2	:= DIFpr(CSP+Bp,BUA1,BUA25,BUA26,BUA27)
	BlueUpAreaPeakSeedC := TruncSeed(BUA2)
	fmt.Println("BlueUpAreaPeakSeed is:      ", BlueUpAreaPeakSeedC)
	//
	//
	//Calculating BlueDown
	//
	y,z,w = ASApr(CSP+Bp,p.NFI(35),Blue,p.NFI(90))
	BlueBaseDown := w
	BlueSaveDown := z
	BlueDownHeight := y

	y,z,w = ASApr(CSP+Bp,p.NFI(7),z,p.NFI(88))
	BlueSave07 := z
	BlueDown := w

	y,z,w = ASApr(CSP+Bp,p.NFI(95),y,p.NFI(2))
	BlueDown = ADDpr(CSP+Bp,BlueDown,w)
	BlueSave08 := z
	BlueSave09 := y
	BlueIndigoSideRight = SUBpr(CSP+Bp,ADDpr(CSP+Bp,BlueSave07,BlueSave08),MULpr(CSP+Bp,IndigoSaveDown,p.NFI(6)))

	y,z,w = ASApr(CSP+Bp,p.NFI(7),BlueIndigoSideRight,p.NFI(97))
	BlueDown = ADDpr(CSP+Bp,BlueDown,w)
	BlueSave10 := y
	BlueSave11 := z
	IndigoPurpleSideLeft = SUBpr(CSP+Bp,ADDpr(CSP+Bp,BlueSave09,BlueSave10),MULpr(CSP+Bp,PurpleSaveUp,p.NFI(4)))

	y,z,w = ASApr(CSP+Bp,p.NFI(104),IndigoPurpleSideLeft,p.NFI(1))
	BlueDown = ADDpr(CSP+Bp,BlueDown,w)
	BlueSave12 := z
	IndigoPurpleSideRight = SUBpr(CSP+Bp,ADDpr(CSP+Bp,BlueSave11,BlueSave12),MULpr(CSP+Bp,PurpleSaveDown,p.NFI(2)))

	y,z,w = ASApr(CSP+Bp,p.NFI(7),IndigoPurpleSideRight,p.NFI(105))
	BlueDown = ADDpr(CSP+Bp,BlueDown,w)
	BD1		:= MULpr(CSP+Bp,p.NFI(6),IndigoDown)
	BD2 	:= MULpr(CSP+Bp,p.NFI(4),PurpleUp)
	BD3 	:= MULpr(CSP+Bp,p.NFI(2),PurpleDown)
	BlueDown = SUMpr(CSP+Bp,BlueDown,BD1,BD2,BD3)
	BlueTotalDown := ADDpr(CSP+Bp,BlueBaseDown,BlueDown)

	BDH1	:= SUBpr(CSP+Bp,MULpr(CSP+Bp,p.NFI(6),IndigoDownHeight),BlueDownHeight)
	BDH21	:= MULpr(CSP+Bp,p.NFI(2),PurpleDownHeight)
	BDH22	:= MULpr(CSP+Bp,p.NFI(4),PurpleUpHeight)
	BDH23	:= SUBpr(CSP+Bp,BDH21,BDH22)
	BDH2	:= ADDpr(CSP+Bp,BDH1,BDH23)
	BlueDownHeightPeakSeedC := TruncSeed(BDH2)
	fmt.Println("BlueDownHeightPeakSeed is:      ", BlueDownHeightPeakSeedC)
	
	BDA11 	:= MULpr(CSP+Bp,p.NFI(15),IndigoDownHeight)
	BDA12 	:= ADDpr(CSP+Bp,BDA11,BlueDownHeight)
	BDA13 	:= MULpr(CSP+Bp,BDA12,Indigo)
	BDA14 	:= MULpr(CSP+Bp,p.NFI(6),IndigoTotalDown)
	BDA15 	:= ADDpr(CSP+Bp,BDA13,BDA14)
	BDA1 	:= SUBpr(CSP+Bp,BlueTotalDown,BDA15)

	BDA21 := MULpr(CSP+Bp,p.NFI(2),BDH1)
	BDA22 := MULpr(CSP+Bp,p.NFI(10),PurpleUpHeight)
	BDA23 := SUMpr(CSP+Bp,BDA21,BDA22,PurpleDownHeight)
	BDA24 := MULpr(CSP+Bp,Purple,BDA23)
	BDA25 := MULpr(CSP+Bp,p.NFI(2),PurpleTotalDown)
	BDA26 := MULpr(CSP+Bp,p.NFI(4),PurpleTotalUp)
	BDA27 := SUMpr(CSP+Bp,BDA24,BDA25,BDA26)
	BDA2  := SUBpr(CSP+Bp,BDA1,BDA27)
	BlueDownAreaPeakSeedC := TruncSeed(BDA2)
	fmt.Println("BlueDownAreaPeakSeed is:    ", BlueDownAreaPeakSeedC)
	fmt.Println("========================================================")
	//
	//
	//Calculating GreenUp
	//
	y,z,w = ASApr(CSP+Gp,p.NFI(90),Green,p.NFI(46))
	GreenSaveUp := y
	GreenBaseUp := w
	GreenUpHeight := z

	y,z,w = ASApr(CSP+Gp,p.NFI(81),y,p.NFI(4))
	GreenUp := w
	GreenBlueSideLeft := SUBpr(CSP+Gp,y,MULpr(CSP+Gp,p.NFI(6),BlueSaveUp))
	GreenSave01 := z

	y,z,w = ASApr(CSP+Gp,p.NFI(85),GreenBlueSideLeft,p.NFI(3))
	GreenUp = ADDpr(CSP+Gp,GreenUp,w)
	GreenSave02 := z
	GreenSave03 := y
	GreenBlueSideRight := ADDpr(CSP+Gp,GreenSave02,GreenSave01)

	y,z,w = ASApr(CSP+Gp,p.NFI(7),GreenBlueSideRight,p.NFI(88))
	GreenUp = ADDpr(CSP+Gp,GreenUp,w)
	GreenSave04 := y
	GreenSave05 := z
	BlueIndigoSideLeft = SUBpr(CSP+Gp,ADDpr(CSP+Gp,GreenSave03,GreenSave04),MULpr(CSP+Gp,p.NFI(3),IndigoSaveUp))

	y,z,w = ASApr(CSP+Gp,p.NFI(95),BlueIndigoSideLeft,p.NFI(2))
	GreenUp = ADDpr(CSP+Gp,GreenUp,w)
	GreenSave06 := z
	GreenSave07 := y
	BlueIndigoSideRight = SUBpr(CSP+Gp,ADDpr(CSP+Gp,GreenSave05,GreenSave06),MULpr(CSP+Gp,p.NFI(3),IndigoSaveDown))

	y,z,w = ASApr(CSP+Gp,p.NFI(7),BlueIndigoSideRight,p.NFI(97))
	GreenUp = ADDpr(CSP+Gp,GreenUp,w)
	GreenSave08 := y
	GreenSave09 := z
	IndigoPurpleSideLeft = SUBpr(CSP+Gp,ADDpr(CSP+Gp,GreenSave07,GreenSave08),PurpleSaveUp)

	y,z,w = ASApr(CSP+Gp,p.NFI(104),IndigoPurpleSideLeft,p.NFI(1))
	GreenUp = ADDpr(CSP+Gp,GreenUp,w)
	GreenSave10 := z
	IndigoPurpleSideRight = SUBpr(CSP+Gp,ADDpr(CSP+Gp,GreenSave09,GreenSave10),MULpr(CSP+Gp,p.NFI(5),PurpleSaveDown))

	y,z,w = ASApr(CSP+Gp,p.NFI(7),IndigoPurpleSideRight,p.NFI(105))
	GreenUp = ADDpr(CSP+Gp,GreenUp,w)
	GU1		:= MULpr(CSP+Gp,p.NFI(6),BlueUp)
	GU2		:= MULpr(CSP+Gp,p.NFI(3),IndigoUp)
	GU3 	:= MULpr(CSP+Gp,p.NFI(5),PurpleDown)
	GU4 	:= MULpr(CSP+Gp,p.NFI(3),IndigoDown)
	GreenUp = SUMpr(CSP+Gp,GreenUp,PurpleUp,GU1,GU2,GU3,GU4)
	GreenTotalUp := ADDpr(CSP+Gp,GreenBaseUp,GreenUp)

	GUH1  	:= SUBpr(CSP+Gp,GreenUpHeight,MULpr(CSP+Gp,p.NFI(6),BlueUpHeight))
	GUH21 	:= MULpr(CSP+Gp,p.NFI(3),IndigoDownHeight)
	GUH22 	:= ADDpr(CSP+Gp,GUH1,GUH21)
	GUH23 	:= MULpr(CSP+Gp,p.NFI(3),IndigoUpHeight)
	GUH2  	:= SUBpr(CSP+Gp,GUH23,GUH22)
	GUH3  	:= SUBpr(CSP+Gp,MULpr(CSP+Gp,p.NFI(5),PurpleDownHeight),ADDpr(CSP+Gp,GUH2,PurpleUpHeight))
	GreenUpHeightPeakSeedC := TruncSeed(GUH3)
	fmt.Println("GreenUpHeightPeakSeed is:      ", GreenUpHeightPeakSeedC)

	GUA11 	:= PRDpr(CSP+Gp,Blue,p.NFI(21),BlueUpHeight)
	GUA12	:= MULpr(CSP+Gp,p.NFI(6),BlueTotalUp)
	GUA13	:= ADDpr(CSP+Gp,GUA11,GUA12)
	GUA1	:= SUBpr(CSP+Gp,GreenTotalUp,GUA13)

	GUA21 	:= MULpr(CSP+Gp,p.NFI(4),GUH1)
	GUA22 	:= MULpr(CSP+Gp,p.NFI(6),IndigoDownHeight)
	GUA23 	:= MULpr(CSP+Gp,p.NFI(3),IndigoUpHeight)
	GUA24 	:= SUMpr(CSP+Gp,GUA21,GUA22,GUA23)
	GUA25 	:= MULpr(CSP+Gp,Indigo,GUA24)
	GUA26 	:= MULpr(CSP+Gp,p.NFI(3),ADDpr(CSP+Gp,IndigoTotalDown,IndigoTotalUp))
	GUA27 	:= ADDpr(CSP+Gp,GUA26,GUA25)
	GUA2 	:= SUBpr(CSP+Gp,GUA1,GUA27)

	GUA31 	:= MULpr(CSP+Gp,p.NFI(2),GUH2)
	GUA32 	:= MULpr(CSP+Gp,p.NFI(10),PurpleDownHeight)
	GUA33 	:= SUMpr(CSP+Gp,GUA31,GUA32,PurpleUpHeight)
	GUA34 	:= MULpr(CSP+Gp,Purple,GUA33)
	GUA35 	:= MULpr(CSP+Gp,p.NFI(5),PurpleTotalDown)
	GUA36 	:= SUMpr(CSP+Gp,PurpleTotalUp,GUA34,GUA35)
	GUA3 	:= SUBpr(CSP+Gp,GUA2,GUA36)
	GreenUpAreaPeakSeedC := TruncSeed(GUA3)
	fmt.Println("GreenUpAreaPeakSeed is:     ", GreenUpAreaPeakSeedC)
	//
	//
	//Calculating GreenDown
	//
	y,z,w = ASApr(CSP+Gp,p.NFI(28),Green,p.NFI(90))
	GreenSaveDown := z
	GreenBaseDown := w
	GreenDownHeight := y

	y,z,w = ASApr(CSP+Gp,p.NFI(7),z,p.NFI(78))
	GreenDown := w
	GreenSave11 := z

	y,z,w = ASApr(CSP+Gp,p.NFI(85),y,p.NFI(3))
	GreenDown = ADDpr(CSP+Gp,GreenDown,w)
	GreenSave12 := z
	GreenSave13 := y
	GreenBlueSideRight = SUBpr(CSP+Gp,ADDpr(CSP+Gp,GreenSave11,GreenSave12),MULpr(CSP+Gp,p.NFI(6),BlueSaveDown))

	y,z,w = ASApr(CSP+Gp,p.NFI(7),GreenBlueSideRight,p.NFI(88))
	GreenDown = ADDpr(CSP+Gp,GreenDown,w)
	GreenSave14 := y
	GreenSave15 := z
	BlueIndigoSideLeft = SUBpr(CSP+Gp,ADDpr(CSP+Gp,GreenSave13,GreenSave14),MULpr(CSP+Gp,p.NFI(4),IndigoSaveUp))

	y,z,w = ASApr(CSP+Gp,p.NFI(95),BlueIndigoSideLeft,p.NFI(2))
	GreenDown = ADDpr(CSP+Gp,GreenDown,w)
	GreenSave16 := z
	GreenSave17 := y
	BlueIndigoSideRight = SUBpr(CSP+Gp,ADDpr(CSP+Gp,GreenSave15,GreenSave16),MULpr(CSP+Gp,p.NFI(2),IndigoSaveDown))

	y,z,w = ASApr(CSP+Gp,p.NFI(7),BlueIndigoSideRight,p.NFI(97))
	GreenDown = ADDpr(CSP+Gp,GreenDown,w)
	GreenSave18 := y
	GreenSave19 := z
	IndigoPurpleSideLeft = SUBpr(CSP+Gp,ADDpr(CSP+Gp,GreenSave17,GreenSave18),MULpr(CSP+Gp,p.NFI(2),PurpleSaveUp))

	y,z,w = ASApr(CSP+Gp,p.NFI(104),IndigoPurpleSideLeft,p.NFI(1))
	GreenDown = ADDpr(CSP+Gp,GreenDown,w)
	GreenSave20 := z
	IndigoPurpleSideRight = SUBpr(CSP+Gp,ADDpr(CSP+Gp,GreenSave19,GreenSave20),MULpr(CSP+Gp,p.NFI(4),PurpleSaveDown))

	y,z,w = ASApr(CSP+Gp,p.NFI(7),IndigoPurpleSideRight,p.NFI(105))
	GreenDown = ADDpr(CSP+Gp,GreenDown,w)
	GD1 	:= MULpr(CSP+Gp,p.NFI(4),IndigoUp)
	GD2 	:= MULpr(CSP+Gp,p.NFI(2),PurpleUp)
	GD3 	:= MULpr(CSP+Gp,p.NFI(4),PurpleDown)
	GD4 	:= MULpr(CSP+Gp,p.NFI(2),IndigoDown)
	GD5 	:= MULpr(CSP+Gp,p.NFI(6),BlueDown)
	GreenDown = SUMpr(CSP+Gp,GreenDown,GD1,GD2,GD3,GD4,GD5)
	GreenTotalDown := ADDpr(CSP+Gp,GreenBaseDown,GreenDown)

	GDH1 	:= SUBpr(CSP+Gp,MULpr(CSP+Gp,p.NFI(6),BlueDownHeight),GreenDownHeight)
	GDH21 	:= ADDpr(CSP+Gp,MULpr(CSP+Gp,p.NFI(2),IndigoDownHeight),GDH1)
	GDH2 	:= SUBpr(CSP+Gp,MULpr(CSP+Gp,p.NFI(4),IndigoUpHeight),GDH21)
	GDH31 	:= ADDpr(CSP+Gp,MULpr(CSP+Gp,p.NFI(2),PurpleUpHeight),GDH2)
	GreenDownH3 := SUBpr(CSP+Gp,MULpr(CSP+Gp,p.NFI(4),PurpleDownHeight),GDH31)
	GreenDownHeightPeakSeedC := TruncSeed(GreenDownH3)
	fmt.Println("GreenDownHeightPeakSeed is:    ", GreenDownHeightPeakSeedC)

	GDA11 	:= ADDpr(CSP+Gp,MULpr(CSP+Gp,p.NFI(15),BlueDownHeight),GreenDownHeight)
	GDA12 	:= MULpr(CSP+Gp,Blue,GDA11)
	GDA13 	:= MULpr(CSP+Gp,BlueTotalDown,p.NFI(6))
	GDA14 	:= ADDpr(CSP+Gp,GDA12,GDA13)
	GA1 	:= SUBpr(CSP+Gp,GreenTotalDown,GDA14)

	GDA21 	:= MULpr(CSP+Gp,GDH1,p.NFI(3))
	GDA22 	:= MULpr(CSP+Gp,IndigoDownHeight,p.NFI(3))
	GDA23 	:= MULpr(CSP+Gp,IndigoUpHeight,p.NFI(6))
	GDA24 	:= SUMpr(CSP+Gp,GDA21,GDA22,GDA23)
	GDA25 	:= MULpr(CSP+Gp,Indigo,GDA24)
	GDA26 	:= MULpr(CSP+Gp,IndigoTotalDown,p.NFI(2))
	GDA27 	:= MULpr(CSP+Gp,IndigoTotalUp,p.NFI(4))
	GDA28 	:= SUMpr(CSP+Gp,GDA25,GDA26,GDA27)
	GA2 	:= SUBpr(CSP+Gp,GA1,GDA28)

	GDA31 	:= MULpr(CSP+Gp,GDH2,p.NFI(3))
	GDA32 	:= MULpr(CSP+Gp,PurpleUpHeight,p.NFI(3))
	GDA33 	:= MULpr(CSP+Gp,PurpleDownHeight,p.NFI(6))
	GDA34 	:= SUMpr(CSP+Gp,GDA31,GDA32,GDA33)
	GDA35 	:= MULpr(CSP+Gp,Purple,GDA34)
	GDA36 	:= MULpr(CSP+Gp,PurpleTotalDown,p.NFI(4))
	GDA37 	:= MULpr(CSP+Gp,PurpleTotalUp,p.NFI(2))
	GDA38 	:= SUMpr(CSP+Gp,GDA35,GDA36,GDA37)
	GDA3 	:= SUBpr(CSP+Gp,GA2,GDA38)
	GreenDownAreaPeakSeedC := TruncSeed(GDA3)

	fmt.Println("GreenDownAreaPeakSeed is:   ", GreenDownAreaPeakSeedC)
	fmt.Println("========================================================")
	//
	//
	//Calculating YellowUp
	//
	y,z,w = ASApr(CSP+YUp,p.NFI(90),Yellow,p.NFI(41))
	YellowSaveUp := y
	YellowBaseUp := w
	YellowUpHeight := z

	y,z,w = ASApr(CSP+YUp,p.NFI(69),y,p.NFI(5))
	YellowUp := w
	YellowGreenSideLeft := SUBpr(CSP+YUp,y,MULpr(CSP+YUp,p.NFI(6),GreenSaveUp))
	YellowSave01 := z

	y,z,w = ASApr(CSP+YUp,p.NFI(74),YellowGreenSideLeft,p.NFI(4))
	YellowUp = ADDpr(CSP+YUp,YellowUp,w)
	YellowSave02 := z
	YellowSave03 := y
	YellowGreenSideRight := ADDpr(CSP+YUp,YellowSave02,YellowSave01)

	y,z,w = ASApr(CSP+YUp,p.NFI(7),YellowGreenSideRight,p.NFI(78))
	YellowUp = ADDpr(CSP+YUp,YellowUp,w)
	YellowSave04 := y
	YellowSave05 := z
	GreenBlueSideLeft = SUBpr(CSP+YUp,ADDpr(CSP+YUp,YellowSave03,YellowSave04),MULpr(CSP+YUp,BlueSaveUp,p.NFI(2)))

	y,z,w = ASApr(CSP+YUp,p.NFI(85),GreenBlueSideLeft,p.NFI(3))
	YellowUp = ADDpr(CSP+YUp,YellowUp,w)
	YellowSave06 := z
	YellowSave07 := y
	GreenBlueSideRight = SUBpr(CSP+YUp,ADDpr(CSP+YUp,YellowSave05,YellowSave06),MULpr(CSP+YUp,BlueSaveDown,p.NFI(4)))

	y,z,w = ASApr(CSP+YUp,p.NFI(7),GreenBlueSideRight,p.NFI(88))
	YellowUp = ADDpr(CSP+YUp,YellowUp,w)
	YellowSave08 := y
	YellowSave09 := z
	BlueIndigoSideLeft = SUBpr(CSP+YUp,ADDpr(CSP+YUp,YellowSave07,YellowSave08),IndigoSaveUp)

	y,z,w = ASApr(CSP+YUp,p.NFI(95),BlueIndigoSideLeft,p.NFI(2))
	YellowUp = ADDpr(CSP+YUp,YellowUp,w)
	YellowSave10 := z
	YellowSave11 := y
	BlueIndigoSideRight = SUBpr(CSP+YUp,ADDpr(CSP+YUp,YellowSave09,YellowSave10),MULpr(CSP+YUp,IndigoSaveDown,p.NFI(5)))

	y,z,w = ASApr(CSP+YUp,p.NFI(7),BlueIndigoSideRight,p.NFI(97))
	YellowUp = ADDpr(CSP+YUp,YellowUp,w)
	YellowSave12 := y
	YellowSave13 := z
	IndigoPurpleSideLeft = SUBpr(CSP+YUp,ADDpr(CSP+YUp,YellowSave11,YellowSave12),MULpr(CSP+YUp,PurpleSaveUp,p.NFI(2)))

	y,z,w = ASApr(CSP+YUp,p.NFI(104),IndigoPurpleSideLeft,p.NFI(1))
	YellowUp = ADDpr(CSP+YUp,YellowUp,w)
	YellowSave14 := z
	IndigoPurpleSideRight = SUBpr(CSP+YUp,ADDpr(CSP+YUp,YellowSave13,YellowSave14),MULpr(CSP+YUp,PurpleSaveDown,p.NFI(4)))

	y,z,w = ASApr(CSP+YUp,p.NFI(7),IndigoPurpleSideRight,p.NFI(105))
	YellowUp = ADDpr(CSP+YUp,YellowUp,w)
	YU1 	:= MULpr(CSP+YUp,GreenUp,p.NFI(6))
	YU2 	:= MULpr(CSP+YUp,BlueUp,p.NFI(2))
	YU3 	:= MULpr(CSP+YUp,PurpleUp,p.NFI(2))
	YU4 	:= MULpr(CSP+YUp,PurpleDown,p.NFI(4))
	YU5 	:= MULpr(CSP+YUp,IndigoDown,p.NFI(5))
	YU6 	:= MULpr(CSP+YUp,BlueDown,p.NFI(4))
	YellowUp = SUMpr(CSP+YUp,YellowUp,IndigoUp,YU1,YU2,YU3,YU4,YU5,YU6)
	YellowTotalUp := ADDpr(CSP+YUp,YellowBaseUp,YellowUp)

	YUH1 	:= SUBpr(CSP+YUp,MULpr(CSP+YUp,GreenUpHeight,p.NFI(6)),YellowUpHeight)
	YUH2 	:= SUBpr(CSP+YUp,ADDpr(CSP+YUp,MULpr(CSP+YUp,BlueUpHeight,p.NFI(2)),YUH1),MULpr(CSP+YUp,BlueDownHeight,p.NFI(4)))
	YUH3 	:= SUBpr(CSP+YUp,ADDpr(CSP+YUp,YUH2,IndigoUpHeight),MULpr(CSP+YUp,IndigoDownHeight,p.NFI(5)))
	YUH4 	:= SUBpr(CSP+YUp,MULpr(CSP+YUp,PurpleDownHeight,p.NFI(4)),ADDpr(CSP+YUp,YUH3,MULpr(CSP+YUp,PurpleUpHeight,p.NFI(2))))
	YellowUpHeightPeakSeedC := TruncSeed(YUH4)
	fmt.Println("YellowUpHeightPeakSeed is:     ", YellowUpHeightPeakSeedC)

	YUA11 	:= ADDpr(CSP+YUp,MULpr(CSP+YUp,GreenUpHeight,p.NFI(15)),YellowUpHeight)
	YUA12 	:= ADDpr(CSP+YUp,MULpr(CSP+YUp,GreenTotalUp,p.NFI(6)),MULpr(CSP+YUp,YUA11,Green))
	YUA1 	:= SUBpr(CSP+YUp,YellowTotalUp,YUA12)

	YUA21 	:= SUMpr(CSP+YUp,BlueUpHeight,MULpr(CSP+YUp,BlueDownHeight,p.NFI(10)),MULpr(CSP+YUp,YUH1,p.NFI(2)))
	YUA22 	:= SUMpr(CSP+YUp,MULpr(CSP+YUp,Blue,YUA21),MULpr(CSP+YUp,BlueTotalDown,p.NFI(4)),MULpr(CSP+YUp,BlueTotalUp,p.NFI(2)))
	YUA2 	:= SUBpr(CSP+YUp,YUA1,YUA22)

	YUA31 	:= ADDpr(CSP+YUp,MULpr(CSP+YUp,IndigoDownHeight,p.NFI(15)),YUH2)
	YUA32 	:= SUMpr(CSP+YUp,MULpr(CSP+YUp,Indigo,YUA31),MULpr(CSP+YUp,IndigoTotalDown,p.NFI(5)),IndigoTotalUp)
	YUA3 	:= SUBpr(CSP+YUp,YUA2,YUA32)

	YUA41 	:= SUMpr(CSP+YUp,MULpr(CSP+YUp,PurpleDownHeight,p.NFI(6)),MULpr(CSP+YUp,PurpleUpHeight,p.NFI(3)),MULpr(CSP+YUp,YUH3,p.NFI(3)))
	YUA42 	:= SUMpr(CSP+YUp,MULpr(CSP+YUp,Purple,YUA41),MULpr(CSP+YUp,PurpleTotalDown,p.NFI(4)),MULpr(CSP+YUp,PurpleTotalUp,p.NFI(2)))
	YUA4 	:= SUBpr(CSP+YUp,YUA3,YUA42)
	YellowUpAreaPeakSeedC := TruncSeed(YUA4)
	fmt.Println("YellowUpAreaPeakSeed is:    ", YellowUpAreaPeakSeedC)
	//
	//
	//Calculating YellowDown
	//
	y,z,w = ASApr(CSP+YDp,p.NFI(21),Yellow,p.NFI(90))
	YellowBaseDown := w
	YellowSaveDown := z
	YellowDownHeight := y

	y,z,w = ASApr(CSP+YDp,p.NFI(7),z,p.NFI(67))
	YellowDown := w
	YellowSave15 := z

	y,z,w = ASApr(CSP+YDp,p.NFI(74),y,p.NFI(4))
	YellowDown = ADDpr(CSP+YDp,YellowDown,w)
	YellowSave16 := z
	YellowSave17 := y
	YellowGreenSideRight = SUBpr(CSP+YDp,ADDpr(CSP+YDp,YellowSave15,YellowSave16),MULpr(CSP+YDp,GreenSaveDown,p.NFI(6)))

	y,z,w = ASApr(CSP+YDp,p.NFI(7),YellowGreenSideRight,p.NFI(78))
	YellowDown = ADDpr(CSP+YDp,YellowDown,w)
	YellowSave18 := y
	YellowSave19 := z
	GreenBlueSideLeft = SUBpr(CSP+YDp,ADDpr(CSP+YDp,YellowSave17,YellowSave18),MULpr(CSP+YDp,BlueSaveUp,p.NFI(4)))

	y,z,w = ASApr(CSP+YDp,p.NFI(85),GreenBlueSideLeft,p.NFI(3))
	YellowDown = ADDpr(CSP+YDp,YellowDown,w)
	YellowSave20 := z
	YellowSave21 := y
	GreenBlueSideRight = SUBpr(CSP+YDp,ADDpr(CSP+YDp,YellowSave19,YellowSave20),MULpr(CSP+YDp,BlueSaveDown,p.NFI(2)))

	y,z,w = ASApr(CSP+YDp,p.NFI(7),GreenBlueSideRight,p.NFI(88))
	YellowDown = ADDpr(CSP+YDp,YellowDown,w)
	YellowSave22 := y
	YellowSave23 := z
	BlueIndigoSideLeft = SUBpr(CSP+YDp,ADDpr(CSP+YDp,YellowSave21,YellowSave22),MULpr(CSP+YDp,IndigoSaveUp,p.NFI(3)))

	y,z,w = ASApr(CSP+YDp,p.NFI(95),BlueIndigoSideLeft,p.NFI(2))
	YellowDown = ADDpr(CSP+YDp,YellowDown,w)
	YellowSave24 := z
	YellowSave25 := y
	BlueIndigoSideRight = SUBpr(CSP+YDp,ADDpr(CSP+YDp,YellowSave23,YellowSave24),MULpr(CSP+YDp,IndigoSaveDown,p.NFI(3)))

	y,z,w = ASApr(CSP+YDp,p.NFI(7),BlueIndigoSideRight,p.NFI(97))
	YellowDown = ADDpr(CSP+YUp,YellowDown,w)
	YellowSave26 := y
	YellowSave27 := z
	IndigoPurpleSideLeft = SUBpr(CSP+YDp,ADDpr(CSP+YDp,YellowSave25,YellowSave26),MULpr(CSP+YDp,PurpleSaveUp,p.NFI(2)))

	y,z,w = ASApr(CSP+YDp,p.NFI(104),IndigoPurpleSideLeft,p.NFI(1))
	YellowDown = ADDpr(CSP+YDp,YellowDown,w)
	YellowSave28 := z
	IndigoPurpleSideRight = SUBpr(CSP+YDp,ADDpr(CSP+YDp,YellowSave27,YellowSave28),MULpr(CSP+YDp,PurpleSaveDown,p.NFI(4)))

	y,z,w = ASApr(CSP+YDp,p.NFI(7),IndigoPurpleSideRight,p.NFI(105))
	YellowDown = ADDpr(CSP+YDp,YellowDown,w)
	YD1 	:= MULpr(CSP+YDp,BlueUp,p.NFI(4))
	YD2 	:= MULpr(CSP+YDp,IndigoUp,p.NFI(3))
	YD3 	:= MULpr(CSP+YDp,PurpleUp,p.NFI(2))
	YD4 	:= MULpr(CSP+YDp,PurpleDown,p.NFI(4))
	YD5 	:= MULpr(CSP+YDp,IndigoDown,p.NFI(3))
	YD6 	:= MULpr(CSP+YDp,BlueDown,p.NFI(2))
	YD7 	:= MULpr(CSP+YDp,GreenDown,p.NFI(6))
	YellowDown = SUMpr(CSP+YDp,YellowDown,YD1,YD2,YD3,YD4,YD5,YD6,YD7)
	YellowTotalDown := ADDpr(CSP+YDp,YellowBaseDown,YellowDown)

	YDH1 	:= SUBpr(CSP+YDp,YellowDownHeight,MULpr(CSP+YDp,GreenDownHeight,p.NFI(5)))
	YDH21 	:= ADDpr(CSP+YDp,GreenDownHeight,MULpr(CSP+YDp,BlueDownHeight,p.NFI(2)))
	YDH22 	:= ADDpr(CSP+YDp,YDH1,MULpr(CSP+YDp,BlueUpHeight,p.NFI(4)))
	YDH2 	:= SUBpr(CSP+YDp,YDH21,YDH22)

	YDH31 	:= ADDpr(CSP+YDp,YDH2,MULpr(CSP+YDp,IndigoDownHeight,p.NFI(3)))
	YDH3 	:= SUBpr(CSP+YDp,MULpr(CSP+YDp,IndigoUpHeight,p.NFI(3)),YDH31)

	YDH41 	:= ADDpr(CSP+YDp,YDH3,MULpr(CSP+YDp,PurpleUpHeight,p.NFI(2)))
	YDH4 	:= SUBpr(CSP+YDp,MULpr(CSP+YDp,PurpleDownHeight,p.NFI(4)),YDH41)
	YellowDownHeightPeakSeedC := TruncSeed(YDH4)
	fmt.Println("YellowDownHeightPeakSeed is:   ", YellowDownHeightPeakSeedC)

	YellowDownArea101 := MULpr(CSP+YDp,MULpr(CSP+YDp,Green,GreenDownHeight),p.NFI(20))
	YellowDownArea102 := ADDpr(CSP+YDp,MULpr(CSP+YDp,GreenTotalDown,p.NFI(6)),YellowDownArea101)
	YellowDownArea1 := SUBpr(CSP+YDp,YellowTotalDown,YellowDownArea102)

	YDA21 	:= MULpr(CSP+YDp,YDH1,p.NFI(5))
	YDA22 	:= MULpr(CSP+YDp,GreenDownHeight,p.NFI(2))
	YDA23 	:= MULpr(CSP+YDp,BlueUpHeight,p.NFI(10))
	YDA24 	:= SUMpr(CSP+YDp,BlueDownHeight,YDA21,YDA22,YDA23)
	YDA25 	:= MULpr(CSP+YDp,Blue,YDA24)
	YDA26 	:= MULpr(CSP+YDp,BlueTotalDown,p.NFI(2))
	YDA27 	:= MULpr(CSP+YDp,BlueTotalUp,p.NFI(4))
	YDA28 	:= SUMpr(CSP+YDp,YDA25,YDA26,YDA27)
	YDA2 	:= SUBpr(CSP+YDp,YellowDownArea1,YDA28)

	YDA31 	:= MULpr(CSP+YDp,YDH2,p.NFI(4))
	YDA32 	:= MULpr(CSP+YDp,IndigoDownHeight,p.NFI(6))
	YDA33 	:= MULpr(CSP+YDp,IndigoUpHeight,p.NFI(3))
	YDA34 	:= SUMpr(CSP+YDp,YDA31,YDA32,YDA33)
	YDA35 	:= MULpr(CSP+YDp,Indigo,YDA34)
	YDA36 	:= MULpr(CSP+YDp,ADDpr(CSP+YDp,IndigoTotalUp,IndigoTotalDown),p.NFI(3))
	YDA3 	:= SUBpr(CSP+YDp,YDA2,ADDpr(CSP+YDp,YDA35,YDA36))

	YDA41 	:= MULpr(CSP+YDp,PurpleDownHeight,p.NFI(6))
	YDA42 	:= MULpr(CSP+YDp,YDH3,p.NFI(3))
	YDA43 	:= MULpr(CSP+YDp,PurpleUpHeight,p.NFI(3))
	YDA44 	:= SUMpr(CSP+YDp,YDA41,YDA42,YDA43)
	YDA45 	:= MULpr(CSP+YDp,Purple,YDA44)
	YDA46 	:= MULpr(CSP+YDp,PurpleTotalDown,p.NFI(4))
	YDA47 	:= MULpr(CSP+YDp,PurpleTotalUp,p.NFI(2))
	YDA48 	:= SUMpr(CSP+YDp,YDA45,YDA46,YDA47)
	YDA4 	:= SUBpr(CSP+YDp,YDA3,YDA48)
	YellowDownAreaPeakSeedC := TruncSeed(YDA4)
	fmt.Println("YellowDownAreaPeakSeed is:  ", YellowDownAreaPeakSeedC)
	fmt.Println("========================================================")
	//
	//
	//Calculating OrangeUp
	//
	y,z,w = ASApr(CSP+Op,p.NFI(90),Orange,p.NFI(35))
	OrangeSaveUp := y
	OrangeBaseUp := w
	OrangeUpHeight := z

	y,z,w = ASApr(CSP+Op,p.NFI(56),y,p.NFI(6))
	OrangeUp := w
	OrangeYellowSideLeft := SUBpr(CSP+Op,y,MULpr(CSP+Op,YellowSaveUp,p.NFI(6)))
	OrangeSave01 := z

	y,z,w = ASApr(CSP+Op,p.NFI(62),OrangeYellowSideLeft,p.NFI(5))
	OrangeUp = ADDpr(CSP+Op,OrangeUp,w)
	OrangeSave02 := z
	OrangeSave03 := y
	OrangeYellowSideRight := ADDpr(CSP+Op,OrangeSave01,OrangeSave02)

	y,z,w = ASApr(CSP+Op,p.NFI(7),OrangeYellowSideRight,p.NFI(67))
	OrangeUp = ADDpr(CSP+Op,OrangeUp,w)
	OrangeSave04 := z
	OrangeSave05 := y
	YellowGreenSideRight = SUBpr(CSP+Op,OrangeSave04,MULpr(CSP+Op,p.NFI(6),GreenSaveDown))

	y,z,w = ASApr(CSP+Op,p.NFI(7),YellowGreenSideRight,p.NFI(74))
	OrangeUp = ADDpr(CSP+Op,OrangeUp,w)
	OrangeSave06 := y
	OrangeSave07 := z
	YellowGreenSideLeft = SUMpr(CSP+Op,OrangeSave03,OrangeSave05,OrangeSave06)

	y,z,w = ASApr(CSP+Op,p.NFI(81),YellowGreenSideLeft,p.NFI(4))
	OrangeUp = ADDpr(CSP+Op,OrangeUp,w)
	OrangeSave08 := y
	OrangeSave09 := z
	GreenBlueSideLeft = SUBpr(CSP+Op,OrangeSave08,MULpr(CSP+Op,p.NFI(6),BlueSaveUp))

	y,z,w = ASApr(CSP+Op,p.NFI(85),GreenBlueSideLeft,p.NFI(3))
	OrangeUp = ADDpr(CSP+Op,OrangeUp,w)
	OrangeSave10 := z
	OrangeSave11 := y
	GreenBlueSideRight =  SUMpr(CSP+Op,OrangeSave07,OrangeSave09,OrangeSave10)

	y,z,w = ASApr(CSP+Op,p.NFI(7),GreenBlueSideRight,p.NFI(88))
	OrangeUp = ADDpr(CSP+Op,OrangeUp,w)
	OrangeSave12 := y
	OrangeSave13 := z
	BlueIndigoSideLeft = SUBpr(CSP+Op,ADDpr(CSP+Op,OrangeSave11,OrangeSave12),MULpr(CSP+Op,IndigoSaveUp,p.NFI(2)))

	y,z,w = ASApr(CSP+Op,p.NFI(95),BlueIndigoSideLeft,p.NFI(2))
	OrangeUp = ADDpr(CSP+Op,OrangeUp,w)
	OrangeSave14 := z
	OrangeSave15 := y
	BlueIndigoSideRight = SUBpr(CSP+Op,ADDpr(CSP+Op,OrangeSave13,OrangeSave14),MULpr(CSP+Op,IndigoSaveDown,p.NFI(4)))

	y,z,w = ASApr(CSP+Op,p.NFI(7),BlueIndigoSideRight,p.NFI(97))
	OrangeUp = ADDpr(CSP+Op,OrangeUp,w)
	OrangeSave16 := y
	OrangeSave17 := z
	IndigoPurpleSideLeft = SUBpr(CSP+Op,ADDpr(CSP+Op,OrangeSave15,OrangeSave16),PurpleSaveUp)

	y,z,w = ASApr(CSP+Op,p.NFI(104),IndigoPurpleSideLeft,p.NFI(1))
	OrangeUp = ADDpr(CSP+Op,OrangeUp,w)
	OrangeSave18 := z
	IndigoPurpleSideRight = SUBpr(CSP+Op,ADDpr(CSP+Op,OrangeSave17,OrangeSave18),MULpr(CSP+Op,PurpleSaveDown,p.NFI(5)))

	y,z,w = ASApr(CSP+Op,p.NFI(7),IndigoPurpleSideRight,p.NFI(105))
	OrangeUp = ADDpr(CSP+Op,OrangeUp,w)
	OU1 	:= MULpr(CSP+Op,p.NFI(6),YellowUp)
	OU2 	:= MULpr(CSP+Op,p.NFI(6),BlueUp)
	OU3 	:= MULpr(CSP+Op,p.NFI(2),IndigoUp)
	OU4 	:= MULpr(CSP+Op,p.NFI(5),PurpleDown)
	OU5 	:= MULpr(CSP+Op,p.NFI(4),IndigoDown)
	OU6 	:= MULpr(CSP+Op,p.NFI(6),GreenDown)
	OrangeUp = SUMpr(CSP+Op,OrangeUp,PurpleUp,OU1,OU2,OU3,OU4,OU5,OU6)
	OrangeTotalUp := ADDpr(CSP+Op,OrangeBaseUp,OrangeUp)


	OUH1 	:= SUBpr(CSP+Op,MULpr(CSP+Op,p.NFI(6),YellowUpHeight),OrangeUpHeight)
	OUH2 	:= SUBpr(CSP+Op,MULpr(CSP+Op,p.NFI(6),GreenDownHeight),OUH1)
	OUH3 	:= SUBpr(CSP+Op,MULpr(CSP+Op,p.NFI(6),BlueUpHeight),OUH2)
	OUH4 	:= SUBpr(CSP+Op,ADDpr(CSP+Op,OUH3,MULpr(CSP+Op,p.NFI(2),IndigoUpHeight)),MULpr(CSP+Op,p.NFI(4),IndigoDownHeight))
	OUH5 	:= DIFpr(CSP+Op,MULpr(CSP+Op,p.NFI(5),PurpleDownHeight),OUH4,PurpleUpHeight)
	OrangeUpHeightPeakSeedC := TruncSeed(OUH5)
	fmt.Println("OrangeUpHeightPeakSeed is:     ", OrangeUpHeightPeakSeedC)

	OUA11 	:= ADDpr(CSP+Op,MULpr(CSP+Op,p.NFI(15),YellowUpHeight),OrangeUpHeight)
	OUA12	:= ADDpr(CSP+Op,MULpr(CSP+Op,p.NFI(6),YellowTotalUp),MULpr(CSP+Op,Yellow,OUA11))
	OUA1	:= SUBpr(CSP+Op,OrangeTotalUp,OUA12)

	OUA21	:= ADDpr(CSP+Op,MULpr(CSP+Op,p.NFI(15),GreenDownHeight),OUH1)
	OUA22	:= ADDpr(CSP+Op,MULpr(CSP+Op,p.NFI(6),GreenTotalDown),MULpr(CSP+Op,Green,OUA21))
	OUA2	:= SUBpr(CSP+Op,OUA1,OUA22)

	OUA31	:= ADDpr(CSP+Op,MULpr(CSP+Op,p.NFI(15),BlueUpHeight),OUH2)
	OUA32	:= ADDpr(CSP+Op,MULpr(CSP+Op,p.NFI(6),BlueTotalUp),MULpr(CSP+Op,Blue,OUA31))
	OUA3	:= SUBpr(CSP+Op,OUA2,OUA32)

	OUA41 	:= MULpr(CSP+Op,p.NFI(10),IndigoDownHeight)
	OUA42 	:= MULpr(CSP+Op,p.NFI(2),OUH3)
	OUA43	:= SUMpr(CSP+Op,IndigoUpHeight,OUA41,OUA42)
	OUA44	:= MULpr(CSP+Op,Indigo,OUA43)
	OUA45 	:= MULpr(CSP+Op,p.NFI(4),IndigoTotalDown)
	OUA46 	:= MULpr(CSP+Op,p.NFI(2),IndigoTotalUp)
	OUA47	:= SUMpr(CSP+Op,OUA44,OUA45,OUA46)
	OUA4	:= SUBpr(CSP+Op,OUA3,OUA47)

	OUA51 	:= MULpr(CSP+Op,p.NFI(10),PurpleDownHeight)
	OUA52 	:= MULpr(CSP+Op,p.NFI(2),OUH4)
	OUA53	:= SUMpr(CSP+Op,PurpleUpHeight,OUA51,OUA52)
	OUA54	:= MULpr(CSP+Op,Purple,OUA53)
	OUA55 	:= MULpr(CSP+Op,p.NFI(5),PurpleTotalDown)
	OUA56	:= SUMpr(CSP+Op,PurpleTotalUp,OUA55,OUA54)
	OUA5	:= SUBpr(CSP+Op,OUA4,OUA56)
	OrangeUpAreaPeakSeedC := TruncSeed(OUA5)
	fmt.Println("OrangeUpAreaPeakSeed is:    ", OrangeUpAreaPeakSeedC)
	//
	//
	//Calculating OrangeDown
	//
	y,z,w = ASApr(CSP+Op,p.NFI(14),Orange,p.NFI(90))
	OrangeBaseDown := w
	OrangeSaveDown := z
	OrangeDownHeight := y

	y,z,w = ASApr(CSP+Op,p.NFI(7),z,p.NFI(55))
	OrangeDown := w
	OrangeSave19 := z

	y,z,w = ASApr(CSP+Op,p.NFI(62),y,p.NFI(5))
	OrangeDown = ADDpr(CSP+Op,OrangeDown,w)
	OrangeSave20 := z
	OrangeSave21 := y
	OrangeYellowSideRight = SUBpr(CSP+Op,ADDpr(CSP+Op,OrangeSave19,OrangeSave20),MULpr(CSP+Op,YellowSaveDown,p.NFI(6)))

	y,z,w = ASApr(CSP+Op,p.NFI(7),OrangeYellowSideRight,p.NFI(67))
	OrangeDown = ADDpr(CSP+Op,OrangeDown,w)
	OrangeSave22 := y
	OrangeSave23 := z
	YellowGreenSideRight = SUBpr(CSP+Op,ADDpr(CSP+Op,OrangeSave21,OrangeSave22),MULpr(CSP+Op,GreenSaveUp,p.NFI(4)))

	y,z,w = ASApr(CSP+Op,p.NFI(74),YellowGreenSideRight,p.NFI(4))
	OrangeDown = ADDpr(CSP+Op,OrangeDown,w)
	OrangeSave24 := z
	OrangeSave25 := y
	YellowGreenSideLeft = SUBpr(CSP+Op,ADDpr(CSP+Op,OrangeSave23,OrangeSave24),MULpr(CSP+Op,GreenSaveDown,p.NFI(2)))

	y,z,w = ASApr(CSP+Op,p.NFI(7),YellowGreenSideLeft,p.NFI(78))
	OrangeDown = ADDpr(CSP+Op,OrangeDown,w)
	OrangeSave26 := y
	OrangeSave27 := z
	GreenBlueSideLeft = SUBpr(CSP+Op,ADDpr(CSP+Op,OrangeSave25,OrangeSave26),MULpr(CSP+Op,BlueSaveUp,p.NFI(5)))

	y,z,w = ASApr(CSP+Op,p.NFI(85),GreenBlueSideLeft,p.NFI(3))
	OrangeDown = ADDpr(CSP+Op,OrangeDown,w)
	OrangeSave28 := z
	OrangeSave29 := y
	GreenBlueSideRight = SUBpr(CSP+Op,ADDpr(CSP+Op,OrangeSave27,OrangeSave28),BlueSaveDown)

	y,z,w = ASApr(CSP+Op,p.NFI(7),GreenBlueSideRight,p.NFI(88))
	OrangeDown = ADDpr(CSP+Op,OrangeDown,w)
	OrangeSave30 := y
	OrangeSave31 := z
	BlueIndigoSideLeft = SUBpr(CSP+Op,ADDpr(CSP+Op,OrangeSave29,OrangeSave30),MULpr(CSP+Op,IndigoSaveUp,p.NFI(4)))

	y,z,w = ASApr(CSP+Op,p.NFI(95),BlueIndigoSideLeft,p.NFI(2))
	OrangeDown = ADDpr(CSP+Op,OrangeDown,w)
	OrangeSave32 := z
	OrangeSave33 := y
	BlueIndigoSideRight = SUBpr(CSP+Op,ADDpr(CSP+Op,OrangeSave31,OrangeSave32),MULpr(CSP+Op,IndigoSaveDown,p.NFI(2)))

	y,z,w = ASApr(CSP+Op,p.NFI(7),BlueIndigoSideRight,p.NFI(97))
	OrangeDown = ADDpr(CSP+Op,OrangeDown,w)
	OrangeSave34 := y
	OrangeSave35 := z
	IndigoPurpleSideLeft = SUBpr(CSP+Op,ADDpr(CSP+Op,OrangeSave33,OrangeSave34),MULpr(CSP+Op,PurpleSaveUp,p.NFI(3)))

	y,z,w = ASApr(CSP+Op,p.NFI(104),IndigoPurpleSideLeft,p.NFI(1))
	OrangeDown = ADDpr(CSP+Op,OrangeDown,w)
	OrangeSave36 := z
	IndigoPurpleSideRight = SUBpr(CSP+Op,ADDpr(CSP+Op,OrangeSave35,OrangeSave36),MULpr(CSP+Op,PurpleSaveDown,p.NFI(3)))

	y,z,w = ASApr(CSP+Op,p.NFI(7),IndigoPurpleSideRight,p.NFI(105))
	OrangeDown = ADDpr(CSP+Op,OrangeDown,w)
	OD1 	:= MULpr(CSP+Op,p.NFI(4),GreenUp)
	OD2 	:= MULpr(CSP+Op,p.NFI(5),BlueUp)
	OD3 	:= MULpr(CSP+Op,p.NFI(4),IndigoUp)
	OD4 	:= MULpr(CSP+Op,p.NFI(3),PurpleUp)
	OD5 	:= MULpr(CSP+Op,p.NFI(3),PurpleDown)
	OD6 	:= MULpr(CSP+Op,p.NFI(2),IndigoDown)
	OD7 	:= MULpr(CSP+Op,p.NFI(2),GreenDown)
	OD8 	:= MULpr(CSP+Op,p.NFI(6),YellowDown)
	OrangeDown = SUMpr(CSP+Op,OrangeDown,BlueDown,OD1,OD2,OD3,OD4,OD5,OD6,OD7,OD8)
	OrangeTotalDown := ADDpr(CSP+Op,OrangeBaseDown,OrangeDown)

	ODH1 	:= SUBpr(CSP+Op,MULpr(CSP+Op,p.NFI(6),YellowDownHeight),OrangeDownHeight)

	ODH21 	:= MULpr(CSP+Op,p.NFI(2),GreenDownHeight)
	ODH22 	:= MULpr(CSP+Op,p.NFI(4),GreenUpHeight)
	ODH23	:= SUBpr(CSP+Op,ODH21,ODH22)
	ODH2	:= ADDpr(CSP+Op,ODH1,ODH23)

	ODH31 	:= MULpr(CSP+Op,p.NFI(5),BlueUpHeight)
	ODH32 	:= ADDpr(CSP+Op,ODH2,BlueDownHeight)
	ODH3	:= SUBpr(CSP+Op,ODH32,ODH31)

	ODH41 	:= MULpr(CSP+Op,p.NFI(2),IndigoDownHeight)
	ODH42 	:= MULpr(CSP+Op,p.NFI(4),IndigoUpHeight)
	ODH43	:= SUBpr(CSP+Op,ODH41,ODH42)
	ODH4	:= ADDpr(CSP+Op,ODH3,ODH43)

	ODH51 	:= MULpr(CSP+Op,p.NFI(3),PurpleDownHeight)
	ODH52 	:= MULpr(CSP+Op,p.NFI(3),PurpleUpHeight)
	ODH53	:= SUBpr(CSP+Op,ODH51,ODH52)
	ODH5	:= ADDpr(CSP+Op,ODH4,ODH53)
	OrangeDownHeightPeakSeedC := TruncSeed(ODH5)
	fmt.Println("OrangeDownHeightPeakSeed:      ", OrangeDownHeightPeakSeedC)

	ODA11	:= MULpr(CSP+Op,p.NFI(15),YellowDownHeight)
	ODA12	:= ADDpr(CSP+Op,ODA11,OrangeDownHeight)
	ODA13	:= MULpr(CSP+Op,Yellow,ODA12)
	ODA14	:= MULpr(CSP+Op,p.NFI(6),YellowTotalDown)
	ODA15	:= ADDpr(CSP+Op,ODA13,ODA14)
	ODA1	:= SUBpr(CSP+Op,OrangeTotalDown,ODA15)

	ODA21	:= MULpr(CSP+Op,p.NFI(2),ODH1)
	ODA22	:= MULpr(CSP+Op,p.NFI(10),GreenUpHeight)
	ODA23 	:= SUMpr(CSP+Op,GreenDownHeight,ODA21,ODA22)
	ODA24 	:= MULpr(CSP+Op,Green,ODA23)
	ODA25	:= MULpr(CSP+Op,p.NFI(2),GreenTotalDown)
	ODA26	:= MULpr(CSP+Op,p.NFI(4),GreenTotalUp)
	ODA27 	:= SUMpr(CSP+Op,ODA24,ODA25,ODA26)
	ODA2	:= SUBpr(CSP+Op,ODA1,ODA27)

	ODA31	:= ADDpr(CSP+Op,MULpr(CSP+Op,p.NFI(15),BlueUpHeight),ODH2)
	ODA32	:= MULpr(CSP+Op,Blue,ODA31)
	ODA33	:= MULpr(CSP+Op,p.NFI(5),BlueTotalUp)
	ODA34	:= SUMpr(CSP+Op,BlueTotalDown,ODA32,ODA33)
	ODA3	:= SUBpr(CSP+Op,ODA2,ODA34)

	ODA41	:= MULpr(CSP+Op,p.NFI(2),ODH3)
	ODA42	:= MULpr(CSP+Op,p.NFI(10),IndigoUpHeight)
	ODA43	:= SUMpr(CSP+Op,IndigoDownHeight,ODA41,ODA42)
	ODA44	:= MULpr(CSP+Op,Indigo,ODA43)
	ODA45	:= MULpr(CSP+Op,p.NFI(2),IndigoTotalDown)
	ODA46	:= MULpr(CSP+Op,p.NFI(4),IndigoTotalUp)
	ODA47	:= SUMpr(CSP+Op,ODA44,ODA45,ODA46)
	ODA4	:= SUBpr(CSP+Op,ODA3,ODA47)

	ODA51	:= MULpr(CSP+Op,p.NFI(3),ODH4)
	ODA52	:= MULpr(CSP+Op,p.NFI(3),PurpleDownHeight)
	ODA53	:= MULpr(CSP+Op,p.NFI(6),PurpleUpHeight)
	ODA54	:= SUMpr(CSP+Op,ODA51,ODA52,ODA53)
	ODA55	:= MULpr(CSP+Op,Purple,ODA54)
	ODA56	:= ADDpr(CSP+Op,PurpleTotalUp,PurpleTotalDown)
	ODA57	:= MULpr(CSP+Op,p.NFI(3),ODA56)
	ODA58	:= ADDpr(CSP+Op,ODA55,ODA57)
	ODA5	:= SUBpr(CSP+Op,ODA4,ODA58)
	OrangeDownAreaPeakSeedC := TruncSeed(ODA5)
	fmt.Println("OrangeDownAreaPeakSeed is:  ", OrangeDownAreaPeakSeedC)
	fmt.Println("========================================================")
	//
	//
	//Calculating RedUp
	//
	y,z,w = ASApr(CSP+RUp,p.NFI(90),Red,p.NFI(28))
	RedSaveUp := y
	RedBaseUp := w
	RedUpHeight := z

	y,z,w = ASApr(CSP+RUp,p.NFI(42),y,p.NFI(7))
	RedUp := w
	RedOrangeSideLeft := SUBpr(CSP+RUp,y,MULpr(CSP+RUp,p.NFI(5),OrangeSaveUp))
	RedSave01 := z

	y,z,w = ASApr(CSP+RUp,p.NFI(49),RedOrangeSideLeft,p.NFI(6))
	RedUp = ADDpr(CSP+RUp,RedUp,w)
	RedSave02 := z
	RedSave03 := y
	RedOrangeSideRight := SUBpr(CSP+RUp,ADDpr(CSP+RUp,RedSave01,RedSave02),OrangeSaveDown)

	y,z,w = ASApr(CSP+RUp,p.NFI(7),RedOrangeSideRight,p.NFI(55))
	RedUp = ADDpr(CSP+RUp,RedUp,w)
	RedSave04 := y
	RedSave05 := z
	OrangeYellowSideLeft = SUBpr(CSP+RUp,ADDpr(CSP+RUp,RedSave03,RedSave04),MULpr(CSP+RUp,YellowSaveUp,p.NFI(4)))

	y,z,w = ASApr(CSP+RUp,p.NFI(62),OrangeYellowSideLeft,p.NFI(5))
	RedUp = ADDpr(CSP+RUp,RedUp,w)
	RedSave06 := z
	RedSave07 := y
	OrangeYellowSideRight = SUBpr(CSP+RUp,ADDpr(CSP+RUp,RedSave05,RedSave06),MULpr(CSP+RUp,YellowSaveDown,p.NFI(2)))

	y,z,w = ASApr(CSP+RUp,p.NFI(7),OrangeYellowSideRight,p.NFI(67))
	RedUp = ADDpr(CSP+RUp,RedUp,w)
	RedSave08 := y
	RedSave09 := z
	YellowGreenSideLeft = SUBpr(CSP+RUp,ADDpr(CSP+RUp,RedSave07,RedSave08),MULpr(CSP+RUp,GreenSaveUp,p.NFI(4)))

	y,z,w = ASApr(CSP+RUp,p.NFI(74),YellowGreenSideLeft,p.NFI(4))
	RedUp = ADDpr(CSP+RUp,RedUp,w)
	RedSave10 := z
	RedSave11 := y
	YellowGreenSideRight = SUBpr(CSP+RUp,ADDpr(CSP+RUp,RedSave09,RedSave10),MULpr(CSP+RUp,GreenSaveDown,p.NFI(2)))

	y,z,w = ASApr(CSP+RUp,p.NFI(7),YellowGreenSideRight,p.NFI(78))
	RedUp = ADDpr(CSP+RUp,RedUp,w)
	RedSave12 := y
	RedSave13 := z
	GreenBlueSideLeft = SUBpr(CSP+RUp,ADDpr(CSP+RUp,RedSave11,RedSave12),MULpr(CSP+RUp,BlueSaveUp,p.NFI(6)))

	y,z,w = ASApr(CSP+RUp,p.NFI(85),GreenBlueSideLeft,p.NFI(3))
	RedUp = ADDpr(CSP+RUp,RedUp,w)
	RedSave14 := z
	RedSave15 := y
	GreenBlueSideRight = ADDpr(CSP+RUp,RedSave13,RedSave14)

	y,z,w = ASApr(CSP+RUp,p.NFI(7),GreenBlueSideRight,p.NFI(88))
	RedUp = ADDpr(CSP+RUp,RedUp,w)
	RedSave16 := y
	RedSave17 := z
	BlueIndigoSideLeft = SUBpr(CSP+RUp,ADDpr(CSP+RUp,RedSave15,RedSave16),MULpr(CSP+RUp,IndigoSaveUp,p.NFI(2)))

	y,z,w = ASApr(CSP+RUp,p.NFI(95),BlueIndigoSideLeft,p.NFI(2))
	RedUp = ADDpr(CSP+RUp,RedUp,w)
	RedSave18 := z
	RedSave19 := y
	BlueIndigoSideRight = SUBpr(CSP+RUp,ADDpr(CSP+RUp,RedSave17,RedSave18),MULpr(CSP+RUp,IndigoSaveDown,p.NFI(4)))

	y,z,w = ASApr(CSP+RUp,p.NFI(7),BlueIndigoSideRight,p.NFI(97))
	RedUp = ADDpr(CSP+RUp,RedUp,w)
	RedSave20 := y
	RedSave21 := z
	IndigoPurpleSideLeft = SUBpr(CSP+RUp,ADDpr(CSP+RUp,RedSave19,RedSave20),MULpr(CSP+RUp,PurpleSaveUp,p.NFI(2)))

	y,z,w = ASApr(CSP+RUp,p.NFI(104),IndigoPurpleSideLeft,p.NFI(1))
	RedUp = ADDpr(CSP+RUp,RedUp,w)
	RedSave22 := z
	IndigoPurpleSideRight = SUBpr(CSP+RUp,ADDpr(CSP+RUp,RedSave21,RedSave22),MULpr(CSP+RUp,PurpleSaveDown,p.NFI(4)))

	y,z,w = ASApr(CSP+RUp,p.NFI(7),IndigoPurpleSideRight,p.NFI(105))
	RedUp = ADDpr(CSP+RUp,RedUp,w)
	RU1 	:= MULpr(CSP+RUp,p.NFI(5),OrangeUp)
	RU2 	:= MULpr(CSP+RUp,p.NFI(6),BlueUp)
	RU3		:= SUMpr(CSP+RUp,YellowUp,GreenUp,PurpleDown,IndigoDown)
	RU4 	:= SUMpr(CSP+RUp,IndigoUp,PurpleUp,GreenDown,YellowDown)
	RU5 	:= MULpr(CSP+RUp,p.NFI(4),RU3)
	RU6 	:= MULpr(CSP+RUp,p.NFI(2),RU4)
	RedUp	 = SUMpr(CSP+RUp,RedUp,OrangeDown,RU1,RU2,RU5,RU6)
	RedTotalUp := ADDpr(CSP+RUp,RedBaseUp,RedUp)

	RUH11 	:= MULpr(CSP+RUp,p.NFI(5),OrangeUpHeight)
	RUH12	:= ADDpr(CSP+RUp,RedUpHeight,OrangeDownHeight)
	RUH1	:= SUBpr(CSP+RUp,RUH12,RUH11)

	RUH21 	:= MULpr(CSP+RUp,p.NFI(4),YellowUpHeight)
	RUH22 	:= MULpr(CSP+RUp,p.NFI(2),YellowDownHeight)
	RUH23	:= SUBpr(CSP+RUp,RUH22,RUH21)
	RUH2	:= ADDpr(CSP+RUp,RUH1,RUH23)

	RUH31 	:= MULpr(CSP+RUp,p.NFI(4),GreenUpHeight)
	RUH32 	:= MULpr(CSP+RUp,p.NFI(2),GreenDownHeight)
	RUH33	:= SUBpr(CSP+RUp,RUH32,RUH31)
	RUH3	:= ADDpr(CSP+RUp,RUH2,RUH33)

	RUH4	:= SUBpr(CSP+RUp,MULpr(CSP+RUp,p.NFI(6),BlueUpHeight),RUH3)

	RUH51 	:= MULpr(CSP+RUp,p.NFI(4),IndigoDownHeight)
	RUH52 	:= MULpr(CSP+RUp,p.NFI(2),IndigoUpHeight)
	RUH53	:= SUBpr(CSP+RUp,RUH52,RUH51)
	RUH5	:= ADDpr(CSP+RUp,RUH4,RUH53)

	RUH61 	:= MULpr(CSP+RUp,p.NFI(4),PurpleDownHeight)
	RUH62 	:= MULpr(CSP+RUp,p.NFI(2),PurpleUpHeight)
	RUH6 	:= DIFpr(CSP+RUp,RUH61,RUH62,RUH5)
	RedUpHeightPeakSeedC := TruncSeed(RUH6)
	fmt.Println("RedUpHeightPeakSeed is:        ", RedUpHeightPeakSeedC)

	RUA11	:= MULpr(CSP+RUp,p.NFI(15),OrangeUpHeight)
	RUA12 	:= ADDpr(CSP+RUp,RedUpHeight,RUA11)
	RUA13	:= MULpr(CSP+RUp,Orange,RUA12)
	RUA14	:= MULpr(CSP+RUp,p.NFI(5),OrangeTotalUp)
	RUA15 	:= SUMpr(CSP+RUp,OrangeTotalDown,RUA13,RUA14)
	RUA1	:= SUBpr(CSP+RUp,RedTotalUp,RUA15)

	RUA21	:= MULpr(CSP+RUp,p.NFI(2),RUH1)
	RUA22	:= MULpr(CSP+RUp,p.NFI(10),YellowUpHeight)
	RUA23	:= SUMpr(CSP+RUp,YellowDownHeight,RUA21,RUA22)
	RUA24	:= MULpr(CSP+RUp,Yellow,RUA23)
	RUA25	:= MULpr(CSP+RUp,p.NFI(2),YellowTotalDown)
	RUA26	:= MULpr(CSP+RUp,p.NFI(4),YellowTotalUp)
	RUA27	:= SUMpr(CSP+RUp,RUA24,RUA25,RUA26)
	RUA2	:= SUBpr(CSP+RUp,RUA1,RUA27)

	RUA31	:= MULpr(CSP+RUp,p.NFI(2),RUH2)
	RUA32	:= MULpr(CSP+RUp,p.NFI(10),GreenUpHeight)
	RUA33	:= SUMpr(CSP+RUp,GreenDownHeight,RUA31,RUA32)
	RUA34	:= MULpr(CSP+RUp,Green,RUA33)
	RUA35	:= MULpr(CSP+RUp,p.NFI(2),GreenTotalDown)
	RUA36	:= MULpr(CSP+RUp,p.NFI(4),GreenTotalUp)
	RUA37	:= SUMpr(CSP+RUp,RUA34,RUA35,RUA36)
	RUA3	:= SUBpr(CSP+RUp,RUA2,RUA37)

	RUA41	:= ADDpr(CSP+RUp,MULpr(CSP+RUp,p.NFI(15),BlueUpHeight),RUH3)
	RUA42	:= MULpr(CSP+RUp,Blue,RUA41)
	RUA43	:= MULpr(CSP+RUp,p.NFI(6),BlueTotalUp)
	RUA44	:= ADDpr(CSP+RUp,RUA42,RUA43)
	RUA4	:= SUBpr(CSP+RUp,RUA3,RUA44)

	RUA51	:= MULpr(CSP+RUp,p.NFI(2),RUH4)
	RUA52	:= MULpr(CSP+RUp,p.NFI(10),IndigoDownHeight)
	RUA53 	:= SUMpr(CSP+RUp,IndigoUpHeight,RUA51,RUA52)
	RUA54 	:= MULpr(CSP+RUp,Indigo,RUA53)
	RUA55	:= MULpr(CSP+RUp,p.NFI(2),IndigoTotalUp)
	RUA56	:= MULpr(CSP+RUp,p.NFI(4),IndigoTotalDown)
	RUA57	:= SUMpr(CSP+RUp,RUA54,RUA55,RUA56)
	RUA5	:= SUBpr(CSP+RUp,RUA4,RUA57)

	RUA61	:= MULpr(CSP+RUp,p.NFI(3),RUH5)
	RUA62	:= MULpr(CSP+RUp,p.NFI(6),PurpleDownHeight)
	RUA63	:= MULpr(CSP+RUp,p.NFI(3),PurpleUpHeight)
	RUA64 	:= SUMpr(CSP+RUp,RUA61,RUA62,RUA63)
	RUA65 	:= MULpr(CSP+RUp,Purple,RUA64)
	RUA66	:= MULpr(CSP+RUp,p.NFI(2),PurpleTotalUp)
	RUA67	:= MULpr(CSP+RUp,p.NFI(4),PurpleTotalDown)
	RUA68	:= SUMpr(CSP+RUp,RUA65,RUA66,RUA67)
	RUA6	:= SUBpr(CSP+RUp,RUA5,RUA68)
	RedUpAreaPeakSeedC := TruncSeed(RUA6)
	fmt.Println("RedUpAreaPeakSeed is:       ", RedUpAreaPeakSeedC)
	//
	//
	//Calculating RedUp
	//
	y,z,w = ASApr(CSP+RDp,p.NFI(7),Red,p.NFI(90))
	RedDownHeight := y
	RedSaveDown := z
	RedBaseDown := w

	y,z,w = ASApr(CSP+RDp,p.NFI(7),z,p.NFI(42))
	RedDown := w
	RedSave23 := z

	y,z,w = ASApr(CSP+RDp,p.NFI(49),y,p.NFI(6))
	RedDown = ADDpr(CSP+RDp,RedDown,w)
	RedSave24 := z
	RedSave25 := y
	RedOrangeSideRight = SUBpr(CSP+RDp,ADDpr(CSP+RDp,RedSave23,RedSave24),MULpr(CSP+RDp,OrangeSaveDown,p.NFI(6)))

	y,z,w = ASApr(CSP+RDp,p.NFI(7),RedOrangeSideRight,p.NFI(55))
	RedDown = ADDpr(CSP+RDp,RedDown,w)
	RedSave26 := y
	RedSave27 := z
	OrangeYellowSideLeft = SUBpr(CSP+RDp,ADDpr(CSP+RDp,RedSave25,RedSave26),MULpr(CSP+RDp,YellowSaveUp,p.NFI(5)))

	y,z,w = ASApr(CSP+RDp,p.NFI(62),OrangeYellowSideLeft,p.NFI(5))
	RedDown = ADDpr(CSP+RDp,RedDown,w)
	RedSave28 := z
	RedSave29 := y
	OrangeYellowSideRight = SUBpr(CSP+RDp,ADDpr(CSP+RDp,RedSave27,RedSave28),YellowSaveDown)

	y,z,w = ASApr(CSP+RDp,p.NFI(7),OrangeYellowSideRight,p.NFI(67))
	RedDown = ADDpr(CSP+RDp,RedDown,w)
	RedSave30 := y
	RedSave31 := z
	YellowGreenSideLeft = SUBpr(CSP+RDp,ADDpr(CSP+RDp,RedSave29,RedSave30),MULpr(CSP+RDp,GreenSaveUp,p.NFI(4)))

	y,z,w = ASApr(CSP+RDp,p.NFI(74),YellowGreenSideLeft,p.NFI(4))
	RedDown = ADDpr(CSP+RDp,RedDown,w)
	RedSave32 := z
	RedSave33 := y
	YellowGreenSideRight = SUBpr(CSP+RDp,ADDpr(CSP+RDp,RedSave31,RedSave32),MULpr(CSP+RDp,GreenSaveDown,p.NFI(2)))

	y,z,w = ASApr(CSP+RDp,p.NFI(7),YellowGreenSideRight,p.NFI(78))
	RedDown = ADDpr(CSP+RDp,RedDown,w)
	RedSave34 := y
	RedSave35 := z
	GreenBlueSideLeft = SUBpr(CSP+RDp,ADDpr(CSP+RDp,RedSave33,RedSave34),MULpr(CSP+RDp,BlueSaveUp,p.NFI(3)))

	y,z,w = ASApr(CSP+RDp,p.NFI(85),GreenBlueSideLeft,p.NFI(3))
	RedDown = ADDpr(CSP+RDp,RedDown,w)
	RedSave36 := z
	RedSave37 := y
	GreenBlueSideRight = SUBpr(CSP+RDp,ADDpr(CSP+RDp,RedSave35,RedSave36),MULpr(CSP+RDp,BlueSaveDown,p.NFI(3)))

	y,z,w = ASApr(CSP+RDp,p.NFI(7),GreenBlueSideRight,p.NFI(88))
	RedDown = ADDpr(CSP+RDp,RedDown,w)
	RedSave38 := y
	RedSave39 := z
	BlueIndigoSideLeft = SUBpr(CSP+RDp,ADDpr(CSP+RDp,RedSave37,RedSave38),MULpr(CSP+RDp,IndigoSaveUp,p.NFI(6)))

	y,z,w = ASApr(CSP+RDp,p.NFI(95),BlueIndigoSideLeft,p.NFI(2))
	RedDown = ADDpr(CSP+RDp,RedDown,w)
	RedSave40 := z
	RedSave41 := y
	BlueIndigoSideRight = ADDpr(CSP+RDp,RedSave39,RedSave40)

	y,z,w = ASApr(CSP+RDp,p.NFI(7),BlueIndigoSideRight,p.NFI(97))
	RedDown = ADDpr(CSP+RDp,RedDown,w)
	RedSave42 := y
	RedSave43 := z
	IndigoPurpleSideLeft = SUBpr(CSP+RDp,ADDpr(CSP+RDp,RedSave41,RedSave42),MULpr(CSP+RDp,PurpleSaveUp,p.NFI(2)))

	y,z,w = ASApr(CSP+RDp,p.NFI(104),IndigoPurpleSideLeft,p.NFI(1))
	RedDown = ADDpr(CSP+RDp,RedDown,w)
	RedSave44 := z
	IndigoPurpleSideRight = SUBpr(CSP+RDp,ADDpr(CSP+RDp,RedSave43,RedSave44),MULpr(CSP+RDp,PurpleSaveDown,p.NFI(4)))

	y,z,w = ASApr(CSP+RDp,p.NFI(7),IndigoPurpleSideRight,p.NFI(105))
	RedDown = ADDpr(CSP+RDp,RedDown,w)
	RD1		:= MULpr(CSP+RDp,p.NFI(5),YellowUp)
	RD2		:= MULpr(CSP+RDp,p.NFI(4),GreenUp)
	RD3		:= MULpr(CSP+RDp,p.NFI(3),BlueUp)
	RD4		:= MULpr(CSP+RDp,p.NFI(6),IndigoUp)
	RD5		:= MULpr(CSP+RDp,p.NFI(2),PurpleUp)
	RD6		:= MULpr(CSP+RDp,p.NFI(6),OrangeDown)
	RD7		:= MULpr(CSP+RDp,p.NFI(2),GreenDown)
	RD8		:= MULpr(CSP+RDp,p.NFI(3),BlueDown)
	RD9		:= MULpr(CSP+RDp,p.NFI(4),PurpleDown)
	RedDown = SUMpr(CSP+RDp,RedDown,YellowDown,RD1,RD2,RD3,RD4,RD5,RD6,RD7,RD8,RD9)
	RedTotalDown := ADDpr(CSP+RDp,RedBaseDown,RedDown)

	RDH1 	:= SUBpr(CSP+RDp,MULpr(CSP+RDp,p.NFI(6),OrangeDownHeight),RedDownHeight)

	RDH21	:= MULpr(CSP+RDp,p.NFI(5),YellowUpHeight)
	RDH22	:= SUBpr(CSP+RDp,YellowDownHeight,RDH21)
	RDH2	:= ADDpr(CSP+RDp,RDH1,RDH22)

	RDH31	:= MULpr(CSP+RDp,p.NFI(4),GreenUpHeight)
	RDH32	:= MULpr(CSP+RDp,p.NFI(2),GreenDownHeight)
	RDH33	:= SUBpr(CSP+RDp,RDH32,RDH31)
	RDH3	:= ADDpr(CSP+RDp,RDH2,RDH33)

	RDH41	:= MULpr(CSP+RDp,p.NFI(3),BlueUpHeight)
	RDH42	:= MULpr(CSP+RDp,p.NFI(3),BlueDownHeight)
	RDH43	:= SUBpr(CSP+RDp,RDH42,RDH41)
	RDH4	:= ADDpr(CSP+RDp,RDH3,RDH43)

	RDH5	:= SUBpr(CSP+RDp,MULpr(CSP+RDp,p.NFI(6),IndigoUpHeight),RDH4)

	RDH61	:= ADDpr(CSP+RDp,MULpr(CSP+RDp,p.NFI(2),PurpleUpHeight),RDH5)
	RDH6	:= SUBpr(CSP+RDp,MULpr(CSP+RDp,p.NFI(4),PurpleDownHeight),RDH61)
	RedDownHeightPeakSeedC := TruncSeed(RDH6)
	fmt.Println("RedDownHeightPeakSeed:         ", RedDownHeightPeakSeedC)

	RDA11	:= ADDpr(CSP+RDp,MULpr(CSP+RDp,p.NFI(15),OrangeDownHeight),RedDownHeight)
	RDA12	:= ADDpr(CSP+RDp,MULpr(CSP+RDp,p.NFI(6),OrangeTotalDown),MULpr(CSP+RDp,Orange,RDA11))
	RDA1	:= SUBpr(CSP+RDp,RedTotalDown,RDA12)

	RDA21	:= ADDpr(CSP+RDp,MULpr(CSP+RDp,p.NFI(15),YellowUpHeight),RDH1)
	RDA22	:= MULpr(CSP+RDp,Yellow,RDA21)
	RDA23	:= MULpr(CSP+RDp,p.NFI(5),YellowTotalUp)
	RDA24	:= SUMpr(CSP+RDp,YellowTotalDown,RDA22,RDA23)
	RDA2	:= SUBpr(CSP+RDp,RDA1,RDA24)

	RDA31	:= MULpr(CSP+RDp,p.NFI(2),RDH2)
	RDA32	:= MULpr(CSP+RDp,p.NFI(10),GreenUpHeight)
	RDA33	:= SUMpr(CSP+RDp,GreenDownHeight,RDA31,RDA32)
	RDA34 	:= MULpr(CSP+RDp,Green,RDA33)
	RDA35	:= MULpr(CSP+RDp,p.NFI(2),GreenTotalDown)
	RDA36	:= MULpr(CSP+RDp,p.NFI(4),GreenTotalUp)
	RDA37	:= SUMpr(CSP+RDp,RDA34,RDA35,RDA36)
	RDA3	:= SUBpr(CSP+RDp,RDA2,RDA37)

	RDA41	:= MULpr(CSP+RDp,p.NFI(3),RDH3)
	RDA42	:= MULpr(CSP+RDp,p.NFI(3),BlueDownHeight)
	RDA43	:= MULpr(CSP+RDp,p.NFI(6),BlueUpHeight)
	RDA44	:= SUMpr(CSP+RDp,RDA41,RDA42,RDA43)
	RDA45	:= MULpr(CSP+RDp,Blue,RDA44)
	RDA46	:= ADDpr(CSP+RDp,BlueTotalUp,BlueTotalDown)
	RDA47	:= MULpr(CSP+RDp,p.NFI(3),RDA46)
	RDA48	:= ADDpr(CSP+RDp,RDA45,RDA47)
	RDA4	:= SUBpr(CSP+RDp,RDA3,RDA48)

	RDA51 	:= ADDpr(CSP+RDp,MULpr(CSP+RDp,p.NFI(15),IndigoUpHeight),RDH4)
	RDA52	:= MULpr(CSP+RDp,Indigo,RDA51)
	RDA53	:= MULpr(CSP+RDp,p.NFI(6),IndigoTotalUp)
	RDA54	:= ADDpr(CSP+RDp,RDA52,RDA53)
	RDA5	:= SUBpr(CSP+RDp,RDA4,RDA54)

	RDA61	:= MULpr(CSP+RDp,p.NFI(3),RDH5)
	RDA62	:= MULpr(CSP+RDp,p.NFI(6),PurpleDownHeight)
	RDA63	:= MULpr(CSP+RDp,p.NFI(3),PurpleUpHeight)
	RDA64	:= SUMpr(CSP+RDp,RDA61,RDA62,RDA63)
	RDA65	:= MULpr(CSP+RDp,Purple,RDA64)
	RDA66	:= MULpr(CSP+RDp,p.NFI(4),PurpleTotalDown)
	RDA67	:= MULpr(CSP+RDp,p.NFI(2),PurpleTotalUp)
	RDA68	:= SUMpr(CSP+RDp,RDA65,RDA66,RDA67)
	RDA6	:= SUBpr(CSP+RDp,RDA5,RDA68)
	RedDownAreaPeakSeedC := TruncSeed(RDA6)
	fmt.Println("RedDownAreaPeakSeed is:     ", RedDownAreaPeakSeedC)
	fmt.Println("========================================================")
	//
	//
	//Calculating White - TotalArea
	//
	y,z,w = ASApr(CSP+Whp,p.NFI(7),White,p.NFI(28))
	TotalArea := w
	WhiteSave01 := z
	WhiteRedSideLeft := SUBpr(CSP+Whp,y,RedSaveUp)

	y,z,w = ASApr(CSP+Whp,p.NFI(35),WhiteRedSideLeft,p.NFI(7))
	TotalArea = ADDpr(CSP+Whp,TotalArea,w)
	WhiteSave02 := z
	WhiteSave03	:= y
	WhiteRedSideRight := SUBpr(CSP+Whp,ADDpr(CSP+Whp,WhiteSave01,WhiteSave02),MULpr(CSP+Whp,RedSaveDown,p.NFI(5)))

	y,z,w = ASApr(CSP+Whp,p.NFI(7),WhiteRedSideRight,p.NFI(42))
	TotalArea = ADDpr(CSP+Whp,TotalArea,w)
	WhiteSave04 := y
	WhiteSave05	:= z
	RedOrangeSideLeft = SUBpr(CSP+Whp,ADDpr(CSP+Whp,WhiteSave03,WhiteSave04),MULpr(CSP+Whp,OrangeSaveUp,p.NFI(2)))

	y,z,w = ASApr(CSP+Whp,p.NFI(49),RedOrangeSideLeft,p.NFI(6))
	TotalArea = ADDpr(CSP+Whp,TotalArea,w)
	WhiteSave06 := z
	WhiteSave07	:= y
	RedOrangeSideRight = SUBpr(CSP+Whp,ADDpr(CSP+Whp,WhiteSave05,WhiteSave06),MULpr(CSP+Whp,OrangeSaveDown,p.NFI(4)))

	y,z,w = ASApr(CSP+Whp,p.NFI(7),RedOrangeSideRight,p.NFI(55))
	TotalArea = ADDpr(CSP+Whp,TotalArea,w)
	WhiteSave08 := y
	WhiteSave09	:= z
	OrangeYellowSideLeft = SUBpr(CSP+Whp,ADDpr(CSP+Whp,WhiteSave07,WhiteSave08),MULpr(CSP+Whp,YellowSaveUp,p.NFI(3)))

	y,z,w = ASApr(CSP+Whp,p.NFI(62),OrangeYellowSideLeft,p.NFI(5))
	TotalArea = ADDpr(CSP+Whp,TotalArea,w)
	WhiteSave10 := z
	WhiteSave11	:= y
	OrangeYellowSideRight = SUBpr(CSP+Whp,ADDpr(CSP+Whp,WhiteSave09,WhiteSave10),MULpr(CSP+Whp,YellowSaveDown,p.NFI(3)))

	y,z,w = ASApr(CSP+Whp,p.NFI(7),OrangeYellowSideRight,p.NFI(67))
	TotalArea = ADDpr(CSP+Whp,TotalArea,w)
	WhiteSave12 := y
	WhiteSave13	:= z
	YellowGreenSideLeft = SUBpr(CSP+Whp,ADDpr(CSP+Whp,WhiteSave11,WhiteSave12),GreenSaveUp)

	y,z,w = ASApr(CSP+Whp,p.NFI(74),YellowGreenSideLeft,p.NFI(4))
	TotalArea = ADDpr(CSP+Whp,TotalArea,w)
	WhiteSave14 := z
	WhiteSave15	:= y
	YellowGreenSideRight = SUBpr(CSP+Whp,ADDpr(CSP+Whp,WhiteSave13,WhiteSave14),MULpr(CSP+Whp,GreenSaveDown,p.NFI(5)))

	y,z,w = ASApr(CSP+Whp,p.NFI(7),YellowGreenSideRight,p.NFI(78))
	TotalArea = ADDpr(CSP+Whp,TotalArea,w)
	WhiteSave16 := y
	WhiteSave17	:= z
	GreenBlueSideLeft = SUBpr(CSP+Whp,ADDpr(CSP+Whp,WhiteSave15,WhiteSave16),MULpr(CSP+Whp,BlueSaveUp,p.NFI(2)))

	y,z,w = ASApr(CSP+Whp,p.NFI(85),GreenBlueSideLeft,p.NFI(3))
	TotalArea = ADDpr(CSP+Whp,TotalArea,w)
	WhiteSave18 := z
	WhiteSave19	:= y
	GreenBlueSideRight = SUBpr(CSP+Whp,ADDpr(CSP+Whp,WhiteSave17,WhiteSave18),MULpr(CSP+Whp,BlueSaveDown,p.NFI(4)))

	y,z,w = ASApr(CSP+Whp,p.NFI(7),GreenBlueSideRight,p.NFI(88))
	TotalArea = ADDpr(CSP+Whp,TotalArea,w)
	WhiteSave20 := y
	WhiteSave21	:= z
	BlueIndigoSideLeft = SUBpr(CSP+Whp,ADDpr(CSP+Whp,WhiteSave19,WhiteSave20),IndigoSaveUp)

	y,z,w = ASApr(CSP+Whp,p.NFI(95),BlueIndigoSideLeft,p.NFI(2))
	TotalArea = ADDpr(CSP+Whp,TotalArea,w)
	WhiteSave22 := z
	WhiteSave23	:= y
	BlueIndigoSideRight = SUBpr(CSP+Whp,ADDpr(CSP+Whp,WhiteSave21,WhiteSave22),MULpr(CSP+Whp,IndigoSaveDown,p.NFI(5)))

	y,z,w = ASApr(CSP+Whp,p.NFI(7),BlueIndigoSideRight,p.NFI(97))
	TotalArea = ADDpr(CSP+Whp,TotalArea,w)
	WhiteSave24 := y
	WhiteSave25	:= z
	IndigoPurpleSideLeft = SUBpr(CSP+Whp,ADDpr(CSP+Whp,WhiteSave23,WhiteSave24),MULpr(CSP+Whp,PurpleSaveUp,p.NFI(3)))

	y,z,w = ASApr(CSP+Whp,p.NFI(104),IndigoPurpleSideLeft,p.NFI(1))
	TotalArea = ADDpr(CSP+Whp,TotalArea,w)
	WhiteSave26 := z
	IndigoPurpleSideRight = SUBpr(CSP+Whp,ADDpr(CSP+Whp,WhiteSave25,WhiteSave26),MULpr(CSP+Whp,PurpleSaveDown,p.NFI(3)))

	y,z,w = ASApr(CSP+Whp,p.NFI(7),IndigoPurpleSideRight,p.NFI(105))
	TotalArea = ADDpr(CSP+Whp,TotalArea,w)
	TA1		:= MULpr(CSP+Whp,p.NFI(2),OrangeUp)
	TA2		:= MULpr(CSP+Whp,p.NFI(2),BlueUp)
	TA3		:= SUMpr(CSP+Whp,YellowUp,PurpleUp,PurpleDown,YellowDown)
	TA4		:= MULpr(CSP+Whp,p.NFI(3),TA3)
	TA5		:= MULpr(CSP+Whp,p.NFI(4),BlueDown)
	TA6		:= MULpr(CSP+Whp,p.NFI(4),OrangeDown)
	TA7		:= SUMpr(CSP+Whp,IndigoDown,GreenDown,RedDown)
	TA8 	:= MULpr(CSP+Whp,p.NFI(5),TA7)
	TotalArea = SUMpr(CSP+Whp,TotalArea,RedUp,GreenUp,IndigoUp,TA1,TA2,TA4,TA5,TA6,TA8)

	WH1 	:= SUBpr(CSP+Whp,MULpr(CSP+Whp,p.NFI(5),RedDownHeight),RedUpHeight)

	WH21	:= SUBpr(CSP+Whp,MULpr(CSP+Whp,p.NFI(4),OrangeDownHeight),MULpr(CSP+Whp,p.NFI(2),OrangeUpHeight))
	WH2		:= ADDpr(CSP+Whp,WH1,WH21)

	WH31	:= MULpr(CSP+Whp,p.NFI(3),YellowUpHeight)
	WH32	:= MULpr(CSP+Whp,p.NFI(3),YellowDownHeight)
	WH3		:= DIFpr(CSP+Whp,WH31,WH32,WH2)

	WH41	:= SUBpr(CSP+Whp,GreenUpHeight,MULpr(CSP+Whp,p.NFI(5),GreenDownHeight))
	WH4		:= ADDpr(CSP+Whp,WH3,WH41)

	WH51	:= MULpr(CSP+Whp,p.NFI(2),BlueUpHeight)
	WH52	:= MULpr(CSP+Whp,p.NFI(4),BlueDownHeight)
	WH53	:= SUBpr(CSP+Whp,WH51,WH52)
	WH5		:= ADDpr(CSP+Whp,WH4,WH53)

	WH61	:= MULpr(CSP+Whp,p.NFI(5),IndigoDownHeight)
	WH6		:= DIFpr(CSP+Whp,WH61,IndigoUpHeight,WH5)

	WH71	:= MULpr(CSP+Whp,p.NFI(3),PurpleUpHeight)
	WH72	:= MULpr(CSP+Whp,p.NFI(3),PurpleDownHeight)
	WH7		:= DIFpr(CSP+Whp,WH71,WH72,WH6)
	WhiteHeightPeakSeedC := TruncSeed(WH7)
	fmt.Println("WhiteHeightPeakSeed is:        ", WhiteHeightPeakSeedC)

	WA11	:= ADDpr(CSP+Whp,MULpr(CSP+Whp,p.NFI(10),RedDownHeight),RedUpHeight)
	WA12	:= MULpr(CSP+Whp,Red,WA11)
	WA13	:= MULpr(CSP+Whp,p.NFI(5),RedTotalDown)
	WA14	:= SUMpr(CSP+Whp,RedTotalUp,WA13,WA12)
	WA1		:= SUBpr(CSP+Whp,TotalArea,WA14)

	WA21	:= MULpr(CSP+Whp,p.NFI(4),WH1)
	WA22	:= MULpr(CSP+Whp,p.NFI(6),OrangeDownHeight)
	WA23	:= MULpr(CSP+Whp,p.NFI(3),OrangeUpHeight)
	WA24	:= SUMpr(CSP+Whp,WA21,WA22,WA23)
	WA25	:= MULpr(CSP+Whp,Orange,WA24)
	WA26	:= MULpr(CSP+Whp,p.NFI(4),OrangeTotalDown)
	WA27	:= MULpr(CSP+Whp,p.NFI(2),OrangeTotalUp)
	WA28	:= SUMpr(CSP+Whp,WA25,WA26,WA27)
	WA2		:= SUBpr(CSP+Whp,WA1,WA28)

	WA31	:= MULpr(CSP+Whp,p.NFI(4),WH2)
	WA32	:= MULpr(CSP+Whp,p.NFI(6),YellowDownHeight)
	WA33	:= MULpr(CSP+Whp,p.NFI(3),YellowUpHeight)
	WA34	:= SUMpr(CSP+Whp,WA31,WA32,WA33)
	WA35	:= MULpr(CSP+Whp,Yellow,WA34)
	WA36	:= MULpr(CSP+Whp,p.NFI(3),YellowTotalDown)
	WA37	:= MULpr(CSP+Whp,p.NFI(3),YellowTotalUp)
	WA38	:= SUMpr(CSP+Whp,WA35,WA36,WA37)
	WA3		:= SUBpr(CSP+Whp,WA2,WA38)

	WA41	:= ADDpr(CSP+Whp,MULpr(CSP+Whp,p.NFI(15),GreenDownHeight),WH3)
	WA42	:= MULpr(CSP+Whp,Green,WA41)
	WA43	:= MULpr(CSP+Whp,p.NFI(5),GreenTotalDown)
	WA44	:= SUMpr(CSP+Whp,GreenTotalUp,WA43,WA42)
	WA4		:= SUBpr(CSP+Whp,WA3,WA44)

	WA51	:= MULpr(CSP+Whp,p.NFI(2),WH4)
	WA52	:= MULpr(CSP+Whp,p.NFI(10),BlueDownHeight)
	WA53	:= SUMpr(CSP+Whp,WA51,WA52,BlueUpHeight)
	WA54	:= MULpr(CSP+Whp,Blue,WA53)
	WA55	:= MULpr(CSP+Whp,p.NFI(4),BlueTotalDown)
	WA56	:= MULpr(CSP+Whp,p.NFI(2),BlueTotalUp)
	WA57	:= SUMpr(CSP+Whp,WA54,WA55,WA56)
	WA5		:= SUBpr(CSP+Whp,WA4,WA57)

	WA61	:= MULpr(CSP+Whp,p.NFI(2),WH5)
	WA62	:= MULpr(CSP+Whp,p.NFI(10),IndigoDownHeight)
	WA63	:= SUMpr(CSP+Whp,WA61,WA62,IndigoUpHeight)
	WA64	:= MULpr(CSP+Whp,Indigo,WA63)
	WA65	:= MULpr(CSP+Whp,p.NFI(5),IndigoTotalDown)
	WA66	:= SUMpr(CSP+Whp,WA64,WA65,IndigoTotalUp)
	WA6		:= SUBpr(CSP+Whp,WA5,WA66)

	WA71	:= MULpr(CSP+Whp,p.NFI(4),WH6)
	WA72	:= MULpr(CSP+Whp,p.NFI(6),PurpleDownHeight)
	WA73	:= MULpr(CSP+Whp,p.NFI(3),PurpleUpHeight)
	WA74	:= SUMpr(CSP+Whp,WA71,WA72,WA73)
	WA75	:= MULpr(CSP+Whp,Purple,WA74)
	WA76	:= MULpr(CSP+Whp,p.NFI(3),PurpleTotalDown)
	WA77	:= MULpr(CSP+Whp,p.NFI(3),PurpleTotalUp)
	WA78	:= SUMpr(CSP+Whp,WA75,WA76,WA77)
	WA7		:= SUBpr(CSP+Whp,WA6,WA78)
	WhiteAreaPeakSeedC := TruncSeed(WA7)
	fmt.Println("WhiteAreaPeakSeed is:       ", WhiteAreaPeakSeedC)
	fmt.Println("========================================================")
	//
	//
	//Computing Final Seeds
	//
	y,z,w = ASApr(CSP+Whp,p.NFI(56),p.NFI(1),p.NFI(90))
	BaseAreaSeedC 	:= TruncSeed(w)
	BaseHeightSeedC	:= TruncSeed(y)
	CPAreaRatio 	:= DIVpr(CSP,CamelReward,TotalArea)
	CPAreaRatioSeedC := TruncSeed(CPAreaRatio)

	fmt.Println("Primary Decimal Seeds have been computed")
	fmt.Println("Seed1st, ", BaseAreaSeedC)
	fmt.Println("Seed2nd, ", BaseHeightSeedC)
	fmt.Println("Seed3nd, ", CPAreaRatioSeedC)
	fmt.Println("========================================================")

	elapsed := time.Since(start)
	fmt.Println("Computing took", elapsed)
}