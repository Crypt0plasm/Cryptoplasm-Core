package main

import (
    el "Cryptoplasm-Core/Cryptoplasm_Elliptic"
)

func main () {
    var (
	PrimeNumber el.PrimePowerTwo
	VerifyLocation = ".EllipseComputing/"
	ModeU = "Unity"
	ModeT = "Trinity"
	Mode4 = "Easy4"
	Mode8 = "Easy8"
	Mode16 = "Easy16"
    )

    PrimeNumber.Power = 6405
    PrimeNumber.RestString = "95"
    PrimeNumber.Sign = true

    //Verify Ryzen Ellipse
    _,_ = el.VerifyEllipseResultsRyzen(VerifyLocation,PrimeNumber,1,ModeU)
    _,_ = el.VerifyEllipseResultsRyzen(VerifyLocation,PrimeNumber,1,ModeT)
    _,_ = el.VerifyEllipseResultsRyzen(VerifyLocation,PrimeNumber,1,Mode4)
    _,_ = el.VerifyEllipseResultsRyzen(VerifyLocation,PrimeNumber,1,Mode8)
    _,_ = el.VerifyEllipseResultsRyzen(VerifyLocation,PrimeNumber,1,Mode16)

}