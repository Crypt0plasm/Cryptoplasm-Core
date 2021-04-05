package main

import (
    ca "Cryptoplasm-Core/Cryptoplasm_Cardinality"
    "fmt"
    "log"
    "os"
    "strconv"
)

func main () {
    var (
        DPMatrix []ca.DivisionPolynom
        DP ca.DivisionPolynom
    )

    DPExportName := "DivisionPolynomials.txt"

    OutputFile, err := os.OpenFile(DPExportName,os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        log.Fatal(err)
    }
    defer OutputFile.Close()



    for i:=1; i<2504; i++ {
        if i == 1 {
            DP = ca.DivisionPolynom{NonY: ca.POne, Y: ca.NoY}
            fmt.Println("   DP", i, "is:", ca.PrintDPolynom(DP))
            DPMatrix = ca.AppendDPolynom(DPMatrix, DP)

            //Writing Externally
            StringToWrite := "DP_" + strconv.Itoa(i) + "(X,Y) = " + ca.PrintDPolynom(DP)
            if _, errr := OutputFile.WriteString(StringToWrite+"\n"); errr != nil {
                log.Println(errr)
            }
        } else if i == 2 {
            DP = ca.DPTwo
            fmt.Println("   DP", i, "is:", ca.PrintDPolynom(DP))
            DPMatrix = ca.AppendDPolynom(DPMatrix, DP)

            //Writing Externally
            StringToWrite := "DP_" + strconv.Itoa(i) + "(X,Y) = " + ca.PrintDPolynom(DP)
            if _, errr := OutputFile.WriteString(StringToWrite+"\n"); errr != nil {
                log.Println(errr)
            }
        }else if i == 3 {
            DP = ca.DPThree
            fmt.Println("   DP", i, "is:", ca.PrintDPolynom(DP))
            DPMatrix = ca.AppendDPolynom(DPMatrix, DP)

            //Writing Externally
            StringToWrite := "DP_" + strconv.Itoa(i) + "(X,Y) = " + ca.PrintDPolynom(DP)
            if _, errr := OutputFile.WriteString(StringToWrite+"\n"); errr != nil {
                log.Println(errr)
            }
        }else if i == 4 {
            DP = ca.DP4
            fmt.Println("   DP", i, "is:", ca.PrintDPolynom(DP))
            DPMatrix = ca.AppendDPolynom(DPMatrix, DP)

            //Writing Externally
            StringToWrite := "DP_" + strconv.Itoa(i) + "(X,Y) = " + ca.PrintDPolynom(DP)
            if _, errr := OutputFile.WriteString(StringToWrite+"\n"); errr != nil {
                log.Println(errr)
            }
        } else if i%2 != 0 {
            n := (i - 1) / 2
            fmt.Println("START================DP",i,"n:=",n,"================")
            T1 := DPMatrix[n+2-1]
            T2 := ca.DPolynomCube(DPMatrix[n-1])
            T3 := ca.DPolynomCube(DPMatrix[n])
            T4 := DPMatrix[n-2]
            T12 := ca.DPolynomMul(T1,T2)
            T34 := ca.DPolynomMul(T3,T4)
            T12r := ca.ReduceDivisionPolynom(T12)
            T34r := ca.ReduceDivisionPolynom(T34)
            T := ca.PolynomSub(T12r,T34r)
            DP = ca.DivisionPolynom{NonY: T,Y: ca.NoY}

            //fmt.Println("T1 is DP",n+2,"^1:",ca.PrintDPolynom(T1))
            //fmt.Println("T2 is DP",n,"^3:",ca.PrintDPolynom(T2))
            //fmt.Println("T3 is DP",n+1,"^3:",ca.PrintDPolynom(T3))
            //fmt.Println("T4 is DP",n-1,"^1:",ca.PrintDPolynom(T4))
            //fmt.Println("T12 is",ca.PrintDPolynom(T12))
            //fmt.Println("T34 is",ca.PrintDPolynom(T34))
            //fmt.Println("   DP", i, "is:", ca.PrintDPolynom(DP))
            DPMatrix = ca.AppendDPolynom(DPMatrix, DP)

            //Writing Externally
            StringToWrite := "DP_" + strconv.Itoa(i) + "(X,Y) = " + ca.PrintDPolynom(DP)
            if _, errr := OutputFile.WriteString(StringToWrite+"\n"); errr != nil {
                log.Println(errr)
            }
            fmt.Println("END==================DP",i,"n:=",n,"================")
        } else {
            n := i / 2
            fmt.Println("START================DP",i,"n:=",n,"================")
            T0 := DPMatrix[n-1]
            T1 := DPMatrix[n+1]
            T2 := ca.DPolynomSqr(DPMatrix[n-2])
            T3 := DPMatrix[n-3]
            T4 := ca.DPolynomSqr(DPMatrix[n])
            T5 := ca.OneUnder2YDP
            T12 := ca.DPolynomMul(T1,T2)
            T34 := ca.DPolynomMul(T3,T4)
            //Computing T12-T34
            T12mT34 := ca.DivisionPolynomSub(T12,T34)
            //Multiplying with T0
            T0mT12mT34 := ca.DPolynomMul(T0,T12mT34)
            //Computing Final DP
            DP = ca.DPolynomMul(T0mT12mT34,T5)

            //fmt.Println("T0 is DP",n,"^1:",ca.PrintDPolynom(T0))
            //fmt.Println("T1 is DP",n+2,"^1:",ca.PrintDPolynom(T1))
            //fmt.Println("T2 is DP",n-1,"^2:",ca.PrintDPolynom(T2))
            //fmt.Println("T3 is DP",n-2,"^1:",ca.PrintDPolynom(T3))
            //fmt.Println("T4 is DP",n+1,"^3:",ca.PrintDPolynom(T4))
            //fmt.Println("T5 is",ca.PrintDPolynom(T5))
            //fmt.Println("T12 is",ca.PrintDPolynom(T12))
            //fmt.Println("T34 is",ca.PrintDPolynom(T34))
            //fmt.Println("T12mT34 is",ca.PrintDPolynom(T12mT34))
            //fmt.Println("T0mT12mT34 is",ca.PrintDPolynom(T0mT12mT34))
            //fmt.Println("   DP", i, "is:", ca.PrintDPolynom(DP))
            DPMatrix = ca.AppendDPolynom(DPMatrix, DP)

            //Writing Externally
            StringToWrite := "DP_" + strconv.Itoa(i) + "(X,Y) = " + ca.PrintDPolynom(DP)
            if _, errr := OutputFile.WriteString(StringToWrite+"\n"); errr != nil {
                log.Println(errr)
            }
            fmt.Println("END==================DP",i,"n:=",n,"================")
        }
    }
}