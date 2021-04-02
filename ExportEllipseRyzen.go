package main

import el "Cryptoplasm-Core/Cryptoplasm_Elliptic"

func main () {
    var (
	PrimeNumber el.PrimePowerTwo
    )

    PrimeNumber.Power = 1605
    PrimeNumber.RestString = "2315"
    PrimeNumber.Sign = true

    //Export for Ryzen (16 Cores, Negative Coefficients)
    el.ExportEllipse(PrimeNumber,0,12000,16,false)
}