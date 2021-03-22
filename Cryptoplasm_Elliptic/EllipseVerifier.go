package Cryptoplasm_Elliptic

import (
    "bufio"
    "fmt"
    "log"
    "math/big"
    "os"
    "strconv"
    "strings"
)

func SafeScalarComputer(Prime, Trace, Cofactor *big.Int) (uint64, string) {
    var (
	Q = new(big.Int)
	Qs string
	Ss string
	Remainder = new(big.Int)
	X uint64
    )
    CofactorBase2 := Cofactor.Text(2)
    CofactorBase2trimmed := TrimFirstRune(CofactorBase2)
    CofactorBitSize := uint64(len(CofactorBase2trimmed))


    v1 := InferiorTrace(Prime,Trace)
    Q.QuoRem(v1,Cofactor,Remainder)
    Power,Sign,Rest := Power2DistanceChecker(Q)

    if Sign == false {
        X = Power - (2 + CofactorBitSize)
        Ss = "-"
    } else if Sign == true {
	X = Power - (1 + CofactorBitSize)
	Ss = "+"
    }
    PowerString := strconv.FormatInt(int64(Power),10)
    RestS := Rest.Text(10)

    Qs = "2^" + PowerString + Ss + RestS
    return X,Qs
}

func Power2DistanceChecker(Number *big.Int) (uint64, bool, *big.Int) {
    var (
    	BetweenInt = new(big.Int)
    	HalfBetween = new(big.Int)
	LowerPower = new(big.Int)
	HigherPower = new(big.Int)

	Rest = new(big.Int)
	Sign bool
	Power uint64
    )

    NumberBase2 := Number.Text(2)
    Between := TrimFirstRune(NumberBase2)
    BetweenInt.SetString(Between,2)	//22

    LowerPower.Sub(Number,BetweenInt)		//32
    HigherPower.Mul(LowerPower,Two)		//64
    HalfBetween.Quo(LowerPower,Two)		//16

    HigherPowerBin := HigherPower.Text(2)

    Cmp := BetweenInt.Cmp(HalfBetween)
    if Cmp == 1 {
        Rest.Sub(LowerPower,BetweenInt)
        Sign = false
	Power = uint64(len(HigherPowerBin))-1
    } else if Cmp == -1 {
	Rest = BetweenInt
	Sign = true
	Power = uint64(len(HigherPowerBin))-2
    } else {
	Rest = BetweenInt
	Sign = true
	Power = uint64(len(HigherPowerBin))-2
    }

    return Power, Sign, Rest
}

//Cores1 = 16 (Ryzen CPU, negative coefficients)
//Cores2 = 32 (Epyc  CPU, positive coefficients)
//Function Verifies all Result Files

//Unity 	Mode verifies the Trace for Prime ≡ 1 mod 4	Hard Verification
//Trinity	Mode verifies the Trace for Prime ≡ 3 mod 4	Hard Verification
//Easy4		Mode verifies only (p+1-t)/4 to be prime
//Easy8		Mode verifies only (p+1-t)/8 to be prime
//Easy16	Mode verifies only (p+1-t)/16 to be prime

//Hard verification verifies p+1-t div by cofactor (4 or 8) to be prime as wel as p+1+t.
//However for the Ellipse to be viable, verifying only p+1-t suffices.
//That is the Easy verification.

func VerifyEllipseResultsV2 (Location string, Prime PrimePowerTwo, RyzenCores, EpycCores uint64, Mode string) (T, F uint64) {
    var (
	Suffix		string
	RyzenV		string
	EpycV		string
	True		uint64
	False		uint64
	RyzenT		uint64
	RyzenF		uint64
	EpycT		uint64
	EpycF		uint64
    )

    Ryzen 	:= int64(RyzenCores) + int64(RyzenCores) / 2
    Epyc 	:= int64(EpycCores) + int64(EpycCores) / 2
    PrimeNumber := MakePrime(Prime)

    FirstRoot 	:= MakeEllipticExportString(Prime)
    NegRoot 	:= MakeEllipticExportStringNeg(FirstRoot)
    PosRoot 	:= MakeEllipticExportStringPos(FirstRoot)

    //Ryzen Verification
    for k:=int64(1); k<=Ryzen; k++ {
	if k < 10 {
	    Suffix = "0" + strconv.FormatInt(k, 10)
	} else {
	    Suffix = strconv.FormatInt(k, 10)
	}
	//fmt.Println("Verifying Ryzen file",k)
	RyzenV = NegRoot + Suffix + ".txt"
	True,False = VerifyEllipseResultFile(Location,RyzenV,&PrimeNumber,Mode)
	RyzenT = RyzenT + True
	RyzenF = RyzenF + False
    }
    
    //Epyc Verification
    for k:=int64(1); k<=Epyc; k++ {
	if k < 10 {
	    Suffix = "0" + strconv.FormatInt(k, 10)
	} else {
	    Suffix = strconv.FormatInt(k, 10)
	}
	//fmt.Println("Verifying Epyc file",k)
	EpycV = PosRoot + Suffix + ".txt"
	True,False = VerifyEllipseResultFile(Location,EpycV,&PrimeNumber,Mode)
	EpycT = EpycT + True
	EpycF = EpycF + False
    }
    fmt.Println("=======Summary=======")
    fmt.Println("RyzenTotal         =",RyzenT+RyzenF,"| True =",RyzenT,"| False =",RyzenF)
    fmt.Println("EpycTotal          =",EpycT+EpycF,"| True =",EpycT,"| False =",EpycF)
    T=RyzenT+EpycT
    F=RyzenF+EpycF
    fmt.Println("TotalVerifications =",T+F,"| True =",T,"| False =",F)

    return T,F
}

