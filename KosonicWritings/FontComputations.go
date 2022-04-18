package KosonicWritings

import (
    b "Cryptoplasm-Core/Cryptoplasm_Blockchain_Constants"
    "fmt"
    "log"

    //b "Cryptoplasm-Core/Cryptoplasm_Blockchain_Constants"
    p "github.com/Crypt0plasm/Firefly-APD"
    svg "github.com/ajstarks/svgo"
    "os"
    "strings"
)

//KerningPairs
type KerningPair struct {
    Pair1 string
    Pair2 string
}
//KerningPairs Definitions
//Kerning Pairs
var (
    //Kerning Pairs that decrease the length by 2
    KP2U = [] KerningPair{
	{"a","v"},
	{"ä","v"},
	{"v","a"},
	{"v","ä"},
	{"t","."},
	{".","t"},
	{"t",","},
	{",","t"},
	{"v","."},
	{".","v"},
	{"v",","},
	{",","v"},
	{"p","."},
	{"p",","},
	{"f","."},
	{"f",","},
	{"y","."},
	{".","y"},
	{"y",","},
	{",","y"},
	{"l","t"},
    }
)

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

func DrawKosonStar (Unit Kanon) *svg.SVG {
    var (
        Width = 18100
        Height = 18100
	S1 = "fill=\"none\""
	S2 = "stroke=\"black\""
	S3 = "stroke-width=\"3\""
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
    TextLines := SplitString(CorrectTabletSize, Unit.Value)
    //Offset := ComputeStartingOffset(CorrectTabletSize,TextLines)

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

    //Kanon Text

    vartoprint := DecimalToInt(CorrectTabletSize.AnchorX)
    fmt.Println("AnchorX for Text is",vartoprint)
    KosonStar.Translate(DecimalToInt(CorrectTabletSize.AnchorX),DecimalToInt(CorrectTabletSize.AnchorY))
    ScalingFactor,_ :=  CorrectTabletSize.ScalingFactor.Float64()
    KosonStar.Scale(ScalingFactor)
    DrawKosonStarText(Unit.Value,CorrectTabletSize.ScalingFactor,OutputVariable)
    KosonStar.Gend()
    KosonStar.Gend()

    //Kanon Name
    KosonStar.Translate(KanonNamePlateX,KanonNamePlateY)
    ScalingFactorKanon,_ := Tablet400.ScalingFactor.Float64()
    KosonStar.Scale(ScalingFactorKanon)
    DrawKosonKanonText(Unit.Name,Tablet400.ScalingFactor,OutputVariable)
    KosonStar.Gend()
    KosonStar.Gend()

    //OM Sign
    if FreeLines > 0 {
	KosonStar.Translate(OMXint,OMYint)
	ScalingFactorOM,_ := OMScalingFactor.Float64()
	KosonStar.Scale(ScalingFactorOM)
	DrawOM(Zero,Zero,OMScalingFactor,OutputVariable)
	KosonStar.Gend()
	KosonStar.Gend()
    }


    KosonStar.End()


    return KosonStar
}

func DrawKosonKanonText(Text string, StrokeWidthScalingFactor *p.Decimal, OutputVariable *os.File) {
    V1 := p.NFI(0)

    DrawWord(V1, p.NFI(0), Text, StrokeWidthScalingFactor, OutputVariable)


}

func DrawKosonStarText (Text string, StrokeWidthScalingFactor *p.Decimal, OutputVariable *os.File) {

    CorrectTabletSize := GetCorrectBoardSize(Text)
    TextLines := SplitString(CorrectTabletSize, Text)
    Offset := ComputeStartingOffset(CorrectTabletSize,TextLines)





    //Width := int(TabletSize.Lines[0]) * 200
    //Height := len(TabletSize.Lines) * 1600
    WLO := new(p.Decimal)
    WLOS := new(p.Decimal)
    //KosonText := svg.New(os.Stdout)

    //S1:="fill=\"none\""
    //S2:="stroke=\"black\""
    //S3:="stroke-width=\"7\""
    //S4:="stroke-width=\"20\""


    //KosonText.Start(Width,Height)

    //KosonText.Rect(0,0,Width,Height,S1,S2,S4)

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
		WLO = DrawWord(X, Y, WordsChainToPrint[j], StrokeWidthScalingFactor, OutputVariable)
		WLOS = b.ADDxc(WLO,p.NFI(1600))		//Offset given by space
	    } else {
		WLO = DrawWord(WLOS, Y, WordsChainToPrint[j], StrokeWidthScalingFactor, OutputVariable)
		WLOS = b.ADDxc(WLO,p.NFI(1600))		//Offset given by space
	    }
	}
    }
    //KosonText.End()
    return
}

