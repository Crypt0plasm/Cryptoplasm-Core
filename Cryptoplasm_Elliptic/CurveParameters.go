package Cryptoplasm_Elliptic

import (
    "math/big"
)
//
// Package Cryptoplasm_Elliptic
// implements Elliptic Point Addition and Multiplication only for the Twisted Edwards Curves.
// Elliptic Point Multiplication is done with a 7^2 Window, by converting the scalar in base 49,
// and using 7^2-1 precomputed points saved in a precomputed [7][7]Matrix.
//
// Although not used in the code, AdditionZ2OneV2, DoubleWithZOne functions are implemented.
// All elliptic point math (Adding, Doubling, Tripling) are taken from http://hyperelliptic.org/EFD/.
//
// Package Cryptoplasm_Elliptic defines:
//      3 Montgomery Curves
//          M383                                        //Cofactor 8
//          Curve383187                                 //Cofactor 8
//          M511                                        //Cofactor 8
//      4 Standard known Twisted Edwards Curves,
//          E382                                        //Cofactor 4
//          Curve41417                                  //Cofactor 8
//          Goldilocks                                  //Cofactor 4
//          E521                                        //Cofactor 4
//      as defined by Safe Curves https://safecurves.cr.yp.to/index.html, and
//      4 novel never before used Twisted Edwards Curves,
//          TEC_S545_Pr551p335_m1874                    //Cofactor 4
//          TEC_S1021_Pr1029p639_p403                   //Cofactor 8
//          TEC_S1023_Pr1029p639_m200                   //Cofactor 4
//          TEC_S1024_Pr1029p639_m729                   //Cofactor 4
//          TEC_S1600_Pr1605p2315_m26                   //Cofactor 4
//
//  The idea behind the Novel TEC curves, is that they have a higher number of Possible unique Private Keys.
//  Bitcoin uses secp256k1 Curve, the Generator Point's order is  2^256 - 432420386565659656852420866394968145599.
//  Since secp256k1 has a Cofactor of 1, one could theoretically use 2^256 - 432420386565659656852420866394968145599
//  different scalars as private keys. In practice the Bitcoin Private Key is a 256 bit, 32 byte number.
//
//  While Curve E521 has an S of 515 meaning 2^515 different private key, the TEC_1617m1874 has an S of 545, allowing
//  2^545 different private key (before creating the clamped scalar which is needed because it has a cofactor of 4).
//  This is several hundreds orders of magnitude larger than the number of possible Bitcoin Private keys.
//
//  Also,and ever larger prime field makes solving the DLP even harder, and increases the number of qubits required to
//  break it using a quantum computer. Cryptoplasm is aiming for an S of at least 1024 (meaning 2^1024 different private keys)
//  wherewith, after clamping the Scalar would become a 1027 bit size number (if a cofactor of 4 elliptic curve will be used)
//  or a 1028 bit size number (for a cofactor of 8).
//
//  Since an S of 1024 would allow 2^1024 different private keys, a private Key could be represented by a grid of
//  32x32 squares, where a filled square would be a 1 bit and an empty square would be a 0 bit.
//  From this point on, for bigger curves with bigger S values, the Private Key could be represented by an ever increasing,
//  ever bigger grid of full and empty squares.
//
//  Curve TEC_S545_Pr551p335_m1874 has an S value of 545, which can be represented by an intercalated grid of 16*16 and 17*17
//  diagonal squares; Pr551p335 represent the prime field, 2^551+335, whereas m1875 represents the D coefficient (minus 1875).
//  16*16+17*17=545.
//
//  The Novel TEC curves creation aim is based on getting a higher S, while keeping Rigidity and Safety Criteria as
//  defined by https://safecurves.cr.yp.to/rigid.html.
//      The prime field is chosen to be the closest to a power of 2 while being congruent to 3 mod 4
//          (optimal condition for cofactor 4 TEC curves with A coefficient 1)
//      Cofactor 4 Curves are searched for, this is the smallest possible Cofactor for TEC curves, to maximize Q size.
//      D is chosen to be the smallest absolute value that produces the intended S size,
//      while the largest Subgroup Order Q is of course prime.
//      The Generator is then chosen as the Point with the smallest X Coordinate that lies on the Prime Order Subgroup G
//      Method MakeGenMin return this point. However its return is saved directly as values.
//
//  All Novel TEC Curves have the "Complex-Multiplication Field Discriminant" larger than 2^100 as required by https://safecurves.cr.yp.to/disc.html
//  Their "rho cost" is higher than every curve listed by SafeCurves https://safecurves.cr.yp.to/rho.html
//  Their "embedding degree" is also higher than every curve listed by SafeCurves https://safecurves.cr.yp.to/transfer.html
//
//  No further consideration is given while choosing the Prime field form or the D coefficient for speed purposes,
//  because all Elliptic Point math is implemented in pure golang using directly big.Int. The speed of implementation is
//  sufficient (a Schnorr signature Creation and Verification takes about 100 ms), and a field point multiplication takes
//  about 35 ms for the TEC_545_551p335_m1874 (test computer is an AMD EPYC 7551P 32-Core Processor which has mediocre Freq.)
//  These running times are deemed for now more than sufficient for single time use.
//
//===============================================================================================
// Twisted-Edwards-Curves are elliptic curves satisfying the equation:
//
//	A * x^2 + y^2 = C^2 * (1 + D * x^2 * y^2)
//
// for some scalars A, C, D over a finite large prime field P.
// We also assume C == 1 because all curves in the generalized form are isomorphic to curves having C == 1.
//
// For details see Bernstein et al, "Twisted Edwards Curves", http://eprint.iacr.org/2008/013.pdf
//===============================================================================================
// Montgomery-Curves are elliptic curves satisfying the equation:
//
//	B * y^2 = x^3 + A * x^2 + x
//
// for some scalars B, A over a finite large prime field P.
// We also assume B == 1 because all curves in the generalized form are isomorphic to curves having B == 1.
//
// For details see Bernstein et al, "Twisted Edwards Curves", http://eprint.iacr.org/2008/013.pdf
//
type FiniteFieldEllipticCurve struct {
    Name string                 // Name of curve

    // Prime Numbers
    P big.Int                   // Prime defining the underlying field
    Q big.Int                   // Generator (Base-Point) Order
    T big.Int                   // Trace of the Curve
    R big.Int      		// Cofactor: R*Q  = P + 1 - T

    // Parameters and Coefficients
    A big.Int                   // x^2 Parameter; Montgomery and Twisted Eduards Curve
    D big.Int 		        // x^2 * y^2 Coefficient; Twisted Eduards Curve only

    //Curve safe scalar size in bits
    //The "private key bit size" is as such that considering first bit 1, and the last n bits 0 (clamped scalar)
    //  the total resulting number in decimal must be lower than the Generator (Base-Point) Order Q
    //  where n is log(2,R) (for cofactor 4 n is 2, for cofactor 8 n is 3)
    //S is computed with the function SafeScalarComputer.
    //  For instance S for Curve E-521 is 515 bits. and the resulting clamped scalar is 518 bits in size.
    //  The resulting 518 bit Scalar is smaller than Q. Therefore there are 2^515 possible private keys.
    //  The bigger the S, the more unique Private Keys exist on that curve.
    S uint64

    //Point coordinates
    FBX, FBY big.Int            // Standard base point for full group
    PBX, PBY big.Int 		// Standard base point for prime-order subgroup

    //Elligators
    Elligator1s big.Int 	// Optional s parameter for Elligator 1
    Elligator2u big.Int 	// Optional u parameter for Elligator 2
}

