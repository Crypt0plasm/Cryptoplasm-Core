package Cryptoplasm_Elliptic

import (
    "crypto/rand"
    "math/big"

    blake3 "github.com/Crypt0plasm/Cryptographic-Hash-Functions/Blake3"
)

func GenerateCryptoPlasmKeys () (PrivKey, PublicKey string){
    PrivKeyNumber := RandCurveE521ModuloPrime()
    PrivKey = PrivKeyNumber.Text(49)
    x,y := GetPublicKeyOnE521(PrivKeyNumber)
    PublicKeyHash := SevenFoldHash(x,y)
    PublicKey = ConvToLetters(PublicKeyHash)
    return PrivKey,PublicKey
}

func RandCurveE521ModuloPrime () *big.Int {
    RandomNumber,_ := rand.Int(rand.Reader,CurveE521Prime)
    return RandomNumber
}

func GetPublicKeyOnE521 (PrivKey *big.Int) (x,y *big.Int) {
    var (
        //start = time.Now()
        PrivKey49 		= PrivKey.Text(49)
        PrivKey49SliceRune 	= []rune(PrivKey49)
        PrivKey49SliceString 	= make([]string,len(PrivKey49))
        ZeroPoint		= CurveE521InfinityPoint
        Result			TecExtendedCoordinates
    )
    //start := time.Now()
    for i := 0; i < len(PrivKey49); i++ {
        PrivKey49SliceString[i] = string(PrivKey49SliceRune[i])
    }
    Result = ZeroPoint
    PrecMatrix := PrecomputingE521()
    for i := 0; i < len(PrivKey49SliceString); i++ {
        Character := PrivKey49SliceString[i]
        switch Character {
        //At last slice element, a 49x Point multiplication isn't executed.
        //49x Point Multiplication occurs on if i is not the last element in the slice
        case "0":
            Result = AdditionV2(Result,ZeroPoint,CurveE521Prime,CurveE521Constant)
            if i != len(PrivKey49SliceString) - 1 {
                Result = FortyNinerE521(Result)
            }
        case "1":
            Result = AdditionV2(Result,PrecMatrix[0][0],CurveE521Prime,CurveE521Constant)
            if i != len(PrivKey49SliceString) - 1 {
                Result = FortyNinerE521(Result)
            }
        case "2":
            Result = AdditionV2(Result,PrecMatrix[0][1],CurveE521Prime,CurveE521Constant)
            if i != len(PrivKey49SliceString) - 1 {
                Result = FortyNinerE521(Result)
            }
        case "3":
            Result = AdditionV2(Result,PrecMatrix[0][2],CurveE521Prime,CurveE521Constant)
            if i != len(PrivKey49SliceString) - 1 {
                Result = FortyNinerE521(Result)
            }
        case "4":
            Result = AdditionV2(Result,PrecMatrix[0][3],CurveE521Prime,CurveE521Constant)
            if i != len(PrivKey49SliceString) - 1 {
                Result = FortyNinerE521(Result)
            }
        case "5":
            Result = AdditionV2(Result,PrecMatrix[0][4],CurveE521Prime,CurveE521Constant)
            if i != len(PrivKey49SliceString) - 1 {
                Result = FortyNinerE521(Result)
            }
        case "6":
            Result = AdditionV2(Result,PrecMatrix[0][5],CurveE521Prime,CurveE521Constant)
            if i != len(PrivKey49SliceString) - 1 {
                Result = FortyNinerE521(Result)
            }
        case "7":
            Result = AdditionV2(Result,PrecMatrix[0][6],CurveE521Prime,CurveE521Constant)
            if i != len(PrivKey49SliceString) - 1 {
                Result = FortyNinerE521(Result)
            }
        case "8":
            Result = AdditionV2(Result,PrecMatrix[1][0],CurveE521Prime,CurveE521Constant)
            if i != len(PrivKey49SliceString) - 1 {
                Result = FortyNinerE521(Result)
            }
        case "9":
            Result = AdditionV2(Result,PrecMatrix[1][1],CurveE521Prime,CurveE521Constant)
            if i != len(PrivKey49SliceString) - 1 {
                Result = FortyNinerE521(Result)
            }
        case "a":
            Result = AdditionV2(Result,PrecMatrix[1][2],CurveE521Prime,CurveE521Constant)
            if i != len(PrivKey49SliceString) - 1 {
                Result = FortyNinerE521(Result)
            }
        case "b":
            Result = AdditionV2(Result,PrecMatrix[1][3],CurveE521Prime,CurveE521Constant)
            if i != len(PrivKey49SliceString) - 1 {
                Result = FortyNinerE521(Result)
            }
        case "c":
            Result = AdditionV2(Result,PrecMatrix[1][4],CurveE521Prime,CurveE521Constant)
            if i != len(PrivKey49SliceString) - 1 {
                Result = FortyNinerE521(Result)
            }
        case "d":
            Result = AdditionV2(Result,PrecMatrix[1][5],CurveE521Prime,CurveE521Constant)
            if i != len(PrivKey49SliceString) - 1 {
                Result = FortyNinerE521(Result)
            }
        case "e":
            Result = AdditionV2(Result,PrecMatrix[1][6],CurveE521Prime,CurveE521Constant)
            if i != len(PrivKey49SliceString) - 1 {
                Result = FortyNinerE521(Result)
            }
        case "f":
            Result = AdditionV2(Result,PrecMatrix[2][0],CurveE521Prime,CurveE521Constant)
            if i != len(PrivKey49SliceString) - 1 {
                Result = FortyNinerE521(Result)
            }
        case "g":
            Result = AdditionV2(Result,PrecMatrix[2][1],CurveE521Prime,CurveE521Constant)
            if i != len(PrivKey49SliceString) - 1 {
                Result = FortyNinerE521(Result)
            }
        case "h":
            Result = AdditionV2(Result,PrecMatrix[2][2],CurveE521Prime,CurveE521Constant)
            if i != len(PrivKey49SliceString) - 1 {
                Result = FortyNinerE521(Result)
            }
        case "i":
            Result = AdditionV2(Result,PrecMatrix[2][3],CurveE521Prime,CurveE521Constant)
            if i != len(PrivKey49SliceString) - 1 {
                Result = FortyNinerE521(Result)
            }
        case "j":
            Result = AdditionV2(Result,PrecMatrix[2][4],CurveE521Prime,CurveE521Constant)
            if i != len(PrivKey49SliceString) - 1 {
                Result = FortyNinerE521(Result)
            }
        case "k":
            Result = AdditionV2(Result,PrecMatrix[2][5],CurveE521Prime,CurveE521Constant)
            if i != len(PrivKey49SliceString) - 1 {
                Result = FortyNinerE521(Result)
            }
        case "l":
            Result = AdditionV2(Result,PrecMatrix[2][6],CurveE521Prime,CurveE521Constant)
            if i != len(PrivKey49SliceString) - 1 {
                Result = FortyNinerE521(Result)
            }
        case "m":
            Result = AdditionV2(Result,PrecMatrix[3][0],CurveE521Prime,CurveE521Constant)
            if i != len(PrivKey49SliceString) - 1 {
                Result = FortyNinerE521(Result)
            }
        case "n":
            Result = AdditionV2(Result,PrecMatrix[3][1],CurveE521Prime,CurveE521Constant)
            if i != len(PrivKey49SliceString) - 1 {
                Result = FortyNinerE521(Result)
            }
        case "o":
            Result = AdditionV2(Result,PrecMatrix[3][2],CurveE521Prime,CurveE521Constant)
            if i != len(PrivKey49SliceString) - 1 {
                Result = FortyNinerE521(Result)
            }
        case "p":
            Result = AdditionV2(Result,PrecMatrix[3][3],CurveE521Prime,CurveE521Constant)
            if i != len(PrivKey49SliceString) - 1 {
                Result = FortyNinerE521(Result)
            }
        case "q":
            Result = AdditionV2(Result,PrecMatrix[3][4],CurveE521Prime,CurveE521Constant)
            if i != len(PrivKey49SliceString) - 1 {
                Result = FortyNinerE521(Result)
            }
        case "r":
            Result = AdditionV2(Result,PrecMatrix[3][5],CurveE521Prime,CurveE521Constant)
            if i != len(PrivKey49SliceString) - 1 {
                Result = FortyNinerE521(Result)
            }
        case "s":
            Result = AdditionV2(Result,PrecMatrix[3][6],CurveE521Prime,CurveE521Constant)
            if i != len(PrivKey49SliceString) - 1 {
                Result = FortyNinerE521(Result)
            }
        case "t":
            Result = AdditionV2(Result,PrecMatrix[4][0],CurveE521Prime,CurveE521Constant)
            if i != len(PrivKey49SliceString) - 1 {
                Result = FortyNinerE521(Result)
            }
        case "u":
            Result = AdditionV2(Result,PrecMatrix[4][1],CurveE521Prime,CurveE521Constant)
            if i != len(PrivKey49SliceString) - 1 {
                Result = FortyNinerE521(Result)
            }
        case "v":
            Result = AdditionV2(Result,PrecMatrix[4][2],CurveE521Prime,CurveE521Constant)
            if i != len(PrivKey49SliceString) - 1 {
                Result = FortyNinerE521(Result)
            }
        case "w":
            Result = AdditionV2(Result,PrecMatrix[4][3],CurveE521Prime,CurveE521Constant)
            if i != len(PrivKey49SliceString) - 1 {
                Result = FortyNinerE521(Result)
            }
        case "x":
            Result = AdditionV2(Result,PrecMatrix[4][4],CurveE521Prime,CurveE521Constant)
            if i != len(PrivKey49SliceString) - 1 {
                Result = FortyNinerE521(Result)
            }
        case "y":
            Result = AdditionV2(Result,PrecMatrix[4][5],CurveE521Prime,CurveE521Constant)
            if i != len(PrivKey49SliceString) - 1 {
                Result = FortyNinerE521(Result)
            }
        case "z":
            Result = AdditionV2(Result,PrecMatrix[4][6],CurveE521Prime,CurveE521Constant)
            if i != len(PrivKey49SliceString) - 1 {
                Result = FortyNinerE521(Result)
            }
        case "A":
            Result = AdditionV2(Result,PrecMatrix[5][0],CurveE521Prime,CurveE521Constant)
            if i != len(PrivKey49SliceString) - 1 {
                Result = FortyNinerE521(Result)
            }
        case "B":
            Result = AdditionV2(Result,PrecMatrix[5][1],CurveE521Prime,CurveE521Constant)
            if i != len(PrivKey49SliceString) - 1 {
                Result = FortyNinerE521(Result)
            }
        case "C":
            Result = AdditionV2(Result,PrecMatrix[5][2],CurveE521Prime,CurveE521Constant)
            if i != len(PrivKey49SliceString) - 1 {
                Result = FortyNinerE521(Result)
            }
        case "D":
            Result = AdditionV2(Result,PrecMatrix[5][3],CurveE521Prime,CurveE521Constant)
            if i != len(PrivKey49SliceString) - 1 {
                Result = FortyNinerE521(Result)
            }
        case "E":
            Result = AdditionV2(Result,PrecMatrix[5][4],CurveE521Prime,CurveE521Constant)
            if i != len(PrivKey49SliceString) - 1 {
                Result = FortyNinerE521(Result)
            }
        case "F":
            Result = AdditionV2(Result,PrecMatrix[5][5],CurveE521Prime,CurveE521Constant)
            if i != len(PrivKey49SliceString) - 1 {
                Result = FortyNinerE521(Result)
            }
        case "G":
            Result = AdditionV2(Result,PrecMatrix[5][6],CurveE521Prime,CurveE521Constant)
            if i != len(PrivKey49SliceString) - 1 {
                Result = FortyNinerE521(Result)
            }
        case "H":
            Result = AdditionV2(Result,PrecMatrix[6][0],CurveE521Prime,CurveE521Constant)
            if i != len(PrivKey49SliceString) - 1 {
                Result = FortyNinerE521(Result)
            }
        case "I":
            Result = AdditionV2(Result,PrecMatrix[6][1],CurveE521Prime,CurveE521Constant)
            if i != len(PrivKey49SliceString) - 1 {
                Result = FortyNinerE521(Result)
            }
        case "J":
            Result = AdditionV2(Result,PrecMatrix[6][2],CurveE521Prime,CurveE521Constant)
            if i != len(PrivKey49SliceString) - 1 {
                Result = FortyNinerE521(Result)
            }
        case "K":
            Result = AdditionV2(Result,PrecMatrix[6][3],CurveE521Prime,CurveE521Constant)
            if i != len(PrivKey49SliceString) - 1 {
                Result = FortyNinerE521(Result)
            }
        case "L":
            Result = AdditionV2(Result,PrecMatrix[6][4],CurveE521Prime,CurveE521Constant)
            if i != len(PrivKey49SliceString) - 1 {
                Result = FortyNinerE521(Result)
            }
        case "M":
            Result = AdditionV2(Result,PrecMatrix[6][5],CurveE521Prime,CurveE521Constant)
            if i != len(PrivKey49SliceString) - 1 {
                Result = FortyNinerE521(Result)
            }
        }

    }
    ResultAff := TecExtended2TecAffine(Result,CurveE521Prime)
    x = ResultAff.AX
    y = ResultAff.AY

    //elapsed := time.Since(start)
    //fmt.Println("")
    //fmt.Println("Computing PublicKey points took:", elapsed)
    return x, y
}

