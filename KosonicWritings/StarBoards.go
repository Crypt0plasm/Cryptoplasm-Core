package KosonicWritings

import (
    p "github.com/Crypt0plasm/Firefly-APD"
    "strings"
)

type Board struct {
    Name string
    Lines []int64
    Size int64
    ScalingFactor *p.Decimal
    AnchorX *p.Decimal
    AnchorY *p.Decimal
}

var (
    Board1600 = []int64{64,54,36}
    Board1200 = []int64{88,80,68,48}
    Board960 = []int64{112,106,96,82,60}
    Board800 = []int64{136,130,122,110,94,72}
    Board600 = []int64{184,178,170,162,150,138,120,96}
    Board480 = []int64{230,226,220,212,202,192,180,164,144,122}
    Board400 = []int64{278,274,268,260,252,244,232,220,206,190,170,146}
    Board320 = []int64{348,344,340,334,326,318,310,300,288,276,262,246,228,206,182}
    Board300 = []int64{372,368,362,358,350,342,334,326,314,302,290,276,258,240,220,194}
    Board240 = []int64{466,462,458,452,446,440,432,424,416,406,396,384,372,360,344,328,310,290,268,244}
    Board200 = []int64{560,556,552,548,542,536,530,522,514,506,498,488,478,466,454,442,428,414,398,380,362,340,318,292}
    Board192 = []int64{584,580,576,572,566,560,554,546,540,532,522,514,504,492,482,470,456,442,426,410,392,374,352,330,304}
    Board160 = []int64{702,698,694,690,684,680,674,668,660,654,646,638,630,620,610,600,590,578,566,552,540,524,510,492,476,456,436,414,390,366}
    Board150 = []int64{748,746,742,736,732,726,722,716,708,702,694,686,678,670,660,652,642,630,618,606,594,580,566,552,536,518,500,482,462,440,416,390}
    Board120 = []int64{938,934,930,926,922,916,912,906,900,894,888,882,874,866,858,850,842,834,824,814,804,794,782,770,758,746,734,720,704,690,674,658,640,622,602,582,560,538,514,488}
    Board100 = []int64{1126,1122,1118,1114,1110,1106,1100,1048,1090,1086,1080,1074,1066,1060,1054,1046,1038,1030,1022,1014,1006,996,988,978,968,956,946,934,922,910,898,886,872,858,844,828,812,796,778,760,742,724,702,682,660,636,610,584}

    Tablet1600 = Board {"Board1600",Board1600, GetBoardSize(Board1600),p.NFS("1"),p.NFS("2650") ,p.NFS("10450")}
    Tablet1200 = Board {"Board1200",Board1200, GetBoardSize(Board1200),p.NFS("0.75"), p.NFS("2450"),p.NFS("10450")}
    Tablet960 = Board {"Board960",Board960, GetBoardSize(Board960),p.NFS("0.6"),p.NFS("2330"),p.NFS("10450")}
    Tablet800 = Board {"Board800",Board800, GetBoardSize(Board800),p.NFS("0.5"),p.NFS("2250"),p.NFS("10450")}
    Tablet600 = Board {"Board600",Board600, GetBoardSize(Board600),p.NFS("0.375"), p.NFS("2150"),p.NFS("10450")}
    Tablet480 = Board {"Board480",Board480, GetBoardSize(Board480),p.NFS("0.3"),p.NFS("2150"),p.NFS("10450")}
    Tablet400 = Board {"Board400",Board400, GetBoardSize(Board400),p.NFS("0.25"),p.NFS("2100"),p.NFS("10450")}
    Tablet320 = Board {"Board320",Board320, GetBoardSize(Board320),p.NFS("0.2"),p.NFS("2090"),p.NFS("10450")}
    Tablet300 = Board {"Board300",Board300, GetBoardSize(Board300),p.NFS("0.1875"),p.NFS("2075"),p.NFS("10450")}
    Tablet240 = Board {"Board240",Board240, GetBoardSize(Board240),p.NFS("0.15"),p.NFS("2015"),p.NFS("10450")}
    Tablet200 = Board {"Board200",Board200, GetBoardSize(Board200),p.NFS("0.125"),p.NFS("2060"),p.NFS("10450")}
    Tablet192 = Board {"Board192",Board192, GetBoardSize(Board192),p.NFS("0.12"),p.NFS("2050"),p.NFS("10450")}
    Tablet160 = Board {"Board160",Board160, GetBoardSize(Board160),p.NFS("0.1"),p.NFS("2042"),p.NFS("10450")}
    Tablet150 = Board {"Board150",Board150, GetBoardSize(Board150),p.NFS("0.09375"),p.NFS("2030"),p.NFS("10450")}
    Tablet120 = Board {"Board120",Board120, GetBoardSize(Board120),p.NFS("0.075"),p.NFS("2037"),p.NFS("10450")}
    Tablet100 = Board {"Board100",Board100, GetBoardSize(Board100),p.NFS("0.0625"),p.NFS("2012"),p.NFS("10450")}

    TabletChain = []Board{Tablet1600,Tablet1200,Tablet960,Tablet800,Tablet600,Tablet480,Tablet400,Tablet320,Tablet300,Tablet240,Tablet200,Tablet192,Tablet160,Tablet150,Tablet120,Tablet100}

    KanonNamePlateX = 7400
    KanonNamePlateY = 15650
    KanonNamePlatePositions = int64(66)
)

func GetNextSmallerTablet (Tablet Board) Board {
    var OutputBoard Board
    for i:=0; i<len(TabletChain); i++ {
        if TabletChain[i].Name == "Board100" {
            OutputBoard = Tablet100
            break
        } else if TabletChain[i].Name == Tablet.Name {
            OutputBoard = TabletChain[i+1]
            break
        }
    }
    return OutputBoard
}

func GetBoardSize (Board []int64) (TotalSize int64) {
    for i:=0; i<len(Board); i++ {
        TotalSize = TotalSize + Board[i]
    }
    return TotalSize
}

func GetCorrectBoardSize(Text string) Board {
    var(
        TabletSize, UsedTabletSize Board
    )

    TabletSize = GetRequiredBoardSize(Text)
    SmallerTabletSize := GetNextSmallerTablet(TabletSize)

    Lines1 := SplitString(TabletSize,Text)
    Lines2 := SplitString(SmallerTabletSize,Text)

    LastLine1 := Lines1[len(Lines1)-1]
    LastLine2 := Lines2[len(Lines2)-1]

    LastLine1Chain := strings.Split(LastLine1," ")
    LastLine2Chain := strings.Split(LastLine2," ")

    LastLine1ChainLastElement := LastLine1Chain[len(LastLine1Chain)-1]
    LastLine2ChainLastElement := LastLine2Chain[len(LastLine2Chain)-1]


    Offset := ComputeStartingOffset(TabletSize,Lines1)
    //Jumps to the next Smaller Board if the initial computed Board doesnt fit.
    //It doesnt fit either by last negative offset, or
    //If the last word from the last line as computed using the initial board and smaller board are different.
    if Offset[len(Offset)-1] < 0 || LastLine1ChainLastElement != LastLine2ChainLastElement {
        //fmt.Println("Doesnt Fit",Offset[len(Offset)-1])
        UsedTabletSize = SmallerTabletSize
    } else {
        UsedTabletSize = TabletSize
        //fmt.Println("Fits")
    }
    return UsedTabletSize
}
