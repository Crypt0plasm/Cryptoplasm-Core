package Firefly_Precision

import "fmt"

//=================================================================================================
//=================================================================================================
// All Cryptoplasm_Firefly_Precision ErrDecimal Functions are listed below:
//=================================================================================================
//
// Function 01 - Abs
//
// Abs performs e.Ctx.Abs(d, x) and returns d.
func (e *ErrDecimal) Abs(d, x *Decimal) *Decimal {
	return e.op2(d, x, e.Ctx.Abs)
}
//================================================
//
// Function 02 - Add
//
// Add performs e.Ctx.Add(d, x, y) and returns d.
func (e *ErrDecimal) Add(d, x, y *Decimal) *Decimal {
	return e.op3(d, x, y, e.Ctx.Add)
}
//================================================
//
// Function 03 - Ceil
//
// Ceil performs e.Ctx.Ceil(d, x) and returns d.
func (e *ErrDecimal) Ceil(d, x *Decimal) *Decimal {
	return e.op2(d, x, e.Ctx.Ceil)
}
//================================================
//
// Function 04 - Err
//
// Err returns the first error encountered or the context's trap error
// if present.
func (e *ErrDecimal) Err() error {
	if e.err != nil {
		return e.err
	}
	if e.Ctx != nil {
		_, e.err = e.Ctx.goError(e.Flags)
		return e.err
	}
	return nil
}
//================================================
//
// Function 05 - Exp
//
// Exp performs e.Ctx.Exp(d, x) and returns d.
func (e *ErrDecimal) Exp(d, x *Decimal) *Decimal {
	return e.op2(d, x, e.Ctx.Exp)
}
//================================================
//
// Function 06 - Floor
//
// Floor performs e.Ctx.Floor(d, x) and returns d.
func (e *ErrDecimal) Floor(d, x *Decimal) *Decimal {
	return e.op2(d, x, e.Ctx.Floor)
}
//================================================
//
// Function 07 - Int64
//
// Int64 returns 0 if err is set. Otherwise returns d.Int64().
func (e *ErrDecimal) Int64(d *Decimal) int64 {
	if e.Err() != nil {
		return 0
	}
	var r int64
	r, e.err = d.Int64()
	return r
}
//================================================
//
// Function 08 - Ln
//
// Ln performs e.Ctx.Ln(d, x) and returns d.
func (e *ErrDecimal) Ln(d, x *Decimal) *Decimal {
	return e.op2(d, x, e.Ctx.Ln)
}
//================================================
//
// Function 09 - Log10
//
// Log10 performs d.Log10(x) and returns d.
func (e *ErrDecimal) Log10(d, x *Decimal) *Decimal {
	return e.op2(d, x, e.Ctx.Log10)
}
//================================================
//
// Function 10 - Mul
//
// Mul performs e.Ctx.Mul(d, x, y) and returns d.
func (e *ErrDecimal) Mul(d, x, y *Decimal) *Decimal {
	return e.op3(d, x, y, e.Ctx.Mul)
}
//================================================
//
// Function 11 - Neg
//
// Neg performs e.Ctx.Neg(d, x) and returns d.
func (e *ErrDecimal) Neg(d, x *Decimal) *Decimal {
	return e.op2(d, x, e.Ctx.Neg)
}
//================================================
//
// Function 12 - Pow
//
// Pow performs e.Ctx.Pow(d, x, y) and returns d.
func (e *ErrDecimal) Pow(d, x, y *Decimal) *Decimal {
	return e.op3(d, x, y, e.Ctx.Pow)
}
//================================================
//
// Function 13 - Quantize
//
// Quantize performs e.Ctx.Quantize(d, v, exp) and returns d.
func (e *ErrDecimal) Quantize(d, v *Decimal, exp int32) *Decimal {
	if e.Err() != nil {
		return d
	}
	res, err := e.Ctx.Quantize(d, v, exp)
	e.Flags |= res
	e.err = err
	return d
}
//================================================
//
// Function 14 - Quo
//
// Quo performs e.Ctx.Quo(d, x, y) and returns d.
func (e *ErrDecimal) Quo(d, x, y *Decimal) *Decimal {
	return e.op3(d, x, y, e.Ctx.Quo)
}
//================================================
//
// Function 15 - QuoInteger
//
// QuoInteger performs e.Ctx.QuoInteger(d, x, y) and returns d.
func (e *ErrDecimal) QuoInteger(d, x, y *Decimal) *Decimal {
	return e.op3(d, x, y, e.Ctx.QuoInteger)
}
//================================================
//
// Function 16 - Reduce
//
// Reduce performs e.Ctx.Reduce(d, x) and returns
// the number of zeros removed and d.
func (e *ErrDecimal) Reduce(d, x *Decimal) (int, *Decimal) {
	if e.Err() != nil {
		return 0, d
	}
	n, res, err := e.Ctx.Reduce(d, x)
	e.Flags |= res
	e.err = err
	return n, d
}
//================================================
//
// Function 17 - Rem
//
// Rem performs e.Ctx.Rem(d, x, y) and returns d.
func (e *ErrDecimal) Rem(d, x, y *Decimal) *Decimal {
	return e.op3(d, x, y, e.Ctx.Rem)
}
//================================================
//
// Function 18 - Round
//
// Round performs e.Ctx.Round(d, x) and returns d.
func (e *ErrDecimal) Round(d, x *Decimal) *Decimal {
	return e.op2(d, x, e.Ctx.Round)
}
//================================================
//
// Function 19 - RoundToIntegralExact
//
// performs e.Ctx.RoundToIntegralExact(d, x) and returns d.
func (e *ErrDecimal) RoundToIntegralExact(d, x *Decimal) *Decimal {
	return e.op2(d, x, e.Ctx.RoundToIntegralExact)
}
//================================================
//
// Function 20 - RoundToIntegralValue
//
// performs e.Ctx.RoundToIntegralValue(d, x) and returns d.
func (e *ErrDecimal) RoundToIntegralValue(d, x *Decimal) *Decimal {
	return e.op2(d, x, e.Ctx.RoundToIntegralValue)
}
//================================================
//
// Function 21 - Sqrt
//
// Sqrt performs e.Ctx.Sqrt(d, x) and returns d.
func (e *ErrDecimal) Sqrt(d, x *Decimal) *Decimal {
	return e.op2(d, x, e.Ctx.Sqrt)
}
//================================================
//
// Function 21 - Sub
//
// Sub performs e.Ctx.Sub(d, x, y) and returns d.
func (e *ErrDecimal) Sub(d, x, y *Decimal) *Decimal {
	return e.op3(d, x, y, e.Ctx.Sub)
}
//================================================
//
// Function 22 - String
//
func (i Form) String() string {
	if i < 0 || i >= Form(len(formIndex)-1) {
		return fmt.Sprintf("Form(%d)", i)
	}
	return formName[formIndex[i]:formIndex[i+1]]
}
//=================================================================================================
//=================================================================================================
// All Cryptoplasm_Firefly_Precision ErrDecimal Secondary Functions are listed below:
//================================================
//
// Function op2 - String
//
func (e *ErrDecimal) op2(d, x *Decimal, f func(a, b *Decimal) (Condition, error)) *Decimal {
	if e.Err() != nil {
		return d
	}
	res, err := f(d, x)
	e.Flags |= res
	e.err = err
	return d
}
//================================================
//
// Function op2 - String
//
func (e *ErrDecimal) op3(d, x, y *Decimal, f func(a, b, c *Decimal) (Condition, error)) *Decimal {
	if e.Err() != nil {
		return d
	}
	res, err := f(d, x, y)
	e.Flags |= res
	e.err = err
	return d
}



