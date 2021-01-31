package Firefly_Precision

import (
	"fmt"
	"github.com/pkg/errors"
	"math"
	"math/big"
)
//=================================================================================================
//=================================================================================================
// Cryptoplasm_Firefly_Precision Numerical Constants (Paragraph 00)
// Standard Cryptoplasm_Firefly_Precision Types/Constants/Variables (Paragraph 01...07)

// 00 - constWithPrecision
// 01 - Condition
// 02 - Context
// 03 - Decimal
// 04 - Form
// 05 - ErrDecimal
// 06 - Rounder
// 07 - LookUpTables
//		Part I - digitsLookupTable
//		Part II - pow10LookupTable
// 08 - loop
//=================================================================================================
//=================================================================================================
//
// 00 - constWithPrecision
//
// Feature:
//
// constWithPrecision implements a look-up table for a constant, rounded-down to
// various precisions. The point is to avoid doing calculations with all the
// digits of the constant when a smaller precision is required.
// Using the following constants, 3 Functions are need for this
//
// 1)Types, 2)Constants, 3)Variables, 4)Functions
//
//===========
//  1)Types
//===========
type constWithPrecision struct {
	unrounded Decimal
	vals      []Decimal
}
//
//===========
//2)Constants
//===========
const strLn10 		= "2.3025850929940456840179914546843642076011014886287729760333279009675726096773524802359972050895982983419677840422862486334095254650828067566662873690987816894829072083255546808437998948262331985283935053089653777326288461633662222876982198867465436674744042432743651550489343149393914796194044002221051017141748003688084012647080685567743216228355220114804663715659121373450747856947683463616792101806445070648000277502684916746550586856935673420670581136429224554405758925724208241314695689016758940256776311356919292033376587141660230105703089634572075440370847469940168269282808481184289314848524948644871927809676271275775397027668605952496716674183485704422507197965004714951050492214776567636938662976979522110718264549734772662425709429322582798502585509785265383207606726317164309505995087807523710333101197857547331541421808427543863591778117054309827482385045648019095610299291824318237525357709750539565187697510374970888692180205189339507238539205144634197265287286965110862571492198849978748873771345686209167058498078280597511938544450099781311469159346662410718466923101075984383191912922307925037472986509290098803919417026544168163357275557031515961135648465461908970428197633658369837163289821744073660091621778505417792763677311450417821376601110107310423978325218948988175979217986663943195239368559164471182467532456309125287783309636042629821530408745609277607266413547875766162629265682987049579549139549180492090694385807900327630179415031178668620924085379498612649334793548717374516758095370882810674524401058924449764796860751202757241818749893959716431055188481952883307466993178146349300003212003277656541304726218839705967944579434683432183953044148448037013057536742621536755798147704580314136377932362915601281853364984669422614652064599420729171193706024449293580370077189810973625332245483669885055282859661928050984471751985036666808749704969822732202448233430971691111368135884186965493237149969419796878030088504089796185987565798948364452120436982164152929878117429733325886079159125109671875109292484750239305726654462762009230687915181358034777012955936462984123664970233551745861955647724618577173693684046765770478743197805738532718109338834963388130699455693993461010907456160333122479493604553618491233330637047517248712763791409243983318101647378233796922656376820717069358463945316169494117018419381194054164494661112747128197058177832938417422314099300229115023621921867233372683856882735333719251034129307056325444266114297653883018223840910261985828884335874559604530045483707890525784731662837019533922310475275649981192287427897137157132283196410034221242100821806795252766898581809561192083917607210809199234615169525990994737827806481280587927319938934534153201859697110214075422827962982370689417647406422257572124553925261793736524344405605953365915391603125244801493132345724538795243890368392364505078817313597112381453237015084134911223243909276817247496079557991513639828810582857405380006533716555530141963322419180876210182049194926514838926922937079"
const strInvLn10 	= "0.4342944819032518276511289189166050822943970058036665661144537831658646492088707747292249493384317483187061067447663037336416792871589639065692210646628122658521270865686703295933708696588266883311636077384905142844348666768646586085135561482123487653435434357317253835622281395603048646652366095539377356176323431916710991411597894962993512457934926357655469077671082419150479910989674900103277537653570270087328550951731440674697951899513594088040423931518868108402544654089797029863286828762624144013457043546132920600712605104028367125954846287707861998992326748439902348171535934551079475492552482577820679220140931468164467381030560475635720408883383209488996522717494541331791417640247407505788767860971099257547730046048656049515610057985741340272675201439247917970859047931285212493341197329877226463885350226083881626316463883553685501768460295286399391633510647555704050513182342988874882120643595023818902643317711537382203362634416478397146001858396093006317333986134035135741787144971453076492968331392399810608505734816169809280016199523523117237676561989228127013815804248715978344927215947562057179993483814031940166771520104787197582531617951490375597514246570736646439756863149325162498727994852637448791165959219701720662704559284657036462635675733575739369673994570909602526350957193468839951236811356428010958778313759442713049980643798750414472095974872674060160650105375287000491167867133309154761441005054775930890767885596533432190763128353570304854020979941614010807910607498871752495841461303867532086001324486392545573072842386175970677989354844570318359336523016027971626535726514428519866063768635338181954876389161343652374759465663921380736144503683797876824369028804493640496751871720614130731804417180216440993200651069696951247072666224570004229341407923361685302418860272411867806272570337552562870767696632173672454758133339263840130320038598899947332285703494195837691472090608812447825078736711573033931565625157907093245370450744326623349807143038059581776957944070042202545430531910888982754062263600601879152267477788232096025228766762416332296812464502577295040226623627536311798532153780883272326920785980990757434437367248710355853306546581653535157943990070326436222520010336980419843015524524173190520247212241110927324425302930200871037337504867498689117225672067268275246578790446735268575794059983346595878592624978725380185506389602375304294539963737367434680767515249986297676732404903363175488195323680087668648666069282082342536311304939972702858872849086258458687045569244548538607202497396631126372122497538854967981580284810494724140453341192674240839673061167234256843129624666246259542760677182858963306586513950932049023032806357536242804315480658368852257832901530787483141985929074121415344772165398214847619288406571345438798607895199435011532826457742311266817183284968697890904324421005272233475053141625981646457044538901148313760708445483457955728303866473638468537587172210685993933008378534367552699899185150879055911525282664"
const (
	// Cbrt uses a quadratic polynomial that approximates the cube root
	// of x when 0.125 <= x <= 1. This approximation is the starting point
	// of the convergence loop. Coefficients are from:
	// https://people.freebsd.org/~lstewart/references/apple_tr_kt32_cuberoot.pdf
	strCbrtC1 = "-0.46946116"
	strCbrtC2 = "1.072302"
	strCbrtC3 = "0.3812513"
)
//
//===========
//3)Variables
//===========
var (
	bigOne  = big.NewInt(1)
	bigTwo  = big.NewInt(2)
	bigFive = big.NewInt(5)
	bigTen  = big.NewInt(10)

	decimalZero      = New(0, 0)
	decimalOneEighth = New(125, -3)
	decimalHalf      = New(5, -1)
	decimalOne       = New(1, 0)
	decimalTwo       = New(2, 0)
	decimalThree     = New(3, 0)
	decimalEight     = New(8, 0)

	decimalCbrtC1 = makeConst(strCbrtC1)
	decimalCbrtC2 = makeConst(strCbrtC2)
	decimalCbrtC3 = makeConst(strCbrtC3)

	// ln(10)
	decimalLn10 = makeConstWithPrecision(strLn10)
	// 1/ln(10)
	decimalInvLn10 = makeConstWithPrecision(strInvLn10)
)
//
//===========
//4)Functions
//===========
// Function 1 - makeConst
//
// First Function needed to implement constWithPrecision Feature
func makeConst(strVal string) *Decimal {
	d := &Decimal{}
	_, _, err := d.SetString(strVal)
	if err != nil {
		panic(err)
	}
	return d
}
//===========
// Function 2 - makeConstWithPrecision
//
// Second Function needed to implement constWithPrecision Feature
func makeConstWithPrecision(strVal string) *constWithPrecision {
	c := &constWithPrecision{}
	if _, _, err := c.unrounded.SetString(strVal); err != nil {
		panic(err)
	}
	// The length of the string might be one higher than the available precision
	// (because of the decimal point), but that's ok.
	maxPrec := uint32(len(strVal))
	for p := uint32(1); p < maxPrec; p *= 2 {
		var d Decimal

		ctx := Context{
			Precision:   p,
			Rounding:    RoundHalfUp,
			MaxExponent: MaxExponent,
			MinExponent: MinExponent,
		}
		_, err := ctx.Round(&d, &c.unrounded)
		if err != nil {
			panic(err)
		}
		c.vals = append(c.vals, d)
	}
	return c
}
//===========
// Function 3 - get
//
// get returns the given constant, rounded down to a precision at least as high
// as the given precision.
// Third Function needed to implement constWithPrecision Feature
func (c *constWithPrecision) get(precision uint32) *Decimal {
	i := 0
	// Find the smallest precision available that's at least as high as precision,
	// i.e. Ceil[ log2(p) ] = 1 + Floor[ log2(p-1) ]
	if precision > 1 {
		precision--
		i++
	}
	for precision >= 16 {
		precision /= 16
		i += 4
	}
	for precision >= 2 {
		precision /= 2
		i++
	}
	if i >= len(c.vals) {
		return &c.unrounded
	}
	return &c.vals[i]
}
//=================================================================================================
//
// 01 - Condition
//
// 1)Types, 2)Constants, 3)..., 4)...
//
//===========
//  1)Types
//===========
type Condition uint32
// Condition holds condition flags.
//
//===========
//2)Constants
//===========
const (
	// SystemOverflow is raised when an exponent is greater than MaxExponent.
	SystemOverflow Condition = 1 << iota
	// SystemUnderflow is raised when an exponent is less than MinExponent.
	SystemUnderflow
	// Overflow is raised when the exponent of a result is too large to be
	// represented.
	Overflow
	// Underflow is raised when a result is both subnormal and inexact.
	Underflow
	// Inexact is raised when a result is not exact (one or more non-zero
	// coefficient digits were discarded during rounding).
	Inexact
	// Subnormal is raised when a result is subnormal (its adjusted exponent is
	// less than Emin), before any rounding.
	Subnormal
	// Rounded is raised when a result has been rounded (that is, some zero or
	// non-zero coefficient digits were discarded).
	Rounded
	// DivisionUndefined is raised when both division operands are 0.
	DivisionUndefined
	// DivisionByZero is raised when a non-zero dividend is divided by zero.
	DivisionByZero
	// DivisionImpossible is raised when integer division cannot be exactly
	// represented with the given precision.
	DivisionImpossible
	// InvalidOperation is raised when a result would be undefined or impossible.
	InvalidOperation
	// Clamped is raised when the exponent of a result has been altered or
	// constrained in order to fit the constraints of the Decimal representation.
	Clamped
)
const (
	errExponentOutOfRangeStr = "exponent out of range"
)
//=================================================================================================
//
// 02 - Context
//
// Context maintains options for Decimal operations. It can safely be used
// concurrently, but not modified concurrently. Arguments for any method
// can safely be used as both result and operand.
//
// 1)Types, 2)Constants, 3)Variables, 4)...
//
//===========
//  1)Types
//===========
type Context struct {
	// Precision is the number of places to round during rounding; this is
	// effectively the total number of digits (before and after the decimal
	// point).
	Precision uint32
	// MaxExponent specifies the largest effective exponent. The
	// effective exponent is the value of the Decimal in scientific notation. That
	// is, for 10e2, the effective exponent is 3 (1.0e3). Zero (0) is not a special
	// value; it does not disable this check.
	MaxExponent int32
	// MinExponent is similar to MaxExponent, but for the smallest effective
	// exponent.
	MinExponent int32
	// Traps are the conditions which will trigger an error result if the
	// corresponding Flag condition occurred.
	Traps Condition
	// Rounding specifies the Rounder to use during rounding. RoundHalfUp is used if
	// empty or not present in Roundings.
	Rounding string
}
//
//===========
//2)Constants
//===========
const (
	// DefaultTraps is the default trap set used by BaseContext.
	DefaultTraps = SystemOverflow |
		SystemUnderflow |
		Overflow |
		Underflow |
		Subnormal |
		DivisionUndefined |
		DivisionByZero |
		DivisionImpossible |
		InvalidOperation

	errZeroPrecisionStr = "Context may not have 0 Precision for this operation"
)
//
//===========
//3)Variables
//===========
// BaseContext is a useful default Context. Should not be mutated.
var BaseContext = Context{
	// Disable rounding.
	Precision: 0,
	// MaxExponent and MinExponent are set to the packages's limits.
	MaxExponent: MaxExponent,
	MinExponent: MinExponent,
	// Default error conditions.
	Traps: DefaultTraps,
}
//=================================================================================================
//
// 03 - Decimal
//
// Decimal is an arbitrary-precision decimal. Its value is:
//
//     Negative × Coeff × 10**Exponent
//
// Coeff must be positive. If it is negative results may be incorrect and
// apd may panic.
//
// 1)Types, 2)Constants, 3)Variables, 4)...
//
//===========
//  1)Types
//===========
type Decimal struct {
	Form     Form
	Negative bool
	Exponent int32
	Coeff    big.Int
}
//
//===========
//2)Constants
//===========
const (
	// MaxExponent is set because both upscale and Round
	// perform a calculation of 10^x, where x is an exponent. This is done by
	// big.Int.Exp. This restriction could be lifted if better algorithms were
	// determined during upscale and Round that don't need to perform Exp.

	// MaxExponent is the highest exponent supported. Exponents near this range will
	// perform very slowly (many seconds per operation).
	MaxExponent = 100000
	// MinExponent is the lowest exponent supported with the same limitations as
	// MaxExponent.
	MinExponent = -MaxExponent
)
//
//===========
//3)Variables
//===========
var (
	decimalNaN      = &Decimal{Form: NaN}
	decimalInfinity = &Decimal{Form: Infinite}
)
var _ fmt.Formatter = decimalZero // *Decimal must implement fmt.Formatter
//=================================================================================================
//
// 04 - Form
//
// Form specifies the form of a Decimal.
// 1)Types, 2)Constants, 3)Variables, 4)...
//
//===========
//  1)Types
//===========
type Form int
//
//===========
//2)Constants
//===========
const (
	// These constants must be in the following order. CmpTotal assumes that
	// the order of these constants reflects the total order on decimals.

	// Finite is the finite form.
	Finite Form = iota
	// Infinite is the infinite form.
	Infinite
	// NaNSignaling is the signaling NaN form. It will always raise the
	// InvalidOperation condition during an operation.
	NaNSignaling
	// NaN is the NaN form.
	NaN
)
const formName = "FiniteInfiniteNaNSignalingNaN"
//
//===========
//3)Variables
//===========
var formIndex = [...]uint8{0, 6, 14, 26, 29}
//=================================================================================================
//
// 05 - ErrDecimal
//
// ErrDecimal performs operations on decimals and collects errors during
// operations. If an error is already set, the operation is skipped. Designed to
// be used for many operations in a row, with a single error check at the end.
//
// 1)Types, 2)..., 3)..., 4)...
//
//===========
//  1)Types
//===========
type ErrDecimal struct {
	err error
	Ctx *Context
	// Flags are the accumulated flags from operations.
	Flags Condition
}
//=================================================================================================
//
// 06 - Rounder
//
// Rounder defines a function that returns true if 1 should be added to the
// absolute value of a number being rounded. result is the result to which
// the 1 would be added. neg is true if the number is negative. half is -1
// if the discarded digits are < 0.5, 0 if = 0.5, or 1 if > 0.5.
// 1)Types, 2)Constants, 3)Variables, 4)Functions
//
//===========
//  1)Types
//===========
type Rounder func(result *big.Int, neg bool, half int) bool
//
//===========
//2)Constants
//===========
const (
	// RoundDown rounds toward 0; truncate.
	RoundDown = "down"
	// RoundHalfUp rounds up if the digits are >= 0.5.
	RoundHalfUp = "half_up"
	// RoundHalfEven rounds up if the digits are > 0.5. If the digits are equal
	// to 0.5, it rounds up if the previous digit is odd, always producing an
	// even digit.
	RoundHalfEven = "half_even"
	// RoundCeiling towards +Inf: rounds up if digits are > 0 and the number
	// is positive.
	RoundCeiling = "ceiling"
	// RoundFloor towards -Inf: rounds up if digits are > 0 and the number
	// is negative.
	RoundFloor = "floor"
	// RoundHalfDown rounds up if the digits are > 0.5.
	RoundHalfDown = "half_down"
	// RoundUp rounds away from 0.
	RoundUp = "up"
	// Round05Up rounds zero or five away from 0; same as round-up, except that
	// rounding up only occurs if the digit to be rounded up is 0 or 5.
	Round05Up = "05up"
)
//
//===========
//3)Variables
//===========
var (
	// Roundings defines the set of Rounders used by Context. Users may add their
	// own, but modification of this map is not safe during any other parallel
	// Context operations.
	Roundings = map[string]Rounder{
		RoundDown:     roundDown,
		RoundHalfUp:   roundHalfUp,
		RoundHalfEven: roundHalfEven,
		RoundCeiling:  roundCeiling,
		RoundFloor:    roundFloor,
		RoundHalfDown: roundHalfDown,
		RoundUp:       roundUp,
		Round05Up:     round05Up,
	}
)
//
//===========
//4)Functions
//===========
func roundDown(*big.Int, bool, int) bool {
	return false
}
//===========
func roundHalfUp(_ *big.Int, _ bool, half int) bool {
	return half >= 0
}
//===========
func roundHalfEven(result *big.Int, _ bool, half int) bool {
	if half > 0 {
		return true
	}
	if half < 0 {
		return false
	}
	return result.Bit(0) == 1
}
//===========
func roundCeiling(_ *big.Int, neg bool, _ int) bool {
	return !neg
}
//===========
func roundFloor(_ *big.Int, neg bool, _ int) bool {
	return neg
}
//===========
func roundHalfDown(_ *big.Int, _ bool, half int) bool {
	return half > 0
}
//===========
func roundUp(*big.Int, bool, int) bool {
	return true
}
//===========
func round05Up(result *big.Int, _ bool, _ int) bool {
	z := new(big.Int)
	z.Rem(result, bigFive)
	if z.Sign() == 0 {
		return true
	}
	z.Rem(result, bigTen)
	return z.Sign() == 0
}
//===========
// roundAddOne adds 1 to abs(b).
func roundAddOne(b *big.Int, diff *int64) {
	if b.Sign() < 0 {
		panic("unexpected negative")
	}
	nd := NumDigits(b)
	b.Add(b, bigOne)
	nd2 := NumDigits(b)
	if nd2 > nd {
		b.Quo(b, bigTen)
		*diff++
	}
}
//=================================================================================================
//
// 07 - LookUpTables
//
// 1)Types, 2)..., 3)..., 4)...
//
//===========
//  1)Types
//===========
type tableVal struct {
	digits  int64
	border  big.Int
	nborder big.Int
}
//================================================
//
// Part I - digitsLookupTable
//
// 1)..., 2)Constants, 3)Variables
//
//===========
//  1)Types
//===========
const digitsTableSize = 128

