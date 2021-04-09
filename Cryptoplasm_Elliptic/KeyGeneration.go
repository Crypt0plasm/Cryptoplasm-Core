package Cryptoplasm_Elliptic

import (
    "fmt"
    blake3 "github.com/Crypt0plasm/Cryptographic-Hash-Functions/Blake3"
    "math/big"
    "strings"
)

// PrivateKeyInt and PublicKeyInt are    big.Int(integers)   representing a number in Base10.
// PrivateKey    and PublicKey    are    strings             representing a number in Base49.
// CryptoPlasm Address is computed from the PublicKeyInt.
//
// PublicKey string has a special form.

var CryptoplasmCurve = TEC_S1600_Pr1605p2315_m26()
//var CryptoplasmCurve = LittleOne()

//A CPKeyPair (CryptoPlasm Key Pair) is a pair consisting of two strings both representing a number in base 49.
type CPKeyPair struct {
    PrivateKey string
    PublicKey string
}
//Generic Convert Functions:
func ConvertBase49toBase10 (NumberBase49 string) *big.Int {
    var Result = new(big.Int)
    Result.SetString(NumberBase49,49)
    return Result
}
func ConvertBase10toBase49 (NumberBase10 *big.Int) string {
    var Result string
    Result = NumberBase10.Text(49)
    return Result
}
//======================================================================================================================
//
//Main Generation Function
//CryptoPlasmKeys are generated using CryptoplasmCurve defined above.
//
//Generates a CryptoPlasm PublicKey and Address from a compatible BitString
func BitString2CRYPTOPLASM (BitString string) (Keys CPKeyPair, Address string) {

    //First Main Function, Getting Keys
    Keys = CryptoplasmCurve.GetKeys(BitString)

    //Second Main Function, Getting the Address
    Address = PublicKey2CRYPTOPLASMAddress(Keys.PublicKey)
    return Keys, Address
}
//======================================================================================================================
//
// VI - Key Generation Methods - Part II
// Functions required for the First Main Function:
//
// VI.5
//
// Generated a CryptoPlasm PublicKey and Address from a Random BitString
func (k *FiniteFieldEllipticCurve) RBS2CRYPTOPLASM () (Keys CPKeyPair, Address string) {

    //First Main Function, Getting Keys
    BitString := k.GetRandomBitsOnCurve()
    Keys,Address = BitString2CRYPTOPLASM(BitString)
    return Keys, Address
}
// VI.6
func (k *FiniteFieldEllipticCurve) GetKeys (BitString string) (Keys CPKeyPair) {
    PrivateKeyInt := k.ClampBitString(BitString)
    Keys.PrivateKey = ConvertBase10toBase49(PrivateKeyInt)
    Keys.PublicKey = k.PrivKey2PubKey(Keys.PrivateKey)
    return Keys
}
// VI.7
func (k *FiniteFieldEllipticCurve) PrivKey2PubKey (PrivateKey string) (PublicKey string) {
    PrivateKeyInt := ConvertBase49toBase10(PrivateKey)
    PublicKey = k.PrivKeyInt2PubKey(PrivateKeyInt)
    return PublicKey
}
// VI.8
func (k *FiniteFieldEllipticCurve) PrivKeyInt2PubKey (PrivateKeyInt *big.Int) (PublicKey string) {
    var (
    	PublicKeyInt = new(big.Int)
        XStringLengthBig = new(big.Int)
    )

    //Scalar Multiplication returns Ext Coordinates, and we convert them back to Affine
    PublicKeyPoints := k.ScalarMultiplierG(PrivateKeyInt)
    PublicKeyPointsAff := k.Extended2Affine(PublicKeyPoints)
    //To verify that PublicKey2Affine returns the same Points as those computed from Private Key
    //Remove the following 1 Comments:
    //fmt.Println("Points:",PublicKeyPoints)

    XString := PublicKeyPointsAff.AX.String()
    XStringLength := int64(len(XString))
    XStringLengthBig.SetInt64(XStringLength)
    PublicKeyPrefix := ConvertBase10toBase49(XStringLengthBig)
    //The PublicKeyPrefix is the Base49 value of the Length of X.
    //This information is needed in order to remake the Affine Coordinates
    //representing the PublicKey from the PublicKey string

    YString := PublicKeyPointsAff.AY.String()
    XYString := XString+YString
    PublicKeyInt.SetString(XYString,10)
    PublicKey = ConvertBase10toBase49(PublicKeyInt)
    PublicKey = PublicKeyPrefix + "." + PublicKey
    return PublicKey
}
//======================================================================================================================
//
// Functions required for the Second Main Function
//
func PublicKey2CRYPTOPLASMAddress (PublicKey string) string {
    //From the PublicKey string, the Prefix is removed in order to obtain the PublicKeyInt
    SplitString := strings.Split(PublicKey,".")
    PublicKeyIntStr := SplitString[1]

    //To verify that PublicKey2Affine returns the same Points as those computed from Private Key
    //Remove the following 2 Comments:
    //RemadePoints := PublicKey2Affine(PublicKey)
    //fmt.Println("Points:",RemadePoints)
    PublicKeyInt := ConvertBase49toBase10(PublicKeyIntStr)
    CryptoPlasmAddress := PublicKeyInt2CRYPTOPLASMAddress(PublicKeyInt)
    return CryptoPlasmAddress
}

