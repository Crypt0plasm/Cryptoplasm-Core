package main

import el "Cryptoplasm-Core/Cryptoplasm_Elliptic"

func main () {
    var (
	PrimeNumber el.PrimePowerTwo
    )

    PrimeNumber.Power = 1029
    PrimeNumber.RestString = "639"
    PrimeNumber.Sign = true

    //Export for Epyc (32 Cores, Positive Coefficients)
    el.ExportEllipse(PrimeNumber,0,1200000,16,true)
}