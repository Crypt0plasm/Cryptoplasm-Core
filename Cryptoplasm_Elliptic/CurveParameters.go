package Cryptoplasm_Elliptic

import "math/big"

// Package Cryptoplasm_Elliptic contains several implementations of Twisted Edwards Curves,
// from general and unoptimized to highly specialized and optimized.
//
// Twisted Edwards curves are elliptic curves satisfying the equation:
//
//	a * x^2 + y^2 = c^2 * (1 + d * x^2 * y^2)
//
// for some scalars c, d over some field K. We assume K is a (finite) prime field for a
// large prime p. We also assume c == 1 because all curves in the generalized form
// are isomorphic to curves having c == 1.
//
// For details see Bernstein et al, "Twisted Edwards Curves", http://eprint.iacr.org/2008/013.pdf

// Param defines a Twisted Edwards curve (TEC).
type FiniteFieldEllipticCurve struct {
    Name string 		// Name of curve

    P big.Int 			// Prime defining the underlying field
    Q big.Int 			// Order of the prime-order base point
    R int     			// Cofactor: Q*R is the total size of the curve

    A, D big.Int 		// Edwards curve equation parameters

    FBX, FBY big.Int 		// Standard base point for full group
    PBX, PBY big.Int 		// Standard base point for prime-order subgroup

    Elligator1s big.Int 	// Optional s parameter for Elligator 1
    Elligator2u big.Int 	// Optional u parameter for Elligator 2
}
// ParamE382 defines the E-382 curve specified in:
// Aranha et al, "A note on high-security general-purpose elliptic curves",
// http://eprint.iacr.org/2013/647.pdf
//
// and more recently in:
//
// "Additional Elliptic Curves for IETF protocols"
// http://tools.ietf.org/html/draft-ladd-safecurves-02
//
// Name: 			E-382
// Equation is: 		x^2 + y^2 = 1 - 67254 * x^2 * y^2
// Prime field (P)		2^382 - 105
//	Decimal Form		9850501549098619803069760025035903451269934817616361666987073351061430442874302652853566563721228910201656997576599
//	Hexadecimal Form	0x3fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff97
//
// Base-Point Prime Order (Q)	2^380 - 1030303207694556153926491950732314247062623204330168346855
//	Decimal Form		2462625387274654950767440006258975862817483704404090416745738034557663054564649171262659326683244604346084081047321
//	Hexadecimal Form	0xfffffffffffffffffffffffffffffffffffffffffffffffd5fb21f21e95eee17c5e69281b102d2773e27e13fd3c9719
//
// Cofactor R			8
// Curve A Parameter		1
// Curve D Parameter		-67254
//
// Base-Point X			3914921414754292646847594472454013487047137431784830634731377862923477302047857640522480241298429278603678181725699
//	Hexadecimal Form	0x196f8dd0eab20391e5f05be96e8d20ae68f840032b0b64352923bab85364841193517dbce8105398ebc0cc9470f79603
// Base-Point Y			17
//	Hexadecimal Form	0x11
func ParamE382() *FiniteFieldEllipticCurve {
    var p FiniteFieldEllipticCurve
    var qs big.Int
    p.Name = "E-382"
    p.P.SetBit(Zero, 382, 1).Sub(&p.P, big.NewInt(105))
    qs.SetString("1030303207694556153926491950732314247062623204330168346855", 10)
    p.Q.SetBit(Zero, 380, 1).Sub(&p.Q, &qs)
    p.R = 8
    p.A.SetInt64(1)
    p.D.SetInt64(-67254)
    p.PBX.SetString("3914921414754292646847594472454013487047137431784830634731377862923477302047857640522480241298429278603678181725699", 10)
    p.PBY.SetString("17", 10)
    return &p
}