func MakeRuneChain (Text string) []rune {
    Result := []rune(Text)
    return Result
}

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

//Doesnt Use Kerning
func GetTextLengthNoKerning (Text string) int64 {
    var Sum int64
    RuneChain := MakeRuneChain(Text)
    for i:=0; i<len(RuneChain); i++ {
	Length := GetRuneWidth(RuneChain[i])
	Sum = Sum + Length
    }
    return Sum
}

//Uses Kerning
func GetTextLengthWithKerning (Word string) int64 {
    var (
    	Sum int64
    	Length int64
    )
    RuneChain := MakeRuneChain(Word)
    for i:=0; i<len(RuneChain); i++ {
        if i == 0 {
	    Length = GetRuneWidth(RuneChain[i])
	    Sum = Sum + Length
	}else if i > 0 {
	    KerningValue := ComputeKerning(string(RuneChain[i-1]),string(RuneChain[i]))
	    Length = GetRuneWidth(RuneChain[i])
	    Sum = Sum + Length - KerningValue
	}
    }
    return Sum
}

func ComputeKerning (FirstGlyph, SecondGlyph string) (KerningValue int64) {
    var (
    	Pair = KerningPair {FirstGlyph,SecondGlyph}
    )
    for i:=0; i<len(KP2U); i++ {
        if Pair == KP2U[i] {
            KerningValue = 2
            break
	}
	// Else if for Other Kerning Pairs with value different from 2
    }

    return KerningValue
}

func GetRequiredBoardSize (Text string) (Result Board) {
    switch TextSize := GetTextLengthWithKerning(Text);{
    case TextSize < Tablet1600.Size:
    	Result = Tablet1600
    case TextSize >= Tablet1600.Size && TextSize < Tablet1200.Size:
    	Result = Tablet1200
    case TextSize >= Tablet1200.Size && TextSize < Tablet960.Size:
	Result = Tablet960
    case TextSize >= Tablet960.Size && TextSize < Tablet800.Size:
	Result = Tablet800
    case TextSize >= Tablet800.Size && TextSize < Tablet600.Size:
	Result = Tablet600
    case TextSize >= Tablet600.Size && TextSize < Tablet480.Size:
	Result = Tablet480
    case TextSize >= Tablet480.Size && TextSize < Tablet400.Size:
	Result = Tablet400
    case TextSize >= Tablet400.Size && TextSize < Tablet320.Size:
	Result = Tablet320
    case TextSize >= Tablet320.Size && TextSize < Tablet300.Size:
	Result = Tablet300
    case TextSize >= Tablet300.Size && TextSize < Tablet240.Size:
	Result = Tablet240
    case TextSize >= Tablet240.Size && TextSize < Tablet200.Size:
	Result = Tablet200
    case TextSize >= Tablet200.Size && TextSize < Tablet192.Size:
	Result = Tablet192
    case TextSize >= Tablet192.Size && TextSize < Tablet160.Size:
	Result = Tablet160
    case TextSize >= Tablet160.Size && TextSize < Tablet150.Size:
	Result = Tablet150
    case TextSize >= Tablet150.Size && TextSize < Tablet120.Size:
	Result = Tablet120
    case TextSize >= Tablet120.Size && TextSize < Tablet100.Size:
	Result = Tablet100
}
    return Result
}