type PrimePowerTwo struct {
    Power int
    RestString string
    Sign bool
}

func CurveCofactor (P, Q, T big.Int) big.Int {
    var h = new(big.Int)
    h.Add(&P,One).Sub(h,&T).Quo(h,&Q)
    return *h
}

func MakePrime (PrimeNumber PrimePowerTwo) big.Int {
    var (
        Prime = new(big.Int)
        Rest = new(big.Int)
    )
    Rest.SetString(PrimeNumber.RestString,10)

    if PrimeNumber.Sign == true {
        Prime.SetBit(Zero, PrimeNumber.Power, 1).Add(Prime, Rest)
    } else {
        Prime.SetBit(Zero, PrimeNumber.Power, 1).Sub(Prime, Rest)
    }

    return *Prime
}
//
// MONTGOMERY CURVES DEFINITIONS
//
// Curve M1:
//
// Function M383 defines the M-383 curve specified in:
// Aranha et al, "A note on high-security general-purpose elliptic curves",
// http://eprint.iacr.org/2013/647.pdf and https://safecurves.cr.yp.to/index.html
//
//
// Name: 			M383
// Equation is: 		y^2 = x^3 + 2065150 * x^2 + x
// Prime field (P)		2^383 - 187
//	Decimal Form		19701003098197239606139520050071806902539869635232723333974146702122860885748605305707133127442457820403313995153221
//	Hexadecimal Form	0x7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff45
//
// Base-Point Prime Order (Q)	2^380 + 166236275931373516105219794935542153308039234455761613271
//	Decimal Form		2462625387274654950767440006258975862817483704404090416746934574041288984234680883008327183083615266784870011007447
//	Hexadecimal Form	0x10000000000000000000000000000000000000000000000006c79673ac36ba6e7a32576f7b1b249e46bbc225be9071d7
//
// Trace                        -1329890207450988128841758359484337226464313875646092906354
// Cofactor R			8
//
// Curve B Parameter		1
// Curve A Parameter		2065150
//
// Base-Point X			12
//	Hexadecimal Form	0xc
// Base-Point Y			4737623401891753997660546300375902576839617167257703725630389791524463565757299203154901655432096558642117242906494
//	Hexadecimal Form	0x1ec7ed04aaf834af310e304b2da0f328e7c165f0e8988abd3992861290f617aa1f1b2e7d0b6e332e969991b62555e77e
func M383() *FiniteFieldEllipticCurve {
    var (
        p FiniteFieldEllipticCurve
        P PrimePowerTwo
        Q PrimePowerTwo
    )
    p.Name = "M383"

    //P=2^383-187
    P.Power = 383
    P.RestString = "187"
    P.Sign = false
    p.P = MakePrime(P)

    //Q=2^380 + 166236275931373516105219794935542153308039234455761613271
    Q.Power = 380
    Q.RestString = "166236275931373516105219794935542153308039234455761613271"
    Q.Sign = true
    p.Q = MakePrime(Q)

    //Trace and Cofactor
    p.T.SetString("-1329890207450988128841758359484337226464313875646092906354", 10)
    p.R = CurveCofactor(p.P,p.Q,p.T)

    //Safe Scalar Size in bits = 376
    p.S,_ = SafeScalarComputer(&p.P,&p.T,&p.R)

    //A Coefficient
    p.A.SetInt64(2065150)

    //Generator Coordinates
    p.PBX.SetString("12", 10)
    p.PBY.SetString("4737623401891753997660546300375902576839617167257703725630389791524463565757299203154901655432096558642117242906494", 10)

    return &p
}
//
// Curve M2:
//
// Function Curve383187 defines the Curve383187 curve specified in:
// https://safecurves.cr.yp.to/index.html
// ---
//
//
// Name: 			Curve383187
// Equation is: 		y^2 = x^3 + 229969 * x^2 + x
// Prime field (P)		2^383 - 187
//	Decimal Form		19701003098197239606139520050071806902539869635232723333974146702122860885748605305707133127442457820403313995153221
//	Hexadecimal Form	0x7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff45
//
// Base-Point Prime Order (Q)	2^380 + 356080847217269887368687156533236720299699248977882517025
//	Decimal Form		2462625387274654950767440006258975862817483704404090416747124418612574880605944350369924877650606926799392131911201
//	Hexadecimal Form	0x1000000000000000000000000000000000000000000000000e85a85287a1488acd41ae84b2b7030446f72088b00a0e21
//
// Trace                        -2848646777738159098949497252265893762397593991823060136386
// Cofactor R			8
//
// Curve B Parameter		1
// Curve A Parameter		2065150
//
// Base-Point X			5
//	Hexadecimal Form	0x5
// Base-Point Y			4759238150142744228328102229734187233490253962521130945928672202662038422584867624507245060283757321006861735839455
//	Hexadecimal Form	0x1eebe07dc1871896732b12d5504a32370471965c7a11f2c89865f855ab3cbd7c224e3620c31af3370788457dd5ce46df
func Curve383187() *FiniteFieldEllipticCurve {
    var (
        p FiniteFieldEllipticCurve
        P PrimePowerTwo
        Q PrimePowerTwo
    )
    p.Name = "Curve383187"

    //P=2^383-187
    P.Power = 383
    P.RestString = "187"
    P.Sign = false
    p.P = MakePrime(P)

    //Q=2^380 + 356080847217269887368687156533236720299699248977882517025
    Q.Power = 380
    Q.RestString = "356080847217269887368687156533236720299699248977882517025"
    Q.Sign = true
    p.Q = MakePrime(Q)

    //Trace and Cofactor
    p.T.SetString("-2848646777738159098949497252265893762397593991823060136386", 10)
    p.R = CurveCofactor(p.P,p.Q,p.T)

    //Safe Scalar Size in bits = 376
    p.S,_ = SafeScalarComputer(&p.P,&p.T,&p.R)

    //A Coefficient
    p.A.SetInt64(229969)

    //Generator Coordinates
    p.PBX.SetString("5", 10)
    p.PBY.SetString("4759238150142744228328102229734187233490253962521130945928672202662038422584867624507245060283757321006861735839455", 10)
    return &p
}
//
// Curve M3:
//
// Function M511 defines the M-511 curve specified in:
// Aranha et al, "A note on high-security general-purpose elliptic curves",
// http://eprint.iacr.org/2013/647.pdf and https://safecurves.cr.yp.to/index.html
//
//
// Name: 			M511
// Equation is: 		y^2 = x^3 + 530438 * x^2 + x
// Prime field (P)		2^511 - 187
//	Decimal Form		6703903964971298549787012499102923063739682910296196688861780721860882015036773488400937149083451713845015929093243025426876941405973284973216824503041861
//	Hexadecimal Form	0x7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff45
//
// Base-Point Prime Order (Q)	2^508 + 10724754759635747624044531514068121842070756627434833028965540808827675062043
//	Decimal Form		837987995621412318723376562387865382967460363787024586107722590232610251879607410804876779383055508762141059258497448934987052508775626162460930737942299
//	Hexadecimal Form	0x100000000000000000000000000000000000000000000000000000000000000017b5feff30c7f5677ab2aeebd13779a2ac125042a6aa10bfa54c15bab76baf1b
//
// Trace                        -85798038077085980992356252112544974736566053019478664231724326470621400496530
// Cofactor R			8
//
// Curve B Parameter		1
// Curve A Parameter		530438
//
// Base-Point X			5
//	Hexadecimal Form	0x5
// Base-Point Y			2500410645565072423368981149139213252211568685173608590070979264248275228603899706950518127817176591878667784247582124505430745177116625808811349787373477
//	Hexadecimal Form	0x2fbdc0ad8530803d28fdbad354bb488d32399ac1cf8f6e01ee3f96389b90c809422b9429e8a43dbf49308ac4455940abe9f1dbca542093a895e30a64af056fa5
func M511() *FiniteFieldEllipticCurve {
    var (
        p FiniteFieldEllipticCurve
        P PrimePowerTwo
        Q PrimePowerTwo
    )
    p.Name = "M511"

    //P=2^511-187
    P.Power = 511
    P.RestString = "187"
    P.Sign = false
    p.P = MakePrime(P)

    //Q=2^508 + 10724754759635747624044531514068121842070756627434833028965540808827675062043
    Q.Power = 508
    Q.RestString = "10724754759635747624044531514068121842070756627434833028965540808827675062043"
    Q.Sign = true
    p.Q = MakePrime(Q)

    //Trace and Cofactor
    p.T.SetString("-85798038077085980992356252112544974736566053019478664231724326470621400496530", 10)
    p.R = CurveCofactor(p.P,p.Q,p.T)

    //Safe Scalar Size in bits = 504
    p.S,_ = SafeScalarComputer(&p.P,&p.T,&p.R)

    //A Coefficients
    p.A.SetInt64(-530438)

    //Generator Coordinates
    p.PBX.SetString("5", 10)
    p.PBY.SetString("2500410645565072423368981149139213252211568685173608590070979264248275228603899706950518127817176591878667784247582124505430745177116625808811349787373477", 10)
    return &p
}
//
// TWISTED EDWARDS CURVES DEFINITIONS
//
// Curve TEC-Standard-1:
//
// Function E382 defines the E-382 curve specified in:
// Aranha et al, "A note on high-security general-purpose elliptic curves",
// http://eprint.iacr.org/2013/647.pdf , https://safecurves.cr.yp.to/index.html
//
// and more recently in:
//
// "Additional Elliptic Curves for IETF protocols"
// http://tools.ietf.org/html/draft-ladd-safecurves-02
//
// Name: 			E382
// Equation is: 		x^2 + y^2 = 1 - 67254 * x^2 * y^2
// Prime field (P)		2^382 - 105
//	Decimal Form		9850501549098619803069760025035903451269934817616361666987073351061430442874302652853566563721228910201656997576599
//	Hexadecimal Form	0x3fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff97
//
// Base-Point Prime Order (Q)	2^380 - 1030303207694556153926491950732314247062623204330168346855
//	Decimal Form		2462625387274654950767440006258975862817483704404090416745738034557663054564649171262659326683244604346084081047321
//	Hexadecimal Form	0xfffffffffffffffffffffffffffffffffffffffffffffffd5fb21f21e95eee17c5e69281b102d2773e27e13fd3c9719
//
// Trace                        4121212830778224615705967802929256988250492817320673387316
// Cofactor R			4
//
// Curve A Parameter		1
// Curve D Parameter		-67254
//
// Base-Point X			3914921414754292646847594472454013487047137431784830634731377862923477302047857640522480241298429278603678181725699
//	Hexadecimal Form	0x196f8dd0eab20391e5f05be96e8d20ae68f840032b0b64352923bab85364841193517dbce8105398ebc0cc9470f79603
// Base-Point Y			17
//	Hexadecimal Form	0x11
func E382() *FiniteFieldEllipticCurve {
    var (
    	p FiniteFieldEllipticCurve
        P PrimePowerTwo
        Q PrimePowerTwo
    )
    p.Name = "E382"

    //P=2^382 - 105
    P.Power = 382
    P.RestString = "105"
    P.Sign = false
    p.P = MakePrime(P)

    //Q=2^380 - 1030303207694556153926491950732314247062623204330168346855
    Q.Power = 380
    Q.RestString = "1030303207694556153926491950732314247062623204330168346855"
    Q.Sign = false
    p.Q = MakePrime(Q)

    //Trace and Cofactor
    p.T.SetString("4121212830778224615705967802929256988250492817320673387316", 10)
    p.R = CurveCofactor(p.P,p.Q,p.T)

    //Safe Scalar Size in bits = 376
    p.S,_ = SafeScalarComputer(&p.P,&p.T,&p.R)

    //A and D Coefficients
    p.A.SetInt64(1)
    p.D.SetInt64(-67254)

    //Generator Coordinates
    p.PBX.SetString("3914921414754292646847594472454013487047137431784830634731377862923477302047857640522480241298429278603678181725699", 10)
    p.PBY.SetString("17", 10)
    //p.PBX.SetString("17", 10)
    //p.PBY.SetString("3914921414754292646847594472454013487047137431784830634731377862923477302047857640522480241298429278603678181725699", 10)
    return &p
}
//
// Curve TEC-Standard-2:
//
// Function Curve41417 defines the Curve41417 curve, as specified in:
// Bernstein et al, "Curve41417: Karatsuba revisited",
// http://eprint.iacr.org/2014/526.pdf and https://safecurves.cr.yp.to/index.html
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
// Trace                        266913126910041140166481421552787081431877817603289668716758056
// Cofactor R			8
//
// Curve A Parameter		1
// Curve D Parameter		3617
//
// Base-Point X			17319886477121189177719202498822615443556957307604340815256226171904769976866975908866528699294134494857887698432266169206165
//	Hexadecimal Form	0x1a334905141443300218c0631c326e5fcd46369f44c03ec7f57ff35498a4ab4d6d6ba111301a73faa8537c64c4fd3812f3cbc595
// Base-Point Y			34
//	Hexadecimal Form	0x22
func Curve41417() *FiniteFieldEllipticCurve {
    var (
        p FiniteFieldEllipticCurve
        P PrimePowerTwo
        Q PrimePowerTwo
    )
    p.Name = "Curve41417"

    //P=2^414 - 17
    P.Power = 414
    P.RestString = "17"
    P.Sign = false
    p.P = MakePrime(P)

    //Q=2^411 - 33364140863755142520810177694098385178984727200411208589594759
    Q.Power = 411
    Q.RestString = "33364140863755142520810177694098385178984727200411208589594759"
    Q.Sign = false
    p.Q = MakePrime(Q)

    //Trace and Cofactor
    p.T.SetString("266913126910041140166481421552787081431877817603289668716758056", 10)
    p.R = CurveCofactor(p.P,p.Q,p.T)

    //Safe Scalar Size in bits = 406
    p.S,_ = SafeScalarComputer(&p.P,&p.T,&p.R)

    //A and D Coefficients
    p.A.SetInt64(1)
    p.D.SetInt64(3617)

    //Generator Coordinates
    p.PBX.SetString("17319886477121189177719202498822615443556957307604340815256226171904769976866975908866528699294134494857887698432266169206165", 10)
    p.PBY.SetString("34", 10)
    return &p
}
//
// Curve TEC-Standard-3:
//
// Function Goldilocks defines the Ed448-Goldilocks curve, as specified in:
// Bernstein et al, "Ed448-Goldilocks, a new elliptic curve",
// https://eprint.iacr.org/2015/625.pdf and https://safecurves.cr.yp.to/index.html
//
// Name: 			Goldilocks
// Equation is: 		x^2 + y^2 = 1 - 39081 * x^2 * y^2
// Prime field (P)		2^448 - 2^224 - 1
//	Decimal Form		726838724295606890549323807888004534353641360687318060281490199180612328166730772686396383698676545930088884461843637361053498018365439
//	Hexadecimal Form	0xfffffffffffffffffffffffffffffffffffffffffffffffffffffffeffffffffffffffffffffffffffffffffffffffffffffffffffffffff
//
// Base-Point Prime Order (Q)	2^446 - 13818066809895115352007386748515426880336692474882178609894547503885
//	Decimal Form		181709681073901722637330951972001133588410340171829515070372549795146003961539585716195755291692375963310293709091662304773755859649779
//	Hexadecimal Form	0x3fffffffffffffffffffffffffffffffffffffffffffffffffffffff7cca23e9c44edb49aed63690216cc2728dc58f552378c292ab5844f3
//
// Trace                        28312320572429821613362531907042076847709625476988141958474579766324
// Cofactor R			4
//
// Curve A Parameter		1
// Curve D Parameter		-39081
//
// Base-Point X			117812161263436946737282484343310064665180535357016373416879082147939404277809514858788439644911793978499419995990477371552926308078495
//	Hexadecimal Form	0x297ea0ea2692ff1b4faff46098453a6a26adf733245f065c3c59d0709cecfa96147eaaf3932d94c63d96c170033f4ba0c7f0de840aed939f
// Base-Point Y			19
//	Hexadecimal Form	0x13
func Goldilocks() *FiniteFieldEllipticCurve {
    var (
        p FiniteFieldEllipticCurve
        P PrimePowerTwo
        R PrimePowerTwo
        Q PrimePowerTwo
    )
    p.Name = "Goldilocks"

    //Rp=2^224 + 1
    R.Power = 224
    R.RestString = "1"
    R.Sign = true
    Rp := MakePrime(R)
    //P=2^448-Rp=2^448-(2^224+1)=2^448-2^224-1
    P.Power = 448
    P.RestString = Rp.Text(10)
    P.Sign = false
    p.P = MakePrime(P)

    //Q=2^446 - 13818066809895115352007386748515426880336692474882178609894547503885
    Q.Power = 446
    Q.RestString = "13818066809895115352007386748515426880336692474882178609894547503885"
    Q.Sign = false
    p.Q = MakePrime(Q)

    //Trace and Cofactor
    p.T.SetString("28312320572429821613362531907042076847709625476988141958474579766324", 10)
    p.R = CurveCofactor(p.P,p.Q,p.T)

    //Safe Scalar Size in bits = 442
    p.S,_ = SafeScalarComputer(&p.P,&p.T,&p.R)

    //A and D Coefficients
    p.A.SetInt64(1)
    p.D.SetInt64(-39081)

    //Generator Coordinates
    p.PBX.SetString("117812161263436946737282484343310064665180535357016373416879082147939404277809514858788439644911793978499419995990477371552926308078495", 10)
    p.PBY.SetString("19", 10)
    return &p
}
//
// Curve TEC-Standard-4:
//
// Function E521 defines the E-521 curve specified in:
// Aranha et al, "A note on high-security general-purpose elliptic curves",
// http://eprint.iacr.org/2013/647.pdf , https://safecurves.cr.yp.to/index.html
//
// and more recently included in:
// "Additional Elliptic Curves for IETF protocols"
// http://tools.ietf.org/html/draft-ladd-safecurves-02
//
// Name: 			E521
// Equation is: 		x^2 + y^2 = 1 - 376014 * x^2 * y^2
// Prime field (P)		2^521 - 1
//	Decimal Form		6864797660130609714981900799081393217269435300143305409394463459185543183397656052122559640661454554977296311391480858037121987999716643812574028291115057151
//	Hexadecimal Form	0x1ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff
//
// Base-Point Prime Order (Q)	2^519 - 337554763258501705789107630418782636071904961214051226618635150085779108655765
//	Decimal Form		1716199415032652428745475199770348304317358825035826352348615864796385795849413675475876651663657849636693659065234142604319282948702542317993421293670108523
//	Hexadecimal Form	0x7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffd15b6c64746fc85f736b8af5e7ec53f04fbd8c4569a8f1f4540ea2435f5180d6b
//
// Trace                        1350219053034006823156430521675130544287619844856204906474540600343116434623060
// Cofactor R			4
//
// Curve A Parameter	1
// Curve D Parameter	-376014
//
// Base-Point X			1571054894184995387535939749894317568645297350402905821437625181152304994381188529632591196067604100772673927915114267193389905003276673749012051148356041324
//	Hexadecimal Form	0x752cb45c48648b189df90cb2296b2878a3bfd9f42fc6c818ec8bf3c9c0c6203913f6ecc5ccc72434b1ae949d568fc99c6059d0fb13364838aa302a940a2f19ba6c
// Base-Point Y			12
//	Hexadecimal Form	0xc
func E521() *FiniteFieldEllipticCurve {
    var (
        p FiniteFieldEllipticCurve
        P PrimePowerTwo
        Q PrimePowerTwo
    )
    p.Name = "E521"

    //P=2^521 - 1
    P.Power = 521
    P.RestString = "1"
    P.Sign = false
    p.P = MakePrime(P)

    //Q=2^519 - 337554763258501705789107630418782636071904961214051226618635150085779108655765
    Q.Power = 519
    Q.RestString = "337554763258501705789107630418782636071904961214051226618635150085779108655765"
    Q.Sign = false
    p.Q = MakePrime(Q)

    //Trace and Cofactor
    p.T.SetString("1350219053034006823156430521675130544287619844856204906474540600343116434623060", 10)
    p.R = CurveCofactor(p.P,p.Q,p.T)

    //Safe Scalar Size in bits = 515
    p.S,_ = SafeScalarComputer(&p.P,&p.T,&p.R)

    //A and D Coefficients
    p.A.SetInt64(1)
    p.D.SetInt64(-376014)

    //Generator Coordinates
    p.PBX.SetString("1571054894184995387535939749894317568645297350402905821437625181152304994381188529632591196067604100772673927915114267193389905003276673749012051148356041324", 10)
    p.PBY.SetString("12", 10)
    return &p
}
//======================================================================================================================
// Novel never before used Twisted Edwards Curves
//
// Curve TEC-Novel-0000:
//
// Function LittleOne defines an elliptic curve of Cofactor 4 over a small prime field of 997,
// used for testing and observing purposes.
//
// Name: 			LittleOne
// Equation is: 		x^2 + y^2 = 1 + 446 * x^2 * y^2
// Prime field (P)		2^10 - 27
//	Decimal Form		997
//	Hexadecimal Form	??
//
// Base-Point Prime Order (Q)	2^8+1
//	Decimal Form		257
//	Hexadecimal Form	??
// Trace (T)                    -30
// Cofactor R			4
// Safe Scalar Size S           5
//
// Curve A Parameter		1
// Curve D Parameter		446
//
// Base-Point X                 24
//	Hexadecimal Form        ??
// Base-Point Y                 440
//	Hexadecimal Form        ??
func LittleOne() *FiniteFieldEllipticCurve {
    var (
        p FiniteFieldEllipticCurve
        P PrimePowerTwo
        Q PrimePowerTwo
    )
    p.Name = "LittleOne"

    //P=2^10 - 27
    P.Power = 10
    P.RestString = "27"
    P.Sign = false
    p.P = MakePrime(P)

    //Q=2^8+1
    Q.Power = 8
    Q.RestString = "1"
    Q.Sign = true
    p.Q = MakePrime(Q)

    //Trace and Cofactor
    p.T.SetString("-30", 10)
    p.R = CurveCofactor(p.P,p.Q,p.T)

    //Safe Scalar Size in bits = 5
    p.S,_ = SafeScalarComputer(&p.P,&p.T,&p.R)

    //A and D Coefficients
    p.A.SetInt64(1)
    p.D.SetInt64(446)

    //Generator Coordinates
    p.PBX.SetString("24", 10)
    p.PBY.SetString("440", 10)
    return &p
}
//======================================================================================================================
//
// Curve TEC-Novel-0001:
//
// Function TEC_S545_Pr551p335_m1874 defines the first Curve discovered with an S-value higher than the highest S-Value
// Public TEC Curve, the E521. It's D coefficient was discovered fairly easy after computing approximately 15000
// Coefficients (Computing Coefficients refers to computing the Trace for a given D coefficient and prime field) with
// Sage on 2 CPUs (Ryzen 5950 testing negative D coefficients, and an Epyc 7551P testing positive Coefficients), over
// the course of approximately 1 and a half day.
//
// Name: 			TEC_S545_Pr551p335_m1874
// Equation is: 		x^2 + y^2 = 1 - 1874 * x^2 * y^2
// Prime field (P)		2^551 + 335
//	Decimal Form		7371020360979572953596786290992712677572111758625860211672277930167234692172165726730716260112614780354430419981960634569864423105321860610471551272329484460252725583
//	Hexadecimal Form	??
//
// Base-Point Prime Order (Q)	2^549-32999719876419924862440765771944715506860861139489669592317112655962959048275399831
//	Decimal Form		1842755090244893238399196572748178169393027939656465052918069482541808673043041431649679345151733770226166839223545443135605244636840795560300775162119412066787781481
//	Hexadecimal Form	??
// Trace (T)                    131998879505679699449763063087778862027443444557958678369268450623851836193101599660
// Cofactor R			4
// Safe Scalar Size S           545
//
// Curve A Parameter		1
// Curve D Parameter		-1874
//
// Base-Point X                 5
//	Hexadecimal Form        ??
// Base-Point Y                 4518488039903337342061416616304793185577751419009710712882273229786958102867981468569632796010757506367702155510163192445896310025029799220155797291359909742186717128
//	Hexadecimal Form        ??
func TEC_S545_Pr551p335_m1874() *FiniteFieldEllipticCurve {
    var (
        p FiniteFieldEllipticCurve
        P PrimePowerTwo
        Q PrimePowerTwo
    )
    p.Name = "TEC_545_551p335_m1874"

    //P=2^551 + 335
    P.Power = 551
    P.RestString = "335"
    P.Sign = true
    p.P = MakePrime(P)

    //Q=2^549 - 32999719876419924862440765771944715506860861139489669592317112655962959048275399831
    Q.Power = 549
    Q.RestString = "32999719876419924862440765771944715506860861139489669592317112655962959048275399831"
    Q.Sign = false
    p.Q = MakePrime(Q)

    //Trace and Cofactor
    p.T.SetString("131998879505679699449763063087778862027443444557958678369268450623851836193101599660", 10)
    p.R = CurveCofactor(p.P,p.Q,p.T)

    //Safe Scalar Size in bits = 545
    p.S,_ = SafeScalarComputer(&p.P,&p.T,&p.R)

    //A and D Coefficients
    p.A.SetInt64(1)
    p.D.SetInt64(-1874)

    //Generator Coordinates
    p.PBX.SetString("5", 10)
    p.PBY.SetString("4518488039903337342061416616304793185577751419009710712882273229786958102867981468569632796010757506367702155510163192445896310025029799220155797291359909742186717128", 10)
    return &p
}
//
// Curve TEC-Novel-0002:
//
// Function TEC_S1021_Pr1029p639_p403 defines the first Curve discovered while searching for a Curve with an S value of
// 1024 on the same systems used for finding the D coefficient for the TEC_S545_Pr551p335_m1874. The Coefficient appeared
// after approximately 1 day of Testing. At the point of its discovery, it became the biggest S-Value curve defined here.
//
// Name: 			TEC_S1021_Pr1029p639_p403
// Equation is: 		x^2 + y^2 = 1 + 403 * x^2 * y^2
// Prime field (P)		2^1029 + 639
//	Decimal Form		5752618031559410904733776610524879147577526332615381032749762597047445625776030820246671274317041152675843644155884587445081272602061331919771117780463171980088572589595695528841671027239875011822498654466720184602820821834958812207165219537306471589227216341906761543678311870031350921754731402547975172391551
//	Hexadecimal Form	??
//
// Base-Point Prime Order (Q)	2^1026-18132002252546586200370882567476047696677474403358385255776356556256228841222610345936048352899569081102195019149436007172657992929987550174537497129476947
//	Decimal Form		719077253944926363091722076315609893447190791576922629093720324630930703222003852530833909289630144084480455519485573430635159075257666489971389722557896479379069321152875740734326310928936679800337928449954767298996046473141010303285306506114956049084320940543326043523781811095925935231791250780999767071917
//	Hexadecimal Form	??
// Trace (T)                    145056018020372689602967060539808381573419795226867082046210852450049830729780882767488386823196552648817560153195488057381263943439900401396299977035816216
// Cofactor R			8
// Safe Scalar Size S           1021
//
// Curve A Parameter		1
// Curve D Parameter		403
//
// Base-Point X                 39
//	Hexadecimal Form        ??
// Base-Point Y                 5576533075937720038098979394411727019827964480087178017371868264528072186969486283264858007407820973699691992666589452515693225782545746655700587813300615359050446313584121584063198554946923237708227614531033897522367303509581575612135794007129867293613969915343060843254041982275099993760177593077210810921912
//	Hexadecimal Form        ??
func TEC_S1021_Pr1029p639_p403() *FiniteFieldEllipticCurve {
    var (
        p FiniteFieldEllipticCurve
        P PrimePowerTwo
        Q PrimePowerTwo
    )
    p.Name = "TEC_S1021_Pr1029p639_p403"

    //P=2^1029 + 639
    P.Power = 1029
    P.RestString = "639"
    P.Sign = true
    p.P = MakePrime(P)

    //Q=2^1026-18132002252546586200370882567476047696677474403358385255776356556256228841222610345936048352899569081102195019149436007172657992929987550174537497129476947
    Q.Power = 1026
    Q.RestString = "18132002252546586200370882567476047696677474403358385255776356556256228841222610345936048352899569081102195019149436007172657992929987550174537497129476947"
    Q.Sign = false
    p.Q = MakePrime(Q)

    //Trace and Cofactor
    p.T.SetString("145056018020372689602967060539808381573419795226867082046210852450049830729780882767488386823196552648817560153195488057381263943439900401396299977035816216", 10)
    p.R = CurveCofactor(p.P,p.Q,p.T)

    //Safe Scalar Size in bits = 1021
    p.S,_ = SafeScalarComputer(&p.P,&p.T,&p.R)

    //A and D Coefficients
    p.A.SetInt64(1)
    p.D.SetInt64(403)

    //Generator Coordinates
    p.PBX.SetString("39", 10)
    p.PBY.SetString("5576533075937720038098979394411727019827964480087178017371868264528072186969486283264858007407820973699691992666589452515693225782545746655700587813300615359050446313584121584063198554946923237708227614531033897522367303509581575612135794007129867293613969915343060843254041982275099993760177593077210810921912", 10)
    return &p
}
// Curve TEC-Novel-0003:
//
// Function TEC_S1023_Pr1029p639_m200 defines the second Curve discovered while searching for a Curve with an S value of
// 1024 on the same systems used for finding the D coefficient for the TEC_S545_Pr551p335_m1874. The Coefficient appeared
// after approximately 2 days of Testing. At the point of its discovery, it became the biggest S-Value curve defined here.
//
// Name: 			TEC_S1023_Pr1029p639_m200
// Equation is: 		x^2 + y^2 = 1 - 200 * x^2 * y^2
// Prime field (P)		2^1029 + 639
//	Decimal Form		5752618031559410904733776610524879147577526332615381032749762597047445625776030820246671274317041152675843644155884587445081272602061331919771117780463171980088572589595695528841671027239875011822498654466720184602820821834958812207165219537306471589227216341906761543678311870031350921754731402547975172391551
//	Hexadecimal Form	??
//
// Base-Point Prime Order (Q)	2^1027-13048810356164098687722578038659254541745638134607534327178785488911851451225729494174990871326244818073066727207431577242595505052831067258628249533735995
//	Decimal Form            1438154507889852726183444152631219786894381583153845258187440649261861406444007705061667818579260288168960911038971146861270318150515332979942779445115792981973332791234825194487839718150714211209986529009145718971919716546888251826061810709335746571061986012409963178488000724912332677607615592008744259361733
//	Hexadecimal Form	??
// Trace (T)                    52195241424656394750890312154637018166982552538430137308715141955647405804902917976699963485304979272292266908829726308970382020211324269034512998134944620
// Cofactor R			4
// Safe Scalar Size S           1023
//
// Curve A Parameter		1
// Curve D Parameter		-200
//
// Base-Point X                 18
//	Hexadecimal Form        ??
// Base-Point Y                 5006392512810367543241026017186205828475671321699765938632799901604288413670061260105487647663536568022230479638139010351335665470173712718298837530633945899923869302110921390691280063917337749102198086546109683731403172016859789550276802795383170944526602213977392860793115308281053135496569817870067300616902
//	Hexadecimal Form        ??
func TEC_S1023_Pr1029p639_m200() *FiniteFieldEllipticCurve {
    var (
        p FiniteFieldEllipticCurve
        P PrimePowerTwo
        Q PrimePowerTwo
    )
    p.Name = "TEC_S1023_Pr1029p639_m200"

    //P=2^1029 + 639
    P.Power = 1029
    P.RestString = "639"
    P.Sign = true
    p.P = MakePrime(P)

    //Q=2^1027-13048810356164098687722578038659254541745638134607534327178785488911851451225729494174990871326244818073066727207431577242595505052831067258628249533735995
    Q.Power = 1027
    Q.RestString = "13048810356164098687722578038659254541745638134607534327178785488911851451225729494174990871326244818073066727207431577242595505052831067258628249533735995"
    Q.Sign = false
    p.Q = MakePrime(Q)

    //Trace and Cofactor
    p.T.SetString("52195241424656394750890312154637018166982552538430137308715141955647405804902917976699963485304979272292266908829726308970382020211324269034512998134944620", 10)
    p.R = CurveCofactor(p.P,p.Q,p.T)

    //Safe Scalar Size in bits = 1023
    p.S,_ = SafeScalarComputer(&p.P,&p.T,&p.R)

    //A and D Coefficients
    p.A.SetInt64(1)
    p.D.SetInt64(-200)

    //Generator Coordinates
    p.PBX.SetString("18", 10)
    p.PBY.SetString("5006392512810367543241026017186205828475671321699765938632799901604288413670061260105487647663536568022230479638139010351335665470173712718298837530633945899923869302110921390691280063917337749102198086546109683731403172016859789550276802795383170944526602213977392860793115308281053135496569817870067300616902", 10)
    return &p
}
//
// Curve TEC-Novel-0004:
//
// Function TEC_S1024_Pr1029p639_m2729 defines the first Cofactor 4 TEC Curve discovered  with an S value of
// 1024 on the same systems used for finding the D coefficient for the TEC_S545_Pr551p335_m1874. The Coefficient appeared
// after approximately 3 days of Testing. At the point of its discovery, it became the biggest S-Value curve defined here.
//
// Name: 			TEC_S1024_Pr1029p639_m729
// Equation is: 		x^2 + y^2 = 1 - 729 * x^2 * y^2
// Prime field (P)		2^1029 + 639
//	Decimal Form		5752618031559410904733776610524879147577526332615381032749762597047445625776030820246671274317041152675843644155884587445081272602061331919771117780463171980088572589595695528841671027239875011822498654466720184602820821834958812207165219537306471589227216341906761543678311870031350921754731402547975172391551
//	Hexadecimal Form	??
//
// Base-Point Prime Order (Q)	2^1027+9418258840691661048958693280848051387209299408194089012775090979625514940291099061879232380228215863338991692577868713164205283324554730781862605682126581
//	Decimal Form            1438154507889852726183444152631219786894381583153845258187440649261861406444007705061667818579260288168960911038971146861270318150515332979942779445115793004440401988090584931169111037658020140164924071810769058925796185084254643342890366763558998125522667424468382963788291131713121054993413632499599475224309
//	Hexadecimal Form	??
// Trace (T)                    -37673035362766644195834773123392205548837197632776356051100363918502059761164396247516929520912863453355966770311474852656821133298218923127450422728505684
// Cofactor R			4
// Safe Scalar Size S           1024
//
// Curve A Parameter		1
// Curve D Parameter		-729
//
// Base-Point X                 18
//	Hexadecimal Form        ??
// Base-Point Y                 215278369951571488896917596155404324026002302190181825431681175816997859909367143653000579666656344456792257669605762582468610930751115704301503268336066379058325768607564533090162357247378501333085803173440477981455490888754538866823680129180124913908161391361773138634347515375569488540295649449731695734303
//	Hexadecimal Form        ??
func TEC_S1024_Pr1029p639_m729() *FiniteFieldEllipticCurve {
    var (
        p FiniteFieldEllipticCurve
        P PrimePowerTwo
        Q PrimePowerTwo
    )
    p.Name = "TEC_S1023_Pr1029p639_m200"

    //P=2^1029 + 639
    P.Power = 1029
    P.RestString = "639"
    P.Sign = true
    p.P = MakePrime(P)

    //Q=2^1027+9418258840691661048958693280848051387209299408194089012775090979625514940291099061879232380228215863338991692577868713164205283324554730781862605682126581
    Q.Power = 1027
    Q.RestString = "9418258840691661048958693280848051387209299408194089012775090979625514940291099061879232380228215863338991692577868713164205283324554730781862605682126581"
    Q.Sign = true
    p.Q = MakePrime(Q)

    //Trace and Cofactor
    p.T.SetString("-37673035362766644195834773123392205548837197632776356051100363918502059761164396247516929520912863453355966770311474852656821133298218923127450422728505684", 10)
    p.R = CurveCofactor(p.P,p.Q,p.T)

    //Safe Scalar Size in bits = 1024
    p.S,_ = SafeScalarComputer(&p.P,&p.T,&p.R)

    //A and D Coefficients
    p.A.SetInt64(1)
    p.D.SetInt64(-729)

    //Generator Coordinates
    p.PBX.SetString("18", 10)
    p.PBY.SetString("215278369951571488896917596155404324026002302190181825431681175816997859909367143653000579666656344456792257669605762582468610930751115704301503268336066379058325768607564533090162357247378501333085803173440477981455490888754538866823680129180124913908161391361773138634347515375569488540295649449731695734303", 10)
    return &p
}


