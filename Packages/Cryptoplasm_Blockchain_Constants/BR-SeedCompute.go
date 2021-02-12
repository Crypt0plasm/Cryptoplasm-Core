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
	PurpleTotalUp := ADDx(CSP+Pp,PurpleBaseUp, PurpleUp)

	PurpleUpHeightSeedC := TruncSeed(PurpleUpHeight)
	fmt.Println("PurpleUpHeightSeed is:         ", PurpleUpHeightSeedC)

	PurpleUpAreaSeedC := TruncSeed(PurpleTotalUp)
	fmt.Println("PurpleUpAreaSeed is:        ", PurpleUpAreaSeedC)
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
	PurpleTotalDown := ADDx(CSP+Pp,PurpleBaseDown, PurpleDown)

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
	IndigoPurpleSideLeft := SUBx(CSP+Ip,y, MULx(CSP+Ip,PurpleSaveUp, p.NFI(6)))
	IndigoSave01 := z

	y, z, w = ASApr(CSP+Ip,p.NFI(104), IndigoPurpleSideLeft, p.NFI(1))
	IndigoUp = ADDx(CSP+Ip,IndigoUp, w)
	IndigoSave02 := z
	IndigoPurpleSideRight := ADDx(CSP+Ip,IndigoSave01, IndigoSave02)

	y, z, w = ASApr(CSP+Ip,p.NFI(105), IndigoPurpleSideRight, p.NFI(7))
	IndigoUp = ADDx(CSP+Ip,IndigoUp, w)
	IndigoUp = ADDx(CSP+Ip,IndigoUp, MULx(CSP+Ip,PurpleUp, p.NFI(6)))
	IndigoTotalUp := ADDx(CSP+Ip,IndigoBaseUp, IndigoUp)

	IndigoUpHeightPeak := SUBx(CSP+Ip,IndigoUpHeight, MULx(CSP+Ip,PurpleUpHeight, p.NFI(6)))
	IndigoUpHeightPeakSeedC := TruncSeed(IndigoUpHeightPeak)
	fmt.Println("IndigoUpHeightPeakSeed is:     ", IndigoUpHeightPeakSeedC)
	//fmt.Println("Purple este",Purple)
	IUA := ADDx(CSP+Ip,MULx(CSP+Ip,p.NFI(6), PurpleTotalUp), PRDx(CSP+Ip,Purple,p.NFI(21),PurpleUpHeight))

	IndigoUpAreaPeak := SUBx(CSP+Ip,IndigoTotalUp, IUA)
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
	IndigoDown = ADDx(CSP+Ip,IndigoDown, w)
	IndigoSave04 := z
	IndigoPurpleSideRight = SUBx(CSP+Ip,ADDx(CSP+Ip,IndigoSave04, IndigoSave03), MULx(CSP+Ip,p.NFI(6), PurpleSaveDown))

	y, z, w = ASApr(CSP+Ip,p.NFI(7), IndigoPurpleSideRight, p.NFI(105))
	IndigoDown = ADDx(CSP+Ip,IndigoDown, w)
	IndigoDown = ADDx(CSP+Ip,IndigoDown, MULx(CSP+Ip,p.NFI(6), PurpleDown))
	IndigoTotalDown := ADDx(CSP+Ip,IndigoBaseDown, IndigoDown)

	IndigoDownHeightPeak := SUBx(CSP+Ip,MULx(CSP+Ip,p.NFI(6), PurpleDownHeight), IndigoDownHeight)
	IndigoDownHeightPeakSeedC := TruncSeed(IndigoDownHeightPeak)
	fmt.Println("IndigoDownHeightPeakSeed is:   ", IndigoDownHeightPeakSeedC)

	IDA1 := MULx(CSP+Ip,PurpleDownHeight, p.NFI(21))
	IDA2 := SUBx(CSP+Ip,IDA1, IndigoDownHeightPeak)
	IDA3 := MULx(CSP+Ip,Purple, IDA2)
	IDA4 := MULx(CSP+Ip,PurpleTotalDown, p.NFI(6))
	IDA5 := ADDx(CSP+Ip,IDA3, IDA4)
	IndigoDownAreaPeak := SUBx(CSP+Ip,IndigoTotalDown, IDA5)
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
	BlueIndigoSideLeft := SUBx(CSP+Bp,y,MULx(CSP+Bp,p.NFI(6),IndigoSaveUp))
	BlueSave01 := z

	y,z,w = ASApr(CSP+Bp,p.NFI(95),BlueIndigoSideLeft,p.NFI(2))
	BlueUp = ADDx(CSP+Bp,BlueUp,w)
	BlueSave02 := z
	BlueSave03 := y
	BlueIndigoSideRight := ADDx(CSP+Bp,BlueSave01,BlueSave02)

	y,z,w = ASApr(CSP+Bp,p.NFI(97),BlueIndigoSideRight,p.NFI(7))
	BlueUp = ADDx(CSP+Bp,BlueUp,w)
	BlueSave04 := z
	BlueSave05 := y
	IndigoPurpleSideLeft = SUBx(CSP+Bp,ADDx(CSP+Bp,BlueSave03,BlueSave04),MULx(CSP+Bp,PurpleSaveUp,p.NFI(4)))

	y,z,w = ASApr(CSP+Bp,p.NFI(104),IndigoPurpleSideLeft,p.NFI(1))
	BlueUp = ADDx(CSP+Bp,BlueUp,w)
	BlueSave06 := z
	IndigoPurpleSideRight = SUBx(CSP+Bp,ADDx(CSP+Bp,BlueSave05,BlueSave06),MULx(CSP+Bp,PurpleSaveDown,p.NFI(2)))

	y,z,w = ASApr(CSP+Bp,p.NFI(105),IndigoPurpleSideRight,p.NFI(7))
	BlueUp = ADDx(CSP+Bp,BlueUp,w)
	BU1 	:= MULx(CSP+Bp,p.NFI(6),IndigoUp)
	BU2 	:= MULx(CSP+Bp,p.NFI(4),PurpleUp)
	BU3 	:= MULx(CSP+Bp,p.NFI(2),PurpleDown)
	BlueUp = SUMx(CSP+Bp,BlueUp,BU1,BU2,BU3)
	BlueTotalUp := ADDx(CSP+Bp,BlueBaseUp,BlueUp)

	BUH11 	:= MULx(CSP+Bp,p.NFI(6),IndigoUpHeight)
	BUH1 	:= SUBx(CSP+Bp,BlueUpHeight, BUH11)
	BUH21 	:= MULx(CSP+Bp,p.NFI(2),PurpleDownHeight)
	BUH22 	:= ADDx(CSP+Bp,BUH21,BUH1)
	BUH23	:= MULx(CSP+Bp,p.NFI(4),PurpleUpHeight)
	BUH2 	:= SUBx(CSP+Bp,BUH23,BUH22)
	BlueUpHeightPeakSeedC := TruncSeed(BUH2)
	fmt.Println("BlueUpHeightPeakSeed is:       ", BlueUpHeightPeakSeedC)


	BUA11 	:= PRDx(CSP+Bp,Indigo,IndigoUpHeight,p.NFI(21))
	BUA12 	:= MULx(CSP+Bp,IndigoTotalUp,p.NFI(6))
	BUA13	:= ADDx(CSP+Bp,BUA11,BUA12)
	BUA1	:= SUBx(CSP+Bp,BlueTotalUp,BUA13)

	BUA21	:= MULx(CSP+Bp,BUH1,p.NFI(3))
	BUA22	:= MULx(CSP+Bp,PurpleDownHeight,p.NFI(3))
	BUA23	:= MULx(CSP+Bp,PurpleUpHeight,p.NFI(6))
	BUA24	:= SUMx(CSP+Bp,BUA21,BUA22,BUA23)
	BUA25	:= MULx(CSP+Bp,Purple,BUA24)
	BUA26	:= MULx(CSP+Bp,PurpleTotalDown,p.NFI(2))
	BUA27	:= MULx(CSP+Bp,PurpleTotalUp,p.NFI(4))
	BUA2	:= DIFx(CSP+Bp,BUA1,BUA25,BUA26,BUA27)
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
	BlueDown = ADDx(CSP+Bp,BlueDown,w)
	BlueSave08 := z
	BlueSave09 := y
	BlueIndigoSideRight = SUBx(CSP+Bp,ADDx(CSP+Bp,BlueSave07,BlueSave08),MULx(CSP+Bp,IndigoSaveDown,p.NFI(6)))

	y,z,w = ASApr(CSP+Bp,p.NFI(7),BlueIndigoSideRight,p.NFI(97))
	BlueDown = ADDx(CSP+Bp,BlueDown,w)
	BlueSave10 := y
	BlueSave11 := z
	IndigoPurpleSideLeft = SUBx(CSP+Bp,ADDx(CSP+Bp,BlueSave09,BlueSave10),MULx(CSP+Bp,PurpleSaveUp,p.NFI(4)))

	y,z,w = ASApr(CSP+Bp,p.NFI(104),IndigoPurpleSideLeft,p.NFI(1))
	BlueDown = ADDx(CSP+Bp,BlueDown,w)
	BlueSave12 := z
	IndigoPurpleSideRight = SUBx(CSP+Bp,ADDx(CSP+Bp,BlueSave11,BlueSave12),MULx(CSP+Bp,PurpleSaveDown,p.NFI(2)))

	y,z,w = ASApr(CSP+Bp,p.NFI(7),IndigoPurpleSideRight,p.NFI(105))
	BlueDown = ADDx(CSP+Bp,BlueDown,w)
	BD1		:= MULx(CSP+Bp,p.NFI(6),IndigoDown)
	BD2 	:= MULx(CSP+Bp,p.NFI(4),PurpleUp)
	BD3 	:= MULx(CSP+Bp,p.NFI(2),PurpleDown)
	BlueDown = SUMx(CSP+Bp,BlueDown,BD1,BD2,BD3)
	BlueTotalDown := ADDx(CSP+Bp,BlueBaseDown,BlueDown)

	BDH1	:= SUBx(CSP+Bp,MULx(CSP+Bp,p.NFI(6),IndigoDownHeight),BlueDownHeight)
	BDH21	:= MULx(CSP+Bp,p.NFI(2),PurpleDownHeight)
	BDH22	:= MULx(CSP+Bp,p.NFI(4),PurpleUpHeight)
	BDH23	:= SUBx(CSP+Bp,BDH21,BDH22)
	BDH2	:= ADDx(CSP+Bp,BDH1,BDH23)
	BlueDownHeightPeakSeedC := TruncSeed(BDH2)
	fmt.Println("BlueDownHeightPeakSeed is:      ", BlueDownHeightPeakSeedC)
	
	BDA11 	:= MULx(CSP+Bp,p.NFI(15),IndigoDownHeight)
	BDA12 	:= ADDx(CSP+Bp,BDA11,BlueDownHeight)
	BDA13 	:= MULx(CSP+Bp,BDA12,Indigo)
	BDA14 	:= MULx(CSP+Bp,p.NFI(6),IndigoTotalDown)
	BDA15 	:= ADDx(CSP+Bp,BDA13,BDA14)
	BDA1 	:= SUBx(CSP+Bp,BlueTotalDown,BDA15)

	BDA21 := MULx(CSP+Bp,p.NFI(2),BDH1)
	BDA22 := MULx(CSP+Bp,p.NFI(10),PurpleUpHeight)
	BDA23 := SUMx(CSP+Bp,BDA21,BDA22,PurpleDownHeight)
	BDA24 := MULx(CSP+Bp,Purple,BDA23)
	BDA25 := MULx(CSP+Bp,p.NFI(2),PurpleTotalDown)
	BDA26 := MULx(CSP+Bp,p.NFI(4),PurpleTotalUp)
	BDA27 := SUMx(CSP+Bp,BDA24,BDA25,BDA26)
	BDA2  := SUBx(CSP+Bp,BDA1,BDA27)
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
	GreenBlueSideLeft := SUBx(CSP+Gp,y,MULx(CSP+Gp,p.NFI(6),BlueSaveUp))
	GreenSave01 := z

	y,z,w = ASApr(CSP+Gp,p.NFI(85),GreenBlueSideLeft,p.NFI(3))
	GreenUp = ADDx(CSP+Gp,GreenUp,w)
	GreenSave02 := z
	GreenSave03 := y
	GreenBlueSideRight := ADDx(CSP+Gp,GreenSave02,GreenSave01)

	y,z,w = ASApr(CSP+Gp,p.NFI(7),GreenBlueSideRight,p.NFI(88))
	GreenUp = ADDx(CSP+Gp,GreenUp,w)
	GreenSave04 := y
	GreenSave05 := z
	BlueIndigoSideLeft = SUBx(CSP+Gp,ADDx(CSP+Gp,GreenSave03,GreenSave04),MULx(CSP+Gp,p.NFI(3),IndigoSaveUp))

	y,z,w = ASApr(CSP+Gp,p.NFI(95),BlueIndigoSideLeft,p.NFI(2))
	GreenUp = ADDx(CSP+Gp,GreenUp,w)
	GreenSave06 := z
	GreenSave07 := y
	BlueIndigoSideRight = SUBx(CSP+Gp,ADDx(CSP+Gp,GreenSave05,GreenSave06),MULx(CSP+Gp,p.NFI(3),IndigoSaveDown))

	y,z,w = ASApr(CSP+Gp,p.NFI(7),BlueIndigoSideRight,p.NFI(97))
	GreenUp = ADDx(CSP+Gp,GreenUp,w)
	GreenSave08 := y
	GreenSave09 := z
	IndigoPurpleSideLeft = SUBx(CSP+Gp,ADDx(CSP+Gp,GreenSave07,GreenSave08),PurpleSaveUp)

	y,z,w = ASApr(CSP+Gp,p.NFI(104),IndigoPurpleSideLeft,p.NFI(1))
	GreenUp = ADDx(CSP+Gp,GreenUp,w)
	GreenSave10 := z
	IndigoPurpleSideRight = SUBx(CSP+Gp,ADDx(CSP+Gp,GreenSave09,GreenSave10),MULx(CSP+Gp,p.NFI(5),PurpleSaveDown))

	y,z,w = ASApr(CSP+Gp,p.NFI(7),IndigoPurpleSideRight,p.NFI(105))
	GreenUp = ADDx(CSP+Gp,GreenUp,w)
	GU1		:= MULx(CSP+Gp,p.NFI(6),BlueUp)
	GU2		:= MULx(CSP+Gp,p.NFI(3),IndigoUp)
	GU3 	:= MULx(CSP+Gp,p.NFI(5),PurpleDown)
	GU4 	:= MULx(CSP+Gp,p.NFI(3),IndigoDown)
	GreenUp = SUMx(CSP+Gp,GreenUp,PurpleUp,GU1,GU2,GU3,GU4)
	GreenTotalUp := ADDx(CSP+Gp,GreenBaseUp,GreenUp)

	GUH1  	:= SUBx(CSP+Gp,GreenUpHeight,MULx(CSP+Gp,p.NFI(6),BlueUpHeight))
	GUH21 	:= MULx(CSP+Gp,p.NFI(3),IndigoDownHeight)
	GUH22 	:= ADDx(CSP+Gp,GUH1,GUH21)
	GUH23 	:= MULx(CSP+Gp,p.NFI(3),IndigoUpHeight)
	GUH2  	:= SUBx(CSP+Gp,GUH23,GUH22)
	GUH3  	:= SUBx(CSP+Gp,MULx(CSP+Gp,p.NFI(5),PurpleDownHeight),ADDx(CSP+Gp,GUH2,PurpleUpHeight))
	GreenUpHeightPeakSeedC := TruncSeed(GUH3)
	fmt.Println("GreenUpHeightPeakSeed is:      ", GreenUpHeightPeakSeedC)

	GUA11 	:= PRDx(CSP+Gp,Blue,p.NFI(21),BlueUpHeight)
	GUA12	:= MULx(CSP+Gp,p.NFI(6),BlueTotalUp)
	GUA13	:= ADDx(CSP+Gp,GUA11,GUA12)
	GUA1	:= SUBx(CSP+Gp,GreenTotalUp,GUA13)

	GUA21 	:= MULx(CSP+Gp,p.NFI(4),GUH1)
	GUA22 	:= MULx(CSP+Gp,p.NFI(6),IndigoDownHeight)
	GUA23 	:= MULx(CSP+Gp,p.NFI(3),IndigoUpHeight)
	GUA24 	:= SUMx(CSP+Gp,GUA21,GUA22,GUA23)
	GUA25 	:= MULx(CSP+Gp,Indigo,GUA24)
	GUA26 	:= MULx(CSP+Gp,p.NFI(3),ADDx(CSP+Gp,IndigoTotalDown,IndigoTotalUp))
	GUA27 	:= ADDx(CSP+Gp,GUA26,GUA25)
	GUA2 	:= SUBx(CSP+Gp,GUA1,GUA27)

	GUA31 	:= MULx(CSP+Gp,p.NFI(2),GUH2)
	GUA32 	:= MULx(CSP+Gp,p.NFI(10),PurpleDownHeight)
	GUA33 	:= SUMx(CSP+Gp,GUA31,GUA32,PurpleUpHeight)
	GUA34 	:= MULx(CSP+Gp,Purple,GUA33)
	GUA35 	:= MULx(CSP+Gp,p.NFI(5),PurpleTotalDown)
	GUA36 	:= SUMx(CSP+Gp,PurpleTotalUp,GUA34,GUA35)
	GUA3 	:= SUBx(CSP+Gp,GUA2,GUA36)
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
	GreenDown = ADDx(CSP+Gp,GreenDown,w)
	GreenSave12 := z
	GreenSave13 := y
	GreenBlueSideRight = SUBx(CSP+Gp,ADDx(CSP+Gp,GreenSave11,GreenSave12),MULx(CSP+Gp,p.NFI(6),BlueSaveDown))

	y,z,w = ASApr(CSP+Gp,p.NFI(7),GreenBlueSideRight,p.NFI(88))
	GreenDown = ADDx(CSP+Gp,GreenDown,w)
	GreenSave14 := y
	GreenSave15 := z
	BlueIndigoSideLeft = SUBx(CSP+Gp,ADDx(CSP+Gp,GreenSave13,GreenSave14),MULx(CSP+Gp,p.NFI(4),IndigoSaveUp))

	y,z,w = ASApr(CSP+Gp,p.NFI(95),BlueIndigoSideLeft,p.NFI(2))
	GreenDown = ADDx(CSP+Gp,GreenDown,w)
	GreenSave16 := z
	GreenSave17 := y
	BlueIndigoSideRight = SUBx(CSP+Gp,ADDx(CSP+Gp,GreenSave15,GreenSave16),MULx(CSP+Gp,p.NFI(2),IndigoSaveDown))

	y,z,w = ASApr(CSP+Gp,p.NFI(7),BlueIndigoSideRight,p.NFI(97))
	GreenDown = ADDx(CSP+Gp,GreenDown,w)
	GreenSave18 := y
	GreenSave19 := z
	IndigoPurpleSideLeft = SUBx(CSP+Gp,ADDx(CSP+Gp,GreenSave17,GreenSave18),MULx(CSP+Gp,p.NFI(2),PurpleSaveUp))

	y,z,w = ASApr(CSP+Gp,p.NFI(104),IndigoPurpleSideLeft,p.NFI(1))
	GreenDown = ADDx(CSP+Gp,GreenDown,w)
	GreenSave20 := z
	IndigoPurpleSideRight = SUBx(CSP+Gp,ADDx(CSP+Gp,GreenSave19,GreenSave20),MULx(CSP+Gp,p.NFI(4),PurpleSaveDown))

	y,z,w = ASApr(CSP+Gp,p.NFI(7),IndigoPurpleSideRight,p.NFI(105))
	GreenDown = ADDx(CSP+Gp,GreenDown,w)
	GD1 	:= MULx(CSP+Gp,p.NFI(4),IndigoUp)
	GD2 	:= MULx(CSP+Gp,p.NFI(2),PurpleUp)
	GD3 	:= MULx(CSP+Gp,p.NFI(4),PurpleDown)
	GD4 	:= MULx(CSP+Gp,p.NFI(2),IndigoDown)
	GD5 	:= MULx(CSP+Gp,p.NFI(6),BlueDown)
	GreenDown = SUMx(CSP+Gp,GreenDown,GD1,GD2,GD3,GD4,GD5)
	GreenTotalDown := ADDx(CSP+Gp,GreenBaseDown,GreenDown)

	GDH1 	:= SUBx(CSP+Gp,MULx(CSP+Gp,p.NFI(6),BlueDownHeight),GreenDownHeight)
	GDH21 	:= ADDx(CSP+Gp,MULx(CSP+Gp,p.NFI(2),IndigoDownHeight),GDH1)
	GDH2 	:= SUBx(CSP+Gp,MULx(CSP+Gp,p.NFI(4),IndigoUpHeight),GDH21)
	GDH31 	:= ADDx(CSP+Gp,MULx(CSP+Gp,p.NFI(2),PurpleUpHeight),GDH2)
	GreenDownH3 := SUBx(CSP+Gp,MULx(CSP+Gp,p.NFI(4),PurpleDownHeight),GDH31)
	GreenDownHeightPeakSeedC := TruncSeed(GreenDownH3)
	fmt.Println("GreenDownHeightPeakSeed is:    ", GreenDownHeightPeakSeedC)

	GDA11 	:= ADDx(CSP+Gp,MULx(CSP+Gp,p.NFI(15),BlueDownHeight),GreenDownHeight)
	GDA12 	:= MULx(CSP+Gp,Blue,GDA11)
	GDA13 	:= MULx(CSP+Gp,BlueTotalDown,p.NFI(6))
	GDA14 	:= ADDx(CSP+Gp,GDA12,GDA13)
	GA1 	:= SUBx(CSP+Gp,GreenTotalDown,GDA14)

	GDA21 	:= MULx(CSP+Gp,GDH1,p.NFI(3))
	GDA22 	:= MULx(CSP+Gp,IndigoDownHeight,p.NFI(3))
	GDA23 	:= MULx(CSP+Gp,IndigoUpHeight,p.NFI(6))
	GDA24 	:= SUMx(CSP+Gp,GDA21,GDA22,GDA23)
	GDA25 	:= MULx(CSP+Gp,Indigo,GDA24)
	GDA26 	:= MULx(CSP+Gp,IndigoTotalDown,p.NFI(2))
	GDA27 	:= MULx(CSP+Gp,IndigoTotalUp,p.NFI(4))
	GDA28 	:= SUMx(CSP+Gp,GDA25,GDA26,GDA27)
	GA2 	:= SUBx(CSP+Gp,GA1,GDA28)

	GDA31 	:= MULx(CSP+Gp,GDH2,p.NFI(3))
	GDA32 	:= MULx(CSP+Gp,PurpleUpHeight,p.NFI(3))
	GDA33 	:= MULx(CSP+Gp,PurpleDownHeight,p.NFI(6))
	GDA34 	:= SUMx(CSP+Gp,GDA31,GDA32,GDA33)
	GDA35 	:= MULx(CSP+Gp,Purple,GDA34)
	GDA36 	:= MULx(CSP+Gp,PurpleTotalDown,p.NFI(4))
	GDA37 	:= MULx(CSP+Gp,PurpleTotalUp,p.NFI(2))
	GDA38 	:= SUMx(CSP+Gp,GDA35,GDA36,GDA37)
	GDA3 	:= SUBx(CSP+Gp,GA2,GDA38)
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
	YellowGreenSideLeft := SUBx(CSP+YUp,y,MULx(CSP+YUp,p.NFI(6),GreenSaveUp))
	YellowSave01 := z

	y,z,w = ASApr(CSP+YUp,p.NFI(74),YellowGreenSideLeft,p.NFI(4))
	YellowUp = ADDx(CSP+YUp,YellowUp,w)
	YellowSave02 := z
	YellowSave03 := y
	YellowGreenSideRight := ADDx(CSP+YUp,YellowSave02,YellowSave01)

	y,z,w = ASApr(CSP+YUp,p.NFI(7),YellowGreenSideRight,p.NFI(78))
	YellowUp = ADDx(CSP+YUp,YellowUp,w)
	YellowSave04 := y
	YellowSave05 := z
	GreenBlueSideLeft = SUBx(CSP+YUp,ADDx(CSP+YUp,YellowSave03,YellowSave04),MULx(CSP+YUp,BlueSaveUp,p.NFI(2)))

	y,z,w = ASApr(CSP+YUp,p.NFI(85),GreenBlueSideLeft,p.NFI(3))
	YellowUp = ADDx(CSP+YUp,YellowUp,w)
	YellowSave06 := z
	YellowSave07 := y
	GreenBlueSideRight = SUBx(CSP+YUp,ADDx(CSP+YUp,YellowSave05,YellowSave06),MULx(CSP+YUp,BlueSaveDown,p.NFI(4)))

	y,z,w = ASApr(CSP+YUp,p.NFI(7),GreenBlueSideRight,p.NFI(88))
	YellowUp = ADDx(CSP+YUp,YellowUp,w)
	YellowSave08 := y
	YellowSave09 := z
	BlueIndigoSideLeft = SUBx(CSP+YUp,ADDx(CSP+YUp,YellowSave07,YellowSave08),IndigoSaveUp)

	y,z,w = ASApr(CSP+YUp,p.NFI(95),BlueIndigoSideLeft,p.NFI(2))
	YellowUp = ADDx(CSP+YUp,YellowUp,w)
	YellowSave10 := z
	YellowSave11 := y
	BlueIndigoSideRight = SUBx(CSP+YUp,ADDx(CSP+YUp,YellowSave09,YellowSave10),MULx(CSP+YUp,IndigoSaveDown,p.NFI(5)))

	y,z,w = ASApr(CSP+YUp,p.NFI(7),BlueIndigoSideRight,p.NFI(97))
	YellowUp = ADDx(CSP+YUp,YellowUp,w)
	YellowSave12 := y
	YellowSave13 := z
	IndigoPurpleSideLeft = SUBx(CSP+YUp,ADDx(CSP+YUp,YellowSave11,YellowSave12),MULx(CSP+YUp,PurpleSaveUp,p.NFI(2)))

	y,z,w = ASApr(CSP+YUp,p.NFI(104),IndigoPurpleSideLeft,p.NFI(1))
	YellowUp = ADDx(CSP+YUp,YellowUp,w)
	YellowSave14 := z
	IndigoPurpleSideRight = SUBx(CSP+YUp,ADDx(CSP+YUp,YellowSave13,YellowSave14),MULx(CSP+YUp,PurpleSaveDown,p.NFI(4)))

	y,z,w = ASApr(CSP+YUp,p.NFI(7),IndigoPurpleSideRight,p.NFI(105))
	YellowUp = ADDx(CSP+YUp,YellowUp,w)
	YU1 	:= MULx(CSP+YUp,GreenUp,p.NFI(6))
	YU2 	:= MULx(CSP+YUp,BlueUp,p.NFI(2))
	YU3 	:= MULx(CSP+YUp,PurpleUp,p.NFI(2))
	YU4 	:= MULx(CSP+YUp,PurpleDown,p.NFI(4))
	YU5 	:= MULx(CSP+YUp,IndigoDown,p.NFI(5))
	YU6 	:= MULx(CSP+YUp,BlueDown,p.NFI(4))
	YellowUp = SUMx(CSP+YUp,YellowUp,IndigoUp,YU1,YU2,YU3,YU4,YU5,YU6)
	YellowTotalUp := ADDx(CSP+YUp,YellowBaseUp,YellowUp)

	YUH1 	:= SUBx(CSP+YUp,MULx(CSP+YUp,GreenUpHeight,p.NFI(6)),YellowUpHeight)
	YUH2 	:= SUBx(CSP+YUp,ADDx(CSP+YUp,MULx(CSP+YUp,BlueUpHeight,p.NFI(2)),YUH1),MULx(CSP+YUp,BlueDownHeight,p.NFI(4)))
	YUH3 	:= SUBx(CSP+YUp,ADDx(CSP+YUp,YUH2,IndigoUpHeight),MULx(CSP+YUp,IndigoDownHeight,p.NFI(5)))
	YUH4 	:= SUBx(CSP+YUp,MULx(CSP+YUp,PurpleDownHeight,p.NFI(4)),ADDx(CSP+YUp,YUH3,MULx(CSP+YUp,PurpleUpHeight,p.NFI(2))))
	YellowUpHeightPeakSeedC := TruncSeed(YUH4)
	fmt.Println("YellowUpHeightPeakSeed is:     ", YellowUpHeightPeakSeedC)

	YUA11 	:= ADDx(CSP+YUp,MULx(CSP+YUp,GreenUpHeight,p.NFI(15)),YellowUpHeight)
	YUA12 	:= ADDx(CSP+YUp,MULx(CSP+YUp,GreenTotalUp,p.NFI(6)),MULx(CSP+YUp,YUA11,Green))
	YUA1 	:= SUBx(CSP+YUp,YellowTotalUp,YUA12)

	YUA21 	:= SUMx(CSP+YUp,BlueUpHeight,MULx(CSP+YUp,BlueDownHeight,p.NFI(10)),MULx(CSP+YUp,YUH1,p.NFI(2)))
	YUA22 	:= SUMx(CSP+YUp,MULx(CSP+YUp,Blue,YUA21),MULx(CSP+YUp,BlueTotalDown,p.NFI(4)),MULx(CSP+YUp,BlueTotalUp,p.NFI(2)))
	YUA2 	:= SUBx(CSP+YUp,YUA1,YUA22)

	YUA31 	:= ADDx(CSP+YUp,MULx(CSP+YUp,IndigoDownHeight,p.NFI(15)),YUH2)
	YUA32 	:= SUMx(CSP+YUp,MULx(CSP+YUp,Indigo,YUA31),MULx(CSP+YUp,IndigoTotalDown,p.NFI(5)),IndigoTotalUp)
	YUA3 	:= SUBx(CSP+YUp,YUA2,YUA32)

	YUA41 	:= SUMx(CSP+YUp,MULx(CSP+YUp,PurpleDownHeight,p.NFI(6)),MULx(CSP+YUp,PurpleUpHeight,p.NFI(3)),MULx(CSP+YUp,YUH3,p.NFI(3)))
	YUA42 	:= SUMx(CSP+YUp,MULx(CSP+YUp,Purple,YUA41),MULx(CSP+YUp,PurpleTotalDown,p.NFI(4)),MULx(CSP+YUp,PurpleTotalUp,p.NFI(2)))
	YUA4 	:= SUBx(CSP+YUp,YUA3,YUA42)
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
	YellowDown = ADDx(CSP+YDp,YellowDown,w)
	YellowSave16 := z
	YellowSave17 := y
	YellowGreenSideRight = SUBx(CSP+YDp,ADDx(CSP+YDp,YellowSave15,YellowSave16),MULx(CSP+YDp,GreenSaveDown,p.NFI(6)))

	y,z,w = ASApr(CSP+YDp,p.NFI(7),YellowGreenSideRight,p.NFI(78))
	YellowDown = ADDx(CSP+YDp,YellowDown,w)
	YellowSave18 := y
	YellowSave19 := z
	GreenBlueSideLeft = SUBx(CSP+YDp,ADDx(CSP+YDp,YellowSave17,YellowSave18),MULx(CSP+YDp,BlueSaveUp,p.NFI(4)))

	y,z,w = ASApr(CSP+YDp,p.NFI(85),GreenBlueSideLeft,p.NFI(3))
	YellowDown = ADDx(CSP+YDp,YellowDown,w)
	YellowSave20 := z
	YellowSave21 := y
	GreenBlueSideRight = SUBx(CSP+YDp,ADDx(CSP+YDp,YellowSave19,YellowSave20),MULx(CSP+YDp,BlueSaveDown,p.NFI(2)))

	y,z,w = ASApr(CSP+YDp,p.NFI(7),GreenBlueSideRight,p.NFI(88))
	YellowDown = ADDx(CSP+YDp,YellowDown,w)
	YellowSave22 := y
	YellowSave23 := z
	BlueIndigoSideLeft = SUBx(CSP+YDp,ADDx(CSP+YDp,YellowSave21,YellowSave22),MULx(CSP+YDp,IndigoSaveUp,p.NFI(3)))

	y,z,w = ASApr(CSP+YDp,p.NFI(95),BlueIndigoSideLeft,p.NFI(2))
	YellowDown = ADDx(CSP+YDp,YellowDown,w)
	YellowSave24 := z
	YellowSave25 := y
	BlueIndigoSideRight = SUBx(CSP+YDp,ADDx(CSP+YDp,YellowSave23,YellowSave24),MULx(CSP+YDp,IndigoSaveDown,p.NFI(3)))

	y,z,w = ASApr(CSP+YDp,p.NFI(7),BlueIndigoSideRight,p.NFI(97))
	YellowDown = ADDx(CSP+YUp,YellowDown,w)
	YellowSave26 := y
	YellowSave27 := z
	IndigoPurpleSideLeft = SUBx(CSP+YDp,ADDx(CSP+YDp,YellowSave25,YellowSave26),MULx(CSP+YDp,PurpleSaveUp,p.NFI(2)))

	y,z,w = ASApr(CSP+YDp,p.NFI(104),IndigoPurpleSideLeft,p.NFI(1))
	YellowDown = ADDx(CSP+YDp,YellowDown,w)
	YellowSave28 := z
	IndigoPurpleSideRight = SUBx(CSP+YDp,ADDx(CSP+YDp,YellowSave27,YellowSave28),MULx(CSP+YDp,PurpleSaveDown,p.NFI(4)))

	y,z,w = ASApr(CSP+YDp,p.NFI(7),IndigoPurpleSideRight,p.NFI(105))
	YellowDown = ADDx(CSP+YDp,YellowDown,w)
	YD1 	:= MULx(CSP+YDp,BlueUp,p.NFI(4))
	YD2 	:= MULx(CSP+YDp,IndigoUp,p.NFI(3))
	YD3 	:= MULx(CSP+YDp,PurpleUp,p.NFI(2))
	YD4 	:= MULx(CSP+YDp,PurpleDown,p.NFI(4))
	YD5 	:= MULx(CSP+YDp,IndigoDown,p.NFI(3))
	YD6 	:= MULx(CSP+YDp,BlueDown,p.NFI(2))
	YD7 	:= MULx(CSP+YDp,GreenDown,p.NFI(6))
	YellowDown = SUMx(CSP+YDp,YellowDown,YD1,YD2,YD3,YD4,YD5,YD6,YD7)
	YellowTotalDown := ADDx(CSP+YDp,YellowBaseDown,YellowDown)

	YDH1 	:= SUBx(CSP+YDp,YellowDownHeight,MULx(CSP+YDp,GreenDownHeight,p.NFI(5)))
	YDH21 	:= ADDx(CSP+YDp,GreenDownHeight,MULx(CSP+YDp,BlueDownHeight,p.NFI(2)))
	YDH22 	:= ADDx(CSP+YDp,YDH1,MULx(CSP+YDp,BlueUpHeight,p.NFI(4)))
	YDH2 	:= SUBx(CSP+YDp,YDH21,YDH22)

	YDH31 	:= ADDx(CSP+YDp,YDH2,MULx(CSP+YDp,IndigoDownHeight,p.NFI(3)))
	YDH3 	:= SUBx(CSP+YDp,MULx(CSP+YDp,IndigoUpHeight,p.NFI(3)),YDH31)

	YDH41 	:= ADDx(CSP+YDp,YDH3,MULx(CSP+YDp,PurpleUpHeight,p.NFI(2)))
	YDH4 	:= SUBx(CSP+YDp,MULx(CSP+YDp,PurpleDownHeight,p.NFI(4)),YDH41)
	YellowDownHeightPeakSeedC := TruncSeed(YDH4)
	fmt.Println("YellowDownHeightPeakSeed is:   ", YellowDownHeightPeakSeedC)

	YellowDownArea101 := MULx(CSP+YDp,MULx(CSP+YDp,Green,GreenDownHeight),p.NFI(20))
	YellowDownArea102 := ADDx(CSP+YDp,MULx(CSP+YDp,GreenTotalDown,p.NFI(6)),YellowDownArea101)
	YellowDownArea1 := SUBx(CSP+YDp,YellowTotalDown,YellowDownArea102)

	YDA21 	:= MULx(CSP+YDp,YDH1,p.NFI(5))
	YDA22 	:= MULx(CSP+YDp,GreenDownHeight,p.NFI(2))
	YDA23 	:= MULx(CSP+YDp,BlueUpHeight,p.NFI(10))
	YDA24 	:= SUMx(CSP+YDp,BlueDownHeight,YDA21,YDA22,YDA23)
	YDA25 	:= MULx(CSP+YDp,Blue,YDA24)
	YDA26 	:= MULx(CSP+YDp,BlueTotalDown,p.NFI(2))
	YDA27 	:= MULx(CSP+YDp,BlueTotalUp,p.NFI(4))
	YDA28 	:= SUMx(CSP+YDp,YDA25,YDA26,YDA27)
	YDA2 	:= SUBx(CSP+YDp,YellowDownArea1,YDA28)

	YDA31 	:= MULx(CSP+YDp,YDH2,p.NFI(4))
	YDA32 	:= MULx(CSP+YDp,IndigoDownHeight,p.NFI(6))
	YDA33 	:= MULx(CSP+YDp,IndigoUpHeight,p.NFI(3))
	YDA34 	:= SUMx(CSP+YDp,YDA31,YDA32,YDA33)
	YDA35 	:= MULx(CSP+YDp,Indigo,YDA34)
	YDA36 	:= MULx(CSP+YDp,ADDx(CSP+YDp,IndigoTotalUp,IndigoTotalDown),p.NFI(3))
	YDA3 	:= SUBx(CSP+YDp,YDA2,ADDx(CSP+YDp,YDA35,YDA36))

	YDA41 	:= MULx(CSP+YDp,PurpleDownHeight,p.NFI(6))
	YDA42 	:= MULx(CSP+YDp,YDH3,p.NFI(3))
	YDA43 	:= MULx(CSP+YDp,PurpleUpHeight,p.NFI(3))
	YDA44 	:= SUMx(CSP+YDp,YDA41,YDA42,YDA43)
	YDA45 	:= MULx(CSP+YDp,Purple,YDA44)
	YDA46 	:= MULx(CSP+YDp,PurpleTotalDown,p.NFI(4))
	YDA47 	:= MULx(CSP+YDp,PurpleTotalUp,p.NFI(2))
	YDA48 	:= SUMx(CSP+YDp,YDA45,YDA46,YDA47)
	YDA4 	:= SUBx(CSP+YDp,YDA3,YDA48)
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
	OrangeYellowSideLeft := SUBx(CSP+Op,y,MULx(CSP+Op,YellowSaveUp,p.NFI(6)))
	OrangeSave01 := z

	y,z,w = ASApr(CSP+Op,p.NFI(62),OrangeYellowSideLeft,p.NFI(5))
	OrangeUp = ADDx(CSP+Op,OrangeUp,w)
	OrangeSave02 := z
	OrangeSave03 := y
	OrangeYellowSideRight := ADDx(CSP+Op,OrangeSave01,OrangeSave02)

	y,z,w = ASApr(CSP+Op,p.NFI(7),OrangeYellowSideRight,p.NFI(67))
	OrangeUp = ADDx(CSP+Op,OrangeUp,w)
	OrangeSave04 := z
	OrangeSave05 := y
	YellowGreenSideRight = SUBx(CSP+Op,OrangeSave04,MULx(CSP+Op,p.NFI(6),GreenSaveDown))

	y,z,w = ASApr(CSP+Op,p.NFI(7),YellowGreenSideRight,p.NFI(74))
	OrangeUp = ADDx(CSP+Op,OrangeUp,w)
	OrangeSave06 := y
	OrangeSave07 := z
	YellowGreenSideLeft = SUMx(CSP+Op,OrangeSave03,OrangeSave05,OrangeSave06)

	y,z,w = ASApr(CSP+Op,p.NFI(81),YellowGreenSideLeft,p.NFI(4))
	OrangeUp = ADDx(CSP+Op,OrangeUp,w)
	OrangeSave08 := y
	OrangeSave09 := z
	GreenBlueSideLeft = SUBx(CSP+Op,OrangeSave08,MULx(CSP+Op,p.NFI(6),BlueSaveUp))

	y,z,w = ASApr(CSP+Op,p.NFI(85),GreenBlueSideLeft,p.NFI(3))
	OrangeUp = ADDx(CSP+Op,OrangeUp,w)
	OrangeSave10 := z
	OrangeSave11 := y
	GreenBlueSideRight =  SUMx(CSP+Op,OrangeSave07,OrangeSave09,OrangeSave10)

	y,z,w = ASApr(CSP+Op,p.NFI(7),GreenBlueSideRight,p.NFI(88))
	OrangeUp = ADDx(CSP+Op,OrangeUp,w)
	OrangeSave12 := y
	OrangeSave13 := z
	BlueIndigoSideLeft = SUBx(CSP+Op,ADDx(CSP+Op,OrangeSave11,OrangeSave12),MULx(CSP+Op,IndigoSaveUp,p.NFI(2)))

	y,z,w = ASApr(CSP+Op,p.NFI(95),BlueIndigoSideLeft,p.NFI(2))
	OrangeUp = ADDx(CSP+Op,OrangeUp,w)
	OrangeSave14 := z
	OrangeSave15 := y
	BlueIndigoSideRight = SUBx(CSP+Op,ADDx(CSP+Op,OrangeSave13,OrangeSave14),MULx(CSP+Op,IndigoSaveDown,p.NFI(4)))

	y,z,w = ASApr(CSP+Op,p.NFI(7),BlueIndigoSideRight,p.NFI(97))
	OrangeUp = ADDx(CSP+Op,OrangeUp,w)
	OrangeSave16 := y
	OrangeSave17 := z
	IndigoPurpleSideLeft = SUBx(CSP+Op,ADDx(CSP+Op,OrangeSave15,OrangeSave16),PurpleSaveUp)

	y,z,w = ASApr(CSP+Op,p.NFI(104),IndigoPurpleSideLeft,p.NFI(1))
	OrangeUp = ADDx(CSP+Op,OrangeUp,w)
	OrangeSave18 := z
	IndigoPurpleSideRight = SUBx(CSP+Op,ADDx(CSP+Op,OrangeSave17,OrangeSave18),MULx(CSP+Op,PurpleSaveDown,p.NFI(5)))

	y,z,w = ASApr(CSP+Op,p.NFI(7),IndigoPurpleSideRight,p.NFI(105))
	OrangeUp = ADDx(CSP+Op,OrangeUp,w)
	OU1 	:= MULx(CSP+Op,p.NFI(6),YellowUp)
	OU2 	:= MULx(CSP+Op,p.NFI(6),BlueUp)
	OU3 	:= MULx(CSP+Op,p.NFI(2),IndigoUp)
	OU4 	:= MULx(CSP+Op,p.NFI(5),PurpleDown)
	OU5 	:= MULx(CSP+Op,p.NFI(4),IndigoDown)
	OU6 	:= MULx(CSP+Op,p.NFI(6),GreenDown)
	OrangeUp = SUMx(CSP+Op,OrangeUp,PurpleUp,OU1,OU2,OU3,OU4,OU5,OU6)
	OrangeTotalUp := ADDx(CSP+Op,OrangeBaseUp,OrangeUp)


	OUH1 	:= SUBx(CSP+Op,MULx(CSP+Op,p.NFI(6),YellowUpHeight),OrangeUpHeight)
	OUH2 	:= SUBx(CSP+Op,MULx(CSP+Op,p.NFI(6),GreenDownHeight),OUH1)
	OUH3 	:= SUBx(CSP+Op,MULx(CSP+Op,p.NFI(6),BlueUpHeight),OUH2)
	OUH4 	:= SUBx(CSP+Op,ADDx(CSP+Op,OUH3,MULx(CSP+Op,p.NFI(2),IndigoUpHeight)),MULx(CSP+Op,p.NFI(4),IndigoDownHeight))
	OUH5 	:= DIFx(CSP+Op,MULx(CSP+Op,p.NFI(5),PurpleDownHeight),OUH4,PurpleUpHeight)
	OrangeUpHeightPeakSeedC := TruncSeed(OUH5)
	fmt.Println("OrangeUpHeightPeakSeed is:     ", OrangeUpHeightPeakSeedC)

	OUA11 	:= ADDx(CSP+Op,MULx(CSP+Op,p.NFI(15),YellowUpHeight),OrangeUpHeight)
	OUA12	:= ADDx(CSP+Op,MULx(CSP+Op,p.NFI(6),YellowTotalUp),MULx(CSP+Op,Yellow,OUA11))
	OUA1	:= SUBx(CSP+Op,OrangeTotalUp,OUA12)

	OUA21	:= ADDx(CSP+Op,MULx(CSP+Op,p.NFI(15),GreenDownHeight),OUH1)
	OUA22	:= ADDx(CSP+Op,MULx(CSP+Op,p.NFI(6),GreenTotalDown),MULx(CSP+Op,Green,OUA21))
	OUA2	:= SUBx(CSP+Op,OUA1,OUA22)

	OUA31	:= ADDx(CSP+Op,MULx(CSP+Op,p.NFI(15),BlueUpHeight),OUH2)
	OUA32	:= ADDx(CSP+Op,MULx(CSP+Op,p.NFI(6),BlueTotalUp),MULx(CSP+Op,Blue,OUA31))
	OUA3	:= SUBx(CSP+Op,OUA2,OUA32)

	OUA41 	:= MULx(CSP+Op,p.NFI(10),IndigoDownHeight)
	OUA42 	:= MULx(CSP+Op,p.NFI(2),OUH3)
	OUA43	:= SUMx(CSP+Op,IndigoUpHeight,OUA41,OUA42)
	OUA44	:= MULx(CSP+Op,Indigo,OUA43)
	OUA45 	:= MULx(CSP+Op,p.NFI(4),IndigoTotalDown)
	OUA46 	:= MULx(CSP+Op,p.NFI(2),IndigoTotalUp)
	OUA47	:= SUMx(CSP+Op,OUA44,OUA45,OUA46)
	OUA4	:= SUBx(CSP+Op,OUA3,OUA47)

	OUA51 	:= MULx(CSP+Op,p.NFI(10),PurpleDownHeight)
	OUA52 	:= MULx(CSP+Op,p.NFI(2),OUH4)
	OUA53	:= SUMx(CSP+Op,PurpleUpHeight,OUA51,OUA52)
	OUA54	:= MULx(CSP+Op,Purple,OUA53)
	OUA55 	:= MULx(CSP+Op,p.NFI(5),PurpleTotalDown)
	OUA56	:= SUMx(CSP+Op,PurpleTotalUp,OUA55,OUA54)
	OUA5	:= SUBx(CSP+Op,OUA4,OUA56)
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
	OrangeDown = ADDx(CSP+Op,OrangeDown,w)
	OrangeSave20 := z
	OrangeSave21 := y
	OrangeYellowSideRight = SUBx(CSP+Op,ADDx(CSP+Op,OrangeSave19,OrangeSave20),MULx(CSP+Op,YellowSaveDown,p.NFI(6)))

	y,z,w = ASApr(CSP+Op,p.NFI(7),OrangeYellowSideRight,p.NFI(67))
	OrangeDown = ADDx(CSP+Op,OrangeDown,w)
	OrangeSave22 := y
	OrangeSave23 := z
	YellowGreenSideRight = SUBx(CSP+Op,ADDx(CSP+Op,OrangeSave21,OrangeSave22),MULx(CSP+Op,GreenSaveUp,p.NFI(4)))

	y,z,w = ASApr(CSP+Op,p.NFI(74),YellowGreenSideRight,p.NFI(4))
	OrangeDown = ADDx(CSP+Op,OrangeDown,w)
	OrangeSave24 := z
	OrangeSave25 := y
	YellowGreenSideLeft = SUBx(CSP+Op,ADDx(CSP+Op,OrangeSave23,OrangeSave24),MULx(CSP+Op,GreenSaveDown,p.NFI(2)))

	y,z,w = ASApr(CSP+Op,p.NFI(7),YellowGreenSideLeft,p.NFI(78))
	OrangeDown = ADDx(CSP+Op,OrangeDown,w)
	OrangeSave26 := y
	OrangeSave27 := z
	GreenBlueSideLeft = SUBx(CSP+Op,ADDx(CSP+Op,OrangeSave25,OrangeSave26),MULx(CSP+Op,BlueSaveUp,p.NFI(5)))

	y,z,w = ASApr(CSP+Op,p.NFI(85),GreenBlueSideLeft,p.NFI(3))
	OrangeDown = ADDx(CSP+Op,OrangeDown,w)
	OrangeSave28 := z
	OrangeSave29 := y
	GreenBlueSideRight = SUBx(CSP+Op,ADDx(CSP+Op,OrangeSave27,OrangeSave28),BlueSaveDown)

	y,z,w = ASApr(CSP+Op,p.NFI(7),GreenBlueSideRight,p.NFI(88))
	OrangeDown = ADDx(CSP+Op,OrangeDown,w)
	OrangeSave30 := y
	OrangeSave31 := z
	BlueIndigoSideLeft = SUBx(CSP+Op,ADDx(CSP+Op,OrangeSave29,OrangeSave30),MULx(CSP+Op,IndigoSaveUp,p.NFI(4)))

	y,z,w = ASApr(CSP+Op,p.NFI(95),BlueIndigoSideLeft,p.NFI(2))
	OrangeDown = ADDx(CSP+Op,OrangeDown,w)
	OrangeSave32 := z
	OrangeSave33 := y
	BlueIndigoSideRight = SUBx(CSP+Op,ADDx(CSP+Op,OrangeSave31,OrangeSave32),MULx(CSP+Op,IndigoSaveDown,p.NFI(2)))

	y,z,w = ASApr(CSP+Op,p.NFI(7),BlueIndigoSideRight,p.NFI(97))
	OrangeDown = ADDx(CSP+Op,OrangeDown,w)
	OrangeSave34 := y
	OrangeSave35 := z
	IndigoPurpleSideLeft = SUBx(CSP+Op,ADDx(CSP+Op,OrangeSave33,OrangeSave34),MULx(CSP+Op,PurpleSaveUp,p.NFI(3)))

	y,z,w = ASApr(CSP+Op,p.NFI(104),IndigoPurpleSideLeft,p.NFI(1))
	OrangeDown = ADDx(CSP+Op,OrangeDown,w)
	OrangeSave36 := z
	IndigoPurpleSideRight = SUBx(CSP+Op,ADDx(CSP+Op,OrangeSave35,OrangeSave36),MULx(CSP+Op,PurpleSaveDown,p.NFI(3)))

	y,z,w = ASApr(CSP+Op,p.NFI(7),IndigoPurpleSideRight,p.NFI(105))
	OrangeDown = ADDx(CSP+Op,OrangeDown,w)
	OD1 	:= MULx(CSP+Op,p.NFI(4),GreenUp)
	OD2 	:= MULx(CSP+Op,p.NFI(5),BlueUp)
	OD3 	:= MULx(CSP+Op,p.NFI(4),IndigoUp)
	OD4 	:= MULx(CSP+Op,p.NFI(3),PurpleUp)
	OD5 	:= MULx(CSP+Op,p.NFI(3),PurpleDown)
	OD6 	:= MULx(CSP+Op,p.NFI(2),IndigoDown)
	OD7 	:= MULx(CSP+Op,p.NFI(2),GreenDown)
	OD8 	:= MULx(CSP+Op,p.NFI(6),YellowDown)
	OrangeDown = SUMx(CSP+Op,OrangeDown,BlueDown,OD1,OD2,OD3,OD4,OD5,OD6,OD7,OD8)
	OrangeTotalDown := ADDx(CSP+Op,OrangeBaseDown,OrangeDown)

	ODH1 	:= SUBx(CSP+Op,MULx(CSP+Op,p.NFI(6),YellowDownHeight),OrangeDownHeight)

	ODH21 	:= MULx(CSP+Op,p.NFI(2),GreenDownHeight)
	ODH22 	:= MULx(CSP+Op,p.NFI(4),GreenUpHeight)
	ODH23	:= SUBx(CSP+Op,ODH21,ODH22)
	ODH2	:= ADDx(CSP+Op,ODH1,ODH23)

	ODH31 	:= MULx(CSP+Op,p.NFI(5),BlueUpHeight)
	ODH32 	:= ADDx(CSP+Op,ODH2,BlueDownHeight)
	ODH3	:= SUBx(CSP+Op,ODH32,ODH31)

	ODH41 	:= MULx(CSP+Op,p.NFI(2),IndigoDownHeight)
	ODH42 	:= MULx(CSP+Op,p.NFI(4),IndigoUpHeight)
	ODH43	:= SUBx(CSP+Op,ODH41,ODH42)
	ODH4	:= ADDx(CSP+Op,ODH3,ODH43)

	ODH51 	:= MULx(CSP+Op,p.NFI(3),PurpleDownHeight)
	ODH52 	:= MULx(CSP+Op,p.NFI(3),PurpleUpHeight)
	ODH53	:= SUBx(CSP+Op,ODH51,ODH52)
	ODH5	:= ADDx(CSP+Op,ODH4,ODH53)
	OrangeDownHeightPeakSeedC := TruncSeed(ODH5)
	fmt.Println("OrangeDownHeightPeakSeed:      ", OrangeDownHeightPeakSeedC)

	ODA11	:= MULx(CSP+Op,p.NFI(15),YellowDownHeight)
	ODA12	:= ADDx(CSP+Op,ODA11,OrangeDownHeight)
	ODA13	:= MULx(CSP+Op,Yellow,ODA12)
	ODA14	:= MULx(CSP+Op,p.NFI(6),YellowTotalDown)
	ODA15	:= ADDx(CSP+Op,ODA13,ODA14)
	ODA1	:= SUBx(CSP+Op,OrangeTotalDown,ODA15)

	ODA21	:= MULx(CSP+Op,p.NFI(2),ODH1)
	ODA22	:= MULx(CSP+Op,p.NFI(10),GreenUpHeight)
	ODA23 	:= SUMx(CSP+Op,GreenDownHeight,ODA21,ODA22)
	ODA24 	:= MULx(CSP+Op,Green,ODA23)
	ODA25	:= MULx(CSP+Op,p.NFI(2),GreenTotalDown)
	ODA26	:= MULx(CSP+Op,p.NFI(4),GreenTotalUp)
	ODA27 	:= SUMx(CSP+Op,ODA24,ODA25,ODA26)
	ODA2	:= SUBx(CSP+Op,ODA1,ODA27)

	ODA31	:= ADDx(CSP+Op,MULx(CSP+Op,p.NFI(15),BlueUpHeight),ODH2)
	ODA32	:= MULx(CSP+Op,Blue,ODA31)
	ODA33	:= MULx(CSP+Op,p.NFI(5),BlueTotalUp)
	ODA34	:= SUMx(CSP+Op,BlueTotalDown,ODA32,ODA33)
	ODA3	:= SUBx(CSP+Op,ODA2,ODA34)

	ODA41	:= MULx(CSP+Op,p.NFI(2),ODH3)
	ODA42	:= MULx(CSP+Op,p.NFI(10),IndigoUpHeight)
	ODA43	:= SUMx(CSP+Op,IndigoDownHeight,ODA41,ODA42)
	ODA44	:= MULx(CSP+Op,Indigo,ODA43)
	ODA45	:= MULx(CSP+Op,p.NFI(2),IndigoTotalDown)
	ODA46	:= MULx(CSP+Op,p.NFI(4),IndigoTotalUp)
	ODA47	:= SUMx(CSP+Op,ODA44,ODA45,ODA46)
	ODA4	:= SUBx(CSP+Op,ODA3,ODA47)

	ODA51	:= MULx(CSP+Op,p.NFI(3),ODH4)
	ODA52	:= MULx(CSP+Op,p.NFI(3),PurpleDownHeight)
	ODA53	:= MULx(CSP+Op,p.NFI(6),PurpleUpHeight)
	ODA54	:= SUMx(CSP+Op,ODA51,ODA52,ODA53)
	ODA55	:= MULx(CSP+Op,Purple,ODA54)
	ODA56	:= ADDx(CSP+Op,PurpleTotalUp,PurpleTotalDown)
	ODA57	:= MULx(CSP+Op,p.NFI(3),ODA56)
	ODA58	:= ADDx(CSP+Op,ODA55,ODA57)
	ODA5	:= SUBx(CSP+Op,ODA4,ODA58)
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
	RedOrangeSideLeft := SUBx(CSP+RUp,y,MULx(CSP+RUp,p.NFI(5),OrangeSaveUp))
	RedSave01 := z

	y,z,w = ASApr(CSP+RUp,p.NFI(49),RedOrangeSideLeft,p.NFI(6))
	RedUp = ADDx(CSP+RUp,RedUp,w)
	RedSave02 := z
	RedSave03 := y
	RedOrangeSideRight := SUBx(CSP+RUp,ADDx(CSP+RUp,RedSave01,RedSave02),OrangeSaveDown)

	y,z,w = ASApr(CSP+RUp,p.NFI(7),RedOrangeSideRight,p.NFI(55))
	RedUp = ADDx(CSP+RUp,RedUp,w)
	RedSave04 := y
	RedSave05 := z
	OrangeYellowSideLeft = SUBx(CSP+RUp,ADDx(CSP+RUp,RedSave03,RedSave04),MULx(CSP+RUp,YellowSaveUp,p.NFI(4)))

	y,z,w = ASApr(CSP+RUp,p.NFI(62),OrangeYellowSideLeft,p.NFI(5))
	RedUp = ADDx(CSP+RUp,RedUp,w)
	RedSave06 := z
	RedSave07 := y
	OrangeYellowSideRight = SUBx(CSP+RUp,ADDx(CSP+RUp,RedSave05,RedSave06),MULx(CSP+RUp,YellowSaveDown,p.NFI(2)))

	y,z,w = ASApr(CSP+RUp,p.NFI(7),OrangeYellowSideRight,p.NFI(67))
	RedUp = ADDx(CSP+RUp,RedUp,w)
	RedSave08 := y
	RedSave09 := z
	YellowGreenSideLeft = SUBx(CSP+RUp,ADDx(CSP+RUp,RedSave07,RedSave08),MULx(CSP+RUp,GreenSaveUp,p.NFI(4)))

	y,z,w = ASApr(CSP+RUp,p.NFI(74),YellowGreenSideLeft,p.NFI(4))
	RedUp = ADDx(CSP+RUp,RedUp,w)
	RedSave10 := z
	RedSave11 := y
	YellowGreenSideRight = SUBx(CSP+RUp,ADDx(CSP+RUp,RedSave09,RedSave10),MULx(CSP+RUp,GreenSaveDown,p.NFI(2)))

	y,z,w = ASApr(CSP+RUp,p.NFI(7),YellowGreenSideRight,p.NFI(78))
	RedUp = ADDx(CSP+RUp,RedUp,w)
	RedSave12 := y
	RedSave13 := z
	GreenBlueSideLeft = SUBx(CSP+RUp,ADDx(CSP+RUp,RedSave11,RedSave12),MULx(CSP+RUp,BlueSaveUp,p.NFI(6)))

	y,z,w = ASApr(CSP+RUp,p.NFI(85),GreenBlueSideLeft,p.NFI(3))
	RedUp = ADDx(CSP+RUp,RedUp,w)
	RedSave14 := z
	RedSave15 := y
	GreenBlueSideRight = ADDx(CSP+RUp,RedSave13,RedSave14)

	y,z,w = ASApr(CSP+RUp,p.NFI(7),GreenBlueSideRight,p.NFI(88))
	RedUp = ADDx(CSP+RUp,RedUp,w)
	RedSave16 := y
	RedSave17 := z
	BlueIndigoSideLeft = SUBx(CSP+RUp,ADDx(CSP+RUp,RedSave15,RedSave16),MULx(CSP+RUp,IndigoSaveUp,p.NFI(2)))

	y,z,w = ASApr(CSP+RUp,p.NFI(95),BlueIndigoSideLeft,p.NFI(2))
	RedUp = ADDx(CSP+RUp,RedUp,w)
	RedSave18 := z
	RedSave19 := y
	BlueIndigoSideRight = SUBx(CSP+RUp,ADDx(CSP+RUp,RedSave17,RedSave18),MULx(CSP+RUp,IndigoSaveDown,p.NFI(4)))

	y,z,w = ASApr(CSP+RUp,p.NFI(7),BlueIndigoSideRight,p.NFI(97))
	RedUp = ADDx(CSP+RUp,RedUp,w)
	RedSave20 := y
	RedSave21 := z
	IndigoPurpleSideLeft = SUBx(CSP+RUp,ADDx(CSP+RUp,RedSave19,RedSave20),MULx(CSP+RUp,PurpleSaveUp,p.NFI(2)))

	y,z,w = ASApr(CSP+RUp,p.NFI(104),IndigoPurpleSideLeft,p.NFI(1))
	RedUp = ADDx(CSP+RUp,RedUp,w)
	RedSave22 := z
	IndigoPurpleSideRight = SUBx(CSP+RUp,ADDx(CSP+RUp,RedSave21,RedSave22),MULx(CSP+RUp,PurpleSaveDown,p.NFI(4)))

	y,z,w = ASApr(CSP+RUp,p.NFI(7),IndigoPurpleSideRight,p.NFI(105))
	RedUp = ADDx(CSP+RUp,RedUp,w)
	RU1 	:= MULx(CSP+RUp,p.NFI(5),OrangeUp)
	RU2 	:= MULx(CSP+RUp,p.NFI(6),BlueUp)
	RU3		:= SUMx(CSP+RUp,YellowUp,GreenUp,PurpleDown,IndigoDown)
	RU4 	:= SUMx(CSP+RUp,IndigoUp,PurpleUp,GreenDown,YellowDown)
	RU5 	:= MULx(CSP+RUp,p.NFI(4),RU3)
	RU6 	:= MULx(CSP+RUp,p.NFI(2),RU4)
	RedUp	 = SUMx(CSP+RUp,RedUp,OrangeDown,RU1,RU2,RU5,RU6)
	RedTotalUp := ADDx(CSP+RUp,RedBaseUp,RedUp)

	RUH11 	:= MULx(CSP+RUp,p.NFI(5),OrangeUpHeight)
	RUH12	:= ADDx(CSP+RUp,RedUpHeight,OrangeDownHeight)
	RUH1	:= SUBx(CSP+RUp,RUH12,RUH11)

	RUH21 	:= MULx(CSP+RUp,p.NFI(4),YellowUpHeight)
	RUH22 	:= MULx(CSP+RUp,p.NFI(2),YellowDownHeight)
	RUH23	:= SUBx(CSP+RUp,RUH22,RUH21)
	RUH2	:= ADDx(CSP+RUp,RUH1,RUH23)

	RUH31 	:= MULx(CSP+RUp,p.NFI(4),GreenUpHeight)
	RUH32 	:= MULx(CSP+RUp,p.NFI(2),GreenDownHeight)
	RUH33	:= SUBx(CSP+RUp,RUH32,RUH31)
	RUH3	:= ADDx(CSP+RUp,RUH2,RUH33)

	RUH4	:= SUBx(CSP+RUp,MULx(CSP+RUp,p.NFI(6),BlueUpHeight),RUH3)

	RUH51 	:= MULx(CSP+RUp,p.NFI(4),IndigoDownHeight)
	RUH52 	:= MULx(CSP+RUp,p.NFI(2),IndigoUpHeight)
	RUH53	:= SUBx(CSP+RUp,RUH52,RUH51)
	RUH5	:= ADDx(CSP+RUp,RUH4,RUH53)

	RUH61 	:= MULx(CSP+RUp,p.NFI(4),PurpleDownHeight)
	RUH62 	:= MULx(CSP+RUp,p.NFI(2),PurpleUpHeight)
	RUH6 	:= DIFx(CSP+RUp,RUH61,RUH62,RUH5)
	RedUpHeightPeakSeedC := TruncSeed(RUH6)
	fmt.Println("RedUpHeightPeakSeed is:        ", RedUpHeightPeakSeedC)

	RUA11	:= MULx(CSP+RUp,p.NFI(15),OrangeUpHeight)
	RUA12 	:= ADDx(CSP+RUp,RedUpHeight,RUA11)
	RUA13	:= MULx(CSP+RUp,Orange,RUA12)
	RUA14	:= MULx(CSP+RUp,p.NFI(5),OrangeTotalUp)
	RUA15 	:= SUMx(CSP+RUp,OrangeTotalDown,RUA13,RUA14)
	RUA1	:= SUBx(CSP+RUp,RedTotalUp,RUA15)

	RUA21	:= MULx(CSP+RUp,p.NFI(2),RUH1)
	RUA22	:= MULx(CSP+RUp,p.NFI(10),YellowUpHeight)
	RUA23	:= SUMx(CSP+RUp,YellowDownHeight,RUA21,RUA22)
	RUA24	:= MULx(CSP+RUp,Yellow,RUA23)
	RUA25	:= MULx(CSP+RUp,p.NFI(2),YellowTotalDown)
	RUA26	:= MULx(CSP+RUp,p.NFI(4),YellowTotalUp)
	RUA27	:= SUMx(CSP+RUp,RUA24,RUA25,RUA26)
	RUA2	:= SUBx(CSP+RUp,RUA1,RUA27)

	RUA31	:= MULx(CSP+RUp,p.NFI(2),RUH2)
	RUA32	:= MULx(CSP+RUp,p.NFI(10),GreenUpHeight)
	RUA33	:= SUMx(CSP+RUp,GreenDownHeight,RUA31,RUA32)
	RUA34	:= MULx(CSP+RUp,Green,RUA33)
	RUA35	:= MULx(CSP+RUp,p.NFI(2),GreenTotalDown)
	RUA36	:= MULx(CSP+RUp,p.NFI(4),GreenTotalUp)
	RUA37	:= SUMx(CSP+RUp,RUA34,RUA35,RUA36)
	RUA3	:= SUBx(CSP+RUp,RUA2,RUA37)

	RUA41	:= ADDx(CSP+RUp,MULx(CSP+RUp,p.NFI(15),BlueUpHeight),RUH3)
	RUA42	:= MULx(CSP+RUp,Blue,RUA41)
	RUA43	:= MULx(CSP+RUp,p.NFI(6),BlueTotalUp)
	RUA44	:= ADDx(CSP+RUp,RUA42,RUA43)
	RUA4	:= SUBx(CSP+RUp,RUA3,RUA44)

	RUA51	:= MULx(CSP+RUp,p.NFI(2),RUH4)
	RUA52	:= MULx(CSP+RUp,p.NFI(10),IndigoDownHeight)
	RUA53 	:= SUMx(CSP+RUp,IndigoUpHeight,RUA51,RUA52)
	RUA54 	:= MULx(CSP+RUp,Indigo,RUA53)
	RUA55	:= MULx(CSP+RUp,p.NFI(2),IndigoTotalUp)
	RUA56	:= MULx(CSP+RUp,p.NFI(4),IndigoTotalDown)
	RUA57	:= SUMx(CSP+RUp,RUA54,RUA55,RUA56)
	RUA5	:= SUBx(CSP+RUp,RUA4,RUA57)

	RUA61	:= MULx(CSP+RUp,p.NFI(3),RUH5)
	RUA62	:= MULx(CSP+RUp,p.NFI(6),PurpleDownHeight)
	RUA63	:= MULx(CSP+RUp,p.NFI(3),PurpleUpHeight)
	RUA64 	:= SUMx(CSP+RUp,RUA61,RUA62,RUA63)
	RUA65 	:= MULx(CSP+RUp,Purple,RUA64)
	RUA66	:= MULx(CSP+RUp,p.NFI(2),PurpleTotalUp)
	RUA67	:= MULx(CSP+RUp,p.NFI(4),PurpleTotalDown)
	RUA68	:= SUMx(CSP+RUp,RUA65,RUA66,RUA67)
	RUA6	:= SUBx(CSP+RUp,RUA5,RUA68)
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
	RedDown = ADDx(CSP+RDp,RedDown,w)
	RedSave24 := z
	RedSave25 := y
	RedOrangeSideRight = SUBx(CSP+RDp,ADDx(CSP+RDp,RedSave23,RedSave24),MULx(CSP+RDp,OrangeSaveDown,p.NFI(6)))

	y,z,w = ASApr(CSP+RDp,p.NFI(7),RedOrangeSideRight,p.NFI(55))
	RedDown = ADDx(CSP+RDp,RedDown,w)
	RedSave26 := y
	RedSave27 := z
	OrangeYellowSideLeft = SUBx(CSP+RDp,ADDx(CSP+RDp,RedSave25,RedSave26),MULx(CSP+RDp,YellowSaveUp,p.NFI(5)))

	y,z,w = ASApr(CSP+RDp,p.NFI(62),OrangeYellowSideLeft,p.NFI(5))
	RedDown = ADDx(CSP+RDp,RedDown,w)
	RedSave28 := z
	RedSave29 := y
	OrangeYellowSideRight = SUBx(CSP+RDp,ADDx(CSP+RDp,RedSave27,RedSave28),YellowSaveDown)

	y,z,w = ASApr(CSP+RDp,p.NFI(7),OrangeYellowSideRight,p.NFI(67))
	RedDown = ADDx(CSP+RDp,RedDown,w)
	RedSave30 := y
	RedSave31 := z
	YellowGreenSideLeft = SUBx(CSP+RDp,ADDx(CSP+RDp,RedSave29,RedSave30),MULx(CSP+RDp,GreenSaveUp,p.NFI(4)))

	y,z,w = ASApr(CSP+RDp,p.NFI(74),YellowGreenSideLeft,p.NFI(4))
	RedDown = ADDx(CSP+RDp,RedDown,w)
	RedSave32 := z
	RedSave33 := y
	YellowGreenSideRight = SUBx(CSP+RDp,ADDx(CSP+RDp,RedSave31,RedSave32),MULx(CSP+RDp,GreenSaveDown,p.NFI(2)))

	y,z,w = ASApr(CSP+RDp,p.NFI(7),YellowGreenSideRight,p.NFI(78))
	RedDown = ADDx(CSP+RDp,RedDown,w)
	RedSave34 := y
	RedSave35 := z
	GreenBlueSideLeft = SUBx(CSP+RDp,ADDx(CSP+RDp,RedSave33,RedSave34),MULx(CSP+RDp,BlueSaveUp,p.NFI(3)))

	y,z,w = ASApr(CSP+RDp,p.NFI(85),GreenBlueSideLeft,p.NFI(3))
	RedDown = ADDx(CSP+RDp,RedDown,w)
	RedSave36 := z
	RedSave37 := y
	GreenBlueSideRight = SUBx(CSP+RDp,ADDx(CSP+RDp,RedSave35,RedSave36),MULx(CSP+RDp,BlueSaveDown,p.NFI(3)))

	y,z,w = ASApr(CSP+RDp,p.NFI(7),GreenBlueSideRight,p.NFI(88))
	RedDown = ADDx(CSP+RDp,RedDown,w)
	RedSave38 := y
	RedSave39 := z
	BlueIndigoSideLeft = SUBx(CSP+RDp,ADDx(CSP+RDp,RedSave37,RedSave38),MULx(CSP+RDp,IndigoSaveUp,p.NFI(6)))

	y,z,w = ASApr(CSP+RDp,p.NFI(95),BlueIndigoSideLeft,p.NFI(2))
	RedDown = ADDx(CSP+RDp,RedDown,w)
	RedSave40 := z
	RedSave41 := y
	BlueIndigoSideRight = ADDx(CSP+RDp,RedSave39,RedSave40)

	y,z,w = ASApr(CSP+RDp,p.NFI(7),BlueIndigoSideRight,p.NFI(97))
	RedDown = ADDx(CSP+RDp,RedDown,w)
	RedSave42 := y
	RedSave43 := z
	IndigoPurpleSideLeft = SUBx(CSP+RDp,ADDx(CSP+RDp,RedSave41,RedSave42),MULx(CSP+RDp,PurpleSaveUp,p.NFI(2)))

	y,z,w = ASApr(CSP+RDp,p.NFI(104),IndigoPurpleSideLeft,p.NFI(1))
	RedDown = ADDx(CSP+RDp,RedDown,w)
	RedSave44 := z
	IndigoPurpleSideRight = SUBx(CSP+RDp,ADDx(CSP+RDp,RedSave43,RedSave44),MULx(CSP+RDp,PurpleSaveDown,p.NFI(4)))

	y,z,w = ASApr(CSP+RDp,p.NFI(7),IndigoPurpleSideRight,p.NFI(105))
	RedDown = ADDx(CSP+RDp,RedDown,w)
	RD1		:= MULx(CSP+RDp,p.NFI(5),YellowUp)
	RD2		:= MULx(CSP+RDp,p.NFI(4),GreenUp)
	RD3		:= MULx(CSP+RDp,p.NFI(3),BlueUp)
	RD4		:= MULx(CSP+RDp,p.NFI(6),IndigoUp)
	RD5		:= MULx(CSP+RDp,p.NFI(2),PurpleUp)
	RD6		:= MULx(CSP+RDp,p.NFI(6),OrangeDown)
	RD7		:= MULx(CSP+RDp,p.NFI(2),GreenDown)
	RD8		:= MULx(CSP+RDp,p.NFI(3),BlueDown)
	RD9		:= MULx(CSP+RDp,p.NFI(4),PurpleDown)
	RedDown = SUMx(CSP+RDp,RedDown,YellowDown,RD1,RD2,RD3,RD4,RD5,RD6,RD7,RD8,RD9)
	RedTotalDown := ADDx(CSP+RDp,RedBaseDown,RedDown)

	RDH1 	:= SUBx(CSP+RDp,MULx(CSP+RDp,p.NFI(6),OrangeDownHeight),RedDownHeight)

	RDH21	:= MULx(CSP+RDp,p.NFI(5),YellowUpHeight)
	RDH22	:= SUBx(CSP+RDp,YellowDownHeight,RDH21)
	RDH2	:= ADDx(CSP+RDp,RDH1,RDH22)

	RDH31	:= MULx(CSP+RDp,p.NFI(4),GreenUpHeight)
	RDH32	:= MULx(CSP+RDp,p.NFI(2),GreenDownHeight)
	RDH33	:= SUBx(CSP+RDp,RDH32,RDH31)
	RDH3	:= ADDx(CSP+RDp,RDH2,RDH33)

	RDH41	:= MULx(CSP+RDp,p.NFI(3),BlueUpHeight)
	RDH42	:= MULx(CSP+RDp,p.NFI(3),BlueDownHeight)
	RDH43	:= SUBx(CSP+RDp,RDH42,RDH41)
	RDH4	:= ADDx(CSP+RDp,RDH3,RDH43)

	RDH5	:= SUBx(CSP+RDp,MULx(CSP+RDp,p.NFI(6),IndigoUpHeight),RDH4)

	RDH61	:= ADDx(CSP+RDp,MULx(CSP+RDp,p.NFI(2),PurpleUpHeight),RDH5)
	RDH6	:= SUBx(CSP+RDp,MULx(CSP+RDp,p.NFI(4),PurpleDownHeight),RDH61)
	RedDownHeightPeakSeedC := TruncSeed(RDH6)
	fmt.Println("RedDownHeightPeakSeed:         ", RedDownHeightPeakSeedC)

	RDA11	:= ADDx(CSP+RDp,MULx(CSP+RDp,p.NFI(15),OrangeDownHeight),RedDownHeight)
	RDA12	:= ADDx(CSP+RDp,MULx(CSP+RDp,p.NFI(6),OrangeTotalDown),MULx(CSP+RDp,Orange,RDA11))
	RDA1	:= SUBx(CSP+RDp,RedTotalDown,RDA12)

	RDA21	:= ADDx(CSP+RDp,MULx(CSP+RDp,p.NFI(15),YellowUpHeight),RDH1)
	RDA22	:= MULx(CSP+RDp,Yellow,RDA21)
	RDA23	:= MULx(CSP+RDp,p.NFI(5),YellowTotalUp)
	RDA24	:= SUMx(CSP+RDp,YellowTotalDown,RDA22,RDA23)
	RDA2	:= SUBx(CSP+RDp,RDA1,RDA24)

	RDA31	:= MULx(CSP+RDp,p.NFI(2),RDH2)
	RDA32	:= MULx(CSP+RDp,p.NFI(10),GreenUpHeight)
	RDA33	:= SUMx(CSP+RDp,GreenDownHeight,RDA31,RDA32)
	RDA34 	:= MULx(CSP+RDp,Green,RDA33)
	RDA35	:= MULx(CSP+RDp,p.NFI(2),GreenTotalDown)
	RDA36	:= MULx(CSP+RDp,p.NFI(4),GreenTotalUp)
	RDA37	:= SUMx(CSP+RDp,RDA34,RDA35,RDA36)
	RDA3	:= SUBx(CSP+RDp,RDA2,RDA37)

	RDA41	:= MULx(CSP+RDp,p.NFI(3),RDH3)
	RDA42	:= MULx(CSP+RDp,p.NFI(3),BlueDownHeight)
	RDA43	:= MULx(CSP+RDp,p.NFI(6),BlueUpHeight)
	RDA44	:= SUMx(CSP+RDp,RDA41,RDA42,RDA43)
	RDA45	:= MULx(CSP+RDp,Blue,RDA44)
	RDA46	:= ADDx(CSP+RDp,BlueTotalUp,BlueTotalDown)
	RDA47	:= MULx(CSP+RDp,p.NFI(3),RDA46)
	RDA48	:= ADDx(CSP+RDp,RDA45,RDA47)
	RDA4	:= SUBx(CSP+RDp,RDA3,RDA48)

	RDA51 	:= ADDx(CSP+RDp,MULx(CSP+RDp,p.NFI(15),IndigoUpHeight),RDH4)
	RDA52	:= MULx(CSP+RDp,Indigo,RDA51)
	RDA53	:= MULx(CSP+RDp,p.NFI(6),IndigoTotalUp)
	RDA54	:= ADDx(CSP+RDp,RDA52,RDA53)
	RDA5	:= SUBx(CSP+RDp,RDA4,RDA54)

	RDA61	:= MULx(CSP+RDp,p.NFI(3),RDH5)
	RDA62	:= MULx(CSP+RDp,p.NFI(6),PurpleDownHeight)
	RDA63	:= MULx(CSP+RDp,p.NFI(3),PurpleUpHeight)
	RDA64	:= SUMx(CSP+RDp,RDA61,RDA62,RDA63)
	RDA65	:= MULx(CSP+RDp,Purple,RDA64)
	RDA66	:= MULx(CSP+RDp,p.NFI(4),PurpleTotalDown)
	RDA67	:= MULx(CSP+RDp,p.NFI(2),PurpleTotalUp)
	RDA68	:= SUMx(CSP+RDp,RDA65,RDA66,RDA67)
	RDA6	:= SUBx(CSP+RDp,RDA5,RDA68)
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
	WhiteRedSideLeft := SUBx(CSP+Whp,y,RedSaveUp)

	y,z,w = ASApr(CSP+Whp,p.NFI(35),WhiteRedSideLeft,p.NFI(7))
	TotalArea = ADDx(CSP+Whp,TotalArea,w)
	WhiteSave02 := z
	WhiteSave03	:= y
	WhiteRedSideRight := SUBx(CSP+Whp,ADDx(CSP+Whp,WhiteSave01,WhiteSave02),MULx(CSP+Whp,RedSaveDown,p.NFI(5)))

	y,z,w = ASApr(CSP+Whp,p.NFI(7),WhiteRedSideRight,p.NFI(42))
	TotalArea = ADDx(CSP+Whp,TotalArea,w)
	WhiteSave04 := y
	WhiteSave05	:= z
	RedOrangeSideLeft = SUBx(CSP+Whp,ADDx(CSP+Whp,WhiteSave03,WhiteSave04),MULx(CSP+Whp,OrangeSaveUp,p.NFI(2)))

	y,z,w = ASApr(CSP+Whp,p.NFI(49),RedOrangeSideLeft,p.NFI(6))
	TotalArea = ADDx(CSP+Whp,TotalArea,w)
	WhiteSave06 := z
	WhiteSave07	:= y
	RedOrangeSideRight = SUBx(CSP+Whp,ADDx(CSP+Whp,WhiteSave05,WhiteSave06),MULx(CSP+Whp,OrangeSaveDown,p.NFI(4)))

	y,z,w = ASApr(CSP+Whp,p.NFI(7),RedOrangeSideRight,p.NFI(55))
	TotalArea = ADDx(CSP+Whp,TotalArea,w)
	WhiteSave08 := y
	WhiteSave09	:= z
	OrangeYellowSideLeft = SUBx(CSP+Whp,ADDx(CSP+Whp,WhiteSave07,WhiteSave08),MULx(CSP+Whp,YellowSaveUp,p.NFI(3)))

	y,z,w = ASApr(CSP+Whp,p.NFI(62),OrangeYellowSideLeft,p.NFI(5))
	TotalArea = ADDx(CSP+Whp,TotalArea,w)
	WhiteSave10 := z
	WhiteSave11	:= y
	OrangeYellowSideRight = SUBx(CSP+Whp,ADDx(CSP+Whp,WhiteSave09,WhiteSave10),MULx(CSP+Whp,YellowSaveDown,p.NFI(3)))

	y,z,w = ASApr(CSP+Whp,p.NFI(7),OrangeYellowSideRight,p.NFI(67))
	TotalArea = ADDx(CSP+Whp,TotalArea,w)
	WhiteSave12 := y
	WhiteSave13	:= z
	YellowGreenSideLeft = SUBx(CSP+Whp,ADDx(CSP+Whp,WhiteSave11,WhiteSave12),GreenSaveUp)

	y,z,w = ASApr(CSP+Whp,p.NFI(74),YellowGreenSideLeft,p.NFI(4))
	TotalArea = ADDx(CSP+Whp,TotalArea,w)
	WhiteSave14 := z
	WhiteSave15	:= y
	YellowGreenSideRight = SUBx(CSP+Whp,ADDx(CSP+Whp,WhiteSave13,WhiteSave14),MULx(CSP+Whp,GreenSaveDown,p.NFI(5)))

	y,z,w = ASApr(CSP+Whp,p.NFI(7),YellowGreenSideRight,p.NFI(78))
	TotalArea = ADDx(CSP+Whp,TotalArea,w)
	WhiteSave16 := y
	WhiteSave17	:= z
	GreenBlueSideLeft = SUBx(CSP+Whp,ADDx(CSP+Whp,WhiteSave15,WhiteSave16),MULx(CSP+Whp,BlueSaveUp,p.NFI(2)))

	y,z,w = ASApr(CSP+Whp,p.NFI(85),GreenBlueSideLeft,p.NFI(3))
	TotalArea = ADDx(CSP+Whp,TotalArea,w)
	WhiteSave18 := z
	WhiteSave19	:= y
	GreenBlueSideRight = SUBx(CSP+Whp,ADDx(CSP+Whp,WhiteSave17,WhiteSave18),MULx(CSP+Whp,BlueSaveDown,p.NFI(4)))

	y,z,w = ASApr(CSP+Whp,p.NFI(7),GreenBlueSideRight,p.NFI(88))
	TotalArea = ADDx(CSP+Whp,TotalArea,w)
	WhiteSave20 := y
	WhiteSave21	:= z
	BlueIndigoSideLeft = SUBx(CSP+Whp,ADDx(CSP+Whp,WhiteSave19,WhiteSave20),IndigoSaveUp)

	y,z,w = ASApr(CSP+Whp,p.NFI(95),BlueIndigoSideLeft,p.NFI(2))
	TotalArea = ADDx(CSP+Whp,TotalArea,w)
	WhiteSave22 := z
	WhiteSave23	:= y
	BlueIndigoSideRight = SUBx(CSP+Whp,ADDx(CSP+Whp,WhiteSave21,WhiteSave22),MULx(CSP+Whp,IndigoSaveDown,p.NFI(5)))

	y,z,w = ASApr(CSP+Whp,p.NFI(7),BlueIndigoSideRight,p.NFI(97))
	TotalArea = ADDx(CSP+Whp,TotalArea,w)
	WhiteSave24 := y
	WhiteSave25	:= z
	IndigoPurpleSideLeft = SUBx(CSP+Whp,ADDx(CSP+Whp,WhiteSave23,WhiteSave24),MULx(CSP+Whp,PurpleSaveUp,p.NFI(3)))

	y,z,w = ASApr(CSP+Whp,p.NFI(104),IndigoPurpleSideLeft,p.NFI(1))
	TotalArea = ADDx(CSP+Whp,TotalArea,w)
	WhiteSave26 := z
	IndigoPurpleSideRight = SUBx(CSP+Whp,ADDx(CSP+Whp,WhiteSave25,WhiteSave26),MULx(CSP+Whp,PurpleSaveDown,p.NFI(3)))

	y,z,w = ASApr(CSP+Whp,p.NFI(7),IndigoPurpleSideRight,p.NFI(105))
	TotalArea = ADDx(CSP+Whp,TotalArea,w)
	TA1		:= MULx(CSP+Whp,p.NFI(2),OrangeUp)
	TA2		:= MULx(CSP+Whp,p.NFI(2),BlueUp)
	TA3		:= SUMx(CSP+Whp,YellowUp,PurpleUp,PurpleDown,YellowDown)
	TA4		:= MULx(CSP+Whp,p.NFI(3),TA3)
	TA5		:= MULx(CSP+Whp,p.NFI(4),BlueDown)
	TA6		:= MULx(CSP+Whp,p.NFI(4),OrangeDown)
	TA7		:= SUMx(CSP+Whp,IndigoDown,GreenDown,RedDown)
	TA8 	:= MULx(CSP+Whp,p.NFI(5),TA7)
	TotalArea = SUMx(CSP+Whp,TotalArea,RedUp,GreenUp,IndigoUp,TA1,TA2,TA4,TA5,TA6,TA8)

	WH1 	:= SUBx(CSP+Whp,MULx(CSP+Whp,p.NFI(5),RedDownHeight),RedUpHeight)

	WH21	:= SUBx(CSP+Whp,MULx(CSP+Whp,p.NFI(4),OrangeDownHeight),MULx(CSP+Whp,p.NFI(2),OrangeUpHeight))
	WH2		:= ADDx(CSP+Whp,WH1,WH21)

	WH31	:= MULx(CSP+Whp,p.NFI(3),YellowUpHeight)
	WH32	:= MULx(CSP+Whp,p.NFI(3),YellowDownHeight)
	WH3		:= DIFx(CSP+Whp,WH31,WH32,WH2)

	WH41	:= SUBx(CSP+Whp,GreenUpHeight,MULx(CSP+Whp,p.NFI(5),GreenDownHeight))
	WH4		:= ADDx(CSP+Whp,WH3,WH41)

	WH51	:= MULx(CSP+Whp,p.NFI(2),BlueUpHeight)
	WH52	:= MULx(CSP+Whp,p.NFI(4),BlueDownHeight)
	WH53	:= SUBx(CSP+Whp,WH51,WH52)
	WH5		:= ADDx(CSP+Whp,WH4,WH53)

	WH61	:= MULx(CSP+Whp,p.NFI(5),IndigoDownHeight)
	WH6		:= DIFx(CSP+Whp,WH61,IndigoUpHeight,WH5)

	WH71	:= MULx(CSP+Whp,p.NFI(3),PurpleUpHeight)
	WH72	:= MULx(CSP+Whp,p.NFI(3),PurpleDownHeight)
	WH7		:= DIFx(CSP+Whp,WH71,WH72,WH6)
	WhiteHeightPeakSeedC := TruncSeed(WH7)
	fmt.Println("WhiteHeightPeakSeed is:        ", WhiteHeightPeakSeedC)

	WA11	:= ADDx(CSP+Whp,MULx(CSP+Whp,p.NFI(10),RedDownHeight),RedUpHeight)
	WA12	:= MULx(CSP+Whp,Red,WA11)
	WA13	:= MULx(CSP+Whp,p.NFI(5),RedTotalDown)
	WA14	:= SUMx(CSP+Whp,RedTotalUp,WA13,WA12)
	WA1		:= SUBx(CSP+Whp,TotalArea,WA14)

	WA21	:= MULx(CSP+Whp,p.NFI(4),WH1)
	WA22	:= MULx(CSP+Whp,p.NFI(6),OrangeDownHeight)
	WA23	:= MULx(CSP+Whp,p.NFI(3),OrangeUpHeight)
	WA24	:= SUMx(CSP+Whp,WA21,WA22,WA23)
	WA25	:= MULx(CSP+Whp,Orange,WA24)
	WA26	:= MULx(CSP+Whp,p.NFI(4),OrangeTotalDown)
	WA27	:= MULx(CSP+Whp,p.NFI(2),OrangeTotalUp)
	WA28	:= SUMx(CSP+Whp,WA25,WA26,WA27)
	WA2		:= SUBx(CSP+Whp,WA1,WA28)

	WA31	:= MULx(CSP+Whp,p.NFI(4),WH2)
	WA32	:= MULx(CSP+Whp,p.NFI(6),YellowDownHeight)
	WA33	:= MULx(CSP+Whp,p.NFI(3),YellowUpHeight)
	WA34	:= SUMx(CSP+Whp,WA31,WA32,WA33)
	WA35	:= MULx(CSP+Whp,Yellow,WA34)
	WA36	:= MULx(CSP+Whp,p.NFI(3),YellowTotalDown)
	WA37	:= MULx(CSP+Whp,p.NFI(3),YellowTotalUp)
	WA38	:= SUMx(CSP+Whp,WA35,WA36,WA37)
	WA3		:= SUBx(CSP+Whp,WA2,WA38)

	WA41	:= ADDx(CSP+Whp,MULx(CSP+Whp,p.NFI(15),GreenDownHeight),WH3)
	WA42	:= MULx(CSP+Whp,Green,WA41)
	WA43	:= MULx(CSP+Whp,p.NFI(5),GreenTotalDown)
	WA44	:= SUMx(CSP+Whp,GreenTotalUp,WA43,WA42)
	WA4		:= SUBx(CSP+Whp,WA3,WA44)

	WA51	:= MULx(CSP+Whp,p.NFI(2),WH4)
	WA52	:= MULx(CSP+Whp,p.NFI(10),BlueDownHeight)
	WA53	:= SUMx(CSP+Whp,WA51,WA52,BlueUpHeight)
	WA54	:= MULx(CSP+Whp,Blue,WA53)
	WA55	:= MULx(CSP+Whp,p.NFI(4),BlueTotalDown)
	WA56	:= MULx(CSP+Whp,p.NFI(2),BlueTotalUp)
	WA57	:= SUMx(CSP+Whp,WA54,WA55,WA56)
	WA5		:= SUBx(CSP+Whp,WA4,WA57)

	WA61	:= MULx(CSP+Whp,p.NFI(2),WH5)
	WA62	:= MULx(CSP+Whp,p.NFI(10),IndigoDownHeight)
	WA63	:= SUMx(CSP+Whp,WA61,WA62,IndigoUpHeight)
	WA64	:= MULx(CSP+Whp,Indigo,WA63)
	WA65	:= MULx(CSP+Whp,p.NFI(5),IndigoTotalDown)
	WA66	:= SUMx(CSP+Whp,WA64,WA65,IndigoTotalUp)
	WA6		:= SUBx(CSP+Whp,WA5,WA66)

	WA71	:= MULx(CSP+Whp,p.NFI(4),WH6)
	WA72	:= MULx(CSP+Whp,p.NFI(6),PurpleDownHeight)
	WA73	:= MULx(CSP+Whp,p.NFI(3),PurpleUpHeight)
	WA74	:= SUMx(CSP+Whp,WA71,WA72,WA73)
	WA75	:= MULx(CSP+Whp,Purple,WA74)
	WA76	:= MULx(CSP+Whp,p.NFI(3),PurpleTotalDown)
	WA77	:= MULx(CSP+Whp,p.NFI(3),PurpleTotalUp)
	WA78	:= SUMx(CSP+Whp,WA75,WA76,WA77)
	WA7		:= SUBx(CSP+Whp,WA6,WA78)
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
	CPAreaRatio 	:= DIVx(CSP,CamelReward,TotalArea)
	CPAreaRatioSeedC := TruncSeed(CPAreaRatio)

	fmt.Println("Primary Decimal Seeds have been computed")
	fmt.Println("Seed1st, ", BaseAreaSeedC)
	fmt.Println("Seed2nd, ", BaseHeightSeedC)
	fmt.Println("Seed3nd, ", CPAreaRatioSeedC)
	fmt.Println("========================================================")

	elapsed := time.Since(start)
	fmt.Println("Computing took", elapsed)
}