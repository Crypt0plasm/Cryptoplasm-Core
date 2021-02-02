package precision

import "math/big"

// All Cryptoplasm_Firefly_Precision Spawner Functions are listed below
// Spawner Functions create a Decimal from another type of variable
//=================================================================================================
//
// Function 01 - NewFromString
//
// NewFromString creates a new decimal from a string s.
// It has no restrictions on exponents or precision.
//
func NewFromString(s string) (*Decimal, Condition, error) {
	return BaseContext.NewFromString(s)
}

//================================================
//
// Function 02 - New
//
// New creates a new decimal with the given coefficient and exponent.
func New(coeff int64, exponent int32) *Decimal {
	d := &Decimal{
		Negative: coeff < 0,
		Coeff:    *big.NewInt(coeff),
		Exponent: exponent,
	}
	d.Coeff.Abs(&d.Coeff)
	return d
}

//================================================
//
// Function 03 - NewWithBigInt
//
// NewWithBigInt creates a new decimal with the given coefficient and exponent.
func NewWithBigInt(coeff *big.Int, exponent int32) *Decimal {
	d := &Decimal{
		Exponent: exponent,
	}
	d.Coeff.Set(coeff)
	if d.Coeff.Sign() < 0 {
		d.Negative = true
		d.Coeff.Abs(&d.Coeff)
	}
	return d
}
