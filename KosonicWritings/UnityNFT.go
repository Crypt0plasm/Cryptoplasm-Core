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
    Period  = RGBColor  {20,    220,    220}
    Black   = RGBColor  {0,     0,      0}
    White   = RGBColor  {255,   255,    255}
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
    UnityNFTStroke = "2"
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
    S10 := S6

    //Get VOP Node Numbers Sum
    N1 := AddDigits(Input.Network.Validator)
    N2 := AddDigits(Input.Network.Observer)
    N3 := AddDigits(Input.Network.Power)
    S11 := b.SUMxc(N1,N2,N3)


    //Get Final Score
    Score = b.SUMxc(S1,S2,S3,S4,S5,S6,S7,S8,S9,S10,S11)

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
    //NFTScore := GetUnityNFTScore(UnityInput)

    IntDecToInt := func(Decimal *p.Decimal) int {
        Integer64, _ := Decimal.Int64()
        return int(Integer64)
    }

    DrawRectangle := func(Rect Rectangle, Vector *svg.SVG, Style StyleSVG) () {
        C1 := IntDecToInt(Rect.X)
        C2 := IntDecToInt(Rect.Y)
        C3 := IntDecToInt(Rect.Length)
        C4 := IntDecToInt(Rect.Height)

        Vector.Rect(C1,C2,C3,C4,SVGStyleToString(Style))
    }

    var (
        Width = 10800
        Height = 10800
        PrimaryUpperRectangle Rectangle
        SecondaryUpperRectangle Rectangle
        PrimaryMiddleRectangle Rectangle
        SecondaryMiddleRectangle Rectangle
        PrimaryLowerRectangle Rectangle
        SecondaryLowerRectangle Rectangle
        MintingPowerRectangle Rectangle
        NodeRectangle Rectangle
    )

    //Outputs SVG content directly to external file
    OutputSVGName := "Day" + Day + ".svg"
    OutputVariable, err := os.Create(OutputSVGName)
    if err != nil {
	log.Fatal(err)
    }
    defer OutputVariable.Close()

    UnityDailyNFT := svg.New(OutputVariable)

    //0 Starts creating the SVG
    UnityDailyNFT.Start(Width,Height)

    //1 Draw Outer Rectangle, that is 10800x10800 exactly the size of the board.
    BlankoStyle := StyleSVG{
        HexFromRGB(Black).Hex,
        UnityNFTStroke,
        "none"}
    UnityDailyNFT.Rect(0,0,Width,Height,SVGStyleToString(BlankoStyle))

    //Draw Day Number
    //Size is 1280 for this Number. Represents 0.8 out of 1600 (Glyph code definition).
    //Therefore a scaling factor of 0.8 must be used
        //Convert string Day to thousand separators
        DayConvTT := ConvertToThousandSeparator(p.NFS(Day))
        //Get Length of the resulted Day Display (1 unit is 200), then multiply it by 0.8 to get actual Length in SVG
        //As such multiplication with 200*0.8 which is 160 is performed
        DayConvTTT := GetTextLengthWithKerning(DayConvTT)
        DayConvTTL := b.MULxc(p.NFI(DayConvTTT),p.NFS("160"))
        //Get X,Y Coordinates as starting point for the Day Number Write
        //Coordinates must be scaled to reflect the future downscale, as such must be divided by 0.8
        DayX := b.SUBxc(p.NFS("5400"),b.DIVxc(DayConvTTL,p.NFS("2")))
        DayY := p.NFS("210")
        DayXX := b.MULxc(b.DIVxc(DayX,p.NFS("8")),p.NFS("10"))
        DayYY := b.DIVxc(DayY,p.NFS("0.8"))
        //Drawing the Day number, scaling it before
        ScalingFactorDay,_ := p.NFS("0.8").Float64()
        UnityDailyNFT.Scale(ScalingFactorDay)
            //Defining Style Used for Draw
            DayStyle := StyleSVG{
            HexFromRGB(Period).Hex,
            b.DTS(b.DIVxc(p.NFS(UnityNFTStroke),p.NFS("0.8"))),
            HexFromRGB(Period).Hex}
            //Effective Draw
            DrawWord(DayXX,DayYY, DayConvTT, DayStyle, OutputVariable)
        UnityDailyNFT.Gend()

    //Draw Year Number
    //Size is 640 for this Number. Represents 0.4 out of 1600 (Glyph code definition).
    //Therefore a scaling factor of 0.8 must be used
        //Convert string Day to thousand separators
        YearConvTT := ConvertToThousandSeparator(UnityOutput.Time.Year)
        //Get Length of the resulted Day Display (1 unit is 200), then multiply it by 0.4 to get actual Length in SVG
        //As such multiplication with 200*0.4 which is 80 is performed
        YearConvTTT := GetTextLengthWithKerning(YearConvTT)
        YearConvTTL := b.MULxc(p.NFI(YearConvTTT),p.NFS("80"))
        //Get X,Y Coordinates as starting point for the Year Number Write
        //Coordinates must be scaled to reflect the future downscale, as such must be divided by 0.4
        YearX := b.SUBxc(p.NFS("5400"),b.DIVxc(YearConvTTL,p.NFS("2")))
        YearY := p.NFS("1505")
        YearXX := b.DIVxc(YearX,p.NFS("0.4"))
        YearYY := b.DIVxc(YearY,p.NFS("0.4"))
        //Drawing the Year number, scaling it before
        ScalingFactorYear,_ := p.NFS("0.4").Float64()
        UnityDailyNFT.Scale(ScalingFactorYear)
        //Defining Style Used for Draw
            YearStyle := StyleSVG{
            HexFromRGB(Period).Hex,
            b.DTS(b.DIVxc(p.NFS(UnityNFTStroke),p.NFS("0.4"))),
            HexFromRGB(Period).Hex}
            //Same Style as Day
            //Effective Draw
            DrawWord(YearXX,YearYY, YearConvTT, YearStyle, OutputVariable)
        UnityDailyNFT.Gend()

    //Draw Day of Year Number
    //Size is 512 for this Number. Represents 0.32 out of 1600 (Glyph code definition).
    //Therefore a scaling factor of 0.32 must be used
        //Convert string Day to thousand separators
        DayOfYearConvTT := ConvertToThousandSeparator(UnityOutput.Time.DayOfYear)
        //Get Length of the resulted Day Display (1 unit is 200), then multiply it by 0.32 to get actual Length in SVG
        //As such multiplication with 200*0.32 which is 64 is performed
        DayOfYearConvTTT := GetTextLengthWithKerning(DayOfYearConvTT)
        DayOfYearConvTTL := b.MULxc(p.NFI(DayOfYearConvTTT),p.NFS("64"))
        //Get X,Y Coordinates as starting point for the DayOfYear Number Write
        //Coordinates must be scaled to reflect the future downscale, as such must be divided by 0.32
        DayOfYearX := b.SUBxc(p.NFS("5400"),b.DIVxc(DayOfYearConvTTL,p.NFS("2")))
        DayOfYearY := p.NFS("2219")
        DayOfYearXX := b.DIVxc(DayOfYearX,p.NFS("0.32"))
        DayOfYearYY := b.DIVxc(DayOfYearY,p.NFS("0.32"))
        //Drawing the DayOfYear number, scaling it before
        ScalingFactorDayOfYear,_ := p.NFS("0.32").Float64()
        UnityDailyNFT.Scale(ScalingFactorDayOfYear)
        //Defining Style Used for Draw
        DayOfYearStyle := StyleSVG{
            HexFromRGB(Period).Hex,
            b.DTS(b.DIVxc(p.NFS(UnityNFTStroke),p.NFS("0.32"))),
            HexFromRGB(Period).Hex}
        //Same Style as Day
        //Effective Draw
        DrawWord(DayOfYearXX,DayOfYearYY, DayOfYearConvTT, DayOfYearStyle, OutputVariable)
        UnityDailyNFT.Gend()

    //Get Board Background Rectangle point of origin and length.
    //Integer Part of the Number has a text size of 640. Represents 0.4 out of 1600 (Glyph code definition).
    //Therefore a scaling factor of 0.4 must be used
    OutputMajorConvTT := ConvertToThousandSeparator(UnityOutput.Major)
    OutputMajorConvTTT := GetTextLengthWithKerning(OutputMajorConvTT)
    //Get Length of the resulted Number Display (1 unit is 200), then multiply it by 0.4 to get actual Length in SVG
    //As such multiplication with 200*0.4 which is 80 is performed
    OutputMajorConvTTL := b.MULxc(p.NFI(OutputMajorConvTTT),p.NFS("80"))
    //80 Units (Half of "," Character are added to the length
    OutputMajorConvTTLF := b.ADDxc(OutputMajorConvTTL,p.NFS("80"))
    LastX := b.SUBxc(p.NFS("5400"),OutputMajorConvTTLF)
    LastY := p.NFS("2805")
    //Getting Rectangle Sizes.
    PrimaryUpperRectangle.X = b.ADDxc(LastX,p.NFS("40"))
    PrimaryUpperRectangle.Y = b.SUBxc(LastY,p.NFS("45"))
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
    PrimaryLowerRectangle.Height = b.SUBxc(p.NFS("8670"),PrimaryLowerRectangle.Y)

    SecondaryLowerRectangle.X = SecondaryUpperRectangle.X
    SecondaryLowerRectangle.Y = b.ADDxc(SecondaryMiddleRectangle.Y,p.NFS("2275"))
    SecondaryLowerRectangle.Length = SecondaryUpperRectangle.Length
    SecondaryLowerRectangle.Height = b.SUBxc(PrimaryLowerRectangle.Height,p.NFS("40"))

    MintingPowerRectangle.X = p.NFS("2760")
    MintingPowerRectangle.Y = p.NFS("8935")
    MintingPowerRectangle.Length = p.NFS("5280")
    MintingPowerRectangle.Height = p.NFS("1380")

    //Get Node Rectangle Size
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
    NodeLastY := p.NFS("9260")

    NodeRectangle.X = NodeLastX
    NodeRectangle.Y = NodeLastY
    NodeRectangle.Length = NodePowerConvTTLFT
    NodeRectangle.Height = p.NFS("1055")

    //Draw Rectangles
    DrawRectangle(PrimaryUpperRectangle,UnityDailyNFT,BlankoStyle)
    DrawRectangle(SecondaryUpperRectangle,UnityDailyNFT,BlankoStyle)
    DrawRectangle(PrimaryMiddleRectangle,UnityDailyNFT,BlankoStyle)
    DrawRectangle(SecondaryMiddleRectangle,UnityDailyNFT,BlankoStyle)
    DrawRectangle(PrimaryLowerRectangle,UnityDailyNFT,BlankoStyle)
    DrawRectangle(SecondaryLowerRectangle,UnityDailyNFT,BlankoStyle)
    DrawRectangle(MintingPowerRectangle,UnityDailyNFT,BlankoStyle)
    DrawRectangle(NodeRectangle,UnityDailyNFT,BlankoStyle)

    //UnityDailyNFT.Rect(0,0,Width,Height,SVGStyleToString(BlankoStyle))


    //Last Command from the SVG
    UnityDailyNFT.End()
    return UnityDailyNFT
}

