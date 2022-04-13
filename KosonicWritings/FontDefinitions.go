package KosonicWritings

import (
    b "Cryptoplasm-Core/Cryptoplasm_Blockchain_Constants"
    p "github.com/Crypt0plasm/Firefly-APD"
    svg "github.com/ajstarks/svgo"
    "os"
)

var (
    WTier8 = []string{"A","Ä","B","C","D","E","F","G","H","I","J","K","L","M","N","O","Ö","P","Q","R","S","T","U","Ü","V","W","X","Y","Z"," "}
    WTier7 = []string{"a","ä","q","o","ö","u","ü","v"}
    WTier6 = []string{"c","d","g","w","m","z","t","x","y","¤"}
    WTier5 = []string{"0","1","2","3","4","5","6","7","8","9","h","n","k","£"}
    WTier4 = []string{"b","ß","e","f","p","r","s","$","l"}
    WTier3 = []string{"(",")","[","]","{","}"}
    WTier2 = []string{"i","j",".",",",":",";","!"}
    Capital = append(WTier8[:29])

)



type GlyphGraphic struct {
    Width int64
    SX *p.Decimal
    SY *p.Decimal
    Code string
}

type SVGStartCoord struct {
    X int
    Y int
}

//Coin Geometry Definitions
var (
    CenterBasorelief            = GlyphGraphic{18100,p.NFS("6850"),     p.NFS("3650"),          CreateSVGCode(ConvertRawToRelative(CenterBasoreliefRaw))}
    LeftMinorBasorelief         = GlyphGraphic{18100,p.NFS("7650"),     p.NFS("10250"),         CreateSVGCode(ConvertRawToRelative(LeftMinorBasoreliefRaw))}
    RightMinorBasorelief        = GlyphGraphic{18100,p.NFS("10450"),    p.NFS("10250"),         CreateSVGCode(ConvertRawToRelative(RightMinorBasoreliefRaw))}
    LeftMajorBasorelief         = GlyphGraphic{18100,p.NFS("6650"),     p.NFS("3650"),          CreateSVGCode(ConvertRawToRelative(LeftMajorBasoreliefRaw))}
    RightMajorBasorelief        = GlyphGraphic{18100,p.NFS("11450"),    p.NFS("3650"),          CreateSVGCode(ConvertRawToRelative(RightMajorBasoreliefRaw))}
    CrownBasorelief             = GlyphGraphic{18100,p.NFS("4524.5166"),p.NFS("3450"),          CreateSVGCode(ConvertRawToRelative(CrownBasoreliefRaw))}
    TextBasorelief              = GlyphGraphic{18100,p.NFS("5389.3990"),p.NFS("15250"),         CreateSVGCode(ConvertRawToRelative(TextBasoreliefRaw))}
    KanonBasorelief             = GlyphGraphic{18100,p.NFS("12348.4845"),p.NFS("15450"),        CreateSVGCode(ConvertRawToRelative(KanonBasoreliefRaw))}
)


