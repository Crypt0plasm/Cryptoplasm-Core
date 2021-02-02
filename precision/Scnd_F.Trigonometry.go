package precision

//
// Constants used in Trigonometry Functions
//

//250 Decimals PI
const PiString = "3.1415926535897932384626433832795028841971693993751058209749445923078164062862089986280348253421170679821480865132823066470938446095505822317253594081284811174502841027019385211055596446229489549303819644288109756659334461284756482337867831652712019091"
const HalfCircle = "180"

var PI, _, _ = NewFromString(PiString)
var HfCi, _, _ = NewFromString(HalfCircle)

//=================================================================================================
//Trigonometric Functions
//
//=================================================================================================
//
//Function 001 RADS
//
//Converts Degrees to Radians
func (c *Context) RADS(x *Decimal) *Decimal {
	rads, grads, multiplier := decimalZero, x, decimalZero
	_, _ = c.Quo(multiplier, PI, HfCi)
	_, _ = c.Mul(rads, grads, multiplier)
	return rads
}

//================================================
//
//Function 002 SIN
//
//Computes Sinus using Taylor Expansion
func (c *Context) SIN(rads *Decimal) *Decimal {
	HigherPrecision := c.Precision + 10
	QuantizeAmount := 0 - c.Precision
	//Sinus math internal Precision
	//10 units higher than the context precision
	SPC := c.WithPrecision(HigherPrecision)

	fact, _, _ := NewFromString("1")
	sign := new(Decimal)
	y := new(Decimal)
	yminusone := new(Decimal)
	factH := new(Decimal)
	pow := new(Decimal)
	NextTermH := new(Decimal)
	TermOne := new(Decimal)
	NextTerm := new(Decimal)
	SumBig := new(Decimal)
	SumSmall := new(Decimal)
	SumBigQ := new(Decimal)
	CompareRes := new(Decimal)

	//Part I
	//First term is TermZero
	//For i = 0, TermZero is rads
	TermZero := rads

	//Part II
	//Computing Term One
	//TermOne is the term for i = 1

	//i = 1, must be declared like this
	//otherwise it doesnt behave as expected
	i, _, _ := NewFromString("1")

	//Sign of the term, its minus for odd numbered terms
	//and plus for even numbered term
	//Therefore for i = 1, it must be minus
	//Here: (-1)**i (which is one) equal -1
	_, _ = SPC.Pow(sign, New(-1, 0), i)
	//fmt.Println("Sign", sign)

	//The Factorial part of the Term
	//It is computed by setting y as i + 2
	//Then doing y * (y - 1)
	//Which is further multiplied by the factorial
	//Which is (the factorial) initialised with 1
	_, _ = SPC.Add(y, i, New(2, 0))
	_, _ = SPC.Sub(yminusone, y, New(1, 0))
	_, _ = SPC.Mul(factH, y, yminusone)
	_, _ = SPC.Mul(fact, fact, factH)

	//Exponent is rads ** y
	//Since y is increased by 2 each iteration
	//It can serve as exponent for x (radians)
	_, _ = SPC.Pow(pow, rads, y)
	//fmt.Println("pow este ", pow)

	//Now that we have all the components computed
	//We can use them to compute the Term
	//A Helper variable is used because of apd Syntax
	//The Term is (sign / fact) * pow
	_, _ = SPC.Quo(NextTermH, sign, fact)
	_, _ = SPC.Mul(TermOne, NextTermH, pow)

	//Now that we have TermOne, we can use it
	//Together with TermZero to compute their Sum
	//This Sum is designated SumBig
	_, _ = SPC.Add(SumBig, TermZero, TermOne)

	//Next we compare the Sum thus resulted with TermZero
	//If they would be identical (which is the condition
	//for the Loop to break, CompareRes would be Zero
	//Therefore "Truth" (boolean) would be true
	_, _ = SPC.Cmp(CompareRes, TermZero, SumBig)
	Truth := CompareRes.IsZero()

	//So long as the two values Compared are different
	//The Loop continues
	//Initially TermZero is compared against TermZero+TermOne
	//As these are different initially, the loop will enter its first iteration
	//
	//Once inside the loop, the previous SumBig = TermZero+TermOne
	//Is made into SumSmall as it is quantized with Cryptoplasm Precision
	//This ensures it can be compared against at the end of the loop
	//
	//Inside the Loop, the Next Term is being calculated
	//Once it is computed, it is added to the SumSmall which was quantized at 250 decimals
	//The Result ist termed SumBig, and is quantized before comparing it to SumSmall
	//At the end of the loop
	//
	//The loop continues for so many iterations until the
	//Quantized SumSmall and SumBig are equal up to the last decimal specified by CryptoplasmPrecision
	//Once this happens the loop is broken, and the SumBig is returned
	//The SumBig is the Sinus Result obtained from the input in radians
	//
	//As the internal computations are done with 10 more decimal precision than CryptoplasmPrecision
	//The result should be precise to CryptoplasmPrecision
	for Truth == false {
		//for kk := 1; kk < 5; kk++ {

		//Quantizing SmallSum in preparation for the Comparison
		//at the end of the loop
		_, _ = SPC.Quantize(SumSmall, SumBig, int32(QuantizeAmount))

		//Incrementing i: i = i + 1
		_, _ = SPC.Add(i, i, New(1, 0))
		//Computing the sign
		_, _ = SPC.Pow(sign, New(-1, 0), i)
		//Computing Factorial
		_, _ = SPC.Add(y, y, New(2, 0))
		_, _ = SPC.Sub(yminusone, y, New(1, 0))
		_, _ = SPC.Mul(factH, y, yminusone)
		_, _ = SPC.Mul(fact, fact, factH)
		//Computing the Exponent
		_, _ = SPC.Pow(pow, rads, y)

		//Now that we have all the components computed
		//We can use them to compute the Term
		_, _ = SPC.Quo(NextTermH, sign, fact)
		_, _ = SPC.Mul(NextTerm, NextTermH, pow)

		//Once we have the NextTerm, we add it to the SumSmall
		//Obtaining the next sum which is preciser than the one before it
		_, _ = SPC.Add(SumBig, SumSmall, NextTerm)
		//The Sum is quantized to CryptoplasmPrecision precision
		//Its result is saved in SumBigQ parameter
		_, _ = SPC.Quantize(SumBigQ, SumBig, int32(QuantizeAmount))

		//The end loop comparison takes place between
		//SumSmall and SumBigQ
		//Once they are perfectly equal to the last decimal
		//Specified by CryptoplasmPrecision
		//The loop will be broken and the Sinus will have been calculated.
		_, _ = SPC.Cmp(CompareRes, SumSmall, SumBigQ)
		Truth = CompareRes.IsZero()
	}
	return SumBigQ
}
