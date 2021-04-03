package Cryptoplasm_Cardinality

import "math/big"

var (
    DPmOne = big.NewInt(-1)
    DPZero = big.NewInt(0)
    DPOne = big.NewInt(1)
    Two  = big.NewInt(2)

    //Defining Y^2
    A0 = Letter{
	Letter: "A",
	Exponent: DPZero,
    }
    B0 = Letter{
	Letter: "B",
	Exponent: DPZero,
    }
    A1 = Letter{
	Letter: "A",
	Exponent: DPOne,
    }
    B1 = Letter{
	Letter: "B",
	Exponent: DPOne,
    }
    A2 = Letter{
	Letter: "A",
	Exponent: Two,
    }
    Y0 = Letter{
	Letter: "Y",
	Exponent: DPZero,
    }
    Y1 = Letter{
	Letter: "Y",
	Exponent: DPOne,
    }

    NoY = YCoefficient{DPOne,DPOne,Y0}

    YSqX0 = Coefficient{DPOne,A0,B1}
    YSqX1 = Coefficient{DPOne,A1,B0}
    YSqX2 = Coefficient{DPZero,A0,B0}
    YSqX3 = Coefficient{DPOne,A0,B0}

    YSquared = Polynom {
	3,
	[]uint64{0, 1, 2, 3},
	[][]Coefficient{
	    {YSqX0},
	    {YSqX1},
	    {YSqX2},
	    {YSqX3},
	},
    }

    POne = Polynom {
        0,
        []uint64{0},
        [][]Coefficient{{YSqX3}},
    }

    DPTwo = DivisionPolynom{
        POne,
        YCoefficient{Two,DPOne,Y1},
    }

    DP3X0 = Coefficient{DPmOne,A2,B0}
    DP3X1 = Coefficient{big.NewInt(12),A0,B1}
    DP3X2 = Coefficient{big.NewInt(6),A1,B0}
    DP3X3 = YSqX2
    DP3X4 = Coefficient{big.NewInt(3),A0,B0}

    PD3 = Polynom{
        4,
        []uint64{0,1,2,3,4},
        [][]Coefficient{
            {DP3X0},
	    {DP3X1},
	    {DP3X2},
	    {DP3X3},
	    {DP3X4},
	},
    }

    DPThree = DivisionPolynom{
        PD3,
        NoY,
    }
)