// Param41417 defines the Curve41417 curve, as specified in:
// Bernstein et al, "Curve41417: Karatsuba revisited",
// http://eprint.iacr.org/2014/526.pdf
//
// Name: 			Curve41417
// Equation is: 		x^2 + y^2 = 1 + 3617 * x^2 * y^2
// Prime field (P)		2^414 - 17
//	Decimal Form		42307582002575910332922579714097346549017899709713998034217522897561970639123926132812109468141778230245837569601494931472367
//	Hexadecimal Form	0x3fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffef
//
// Base-Point Prime Order (Q)	2^411 - 33364140863755142520810177694098385178984727200411208589594759
//	Decimal Form		5288447750321988791615322464262168318627237463714249754277190328831105466135348245791335989419337099796002495788978276839289
//	Hexadecimal Form	0x7ffffffffffffffffffffffffffffffffffffffffffffffffffeb3cc92414cf706022b36f1c0338ad63cf181b0e71a5e106af79
//
// Cofactor R			8
// Curve A Parameter		1
// Curve D Parameter		3617
//
// Base-Point X			17319886477121189177719202498822615443556957307604340815256226171904769976866975908866528699294134494857887698432266169206165
//	Hexadecimal Form	0x1a334905141443300218c0631c326e5fcd46369f44c03ec7f57ff35498a4ab4d6d6ba111301a73faa8537c64c4fd3812f3cbc595
// Base-Point Y			34
//	Hexadecimal Form	0x22
func Param41417() *FiniteFieldEllipticCurve {
    var p FiniteFieldEllipticCurve
    var qs big.Int
    p.Name = "Curve41417"
    p.P.SetBit(Zero, 414, 1).Sub(&p.P, big.NewInt(17))
    qs.SetString("33364140863755142520810177694098385178984727200411208589594759", 10)
    p.Q.SetBit(Zero, 411, 1).Sub(&p.Q, &qs)
    p.R = 8
    p.A.SetInt64(1)
    p.D.SetInt64(3617)
    p.PBX.SetString("17319886477121189177719202498822615443556957307604340815256226171904769976866975908866528699294134494857887698432266169206165", 10)
    p.PBY.SetString("34", 10)
    return &p
}

// ParamE521 defines the E-521 curve specified in:
// Aranha et al, "A note on high-security general-purpose elliptic curves",
// http://eprint.iacr.org/2013/647.pdf
//
// and more recently included in:
// "Additional Elliptic Curves for IETF protocols"
// http://tools.ietf.org/html/draft-ladd-safecurves-02
//
// Name: 			E-521
// Equation is: 		x^2 + y^2 = 1 - 376014 * x^2 * y^2
// Prime field (P)		2^521 - 1
//	Decimal Form		6864797660130609714981900799081393217269435300143305409394463459185543183397656052122559640661454554977296311391480858037121987999716643812574028291115057151
//	Hexadecimal Form	0x1ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff
//
// Base-Point Prime Order (Q)	2^519 - 337554763258501705789107630418782636071904961214051226618635150085779108655765
//	Decimal Form		1716199415032652428745475199770348304317358825035826352348615864796385795849413675475876651663657849636693659065234142604319282948702542317993421293670108523
//	Hexadecimal Form	0x7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffd15b6c64746fc85f736b8af5e7ec53f04fbd8c4569a8f1f4540ea2435f5180d6b
//
// Cofactor R			8
// Curve A Parameter		1
// Curve D Parameter		-376014
//
// Base-Point X			1571054894184995387535939749894317568645297350402905821437625181152304994381188529632591196067604100772673927915114267193389905003276673749012051148356041324
//	Hexadecimal Form	0x752cb45c48648b189df90cb2296b2878a3bfd9f42fc6c818ec8bf3c9c0c6203913f6ecc5ccc72434b1ae949d568fc99c6059d0fb13364838aa302a940a2f19ba6c
// Base-Point Y			12
//	Hexadecimal Form	0xc
func ParamE521() *FiniteFieldEllipticCurve {
    var p FiniteFieldEllipticCurve
    var qs big.Int
    p.Name = "E-521"
    p.P.SetBit(Zero, 521, 1).Sub(&p.P, One)
    qs.SetString("337554763258501705789107630418782636071904961214051226618635150085779108655765", 10)
    p.Q.SetBit(Zero, 519, 1).Sub(&p.Q, &qs)
    p.R = 8
    p.A.SetInt64(1)
    p.D.SetInt64(-376014)
    p.PBX.SetString("1571054894184995387535939749894317568645297350402905821437625181152304994381188529632591196067604100772673927915114267193389905003276673749012051148356041324", 10)
    p.PBY.SetString("12", 10)
    return &p
}