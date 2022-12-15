package KosonicWritings

import (
    b "Cryptoplasm-Core/Cryptoplasm_Blockchain_Constants"
    "encoding/json"
    "fmt"
    p "github.com/Crypt0plasm/Firefly-APD"
    svg "github.com/ajstarks/svgo"
    "log"
    "os"
)

type StyleSVG struct {
    StrokeColour string
    StrokeWidth string
    StrokeFill string
}

type Rectangle struct {
    X *p.Decimal
    Y *p.Decimal
    Length *p.Decimal
    Height *p.Decimal
}


func SVGStyleToString (InputStyle StyleSVG) string {
    S1 := "stroke:" + InputStyle.StrokeColour
    S2 := "stroke-width:" + InputStyle.StrokeWidth
    S3 := "fill:" + InputStyle.StrokeFill
    S4 := S1 + "; " + S2 + "; " + S3
    //S5 := "\"" + S4 + "\""
    return S4
}

type RGBColor struct {
    R uint8
    G uint8
    B uint8
}

type HEXColor struct {
    Hex string
}

//Colour Variables
var (
    Period      = RGBColor  {20,    220,    220}
    Black       = RGBColor  {0,     0,      0}
    White       = RGBColor  {255,   255,    255}
    Yellow      = RGBColor  {240,   220,    20}
    Red         = RGBColor  {240,   70,     20}
    Blue        = RGBColor  {20,    170,    240}
    Green       = RGBColor  {20,    240,    140}
    DarkBlue    = RGBColor  {20,    90,     240}
    Mov         = RGBColor  {90,    20,     240}
    LightBrown  = RGBColor  {220,   190,    140}
    HardBrown   = RGBColor  {220,   170,    80}
    Pink        = RGBColor  {240,   20,     170}
    DarkBrown   = RGBColor  {133,   72,     30}
    PureRed     = RGBColor  {255,   0,      0}


)

func HexFromRGB (Input RGBColor) HEXColor {
    return HEXColor{Hex: fmt.Sprintf("#%02x%02x%02x", Input.R, Input.G, Input.B)}
}



type ElrondAccount struct {
    Address         string `json:"address"`
    Balance         string `json:"balance"`
    Nonce           int    `json:"nonce"`
    Shard           int    `json:"shard"`
    RootHash        string `json:"rootHash"`
    TxCount         int    `json:"txCount"`
    ScrCount        int    `json:"scrCount"`
    Username        string `json:"username"`
    DeveloperReward string `json:"developerReward"`
}

var (
    UnityNFTStroke = "1"
)

func GetHeroTag (Addy b.ElrondAddress) string {
    var (
    	AccountPing = "https://api.elrond.com/accounts/"
    	ScannedJSON ElrondAccount
    	HeroTag string
    )

    Link := AccountPing + string(Addy)
    Snapshot := b.OnPage(Link)
    _ = json.Unmarshal([]byte(Snapshot), &ScannedJSON)

    //Returns NULL if Address has no HeroTag
    if ScannedJSON.Username == "" {
        HeroTag = "NULL"
    } else {
	HeroTag = ScannedJSON.Username
    }

    return HeroTag
}

func AddDigits (Number *p.Decimal) *p.Decimal {
    Sum := p.NFS("0")

    //Creating the Chain of runes representing the Decimal Number
    Coefficient := Number.Coeff
    DecimalAsLongText := Coefficient.Text(10)
    OriginalRuneSlice := MakeRuneChain(DecimalAsLongText)

    for i:=0; i<len(OriginalRuneSlice); i++ {
        Sum = b.ADDxc(Sum,p.NFS(string(OriginalRuneSlice[i])))
    }

    return Sum
}

func GetUnityNFTScore (Input b.UnityNftInput) (Score *p.Decimal) {
    Output := b.UnityNftComputer(Input)

    //Get Summed Time
    DaySummed := AddDigits(Output.Time.Day)
    YearSummer := AddDigits(Output.Time.Year)
    DayOfYearSummed := AddDigits(Output.Time.DayOfYear)
    S1 := b.SUMxc(DaySummed,YearSummer,DayOfYearSummed)

    //Get Raw Daily Mint Sum
    S2 := AddDigits(Input.RawDailyMint)
    //Get Volumetric Daily Tx Tax Sum
    S3 := AddDigits(Input.TxTax)
    //Get Treasury Sum
    S4 := AddDigits(Output.Treasury)

    //Get Minor Daily Total Sum
    S5 := AddDigits(Output.Minor)
    //Get Percentual Node Bonus Sum
    S6 := AddDigits(Input.Network.Bonus)
    //Get Absolute Node Bonus Sum
    S7 := AddDigits(Output.AbsoluteBonus)

    //Get Major Daily Total Sum
    S8 := AddDigits(Output.Major)
    //Get Elite Auryn Sum Total
    S9 := AddDigits(Output.EliteAuryn)

    //Get Minting Power Daily Total Sum
    //S10 := S6

    //Get VOP Node Numbers Sum
    N1 := AddDigits(Input.Network.Validator)
    N2 := AddDigits(Input.Network.Observer)
    N3 := AddDigits(Input.Network.Power)
    S10 := b.SUMxc(N1,N2,N3)


    //Get Final Score
    Score = b.SUMxc(S1,S2,S3,S4,S5,S6,S7,S8,S9,S10)

    return
}

