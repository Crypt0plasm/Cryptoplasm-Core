package main

import el "Cryptoplasm-Core/Cryptoplasm_Elliptic"

func main () {
    var (
	PrimeNumber el.PrimePowerTwo
    )

    PrimeNumber.Power = 2250
    PrimeNumber.RestString = "6727"
    PrimeNumber.Sign = true

    //Export for Epyc (32 Cores, Positive Coefficients)
    el.ExportEllipse(PrimeNumber,0,1200000,32,true)
}