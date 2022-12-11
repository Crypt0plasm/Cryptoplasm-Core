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
    Period = RGBColor {20,220,220}
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
    //UnityInput := b.UnityDayScanner(Day, TxTaxAddy)
    //UnityOutput := b.UnityNftComputer(UnityInput)
    //NFTScore := GetUnityNFTScore(UnityInput)

    var (
        Width = 10800
        Height = 10800
	S1 = "fill=\"none\""
	S2 = "stroke=\"black\""
	S3 = "stroke-width=\"" + UnityNFTStroke + "\""
	//Zero = p.NFS("0")
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
    UnityDailyNFT.Rect(0,0,Width,Height,S1,S2,S3)

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
        DayXX := b.DIVxc(DayX,p.NFS("0.8"))
        DayYY := b.DIVxc(DayY,p.NFS("0.8"))
        //Drawing the Day number, scaling it before
        ScalingFactorDay,_ := p.NFS("0.8").Float64()
        UnityDailyNFT.Scale(ScalingFactorDay)
            //Defining Style Used for Draw
            DayStyle := StyleSVG{
            HexFromRGB(Period).Hex,
            p.DTS(b.DIVxc(p.NFS(UnityNFTStroke),p.NFS("0.8"))),
            HexFromRGB(Period).Hex}
            fmt.Println("ESTE",SVGStyleToString(DayStyle))
            //Effective Draw
            DrawWord(DayXX,DayYY, DayConvTT, DayStyle, OutputVariable)
        UnityDailyNFT.Gend()



    //Last Command from the SVG
    UnityDailyNFT.End()
    return UnityDailyNFT
}