//Glyph Graphic Definitions
var (
    //Tier 8
    CapitalLeftPillar           = GlyphGraphic{1600,p.NFS("560"),       p.NFS("0"),             CreateSVGCode(ConvertRawToRelative(CapitalLeftRawCoordinates))}
    CapitalRightPillar          = GlyphGraphic{1600,p.NFS("1040"),      p.NFS("0"),             CreateSVGCode(ConvertRawToRelative(CapitalRightRawCoordinates))}
    CapitalDoublePillar         = GlyphGraphic{1600,p.NFS("1040"),      p.NFS("0"),             CreateSVGCode(ConvertRawToRelative(CapitalDoubleRawCoordinates))}

    //Tier 7
    LatinSmallLetterA           = GlyphGraphic{1400,p.NFS("700"),       p.NFS("240"),           CreateSVGCode(ConvertRawToRelative(LatinSmallLetterARawCoordinates))}         // U+0061
    LatinSmallLetterQ           = GlyphGraphic{1400,p.NFS("700"),       p.NFS("800"),           CreateSVGCode(ConvertRawToRelative(LatinSmallLetterQRawCoordinates))}         // U+0071
    LatinSmallLetterO           = GlyphGraphic{1400,p.NFS("0"),         p.NFS("799.988"),       CreateSVGCode(ConvertRawToRelative(LatinSmallLetterORawCoordinates))}         // U+0071
    LatinSmallLetterU           = GlyphGraphic{1400,p.NFS("60"),        p.NFS("240"),           CreateSVGCode(ConvertRawToRelative(LatinSmallLetterURawCoordinates))}         // U+0075
    LatinSmallLetterV           = GlyphGraphic{1400,p.NFS("376.684"),   p.NFS("240"),           CreateSVGCode(ConvertRawToRelative(LatinSmallLetterVRawCoordinates))}         // U+0076


    //Tier 6
    LatinSmallLetterM           = GlyphGraphic{1200,p.NFS("20"),        p.NFS("240"),           CreateSVGCode(ConvertRawToRelative(LatinSmallLetterMRawCoordinates))}         //  U+006D
    LatinSmallLetterC           = GlyphGraphic{1200,p.NFS("1109.117"),  p.NFS("404.020"),       CreateSVGCode(ConvertRawToRelative(LatinSmallLetterCRawCoordinates))}         //  U+0063
    LatinSmallLetterD           = GlyphGraphic{1200,p.NFS("124.020"),   p.NFS("240"),           CreateSVGCode(ConvertRawToRelative(LatinSmallLetterDRawCoordinates))}         //  U+0064
    LatinSmallLetterG           = GlyphGraphic{1200,p.NFS("1109.117"),  p.NFS("404.020"),       CreateSVGCode(ConvertRawToRelative(LatinSmallLetterGRawCoordinates))}         //  U+0067
    LatinSmallLetterT           = GlyphGraphic{1200,p.NFS("40"),        p.NFS("160"),           CreateSVGCode(ConvertRawToRelative(LatinSmallLetterTRawCoordinates))}         // U+0074
    LatinSmallLetterW           = GlyphGraphic{1200,p.NFS("20"),        p.NFS("240"),           CreateSVGCode(ConvertRawToRelative(LatinSmallLetterWRawCoordinates))}         // U+0077
    LatinSmallLetterX           = GlyphGraphic{1200,p.NFS("294.558"),   p.NFS("240"),           CreateSVGCode(ConvertRawToRelative(LatinSmallLetterXRawCoordinates))}         // U+0078
    LatinSmallLetterY           = GlyphGraphic{1200,p.NFS("294.558"),   p.NFS("240"),           CreateSVGCode(ConvertRawToRelative(LatinSmallLetterYRawCoordinates))}         // U+0079
    LatinSmallLetterZ           = GlyphGraphic{1200,p.NFS("40"),        p.NFS("160"),           CreateSVGCode(ConvertRawToRelative(LatinSmallLetterZRawCoordinates))}         // U+007A

    //Tier 5
    LatinSmallLetterH           = GlyphGraphic{1000,p.NFS("20"),        p.NFS("240"),           CreateSVGCode(ConvertRawToRelative(LatinSmallLetterHRawCoordinates))}         //  U+0068
    LatinSmallLetterK           = GlyphGraphic{1000,p.NFS("20"),        p.NFS("240"),           CreateSVGCode(ConvertRawToRelative(LatinSmallLetterKRawCoordinates))}         //  U+006B
    LatinSmallLetterN           = GlyphGraphic{1000,p.NFS("20"),        p.NFS("240"),           CreateSVGCode(ConvertRawToRelative(LatinSmallLetterNRawCoordinates))}         //  U+006E
    HyphenMinus                 = GlyphGraphic{1000,p.NFS("20"),        p.NFS("620"),           CreateSVGCode(ConvertRawToRelative(HyphenMinusRawCoordinates))}                //  U+002D
    PlusSign                    = GlyphGraphic{1000,p.NFS("20"),        p.NFS("620"),           CreateSVGCode(ConvertRawToRelative(PlusSignRawCoordinates))}                    //  U+002B
    //Numbers
    DigitZero                   = GlyphGraphic{1000,p.NFS("994"),       p.NFS("544"),           CreateSVGCode(ConvertRawToRelative(DigitZeroRawCoordinates))}                 //  U+0030
    DigitZeroInner              = GlyphGraphic{1000,p.NFS("372"),       p.NFS("544"),           CreateSVGCode(ConvertRawToRelative(DigitZeroInnerRawCoordinates))}            //  U+0030
    DigitOne                    = GlyphGraphic{1000,p.NFS("500"),       p.NFS("160"),           CreateSVGCode(ConvertRawToRelative(DigitOneRawCoordinates))}                  //  U+0031
    DigitTwo                    = GlyphGraphic{1000,p.NFS("466.871"),   p.NFS("667.639"),       CreateSVGCode(ConvertRawToRelative(DigitTwoRawCoordinates))}                  //  U+0032
    DigitThree                  = GlyphGraphic{1000,p.NFS("116"),       p.NFS("160"),           CreateSVGCode(ConvertRawToRelative(DigitThreeRawCoordinates))}                  //  U+0033
    DigitFour                   = GlyphGraphic{1000,p.NFS("756"),       p.NFS("160"),           CreateSVGCode(ConvertRawToRelative(DigitFourRawCoordinates))}                  //  U+0034
    DigitFive                   = GlyphGraphic{1000,p.NFS("884"),       p.NFS("160"),           CreateSVGCode(ConvertRawToRelative(DigitFiveRawCoordinates))}                  //  U+0035
    DigitSix                    = GlyphGraphic{1000,p.NFS("322.008"),   p.NFS("116.822"),       CreateSVGCode(ConvertRawToRelative(DigitSixRawCoordinates))}                  //  U+0036
    DigitSeven                  = GlyphGraphic{1000,p.NFS("116"),       p.NFS("160"),           CreateSVGCode(ConvertRawToRelative(DigitSevenRawCoordinates))}                  //  U+0037
    DigitEight                  = GlyphGraphic{1000,p.NFS("500"),       p.NFS("26.660"),        CreateSVGCodeNoZ(ConvertRawToRelative(DigitEightRawCoordinates))}                  //  U+0038
    DigitNine                   = GlyphGraphic{1000,p.NFS("536.685"),   p.NFS("626.742"),       CreateSVGCode(ConvertRawToRelative(DigitNineRawCoordinates))}                  //  U+0039



    //Tier 4
    LatinSmallLetterB           = GlyphGraphic{800, p.NFS("20"),        p.NFS("240"),           CreateSVGCode(ConvertRawToRelative(LatinSmallLetterBRawCoordinates))}         //  U+0062
    LatinSmallLetterSharpS      = GlyphGraphic{800, p.NFS("20"),        p.NFS("240"),           CreateSVGCode(ConvertRawToRelative(LatinSmallLetterSharpSRawCoordinates))}    //  U+00DF
    LatinSmallLetterE           = GlyphGraphic{800, p.NFS("100"),       p.NFS("240"),           CreateSVGCode(ConvertRawToRelative(LatinSmallLetterERawCoordinates))}         //  U+0065
    LatinSmallLetterF           = GlyphGraphic{800, p.NFS("100"),       p.NFS("240"),           CreateSVGCode(ConvertRawToRelative(LatinSmallLetterFRawCoordinates))}         //  U+0066
    LatinSmallLetterL           = GlyphGraphic{800, p.NFS("20"),        p.NFS("240"),           CreateSVGCode(ConvertRawToRelative(LatinSmallLetterLRawCoordinates))}        //  U+006C
    LatinSmallLetterP           = GlyphGraphic{800, p.NFS("20"),        p.NFS("240"),           CreateSVGCode(ConvertRawToRelative(LatinSmallLetterPRawCoordinates))}        //  U+0070
    LatinSmallLetterR           = GlyphGraphic{800, p.NFS("20"),        p.NFS("240"),           CreateSVGCode(ConvertRawToRelative(LatinSmallLetterRRawCoordinates))}        //  U+0072
    LatinSmallLetterS           = GlyphGraphic{800, p.NFS("730"),       p.NFS("160"),           CreateSVGCode(ConvertRawToRelative(LatinSmallLetterSRawCoordinates))}        // U+0073

    //Tier 2
    LatinSmallLetterI           = GlyphGraphic{400, p.NFS("20"),        p.NFS("240"),           CreateSVGCode(ConvertRawToRelative(LatinSmallLetterIRawCoordinates))}         //  U+0069
    LatinSmallLetterJ           = GlyphGraphic{400, p.NFS("20"),        p.NFS("240"),           CreateSVGCode(ConvertRawToRelative(LatinSmallLetterJRawCoordinates))}         //  U+006A

    //Tier 0 - Diacritics and Orthography
    DieresisLeft                = GlyphGraphic{-1,  p.NFS("-230"),      p.NFS("10"),            CreateSVGCode(ConvertRawToRelative(DieresisLeftRawCoordinates))}
    DieresisRight               = GlyphGraphic{-1,  p.NFS("230"),       p.NFS("10"),            CreateSVGCode(ConvertRawToRelative(DieresisRightRawCoordinates))}
    Comma                       = GlyphGraphic{400, p.NFS("200"),       p.NFS("1060"),          CreateSVGCode(ConvertRawToRelative(CommaRawCoordinates))}
    PointDown                   = GlyphGraphic{400, p.NFS("200"),       p.NFS("1060"),          CreateSVGCode(ConvertRawToRelative(PointDownRawCoordinates))}
    PointUp                     = GlyphGraphic{400, p.NFS("200"),       p.NFS("240"),           CreateSVGCode(ConvertRawToRelative(PointUpRawCoordinates))}
    LeftParenthesis             = GlyphGraphic{600, p.NFS("600"),       p.NFS("0"),             CreateSVGCode(ConvertRawToRelative(LeftParenthesisRawCoordinates))}
    RightParenthesis            = GlyphGraphic{600, p.NFS("0"),         p.NFS("0"),             CreateSVGCode(ConvertRawToRelative(RightParenthesisRawCoordinates))}
    LeftDoubleAngleBracketOne   = GlyphGraphic{1200,p.NFS("40"),        p.NFS("800"),           CreateSVGCode(ConvertRawToRelative(LeftDoubleAngleBracketOneRawCoordinates))}
    LeftDoubleAngleBracketTwo   = GlyphGraphic{1200,p.NFS("464.264"),   p.NFS("800"),           CreateSVGCode(ConvertRawToRelative(LeftDoubleAngleBracketTwoRawCoordinates))}
    RightDoubleAngleBracketOne  = GlyphGraphic{1200,p.NFS("735.736"),   p.NFS("800"),           CreateSVGCode(ConvertRawToRelative(RightDoubleAngleBracketOneRawCoordinates))}
    RightDoubleAngleBracketTwo  = GlyphGraphic{1200,p.NFS("1160"),      p.NFS("800"),           CreateSVGCode(ConvertRawToRelative(RightDoubleAngleBracketTwoRawCoordinates))}


    //Omfalon Murado
    OmPoint                     = GlyphGraphic{1600,p.NFS("1066.038"),          p.NFS("140.095"),             CreateSVGCode(ConvertRawToRelative(OMPointRC))}
    OmPoint2                    = GlyphGraphic{1600,p.NFS("1391.157"),          p.NFS("252.850"),             CreateSVGCode(ConvertRawToRelative(OMPoint2RC))}
    OmMain                      = GlyphGraphic{1600,p.NFS("1422.355"),          p.NFS("521.564"),             CreateSVGCode(ConvertRawToRelative(OMMainRC))}

)