//
// Curve TEC-Novel-0005:
//
// Function TEC_S1600_Pr1605p2315_m26 defines the first Cofactor 4 TEC Curve discovered  with an S value of
// 1600 on the same systems used for finding the D coefficient for the TEC_S545_Pr551p335_m1874. The Coefficient appeared
// after approximately 2 days of Testing, which is incredible, as a single Trace calculation takes between 7 and 9 hours on a Ryzen 5950x
// At the point of its discovery, it became the biggest S-Value curve defined here.
//
// Name: 			TEC_S1600_Pr1605p2315_m26
// Equation is: 		x^2 + y^2 = 1 - 26 * x^2 * y^2
// Prime field (P)		2^1605 + 2315
//	Decimal Form		1422797327267009427840538050096555658106215504388410858221383139000117858232612871647513042437567021359550512676513580279843425806569087338124320534451609923826002722515853311255382582347042785249149784360802682472492979074093381556885921145244739647732065354535899210048956069171919597180088946524547104950094769024061289963739599006148764550591932282504503081289479047974485591701434070586194072480318729612468538986301952269327160804624381735897223889905247070633477140043668654347
//	Hexadecimal Form	??
//
// Base-Point Prime Order (Q)	2^1603+1258387060301909514024042379046449850251725029634697115619073843890705481440046740552204199635883885272944914904655483501916023678206167596650367826811846862157534952990004386839463386963494516862067933899764941962204635259228497801901380413
//	Decimal Form            355699331816752356960134512524138914526553876097102714555345784750029464558153217911878260609391755339887628169128395069960856451642271834531080133612902480956500680628963327813845645586760696312287446090200670618123244768523345389221480286312443298993318248147998844891285467143231624324656933746755850081414397737455369231487103951173075022920927985530781253824285785671827565522008885473360364982237217356107139133414951454295284718018163367874070914438516402917597782812818543421
//	Hexadecimal Form	??
// Trace (T)                    -5033548241207638056096169516185799401006900118538788462476295375562821925760186962208816798543535541091779659618621934007664094712824670386601471307247387448630139811960017547357853547853978067448271735599059767848818541036913991207605519336
// Cofactor R			4
// Safe Scalar Size S           1600
//
// Curve A Parameter		1
// Curve D Parameter		-26
//
// Base-Point X                 2
//	Hexadecimal Form        ??
// Base-Point Y                 479577721234741891316129314062096440203224800598561362604776518993348406897758651324205216647014453759416735508511915279509434960064559686580741767201752370055871770203009254182472722342456597752506165983884867351649283353392919401537107130232654743719219329990067668637876645065665284755295099198801899803461121192253205447281506198423683290960014859350933836516450524873032454015597501532988405894858561193893921904896724509904622632232182531698393484411082218273681226753590907472
//	Hexadecimal Form        ??
func TEC_S1600_Pr1605p2315_m26() *FiniteFieldEllipticCurve {
    var (
        p FiniteFieldEllipticCurve
        P PrimePowerTwo
        Q PrimePowerTwo
    )
    p.Name = "TEC_S1600_Pr1605p2315_m26"

    //P=2^1605 + 2315
    P.Power = 1605
    P.RestString = "2315"
    P.Sign = true
    p.P = MakePrime(P)

    //2^1603+1258387060301909514024042379046449850251725029634697115619073843890705481440046740552204199635883885272944914904655483501916023678206167596650367826811846862157534952990004386839463386963494516862067933899764941962204635259228497801901380413
    Q.Power = 1603
    Q.RestString = "1258387060301909514024042379046449850251725029634697115619073843890705481440046740552204199635883885272944914904655483501916023678206167596650367826811846862157534952990004386839463386963494516862067933899764941962204635259228497801901380413"
    Q.Sign = true
    p.Q = MakePrime(Q)

    //Trace and Cofactor
    p.T.SetString("-5033548241207638056096169516185799401006900118538788462476295375562821925760186962208816798543535541091779659618621934007664094712824670386601471307247387448630139811960017547357853547853978067448271735599059767848818541036913991207605519336", 10)
    p.R = CurveCofactor(p.P,p.Q,p.T)

    //Safe Scalar Size in bits = 1600
    p.S,_ = SafeScalarComputer(&p.P,&p.T,&p.R)

    //A and D Coefficients
    p.A.SetInt64(1)
    p.D.SetInt64(-26)

    //Generator Coordinates
    p.PBX.SetString("2", 10)
    p.PBY.SetString("479577721234741891316129314062096440203224800598561362604776518993348406897758651324205216647014453759416735508511915279509434960064559686580741767201752370055871770203009254182472722342456597752506165983884867351649283353392919401537107130232654743719219329990067668637876645065665284755295099198801899803461121192253205447281506198423683290960014859350933836516450524873032454015597501532988405894858561193893921904896724509904622632232182531698393484411082218273681226753590907472", 10)
    return &p
}