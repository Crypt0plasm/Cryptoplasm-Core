package main

import (
    el "Cryptoplasm-Core/Cryptoplasm_Elliptic"
    "fmt"
)
func main () {
    Curve := el.TEC_S1024_Pr1029p639_m729()
    MadeGenerator := Curve.MakeGenMin()
    fmt.Println("Made Generator of Order Q is",MadeGenerator)
}