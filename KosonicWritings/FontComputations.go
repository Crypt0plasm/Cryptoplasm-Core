package KosonicWritings

import (
    b "Cryptoplasm-Core/Cryptoplasm_Blockchain_Constants"
    "fmt"
    p "github.com/Crypt0plasm/Firefly-APD"
    svg "github.com/ajstarks/svgo"
    "log"
    "os"
    "strings"
)

//================================================
//
// 	CONTENT LIST:
//
//      Structure Definitions
//              01  - KerningPair
//
//      Variable Definitions
//              01  - KP1U                                      Kerning Pairs that decrease length by 1
//              02  - KP2U                                      Kerning Pairs that decrease length by 2
//
//      Functions:
//
//      01 Drawing Functions
//              01  - PrintStar                                 Main calling Function that prints the Koson/Unity Star
//              02  - DrawKosonStar                             The function that is being called by PrintStar Func.
//              02a - DrawKosonKanonText                        1st SubDrawing Function used by the DrawKosonStar Func.
//              02b - DrawKosonStarText                         2nd SubDrawing Function used by the DrawKosonStar Func.
//
//      02 Kerning Computation Functions
//              01  - ComputeKerning                            Kerning Computation Function
//              02  - GetRuneWidth                              Line length computations based on Kerning Values
//              03  - CheckTierWidth                            Computes rune width based on its tier
//
//      03 Kanon Text generating Functions			One Time only used functions
//              01  - MakeKanonString                           Generates and empty Kanon string
//              02  - MakeKanonString2                          Generates and empty Kanon string, another variant
//              03  - PrintKosonicString                        Function used to print the generated string
//
//      04 Testing Function					Functions created to understand svg mechanics
//		01  - ScaleTester				Created to understand svg scaling mechanics
//
//================================================
//KerningPairs
type KerningPair struct {
    Pair1 string
    Pair2 string
}

//Wanted Stroke Size Definition
var (
    KosonStarStrokeSize = "3")

