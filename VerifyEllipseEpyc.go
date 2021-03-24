package main

import (
    el "Cryptoplasm-Core/Cryptoplasm_Elliptic"
)

func main () {
    var (
	PrimeNumber el.PrimePowerTwo
	VerifyLocation = ".ElipseComputing/"
	ModeU = "Unity"
	ModeT = "Trinity"
	Mode4 = "Easy4"
	Mode8 = "Easy8"
	Mode16 = "Easy16"
    )

    PrimeNumber.Power = 1029
    PrimeNumber.RestString = "639"
    PrimeNumber.Sign = true

    //Verify Ryzen Ellipse
    _,_ = el.VerifyEllipseResultsEpyc(VerifyLocation,PrimeNumber,32,ModeU)
    _,_ = el.VerifyEllipseResultsEpyc(VerifyLocation,PrimeNumber,32,ModeT)
    _,_ = el.VerifyEllipseResultsEpyc(VerifyLocation,PrimeNumber,32,Mode4)
    _,_ = el.VerifyEllipseResultsEpyc(VerifyLocation,PrimeNumber,32,Mode8)
    _,_ = el.VerifyEllipseResultsEpyc(VerifyLocation,PrimeNumber,32,Mode16)

}