//Function Verifies individual Files.
func VerifyEllipseResultFile (Location, FileName string, Prime *big.Int, Mode string) (True, False uint64) {
    var (
        Trace = new(big.Int)
    )

    Lines := ReadEllipseComputationLines(Location+FileName)
    for i:=0; i<len(Lines); i++ {
	Line := Lines[i]
	LineElements := strings.Split(Line, ";")
	if len(LineElements) == 6 {
	    TraceString := LineElements[4]
	    Trace.SetString(TraceString,10)     //Trace as big Int

	    if Mode == "Unity" {
		TraceCheck := UnityTraceVerification(Prime,Trace)
		if TraceCheck == true {
		    True = True + 1
		    fmt.Println("D=",LineElements[0],"T=",LineElements[4])
		} else {
		    False = False + 1
		}
	    } else if Mode == "Trinity" {
		TraceCheck := TrinityTraceVerification(Prime,Trace)
		if TraceCheck == true {
		    True = True + 1
		    Ss,Qs := SafeScalarComputer(Prime,Trace,Four)
		    fmt.Println("Ss=",Ss,"; D=",LineElements[0],"; Q=",Qs,"; T=",LineElements[4])
		} else {
		    False = False + 1
		}
	    } else if Mode == "Easy4" {
		TraceCheck := Easy4TraceVerification(Prime,Trace)
		if TraceCheck == true {
		    True = True + 1
		    Ss,Qs := SafeScalarComputer(Prime,Trace,Four)
		    fmt.Println("Ss=",Ss,"; D=",LineElements[0],"; Q=",Qs,"; T=",LineElements[4])
		} else {
		    False = False + 1
		}
	    } else if Mode == "Easy8" {
		TraceCheck := Easy8TraceVerification(Prime,Trace)
		if TraceCheck == true {
		    True = True + 1
		    Ss,Qs := SafeScalarComputer(Prime,Trace,Eight)
		    fmt.Println("Ss=",Ss,"; D=",LineElements[0],"; Q=",Qs,"; T=",LineElements[4])
		} else {
		    False = False + 1
		}
	    } else if Mode == "Easy16" {
		TraceCheck := Easy16TraceVerification(Prime,Trace)
		if TraceCheck == true {
		    True = True + 1
		    Ss,Qs := SafeScalarComputer(Prime,Trace,Sixteen)
		    fmt.Println("Ss=",Ss,"; D=",LineElements[0],"; Q=",Qs,"; T=",LineElements[4])
		} else {
		    False = False + 1
		}
	    }
	}
    }
    return True, False
}

//Function reads all lines from an individual files, and creates a slice of strings,
// where each element of the slice is a line from within the file.
// Each line will be separated in individual elements with the help of the ";" separator
// in order to get the computed trace value
func ReadEllipseComputationLines (FileName string) []string {
    var TextLines []string

    file, err := os.Open(FileName)

    if err != nil {
	log.Fatalf("Failed opening file: %s", err)
    }

    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanLines)

    for scanner.Scan() {
	TextLines = append(TextLines, scanner.Text())
    }

    file.Close()

    //Printing read elements. Dont need it
    //for _, EachLine := range TextLines {
    //    fmt.Println(EachLine)
    //}

    return TextLines
}