func SevenFoldHash (x,y *big.Int)[]byte {
    XString := x.String()
    YString := y.String()
    LongString := XString+YString
    LongStringToByteSlice := []byte(LongString)

    S1 := blake3.SumCustom(LongStringToByteSlice,160)
    S2 := blake3.SumCustom(S1,160)
    S3 := blake3.SumCustom(S2,160)
    S4 := blake3.SumCustom(S3,160)
    S5 := blake3.SumCustom(S4,160)
    S6 := blake3.SumCustom(S5,160)
    S7 := blake3.SumCustom(S6,160)

    return S7
}

func ConvToLetters (hash []byte) string {
    var (
    	result string
        SliceStr []string
    )

    Matrix := CharacterMatrix()
    for i := 0; i < len(hash); i++ {
        Value := hash[i]
        switch Value {
        case 0:SliceStr = append(SliceStr,string(Matrix[0][0]))
        case 1:SliceStr = append(SliceStr,string(Matrix[0][1]))
        case 2:SliceStr = append(SliceStr,string(Matrix[0][2]))
        case 3:SliceStr = append(SliceStr,string(Matrix[0][3]))
        case 4:SliceStr = append(SliceStr,string(Matrix[0][4]))
        case 5:SliceStr = append(SliceStr,string(Matrix[0][5]))
        case 6:SliceStr = append(SliceStr,string(Matrix[0][6]))
        case 7:SliceStr = append(SliceStr,string(Matrix[0][7]))
        case 8:SliceStr = append(SliceStr,string(Matrix[0][8]))
        case 9:SliceStr = append(SliceStr,string(Matrix[0][9]))
        case 10:SliceStr = append(SliceStr,string(Matrix[0][10]))
        case 11:SliceStr = append(SliceStr,string(Matrix[0][11]))
        case 12:SliceStr = append(SliceStr,string(Matrix[0][12]))
        case 13:SliceStr = append(SliceStr,string(Matrix[0][13]))
        case 14:SliceStr = append(SliceStr,string(Matrix[0][14]))
        case 15:SliceStr = append(SliceStr,string(Matrix[0][15]))

        case 16:SliceStr = append(SliceStr,string(Matrix[1][0]))
        case 17:SliceStr = append(SliceStr,string(Matrix[1][1]))
        case 18:SliceStr = append(SliceStr,string(Matrix[1][2]))
        case 19:SliceStr = append(SliceStr,string(Matrix[1][3]))
        case 20:SliceStr = append(SliceStr,string(Matrix[1][4]))
        case 21:SliceStr = append(SliceStr,string(Matrix[1][5]))
        case 22:SliceStr = append(SliceStr,string(Matrix[1][6]))
        case 23:SliceStr = append(SliceStr,string(Matrix[1][7]))
        case 24:SliceStr = append(SliceStr,string(Matrix[1][8]))
        case 25:SliceStr = append(SliceStr,string(Matrix[1][9]))
        case 26:SliceStr = append(SliceStr,string(Matrix[1][10]))
        case 27:SliceStr = append(SliceStr,string(Matrix[1][11]))
        case 28:SliceStr = append(SliceStr,string(Matrix[1][12]))
        case 29:SliceStr = append(SliceStr,string(Matrix[1][13]))
        case 30:SliceStr = append(SliceStr,string(Matrix[1][14]))
        case 31:SliceStr = append(SliceStr,string(Matrix[1][15]))

        case 32:SliceStr = append(SliceStr,string(Matrix[2][0]))
        case 33:SliceStr = append(SliceStr,string(Matrix[2][1]))
        case 34:SliceStr = append(SliceStr,string(Matrix[2][2]))
        case 35:SliceStr = append(SliceStr,string(Matrix[2][3]))
        case 36:SliceStr = append(SliceStr,string(Matrix[2][4]))
        case 37:SliceStr = append(SliceStr,string(Matrix[2][5]))
        case 38:SliceStr = append(SliceStr,string(Matrix[2][6]))
        case 39:SliceStr = append(SliceStr,string(Matrix[2][7]))
        case 40:SliceStr = append(SliceStr,string(Matrix[2][8]))
        case 41:SliceStr = append(SliceStr,string(Matrix[2][9]))
        case 42:SliceStr = append(SliceStr,string(Matrix[2][10]))
        case 43:SliceStr = append(SliceStr,string(Matrix[2][11]))
        case 44:SliceStr = append(SliceStr,string(Matrix[2][12]))
        case 45:SliceStr = append(SliceStr,string(Matrix[2][13]))
        case 46:SliceStr = append(SliceStr,string(Matrix[2][14]))
        case 47:SliceStr = append(SliceStr,string(Matrix[2][15]))

        case 48:SliceStr = append(SliceStr,string(Matrix[3][0]))
        case 49:SliceStr = append(SliceStr,string(Matrix[3][1]))
        case 50:SliceStr = append(SliceStr,string(Matrix[3][2]))
        case 51:SliceStr = append(SliceStr,string(Matrix[3][3]))
        case 52:SliceStr = append(SliceStr,string(Matrix[3][4]))
        case 53:SliceStr = append(SliceStr,string(Matrix[3][5]))
        case 54:SliceStr = append(SliceStr,string(Matrix[3][6]))
        case 55:SliceStr = append(SliceStr,string(Matrix[3][7]))
        case 56:SliceStr = append(SliceStr,string(Matrix[3][8]))
        case 57:SliceStr = append(SliceStr,string(Matrix[3][9]))
        case 58:SliceStr = append(SliceStr,string(Matrix[3][10]))
        case 59:SliceStr = append(SliceStr,string(Matrix[3][11]))
        case 60:SliceStr = append(SliceStr,string(Matrix[3][12]))
        case 61:SliceStr = append(SliceStr,string(Matrix[3][13]))
        case 62:SliceStr = append(SliceStr,string(Matrix[3][14]))
        case 63:SliceStr = append(SliceStr,string(Matrix[3][15]))

        case 64:SliceStr = append(SliceStr,string(Matrix[4][0]))
        case 65:SliceStr = append(SliceStr,string(Matrix[4][1]))
        case 66:SliceStr = append(SliceStr,string(Matrix[4][2]))
        case 67:SliceStr = append(SliceStr,string(Matrix[4][3]))
        case 68:SliceStr = append(SliceStr,string(Matrix[4][4]))
        case 69:SliceStr = append(SliceStr,string(Matrix[4][5]))
        case 70:SliceStr = append(SliceStr,string(Matrix[4][6]))
        case 71:SliceStr = append(SliceStr,string(Matrix[4][7]))
        case 72:SliceStr = append(SliceStr,string(Matrix[4][8]))
        case 73:SliceStr = append(SliceStr,string(Matrix[4][9]))
        case 74:SliceStr = append(SliceStr,string(Matrix[4][10]))
        case 75:SliceStr = append(SliceStr,string(Matrix[4][11]))
        case 76:SliceStr = append(SliceStr,string(Matrix[4][12]))
        case 77:SliceStr = append(SliceStr,string(Matrix[4][13]))
        case 78:SliceStr = append(SliceStr,string(Matrix[4][14]))
        case 79:SliceStr = append(SliceStr,string(Matrix[4][15]))

        case 80:SliceStr = append(SliceStr,string(Matrix[5][0]))
        case 81:SliceStr = append(SliceStr,string(Matrix[5][1]))
        case 82:SliceStr = append(SliceStr,string(Matrix[5][2]))
        case 83:SliceStr = append(SliceStr,string(Matrix[5][3]))
        case 84:SliceStr = append(SliceStr,string(Matrix[5][4]))
        case 85:SliceStr = append(SliceStr,string(Matrix[5][5]))
        case 86:SliceStr = append(SliceStr,string(Matrix[5][6]))
        case 87:SliceStr = append(SliceStr,string(Matrix[5][7]))
        case 88:SliceStr = append(SliceStr,string(Matrix[5][8]))
        case 89:SliceStr = append(SliceStr,string(Matrix[5][9]))
        case 90:SliceStr = append(SliceStr,string(Matrix[5][10]))
        case 91:SliceStr = append(SliceStr,string(Matrix[5][11]))
        case 92:SliceStr = append(SliceStr,string(Matrix[5][12]))
        case 93:SliceStr = append(SliceStr,string(Matrix[5][13]))
        case 94:SliceStr = append(SliceStr,string(Matrix[5][14]))
        case 95:SliceStr = append(SliceStr,string(Matrix[5][15]))

        case 96:SliceStr = append(SliceStr,string(Matrix[6][0]))
        case 97:SliceStr = append(SliceStr,string(Matrix[6][1]))
        case 98:SliceStr = append(SliceStr,string(Matrix[6][2]))
        case 99:SliceStr = append(SliceStr,string(Matrix[6][3]))
        case 100:SliceStr = append(SliceStr,string(Matrix[6][4]))
        case 101:SliceStr = append(SliceStr,string(Matrix[6][5]))
        case 102:SliceStr = append(SliceStr,string(Matrix[6][6]))
        case 103:SliceStr = append(SliceStr,string(Matrix[6][7]))
        case 104:SliceStr = append(SliceStr,string(Matrix[6][8]))
        case 105:SliceStr = append(SliceStr,string(Matrix[6][9]))
        case 106:SliceStr = append(SliceStr,string(Matrix[6][10]))
        case 107:SliceStr = append(SliceStr,string(Matrix[6][11]))
        case 108:SliceStr = append(SliceStr,string(Matrix[6][12]))
        case 109:SliceStr = append(SliceStr,string(Matrix[6][13]))
        case 110:SliceStr = append(SliceStr,string(Matrix[6][14]))
        case 111:SliceStr = append(SliceStr,string(Matrix[6][15]))

        case 112:SliceStr = append(SliceStr,string(Matrix[7][0]))
        case 113:SliceStr = append(SliceStr,string(Matrix[7][1]))
        case 114:SliceStr = append(SliceStr,string(Matrix[7][2]))
        case 115:SliceStr = append(SliceStr,string(Matrix[7][3]))
        case 116:SliceStr = append(SliceStr,string(Matrix[7][4]))
        case 117:SliceStr = append(SliceStr,string(Matrix[7][5]))
        case 118:SliceStr = append(SliceStr,string(Matrix[7][6]))
        case 119:SliceStr = append(SliceStr,string(Matrix[7][7]))
        case 120:SliceStr = append(SliceStr,string(Matrix[7][8]))
        case 121:SliceStr = append(SliceStr,string(Matrix[7][9]))
        case 122:SliceStr = append(SliceStr,string(Matrix[7][10]))
        case 123:SliceStr = append(SliceStr,string(Matrix[7][11]))
        case 124:SliceStr = append(SliceStr,string(Matrix[7][12]))
        case 125:SliceStr = append(SliceStr,string(Matrix[7][13]))
        case 126:SliceStr = append(SliceStr,string(Matrix[7][14]))
        case 127:SliceStr = append(SliceStr,string(Matrix[7][15]))

        case 128:SliceStr = append(SliceStr,string(Matrix[8][0]))
        case 129:SliceStr = append(SliceStr,string(Matrix[8][1]))
        case 130:SliceStr = append(SliceStr,string(Matrix[8][2]))
        case 131:SliceStr = append(SliceStr,string(Matrix[8][3]))
        case 132:SliceStr = append(SliceStr,string(Matrix[8][4]))
        case 133:SliceStr = append(SliceStr,string(Matrix[8][5]))
        case 134:SliceStr = append(SliceStr,string(Matrix[8][6]))
        case 135:SliceStr = append(SliceStr,string(Matrix[8][7]))
        case 136:SliceStr = append(SliceStr,string(Matrix[8][8]))
        case 137:SliceStr = append(SliceStr,string(Matrix[8][9]))
        case 138:SliceStr = append(SliceStr,string(Matrix[8][10]))
        case 139:SliceStr = append(SliceStr,string(Matrix[8][11]))
        case 140:SliceStr = append(SliceStr,string(Matrix[8][12]))
        case 141:SliceStr = append(SliceStr,string(Matrix[8][13]))
        case 142:SliceStr = append(SliceStr,string(Matrix[8][14]))
        case 143:SliceStr = append(SliceStr,string(Matrix[8][15]))

        case 144:SliceStr = append(SliceStr,string(Matrix[9][0]))
        case 145:SliceStr = append(SliceStr,string(Matrix[9][1]))
        case 146:SliceStr = append(SliceStr,string(Matrix[9][2]))
        case 147:SliceStr = append(SliceStr,string(Matrix[9][3]))
        case 148:SliceStr = append(SliceStr,string(Matrix[9][4]))
        case 149:SliceStr = append(SliceStr,string(Matrix[9][5]))
        case 150:SliceStr = append(SliceStr,string(Matrix[9][6]))
        case 151:SliceStr = append(SliceStr,string(Matrix[9][7]))
        case 152:SliceStr = append(SliceStr,string(Matrix[9][8]))
        case 153:SliceStr = append(SliceStr,string(Matrix[9][9]))
        case 154:SliceStr = append(SliceStr,string(Matrix[9][10]))
        case 155:SliceStr = append(SliceStr,string(Matrix[9][11]))
        case 156:SliceStr = append(SliceStr,string(Matrix[9][12]))
        case 157:SliceStr = append(SliceStr,string(Matrix[9][13]))
        case 158:SliceStr = append(SliceStr,string(Matrix[9][14]))
        case 159:SliceStr = append(SliceStr,string(Matrix[9][15]))

        case 160:SliceStr = append(SliceStr,string(Matrix[10][0]))
        case 161:SliceStr = append(SliceStr,string(Matrix[10][1]))
        case 162:SliceStr = append(SliceStr,string(Matrix[10][2]))
        case 163:SliceStr = append(SliceStr,string(Matrix[10][3]))
        case 164:SliceStr = append(SliceStr,string(Matrix[10][4]))
        case 165:SliceStr = append(SliceStr,string(Matrix[10][5]))
        case 166:SliceStr = append(SliceStr,string(Matrix[10][6]))
        case 167:SliceStr = append(SliceStr,string(Matrix[10][7]))
        case 168:SliceStr = append(SliceStr,string(Matrix[10][8]))
        case 169:SliceStr = append(SliceStr,string(Matrix[10][9]))
        case 170:SliceStr = append(SliceStr,string(Matrix[10][10]))
        case 171:SliceStr = append(SliceStr,string(Matrix[10][11]))
        case 172:SliceStr = append(SliceStr,string(Matrix[10][12]))
        case 173:SliceStr = append(SliceStr,string(Matrix[10][13]))
        case 174:SliceStr = append(SliceStr,string(Matrix[10][14]))
        case 175:SliceStr = append(SliceStr,string(Matrix[10][15]))

        case 176:SliceStr = append(SliceStr,string(Matrix[11][0]))
        case 177:SliceStr = append(SliceStr,string(Matrix[11][1]))
        case 178:SliceStr = append(SliceStr,string(Matrix[11][2]))
        case 179:SliceStr = append(SliceStr,string(Matrix[11][3]))
        case 180:SliceStr = append(SliceStr,string(Matrix[11][4]))
        case 181:SliceStr = append(SliceStr,string(Matrix[11][5]))
        case 182:SliceStr = append(SliceStr,string(Matrix[11][6]))
        case 183:SliceStr = append(SliceStr,string(Matrix[11][7]))
        case 184:SliceStr = append(SliceStr,string(Matrix[11][8]))
        case 185:SliceStr = append(SliceStr,string(Matrix[11][9]))
        case 186:SliceStr = append(SliceStr,string(Matrix[11][10]))
        case 187:SliceStr = append(SliceStr,string(Matrix[11][11]))
        case 188:SliceStr = append(SliceStr,string(Matrix[11][12]))
        case 189:SliceStr = append(SliceStr,string(Matrix[11][13]))
        case 190:SliceStr = append(SliceStr,string(Matrix[11][14]))
        case 191:SliceStr = append(SliceStr,string(Matrix[11][15]))

        case 192:SliceStr = append(SliceStr,string(Matrix[12][0]))
        case 193:SliceStr = append(SliceStr,string(Matrix[12][1]))
        case 194:SliceStr = append(SliceStr,string(Matrix[12][2]))
        case 195:SliceStr = append(SliceStr,string(Matrix[12][3]))
        case 196:SliceStr = append(SliceStr,string(Matrix[12][4]))
        case 197:SliceStr = append(SliceStr,string(Matrix[12][5]))
        case 198:SliceStr = append(SliceStr,string(Matrix[12][6]))
        case 199:SliceStr = append(SliceStr,string(Matrix[12][7]))
        case 200:SliceStr = append(SliceStr,string(Matrix[12][8]))
        case 201:SliceStr = append(SliceStr,string(Matrix[12][9]))
        case 202:SliceStr = append(SliceStr,string(Matrix[12][10]))
        case 203:SliceStr = append(SliceStr,string(Matrix[12][11]))
        case 204:SliceStr = append(SliceStr,string(Matrix[12][12]))
        case 205:SliceStr = append(SliceStr,string(Matrix[12][13]))
        case 206:SliceStr = append(SliceStr,string(Matrix[12][14]))
        case 207:SliceStr = append(SliceStr,string(Matrix[12][15]))

        case 208:SliceStr = append(SliceStr,string(Matrix[13][0]))
        case 209:SliceStr = append(SliceStr,string(Matrix[13][1]))
        case 210:SliceStr = append(SliceStr,string(Matrix[13][2]))
        case 211:SliceStr = append(SliceStr,string(Matrix[13][3]))
        case 212:SliceStr = append(SliceStr,string(Matrix[13][4]))
        case 213:SliceStr = append(SliceStr,string(Matrix[13][5]))
        case 214:SliceStr = append(SliceStr,string(Matrix[13][6]))
        case 215:SliceStr = append(SliceStr,string(Matrix[13][7]))
        case 216:SliceStr = append(SliceStr,string(Matrix[13][8]))
        case 217:SliceStr = append(SliceStr,string(Matrix[13][9]))
        case 218:SliceStr = append(SliceStr,string(Matrix[13][10]))
        case 219:SliceStr = append(SliceStr,string(Matrix[13][11]))
        case 220:SliceStr = append(SliceStr,string(Matrix[13][12]))
        case 221:SliceStr = append(SliceStr,string(Matrix[13][13]))
        case 222:SliceStr = append(SliceStr,string(Matrix[13][14]))
        case 223:SliceStr = append(SliceStr,string(Matrix[13][15]))

        case 224:SliceStr = append(SliceStr,string(Matrix[14][0]))
        case 225:SliceStr = append(SliceStr,string(Matrix[14][1]))
        case 226:SliceStr = append(SliceStr,string(Matrix[14][2]))
        case 227:SliceStr = append(SliceStr,string(Matrix[14][3]))
        case 228:SliceStr = append(SliceStr,string(Matrix[14][4]))
        case 229:SliceStr = append(SliceStr,string(Matrix[14][5]))
        case 230:SliceStr = append(SliceStr,string(Matrix[14][6]))
        case 231:SliceStr = append(SliceStr,string(Matrix[14][7]))
        case 232:SliceStr = append(SliceStr,string(Matrix[14][8]))
        case 233:SliceStr = append(SliceStr,string(Matrix[14][9]))
        case 234:SliceStr = append(SliceStr,string(Matrix[14][10]))
        case 235:SliceStr = append(SliceStr,string(Matrix[14][11]))
        case 236:SliceStr = append(SliceStr,string(Matrix[14][12]))
        case 237:SliceStr = append(SliceStr,string(Matrix[14][13]))
        case 238:SliceStr = append(SliceStr,string(Matrix[14][14]))
        case 239:SliceStr = append(SliceStr,string(Matrix[14][15]))

        case 240:SliceStr = append(SliceStr,string(Matrix[15][0]))
        case 241:SliceStr = append(SliceStr,string(Matrix[15][1]))
        case 242:SliceStr = append(SliceStr,string(Matrix[15][2]))
        case 243:SliceStr = append(SliceStr,string(Matrix[15][3]))
        case 244:SliceStr = append(SliceStr,string(Matrix[15][4]))
        case 245:SliceStr = append(SliceStr,string(Matrix[15][5]))
        case 246:SliceStr = append(SliceStr,string(Matrix[15][6]))
        case 247:SliceStr = append(SliceStr,string(Matrix[15][7]))
        case 248:SliceStr = append(SliceStr,string(Matrix[15][8]))
        case 249:SliceStr = append(SliceStr,string(Matrix[15][9]))
        case 250:SliceStr = append(SliceStr,string(Matrix[15][10]))
        case 251:SliceStr = append(SliceStr,string(Matrix[15][11]))
        case 252:SliceStr = append(SliceStr,string(Matrix[15][12]))
        case 253:SliceStr = append(SliceStr,string(Matrix[15][13]))
        case 254:SliceStr = append(SliceStr,string(Matrix[15][14]))
        case 255:SliceStr = append(SliceStr,string(Matrix[15][15]))
        }
    }

    //Converting the Slice of Strings to a String as final step
    for i := 0; i < len(SliceStr); i++ {
        result = result + SliceStr[i]
    }

    return result
}

