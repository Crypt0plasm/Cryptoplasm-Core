package Firefly_Precision

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
// It has no restriction on precision. Tested Accepted characters are "." and "-" for a negative sign
func NFS(s string) *Decimal {
	d := new(Decimal)
    	//Further Tests can be added to detect wrong string construction.
    	//It is always assumed the string number used with this function is of correct construction.
	d,_,_ = NewFromString(s)
	return d
}
//================================================
//
// Function 02 - NFI
//
// NFI creates a new decimal from a string integer
// THe integer must be within int64 range
// From 	-9,223,372,036,854,775,808 to
//		 9,223,372,036,854,775,807
func NFI(coeff int64) *Decimal {
	d := new(Decimal)
	d = New(coeff,0)
	return d
}
//================================================
//
// Function 02 - INT64
//
// INT64 returns the int64 representation of x.
// No error is returned with this function
func INT64(Decimal2Convert *Decimal) int64 {
    var Result int64
    Result, _ = Decimal2Convert.Int64()
    return Result
}
//================================================