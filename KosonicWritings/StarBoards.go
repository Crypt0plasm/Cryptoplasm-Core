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
    AnchorX int
    AnchorY int
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
    Board100 = []int64{1126,1122,1118,1114,1110,1106,1100,1048,1090,1086,1080,1074,1066,1060,1054,1046,1038,1030,1022,1014,1006,996,988,978,968,956,946,934,922,910,898,886,872,858,844,828,812,796,778,760,742,724,702,682,660,636,610,584}

    Tablet1600 = Board {"Board1600",Board1600, GetBoardSize(Board1600),p.NFS("1"),2650,10450}
    Tablet1200 = Board {"Board1200",Board1200, GetBoardSize(Board1200),p.NFS("0.75"), 2450,10450}
    Tablet960 = Board {"Board960",Board960, GetBoardSize(Board960),p.NFS("0.6"),2330,10450}
    Tablet800 = Board {"Board800",Board800, GetBoardSize(Board800),p.NFS("0.5"),2250,10450}
    Tablet600 = Board {"Board600",Board600, GetBoardSize(Board600),p.NFS("0.375"), 2150,10450}
    Tablet480 = Board {"Board480",Board480, GetBoardSize(Board480),p.NFS("0.3"),2150,10450}
    Tablet400 = Board {"Board400",Board400, GetBoardSize(Board400),p.NFS("0.25"),2100,10450}
    Tablet320 = Board {"Board320",Board320, GetBoardSize(Board320),p.NFS("0.2"),2090,10450}
    Tablet300 = Board {"Board300",Board300, GetBoardSize(Board300),p.NFS("0.1875"),2075,10450}
    Tablet100 = Board {"Board100",Board100, GetBoardSize(Board100),p.NFS("0.0625"),20125,10450}

    TabletChain = []Board{Tablet1600,Tablet1200,Tablet960,Tablet800,Tablet600,Tablet480,Tablet400,Tablet320,Tablet300}

    KanonNamePlateX = 7400
    KanonNamePlateY = 15650
    KanonNamePlatePositions = int64(66)
)

func GetNextSmallerTablet (Tablet Board) Board {
    var OutputBoard Board
    for i:=0; i<len(TabletChain); i++ {
        if TabletChain[i].Name == Tablet.Name {
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
