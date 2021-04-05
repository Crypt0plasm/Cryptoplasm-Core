package Cryptoplasm_Elliptic

import (
    "fmt"
    "log"
    "math/big"
    "os"
    "strconv"
)

func MakeEllipticExportString (PrimeNumber PrimePowerTwo) string {
    var (
	S1 = "2^"
	S2 string
    )

    if PrimeNumber.Sign == true {
	S2 = "+"
    } else {
	S2 = "-"
    }

    PowerString := strconv.FormatInt(int64(PrimeNumber.Power),10)

    ExportString := S1 + PowerString + S2 + PrimeNumber.RestString
    return ExportString
}

func MakeEllipticExportStringPos (String string) string {
    Result := String + ".pos_"
    return Result
}

func MakeEllipticExportStringNeg (String string) string {
    Result := String + ".neg_"
    return Result
}

func ExportEllipse(Prime PrimePowerTwo, StartPoint int64, IntervalSize, Cores uint64, Direction bool) {
    //IntervalSize must be a multiple of Instances
    var (
	SecondRoot string
	Suffix     string
    )
    FirstRoot := MakeEllipticExportString(Prime)

    if Direction == true {
	SecondRoot = MakeEllipticExportStringPos(FirstRoot)
    } else {
	SecondRoot = MakeEllipticExportStringNeg(FirstRoot)
    }

    //Define the Prime Number
    PrimeNumber := MakePrime(Prime)
    //fmt.Println("PrimeNumber is",PrimeNumber)
    //fmt.Println("SecondRoot izz",SecondRoot)

    //Define the Number of instances created
    Instances := int64(Cores)	//+ int64(Cores) / 2

    if Direction == true {
	for k:=int64(1); k<=Instances; k++ {
	    if k < 10 {
		Suffix = "0" + strconv.FormatInt(k,10)
	    } else {
		Suffix = strconv.FormatInt(k,10)
	    }
	    ExportFileName := SecondRoot + Suffix + ".txt"
	    fmt.Println("===== k =",k,"=======")
	    for j:=int64(0); j*Instances<int64(IntervalSize); j++ {
		i := StartPoint + Instances * j + k
		fmt.Println("i iz", i)
		ComputeEllipseExport(ExportFileName,i,&PrimeNumber)
	    }
	}
    } else {
	for k:=int64(1); k<=Instances; k++ {
	    if k < 10 {
		Suffix = "0" + strconv.FormatInt(k,10)
	    } else {
		Suffix = strconv.FormatInt(k,10)
	    }
	    ExportFileName := SecondRoot + Suffix + ".txt"
	    fmt.Println("===== k =",k,"=======")
	    for j:=int64(0); j*Instances<int64(IntervalSize); j++ {
		i := StartPoint - Instances * j - k
		fmt.Println("i iz", i)
		ComputeEllipseExport(ExportFileName,i,&PrimeNumber)
	    }
	}
    }

}

func ComputeEllipseExport (Name string, DCoeff int64, Prime *big.Int) {
    var (
	D = new(big.Int)
	E = big.NewInt(1)
    )
    D.SetInt64(DCoeff)

    OutputFile, err := os.OpenFile(Name,os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
	log.Fatal(err)
    }
    defer OutputFile.Close()

    MB,MA := TEC2Montgomery(Prime,E,D)
    W1,W2 := Montgomery2ShortWeierstrass(Prime,MB,MA)
    Dis,_ := DIShortWeierstrass(Prime,W1,W2)

    Cmp := Zero.Cmp(Dis)
    if Cmp != 0 {
	Ds := D.Text(10)
	W1s:= W1.Text(10)
	W2s:= W2.Text(10)
	OutputString := Ds + ";" + W1s + ";" + W2s + ";"
	//fmt.Println("IZ:",OutputString)
	if _, err := OutputFile.WriteString(OutputString+"\n"); err != nil {
	    log.Println(err)
	}
    }
}