//KerningPairs Definitions
//Kerning Pairs
var (

    //Kerning Pairs that decrease the length by 1
    KP1U = [] KerningPair{
        //a Pair
	{"a","c"},
	{"a","g"},
	{"a","o"},
		{"o","a"},
	{"a","ö"},
		{"ö","a"},
	{"a","q"},
		{"q","a"},
	{"a","u"},
		{"u","a"},

		//ä Pair
	{"ä","c"},
	{"ä","g"},
	{"ä","o"},
		{"o","ä"},
	{"ä","ö"},
		{"ö","ä"},
	{"ä","q"},
		{"q","ä"},
	{"ä","u"},
		{"u","ä"},

	//b No B Pairs

	//c Pairs
	{"c","o"},
	{"c","ö"},
	{"c","q"},

	//d Pairs
	{"d","a"},
	{"d","ä"},
	{"d","v"},

	//e No e Pairs

	//f Pairs
	{"f","a"},
	{"f","ä"},

	//g No g Pairs
	//h No g Pairs
	//i No g Pairs
	//j No g Pairs

	//k Pairs
	{"k","o"},
	{"k","ö"},
	{"k","q"},

	//l Pairs - see kerning 4
	{"l","o"},
	{"l","ö"},
	{"l","q"},

	//m No g Pairs
	//m No g Pairs

	//o Pairs
	{"o","t"},
		{"t","o"},
	{"o","v"},
		{"v","o"},
	{"o","x"},
		{"x","o"},
	{"o","y"},
		{"y","o"},

	//ö Pairs
	{"ö","t"},
		{"t","ö"},
	{"ö","v"},
		{"v","ö"},
	{"ö","x"},
		{"x","ö"},
	{"ö","y"},
		{"y","ö"},


	//q Pairs
	{"q","t"},
		{"t","q"},
	{"q","v"},
		{"v","q"},
	{"q","x"},
		{"x","q"},
	{"q","y"},
		{"y","q"},

	//r No p Pairs
	//s No p Pairs
	//u no u Pairs
	//w No w Pairs
	//x Pairs already entered
	//z no z Pairs
    }

    //Kerning Pairs that decrease the length by 2
    KP2U = [] KerningPair{
        //a Pair
	{"a","t"},
		{"t","a"},
	{"a","v"},
		{"v","a"},
	{"a","y"},
		{"y","a"},

		//ä Pair
	{"ä","t"},
		{"t","ä"},
	{"ä","v"},
		{"v","ä"},
	{"ä","y"},
		{"y","ä"},

	//f Pairs
	{"f","."},
	{"f",","},

		//l Pairs
	{"l","t"},
	{"l","v"},
	{"l","y"},

	//p Pair
	{"p","a"},
	{"p","ä"},

	//t pairs
	{"t","."},
		{".","t"},
	{"t",","},
		{",","t"},

	//v Pairs
	{"v","."},
		{".","v"},
	{"v",","},
		{",","v"},

	//p Pairs
	{"p","."},
	{"p",","},

	//y Pairs
	{"y","."},
		{".","y"},
	{"y",","},
		{",","y"},
    }
)
//
//======================================================================================================================
//======================================================================================================================
//
// Drawing Functions
//
//      01.01
//      PrintStar Function
//
//	The 1st Function that prints the Star; Its called by calling:
// 	om.PrintStar(om.KossonSILVER660)
//
func PrintStar (Star Kanon) {
    DrawKosonStar(Star)
}
//
//======================================================================================================================
//
//
//      01.02
//      DrawKosonStar Function
//
//	The Actual Function that prints the Star
//	It is being called by the PrintStar Function
//
func DrawKosonStar (Unit Kanon) *svg.SVG {
    var (
        Width = 18100
        Height = 18100
	S1 = "fill=\"none\""
	S2 = "stroke=\"black\""
	S3 = "stroke-width=\"" + KosonStarStrokeSize +"\""
	Zero = p.NFS("0")
    )

    //Outputs SVG content directly to external File.
    OutputSVGName := Unit.Name + ".svg"
    OutputVariable, err := os.Create(OutputSVGName)

    if err != nil {
	log.Fatal(err)
    }
    defer OutputVariable.Close()
    KosonStar := svg.New(OutputVariable)

    DecimalToInt := func (Number *p.Decimal) int {
        NumberInt64, _ := Number.Int64()
        NumberInt := int(NumberInt64)
        return NumberInt
    }

    //Get Tablet Size and XY for the OM Sign
    CorrectTabletSize := GetCorrectBoardSize(Unit.Value)
    fmt.Println("Correct Tablet Size is",CorrectTabletSize.Name)
    fmt.Println("Board Size is",CorrectTabletSize.Size)
    TextLines := SplitString(CorrectTabletSize, Unit.Value)

    //Print Info about Text Lines
    fmt.Println("Line Number is",len(TextLines))
    //Printing Lines
    for i:=0; i<len(TextLines); i++ {
        if i < 10 {
            fmt.Println("LINE ",i,":: ",TextLines[i])
            }else {
                fmt.Println("LINE",i,":: ",TextLines[i])
                }
    }

    FreeLines := int64(len(CorrectTabletSize.Lines) - len(TextLines))
    OMScalingFactor := b.MULxc(p.NFI(FreeLines),CorrectTabletSize.ScalingFactor)
    OMHeight := b.MULxc(OMScalingFactor,p.NFI(1600))
    OMY := b.ADDxc(b.PRDxc(p.NFI(int64(len(TextLines))),CorrectTabletSize.ScalingFactor,p.NFI(1600)),p.NFI(10450))
    OMYint := DecimalToInt(OMY)

    TextLength := b.PRDxc(CorrectTabletSize.ScalingFactor,p.NFI(200),p.NFI(CorrectTabletSize.Lines[0]))
    V1 := b.SUBxc(TextLength,OMHeight)
    V2 := b.DIVxc(V1,p.NFI(2))
    OMX := b.ADDxc(CorrectTabletSize.AnchorX,V2)
    OMXint := DecimalToInt(OMX)

    fmt.Println("FreeLines are", FreeLines, "HeightLeft is", OMHeight)

    //Starts creating the Star and its interior basorelief geometries

    KosonStar.Start(Width,Height)
    KosonStar.Path(ComputeGlyphCode(CenterBasorelief,Zero,Zero,false), S1, S2, S3)
    KosonStar.Path(ComputeGlyphCode(LeftMinorBasorelief,Zero,Zero,false), S1, S2, S3)
    KosonStar.Path(ComputeGlyphCode(RightMinorBasorelief,Zero,Zero,false), S1, S2, S3)
    KosonStar.Path(ComputeGlyphCode(LeftMajorBasorelief,Zero,Zero,false), S1, S2, S3)
    KosonStar.Path(ComputeGlyphCode(RightMajorBasorelief,Zero,Zero,false), S1, S2, S3)
    KosonStar.Path(ComputeGlyphCode(CrownBasorelief,Zero,Zero,false), S1, S2, S3)
    KosonStar.Path(ComputeGlyphCode(TextBasorelief,Zero,Zero,false), S1, S2, S3)
    KosonStar.Path(ComputeGlyphCode(KanonBasorelief,Zero,Zero,false), S1, S2, S3)
    KosonStar.Circle(9050,9050,1200, S1, S2, S3)
    KosonStar.Circle(9050,9050,7400, S1, S2, S3)
    KosonStar.Circle(9050,9050,8800, S1, S2, S3)

    //Kanon Name - Name is Kanon.Name (its title)
    KosonStar.Translate(KanonNamePlateX,KanonNamePlateY)
    ScalingFactorKanon,_ := Tablet400.ScalingFactor.Float64()
    KosonStar.Scale(ScalingFactorKanon)
	//Defining Style Used for Draw
	KosonKanonTextStyle := StyleSVG{
	    "black",
	    b.DTS(b.DIVxc(p.NFS(KosonStarStrokeSize),Tablet400.ScalingFactor)),
	    "none"}
    	//Effective Draw
    	DrawKosonKanonText(Unit.Name,KosonKanonTextStyle,OutputVariable)
    KosonStar.Gend()
    KosonStar.Gend()

    //Kanon Value(Text) - Text is Kanon.Value (its content)
    vartoprint := DecimalToInt(CorrectTabletSize.AnchorX)
    fmt.Println("AnchorX for Text is",vartoprint)
    KosonStar.Translate(DecimalToInt(CorrectTabletSize.AnchorX),DecimalToInt(CorrectTabletSize.AnchorY))
    ScalingFactor,_ :=  CorrectTabletSize.ScalingFactor.Float64()
    KosonStar.Scale(ScalingFactor)
	//Defining Style Used for Draw
	KosonStarTextStyle := StyleSVG{
	    "black",
	    b.DTS(b.DIVxc(p.NFS(KosonStarStrokeSize),CorrectTabletSize.ScalingFactor)),
	    "none"}
	//Effective Draw
    	DrawKosonStarText(Unit.Value, KosonStarTextStyle,OutputVariable)
    KosonStar.Gend()
    KosonStar.Gend()

    //Draws the OM Sign, using a scaling factor, at given coordinates.
    if FreeLines > 0 {
	KosonStar.Translate(OMXint,OMYint)
	ScalingFactorOM,_ := OMScalingFactor.Float64()
	KosonStar.Scale(ScalingFactorOM)
	    //Defining Style Used for Draw
	    OMStyle := StyleSVG{
		"black",
		b.DTS(b.DIVxc(p.NFS(KosonStarStrokeSize),OMScalingFactor)),
		"none"}
	    //Effective Draw
	    DrawOM(Zero,Zero,OMStyle,OutputVariable)
	KosonStar.Gend()
	KosonStar.Gend()
    }

    KosonStar.End()
    return KosonStar
}
//
//======================================================================================================================
//
//
//      01.02a
//      DrawKosonKanonText Function
//
//	Function that draws a "text" of characters, such as for example "Kanon  9. 999"
//	It is used for short texts, that are taken as a "single word" (spaces included)
//
func DrawKosonKanonText(Text string, Style StyleSVG, OutputVariable *os.File) {
    V1 := p.NFI(0)
    DrawWord(V1, p.NFI(0), Text, Style, OutputVariable)
}
//
//======================================================================================================================
//
//
//      01.02b
//      DrawKosonStarText Function
//
//	Function that draws a "text" of characters, that is extra long
//	It has logic that splits it in lines, and centers the resulting words based on their length.
//
func DrawKosonStarText (Text string, Style StyleSVG, OutputVariable *os.File) {

    CorrectTabletSize := GetCorrectBoardSize(Text)
    TextLines := SplitString(CorrectTabletSize, Text)
    Offset := ComputeStartingOffset(CorrectTabletSize,TextLines)
    fmt.Println("Offset calculat in DrawKosonStarText este",Offset)

    WLO := new(p.Decimal)
    WLOS := new(p.Decimal)

    for i:=0; i<len(TextLines); i++ {
        var (
            X, Y = new(p.Decimal), new(p.Decimal)
	)
        if i == 0 {
            X = b.MULxc(p.NFI(Offset[i]),p.NFI(100))
	    //KosonText.Line(0,0,0,1600, S1, S2, S3)
	} else {
	    V1 := b.SUBxc(p.NFI(CorrectTabletSize.Lines[0]),p.NFI(CorrectTabletSize.Lines[i]))
	    V2 := b.DIVxc(V1,p.NFI(2))
	    V3 := b.DIVxc(p.NFI(Offset[i]),p.NFI(2))
	    V4 := b.ADDxc(V2,V3)
	    X = b.MULxc(V4,p.NFI(200))
	}
	Y = b.MULxc(p.NFI(int64(i)),p.NFI(1600))
	WordsChainToPrint := strings.Split(TextLines[i]," ")
	for j:=0; j<len(WordsChainToPrint); j++ {
	    if j == 0{
		WLO = DrawWord(X, Y, WordsChainToPrint[j], Style, OutputVariable)
		WLOS = b.ADDxc(WLO,p.NFI(1600))		//Offset given by space
	    } else {
		WLO = DrawWord(WLOS, Y, WordsChainToPrint[j], Style, OutputVariable)
		WLOS = b.ADDxc(WLO,p.NFI(1600))		//Offset given by space
	    }
	}
    }
    return
}
//
//======================================================================================================================
//======================================================================================================================
//
// KERNING Computation Functions
//
// 	02.01
//	ComputeKerning Function; Main Kerning Computation Function
//
// 	Computes Kerning Value between 2 glyphs, based on the Kerning Pairs defined at the beginning of the code
//
func ComputeKerning (FirstGlyph, SecondGlyph string) (KerningValue int64) {
    var (
	Pair = KerningPair {FirstGlyph,SecondGlyph}
    )

    for i:=0; i<len(KP1U); i++ {
	if Pair == KP1U[i] {
	    KerningValue = 1
	    break
	} else if KerningValue == 0 {
	    for j:=0; j<len(KP2U); j++ {
		if Pair == KP2U[j] {
		    KerningValue = 2
		    break
		}
	    }
	}
    }
    //fmt.Println("For Pair",FirstGlyph,"and ",SecondGlyph,"Kerning is",KerningValue)
    return KerningValue
}
//
//======================================================================================================================
//
// 	02.02
//	GetRuneWidth Function; computes rune width
//
// 	Used in the Line length computations based on Kerning Values
//	1 Units represent 200 width value (One Eighth). (Glyph Definition is 1600x1600).
//
func GetRuneWidth (Glyph rune) int64 {
    var Result int64
    if CheckTierWidth(Glyph, WTier2) == true {
	Result = 2
    } else if CheckTierWidth(Glyph, WTier3) == true {
	Result = 3
    } else if CheckTierWidth(Glyph, WTier4) == true {
	Result = 4
    } else if CheckTierWidth(Glyph, WTier5) == true {
	Result = 5
    } else if CheckTierWidth(Glyph, WTier6) == true {
	Result = 6
    } else if CheckTierWidth(Glyph, WTier7) == true {
	Result = 7
    } else if CheckTierWidth(Glyph, WTier8) == true {
	Result = 8
    }
    return Result
}
//
//======================================================================================================================
//
// 	02.03
//	CheckTierWidth Function; computes rune width
//
// 	Used in the GetRuneWidth function.
//
func CheckTierWidth (Glyph rune, Tier []string) bool {
    var Result bool
    for i:=0; i<len(Tier); i++ {
	if string(Glyph) == Tier[i] {
	    Result = true
	    break
	}
    }
    return Result
}
//
//======================================================================================================================
//======================================================================================================================
//
// Kanon Text generating Functions			One Time only used functions
//
//	03.01
//	MakeKanonString Function
//
// 	Generates the Empty Kanon String that was used to populate Om.go Kanon List
//
func MakeKanonString (KanonNumber *p.Decimal, KanonSize *p.Decimal, StartingPosition *p.Decimal)  []string {
    DecimalToInt := func (Number *p.Decimal) int {
	NumberInt64, _ := Number.Int64()
	NumberInt := int(NumberInt64)
	return NumberInt
    }

    KanonSizeInt := DecimalToInt(KanonSize)
    KosonicChain := make([]string, KanonSizeInt)

    var (
        //KosonicChain []string
        String01 string
	String02 string
	String03 string
	String04 string
	String05 string
	String06 string
    )
    String01 = "KossonCOPPER"
    String03 = " = Kanon {\"Kanon "
    String04 = KanonNumber.String() + "."
    String06 = "\", \"\"}"

    for i:=0; i<len(KosonicChain); i++ {


	//CopperCoinNumber
	CoinNumber := b.ADDxc(StartingPosition,p.NFI(int64(i)))
	if b.Count4Coma(CoinNumber) == 1 {
	    String02 = "000" + CoinNumber.String()
	} else if b.Count4Coma(CoinNumber) == 2 {
	    String02 = "00" + CoinNumber.String()
	} else if b.Count4Coma(CoinNumber) == 3 {
	    String02 = "0" + CoinNumber.String()
	} else if b.Count4Coma(CoinNumber) == 4 {
	    String02 = CoinNumber.String()
	}

	//KanonNumber
	Verse := b.ADDxc(p.NFI(int64(i)),p.NFI(1))
	if b.Count4Coma(Verse) == 1 {
	    String05 = "   " + Verse.String()
	} else if b.Count4Coma(Verse) == 2 {
	    String05 = "  " + Verse.String()
	} else if b.Count4Coma(Verse) == 3 {
	    String05 = " " + Verse.String()
	} else if b.Count4Coma(Verse) == 4 {
	    String05 = Verse.String()
	}

	KosonicChain[i] = String01 + String02 + String03 + String04 + String05 + String06

    }

    return KosonicChain
}
//
//======================================================================================================================
//
//	03.02
//	MakeKanonString2 Function
//
// 	Generates the Empty Kanon String that was used to populate Om.go Kanon List, second variant
//
func MakeKanonString2 () []string {
    KosonicChain := make([]string, 9100)
    KosonicChainBig := make([]string, 910)
    var (
        WordBase string
        NumberS string
        Word,WordB,K string
        W0,W1,W2,W3,W4,W5,W6,W7,W8,W9 string
    )
    WordBase = "KossonCOPPER"
    K = ","

    for i:=0; i<len(KosonicChain); i++ {
	Number := b.ADDxc(p.NFI(int64(i)),p.NFI(1))
	if b.Count4Coma(Number) == 1 {
	    NumberS = "000" + Number.String()
	} else if b.Count4Coma(Number) == 2 {
	    NumberS = "00" + Number.String()
	} else if b.Count4Coma(Number) == 3 {
	    NumberS = "0" + Number.String()
	} else if b.Count4Coma(Number) == 4 {
	    NumberS = Number.String()
	}
	Word = WordBase + NumberS
	KosonicChain[i] = Word
    }

    for i:=0; i<len(KosonicChainBig); i++ {
        W0 = KosonicChain[i*10]
	W1 = KosonicChain[i*10+1]
	W2 = KosonicChain[i*10+2]
	W3 = KosonicChain[i*10+3]
	W4 = KosonicChain[i*10+4]
	W5 = KosonicChain[i*10+5]
	W6 = KosonicChain[i*10+6]
	W7 = KosonicChain[i*10+7]
	W8 = KosonicChain[i*10+8]
	W9 = KosonicChain[i*10+9]
	WordB = W0+K+W1+K+W2+K+W3+K+W4+K+W5+K+W6+K+W7+K+W8+K+W9+K
	KosonicChainBig[i] = WordB
    }

    return KosonicChainBig

}
//
//======================================================================================================================
//
//	03.03
//	PrintKosonicString Function
//
// 	Used to print the generated Kosonic String
//
func PrintKosonicString (KosonicChain []string) {
    OutputName := "KosonicStrings.txt"
    fmt.Println("Printing String to", OutputName)
    OutputFile, err := os.Create(OutputName)
    if err != nil {
	log.Fatal(err)
    }
    defer OutputFile.Close()

    for i:=0; i<len(KosonicChain); i++ {
	_, _ = fmt.Fprintln(OutputFile, KosonicChain[i])
    }
}
//
//======================================================================================================================
//======================================================================================================================
//
// Testing Functions
//
//	04.01
// 	ScaleTester Function
//
//	A scale testing function;
//	Tests how scaling works in a svg function to better understand scaling implementation in future functions
//
func ScaleTester () *svg.SVG {
    var (
	Width = 1000
	Height = 1000



	S1 = "fill=\"none\""
	S2 = "stroke=\"black\""
	S3 = "stroke-width=\"7\""
	//Zero = p.NFS("0")
    )

    OutputVariable, err := os.Create("data.txt")

    if err != nil {
	log.Fatal(err)
    }

    defer OutputVariable.Close()
    KosonStar := svg.New(OutputVariable)

    KosonStar.Start(Width,Height)
    KosonStar.Square(0,0,1000,S1,S2,S3)

    SquareUpperCornerX := 0
    SquareUpperCornerY := 0
    //Tx := 300
    //Ty := 700

    KosonStar.Translate(0,600)
    KosonStar.Scale(0.5)
    KosonStar.Square(SquareUpperCornerX,SquareUpperCornerY,500,S1,S2,S3)
    KosonStar.Gend()
    KosonStar.Gend()
    KosonStar.End()

    return KosonStar
}
//
//======================================================================================================================