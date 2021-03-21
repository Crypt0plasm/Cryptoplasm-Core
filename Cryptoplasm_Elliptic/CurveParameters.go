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
//          M383
//          Curve383187
//          M511
//      4 Standard known Twisted Edwards Curves,
//          E382
//          Curve41417
//          Ed448-Goldilocks
//          E521
//      as defined by Safe Curves https://safecurves.cr.yp.to/index.html, and
//      1 novel never before used Twisted Edwards Curves,
//          TEC_1617_m1874
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
//  Curve TEC_1617_m1875 has an S value of 545, which can be represented by an intercalated grid of 16*16 and 17*17
//  diagonal squares, hence the name 1617, whereas m1875 represents the D coefficient (minus 1875). 16*16+17*17=545.
//
//  The Novel TEC curves creation aim is based on getting a higher S, while keeping Rigidity and Safety Criteria as
//  defined by https://safecurves.cr.yp.to/rigid.html.
//      The prime field is chosen to be the closest to a power of 2 while being congruent to 3 mod 4
//          (optimal condition for cofactor 4 TEC curves with A coefficient 1)
//      Cofactor 4 Curves are searched for, this is the smallest possible Cofactor for TEC curves, to maximize Q size.
//      D is chosen to be the smallest absolute value that produces the intended S size,
//      while the largest Subgroup Order Q is of course prime.
//      The Generator is then chosen as the Point with the smallest x affine coordinate that is on the Subgroup Q.
//
//  All Novel TEC Curves have the "Complex-Multiplication Field Discriminant" larger than 2^100 as required by https://safecurves.cr.yp.to/disc.html
//  Their "rho cost" is higher than every curve listed by SafeCurves https://safecurves.cr.yp.to/rho.html
//  Their "embedding degree" is also higher than every curve listed by SafeCurves https://safecurves.cr.yp.to/transfer.html
//
//  Further considerations aren't given to the form of the prime field or to the form of the D coefficient,
//  (which are normally considered to speed up math), field point multiplication is implemented in pure golang
//  using directly big.Int without creating limbs, because as these are outside the scope of the Curves intended use.
//
//  No further consideration is given while choosing the Prime field form or the D coefficient for speed purposes,
//  because all Elliptic Point math is implemented in pure golang using directly big.Int. The speed of implementation is
//  sufficient (a Schnorr signature Creation and Verification takes about 100 ms), and a field point multiplication takes
//  about 35 ms for the TEC_1617_m1874 (test computer is an AMD EPYC 7551P 32-Core Processor which has mediocre Freq.)
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
// DefineM383 defines the M-383 curve specified in:
// Aranha et al, "A note on high-security general-purpose elliptic curves",
// http://eprint.iacr.org/2013/647.pdf
//
//
// Name: 			M-383
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
func DefineM383() *FiniteFieldEllipticCurve {
    var (
        p FiniteFieldEllipticCurve
        P PrimePowerTwo
        Q PrimePowerTwo
    )
    p.Name = "M-383"

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
// Define383187 defines the Curve383187 curve specified in:
// ---
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
func Define383187() *FiniteFieldEllipticCurve {
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
// DefineM511 defines the M-511 curve specified in:
// Aranha et al, "A note on high-security general-purpose elliptic curves",
// http://eprint.iacr.org/2013/647.pdf
//
//
// Name: 			M-511
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
func DefineM511() *FiniteFieldEllipticCurve {
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
// DefineE382 defines the E-382 curve specified in:
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
func DefineE382() *FiniteFieldEllipticCurve {
    var (
    	p FiniteFieldEllipticCurve
        P PrimePowerTwo
        Q PrimePowerTwo
    )
    p.Name = "E-382"

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
    return &p
}
//
// Curve TEC-Standard-2:
//
// Define41417 defines the Curve41417 curve, as specified in:
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
func Define41417() *FiniteFieldEllipticCurve {
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
// DefineGoldilocks defines the Ed448-Goldilocks curve, as specified in:
// Bernstein et al, "Ed448-Goldilocks, a new elliptic curve",
// https://eprint.iacr.org/2015/625.pdf
//
// Name: 			Ed448-Goldilocks
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
func DefineGoldilocks() *FiniteFieldEllipticCurve {
    var (
        p FiniteFieldEllipticCurve
        P PrimePowerTwo
        R PrimePowerTwo
        Q PrimePowerTwo
    )
    p.Name = "Ed448-Goldilocks"

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
// DefineE521 defines the E-521 curve specified in:
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
func DefineE521() *FiniteFieldEllipticCurve {
    var (
        p FiniteFieldEllipticCurve
        P PrimePowerTwo
        Q PrimePowerTwo
    )
    p.Name = "E-521"

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
//
// Novel never before used Twisted Edwards Curves
//
// Curve TEC-Novel-0001:
//
// Name: 			TEC_1617_m1874
// Equation is: 		x^2 + y^2 = 1 - 1874 * x^2 * y^2
// Prime field (P)		2^551 + 335
//	Decimal Form		7371020360979572953596786290992712677572111758625860211672277930167234692172165726730716260112614780354430419981960634569864423105321860610471551272329484460252725583
//	Hexadecimal Form	0x209
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
// Base-Point X                 2
//	Hexadecimal Form
// Base-Point Y                 5501022230409387651160486994968341429366041990178351371092942189673622147332545191574204359497248532356106710117122683727294366951792383073058359410033494487297956284
//	Hexadecimal Form
func DefineTec1617M1874() *FiniteFieldEllipticCurve {
    var (
        p FiniteFieldEllipticCurve
        P PrimePowerTwo
        Q PrimePowerTwo
    )
    p.Name = "TEC_1617_m1874"

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
    p.PBX.SetString("2", 10)
    p.PBY.SetString("5372662689287747375095357785063354697686466712559330249539593073273672823349366756432292166344067157200099005518461427402418611188579753765899708983934704551575186856", 10)
    return &p
}