package main

import el "Cryptoplasm-Core/Cryptoplasm_Elliptic"

func main () {
    var (
	PrimeNumber el.PrimePowerTwo
    )

    PrimeNumber.Power = 6405
    PrimeNumber.RestString = "95"
    PrimeNumber.Sign = true

    //Export for Epyc (32 Cores, Positive Coefficients)
    el.ExportEllipse(PrimeNumber,0,12000,1,true)
}