//Remove decimals, and converts resulting integer to a string with thousands separator, using "." as the separator
func ConvertToThousandSeparator(Number *p.Decimal) string {
    var (
    	Separators *p.Decimal
    	NewChain []rune
    	Result string
    )
    Separator := MakeRuneChain(".")[0]

    //Function to insert an element (Value to insert) in a slice (in this example runes) at a given position (Index)
    InsertIntoRuneSlice := func(RuneSlice []rune, Index int, ValueToInsert rune) []rune {
        if len(RuneSlice) == Index { //nil or empty slice or after the last element
            return append(RuneSlice, ValueToInsert)
        }
        RuneSlice = append(RuneSlice[:Index+1], RuneSlice[Index:]...) //Index < len(RuneSlice)
        RuneSlice[Index] = ValueToInsert
        return RuneSlice
    }

    IntegerNumber := b.RemoveDecimals(Number)
    IntegerAsLongText := MakeRuneChain(IntegerNumber.Coeff.Text(10))

    Integer := b.DivInt(p.NFI(IntegerNumber.NumDigits()),p.NFS("3"))
    Rest := b.DivMod(p.NFI(IntegerNumber.NumDigits()),p.NFS("3"))

    //Compute number of Separators
    if b.DecimalEqual(Rest,p.NFS("0")) == true {
        Separators = b.SUBxc(Integer,p.NFS("1"))
    } else {
        Separators = Integer
    }
    SeparatorValue, _ := Separators.Int64()

    //Insert Separators: No separators inserted if smaller than 4 digits
    if SeparatorValue == 0 {
        Result = string(IntegerAsLongText)
    } else {
        //Insert Separators
        for i:=1; i<=int(SeparatorValue); i++ {
            if i == 1 {
                Index := len(IntegerAsLongText) - 3 * i
                NewChain = InsertIntoRuneSlice(IntegerAsLongText,Index,Separator)
            } else {
                Index := len(IntegerAsLongText) - 3 * i - i
                NewChain = InsertIntoRuneSlice(NewChain,Index+i,Separator)
            }

        }
        Result = string(NewChain)
    }
    return Result
}