func CharacterMatrix () [16][16]rune {
    //Digits Block 10 runes
    C000 := '0'     //U+0030    Digit Zero
    C001 := '1'     //U+0031    Digit One
    C002 := '2'     //U+0032    Digit Two
    C003 := '3'     //U+0033    Digit Three
    C004 := '4'     //U+0034    Digit Four
    C005 := '5'     //U+0035    Digit Five
    C006 := '6'     //U+0036    Digit Six
    C007 := '7'     //U+0037    Digit Seven
    C008 := '8'     //U+0038    Digit Eight
    C009 := '9'     //U+0039    Digit Nine

    //Currencies Block 14 runes
    C010 := 'Ѻ'     //U+047A    Cyrillic Capital Letter Round Omega aka Cryptoplasm
    C011 := '₿'     //U+20BF    Bitcoin Sign
    C012 := 'Ł'     //U+0141    Latin Capital Letter L with Stroke aka Litecoin
    C013 := '$'     //U+0024    Dollar Sign
    C014 := '¢'     //U+00A2    Cent Sign
    C015 := '€'     //U+20AC    Euro Sign
    C016 := '£'     //U+00A3    Pound Sign
    C017 := '¥'     //U+00A5    Yen Sign
    C018 := '₹'     //U+20B9    Indian Rupee Sign
    C019 := '₽'     //U+20BD    Ruble Sign
    C020 := '₺'     //U+20BA    Turkish Lira Sign
    C021 := '₱'     //U+20B1    Peso Sign
    C022 := '₳'     //U+20B1    Austral Sign
    C023 := '₣'     //U+20A3    French Franc Sign


    //Lonely Small Letters 36 Runes
    C024 := 'α'     //U+03B1    Greek Small Letter Alpha
    C025 := 'b'     //U+0062    Latin Small Letter B
    C026 := 'β'     //U+03B2    Greek Small Letter Beta
    C027 := 'χ'     //U+03C7    Greek Small Letter Chi
    C028 := 'д'     //U+0434    Cyrillic Small Letter De
    C029 := 'ε'     //U+03B5    Greek Small Letter Epsilon
    C030 := 'γ'     //U+03B3    Greek Small Letter Gamma
    C031 := 'г'     //U+0433    Cyrillic Small Letter Ghe
    C032 := 'h'     //U+0068    Latin Small Letter H
    C033 := 'η'     //U+03B7    Greek Small Letter Eta
    C034 := 'ι'     //U+03B9    Greek Small Letter Iota
    C035 := 'и'     //U+0438    Cyrillic Small Letter I
    C036 := 'й'     //U+0439    Cyrillic Small Letter Short I
    C037 := 'ж'     //U+0436    Cyrillic Small Letter Zhe
    C038 := 'k'     //U+006B    Latin Small Letter K
    C039 := 'л'     //U+043B    Cyrillic Small Letter El
    C040 := 'm'     //U+20B1    Peso Sign
    C041 := 'μ'     //U+03BC    Greek Small Letter Mu
    C042 := 'м'     //U+043C    Cyrillic Small Letter Em
    C043 := 'ν'     //U+03BD    Greek Small Letter Nu
    C044 := 'н'     //U+043D    Cyrillic Small Letter En
    C045 := 'π'     //U+03C0    Greek Small Letter Pi
    C046 := 'ψ'     //U+03C8    Greek Small Letter Psi
    C047 := 'п'     //U+043F    Cyrillic Small Letter Pe
    C048 := 'ρ'     //U+03C1    Greek Small Letter Rho
    C049 := 'ς'     //U+03C2    Greek Small Letter Final Sigma
    C050 := 'ш'     //U+0448    Cyrillic Small Letter Sha
    C051 := 'щ'     //U+0449    Cyrillic Small Letter Shcha
    C052 := 'τ'     //U+03C4    Greek Small Letter Tau
    C053 := 'ц'     //U+0446    Cyrillic Small Letter Tse
    C054 := 'в'     //U+0432    Cyrillic Small Letter Ve
    C055 := 'ю'     //U+044E    Cyrillic Small Letter Yu
    C056 := 'я'     //U+044F    Cyrillic Small Letter Ya
    C057 := 'ß'     //U+00DF    Latin Small Letter Sharp S
    C058 := 'ъ'     //U+044A    Cyrillic Small Letter Hard Sign
    C059 := 'ь'     //U+044C    Cyrillic Small Letter Soft Sign

    //Capital-Small Letters pairs 2*98=196 Runes

    C060 := 'A'     //U+20B1    Peso Sign
    C061 := 'a'     //U+20B1    Peso Sign
    C062 := 'Å'     //U+20B1    Peso Sign
    C063 := 'å'     //U+20B1    Peso Sign
    C064 := 'Ä'     //U+20B1    Peso Sign
    C065 := 'ä'     //U+20B1    Peso Sign
    C066 := 'Â'     //U+20B1    Peso Sign
    C067 := 'â'     //U+20B1    Peso Sign
    C068 := 'Á'     //U+20B1    Peso Sign
    C069 := 'á'     //U+20B1    Peso Sign
    C070 := 'À'     //U+20B1    Peso Sign
    C071 := 'à'     //U+20B1    Peso Sign
    C072 := 'Ă'     //U+20B1    Peso Sign
    C073 := 'ă'     //U+20B1    Peso Sign
    C074 := 'Ã'     //U+20B1    Peso Sign
    C075 := 'ã'     //U+20B1    Peso Sign
    C076 := 'Ā'     //U+20B1    Peso Sign
    C077 := 'ā'     //U+20B1    Peso Sign
    C078 := 'Æ'     //U+20B1    Peso Sign
    C079 := 'æ'     //U+20B1    Peso Sign
    C080 := 'Б'     //U+20B1    Peso Sign
    C081 := 'б'     //U+20B1    Peso Sign
    C082 := 'C'     //U+20B1    Peso Sign
    C083 := 'c'     //U+20B1    Peso Sign
    C084 := 'Ċ'     //U+20B1    Peso Sign
    C085 := 'ċ'     //U+20B1    Peso Sign
    C086 := 'Č'     //U+20B1    Peso Sign
    C087 := 'č'     //U+20B1    Peso Sign
    C088 := 'Ĉ'     //U+20B1    Peso Sign
    C089 := 'ĉ'     //U+20B1    Peso Sign
    C090 := 'Ć'     //U+20B1    Peso Sign
    C091 := 'ć'     //U+20B1    Peso Sign
    C092 := 'Ç'     //U+20B1    Peso Sign
    C093 := 'ç'     //U+20B1    Peso Sign
    C094 := 'Ч'     //U+0427    Cyrillic Capital Letter Che
    C095 := 'ч'     //U+20B1    Peso Sign
    C096 := 'D'     //U+20B1    Peso Sign
    C097 := 'd'     //U+20B1    Peso Sign
    C098 := 'Δ'     //U+0394    Greek Capital Letter Delta
    C099 := 'δ'     //U+20B1    Peso Sign
    C100 := 'E'     //U+20B1    Peso Sign
    C101 := 'e'     //U+20B1    Peso Sign
    C102 := 'Ė'     //U+20B1    Peso Sign
    C103 := 'ė'     //U+20B1    Peso Sign
    C104 := 'Ë'     //U+20B1    Peso Sign
    C105 := 'ë'     //U+20B1    Peso Sign
    C106 := 'Ě'     //U+20B1    Peso Sign
    C107 := 'ě'     //U+20B1    Peso Sign
    C108 := 'Ê'     //U+20B1    Peso Sign
    C109 := 'ê'     //U+20B1    Peso Sign
    C110 := 'É'     //U+20B1    Peso Sign
    C111 := 'é'     //U+20B1    Peso Sign
    C112 := 'È'     //U+20B1    Peso Sign
    C113 := 'è'     //U+20B1    Peso Sign
    C114 := 'Ĕ'     //U+20B1    Peso Sign
    C115 := 'ĕ'     //U+20B1    Peso Sign
    C116 := 'Ē'     //U+20B1    Peso Sign
    C117 := 'ē'     //U+20B1    Peso Sign
    C118 := 'Э'     //U+042D    Cyrillic Capital Letter
    C119 := 'э'     //U+20B1    Peso Sign
    C120 := 'F'     //U+20B1    Peso Sign
    C121 := 'f'     //U+20B1    Peso Sign
    C122 := 'Φ'     //U+03A6    Greek Capital Letter Phi
    C123 := 'φ'     //U+03C6    Greek Small Letter Phi
    C124 := 'G'     //U+20B1    Peso Sign
    C125 := 'g'     //U+20B1    Peso Sign
    C126 := 'Ġ'     //U+20B1    Peso Sign
    C127 := 'ġ'     //U+20B1    Peso Sign
    C128 := 'Ĝ'     //U+20B1    Peso Sign
    C129 := 'ĝ'     //U+20B1    Peso Sign
    C130 := 'Ğ'     //U+20B1    Peso Sign
    C131 := 'ğ'     //U+20B1    Peso Sign
    C132 := 'Ĥ'     //U+20B1    Peso Sign
    C133 := 'ĥ'     //U+20B1    Peso Sign
    C134 := 'I'     //U+20B1    Peso Sign
    C135 := 'i'     //U+20B1    Peso Sign
    C136 := 'Ï'     //U+20B1    Peso Sign
    C137 := 'ï'     //U+20B1    Peso Sign
    C138 := 'Î'     //U+20B1    Peso Sign
    C139 := 'î'     //U+20B1    Peso Sign
    C140 := 'Í'     //U+20B1    Peso Sign
    C141 := 'í'     //U+20B1    Peso Sign
    C142 := 'Ì'     //U+20B1    Peso Sign
    C143 := 'ì'     //U+20B1    Peso Sign
    C144 := 'Ĭ'     //U+20B1    Peso Sign
    C145 := 'ĭ'     //U+20B1    Peso Sign
    C146 := 'Ĩ'     //U+20B1    Peso Sign
    C147 := 'ĩ'     //U+20B1    Peso Sign
    C148 := 'Ī'     //U+20B1    Peso Sign
    C149 := 'ī'     //U+20B1    Peso Sign
    C150 := 'J'     //U+20B1    Peso Sign
    C151 := 'j'     //U+20B1    Peso Sign
    C152 := 'Ĵ'     //U+20B1    Peso Sign
    C153 := 'ĵ'     //U+20B1    Peso Sign
    C154 := 'Λ'     //U+039B    Greek Capital Letter Lamda
    C155 := 'λ'     //U+20B1    Peso Sign
    C156 := 'N'     //U+20B1    Peso Sign
    C157 := 'n'     //U+20B1    Peso Sign
    C158 := 'Ň'     //U+20B1    Peso Sign
    C159 := 'ň'     //U+20B1    Peso Sign
    C160 := 'Ń'     //U+20B1    Peso Sign
    C161 := 'ń'     //U+20B1    Peso Sign
    C162 := 'Ñ'     //U+20B1    Peso Sign
    C163 := 'ñ'     //U+20B1    Peso Sign
    C164 := 'Ņ'     //U+20B1    Peso Sign
    C165 := 'ņ'     //U+20B1    Peso Sign
    C166 := 'O'     //U+20B1    Peso Sign
    C167 := 'o'     //U+20B1    Peso Sign
    C168 := 'Ö'     //U+20B1    Peso Sign
    C169 := 'ö'     //U+20B1    Peso Sign
    C170 := 'Ô'     //U+20B1    Peso Sign
    C171 := 'ô'     //U+20B1    Peso Sign
    C172 := 'Ő'     //U+20B1    Peso Sign
    C173 := 'ő'     //U+20B1    Peso Sign
    C174 := 'Ó'     //U+20B1    Peso Sign
    C175 := 'ó'     //U+20B1    Peso Sign
    C176 := 'Ò'     //U+20B1    Peso Sign
    C177 := 'ò'     //U+20B1    Peso Sign
    C178 := 'Ŏ'     //U+20B1    Peso Sign
    C179 := 'ŏ'     //U+20B1    Peso Sign
    C180 := 'Õ'     //U+20B1    Peso Sign
    C181 := 'õ'     //U+20B1    Peso Sign
    C182 := 'Ō'     //U+20B1    Peso Sign
    C183 := 'ō'     //U+20B1    Peso Sign
    C184 := 'Œ'     //U+20B1    Peso Sign
    C185 := 'œ'     //U+20B1    Peso Sign
    C186 := 'Ω'     //U+03A9    Greek Capital Letter Omega
    C187 := 'ω'     //U+03C9    Greek Small Letter Omega
    C188 := 'P'     //U+20B1    Peso Sign
    C189 := 'p'     //U+20B1    Peso Sign
    C190 := 'Q'     //U+20B1    Peso Sign
    C191 := 'q'     //U+20B1    Peso Sign
    C192 := 'R'     //U+20B1    Peso Sign
    C193 := 'r'     //U+20B1    Peso Sign
    C194 := 'Ř'     //U+20B1    Peso Sign
    C195 := 'ř'     //U+20B1    Peso Sign
    C196 := 'Ŕ'     //U+20B1    Peso Sign
    C197 := 'ŕ'     //U+20B1    Peso Sign
    C198 := 'Ŗ'     //U+20B1    Peso Sign
    C199 := 'ŗ'     //U+20B1    Peso Sign
    C200 := 'S'     //U+20B1    Peso Sign
    C201 := 's'     //U+20B1    Peso Sign
    C202 := 'Š'     //U+20B1    Peso Sign
    C203 := 'š'     //U+20B1    Peso Sign
    C204 := 'Ŝ'     //U+20B1    Peso Sign
    C205 := 'ŝ'     //U+20B1    Peso Sign
    C206 := 'Ś'     //U+20B1    Peso Sign
    C207 := 'ś'     //U+20B1    Peso Sign
    C208 := 'Ş'     //U+20B1    Peso Sign
    C209 := 'ș'     //U+20B1    Peso Sign
    C210 := 'Σ'     //U+03A3    Greek Capital Letter Sigma
    C211 := 'σ'     //U+03C3    Greek Small Letter Sigma
    C212 := 'T'     //U+20B1    Peso Sign
    C213 := 't'     //U+20B1    Peso Sign
    C214 := 'Ț'     //U+20B1    Peso Sign
    C215 := 'ț'     //U+20B1    Peso Sign
    C216 := 'U'     //U+20B1    Peso Sign
    C217 := 'u'     //U+20B1    Peso Sign
    C218 := 'Ů'     //U+20B1    Peso Sign
    C219 := 'ů'     //U+20B1    Peso Sign
    C220 := 'Ü'     //U+20B1    Peso Sign
    C221 := 'ü'     //U+20B1    Peso Sign
    C222 := 'Û'     //U+20B1    Peso Sign
    C223 := 'û'     //U+20B1    Peso Sign
    C224 := 'Ű'     //U+20B1    Peso Sign
    C225 := 'ű'     //U+20B1    Peso Sign
    C226 := 'Ú'     //U+20B1    Peso Sign
    C227 := 'ú'     //U+20B1    Peso Sign
    C228 := 'Ù'     //U+20B1    Peso Sign
    C229 := 'ù'     //U+20B1    Peso Sign
    C230 := 'Ũ'     //U+20B1    Peso Sign
    C231 := 'ũ'     //U+20B1    Peso Sign
    C232 := 'Ū'     //U+20B1    Peso Sign
    C233 := 'ū'     //U+20B1    Peso Sign
    C234 := 'Ʊ'     //U+01B1    Latin Capital Letter Upsilon
    C235 := 'ʊ'     //U+028A    Latin Small Letter Upsilon
    C236 := 'V'     //U+20B1    Peso Sign
    C237 := 'v'     //U+20B1    Peso Sign
    C238 := 'W'     //U+20B1    Peso Sign
    C239 := 'w'     //U+20B1    Peso Sign
    C240 := 'W'     //U+20B1    Peso Sign
    C241 := 'x'     //U+20B1    Peso Sign
    C242 := 'Ξ'     //U+039E    Greek Capital Letter Xi
    C243 := 'ξ'     //U+03BE    Greek Small Letter Xi
    C244 := 'Y'     //U+20B1    Peso Sign
    C245 := 'y'     //U+20B1    Peso Sign
    C246 := 'Ŷ'     //U+20B1    Peso Sign
    C247 := 'ŷ'     //U+20B1    Peso Sign
    C248 := 'Z'     //U+20B1    Peso Sign
    C249 := 'z'     //U+20B1    Peso Sign
    C250 := 'Ż'     //U+20B1    Peso Sign
    C251 := 'ż'     //U+20B1    Peso Sign
    C252 := 'Ž'     //U+20B1    Peso Sign
    C253 := 'ž'     //U+20B1    Peso Sign
    C254 := 'Ź'     //U+20B1    Peso Sign
    C255 := 'ź'     //U+20B1    Peso Sign

    Row00 := [...]rune{C000,C001,C002,C003,C004,C005,C006,C007,C008,C009,C010,C011,C012,C013,C014,C015}
    Row01 := [...]rune{C016,C017,C018,C019,C020,C021,C022,C023,C024,C025,C026,C027,C028,C029,C030,C031}
    Row02 := [...]rune{C032,C033,C034,C035,C036,C037,C038,C039,C040,C041,C042,C043,C044,C045,C046,C047}
    Row03 := [...]rune{C048,C049,C050,C051,C052,C053,C054,C055,C056,C057,C058,C059,C060,C061,C062,C063}
    Row04 := [...]rune{C064,C065,C066,C067,C068,C069,C070,C071,C072,C073,C074,C075,C076,C077,C078,C079}
    Row05 := [...]rune{C080,C081,C082,C083,C084,C085,C086,C087,C088,C089,C090,C091,C092,C093,C094,C095}
    Row06 := [...]rune{C096,C097,C098,C099,C100,C101,C102,C103,C104,C105,C106,C107,C108,C109,C110,C111}
    Row07 := [...]rune{C112,C113,C114,C115,C116,C117,C118,C119,C120,C121,C122,C123,C124,C125,C126,C127}
    Row08 := [...]rune{C128,C129,C130,C131,C132,C133,C134,C135,C136,C137,C138,C139,C140,C141,C142,C143}
    Row09 := [...]rune{C144,C145,C146,C147,C148,C149,C150,C151,C152,C153,C154,C155,C156,C157,C158,C159}
    Row10 := [...]rune{C160,C161,C162,C163,C164,C165,C166,C167,C168,C169,C170,C171,C172,C173,C174,C175}
    Row11 := [...]rune{C176,C177,C178,C179,C180,C181,C182,C183,C184,C185,C186,C187,C188,C189,C190,C191}
    Row12 := [...]rune{C192,C193,C194,C195,C196,C197,C198,C199,C200,C201,C202,C203,C204,C205,C206,C207}
    Row13 := [...]rune{C208,C209,C210,C211,C212,C213,C214,C215,C216,C217,C218,C219,C220,C221,C222,C223}
    Row14 := [...]rune{C224,C225,C226,C227,C228,C229,C230,C231,C232,C233,C234,C235,C236,C237,C238,C239}
    Row15 := [...]rune{C240,C241,C242,C243,C244,C245,C246,C247,C248,C249,C250,C251,C252,C253,C254,C255}

    Matrix := [16][16]rune{Row00,Row01,Row02,Row03,Row04,Row05,Row06,Row07,Row08,Row09,Row10,Row11,Row12,Row13,Row14,Row15}
    return Matrix
}