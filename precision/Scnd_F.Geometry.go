package precision

//===================================================================
//
//Function 001 ASA
//
//Solves a triangle given the Angle(alfa), Side(c), Angle(beta);
//and outputs Angle(gamma), Side(a), Side(B), and Area(area).
func (c *Context) ASA(angleAlfa, sideC, angleBeta *Decimal) (*Decimal, *Decimal, *Decimal, *Decimal) {
	HigherPrecision := c.Precision + 10
	QuantizeAmount := 0 - c.Precision
	//Sinus math internal Precision
	//10 units higher than the context precision
	SPC := c.WithPrecision(HigherPrecision)

	angleA := angleAlfa
	angleB := angleBeta
	sdC := sideC
	//fmt.Println("InsideFunction", angleA)
	//fmt.Println("InsideFunction", angleB)
	//fmt.Println("InsideFunction", sdC)

	radsA := new(Decimal)
	radsB := new(Decimal)
	radsG := new(Decimal)
	angleG := new(Decimal)
	helper00 := new(Decimal)
	helper01 := new(Decimal)
	helper02 := new(Decimal)
	helper03 := new(Decimal)
	helper04 := new(Decimal)
	helper05 := new(Decimal)
	helper06 := new(Decimal)
	helper07 := new(Decimal)
	helper08 := new(Decimal)
	helper09 := new(Decimal)
	helper10 := new(Decimal)
	sdA := new(Decimal)
	sdB := new(Decimal)
	sp := new(Decimal)
	area := new(Decimal)

	//helper02 := new(Decimal)

	_, _ = c.Add(helper00, angleA, angleB)
	_, _ = c.Sub(angleG, HfCi, helper00)

	_, _ = SPC.Quantize(radsA, c.RADS(angleA), int32(QuantizeAmount))
	_, _ = SPC.Quantize(radsB, c.RADS(angleB), int32(QuantizeAmount))
	_, _ = SPC.Quantize(radsG, c.RADS(angleG), int32(QuantizeAmount))

	_, _ = c.Mul(helper01, sdC, c.SIN(radsA))
	_, _ = c.Quo(sdA, helper01, c.SIN(radsG))
	_, _ = c.Mul(helper02, sdC, c.SIN(radsB))
	_, _ = c.Quo(sdB, helper02, c.SIN(radsG))

	_, _ = c.Add(helper03, sdA, sdB)
	_, _ = c.Add(helper04, helper03, sdC)
	_, _ = c.Quo(sp, helper04, New(2, 0))

	_, _ = c.Sub(helper05, sp, sdA)
	_, _ = c.Sub(helper06, sp, sdB)
	_, _ = c.Sub(helper07, sp, sdC)
	_, _ = c.Mul(helper08, helper05, helper06)
	_, _ = c.Mul(helper09, helper07, sp)
	_, _ = c.Mul(helper10, helper08, helper09)
	_, _ = c.Sqrt(area, helper10)

	return angleG, sdA, sdB, area
}
