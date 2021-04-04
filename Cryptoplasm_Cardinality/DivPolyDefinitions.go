package Cryptoplasm_Cardinality

import "math/big"

var (
    //Division Polynom Minus One
    DPmOne = big.NewInt(-1)

    //Division Polynom 0
    DPZero = big.NewInt(0)

    //Division Polynom One
    DPOne = big.NewInt(1)

    Two  = big.NewInt(2)
    Three = big.NewInt(3)

    //Defining Letters
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
    B2 = Letter{
	Letter: "B",
	Exponent: Two,
    }
    A3 = Letter{
	Letter: "A",
	Exponent: Three,
    }
    Y0 = Letter{
	Letter: "Y",
	Exponent: DPZero,
    }
    Ym1 = Letter{
	Letter: "Y",
	Exponent: DPmOne,
    }
    Y1 = Letter{
	Letter: "Y",
	Exponent: DPOne,
    }
    Y3 = Letter{
	Letter: "Y",
	Exponent: Three,
    }


    //No Y Coefficient
    NoY = YCoefficient{Two,DPZero,Y0}

    //Defining Y^2
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

    // 1/2Y as DivisionPolynom
    OneUnder2YDP = DivisionPolynom{
	POne,
	YCoefficient{Two,DPmOne,Ym1},
    }

    //Division Polynom One, defined as Polynom, used in the Definition of Division Polynom 2
    POne = Polynom {
        0,
        []uint64{0},
        [][]Coefficient{{YSqX3}},
    }

    //Division Polynom Two
    DPTwo = DivisionPolynom{
        POne,
        YCoefficient{Two,DPOne,Y1},
    }






    //Division Polynom Three
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

    //Division Polynom Four
    DP4X01 = Coefficient{big.NewInt(-8),A0,B2}
    DP4X02 = Coefficient{big.NewInt(-1),A3,B0}
    DP4X1 = Coefficient{big.NewInt(-4),A1,B1}
    DP4X2 = Coefficient{big.NewInt(-5),A2,B0}
    DP4X3 = Coefficient{big.NewInt(20),A0,B1}
    DP4X4 = Coefficient{big.NewInt(5),A1,B0}
    DP4X5 = YSqX2
    DP4X6 = YSqX3

    PD4 = Polynom{
	6,
	[]uint64{0,1,2,3,4,5,6},
	[][]Coefficient{
	    {DP4X01,DP4X02},
	    {DP4X1},
	    {DP4X2},
	    {DP4X3},
	    {DP4X4},
	    {DP4X5},
	    {DP4X6},
	},
    }

    DP4 = DivisionPolynom{
	PD4,
	YCoefficient{big.NewInt(2),Two,Y1},
    }
)