func SplitString(Tablet Board, TextToSplit string) []string {
    var (
	Line string
	LineNumber int64
	StringList = make([]string, len(Tablet.Lines))
	StopSignal bool
	LengthToAdd int64
	LineLength int64
	CurrentBoardLine = int64(0)
	Stop int
    )
    WordsChain := strings.Split(TextToSplit," ")

    DeleteEmpty := func (s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
		r = append(r, str)
		}
	}
	return r
    }

    for j:=0; j<len(Tablet.Lines); j++ {
        if StopSignal == true {
            break
	}
	Line = ""
	LineLength = 0
	for i:=Stop; i<len(WordsChain); i++ {
	    if i == Stop {
		if i == len(WordsChain)-1 {
		    StopSignal = true
		}
		LengthToAdd = GetTextLengthWithKerning(WordsChain[i])
		LineLength = LineLength + LengthToAdd
		Line = Line + WordsChain[i]
		StringList[LineNumber] = Line


	    } else if i > Stop {
		LengthToAdd = GetTextLengthWithKerning(WordsChain[i])
		LineLength = LineLength + LengthToAdd + 8

		if LineLength > Tablet.Lines[CurrentBoardLine] {
		    Stop = i
		    StringList[LineNumber] = Line
		    LineNumber = LineNumber + 1
		    CurrentBoardLine = CurrentBoardLine + 1
		    break

		} else {
		    Line = Line + " " + WordsChain[i]
		    StringList[LineNumber] = Line

		}
		if i == len(WordsChain)-1 {
		    StopSignal = true
		}
	    }
	}
    }
    TrimmedList := DeleteEmpty(StringList)
    return TrimmedList
}

func ComputeStartingOffset (Tablet Board, TextLines []string) []int64 {
    var(
	OffsetList = make([]int64, len(TextLines))
	LineLength,TotalOffset int64
    )
    for i:=0; i<len(TextLines); i++ {
        LineLength = GetTextLengthWithKerning(TextLines[i])
        TotalOffset = Tablet.Lines[i] - LineLength

        if TotalOffset == Tablet.Lines[i] {
	    TotalOffset = -1
	}
        OffsetList[i] = TotalOffset
    }
    return OffsetList
}

//The Last Function that prints the Star
//Its called by calling :
// 	om.PrintStar(om.KossonSILVER660)
func PrintStar (Star Kanon) {
    var(
	UsedTabletSize Board
	Check bool
    )
    TabletSize := GetRequiredBoardSize(Star.Value)
    SmallerTabletSize := GetNextSmallerTablet(TabletSize)

    Lines := SplitString(TabletSize,Star.Value)
    Lines2 := SplitString(SmallerTabletSize,Star.Value)

    Offset := ComputeStartingOffset(TabletSize,Lines)
    if Offset[len(Offset)-1] < 0 || Lines[len(Lines)-1] != Lines2[len(Lines2)-1] {
	UsedTabletSize = SmallerTabletSize
	Check = true
    } else {
	UsedTabletSize = TabletSize
    }
    fmt.Println("UsedTabletSize is", UsedTabletSize.Name)
    Offset = ComputeStartingOffset(UsedTabletSize,Lines)

    if Check == true {
	for i:=0; i<len(Lines2); i++ {
	    if i < 10 {
		fmt.Println("LINE ",i,":: ",Lines2[i])
	    }else {
		fmt.Println("LINE",i,":: ",Lines2[i])
	    }
	}
    } else {
	for i:=0; i<len(Lines); i++ {
	    if i < 10 {
		fmt.Println("LINE ",i,":: ",Lines[i])
	    }else {
		fmt.Println("LINE",i,":: ",Lines[i])
	    }
	}
    }

    fmt.Println("Offset is",Offset)
    fmt.Println("Line Number", len(Lines))
    fmt.Println("====================")
    fmt.Println("")

    DrawKosonStar(Star)
}

//Generates the Empty Kanon String
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