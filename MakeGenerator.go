package main

import (
    el "Cryptoplasm-Core/Cryptoplasm_Elliptic"
    "fmt"
)
func main () {
    Curve := el.TEC_S1600_Pr1605p2315_m26()
    MadeGenerator := Curve.MakeGenMin()
    fmt.Println("Made Generator of Order Q is",MadeGenerator)
}