func PublicKeyInt2CRYPTOPLASMAddress (PublicKeyInt *big.Int) string {
    PublicKeyIntHashed := SevenFoldHash(PublicKeyInt)
    CryptoPlasmAddress := ConvToLetters(PublicKeyIntHashed)
    return CryptoPlasmAddress
}

func SevenFoldHash (PublicKeyInt *big.Int) []byte {
    PublicKeyIntAsString := PublicKeyInt.String()
    LongStringToByteSlice := []byte(PublicKeyIntAsString)

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

    //Currencies Block 10 runes
    C010 := 'Ѻ'     //U+047A    Cyrillic Capital Letter Round Omega aka Cryptoplasm
    C011 := '₿'     //U+20BF    Bitcoin Sign
    C012 := '$'     //U+0024    Dollar Sign
    C013 := '¢'     //U+00A2    Cent Sign
    C014 := '€'     //U+20AC    Euro Sign
    C015 := '£'     //U+00A3    Pound Sign
    C016 := '¥'     //U+00A5    Yen Sign
    C017 := '₽'     //U+20BD    Ruble Sign
    C018 := '₱'     //U+20B1    Peso Sign
    C019 := '₣'     //U+20A3    French Franc Sign

    //Lonely Small Letters 36 Runes
    C020 := 'α'     //U+03B1    Greek Small Letter Alpha
    C021 := 'b'     //U+0062    Latin Small Letter B
    C022 := 'β'     //U+03B2    Greek Small Letter Beta
    C023 := 'χ'     //U+03C7    Greek Small Letter Chi
    C024 := 'д'     //U+0434    Cyrillic Small Letter De
    C025 := 'ε'     //U+03B5    Greek Small Letter Epsilon
    C026 := 'ɣ'     //U+0263    Latin Small Letter Gamma
    C027 := 'γ'     //U+03B3    Greek Small Letter Gamma
    C028 := 'г'     //U+0433    Cyrillic Small Letter Ghe
    C029 := 'h'     //U+0068    Latin Small Letter H
    C030 := 'η'     //U+03B7    Greek Small Letter Eta
    C031 := 'и'     //U+0438    Cyrillic Small Letter I
    C032 := 'й'     //U+0439    Cyrillic Small Letter Short I
    C033 := 'ж'     //U+0436    Cyrillic Small Letter Zhe
    C034 := 'k'     //U+006B    Latin Small Letter K
    C035 := 'л'     //U+043B    Cyrillic Small Letter El
    C036 := 'm'     //U+20B1    Peso Sign
    C037 := 'μ'     //U+03BC    Greek Small Letter Mu
    C038 := 'м'     //U+043C    Cyrillic Small Letter Em
    C039 := 'ν'     //U+03BD    Greek Small Letter Nu
    C040 := 'н'     //U+043D    Cyrillic Small Letter En
    C041 := 'π'     //U+03C0    Greek Small Letter Pi
    C042 := 'ψ'     //U+03C8    Greek Small Letter Psi
    C043 := 'п'     //U+043F    Cyrillic Small Letter Pe
    C044 := 'ρ'     //U+03C1    Greek Small Letter Rho
    C045 := 'ς'     //U+03C2    Greek Small Letter Final Sigma
    C046 := 'ш'     //U+0448    Cyrillic Small Letter Sha
    C047 := 'щ'     //U+0449    Cyrillic Small Letter Shcha
    C048 := 'τ'     //U+03C4    Greek Small Letter Tau
    C049 := 'ц'     //U+0446    Cyrillic Small Letter Tse
    C050 := 'в'     //U+0432    Cyrillic Small Letter Ve
    C051 := 'ю'     //U+044E    Cyrillic Small Letter Yu
    C052 := 'я'     //U+044F    Cyrillic Small Letter Ya
    C053 := 'ß'     //U+00DF    Latin Small Letter Sharp S
    C054 := 'ъ'     //U+044A    Cyrillic Small Letter Hard Sign
    C055 := 'ь'     //U+044C    Cyrillic Small Letter Soft Sign

    //Capital-Small Letters pairs 2*100=200 Runes
    C056 := 'A'     //U+0041    Latin Capital Letter A
    C057 := 'a'     //U+0061    Latin Small Letter A
    C058 := 'Å'     //U+00C5    Latin Capital Letter A with Ring Above
    C059 := 'å'     //U+00E5    Latin Small Letter A with Ring Above
    C060 := 'Ä'     //U+00C4    Latin Capital Letter A with Diaeresis
    C061 := 'ä'     //U+00E4    Latin Small Letter A with Diaeresis
    C062 := 'Â'     //U+00C2    Latin Capital Letter A with Circumflex
    C063 := 'â'     //U+00E2    Latin Small Letter A with Circumflex
    C064 := 'Á'     //U+00C1    Latin Capital Letter A with Acute
    C065 := 'á'     //U+00E1    Latin Small Letter A with Acute
    C066 := 'À'     //U+00C0    Latin Capital Letter A with Grave
    C067 := 'à'     //U+00E0    Latin Small Letter A with Grave
    C068 := 'Ă'     //U+0102    Latin Capital Letter A with Breve
    C069 := 'ă'     //U+0103    Latin Small Letter A with Breve
    C070 := 'Ã'     //U+00C3    Latin Capital Letter A with Tilde
    C071 := 'ã'     //U+00E3    Latin Small Letter A with Tilde
    C072 := 'Ā'     //U+0100    Latin Capital Letter A with Macron
    C073 := 'ā'     //U+0101    Latin Small Letter A with Macron
    C074 := 'Æ'     //U+00C6    Latin Capital Letter Ae
    C075 := 'æ'     //U+00E6    Latin Small Letter Ae
    C076 := 'Б'     //U+0411    Cyrillic Capital Letter Be
    C077 := 'б'     //U+0431    Cyrillic Small Letter Be
    C078 := 'C'     //U+0043    Latin Capital Letter C
    C079 := 'c'     //U+0063    Latin Small Letter C
    C080 := 'Ċ'     //U+010A    Latin Capital Letter C with Dot Above
    C081 := 'ċ'     //U+010B    Latin Small Letter C with Dot Above
    C082 := 'Č'     //U+010C    Latin Capital Letter C with Caron
    C083 := 'č'     //U+010D    Latin Small Letter C with Caron
    C084 := 'Ĉ'     //U+0108    Latin Capital Letter C with Circumflex
    C085 := 'ĉ'     //U+0109    Latin Small Letter C with Circumflex
    C086 := 'Ć'     //U+0106    Latin Capital Letter C with Acute
    C087 := 'ć'     //U+0107    Latin Small Letter C with Acute
    C088 := 'Ç'     //U+00C7    Latin Capital Letter C with Cedilla
    C089 := 'ç'     //U+00E7    Latin Small Letter C with Cedilla
    C090 := 'Ч'     //U+0427    Cyrillic Capital Letter Che
    C091 := 'ч'     //U+0447    Cyrillic Small Letter Che
    C092 := 'D'     //U+0044    Latin Capital Letter D
    C093 := 'd'     //U+0064    Latin Small Letter D
    C094 := 'Δ'     //U+0394    Greek Capital Letter Delta
    C095 := 'δ'     //U+03B4    Greek Small Letter Delta
    C096 := 'E'     //U+0045    Latin Capital Letter E
    C097 := 'e'     //U+0065    Latin Small Letter E
    C098 := 'Ė'     //U+0116    Latin Capital Letter E with Dot Above
    C099 := 'ė'     //U+0117    Latin Small Letter E with Dot Above
    C100 := 'Ë'     //U+00CB    Latin Capital Letter E with Diaeresis
    C101 := 'ë'     //U+00EB    Latin Small Letter E with Diaeresis
    C102 := 'Ě'     //U+011A    Latin Capital Letter E with Caron
    C103 := 'ě'     //U+011B    Latin Small Letter E with Caron
    C104 := 'Ê'     //U+00CA    Latin Capital Letter E with Circumflex
    C105 := 'ê'     //U+20B1    Latin Small Letter E with Circumflex
    C106 := 'É'     //U+00C9    Latin Capital Letter E with Acute
    C107 := 'é'     //U+00E9    Latin Small Letter E with Acute
    C108 := 'È'     //U+00C8    Latin Capital Letter E with Grave
    C109 := 'è'     //U+00E8    Latin Small Letter E with Grave
    C110 := 'Ĕ'     //U+0114    Latin Capital Letter E with Breve
    C111 := 'ĕ'     //U+0115    Latin Small Letter E with Breve
    C112 := 'Ē'     //U+0112    Latin Capital Letter E with Macron
    C113 := 'ē'     //U+0113    Latin Small Letter E with Macron
    C114 := 'Э'     //U+042D    Cyrillic Capital Letter E
    C115 := 'э'     //U+044D    Cyrillic Small Letter E
    C116 := 'F'     //U+0046    Latin Capital Letter F
    C117 := 'f'     //U+0066    Latin Small Letter F
    C118 := 'Φ'     //U+03A6    Greek Capital Letter Phi
    C119 := 'φ'     //U+03C6    Greek Small Letter Phi
    C120 := 'G'     //U+0047    Latin Capital Letter G
    C121 := 'g'     //U+0067    Latin Small Letter G
    C122 := 'Ġ'     //U+0120    Latin Capital Letter G with Dot Above
    C123 := 'ġ'     //U+0121    Latin Small Letter G with Dot Above
    C124 := 'Ĝ'     //U+011C    Latin Capital Letter G with Circumflex
    C125 := 'ĝ'     //U+011D    Latin Small Letter G with Circumflex
    C126 := 'Ğ'     //U+011E    Latin Capital Letter G with Breve
    C127 := 'ğ'     //U+011F    Latin Small Letter G with BreveF
    C128 := 'Ĥ'     //U+0124    Latin Capital Letter H with Circumflex
    C129 := 'ĥ'     //U+0125    Latin Small Letter H with Circumflex
    C130 := 'I'     //U+0049    Latin Capital Letter I
    C131 := 'i'     //U+0069    Latin Small Letter I
    C132 := 'Ï'     //U+00CF    Latin Capital Letter I with Diaeresis
    C133 := 'ï'     //U+00EF    Latin Small Letter I with Diaeresis
    C134 := 'Î'     //U+00CE    Latin Capital Letter I with Circumflex
    C135 := 'î'     //U+00EE    Latin Small Letter I with Circumflex
    C136 := 'Í'     //U+00CD    Latin Capital Letter I with Acute
    C137 := 'í'     //U+00ED    Latin Small Letter I with Acute
    C138 := 'Ì'     //U+00CC    Latin Capital Letter I with Grave
    C139 := 'ì'     //U+00EC    Latin Small Letter I with Grave
    C140 := 'Ĭ'     //U+012C    Latin Capital Letter I with Breve
    C141 := 'ĭ'     //U+012D    Latin Small Letter I with Breve
    C142 := 'Ĩ'     //U+0128    Latin Capital Letter I with Tilde
    C143 := 'ĩ'     //U+0129    Latin Small Letter I with Tilde
    C144 := 'Ī'     //U+012A    Latin Capital Letter I with Macron
    C145 := 'ī'     //U+012B    Latin Small Letter I with Macron
    C146 := 'J'     //U+004A    Latin Capital Letter J
    C147 := 'j'     //U+006A    Latin Small Letter J
    C148 := 'Ĵ'     //U+0134    Latin Capital Letter J with Circumflex
    C149 := 'ĵ'     //U+0135    Latin Small Letter J with Circumflex
    C150 := 'Λ'     //U+039B    Greek Capital Letter Lamda
    C151 := 'λ'     //U+03BB    Greek Small Letter Lamda
    C152 := 'N'     //U+004E    Latin Capital Letter N
    C153 := 'n'     //U+006E    Latin Small Letter N
    C154 := 'Ň'     //U+0147    Latin Capital Letter N with Caron
    C155 := 'ň'     //U+0148    Latin Small Letter N with Caron
    C156 := 'Ń'     //U+0143    Latin Capital Letter N with Acute
    C157 := 'ń'     //U+0144    Latin Small Letter N with Acute
    C158 := 'Ñ'     //U+00D1    Latin Capital Letter N with Tilde
    C159 := 'ñ'     //U+00F1    Latin Small Letter N with Tilde
    C160 := 'Ņ'     //U+0145    Latin Capital Letter N with Cedilla
    C161 := 'ņ'     //U+0146    Latin Small Letter N with Cedilla
    C162 := 'O'     //U+004F    Latin Capital Letter O
    C163 := 'o'     //U+006F    Latin Small Letter O
    C164 := 'Ö'     //U+00D6    Latin Capital Letter O with Diaeresis
    C165 := 'ö'     //U+00F6    Latin Small Letter O with Diaeresis
    C166 := 'Ô'     //U+00D4    Latin Capital Letter O with Circumflex
    C167 := 'ô'     //U+00F4    Latin Small Letter O with Circumflex
    C168 := 'Ő'     //U+0150    Latin Capital Letter O with Double Acute
    C169 := 'ő'     //U+0151    Latin Small Letter O with Double Acute
    C170 := 'Ó'     //U+00D3    Latin Capital Letter O with Acute
    C171 := 'ó'     //U+00F3    Latin Small Letter O with Acute
    C172 := 'Ò'     //U+00D2    Latin Capital Letter O with Grave
    C173 := 'ò'     //U+00F2    Latin Small Letter O with Grave
    C174 := 'Ŏ'     //U+014E    Latin Capital Letter O with Breve
    C175 := 'ŏ'     //U+014F    Latin Small Letter O with Breve
    C176 := 'Õ'     //U+00D5    Latin Capital Letter O with Tilde
    C177 := 'õ'     //U+00F5    Latin Small Letter O with Tilde
    C178 := 'Ō'     //U+014C    Latin Capital Letter O with Macron
    C179 := 'ō'     //U+014D    Latin Small Letter O with Macron
    C180 := 'Œ'     //U+0152    Latin Capital Ligature Oe
    C181 := 'œ'     //U+0153    Latin Small Ligature Oe
    C182 := 'Ω'     //U+03A9    Greek Capital Letter Omega
    C183 := 'ω'     //U+03C9    Greek Small Letter Omega
    C184 := 'P'     //U+0050    Latin Capital Letter P
    C185 := 'p'     //U+0070    Latin Small Letter P
    C186 := 'Q'     //U+0051    Latin Capital Letter Q
    C187 := 'q'     //U+0071    Latin Small Letter Q
    C188 := 'R'     //U+0052    Latin Capital Letter R
    C189 := 'r'     //U+0072    Latin Small Letter R
    C190 := 'Ř'     //U+0158    Latin Capital Letter R with Caron
    C191 := 'ř'     //U+0159    Latin Small Letter R with Caron
    C192 := 'Ŕ'     //U+0154    Latin Capital Letter R with Acute
    C193 := 'ŕ'     //U+0155    Latin Small Letter R with Acute
    C194 := 'Ŗ'     //U+0156    Latin Capital Letter R with Cedilla
    C195 := 'ŗ'     //U+0157    Latin Small Letter R with Cedilla
    C196 := 'S'     //U+0053    Latin Capital Letter S
    C197 := 's'     //U+0073    Latin Small Letter S
    C198 := 'Š'     //U+0160    Latin Capital Letter S with Caron
    C199 := 'š'     //U+0161    Latin Small Letter S with Caron
    C200 := 'Ŝ'     //U+015C    Latin Capital Letter S with Circumflex
    C201 := 'ŝ'     //U+015D    Latin Small Letter S with Circumflex
    C202 := 'Ś'     //U+015A    Latin Capital Letter S with Acute
    C203 := 'ś'     //U+015B    Latin Small Letter S with Acute
    C204 := 'Ş'     //U+015E    Latin Capital Letter S with Cedilla
    C205 := 'ş'     //U+015F    Latin Small Letter S with Cedilla
    C206 := 'Ș'     //U+0218    Latin Capital Letter S with Comma Below
    C207 := 'ș'     //U+0219    Latin Small Letter S with Comma Below
    C208 := 'Σ'     //U+03A3    Greek Capital Letter Sigma
    C209 := 'σ'     //U+03C3    Greek Small Letter Sigma
    C210 := 'T'     //U+0054    Latin Capital Letter T
    C211 := 't'     //U+0074    Latin Small Letter T
    C212 := 'Ţ'     //U+0162    Latin Capital Letter T with Cedilla
    C213 := 'ţ'     //U+0163    Latin Small Letter T with Cedilla
    C214 := 'Ț'     //U+021A    Latin Capital Letter T with Comma Below
    C215 := 'ț'     //U+021B    Latin Small Letter T with Comma Below
    C216 := 'U'     //U+0055    Latin Capital Letter U
    C217 := 'u'     //U+0075    Latin Small Letter U
    C218 := 'Ů'     //U+016E    Latin Capital Letter U with Ring Above
    C219 := 'ů'     //U+016F    Latin Small Letter U with Ring Above
    C220 := 'Ü'     //U+00DC    Latin Capital Letter U with Diaeresis
    C221 := 'ü'     //U+00FC    Latin Small Letter U with Diaeresis
    C222 := 'Û'     //U+00DB    Latin Capital Letter U with Circumflex
    C223 := 'û'     //U+00FB    Latin Small Letter U with Circumflex
    C224 := 'Ű'     //U+0170    Latin Capital Letter U with Double Acute
    C225 := 'ű'     //U+0171    Latin Small Letter U with Double Acute
    C226 := 'Ú'     //U+00DA    Latin Capital Letter U with Acute
    C227 := 'ú'     //U+00FA    Latin Small Letter U with Acute
    C228 := 'Ù'     //U+00D9    Latin Capital Letter U with Grave
    C229 := 'ù'     //U+00F9    Latin Small Letter U with Grave
    C230 := 'Ũ'     //U+0168    Latin Capital Letter U with Tilde
    C231 := 'ũ'     //U+0169    Latin Small Letter U with Tilde
    C232 := 'Ū'     //U+016A    Latin Capital Letter U with Macron
    C233 := 'ū'     //U+016B    Latin Small Letter U with Macron
    C234 := 'Ʊ'     //U+01B1    Latin Capital Letter Upsilon
    C235 := 'ʊ'     //U+028A    Latin Small Letter Upsilon
    C236 := 'V'     //U+0056    Latin Capital Letter V
    C237 := 'v'     //U+0076    Latin Small Letter V
    C238 := 'W'     //U+0057    Latin Capital Letter W
    C239 := 'w'     //U+0077    Latin Small Letter W
    C240 := 'X'     //U+0058    Latin Capital Letter X
    C241 := 'x'     //U+0078    Latin Small Letter X
    C242 := 'Ξ'     //U+039E    Greek Capital Letter Xi
    C243 := 'ξ'     //U+03BE    Greek Small Letter Xi
    C244 := 'Y'     //U+0059    Latin Capital Letter Y
    C245 := 'y'     //U+0079    Latin Small Letter Y
    C246 := 'Ŷ'     //U+0176    Latin Capital Letter Y with Circumflex
    C247 := 'ŷ'     //U+0177    Latin Small Letter Y with Circumflex
    C248 := 'Z'     //U+005A    Latin Capital Letter Z
    C249 := 'z'     //U+007A    Latin Small Letter Z
    C250 := 'Ż'     //U+017B    Latin Capital Letter Z with Dot Above
    C251 := 'ż'     //U+017C    Latin Small Letter Z with Dot Above
    C252 := 'Ž'     //U+017D    Latin Capital Letter Z with Caron
    C253 := 'ž'     //U+017E    Latin Small Letter Z with Caron
    C254 := 'Ź'     //U+0179    Latin Capital Letter Z with Acute
    C255 := 'ź'     //U+017A    Latin Small Letter Z with Acute

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
//======================================================================================================================
//
func (k *FiniteFieldEllipticCurve) MakeCryptoplasmKeys () {
    var(
        FirstAnswer uint
        Answer11, Answer21, Answer31 uint
        Answer111 uint
        Answer1112 string
        Condition1 bool

        Print1 = `1.Create New Cryptoplasm Keys
2.Restore Cryptoplasm Keys
3.Open Existing Cryptoplasm Keys`
        Print10 = `How do you want to create the Private Key ?`
        Print11 = `1.1.BitString         - create Cryptoplasm Private Key from a BitString
1.2.Number in Base 10 - create Cryptoplasm Private Key from a Base 10 Number
1.3.Number in Base 49 - create Cryptoplasm Private Key from a Base 49 Number
1.4.Monochrome Bitmap - create Cryptoplasm Private Key from a Monochrome Bitmap. Whites are 0, Blacks are 1
1.5.Custom Seed Words - create Cryptoplasm Private Key from a custom List of Words`
        Print111 = `1.1.1 Random BitString - create a Cryptoplasm Private Key from a random BitString
1.1.2 Manual BitString - create a Cryptoplasm Private Key from a manually inputted BitString`
        Print20 = `Restoring assumes the following are password protected.
What do you want to restore the Private Key from?`
        Print21 = `2.1.BitString         - create Cryptoplasm Private Key from a BitString
2.2.Number in Base 10 - create Cryptoplasm Private Key from a Base 10 Number
2.3.Number in Base 49 - create Cryptoplasm Private Key from a Base 49 Number
2.4.Monochrome Bitmap - create Cryptoplasm Private Key from a Monochrome Bitmap. Whites are 0, Blacks are 1
2.5.Custom Seed Words - create Cryptoplasm Private Key from a custom List of Words`
        Print30 = `How do you have the Private Key saved ?`
        Print31 = `1.Key File          - open Cryptoplasm Private Key from a Key-File
2.Monochrome Bitmap - the Cryptoplasm Private Key from a Monochrome-Bitmap`

    )
    fmt.Println("What do you want to do ?")
    fmt.Println(Print1)
    _, _ = fmt.Scanln(&FirstAnswer)
    if FirstAnswer == 1 {
        fmt.Println(Print10)
        fmt.Println(Print11)
        _, _ = fmt.Scanln(&Answer11)
        switch Answer11 {
        case 1 :
            fmt.Println("Case 1")
            fmt.Println(Print111)
            _, _ = fmt.Scanln(&Answer111)
            switch Answer111 {
            case 1 :
                fmt.Println("Case 1")
                Keys, Address := k.RBS2CRYPTOPLASM()
                fmt.Println("Private Key is,",Keys.PrivateKey)
                fmt.Println("Public  Key is,",Keys.PublicKey)
                fmt.Println("Address Str is,",Address)
            case 2 :
                for {
                    fmt.Println("Please enter a string of bits(0s and 1s)",k.S,"digits long:")
                    _, _ = fmt.Scanln(&Answer1112)
                    TT,T1,T2 := k.ValidateBitString(Answer1112)
                    if TT == true {
                        Condition1 = true
                    }
                    if Condition1 == true {
                        break
                    } else {
                        if T1 == false && T2 == true {
                            fmt.Println("Entered BitString didn't have the required length, try again!")
                        } else if T1 == true && T2 == false {
                            fmt.Println("Entered BitString wasn't composed entirely of 0s and 1s, try again!")
                        } else if T1 == false && T2 == false {
                            fmt.Println("Entered BitString neither did have the required length, nor was it composed entirely of 0s and 1s, try again!")
                        }
                    }
                }
                Keys, Address := BitString2CRYPTOPLASM(Answer1112)
                fmt.Println("Private Key is,",Keys.PrivateKey)
                fmt.Println("Public  Key is,",Keys.PublicKey)
                fmt.Println("Address Str is,",Address)
            }
        case 2 :
            fmt.Println("Case 2")
        case 3 :
            fmt.Println("Case 3")
        case 4 :
            fmt.Println("Case 4")
        case 5 :
            fmt.Println("Case 5")
        }
    } else if FirstAnswer == 2 {
        fmt.Println(Print20)
        fmt.Println(Print21)
        _, _ = fmt.Scanln(&Answer21)
        switch Answer21 {
        case 1 :
            fmt.Println("Case 1")
        case 2 :
            fmt.Println("Case 2")
        case 3 :
            fmt.Println("Case 3")
        case 4 :
            fmt.Println("Case 4")
        case 5 :
            fmt.Println("Case 5")
        }
    } else if FirstAnswer == 3 {
        fmt.Println(Print30)
        fmt.Println(Print31)
        _, _ = fmt.Scanln(&Answer31)
        switch Answer31 {
        case 1 :
            fmt.Println("Case 1")
        case 2 :
            fmt.Println("Case 2")
        }
    }
}