func UnityTraceVerification (Prime, Trace *big.Int) (Result bool) {
    var (
        FirstTruth bool
	SecondTruth bool
    )

    InferiorTracedPrime := InferiorTrace (Prime, Trace)
    SuperiorTracedPrime := SuperiorTrace (Prime, Trace)
    //fmt.Println("Inferior is",InferiorTracedPrime)
    //fmt.Println("Superior is",SuperiorTracedPrime)

    ITPDiv4, Rest1 :=  IsDivisibleByFour(InferiorTracedPrime)
    ITPDiv8, Rest2 :=  IsDivisibleByEight(InferiorTracedPrime)

    STPDiv4, Rest3 :=  IsDivisibleByFour(SuperiorTracedPrime)
    STPDiv8, Rest4 :=  IsDivisibleByEight(SuperiorTracedPrime)

    if ITPDiv4 == true && STPDiv8 == true {
	FirstTruth = true		//Cofactor 4
    } else {
	FirstTruth = false
    }

    if ITPDiv8 == true && STPDiv4 == true {
	SecondTruth = true		//Cofactor 8
    } else {
	SecondTruth = false
    }

    if FirstTruth == true || SecondTruth == true {
        if FirstTruth == true {
	    if Rest1.ProbablyPrime(500) == true && Rest4.ProbablyPrime(50) == true {
		Result = true
	    } else {
		Result = false
	    }
	} else if SecondTruth == true {
	    if Rest2.ProbablyPrime(500) == true && Rest3.ProbablyPrime(50) == true {
		Result = true
	    } else {
		Result = false
	    }
	}
    }

    return Result
}

func TrinityTraceVerification (Prime, Trace *big.Int) (Result bool) {
    //Cofactor 4
    InferiorTracedPrime := InferiorTrace (Prime, Trace)
    SuperiorTracedPrime := SuperiorTrace (Prime, Trace)

    //fmt.Println("Inferior is",InferiorTracedPrime)
    //fmt.Println("Superior is",SuperiorTracedPrime)

    ITPDiv4, Rest1 :=  IsDivisibleByFour(InferiorTracedPrime)
    STPDiv4, Rest2 :=  IsDivisibleByFour(SuperiorTracedPrime)

    //fmt.Println("Inferior values",ITPDiv4,Rest1)
    //fmt.Println("Superior values",STPDiv4,Rest2)

    if ITPDiv4 == true && STPDiv4 == true {
        if Rest1.ProbablyPrime(500) == true && Rest2.ProbablyPrime(50) == true {
            Result = true
	} else {
	    Result = false
	}
    }

    return Result
}

func Easy4TraceVerification (Prime, Trace *big.Int) (Result bool) {
    //Cofactor 4
    InferiorTracedPrime := InferiorTrace (Prime, Trace)
    ITPDiv4, Rest1 :=  IsDivisibleByFour(InferiorTracedPrime)

    if ITPDiv4 == true {
	if Rest1.ProbablyPrime(500) == true {
	    Result = true
	} else {
	    Result = false
	}
    }

    return Result
}

func Easy8TraceVerification (Prime, Trace *big.Int) (Result bool) {
    //Cofactor 8
    InferiorTracedPrime := InferiorTrace (Prime, Trace)
    ITPDiv8, Rest1 :=  IsDivisibleByEight(InferiorTracedPrime)

    if ITPDiv8 == true {
	if Rest1.ProbablyPrime(500) == true {
	    Result = true
	} else {
	    Result = false
	}
    }

    return Result
}

func Easy16TraceVerification (Prime, Trace *big.Int) (Result bool) {
    //Cofactor 16
    InferiorTracedPrime := InferiorTrace (Prime, Trace)
    ITPDiv16, Rest1 :=  IsDivisibleBySixteen(InferiorTracedPrime)

    if ITPDiv16 == true {
	if Rest1.ProbablyPrime(500) == true {
	    Result = true
	} else {
	    Result = false
	}
    }

    return Result
}

func SuperiorTrace (Prime, Trace *big.Int) *big.Int {
    var (
	v1 = new(big.Int)
	Result = new(big.Int)
    )
    v1.Add(Prime,One)
    Result.Add(v1,Trace)
    return Result
}

func InferiorTrace (Prime, Trace *big.Int) *big.Int {
    var (
	v1 = new(big.Int)
	Result = new(big.Int)
    )
    v1.Add(Prime,One)
    Result.Sub(v1,Trace)
    return Result
}

func IsDivisibleByFour (Number *big.Int) (bool, *big.Int) {
    var (
	Quotient = new(big.Int)
	Remainder = new(big.Int)
	Result bool
    )

    Quotient.QuoRem(Number,Four,Remainder)
    Cmp := Remainder.Cmp(Zero)
    if Cmp == 0 {
        Result = true
    } else {
        Result = false
    }

    return Result, Quotient
}

func IsDivisibleByEight (Number *big.Int) (bool, *big.Int) {
    var (
	Quotient = new(big.Int)
	Remainder = new(big.Int)
	Result bool
    )

    Quotient.QuoRem(Number,Eight,Remainder)
    Cmp := Remainder.Cmp(Zero)
    if Cmp == 0 {
	Result = true
    } else {
	Result = false
    }

    return Result, Quotient
}

func IsDivisibleBySixteen (Number *big.Int) (bool, *big.Int) {
    var (
	Quotient = new(big.Int)
	Remainder = new(big.Int)
	Result bool
    )

    Quotient.QuoRem(Number,Sixteen,Remainder)
    Cmp := Remainder.Cmp(Zero)
    if Cmp == 0 {
	Result = true
    } else {
	Result = false
    }

    return Result, Quotient
}