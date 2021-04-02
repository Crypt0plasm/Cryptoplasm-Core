package main

import el "Cryptoplasm-Core/Cryptoplasm_Elliptic"

func main () {
    var (
	PrimeNumber el.PrimePowerTwo
    )

    PrimeNumber.Power = 1605
    PrimeNumber.RestString = "2315"
    PrimeNumber.Sign = true

    //Export for Epyc (32 Cores, Positive Coefficients)
    el.ExportEllipse(PrimeNumber,0,120000,12,true)
}