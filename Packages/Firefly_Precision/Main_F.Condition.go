package Firefly_Precision

import (
	"github.com/pkg/errors"
	"strings"
)

//=================================================================================================
//=================================================================================================
// All Cryptoplasm_Firefly_Precision Condition Functions are listed below:
//=================================================================================================
//
// Function 01 - Any
//
// returns true if any flag is true.
func (r Condition) Any() bool { return r != 0 }
//================================================
//
// Function 02 - Clamped
//
// returns true if the Clamped flag is set.
func (r Condition) Clamped() bool { return r&Clamped != 0 }
//================================================
//
// Function 03 - DivisionByZero
//
// returns true if the DivisionByZero flag is set.
func (r Condition) DivisionByZero() bool { return r&DivisionByZero != 0 }
//================================================
//
// Function 04 - DivisionImpossible
//
// returns true if the DivisionImpossible flag is set.
func (r Condition) DivisionImpossible() bool { return r&DivisionImpossible != 0 }
//================================================
//
// Function 05 - DivisionUndefined
//
// DivisionUndefined returns true if the DivisionUndefined flag is set.
func (r Condition) DivisionUndefined() bool { return r&DivisionUndefined != 0 }
//================================================
//
//Function 06 - GoError
//
// converts r to an error based on the given traps and returns
// r. Traps are the conditions which will trigger an error result if the
// corresponding Flag condition occurred.
func (r Condition) GoError(traps Condition) (Condition, error) {
	const (
		systemErrors = SystemOverflow | SystemUnderflow
	)
	var err error
	if r&systemErrors != 0 {
		err = errors.New(errExponentOutOfRangeStr)
	} else if t := r & traps; t != 0 {
		err = errors.New(t.String())
	}
	return r, err
}
//================================================
//
// Function 07 - DivisionUndefined
//
// returns true if the Inexact flag is set.
func (r Condition) Inexact() bool { return r&Inexact != 0 }
//================================================
//
// Function 08 - InvalidOperation
//
// returns true if the InvalidOperation flag is set.
func (r Condition) InvalidOperation() bool { return r&InvalidOperation != 0 }
//================================================
//
// Function 09 - Overflow
//
// returns true if the Overflow flag is set.
func (r Condition) Overflow() bool { return r&Overflow != 0 }
//================================================
//
// Function 10 - Rounded
//
// returns true if the Rounded flag is set.
func (r Condition) Rounded() bool { return r&Rounded != 0 }
//================================================
//
// Function 11 - String
//
// returns the Condition strings.
func (r Condition) String() string {
	var names []string
	for i := Condition(1); r != 0; i <<= 1 {
		if r&i == 0 {
			continue
		}
		r ^= i
		var s string
		switch i {
		case SystemOverflow, SystemUnderflow:
			continue
		case Overflow:
			s = "overflow"
		case Underflow:
			s = "underflow"
		case Inexact:
			s = "inexact"
		case Subnormal:
			s = "subnormal"
		case Rounded:
			s = "rounded"
		case DivisionUndefined:
			s = "division undefined"
		case DivisionByZero:
			s = "division by zero"
		case DivisionImpossible:
			s = "division impossible"
		case InvalidOperation:
			s = "invalid operation"
		case Clamped:
			s = "clamped"
		default:
			panic(errors.Errorf("unknown condition %d", i))
		}
		names = append(names, s)
	}
	return strings.Join(names, ", ")
}
//================================================
//
// Function 12 - Subnormal
//
// returns true if the Subnormal flag is set.
func (r Condition) Subnormal() bool { return r&Subnormal != 0 }
//================================================
//
// Function 13 - SystemOverflow
//
// returns true if the SystemOverflow flag is set.
func (r Condition) SystemOverflow() bool { return r&SystemOverflow != 0 }
//================================================
//
// Function 14 - Underflow
//
// returns true if the Underflow flag is set.
func (r Condition) Underflow() bool { return r&Underflow != 0 }
//================================================
//
// Function 15 - negateOverflowFlags
//
// converts Overflow and SystemOverflow flags into their equivalent Underflows.
func (r Condition) negateOverflowFlags() Condition {
	if r.Overflow() {
		// Underflow always also means Subnormal. See GDA definition.
		r |= Underflow | Subnormal
		r &= ^Overflow
	}
	if r.SystemOverflow() {
		r |= SystemUnderflow
		r &= ^SystemOverflow
	}
	return r
}
//=================================================================================================
//=================================================================================================