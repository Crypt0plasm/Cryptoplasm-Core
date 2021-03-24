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
    _,_ = el.VerifyEllipseResultsRyzen(VerifyLocation,PrimeNumber,16,ModeU)
    _,_ = el.VerifyEllipseResultsRyzen(VerifyLocation,PrimeNumber,16,ModeT)
    _,_ = el.VerifyEllipseResultsRyzen(VerifyLocation,PrimeNumber,16,Mode4)
    _,_ = el.VerifyEllipseResultsRyzen(VerifyLocation,PrimeNumber,16,Mode8)
    _,_ = el.VerifyEllipseResultsRyzen(VerifyLocation,PrimeNumber,16,Mode16)

}