func DrawUnitySVG (Day string, TxTaxAddy, Winner b.ElrondAddress) *svg.SVG {
    //WinnerHeroTag := GetHeroTag(Winner)
    UnityInput := b.UnityDayScanner(Day, TxTaxAddy)
    UnityOutput := b.UnityNftComputer(UnityInput)
    NFTScore := GetUnityNFTScore(UnityInput)


    var (
        Width = 10800
        Height = 10800
        Zero = p.NFS("0")
        PrimaryUpperRectangle Rectangle
        SecondaryUpperRectangle Rectangle
        PrimaryMiddleRectangle Rectangle
        SecondaryMiddleRectangle Rectangle
        PrimaryLowerRectangle Rectangle
        SecondaryLowerRectangle Rectangle
        MintingPowerRectangle Rectangle
        NodeRectangle Rectangle

        //Scaling Factors
        ScalingFactorDay                = p.NFS("0.8")
        ScalingFactorYear               = p.NFS("0.4")
        ScalingFactorDayOfYear          = p.NFS("0.32")
        ScalingFactorLetter             = p.NFS("0.203125")
        ScalingFactorMintingPower       = p.NFS("0.3880596875")
        ScalingFactorSmallNumber        = p.NFS("0.2")

        //SVG Drawing Styles
        BlankoStyle = StyleSVG{
            HexFromRGB(Black).Hex,
            UnityNFTStroke,
            "none"}
        PolygonStyle = StyleSVG{
            HexFromRGB(Black).Hex,
            b.DTS(b.MULxc(p.NFS(UnityNFTStroke),p.NFS("2"))),
            HexFromRGB(Black).Hex}
        DayStyle = StyleSVG{
            HexFromRGB(Period).Hex,
            b.DTS(b.DIVxc(p.NFS(UnityNFTStroke),ScalingFactorDay)),
            HexFromRGB(Period).Hex}
        YearStyle = StyleSVG{
            HexFromRGB(Period).Hex,
            b.DTS(b.DIVxc(p.NFS(UnityNFTStroke),ScalingFactorYear)),
            HexFromRGB(Period).Hex}
        DayOfYearStyle = StyleSVG{
            HexFromRGB(Period).Hex,
            b.DTS(b.DIVxc(p.NFS(UnityNFTStroke),ScalingFactorDayOfYear)),
            HexFromRGB(Period).Hex}
        LettersStyle = StyleSVG{
            HexFromRGB(Yellow).Hex,
            b.DTS(b.DIVxc(p.NFS(UnityNFTStroke),ScalingFactorLetter)),
            HexFromRGB(Yellow).Hex}
        Letters2Style = StyleSVG{
            HexFromRGB(Yellow).Hex,
            b.DTS(b.DIVxc(p.NFS(UnityNFTStroke),ScalingFactorMintingPower)),
            HexFromRGB(Yellow).Hex}
        NodesStyle = StyleSVG{
            HexFromRGB(Red).Hex,
            b.DTS(b.DIVxc(p.NFS(UnityNFTStroke),ScalingFactorLetter)),
            HexFromRGB(Red).Hex}
        ScoreStyle = StyleSVG{
            HexFromRGB(PureRed).Hex,
            b.DTS(b.DIVxc(p.NFS(UnityNFTStroke),ScalingFactorLetter)),
            HexFromRGB(PureRed).Hex}
        Number0StyleB = StyleSVG{
            HexFromRGB(Green).Hex,
            b.DTS(b.DIVxc(p.NFS(UnityNFTStroke),ScalingFactorYear)),
            HexFromRGB(Green).Hex}
        Number0StyleS = StyleSVG{
            HexFromRGB(Green).Hex,
            b.DTS(b.DIVxc(p.NFS(UnityNFTStroke),ScalingFactorSmallNumber)),
            HexFromRGB(Green).Hex}
        Number1StyleB = StyleSVG{
            HexFromRGB(Blue).Hex,
            b.DTS(b.DIVxc(p.NFS(UnityNFTStroke),ScalingFactorYear)),
            HexFromRGB(Blue).Hex}
        Number1StyleS = StyleSVG{
            HexFromRGB(Blue).Hex,
            b.DTS(b.DIVxc(p.NFS(UnityNFTStroke),ScalingFactorSmallNumber)),
            HexFromRGB(Blue).Hex}
        Number2StyleB = StyleSVG{
            HexFromRGB(DarkBlue).Hex,
            b.DTS(b.DIVxc(p.NFS(UnityNFTStroke),ScalingFactorYear)),
            HexFromRGB(DarkBlue).Hex}
        Number2StyleS = StyleSVG{
            HexFromRGB(DarkBlue).Hex,
            b.DTS(b.DIVxc(p.NFS(UnityNFTStroke),ScalingFactorSmallNumber)),
            HexFromRGB(DarkBlue).Hex}
        Number3StyleB = StyleSVG{
            HexFromRGB(Mov).Hex,
            b.DTS(b.DIVxc(p.NFS(UnityNFTStroke),ScalingFactorYear)),
            HexFromRGB(Mov).Hex}
        Number3StyleS = StyleSVG{
            HexFromRGB(Mov).Hex,
            b.DTS(b.DIVxc(p.NFS(UnityNFTStroke),ScalingFactorSmallNumber)),
            HexFromRGB(Mov).Hex}
        Number4StyleB = StyleSVG{
            HexFromRGB(Pink).Hex,
            b.DTS(b.DIVxc(p.NFS(UnityNFTStroke),ScalingFactorYear)),
            HexFromRGB(Pink).Hex}
        Number4StyleS = StyleSVG{
            HexFromRGB(Pink).Hex,
            b.DTS(b.DIVxc(p.NFS(UnityNFTStroke),ScalingFactorSmallNumber)),
            HexFromRGB(Pink).Hex}
        BracketStyle = StyleSVG{
            HexFromRGB(LightBrown).Hex,
            b.DTS(b.DIVxc(p.NFS(UnityNFTStroke),ScalingFactorYear)),
            HexFromRGB(LightBrown).Hex}
        BarStyle = StyleSVG{
            HexFromRGB(HardBrown).Hex,
            b.DTS(b.DIVxc(p.NFS(UnityNFTStroke),ScalingFactorYear)),
            HexFromRGB(HardBrown).Hex}
        CapitalStyle = StyleSVG{
            HexFromRGB(DarkBrown).Hex,
            b.DTS(b.DIVxc(p.NFS(UnityNFTStroke),ScalingFactorLetter)),
            HexFromRGB(DarkBrown).Hex}
    )

    //Outputs SVG content directly to external file
    OutputSVGName := "Day" + Day + ".svg"
    OutputVariable, err := os.Create(OutputSVGName)
    if err != nil {
	log.Fatal(err)
    }
    defer OutputVariable.Close()

    UnityDailyNFT := svg.New(OutputVariable)

    //Functions

    //1)Converts an integer p.Decimal to int
    IntDecToInt := func(Decimal *p.Decimal) int {
        Integer64, _ := Decimal.Int64()
        return int(Integer64)
    }

    //2)Converts a p.Decimal to float
    DecToFloat := func(Decimal *p.Decimal) float64 {
        Float, _ := Decimal.Float64()
        return Float
    }

    //3)Draws a rectangle within the UnityDailyNFT
    DrawRectangle := func(Rect Rectangle, Style StyleSVG) () {
        C1 := IntDecToInt(Rect.X)
        C2 := IntDecToInt(Rect.Y)
        C3 := IntDecToInt(Rect.Length)
        C4 := IntDecToInt(Rect.Height)

        UnityDailyNFT.Rect(C1,C2,C3,C4,SVGStyleToString(Style))
    }

    //4)Draws a Word a X/Y Coordinates that is also scaled
    TSDraw := func (X, Y, SF *p.Decimal, Text string, Style StyleSVG, OV *os.File) () {
        UnityDailyNFT.Translate(IntDecToInt(X),IntDecToInt(Y))
        UnityDailyNFT.Scale(DecToFloat(SF))
        DrawWord(Zero,Zero,Text,Style,OV)
        UnityDailyNFT.Gend()
        UnityDailyNFT.Gend()
    }

    //5)Draws the Day, Year and DayOfYear Numbers
    //Types used are "Day", "Year", "DayOfYear".
    DrawPeriod := func (Type string) () {
        var(
            TS string
            NActualLength, X, Y, SF *p.Decimal
            Style StyleSVG
        )
        switch Variant := Type; {
        case Variant == "Day":
            TS = ConvertToThousandSeparator(p.NFS(Day))
            NActualLength = b.PRDxc(p.NFI(GetTextLengthWithKerning(TS)),p.NFS("200"),ScalingFactorDay)
            Y = p.NFS("210")
            SF = ScalingFactorDay
            Style = DayStyle
        case Variant == "Year":
            TS = ConvertToThousandSeparator(UnityOutput.Time.Year)
            NActualLength = b.PRDxc(p.NFI(GetTextLengthWithKerning(TS)),p.NFS("200"),ScalingFactorYear)
            Y = p.NFS("1505")
            SF = ScalingFactorYear
            Style = YearStyle
        case Variant == "DayOfYear":
            TS = ConvertToThousandSeparator(UnityOutput.Time.DayOfYear)
            NActualLength = b.PRDxc(p.NFI(GetTextLengthWithKerning(TS)),p.NFS("200"),ScalingFactorDayOfYear)
            Y = p.NFS("2219")
            SF = ScalingFactorDayOfYear
            Style = DayOfYearStyle
        }
        X = b.SUBxc(p.NFS("5400"),b.DIVxc(NActualLength,p.NFS("2")))

        //Effective Draw
        TSDraw(X,Y,SF,TS,Style,OutputVariable)
    }

    //6)Draws the Network Node Numbers
    //)Types used are "Validator", "Observer" and "Power"
    DrawNodes := func (Type string) () {
        var (
            TS string
            NumberLength int64
            NActualLength, TActualLength, XN, XT, Y *p.Decimal
        )
        switch Variant := Type; {
        case Variant == "Score":
            TS = ConvertToThousandSeparator(NFTScore)
            Y = p.NFS("8975")
        case Variant == "Validator":
            TS = ConvertToThousandSeparator(UnityInput.Network.Validator)
            Y = p.NFS("9300")
        case Variant == "Observer":
            TS = ConvertToThousandSeparator(UnityInput.Network.Observer)
            Y = p.NFS("9625")
        case Variant == "Power":
            TS = ConvertToThousandSeparator(UnityInput.Network.Power)
            Y = p.NFS("9950")
        }
        NumberLength = GetTextLengthWithKerning(TS)
        NActualLength = b.TruncateCustom(b.MULxc(p.NFI(NumberLength),p.NFS("40.625")),0)
        TActualLength = b.TruncateCustom(b.PRDxc(p.NFS("1.25"),p.NFS("1600"),ScalingFactorLetter),0)

        XN = b.SUBxc(p.NFS("10610"),b.SUMxc(NActualLength,TActualLength,p.NFS("40")))
        XT = b.ADDxc(XN,NActualLength)


        //Draw Letter Part
        switch Variant := Type; {
        case Variant == "Score":
            TSDraw(XN,Y,ScalingFactorLetter,TS,ScoreStyle,OutputVariable)
            TSDraw(XT,Y,ScalingFactorLetter,":S",LettersStyle,OutputVariable)
        case Variant == "Validator":
            TSDraw(XN,Y,ScalingFactorLetter,TS,NodesStyle,OutputVariable)
            TSDraw(XT,Y,ScalingFactorLetter,":V",LettersStyle,OutputVariable)
        case Variant == "Observer":
            TSDraw(XN,Y,ScalingFactorLetter,TS,NodesStyle,OutputVariable)
            TSDraw(XT,Y,ScalingFactorLetter,":O",LettersStyle,OutputVariable)
        case Variant == "Power":
            TSDraw(XN,Y,ScalingFactorLetter,TS,NodesStyle,OutputVariable)
            UnityDailyNFT.Translate(IntDecToInt(XT),IntDecToInt(Y))
            UnityDailyNFT.Scale(DecToFloat(ScalingFactorLetter))

            UnityDailyNFT.Path(ComputeGlyphCode(PointDown, Zero, Zero, false), SVGStyleToString(LettersStyle))
            UnityDailyNFT.Path(ComputeGlyphCode(PointUp, Zero, Zero, false), SVGStyleToString(LettersStyle))
            MovementDistance := p.NFI(PointDown.Width)
            UnityDailyNFT.Path(ComputeGlyphCode(CapitalLeftPillar,MovementDistance,Zero, false), SVGStyleToString(CapitalStyle))
            UnityDailyNFT.Path(ComputeGlyphCode(Power,MovementDistance,Zero, true), SVGStyleToString(LettersStyle))
            UnityDailyNFT.Path(ComputeGlyphCode(CapitalRightPillar,MovementDistance,Zero, false), SVGStyleToString(CapitalStyle))

            UnityDailyNFT.Gend()
            UnityDailyNFT.Gend()
        }
    }

    //7)Draws Numbers with 18 Decimals on the main plane.
    //Types used are
    //      Type 0: Treasury and Percentual Node Bonus
    //      Type 1: Upper Rectangle
    //      Type 2: Middle Rectangle
    //      Type 3: Lower Rectangle
    //      Type 4: Lowest Rectangle
    DrawUnityValue := func (Number, Y *p.Decimal, Type string) () {
        //Type = "0" (invisible numbers)
        //Type = "1" (Upper numbers)
        //Type = "2" (middle numbers)
        //Type = "3" (lower number)
        //Type = "4" (minting power)
        var (
            Integer2Write string
            StyleBig, StyleSmall StyleSVG
        )
        switch Variant := Type; {
        case Variant == "0":
            StyleBig = Number0StyleB
            StyleSmall = Number0StyleS
        case Variant == "1":
            StyleBig = Number1StyleB
            StyleSmall = Number1StyleS
        case Variant == "2":
            StyleBig = Number2StyleB
            StyleSmall = Number2StyleS
        case Variant == "3":
            StyleBig = Number3StyleB
            StyleSmall = Number3StyleS
        case Variant == "4":
            StyleBig = Number4StyleB
            StyleSmall = Number4StyleS
        }

        PrecisionNumber := b.TruncateCustom(Number,18)
        S0 := ConvertToThousandSeparator(PrecisionNumber)
        RC := MakeRuneChain(b.DTS(b.SUBxc(PrecisionNumber,b.RemoveDecimals(PrecisionNumber))))[2:]
        S1 := string(RC[:3])
        S2 := string(RC[3:6])
        S3 := string(RC[6:9])
        S4 := string(RC[9:12])
        S5 := string(RC[12:15])
        S6 := string(RC[15:])

        IntegerLength := GetTextLengthWithKerning(S0)
        IntegerActualLength := b.PRDxc(p.NFI(IntegerLength),p.NFS("200"),ScalingFactorYear)
        X := b.SUBxc(p.NFS("5400"),b.ADDxc(IntegerActualLength,p.NFS("80")))

        if S0 == "0" {
            Integer2Write = ""
        } else {
            Integer2Write = S0 + ","
        }

        TSDraw(X,Y,ScalingFactorYear,Integer2Write,StyleBig,OutputVariable)
        X1 := p.NFS("5680")
        X2 := p.NFS("6440")
        X3 := p.NFS("7200")
        Y2 := b.ADDxc(Y,p.NFS("320"))

        TSDraw(X1,Y,ScalingFactorSmallNumber,S1,StyleSmall,OutputVariable)
        TSDraw(X2,Y,ScalingFactorSmallNumber,S2,StyleSmall,OutputVariable)
        TSDraw(X3,Y,ScalingFactorSmallNumber,S3,StyleSmall,OutputVariable)
        TSDraw(X1,Y2,ScalingFactorSmallNumber,S4,StyleSmall,OutputVariable)
        TSDraw(X2,Y2,ScalingFactorSmallNumber,S5,StyleSmall,OutputVariable)
        TSDraw(X3,Y2,ScalingFactorSmallNumber,S6,StyleSmall,OutputVariable)

        TSDraw(b.SUBxc(X1,p.NFS("240")),Y,ScalingFactorYear, "[", BracketStyle, OutputVariable)
        TSDraw(b.ADDxc(X3,p.NFS("600")),Y,ScalingFactorYear, "]", BracketStyle, OutputVariable)
        TSDraw(b.ADDxc(X1,p.NFS("600")),Y,ScalingFactorYear, "|", BarStyle, OutputVariable)
        TSDraw(b.ADDxc(X2,p.NFS("600")),Y,ScalingFactorYear, "|", BarStyle, OutputVariable)

    }

    //8)Draw Polygon Function
    //Draws the "bar" between Rectangles
    DrawPolygon := func (X,Y *p.Decimal) () {
        MUL2 := func (N *p.Decimal) *p.Decimal {
            return b.MULxc(N,p.NFS("2"))
        }
        P1X := IntDecToInt(MUL2(b.ADDxc(X,p.NFS("120"))))
        P1Y := IntDecToInt(MUL2(b.ADDxc(Y,p.NFS("80"))))
        P2X := IntDecToInt(MUL2(b.ADDxc(X,p.NFS("40"))))
        P2Y := IntDecToInt(MUL2(Y))
        P3X := P1X
        P3Y := IntDecToInt(MUL2(b.SUBxc(Y,p.NFS("80"))))

        P4X := IntDecToInt(MUL2(p.NFS("7960")))
        P4Y := P3Y
        P5X := IntDecToInt(MUL2(p.NFS("8040")))
        P5Y := P2Y
        P6X := P4X
        P6Y := P1Y

        var (
            PX = make([]int,6)
            PY = make([]int,6)
        )
        PX[0] = P1X
        PX[1] = P2X
        PX[2] = P3X
        PX[3] = P4X
        PX[4] = P5X
        PX[5] = P6X

        PY[0] = P1Y
        PY[1] = P2Y
        PY[2] = P3Y
        PY[3] = P4Y
        PY[4] = P5Y
        PY[5] = P6Y

        SF := p.NFS("0.5")

        UnityDailyNFT.Scale(DecToFloat(SF))
        Style:=SVGStyleToString(PolygonStyle)
        UnityDailyNFT.Polygon(PX,PY,Style)

        UnityDailyNFT.Gend()
    }


    //STARTs the DRAWING Process
    //0 Starts creating the SVG
    UnityDailyNFT.Start(Width,Height)

    //1 Draw Outer Rectangle, that is 10800x10800 exactly the size of the board.
    MainSquare := Rectangle{Zero,Zero,p.NFI(int64(Width)),p.NFI(int64(Height))}
    DrawRectangle(MainSquare,BlankoStyle)

    //2 Draw Period
    DrawPeriod("Day")
    DrawPeriod("Year")
    DrawPeriod("DayOfYear")

    //3a)Getting max Main Rectangle Size.
    //Get Board Background Rectangle point of origin and length. (First 3 Rectangle categories)
    //Integer Part of the Number has a text size of 640. Represents 0.4 out of 1600 (Glyph code definition).
    //Therefore a scaling factor of 0.4 must be used
    OutputMajorConvTT := ConvertToThousandSeparator(UnityOutput.Major)
    OutputMajorConvTTT := GetTextLengthWithKerning(OutputMajorConvTT)
    //Get Length of the resulted Number Display (1 unit is 200), then multiply it by 0.4 to get actual Length in SVG
    //As such multiplication with 200*0.4 which is 80 is performed
    OutputMajorConvTTL := b.MULxc(p.NFI(OutputMajorConvTTT),p.NFS("80"))
    //80 Units (Half of "," Character are added to the length
    //40 Units (20 and 20 as rectangle border from the text length must also be added)
    //120 in total must be added
    OutputMajorConvTTLF := b.ADDxc(OutputMajorConvTTL,p.NFS("120"))
    LastX := b.SUBxc(p.NFS("5400"),OutputMajorConvTTLF)
    LastY := p.NFS("2760")

    //3b)Draw the Bar between Rectangles
    DrawPolygon(LastX, p.NFS("4912.5"))
    DrawPolygon(LastX, p.NFS("7187.5"))

    //3c)Computing Rectangles based on Max Rectangle Size
    PrimaryUpperRectangle.X = LastX
    PrimaryUpperRectangle.Y = LastY
    PrimaryUpperRectangle.Length = b.SUBxc(p.NFS("10610"),PrimaryUpperRectangle.X)
    PrimaryUpperRectangle.Height = b.SUBxc(p.NFS("4790"),PrimaryUpperRectangle.Y)

    SecondaryUpperRectangle.X = b.ADDxc(PrimaryUpperRectangle.X,p.NFS("20"))
    SecondaryUpperRectangle.Y = b.ADDxc(PrimaryUpperRectangle.Y,p.NFS("20"))
    SecondaryUpperRectangle.Length = b.SUBxc(p.NFS("8000"),SecondaryUpperRectangle.X)    //Until Capital Starts
    SecondaryUpperRectangle.Height = b.SUBxc(PrimaryUpperRectangle.Height,p.NFS("40"))

    PrimaryMiddleRectangle.X = PrimaryUpperRectangle.X
    PrimaryMiddleRectangle.Y = b.ADDxc(PrimaryUpperRectangle.Y,p.NFS("2275"))
    PrimaryMiddleRectangle.Length = PrimaryUpperRectangle.Length
    PrimaryMiddleRectangle.Height = PrimaryUpperRectangle.Height

    SecondaryMiddleRectangle.X = SecondaryUpperRectangle.X
    SecondaryMiddleRectangle.Y = b.ADDxc(SecondaryUpperRectangle.Y,p.NFS("2275"))
    SecondaryMiddleRectangle.Length = SecondaryUpperRectangle.Length
    SecondaryMiddleRectangle.Height = SecondaryUpperRectangle.Height

    PrimaryLowerRectangle.X = PrimaryUpperRectangle.X
    PrimaryLowerRectangle.Y = b.ADDxc(PrimaryMiddleRectangle.Y,p.NFS("2275"))
    PrimaryLowerRectangle.Length = PrimaryUpperRectangle.Length
    PrimaryLowerRectangle.Height = b.SUBxc(p.NFS("8690"),PrimaryLowerRectangle.Y)

    SecondaryLowerRectangle.X = SecondaryUpperRectangle.X
    SecondaryLowerRectangle.Y = b.ADDxc(SecondaryMiddleRectangle.Y,p.NFS("2275"))
    SecondaryLowerRectangle.Length = SecondaryUpperRectangle.Length
    SecondaryLowerRectangle.Height = b.SUBxc(PrimaryLowerRectangle.Height,p.NFS("40"))

    MintingPowerRectangle.X = p.NFS("2760")
    MintingPowerRectangle.Y = p.NFS("8935")
    MintingPowerRectangle.Length = p.NFS("5280")
    MintingPowerRectangle.Height = p.NFS("1380")

    //3d)Getting max Node Rectangle Size
    //Integer Part of the Number has a text size of 325. Represents 0.203125 out of 1600 (Glyph code definition).
    //Therefore a scaling factor of 0.203125 must be used
    NodePowerConvTT := ConvertToThousandSeparator(UnityInput.Network.Power)
    NodePowerConvTTT := GetTextLengthWithKerning(NodePowerConvTT)
    //Get Length of the resulted Number Display (1 unit is 200), then multiply it by 0.203125 to get actual Length in SVG
    //As such multiplication with 200*0.203125 which is 40.625 is performed
    NodePowerConvTTL := b.MULxc(p.NFI(NodePowerConvTTT),p.NFS("40.625"))
    //406.25 Units (1.25x Glyph Size = Full Glyph plus quarter ":") 1.25*1600*0.203125 = 406.25
    //80 Units also must be added, 20 inner border per side, and 20 offset for a 40 pixel line width
    NodePowerConvTTLF := b.ADDxc(NodePowerConvTTL,p.NFS("486.25"))
    //RoundUp, using TruncateCustom, with 0 precision.
    NodePowerConvTTLFT := b.TruncateCustom(NodePowerConvTTLF,0)
    NodeLastX := b.SUBxc(p.NFS("10610"),NodePowerConvTTLFT)
    NodeLastY := p.NFS("8935")

    //3e)Computing Node Rectangle Size
    NodeRectangle.X = NodeLastX
    NodeRectangle.Y = NodeLastY
    NodeRectangle.Length = NodePowerConvTTLFT
    NodeRectangle.Height = p.NFS("1380")

    //3f)Draw Rectangles
    //DrawRectangle(PrimaryUpperRectangle,BlankoStyle)
    //DrawRectangle(SecondaryUpperRectangle,BlankoStyle)
    //DrawRectangle(PrimaryMiddleRectangle,BlankoStyle)
    //DrawRectangle(SecondaryMiddleRectangle,BlankoStyle)
    //DrawRectangle(PrimaryLowerRectangle,BlankoStyle)
    //DrawRectangle(SecondaryLowerRectangle,BlankoStyle)
    //DrawRectangle(MintingPowerRectangle,BlankoStyle)
    //DrawRectangle(NodeRectangle,BlankoStyle)

    //4)Draw Words.
    LX := p.NFS("8001")
    MX := p.NFS("2800")

    TSDraw(LX,p.NFS("2800"),ScalingFactorLetter,"Raw",LettersStyle,OutputVariable)
    TSDraw(LX,p.NFS("3125"),ScalingFactorLetter,"Daily Mint",LettersStyle,OutputVariable)
    TSDraw(LX,p.NFS("3450"),ScalingFactorLetter,"Volumetric",LettersStyle,OutputVariable)
    TSDraw(LX,p.NFS("3775"),ScalingFactorLetter,"Daily tx tax",LettersStyle,OutputVariable)
    TSDraw(LX,p.NFS("4100"),ScalingFactorLetter,"1% to",LettersStyle,OutputVariable)
    TSDraw(LX,p.NFS("4425"),ScalingFactorLetter,"Treasury",LettersStyle,OutputVariable)
    TSDraw(LX,p.NFS("5075"),ScalingFactorLetter,"Minor",LettersStyle,OutputVariable)
    TSDraw(LX,p.NFS("5400"),ScalingFactorLetter,"Daily Total",LettersStyle,OutputVariable)
    TSDraw(LX,p.NFS("5725"),ScalingFactorLetter,"Percentual",LettersStyle,OutputVariable)
    TSDraw(LX,p.NFS("6050"),ScalingFactorLetter,"Node Bonus",LettersStyle,OutputVariable)
    TSDraw(LX,p.NFS("6375"),ScalingFactorLetter,"Absolute",LettersStyle,OutputVariable)
    TSDraw(LX,p.NFS("6700"),ScalingFactorLetter,"Node Bonus",LettersStyle,OutputVariable)
    TSDraw(LX,p.NFS("7350"),ScalingFactorLetter,"Major",LettersStyle,OutputVariable)
    TSDraw(LX,p.NFS("7675"),ScalingFactorLetter,"Daily Total",LettersStyle,OutputVariable)
    TSDraw(LX,p.NFS("8000"),ScalingFactorLetter,"1‰ to",LettersStyle,OutputVariable)
    TSDraw(LX,p.NFS("8325"),ScalingFactorLetter,"Elite Auryn",LettersStyle,OutputVariable)
    TSDraw(MX,p.NFS("8990"),ScalingFactorMintingPower,"Minting Power",Letters2Style,OutputVariable)

    //5)
    //Network Node Text and Numbers
    DrawNodes("Score")
    DrawNodes("Validator")
    DrawNodes("Observer")
    DrawNodes("Power")

    //6)
    //Draw Unity Numbers
    DrawUnityValue(UnityInput.RawDailyMint,p.NFS("2805"),"1")
    DrawUnityValue(UnityInput.TxTax,p.NFS("3455"),"1")
    DrawUnityValue(UnityOutput.Treasury,p.NFS("4105"),"0")

    DrawUnityValue(UnityOutput.Minor,p.NFS("5080"),"2")
    DrawUnityValue(UnityInput.Network.Bonus,p.NFS("5730"),"0")
    DrawUnityValue(UnityOutput.AbsoluteBonus,p.NFS("6380"),"2")

    DrawUnityValue(UnityOutput.Major,p.NFS("7355"),"3")
    DrawUnityValue(UnityOutput.EliteAuryn,p.NFS("8005"),"3")

    MintingPower := b.ADDxc(UnityInput.Network.Bonus,b.DIVxc(NFTScore,p.NFS("32")))
    DrawUnityValue(MintingPower,p.NFS("9630"),"4")

    //Last Command from the SVG
    UnityDailyNFT.End()
    return UnityDailyNFT
}