//Letter Functions
//True is used for capital, False for non Capital
func ComputeGlyphCode(Glyph GlyphGraphic, StartX, StartY *p.Decimal, Capital bool) string {
    var Result string
    if Capital == false {
        StartingX := b.ADDxc(StartX,Glyph.SX).String()
        StartingY := b.ADDxc(StartY,Glyph.SY).String()
        Head := "M" + StartingX + "," + StartingY + " "
        Result = Head + Glyph.Code
    } else {
        V1:= b.SUBxc(p.NFI(1600),p.NFI(Glyph.Width))
        V2:= b.DIVxc(V1,p.NFI(2))
        StartingX := b.SUMxc(StartX,Glyph.SX,V2).String()
        StartingY := b.ADDxc(StartY,Glyph.SY).String()
        Head := "M" + StartingX + "," + StartingY + " "
        Result = Head + Glyph.Code
    }

    return Result
}

func DrawCapitalGlyphAt (X,Y *p.Decimal, Glyph string, StrokeWidthScalingFactor *p.Decimal, Type int, OutputVariable *os.File) (OutputGlyph *svg.SVG, MovementDistance *p.Decimal) {
    var (
        S1="fill=\"none\""
        S2="stroke=\"black\""
        StrokeWidth = b.DIVxc(p.NFS("3"),StrokeWidthScalingFactor).String()
        S3="stroke-width=\"" + StrokeWidth + "\""

        Canvas             = svg.New(OutputVariable)
    )

    Switcheroo := func () {
        switch Letter := Glyph;{
        case Letter == "A":
            Canvas.Path(ComputeGlyphCode(LatinSmallLetterA,X,Y,true),S1,S2,S3)
        case Letter == "Ä":
            Canvas.Path(ComputeGlyphCode(LatinSmallLetterA,X,Y,true),S1,S2,S3)
            Canvas.Path(ComputeGlyphCode(DieresisLeft,b.ADDxc(X,p.NFS("800")),Y,false),S1,S2,S3)
            Canvas.Path(ComputeGlyphCode(DieresisRight,b.ADDxc(X,p.NFS("800")),Y,false),S1,S2,S3)
        case Letter == "B":
            Canvas.Path(ComputeGlyphCode(LatinSmallLetterB,X,Y,true),S1,S2,S3)
        case Letter == "C":
            Canvas.Path(ComputeGlyphCode(LatinSmallLetterC,X,Y,true),S1,S2,S3)
        case Letter == "D":
            Canvas.Path(ComputeGlyphCode(LatinSmallLetterD,X,Y,true),S1,S2,S3)
        case Letter == "E":
            Canvas.Path(ComputeGlyphCode(LatinSmallLetterE,X,Y,true),S1,S2,S3)
        case Letter == "F":
            Canvas.Path(ComputeGlyphCode(LatinSmallLetterF,X,Y,true),S1,S2,S3)
        case Letter == "G":
            Canvas.Path(ComputeGlyphCode(LatinSmallLetterG,X,Y,true),S1,S2,S3)
        case Letter == "H":
            Canvas.Path(ComputeGlyphCode(LatinSmallLetterH,X,Y,true),S1,S2,S3)
        case Letter == "I":
            Canvas.Path(ComputeGlyphCode(LatinSmallLetterI,X,Y,true),S1,S2,S3)
        case Letter == "J":
            Canvas.Path(ComputeGlyphCode(LatinSmallLetterJ,X,Y,true),S1,S2,S3)
        case Letter == "K":
            Canvas.Path(ComputeGlyphCode(LatinSmallLetterK,X,Y,true),S1,S2,S3)
        case Letter == "L":
            Canvas.Path(ComputeGlyphCode(LatinSmallLetterL,X,Y,true),S1,S2,S3)
        case Letter == "M":
            Canvas.Path(ComputeGlyphCode(LatinSmallLetterM,X,Y,true),S1,S2,S3)
        case Letter == "N":
            Canvas.Path(ComputeGlyphCode(LatinSmallLetterN,X,Y,true),S1,S2,S3)
        case Letter == "O":
            Canvas.Path(ComputeGlyphCode(LatinSmallLetterO,X,Y,true),S1,S2,S3)
        case Letter == "Ö":
            Canvas.Path(ComputeGlyphCode(LatinSmallLetterO,X,Y,true),S1,S2,S3)
            Canvas.Path(ComputeGlyphCode(DieresisLeft,b.ADDxc(X,p.NFS("800")),Y,false),S1,S2,S3)
            Canvas.Path(ComputeGlyphCode(DieresisRight,b.ADDxc(X,p.NFS("800")),Y,false),S1,S2,S3)
        case Letter == "P":
            Canvas.Path(ComputeGlyphCode(LatinSmallLetterP,X,Y,true),S1,S2,S3)
        case Letter == "Q":
            Canvas.Path(ComputeGlyphCode(LatinSmallLetterQ,X,Y,true),S1,S2,S3)
        case Letter == "R":
            Canvas.Path(ComputeGlyphCode(LatinSmallLetterR,X,Y,true),S1,S2,S3)
        case Letter == "S":
            Canvas.Path(ComputeGlyphCode(LatinSmallLetterS,X,Y,true),S1,S2,S3)
        case Letter == "T":
            Canvas.Path(ComputeGlyphCode(LatinSmallLetterT,X,Y,true),S1,S2,S3)
        case Letter == "U":
            Canvas.Path(ComputeGlyphCode(LatinSmallLetterU,X,Y,true),S1,S2,S3)
        case Letter == "Ü":
            Canvas.Path(ComputeGlyphCode(LatinSmallLetterU,X,Y,true),S1,S2,S3)
            Canvas.Path(ComputeGlyphCode(DieresisLeft,b.ADDxc(X,p.NFS("800")),Y,false),S1,S2,S3)
            Canvas.Path(ComputeGlyphCode(DieresisRight,b.ADDxc(X,p.NFS("800")),Y,false),S1,S2,S3)
        case Letter == "V":
            Canvas.Path(ComputeGlyphCode(LatinSmallLetterV,X,Y,true),S1,S2,S3)
        case Letter == "W":
            Canvas.Path(ComputeGlyphCode(LatinSmallLetterW,X,Y,true),S1,S2,S3)
        case Letter == "X":
            Canvas.Path(ComputeGlyphCode(LatinSmallLetterX,X,Y,true),S1,S2,S3)
        case Letter == "Y":
            Canvas.Path(ComputeGlyphCode(LatinSmallLetterY,X,Y,true),S1,S2,S3)
        case Letter == "Z":
            Canvas.Path(ComputeGlyphCode(LatinSmallLetterZ,X,Y,true),S1,S2,S3)
        }
    }

    if Type == 0 {
        Canvas.Path(ComputeGlyphCode(CapitalLeftPillar,X,Y,false),S1,S2,S3)
        Switcheroo()
        Canvas.Path(ComputeGlyphCode(CapitalRightPillar,X,Y,false),S1,S2,S3)
    } else if Type == 1 {
        Canvas.Path(ComputeGlyphCode(CapitalLeftPillar,X,Y,false),S1,S2,S3)
        Switcheroo()
        Canvas.Path(ComputeGlyphCode(CapitalDoublePillar,X,Y,false),S1,S2,S3)
    } else if Type == 2 {
        Switcheroo()
        Canvas.Path(ComputeGlyphCode(CapitalRightPillar,X,Y,false),S1,S2,S3)
    } else if Type == 3 {
        Switcheroo()
        Canvas.Path(ComputeGlyphCode(CapitalDoublePillar,X,Y,false),S1,S2,S3)
    }
    OutputGlyph = Canvas
    MovementDistance = p.NFS("1600")
    return OutputGlyph,MovementDistance
}

