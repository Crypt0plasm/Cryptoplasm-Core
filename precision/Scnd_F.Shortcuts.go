package precision

//
// Shortcuts functions of Main Functions
// Eases up function usage, discarding the Condition and Error outputs
//
//================================================
//
// Function 01 - NFS
//
// NFS creates a new decimal from a string s.
// Same as NewFromString, it only outputs the decimal
func NFS(s string) *Decimal {
	d := new(Decimal)
	d, _, _ = NewFromString(s)
	return d
}

//================================================
//
// Function 02 - NFI
//
// NFI creates a new decimal from a string integer
// THe integer must be within int64 range
// From 	-9,223,372,036,854,775,808 to
//			 9,223,372,036,854,775,807
func NFI(coeff int64) *Decimal {
	d := new(Decimal)
	d = New(coeff, 0)
	return d
}

//================================================
