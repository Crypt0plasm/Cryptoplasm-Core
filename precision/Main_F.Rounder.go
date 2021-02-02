package precision

import "math/big"

//=================================================================================================
//=================================================================================================
// All Cryptoplasm_Firefly_Precision Rounder Functions are listed below:
//=================================================================================================
//=================================================================================================
//
// Function 01 - Rounder
//
// Round sets d to rounded x.
func (r Rounder) Round(c *Context, d, x *Decimal) Condition {
	d.Set(x)
	nd := x.NumDigits()
	xs := x.Sign()
	var res Condition

	// adj is the adjusted exponent: exponent + clength - 1
	if adj := int64(x.Exponent) + nd - 1; xs != 0 && adj < int64(c.MinExponent) {
		// Subnormal is defined before rounding.
		res |= Subnormal
		// setExponent here to prevent double-rounded subnormals.
		res |= d.setExponent(c, res, int64(d.Exponent))
		return res
	}

	diff := nd - int64(c.Precision)
	if diff > 0 {
		if diff > MaxExponent {
			return SystemOverflow | Overflow
		}
		if diff < MinExponent {
			return SystemUnderflow | Underflow
		}
		res |= Rounded
		y := new(big.Int)
		e := tableExp10(diff, y)
		m := new(big.Int)
		y.QuoRem(&d.Coeff, e, m)
		if m.Sign() != 0 {
			res |= Inexact
			discard := NewWithBigInt(m, int32(-diff))
			if r(y, x.Negative, discard.Cmp(decimalHalf)) {
				roundAddOne(y, &diff)
			}
		}
		d.Coeff = *y
	} else {
		diff = 0
	}
	res |= d.setExponent(c, res, int64(d.Exponent), diff)
	return res
}