func DrawGlyphAt (X,Y *p.Decimal, Glyph string,  StrokeWidthScalingFactor *p.Decimal, OutputVariable *os.File) (OutputGlyph *svg.SVG, MovementDistance *p.Decimal) {
    var (
        //Width int
        //Height = 1600
        S1="fill=\"none\""
        S2="stroke=\"black\""

        StrokeWidth = b.DIVxc(p.NFS("3"),StrokeWidthScalingFactor).String()
        S3="stroke-width=\"" + StrokeWidth + "\""
        Canvas             = svg.New(OutputVariable)
        
    )

    Switcheroo := func () {
        switch Letter := Glyph; {
        case Letter == " ":
            MovementDistance = p.NFI(800)
        case Letter == "-":
            Canvas.Path(ComputeGlyphCode(HyphenMinus, X, Y, false), S1, S2, S3)
            MovementDistance = p.NFI(HyphenMinus.Width)
        case Letter == "+":
            Canvas.Path(ComputeGlyphCode(PlusSign, X, Y, false), S1, S2, S3)
            MovementDistance = p.NFI(PlusSign.Width)
        case Letter == "(":
            Canvas.Path(ComputeGlyphCode(LeftParenthesis, X, Y, false), S1, S2, S3)
            MovementDistance = p.NFI(LeftParenthesis.Width)
        case Letter == ")":
            Canvas.Path(ComputeGlyphCode(RightParenthesis, X, Y, false), S1, S2, S3)
            MovementDistance = p.NFI(RightParenthesis.Width)
        case Letter == "《" || Letter == "«":
            Canvas.Path(ComputeGlyphCode(LeftDoubleAngleBracketOne, X, Y, false), S1, S2, S3)
            Canvas.Path(ComputeGlyphCode(LeftDoubleAngleBracketTwo, X, Y, false), S1, S2, S3)
            MovementDistance = p.NFI(LeftDoubleAngleBracketOne.Width)
        case Letter == "》" || Letter == "»":
            Canvas.Path(ComputeGlyphCode(RightDoubleAngleBracketOne, X, Y, false), S1, S2, S3)
            Canvas.Path(ComputeGlyphCode(RightDoubleAngleBracketTwo, X, Y, false), S1, S2, S3)
            MovementDistance = p.NFI(RightDoubleAngleBracketOne.Width)
        case Letter == "0":
            Canvas.Path(ComputeGlyphCode(DigitZero, X, Y, false), S1, S2, S3)
            Canvas.Path(ComputeGlyphCode(DigitZeroInner, X, Y, false), S1, S2, S3)
            MovementDistance = p.NFI(DigitZero.Width)
        case Letter == "1":
            Canvas.Path(ComputeGlyphCode(DigitOne, X, Y, false), S1, S2, S3)
            MovementDistance = p.NFI(DigitOne.Width)
        case Letter == "2":
            Canvas.Path(ComputeGlyphCode(DigitTwo, X, Y, false), S1, S2, S3)
            MovementDistance = p.NFI(DigitTwo.Width)
        case Letter == "3":
            Canvas.Path(ComputeGlyphCode(DigitThree, X, Y, false), S1, S2, S3)
            MovementDistance = p.NFI(DigitThree.Width)
        case Letter == "4":
            Canvas.Path(ComputeGlyphCode(DigitFour, X, Y, false), S1, S2, S3)
            MovementDistance = p.NFI(DigitFour.Width)
        case Letter == "5":
            Canvas.Path(ComputeGlyphCode(DigitFive, X, Y, false), S1, S2, S3)
            MovementDistance = p.NFI(DigitFive.Width)
        case Letter == "6":
            Canvas.Path(ComputeGlyphCode(DigitSix, X, Y, false), S1, S2, S3)
            MovementDistance = p.NFI(DigitSix.Width)
        case Letter == "7":
            Canvas.Path(ComputeGlyphCode(DigitSeven, X, Y, false), S1, S2, S3)
            MovementDistance = p.NFI(DigitSeven.Width)
        case Letter == "8":
            Canvas.Path(ComputeGlyphCode(DigitEight, X, Y, false), S1, S2, S3)
            C1X, _ := b.ADDxc(X,p.NFI(500)).Int64()
            C1Y, _ := b.ADDxc(Y,p.NFI(544)).Int64()
            C2X, _ := b.ADDxc(X,p.NFI(500)).Int64()
            C2Y, _ := b.ADDxc(Y,p.NFI(1056)).Int64()
            Canvas.Circle(int(C1X),int(C1Y),128, S1, S2, S3)
            Canvas.Circle(int(C2X),int(C2Y),128, S1, S2, S3)
            MovementDistance = p.NFI(DigitEight.Width)
        case Letter == "9":
            Canvas.Path(ComputeGlyphCode(DigitNine, X, Y, false), S1, S2, S3)
            MovementDistance = p.NFI(DigitNine.Width)
        case Letter == "a":
            Canvas.Path(ComputeGlyphCode(LatinSmallLetterA, X, Y, false), S1, S2, S3)
            MovementDistance = p.NFI(LatinSmallLetterA.Width)
        case Letter == "ä":
            Canvas.Path(ComputeGlyphCode(LatinSmallLetterA, X, Y, false), S1, S2, S3)
            Canvas.Path(ComputeGlyphCode(DieresisLeft,b.ADDxc(X,b.DIVxc(p.NFI(LatinSmallLetterA.Width),p.NFS("2"))),Y,false),S1,S2,S3)
            Canvas.Path(ComputeGlyphCode(DieresisRight,b.ADDxc(X,b.DIVxc(p.NFI(LatinSmallLetterA.Width),p.NFS("2"))),Y,false),S1,S2,S3)
            MovementDistance = p.NFI(LatinSmallLetterA.Width)
        case Letter == "b":
            Canvas.Path(ComputeGlyphCode(LatinSmallLetterB, X, Y, false), S1, S2, S3)
            MovementDistance = p.NFI(LatinSmallLetterB.Width)
        case Letter == "ß":
            Canvas.Path(ComputeGlyphCode(LatinSmallLetterSharpS, X, Y, false), S1, S2, S3)
            MovementDistance = p.NFI(LatinSmallLetterSharpS.Width)
        case Letter == "c":
            Canvas.Path(ComputeGlyphCode(LatinSmallLetterC, X, Y, false), S1, S2, S3)
            MovementDistance = p.NFI(LatinSmallLetterC.Width)
        case Letter == "d":
            Canvas.Path(ComputeGlyphCode(LatinSmallLetterD, X, Y, false), S1, S2, S3)
            MovementDistance = p.NFI(LatinSmallLetterD.Width)
        case Letter == "e":
            Canvas.Path(ComputeGlyphCode(LatinSmallLetterE, X, Y, false), S1, S2, S3)
            MovementDistance = p.NFI(LatinSmallLetterE.Width)
        case Letter == "f":
            Canvas.Path(ComputeGlyphCode(LatinSmallLetterF, X, Y, false), S1, S2, S3)
            MovementDistance = p.NFI(LatinSmallLetterF.Width)
        case Letter == "g":
            Canvas.Path(ComputeGlyphCode(LatinSmallLetterG, X, Y, false), S1, S2, S3)
            MovementDistance = p.NFI(LatinSmallLetterG.Width)
        case Letter == "h":
            Canvas.Path(ComputeGlyphCode(LatinSmallLetterH, X, Y, false), S1, S2, S3)
            MovementDistance = p.NFI(LatinSmallLetterH.Width)
        case Letter == "i":
            Canvas.Path(ComputeGlyphCode(LatinSmallLetterI, X, Y, false), S1, S2, S3)
            MovementDistance = p.NFI(LatinSmallLetterI.Width)
        case Letter == "j":
            Canvas.Path(ComputeGlyphCode(LatinSmallLetterJ, X, Y, false), S1, S2, S3)
            MovementDistance = p.NFI(LatinSmallLetterJ.Width)
        case Letter == "k":
            Canvas.Path(ComputeGlyphCode(LatinSmallLetterK, X, Y, false), S1, S2, S3)
            MovementDistance = p.NFI(LatinSmallLetterK.Width)
        case Letter == "l":
            Canvas.Path(ComputeGlyphCode(LatinSmallLetterL, X, Y, false), S1, S2, S3)
            MovementDistance = p.NFI(LatinSmallLetterL.Width)
        case Letter == "m":
            Canvas.Path(ComputeGlyphCode(LatinSmallLetterM, X, Y, false), S1, S2, S3)
            MovementDistance = p.NFI(LatinSmallLetterM.Width)
        case Letter == "n":
            Canvas.Path(ComputeGlyphCode(LatinSmallLetterN, X, Y, false), S1, S2, S3)
            MovementDistance = p.NFI(LatinSmallLetterN.Width)
        case Letter == "o":
            Canvas.Path(ComputeGlyphCode(LatinSmallLetterO, X, Y, false), S1, S2, S3)
            MovementDistance = p.NFI(LatinSmallLetterO.Width)
        case Letter == "ö":
            Canvas.Path(ComputeGlyphCode(LatinSmallLetterO, X, Y, false), S1, S2, S3)
            Canvas.Path(ComputeGlyphCode(DieresisLeft,b.ADDxc(X,b.DIVxc(p.NFI(LatinSmallLetterO.Width),p.NFS("2"))),Y,false),S1,S2,S3)
            Canvas.Path(ComputeGlyphCode(DieresisRight,b.ADDxc(X,b.DIVxc(p.NFI(LatinSmallLetterO.Width),p.NFS("2"))),Y,false),S1,S2,S3)
            MovementDistance = p.NFI(LatinSmallLetterO.Width)
        case Letter == "p":
            Canvas.Path(ComputeGlyphCode(LatinSmallLetterP, X, Y, false), S1, S2, S3)
            MovementDistance = p.NFI(LatinSmallLetterP.Width)
        case Letter == "q":
            Canvas.Path(ComputeGlyphCode(LatinSmallLetterQ, X, Y, false), S1, S2, S3)
            MovementDistance = p.NFI(LatinSmallLetterQ.Width)
        case Letter == "r":
            Canvas.Path(ComputeGlyphCode(LatinSmallLetterR, X, Y, false), S1, S2, S3)
            MovementDistance = p.NFI(LatinSmallLetterR.Width)
        case Letter == "s":
            Canvas.Path(ComputeGlyphCode(LatinSmallLetterS, X, Y, false), S1, S2, S3)
            MovementDistance = p.NFI(LatinSmallLetterS.Width)
        case Letter == "t":
            Canvas.Path(ComputeGlyphCode(LatinSmallLetterT, X, Y, false), S1, S2, S3)
            MovementDistance = p.NFI(LatinSmallLetterT.Width)
        case Letter == "u":
            Canvas.Path(ComputeGlyphCode(LatinSmallLetterU, X, Y, false), S1, S2, S3)
            MovementDistance = p.NFI(LatinSmallLetterU.Width)
        case Letter == "ü":
            Canvas.Path(ComputeGlyphCode(LatinSmallLetterU, X, Y, false), S1, S2, S3)
            Canvas.Path(ComputeGlyphCode(DieresisLeft,b.ADDxc(X,b.DIVxc(p.NFI(LatinSmallLetterU.Width),p.NFS("2"))),Y,false),S1,S2,S3)
            Canvas.Path(ComputeGlyphCode(DieresisRight,b.ADDxc(X,b.DIVxc(p.NFI(LatinSmallLetterU.Width),p.NFS("2"))),Y,false),S1,S2,S3)
            MovementDistance = p.NFI(LatinSmallLetterU.Width)
        case Letter == "v":
            Canvas.Path(ComputeGlyphCode(LatinSmallLetterV, X, Y, false), S1, S2, S3)
            MovementDistance = p.NFI(LatinSmallLetterV.Width)
        case Letter == "w":
            Canvas.Path(ComputeGlyphCode(LatinSmallLetterW, X, Y, false), S1, S2, S3)
            MovementDistance = p.NFI(LatinSmallLetterW.Width)
        case Letter == "x":
            Canvas.Path(ComputeGlyphCode(LatinSmallLetterX, X, Y, false), S1, S2, S3)
            MovementDistance = p.NFI(LatinSmallLetterX.Width)
        case Letter == "y":
            Canvas.Path(ComputeGlyphCode(LatinSmallLetterY, X, Y, false), S1, S2, S3)
            MovementDistance = p.NFI(LatinSmallLetterY.Width)
        case Letter == "z":
            Canvas.Path(ComputeGlyphCode(LatinSmallLetterZ, X, Y, false), S1, S2, S3)
            MovementDistance = p.NFI(LatinSmallLetterZ.Width)
        case Letter == ".":
            Canvas.Path(ComputeGlyphCode(PointDown, X, Y, false), S1, S2, S3)
            MovementDistance = p.NFI(PointDown.Width)
        case Letter == ",":
            Canvas.Path(ComputeGlyphCode(Comma, X, Y, false), S1, S2, S3)
            MovementDistance = p.NFI(Comma.Width)
        case Letter == ":":
            Canvas.Path(ComputeGlyphCode(PointDown, X, Y, false), S1, S2, S3)
            Canvas.Path(ComputeGlyphCode(PointUp, X, Y, false), S1, S2, S3)
            MovementDistance = p.NFI(PointDown.Width)
        case Letter == ";":
            Canvas.Path(ComputeGlyphCode(Comma, X, Y, false), S1, S2, S3)
            Canvas.Path(ComputeGlyphCode(PointUp, X, Y, false), S1, S2, S3)
            MovementDistance = p.NFI(PointDown.Width)
        }
    }

    Switcheroo()
    OutputGlyph = Canvas
    return OutputGlyph,MovementDistance
}