var digitsLookupTable [digitsTableSize + 1]tableVal
//================================================
//
// Part II - pow10LookupTable
//
// 1)..., 2)Constants, 3)Variables
//
//===========
//2)Constants
//===========
const powerTenTableSize = 128
//
//===========
//3)Variables
//===========
var pow10LookupTable [powerTenTableSize + 1]big.Int
//=================================================================================================
//
// 08 - Loop
//
// 1)Types, 2)Constants, 3)..., 4)Function
//
//===========
//  1)Types
//===========
type loop struct {
	c             *Context
	name          string // The name of the function we are evaluating.
	i             uint64 // Loop count.
	precision     int32
	maxIterations uint64   // When to give up.
	arg           *Decimal // original argument to function; only used for diagnostic.
	prevZ         *Decimal // Result from the previous iteration.
	delta         *Decimal // |Change| from previous iteration.
}
//
//===========
//2)Constants
//===========
const digitsToBitsRatio = math.Ln10 / math.Ln2
//
//===========
//4)Functions
//===========
// Function done reports whether the loop is done.
// If it does not converge after the maximum number of iterations,
// it returns an error.
func (l *loop) done(z *Decimal) (bool, error) {
	if _, err := l.c.Sub(l.delta, l.prevZ, z); err != nil {
		return false, err
	}
	sign := l.delta.Sign()
	if sign == 0 {
		return true, nil
	}
	if sign < 0 {
		// Convergence can oscillate when the calculation is nearly
		// done and we're running out of bits. This stops that.
		// See next comment.
		l.delta.Neg(l.delta)
	}

	// We stop if the delta is smaller than a change of 1 in the
	// (l.precision)-th digit of z. Examples:
	//
	//   p   = 4
	//   z   = 12345.678 = 12345678 * 10^-3
	//   eps =    10.000 = 10^(-4+8-3)
	//
	//   p   = 3
	//   z   = 0.001234 = 1234 * 10^-6
	//   eps = 0.00001  = 10^(-3+4-6)
	eps := Decimal{Coeff: *bigOne, Exponent: -l.precision + int32(z.NumDigits()) + z.Exponent}
	if l.delta.Cmp(&eps) <= 0 {
		return true, nil
	}
	l.i++
	if l.i == l.maxIterations {
		return false, errors.Errorf(
			"%s %s: did not converge after %d iterations; prev,last result %s,%s delta %s precision: %d",
			l.name, l.arg.String(), l.maxIterations, z, l.prevZ, l.delta, l.precision,
		)
	}
	l.prevZ.Set(z)
	return false, nil
}
//=================================================================================================