func IzCapital (Glyph rune) bool {
    var Output bool
    for i:=0; i<len(Capital); i++ {
        if string(Glyph) == Capital[i] {
            Output = true
            break
        }else {
            Output = false
        }
    }
    return Output
}

func DrawOM (X,Y *p.Decimal, StrokeWidthScalingFactor *p.Decimal, OutputVariable *os.File) {
    var(
        S1="fill=\"none\""
        S2="stroke=\"black\""

        StrokeWidth = b.DIVxc(p.NFS("3"),StrokeWidthScalingFactor).String()
        S3="stroke-width=\"" + StrokeWidth + "\""
        Canvas             = svg.New(OutputVariable)
    )
    //EllipseCenterX, _ := b.ADDxc(X,p.NFS("1006.4838")).Int64()
    //EllipseCenterY, _ := b.ADDxc(Y,p.NFS("212.3908")).Int64()
    //EllipseMinorAx, _ := p.NFS("56.5243").Int64()
    //EllipseMajorAx, _ := p.NFS("93.6657").Int64()


    Canvas.Path(ComputeGlyphCode(OmPoint,X,Y,false),S1,S2,S3)
    Canvas.Path(ComputeGlyphCode(OmPoint2,X,Y,false),S1,S2,S3)
    Canvas.Path(ComputeGlyphCode(OmMain,X,Y,false),S1,S2,S3)
}

func DrawWord(X,Y *p.Decimal, Word string, StrokeWidthScalingFactor *p.Decimal, OutputVariable *os.File) *p.Decimal {
    var(
        //Width = int(GetTextLengthWithKerning(Word)) * 200
        //Height = 1600

        //OutputGlyph = svg.New(os.Stdout)
        //TotalWord = svg.New(os.Stdout)
        KerningValue int64
        MovementDistance = new(p.Decimal)
        DrawingPointX = new(p.Decimal)
        DrawingPointY = new(p.Decimal)


        //S1="fill=\"none\""
        //S2="stroke=\"black\""
        //S3="stroke-width=\"1\""
    )
    DrawingPointX = X
    DrawingPointY = Y
    //TotalWord.Start(Width, Height)
    RuneChain := MakeRuneChain(Word)
    for i:=0; i<len(RuneChain); i++ {

        //First and only Position

        if i==0 && len(RuneChain) == 1 {
            //fmt.Println("First and only Position")
            if IzCapital(RuneChain[i]) == true {
                _, MovementDistance = DrawCapitalGlyphAt(DrawingPointX, DrawingPointY, string(RuneChain[i]), StrokeWidthScalingFactor, 0,OutputVariable)
            } else {
                _, MovementDistance = DrawGlyphAt(DrawingPointX, DrawingPointY, string(RuneChain[i]), StrokeWidthScalingFactor, OutputVariable)
            }

        //First Position of many

        }else if i==0 {
            //fmt.Println("First Position of many")
            if IzCapital(RuneChain[i]) == true {
                if IzCapital(RuneChain[i+1]) == true {
                    _, MovementDistance = DrawCapitalGlyphAt(DrawingPointX, DrawingPointY, string(RuneChain[i]), StrokeWidthScalingFactor,1, OutputVariable)
                } else {
                    _, MovementDistance = DrawCapitalGlyphAt(DrawingPointX, DrawingPointY, string(RuneChain[i]), StrokeWidthScalingFactor,0, OutputVariable)
                }
            } else {
                //Kerning goes here
                _, MovementDistance = DrawGlyphAt(DrawingPointX, DrawingPointY, string(RuneChain[i]), StrokeWidthScalingFactor, OutputVariable)
                KerningValue = ComputeKerning(string(RuneChain[i]),string(RuneChain[i+1]))
                DrawingPointX = b.SUBxc(DrawingPointX,b.MULxc(p.NFI(KerningValue),p.NFS("200")))
            }
        //Intermediary Positions
        } else if i<len(RuneChain)-1 && i != 0 && i != len(RuneChain){
            //fmt.Println("Intermediary Positions")
            if IzCapital(RuneChain[i]) == true {
                if IzCapital(RuneChain[i-1]) == true {
                    if IzCapital(RuneChain[i+1]) == true {
                        _, MovementDistance = DrawCapitalGlyphAt(DrawingPointX, DrawingPointY,string(RuneChain[i]), StrokeWidthScalingFactor,3, OutputVariable)
                    } else {
                        _, MovementDistance = DrawCapitalGlyphAt(DrawingPointX, DrawingPointY,string(RuneChain[i]), StrokeWidthScalingFactor,2, OutputVariable)
                    }
                } else {
                    if IzCapital(RuneChain[i+1]) == true {
                        _, MovementDistance = DrawCapitalGlyphAt(DrawingPointX, DrawingPointY, string(RuneChain[i]), StrokeWidthScalingFactor,1, OutputVariable)
                    } else {
                        _, MovementDistance = DrawCapitalGlyphAt(DrawingPointX, DrawingPointY, string(RuneChain[i]), StrokeWidthScalingFactor,0,OutputVariable)
                    }
                }

            }  else {
                //Kerning goes here
                _, MovementDistance = DrawGlyphAt(DrawingPointX, DrawingPointY, string(RuneChain[i]), StrokeWidthScalingFactor, OutputVariable)
                KerningValue = ComputeKerning(string(RuneChain[i]),string(RuneChain[i+1]))
                DrawingPointX = b.SUBxc(DrawingPointX,b.MULxc(p.NFI(KerningValue),p.NFS("200")))
            }

        //Last Position
        } else if i == len(RuneChain)-1 {
            //fmt.Println("Last Position")
            if IzCapital(RuneChain[i]) == true {
                if IzCapital(RuneChain[i-1]) == true {
                    _, MovementDistance = DrawCapitalGlyphAt(DrawingPointX, DrawingPointY, string(RuneChain[i]), StrokeWidthScalingFactor,2, OutputVariable)
                } else {
                    _, MovementDistance = DrawCapitalGlyphAt(DrawingPointX, DrawingPointY, string(RuneChain[i]), StrokeWidthScalingFactor,0, OutputVariable)
                }
            } else {
                _, MovementDistance = DrawGlyphAt(DrawingPointX, DrawingPointY, string(RuneChain[i]), StrokeWidthScalingFactor, OutputVariable)
            }
        }

        DrawingPointX = b.ADDxc(DrawingPointX,MovementDistance)
    }
    //TotalWord.End()

    return DrawingPointX


}
