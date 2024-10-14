package Cryptoplasm_Elliptic

import (
    aux "Cryptoplasm-Core/Auxiliary"
    "encoding/hex"
    "fmt"
    "github.com/Crypt0plasm/Cryptographic-Hash-Functions/AES"
    "github.com/Crypt0plasm/Cryptographic-Hash-Functions/Blake3"
    "log"
    "math/big"
    "os"
    "strings"
)

// PrivateKeyInt and PublicKeyInt are    big.Int(integers)   representing a number in Base10.
// PrivateKey    and PublicKey    are    strings             representing a number in Base49.
// CryptoPlasm Address is computed from the PublicKeyInt.
//
// PublicKey string has a special form.

var CryptoplasmCurve = TEC_S1600_Pr1605p2315_m26()

//var CryptoplasmCurve = E521()
//var CryptoplasmCurve = LittleOne()

//A CPKeyPair (CryptoPlasm Key Pair) is a pair consisting of two strings both representing a number in base 49.
type CPKeyPair struct {
    PrivateKey string
    PublicKey  string
}

//Generic Convert Functions:
func ConvertBase2toBase10(NumberBase2 string) *big.Int {
    var Result = new(big.Int)
    Result.SetString(NumberBase2, 2)
    return Result
}
func ConvertBase16toBase10(NumberBase16 string) *big.Int {
    var Result = new(big.Int)
    Result.SetString(NumberBase16, 16)
    return Result
}
func ConvertBase49toBase10(NumberBase49 string) *big.Int {
    var Result = new(big.Int)
    Result.SetString(NumberBase49, 49)
    return Result
}
func ConvertBase49toBase2(NumberBase49 string) string {
    Base10 := ConvertBase49toBase10(NumberBase49)
    Result := ConvertBase10toBase2(Base10)
    return Result
}
func ConvertBase2toBase49(NumberBase2 string) string {
    Base10 := ConvertBase2toBase10(NumberBase2)
    Result := ConvertBase10toBase49(Base10)
    return Result
}
func ConvertBase2toBase16(NumberBase2 string) string {
    Base10 := ConvertBase2toBase10(NumberBase2)
    Result := ConvertBase10toBase16(Base10)
    return Result
}
func ConvertBase10toBase49(NumberBase10 *big.Int) string {
    var Result string
    Result = NumberBase10.Text(49)
    return Result
}
func ConvertBase10toBase16(NumberBase10 *big.Int) string {
    var Result string
    Result = NumberBase10.Text(16)
    return Result
}
func ConvertBase10toBase2(NumberBase10 *big.Int) string {
    var Result string
    Result = NumberBase10.Text(2)
    return Result
}

//Key Print Function
func PrintKeysWithAddress(Keys CPKeyPair, Address string) () {
    fmt.Println("Private Key is,", Keys.PrivateKey)
    fmt.Println("Public  Key is,", Keys.PublicKey)
    fmt.Println("Address Str is,", Address)
}

//======================================================================================================================
//
//Main Generation Function
//CryptoPlasmKeys are generated using CryptoplasmCurve defined above.
//
//Generates a CryptoPlasm PublicKey and Address from a compatible BitString
func (k *FiniteFieldEllipticCurve) BS2CRYPTOPLASM(BitString string) (Keys CPKeyPair, Address string) {
    
    //First Main Function, Getting Keys
    Keys = k.GetKeys(BitString)
    
    //Second Main Function, Getting the Address
    Address = PublicKey2CRYPTOPLASMAddress(Keys.PublicKey)
    return Keys, Address
}

//Generates a CryptoPlasm PublicKey and Address from a compatible Number in Base10.
//It is assumed the number is compatible with the Curve in use, that is, is of a clamped form according to
//used Curve Cofactor.Therefore attention must be used not to pass an incompatible number.
//This function is used only after the number was validated previously.
func Number2CRYPTOPLASM(Number *big.Int) (Keys CPKeyPair, Address string) {
    
    //First Main Function, Getting Keys
    Keys = CryptoplasmCurve.GetKeysInt(Number)
    
    //Second Main Function, Getting the Address
    Address = PublicKey2CRYPTOPLASMAddress(Keys.PublicKey)
    return Keys, Address
}

//======================================================================================================================
//
// VI - Key Generation Methods - Part II
//
// VIb.1
// Generates a CryptoPlasm Key-Pair and Address from a Random BitString (which is used as base for the Private-Key)
func (k *FiniteFieldEllipticCurve) RBS2CRYPTOPLASM() (Keys CPKeyPair, Address string) {
    
    //First Main Function, Getting Keys
    BitString := k.GetRandomBitsOnCurve()
    Keys, Address = k.BS2CRYPTOPLASM(BitString)
    return Keys, Address
}

// VIb.2
// Generates a Cryptoplasm Key-Pair from a given BitString (which is used as base for the Private-Key)
func (k *FiniteFieldEllipticCurve) GetKeys(BitString string) (Keys CPKeyPair) {
    PrivateKeyInt := k.ClampBitString(BitString)
    Keys.PrivateKey = ConvertBase10toBase49(PrivateKeyInt)
    Keys.PublicKey = k.PrivKey2PubKey(Keys.PrivateKey)
    return Keys
}

// VIb.3
// Generates a Cryptoplasm Key-Pair from a given Number (which is used as base for the Private-Key)
func (k *FiniteFieldEllipticCurve) GetKeysInt(Number *big.Int) (Keys CPKeyPair) {
    PrivateKeyInt := Number
    Keys.PrivateKey = ConvertBase10toBase49(PrivateKeyInt)
    Keys.PublicKey = k.PrivKey2PubKey(Keys.PrivateKey)
    return Keys
}

// VIb.4
// Generates the PublicKey from a String representing the PrivateKey
func (k *FiniteFieldEllipticCurve) PrivKey2PubKey(PrivateKey string) (PublicKey string) {
    PrivateKeyInt := ConvertBase49toBase10(PrivateKey)
    PublicKey = k.PrivKeyInt2PubKey(PrivateKeyInt)
    return PublicKey
}

// VIb.5
// Extracts the BitString from a Private-Key
func (k *FiniteFieldEllipticCurve) PrivKey2BitString(PrivateKey string) (BitString string) {
    
    PrivateKeyBitString := ConvertBase49toBase2(PrivateKey)
    R := []rune(PrivateKeyBitString)
    R2 := R[:k.S+1]
    for i := 1; i < len(R2); i++ {
        BitString = BitString + string(R2[i])
    }
    return BitString
}

// VIb.6
// Generates the PublicKey in its custom format from a Number representing the PrivateKey.
func (k *FiniteFieldEllipticCurve) PrivKeyInt2PubKey(PrivateKeyInt *big.Int) (PublicKey string) {
    var (
        PublicKeyInt     = new(big.Int)
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
    XYString := XString + YString
    PublicKeyInt.SetString(XYString, 10)
    PublicKey = ConvertBase10toBase49(PublicKeyInt)
    PublicKey = PublicKeyPrefix + "." + PublicKey
    return PublicKey
}

//======================================================================================================================
//
// Functions required for the Second Main Function
//
func PublicKey2CRYPTOPLASMAddress(PublicKey string) string {
    //From the PublicKey string, the Prefix is removed in order to obtain the PublicKeyInt
    SplitString := strings.Split(PublicKey, ".")
    PublicKeyIntStr := SplitString[1]
    
    //To verify that PublicKey2Affine returns the same Points as those computed from Private Key
    //Remove the following 2 Comments:
    //RemadePoints := PublicKey2Affine(PublicKey)
    //fmt.Println("Points:",RemadePoints)
    PublicKeyInt := ConvertBase49toBase10(PublicKeyIntStr)
    CryptoPlasmAddress := PublicKeyInt2CRYPTOPLASMAddress(PublicKeyInt)
    return CryptoPlasmAddress
}

func PublicKeyInt2CRYPTOPLASMAddress(PublicKeyInt *big.Int) string {
    PublicKeyIntHashed := SevenFoldHash(PublicKeyInt)
    CryptoPlasmAddress := ConvToLetters(PublicKeyIntHashed)
    return CryptoPlasmAddress
}

func SevenFoldHash(PublicKeyInt *big.Int) []byte {
    PublicKeyIntAsString := PublicKeyInt.String()
    LongStringToByteSlice := []byte(PublicKeyIntAsString)
    
    S1 := Blake3.SumCustom(LongStringToByteSlice, 160)
    S2 := Blake3.SumCustom(S1, 160)
    S3 := Blake3.SumCustom(S2, 160)
    S4 := Blake3.SumCustom(S3, 160)
    S5 := Blake3.SumCustom(S4, 160)
    S6 := Blake3.SumCustom(S5, 160)
    S7 := Blake3.SumCustom(S6, 160)
    return S7
}

func ConvToLetters(hash []byte) string {
    var (
        result   string
        SliceStr []string
    )
    
    Matrix := CharacterMatrix()
    for i := 0; i < len(hash); i++ {
        Value := hash[i]
        switch Value {
        case 0:
            SliceStr = append(SliceStr, string(Matrix[0][0]))
        case 1:
            SliceStr = append(SliceStr, string(Matrix[0][1]))
        case 2:
            SliceStr = append(SliceStr, string(Matrix[0][2]))
        case 3:
            SliceStr = append(SliceStr, string(Matrix[0][3]))
        case 4:
            SliceStr = append(SliceStr, string(Matrix[0][4]))
        case 5:
            SliceStr = append(SliceStr, string(Matrix[0][5]))
        case 6:
            SliceStr = append(SliceStr, string(Matrix[0][6]))
        case 7:
            SliceStr = append(SliceStr, string(Matrix[0][7]))
        case 8:
            SliceStr = append(SliceStr, string(Matrix[0][8]))
        case 9:
            SliceStr = append(SliceStr, string(Matrix[0][9]))
        case 10:
            SliceStr = append(SliceStr, string(Matrix[0][10]))
        case 11:
            SliceStr = append(SliceStr, string(Matrix[0][11]))
        case 12:
            SliceStr = append(SliceStr, string(Matrix[0][12]))
        case 13:
            SliceStr = append(SliceStr, string(Matrix[0][13]))
        case 14:
            SliceStr = append(SliceStr, string(Matrix[0][14]))
        case 15:
            SliceStr = append(SliceStr, string(Matrix[0][15]))
        
        case 16:
            SliceStr = append(SliceStr, string(Matrix[1][0]))
        case 17:
            SliceStr = append(SliceStr, string(Matrix[1][1]))
        case 18:
            SliceStr = append(SliceStr, string(Matrix[1][2]))
        case 19:
            SliceStr = append(SliceStr, string(Matrix[1][3]))
        case 20:
            SliceStr = append(SliceStr, string(Matrix[1][4]))
        case 21:
            SliceStr = append(SliceStr, string(Matrix[1][5]))
        case 22:
            SliceStr = append(SliceStr, string(Matrix[1][6]))
        case 23:
            SliceStr = append(SliceStr, string(Matrix[1][7]))
        case 24:
            SliceStr = append(SliceStr, string(Matrix[1][8]))
        case 25:
            SliceStr = append(SliceStr, string(Matrix[1][9]))
        case 26:
            SliceStr = append(SliceStr, string(Matrix[1][10]))
        case 27:
            SliceStr = append(SliceStr, string(Matrix[1][11]))
        case 28:
            SliceStr = append(SliceStr, string(Matrix[1][12]))
        case 29:
            SliceStr = append(SliceStr, string(Matrix[1][13]))
        case 30:
            SliceStr = append(SliceStr, string(Matrix[1][14]))
        case 31:
            SliceStr = append(SliceStr, string(Matrix[1][15]))
        
        case 32:
            SliceStr = append(SliceStr, string(Matrix[2][0]))
        case 33:
            SliceStr = append(SliceStr, string(Matrix[2][1]))
        case 34:
            SliceStr = append(SliceStr, string(Matrix[2][2]))
        case 35:
            SliceStr = append(SliceStr, string(Matrix[2][3]))
        case 36:
            SliceStr = append(SliceStr, string(Matrix[2][4]))
        case 37:
            SliceStr = append(SliceStr, string(Matrix[2][5]))
        case 38:
            SliceStr = append(SliceStr, string(Matrix[2][6]))
        case 39:
            SliceStr = append(SliceStr, string(Matrix[2][7]))
        case 40:
            SliceStr = append(SliceStr, string(Matrix[2][8]))
        case 41:
            SliceStr = append(SliceStr, string(Matrix[2][9]))
        case 42:
            SliceStr = append(SliceStr, string(Matrix[2][10]))
        case 43:
            SliceStr = append(SliceStr, string(Matrix[2][11]))
        case 44:
            SliceStr = append(SliceStr, string(Matrix[2][12]))
        case 45:
            SliceStr = append(SliceStr, string(Matrix[2][13]))
        case 46:
            SliceStr = append(SliceStr, string(Matrix[2][14]))
        case 47:
            SliceStr = append(SliceStr, string(Matrix[2][15]))
        
        case 48:
            SliceStr = append(SliceStr, string(Matrix[3][0]))
        case 49:
            SliceStr = append(SliceStr, string(Matrix[3][1]))
        case 50:
            SliceStr = append(SliceStr, string(Matrix[3][2]))
        case 51:
            SliceStr = append(SliceStr, string(Matrix[3][3]))
        case 52:
            SliceStr = append(SliceStr, string(Matrix[3][4]))
        case 53:
            SliceStr = append(SliceStr, string(Matrix[3][5]))
        case 54:
            SliceStr = append(SliceStr, string(Matrix[3][6]))
        case 55:
            SliceStr = append(SliceStr, string(Matrix[3][7]))
        case 56:
            SliceStr = append(SliceStr, string(Matrix[3][8]))
        case 57:
            SliceStr = append(SliceStr, string(Matrix[3][9]))
        case 58:
            SliceStr = append(SliceStr, string(Matrix[3][10]))
        case 59:
            SliceStr = append(SliceStr, string(Matrix[3][11]))
        case 60:
            SliceStr = append(SliceStr, string(Matrix[3][12]))
        case 61:
            SliceStr = append(SliceStr, string(Matrix[3][13]))
        case 62:
            SliceStr = append(SliceStr, string(Matrix[3][14]))
        case 63:
            SliceStr = append(SliceStr, string(Matrix[3][15]))
        
        case 64:
            SliceStr = append(SliceStr, string(Matrix[4][0]))
        case 65:
            SliceStr = append(SliceStr, string(Matrix[4][1]))
        case 66:
            SliceStr = append(SliceStr, string(Matrix[4][2]))
        case 67:
            SliceStr = append(SliceStr, string(Matrix[4][3]))
        case 68:
            SliceStr = append(SliceStr, string(Matrix[4][4]))
        case 69:
            SliceStr = append(SliceStr, string(Matrix[4][5]))
        case 70:
            SliceStr = append(SliceStr, string(Matrix[4][6]))
        case 71:
            SliceStr = append(SliceStr, string(Matrix[4][7]))
        case 72:
            SliceStr = append(SliceStr, string(Matrix[4][8]))
        case 73:
            SliceStr = append(SliceStr, string(Matrix[4][9]))
        case 74:
            SliceStr = append(SliceStr, string(Matrix[4][10]))
        case 75:
            SliceStr = append(SliceStr, string(Matrix[4][11]))
        case 76:
            SliceStr = append(SliceStr, string(Matrix[4][12]))
        case 77:
            SliceStr = append(SliceStr, string(Matrix[4][13]))
        case 78:
            SliceStr = append(SliceStr, string(Matrix[4][14]))
        case 79:
            SliceStr = append(SliceStr, string(Matrix[4][15]))
        
        case 80:
            SliceStr = append(SliceStr, string(Matrix[5][0]))
        case 81:
            SliceStr = append(SliceStr, string(Matrix[5][1]))
        case 82:
            SliceStr = append(SliceStr, string(Matrix[5][2]))
        case 83:
            SliceStr = append(SliceStr, string(Matrix[5][3]))
        case 84:
            SliceStr = append(SliceStr, string(Matrix[5][4]))
        case 85:
            SliceStr = append(SliceStr, string(Matrix[5][5]))
        case 86:
            SliceStr = append(SliceStr, string(Matrix[5][6]))
        case 87:
            SliceStr = append(SliceStr, string(Matrix[5][7]))
        case 88:
            SliceStr = append(SliceStr, string(Matrix[5][8]))
        case 89:
            SliceStr = append(SliceStr, string(Matrix[5][9]))
        case 90:
            SliceStr = append(SliceStr, string(Matrix[5][10]))
        case 91:
            SliceStr = append(SliceStr, string(Matrix[5][11]))
        case 92:
            SliceStr = append(SliceStr, string(Matrix[5][12]))
        case 93:
            SliceStr = append(SliceStr, string(Matrix[5][13]))
        case 94:
            SliceStr = append(SliceStr, string(Matrix[5][14]))
        case 95:
            SliceStr = append(SliceStr, string(Matrix[5][15]))
        
        case 96:
            SliceStr = append(SliceStr, string(Matrix[6][0]))
        case 97:
            SliceStr = append(SliceStr, string(Matrix[6][1]))
        case 98:
            SliceStr = append(SliceStr, string(Matrix[6][2]))
        case 99:
            SliceStr = append(SliceStr, string(Matrix[6][3]))
        case 100:
            SliceStr = append(SliceStr, string(Matrix[6][4]))
        case 101:
            SliceStr = append(SliceStr, string(Matrix[6][5]))
        case 102:
            SliceStr = append(SliceStr, string(Matrix[6][6]))
        case 103:
            SliceStr = append(SliceStr, string(Matrix[6][7]))
        case 104:
            SliceStr = append(SliceStr, string(Matrix[6][8]))
        case 105:
            SliceStr = append(SliceStr, string(Matrix[6][9]))
        case 106:
            SliceStr = append(SliceStr, string(Matrix[6][10]))
        case 107:
            SliceStr = append(SliceStr, string(Matrix[6][11]))
        case 108:
            SliceStr = append(SliceStr, string(Matrix[6][12]))
        case 109:
            SliceStr = append(SliceStr, string(Matrix[6][13]))
        case 110:
            SliceStr = append(SliceStr, string(Matrix[6][14]))
        case 111:
            SliceStr = append(SliceStr, string(Matrix[6][15]))
        
        case 112:
            SliceStr = append(SliceStr, string(Matrix[7][0]))
        case 113:
            SliceStr = append(SliceStr, string(Matrix[7][1]))
        case 114:
            SliceStr = append(SliceStr, string(Matrix[7][2]))
        case 115:
            SliceStr = append(SliceStr, string(Matrix[7][3]))
        case 116:
            SliceStr = append(SliceStr, string(Matrix[7][4]))
        case 117:
            SliceStr = append(SliceStr, string(Matrix[7][5]))
        case 118:
            SliceStr = append(SliceStr, string(Matrix[7][6]))
        case 119:
            SliceStr = append(SliceStr, string(Matrix[7][7]))
        case 120:
            SliceStr = append(SliceStr, string(Matrix[7][8]))
        case 121:
            SliceStr = append(SliceStr, string(Matrix[7][9]))
        case 122:
            SliceStr = append(SliceStr, string(Matrix[7][10]))
        case 123:
            SliceStr = append(SliceStr, string(Matrix[7][11]))
        case 124:
            SliceStr = append(SliceStr, string(Matrix[7][12]))
        case 125:
            SliceStr = append(SliceStr, string(Matrix[7][13]))
        case 126:
            SliceStr = append(SliceStr, string(Matrix[7][14]))
        case 127:
            SliceStr = append(SliceStr, string(Matrix[7][15]))
        
        case 128:
            SliceStr = append(SliceStr, string(Matrix[8][0]))
        case 129:
            SliceStr = append(SliceStr, string(Matrix[8][1]))
        case 130:
            SliceStr = append(SliceStr, string(Matrix[8][2]))
        case 131:
            SliceStr = append(SliceStr, string(Matrix[8][3]))
        case 132:
            SliceStr = append(SliceStr, string(Matrix[8][4]))
        case 133:
            SliceStr = append(SliceStr, string(Matrix[8][5]))
        case 134:
            SliceStr = append(SliceStr, string(Matrix[8][6]))
        case 135:
            SliceStr = append(SliceStr, string(Matrix[8][7]))
        case 136:
            SliceStr = append(SliceStr, string(Matrix[8][8]))
        case 137:
            SliceStr = append(SliceStr, string(Matrix[8][9]))
        case 138:
            SliceStr = append(SliceStr, string(Matrix[8][10]))
        case 139:
            SliceStr = append(SliceStr, string(Matrix[8][11]))
        case 140:
            SliceStr = append(SliceStr, string(Matrix[8][12]))
        case 141:
            SliceStr = append(SliceStr, string(Matrix[8][13]))
        case 142:
            SliceStr = append(SliceStr, string(Matrix[8][14]))
        case 143:
            SliceStr = append(SliceStr, string(Matrix[8][15]))
        
        case 144:
            SliceStr = append(SliceStr, string(Matrix[9][0]))
        case 145:
            SliceStr = append(SliceStr, string(Matrix[9][1]))
        case 146:
            SliceStr = append(SliceStr, string(Matrix[9][2]))
        case 147:
            SliceStr = append(SliceStr, string(Matrix[9][3]))
        case 148:
            SliceStr = append(SliceStr, string(Matrix[9][4]))
        case 149:
            SliceStr = append(SliceStr, string(Matrix[9][5]))
        case 150:
            SliceStr = append(SliceStr, string(Matrix[9][6]))
        case 151:
            SliceStr = append(SliceStr, string(Matrix[9][7]))
        case 152:
            SliceStr = append(SliceStr, string(Matrix[9][8]))
        case 153:
            SliceStr = append(SliceStr, string(Matrix[9][9]))
        case 154:
            SliceStr = append(SliceStr, string(Matrix[9][10]))
        case 155:
            SliceStr = append(SliceStr, string(Matrix[9][11]))
        case 156:
            SliceStr = append(SliceStr, string(Matrix[9][12]))
        case 157:
            SliceStr = append(SliceStr, string(Matrix[9][13]))
        case 158:
            SliceStr = append(SliceStr, string(Matrix[9][14]))
        case 159:
            SliceStr = append(SliceStr, string(Matrix[9][15]))
        
        case 160:
            SliceStr = append(SliceStr, string(Matrix[10][0]))
        case 161:
            SliceStr = append(SliceStr, string(Matrix[10][1]))
        case 162:
            SliceStr = append(SliceStr, string(Matrix[10][2]))
        case 163:
            SliceStr = append(SliceStr, string(Matrix[10][3]))
        case 164:
            SliceStr = append(SliceStr, string(Matrix[10][4]))
        case 165:
            SliceStr = append(SliceStr, string(Matrix[10][5]))
        case 166:
            SliceStr = append(SliceStr, string(Matrix[10][6]))
        case 167:
            SliceStr = append(SliceStr, string(Matrix[10][7]))
        case 168:
            SliceStr = append(SliceStr, string(Matrix[10][8]))
        case 169:
            SliceStr = append(SliceStr, string(Matrix[10][9]))
        case 170:
            SliceStr = append(SliceStr, string(Matrix[10][10]))
        case 171:
            SliceStr = append(SliceStr, string(Matrix[10][11]))
        case 172:
            SliceStr = append(SliceStr, string(Matrix[10][12]))
        case 173:
            SliceStr = append(SliceStr, string(Matrix[10][13]))
        case 174:
            SliceStr = append(SliceStr, string(Matrix[10][14]))
        case 175:
            SliceStr = append(SliceStr, string(Matrix[10][15]))
        
        case 176:
            SliceStr = append(SliceStr, string(Matrix[11][0]))
        case 177:
            SliceStr = append(SliceStr, string(Matrix[11][1]))
        case 178:
            SliceStr = append(SliceStr, string(Matrix[11][2]))
        case 179:
            SliceStr = append(SliceStr, string(Matrix[11][3]))
        case 180:
            SliceStr = append(SliceStr, string(Matrix[11][4]))
        case 181:
            SliceStr = append(SliceStr, string(Matrix[11][5]))
        case 182:
            SliceStr = append(SliceStr, string(Matrix[11][6]))
        case 183:
            SliceStr = append(SliceStr, string(Matrix[11][7]))
        case 184:
            SliceStr = append(SliceStr, string(Matrix[11][8]))
        case 185:
            SliceStr = append(SliceStr, string(Matrix[11][9]))
        case 186:
            SliceStr = append(SliceStr, string(Matrix[11][10]))
        case 187:
            SliceStr = append(SliceStr, string(Matrix[11][11]))
        case 188:
            SliceStr = append(SliceStr, string(Matrix[11][12]))
        case 189:
            SliceStr = append(SliceStr, string(Matrix[11][13]))
        case 190:
            SliceStr = append(SliceStr, string(Matrix[11][14]))
        case 191:
            SliceStr = append(SliceStr, string(Matrix[11][15]))
        
        case 192:
            SliceStr = append(SliceStr, string(Matrix[12][0]))
        case 193:
            SliceStr = append(SliceStr, string(Matrix[12][1]))
        case 194:
            SliceStr = append(SliceStr, string(Matrix[12][2]))
        case 195:
            SliceStr = append(SliceStr, string(Matrix[12][3]))
        case 196:
            SliceStr = append(SliceStr, string(Matrix[12][4]))
        case 197:
            SliceStr = append(SliceStr, string(Matrix[12][5]))
        case 198:
            SliceStr = append(SliceStr, string(Matrix[12][6]))
        case 199:
            SliceStr = append(SliceStr, string(Matrix[12][7]))
        case 200:
            SliceStr = append(SliceStr, string(Matrix[12][8]))
        case 201:
            SliceStr = append(SliceStr, string(Matrix[12][9]))
        case 202:
            SliceStr = append(SliceStr, string(Matrix[12][10]))
        case 203:
            SliceStr = append(SliceStr, string(Matrix[12][11]))
        case 204:
            SliceStr = append(SliceStr, string(Matrix[12][12]))
        case 205:
            SliceStr = append(SliceStr, string(Matrix[12][13]))
        case 206:
            SliceStr = append(SliceStr, string(Matrix[12][14]))
        case 207:
            SliceStr = append(SliceStr, string(Matrix[12][15]))
        
        case 208:
            SliceStr = append(SliceStr, string(Matrix[13][0]))
        case 209:
            SliceStr = append(SliceStr, string(Matrix[13][1]))
        case 210:
            SliceStr = append(SliceStr, string(Matrix[13][2]))
        case 211:
            SliceStr = append(SliceStr, string(Matrix[13][3]))
        case 212:
            SliceStr = append(SliceStr, string(Matrix[13][4]))
        case 213:
            SliceStr = append(SliceStr, string(Matrix[13][5]))
        case 214:
            SliceStr = append(SliceStr, string(Matrix[13][6]))
        case 215:
            SliceStr = append(SliceStr, string(Matrix[13][7]))
        case 216:
            SliceStr = append(SliceStr, string(Matrix[13][8]))
        case 217:
            SliceStr = append(SliceStr, string(Matrix[13][9]))
        case 218:
            SliceStr = append(SliceStr, string(Matrix[13][10]))
        case 219:
            SliceStr = append(SliceStr, string(Matrix[13][11]))
        case 220:
            SliceStr = append(SliceStr, string(Matrix[13][12]))
        case 221:
            SliceStr = append(SliceStr, string(Matrix[13][13]))
        case 222:
            SliceStr = append(SliceStr, string(Matrix[13][14]))
        case 223:
            SliceStr = append(SliceStr, string(Matrix[13][15]))
        
        case 224:
            SliceStr = append(SliceStr, string(Matrix[14][0]))
        case 225:
            SliceStr = append(SliceStr, string(Matrix[14][1]))
        case 226:
            SliceStr = append(SliceStr, string(Matrix[14][2]))
        case 227:
            SliceStr = append(SliceStr, string(Matrix[14][3]))
        case 228:
            SliceStr = append(SliceStr, string(Matrix[14][4]))
        case 229:
            SliceStr = append(SliceStr, string(Matrix[14][5]))
        case 230:
            SliceStr = append(SliceStr, string(Matrix[14][6]))
        case 231:
            SliceStr = append(SliceStr, string(Matrix[14][7]))
        case 232:
            SliceStr = append(SliceStr, string(Matrix[14][8]))
        case 233:
            SliceStr = append(SliceStr, string(Matrix[14][9]))
        case 234:
            SliceStr = append(SliceStr, string(Matrix[14][10]))
        case 235:
            SliceStr = append(SliceStr, string(Matrix[14][11]))
        case 236:
            SliceStr = append(SliceStr, string(Matrix[14][12]))
        case 237:
            SliceStr = append(SliceStr, string(Matrix[14][13]))
        case 238:
            SliceStr = append(SliceStr, string(Matrix[14][14]))
        case 239:
            SliceStr = append(SliceStr, string(Matrix[14][15]))
        
        case 240:
            SliceStr = append(SliceStr, string(Matrix[15][0]))
        case 241:
            SliceStr = append(SliceStr, string(Matrix[15][1]))
        case 242:
            SliceStr = append(SliceStr, string(Matrix[15][2]))
        case 243:
            SliceStr = append(SliceStr, string(Matrix[15][3]))
        case 244:
            SliceStr = append(SliceStr, string(Matrix[15][4]))
        case 245:
            SliceStr = append(SliceStr, string(Matrix[15][5]))
        case 246:
            SliceStr = append(SliceStr, string(Matrix[15][6]))
        case 247:
            SliceStr = append(SliceStr, string(Matrix[15][7]))
        case 248:
            SliceStr = append(SliceStr, string(Matrix[15][8]))
        case 249:
            SliceStr = append(SliceStr, string(Matrix[15][9]))
        case 250:
            SliceStr = append(SliceStr, string(Matrix[15][10]))
        case 251:
            SliceStr = append(SliceStr, string(Matrix[15][11]))
        case 252:
            SliceStr = append(SliceStr, string(Matrix[15][12]))
        case 253:
            SliceStr = append(SliceStr, string(Matrix[15][13]))
        case 254:
            SliceStr = append(SliceStr, string(Matrix[15][14]))
        case 255:
            SliceStr = append(SliceStr, string(Matrix[15][15]))
        }
    }
    
    //Converting the Slice of Strings to a String as final step
    for i := 0; i < len(SliceStr); i++ {
        result = result + SliceStr[i]
    }
    
    return result
}

func CharacterMatrix() [16][16]rune {
    //Digits Block 10 runes
    C000 := '0' //U+0030    [48] Digit Zero
    C001 := '1' //U+0031    [49] Digit One
    C002 := '2' //U+0032    [50] Digit Two
    C003 := '3' //U+0033    [51] Digit Three
    C004 := '4' //U+0034    [52] Digit Four
    C005 := '5' //U+0035    [53] Digit Five
    C006 := '6' //U+0036    [54] Digit Six
    C007 := '7' //U+0037    [55] Digit Seven
    C008 := '8' //U+0038    [56] Digit Eight
    C009 := '9' //U+0039    [57] Digit Nine
    
    //Currencies Block 10 runes
    C010 := 'Ѻ' //U+047A    [209 186] Cyrillic Capital Letter Round Omega (Ourobos Currency)
    C011 := '₿' //U+20BF    [226 130 191] Bitcoin Sign
    C012 := '$' //U+0024    [36] Dollar Sign
    C013 := '¢' //U+00A2    [194 162] Cent Sign
    C014 := '€' //U+20AC    [226 130 172] Euro Sign
    C015 := '£' //U+00A3    [194 163] Pound Sign
    C016 := '¥' //U+00A5    [194 165] Yen Sign
    C017 := '₱' //U+20B1    [226 130 177] Peso Sign
    C018 := '₳' //U+20B3    [226 130 179] Austral Sign (AURYN Currency)
    C019 := '∇' //U+2207    [226 136 135] Nabla (TALOS Currency)
    
    //Latin Capital Letters 26 runes
    C020 := 'A' //U+0041    [65] Latin Capital Letter A
    C021 := 'B' //U+0042    [66] Latin Capital Letter B
    C022 := 'C' //U+0043    [67] Latin Capital Letter C
    C023 := 'D' //U+0044    [68] Latin Capital Letter D
    C024 := 'E' //U+0045    [69] Latin Capital Letter E
    C025 := 'F' //U+0046    [70] Latin Capital Letter F
    C026 := 'G' //U+0047    [71] Latin Capital Letter G
    C027 := 'H' //U+0048    [72] Latin Capital Letter G
    C028 := 'I' //U+0049    [73] Latin Capital Letter I
    C029 := 'J' //U+004A    [74] Latin Capital Letter J
    C030 := 'K' //U+004B    [75] Latin Capital Letter K
    C031 := 'L' //U+004C    [76] Latin Capital Letter L
    C032 := 'M' //U+004D    [77] Latin Capital Letter M
    C033 := 'N' //U+004E    [78] Latin Capital Letter N
    C034 := 'O' //U+004F    [79] Latin Capital Letter O
    C035 := 'P' //U+0050    [80] Latin Capital Letter P
    C036 := 'Q' //U+0051    [81] Latin Capital Letter Q
    C037 := 'R' //U+0052    [82] Latin Capital Letter R
    C038 := 'S' //U+0053    [83] Latin Capital Letter S
    C039 := 'T' //U+0054    [84] Latin Capital Letter T
    C040 := 'U' //U+0055    [85] Latin Capital Letter U
    C041 := 'V' //U+0056    [86] Latin Capital Letter V
    C042 := 'W' //U+0057    [87] Latin Capital Letter W
    C043 := 'X' //U+0058    [88] Latin Capital Letter X
    C044 := 'Y' //U+0059    [89] Latin Capital Letter Y
    C045 := 'Z' //U+005A    [90] Latin Capital Letter Z
    
    //Latin Small Letters 26 runes
    C046 := 'a' //U+0061    [97] Latin Small Letter A
    C047 := 'b' //U+0062    [98] Latin Small Letter B
    C048 := 'c' //U+0063    [99] Latin Small Letter C
    C049 := 'd' //U+0064    [100] Latin Small Letter D
    C050 := 'e' //U+0065    [101] Latin Small Letter E
    C051 := 'f' //U+0066    [102] Latin Small Letter F
    C052 := 'g' //U+0067    [103] Latin Small Letter G
    C053 := 'h' //U+0068    [104] Latin Small Letter H
    C054 := 'i' //U+0069    [105] Latin Small Letter I
    C055 := 'j' //U+006A    [106] Latin Small Letter J
    C056 := 'k' //U+006B    [107] Latin Small Letter K
    C057 := 'l' //U+006C    [108] Latin Small Letter L
    C058 := 'm' //U+006D    [108] Latin Small Letter M
    C059 := 'n' //U+006E    [110] Latin Small Letter N
    C060 := 'o' //U+006F    [111] Latin Small Letter O
    C061 := 'p' //U+0070    [112] Latin Small Letter P
    C062 := 'q' //U+0071    [113] Latin Small Letter Q
    C063 := 'r' //U+0072    [114] Latin Small Letter R
    C064 := 's' //U+0073    [115] Latin Small Letter S
    C065 := 't' //U+0074    [116] Latin Small Letter T
    C066 := 'u' //U+0075    [117] Latin Small Letter U
    C067 := 'v' //U+0076    [118] Latin Small Letter V
    C068 := 'w' //U+0077    [119] Latin Small Letter W
    C069 := 'x' //U+0078    [120] Latin Small Letter X
    C070 := 'y' //U+0079    [121] Latin Small Letter Y
    C071 := 'z' //U+007A    [122] Latin Small Letter Z
    
    //Latin Extended Capital Letters 53 runes
    C072 := 'Æ' //U+00C6    [195 134] Latin Capital Letter Ae                   French, Danish, Norwegian, Latin, Icelanding (missing thorn)
    C073 := 'Œ' //U+0152    [197 146] Latin Capital Letter Oe                   French, Latin
    C074 := 'Á' //U+00C1    [195 129] Latin Capital Letter A with Acute         Spanish, Czech, Portuguese, Icelanding (missing thorn),
    C075 := 'Ă' //U+0102    [196 130] Latin Capital Letter A with Breve         Romanian
    C076 := 'Â' //U+00C2    [195 130] Latin Capital Letter A with Circumflex    French, Romanian, Portuguese, Czech,
    C077 := 'Ä' //U+00C4    [195 132] Latin Capital Letter A with Diaeresis     German, Swedish, Finnish, Estonian
    C078 := 'À' //U+00C0    [195 128] Latin Capital Letter A with Grave         French, Portuguese, Italian
    C079 := 'Ą' //U+0104    [196 132] Latin Capital Letter A with Ogonek        Polish
    C080 := 'Å' //U+00C5    [195 133] Latin Capital Letter A with Ring Above    Danish, Norwegian, Swedish, Finnish
    C081 := 'Ã' //U+00C3    [195 131] Latin Capital Letter A with Tilde         Portuguese
    C082 := 'Ć' //U+0106    [196 134] Latin Capital Letter C with Acute         Polish, Croatian, Serbian-Bosnian(latinized)
    C083 := 'Č' //U+010C    [196 140] Latin Capital Letter C with Caron         Czech, Croatian, Serbian-Bosnian(latinized), Slovenian
    C084 := 'Ç' //U+00C7    [195 135] Latin Capital Letter C with Cedilla       French, Kurmanji, Portuguese, Kurdish, Albanian, Turkish
    C085 := 'Ď' //U+010E    [196 142] Latin Capital Letter D with Caron         Czech, Serbian(latinized)
    C086 := 'Đ' //U+0110    [196 144] Latin Capital Letter D with Stroke        Croatian, Serbian-Bosnian(latinized)
    C087 := 'É' //U+00C9    [195 137] Latin Capital Letter E with Acute         Kurmanji, Icelanding (missing thorn), Spanish, Czech, Portuguese, Italian
    C088 := 'Ě' //U+011A    [196 154] Latin Capital Letter E with Caron         Czech
    C089 := 'Ê' //U+00CA    [195 138] Latin Capital Letter E with Circumflex    French, Kurmanji, Kurdish
    C090 := 'Ë' //U+00CB    [195 139] Latin Capital Letter E with Diaeresis     French, Albanian
    C091 := 'È' //U+00C8    [195 136] Latin Capital Letter E with Grave         French, Portuguese, Italian
    C092 := 'Ę' //U+0118    [196 152] Latin Capital Letter E with Ogonek        Polish
    C093 := 'Ğ' //U+011E    [196 158] Latin Capital Letter G with Breve         Turkish
    C094 := 'Í' //U+00CD    [195 141] Latin Capital Letter I with Acute         Icelanding (missing thorn), Spanish, Czech, Portuguese, Italian
    C095 := 'Î' //U+00CE    [195 142] Latin Capital Letter I with Circumflex    French, Kurmanji, Romanian, Kurdish, Italian
    C096 := 'Ï' //U+00CF    [195 143] Latin Capital Letter I with Diaeresis     French
    C097 := 'Ì' //U+00CC    [195 140] Latin Capital Letter I with Grave         Portuguese, Italian
    C098 := 'Ł' //U+0141    [197 129] Latin Capital Letter L with Stroke        Polish
    C099 := 'Ń' //U+0143    [197 131] Latin Capital Letter N with Acute         Polish
    C100 := 'Ñ' //U+00D1    [195 145] Latin Capital Letter N with Tilde         Spanish, Czech
    C101 := 'Ó' //U+00D3    [195 147] Latin Capital Letter O with Acute         Icelanding (missing thorn), Spanish, Czech, Portuguese, Italian, Polish
    C102 := 'Ô' //U+00D4    [195 148] Latin Capital Letter O with Circumflex    French, Portuguese
    C103 := 'Ö' //U+00D6    [195 150] Latin Capital Letter O with Diaeresis     Icelanding (missing thorn), German, Swedish, Finnish, Turkish
    C104 := 'Ò' //U+00D2    [195 146] Latin Capital Letter O with Grave         Portuguese, Italian
    C105 := 'Ø' //U+00D8    [195 152] Latin Capital Letter O with Stroke        Danish, Norwegian
    C106 := 'Õ' //U+00D5    [195 149] Latin Capital Letter O with Tilde         Portuguese
    C107 := 'Ř' //U+0158    [197 152] Latin Capital Letter R with Caron         Czech
    C108 := 'Ś' //U+015A    [197 154] Latin Capital Letter S with Acute         Polish
    C109 := 'Š' //U+0160    [197 160] Latin Capital Letter S with Caron         Czech, Estonian, Croatian, Serbian-Bosnian(latinized), Slovenian
    C110 := 'Ş' //U+015E    [197 158] Latin Capital Letter S with Cedilla       Kurdish, Turkish
    C111 := 'Ș' //U+0218    [200 152] Latin Capital Letter S with Comma Below   Kurmanji, Romanian
    C112 := 'Þ' //U+00DE    [195 158] Latin Capital Letter Thorn                Icelandic
    C113 := 'Ť' //U+0164    [197 164] Latin Capital Letter T with Caron         Czech
    C114 := 'Ț' //U+021A    [200 154] Latin Capital Letter T with Comma Below   Romanian
    C115 := 'Ú' //U+00DA    [195 154] Latin Capital Letter U with Acute         Icelanding (missing thorn), Spanish, Czech, Portuguese, Italian
    C116 := 'Û' //U+00DB    [195 155] Latin Capital Letter U with Circumflex    French, Kurmanji, Kurdish
    C117 := 'Ü' //U+00DC    [195 156] Latin Capital Letter U with Diaeresis     French, Spanish, German, Estonian, Turkish
    C118 := 'Ù' //U+00D9    [195 153] Latin Capital Letter U with Grave         French, Portuguese, Italian
    C119 := 'Ů' //U+016E    [197 174] Latin Capital Letter U with Ring Above    Czech
    C120 := 'Ý' //U+00DD    [195 157] Latin Capital Letter Y with Acute         Icelanding (missing thorn), Czech
    C121 := 'Ÿ' //U+00DC    [195 184] Latin Capital Letter U with Diaeresis     French
    C122 := 'Ź' //U+0179    [197 185] Latin Capital Letter Z with Acute         Polish
    C123 := 'Ž' //U+017D    [197 189] Latin Capital Letter Z with Caron         Czech, Estonian, Croatian, Serbian-Bosnian(latinized), Slovenian
    C124 := 'Ż' //U+017B    [197 187] Latin Capital Letter Z with Dot Above     Polish
    
    //Latin Extended Small Letters 54 runes
    C125 := 'æ' //U+00E6    [195 166] Latin Small Letter Ae
    C126 := 'œ' //U+0153    [197 147] Latin Small Letter Oe
    C127 := 'á' //U+00E1    [195 161] Latin Small Letter A with Acute
    C128 := 'ă' //U+0103    [196 131] Latin Small Letter A with Breve
    C129 := 'â' //U+00E2    [195 162] Latin Small Letter A with Circumflex
    C130 := 'ä' //U+00E4    [195 164] Latin Small Letter A with Diaeresis
    C131 := 'à' //U+00E0    [195 160] Latin Small Letter A with Grave
    C132 := 'ą' //U+0105    [196 133] Latin Small Letter A with Ogonek
    C133 := 'å' //U+00E5    [195 165] Latin Small Letter A with Ring Above
    C134 := 'ã' //U+00E3    [195 163] Latin Small Letter A with Tilde
    C135 := 'ć' //U+0107    [196 135] Latin Small Letter C with Acute
    C136 := 'č' //U+010D    [196 141] Latin Small Letter C with Caron
    C137 := 'ç' //U+00E7    [195 167] Latin Small Letter C with Cedilla
    C138 := 'ď' //U+010F    [196 143] Latin Small Letter D with Caron
    C139 := 'đ' //U+0111    [196 145] Latin Small Letter D with Stroke
    C140 := 'é' //U+00E9    [195 169] Latin Small Letter E with Acute
    C141 := 'ě' //U+011B    [196 155] Latin Small Letter E with Caron
    C142 := 'ê' //U+00EA    [195 170] Latin Small Letter E with Circumflex
    C143 := 'ë' //U+00EB    [195 171] Latin Small Letter E with Diaeresis
    C144 := 'è' //U+00E8    [195 168] Latin Small Letter E with Grave
    C145 := 'ę' //U+0119    [196 153] Latin Small Letter E with Ogonek
    C146 := 'ğ' //U+011F    [196 159] Latin Small Letter G with Breve
    C147 := 'í' //U+00ED    [195 173] Latin Small Letter I with Acute
    C148 := 'î' //U+00EE    [195 174] Latin Small Letter I with Circumflex
    C149 := 'ï' //U+00EF    [195 175] Latin Small Letter I with Diaeresis
    C150 := 'ì' //U+00EC    [195 172] Latin Small Letter I with Grave
    C151 := 'ł' //U+0142    [197 130] Latin Small Letter L with Stroke
    C152 := 'ń' //U+0144    [197 132] Latin Small Letter N with Acute
    C153 := 'ñ' //U+00F1    [195 177] Latin Small Letter N with Tilde
    C154 := 'ó' //U+00F3    [195 179] Latin Small Letter O with Acute
    C155 := 'ô' //U+00F4    [195 180] Latin Small Letter O with Circumflex
    C156 := 'ö' //U+00F6    [195 182] Latin Small Letter O with Diaeresis
    C157 := 'ò' //U+00F2    [195 178] Latin Small Letter O with Grave
    C158 := 'ø' //U+00F8    [195 184] Latin Small Letter O with Stroke
    C159 := 'õ' //U+00F5    [195 181] Latin Small Letter O with Tilde
    C160 := 'ř' //U+0159    [197 153] Latin Small Letter R with Caron
    C161 := 'ś' //U+015B    [197 155] Latin Small Letter S with Acute
    C162 := 'š' //U+0161    [197 161] Latin Small Letter S with Caron
    C163 := 'ş' //U+015F    [197 159] Latin Small Letter S with Cedilla
    C164 := 'ș' //U+0219    [200 153] Latin Small Letter S with Comma Below
    C165 := 'þ' //U+00FE    [195 190] Latin Small Letter Thorn
    C166 := 'ť' //U+0165    [197 165] Latin Small Letter T with Caron
    C167 := 'ț' //U+021B    [200 155] Latin Small Letter T with Comma Below
    C168 := 'ú' //U+00FA    [195 186] Latin Small Letter U with Acute
    C169 := 'û' //U+00FB    [195 187] Latin Small Letter U with Circumflex
    C170 := 'ü' //U+00FC    [195 188] Latin Small Letter U with Diaeresis
    C171 := 'ù' //U+00F9    [195 185] Latin Small Letter U with Grave
    C172 := 'ů' //U+016F    [197 175] Latin Small Letter U with Ring Above
    C173 := 'ý' //U+00FD    [195 189] Latin Small Letter Y with Acute
    C174 := 'ÿ' //U+00FF    [195 191] Latin Small Letter Y with Diaeresis
    C175 := 'ź' //U+017A    [197 186] Latin Small Letter Z with Acute
    C176 := 'ž' //U+017E    [197 190] Latin Small Letter Z with Caron
    C177 := 'ż' //U+017C    [197 188] Latin Small Letter Z with Dot Above
    C178 := 'ß' //U+00DF    [195 159] Latin Small Letter Sharp S
    
    //Greek Capital Letters 10 runes
    C179 := 'Γ' //U+0393    [206 147] Greek Capital Letter Gamma
    C180 := 'Δ' //U+0394    [206 148] Greek Capital Letter Delta
    C181 := 'Θ' //U+0398    [206 152] Greek Capital Letter Theta
    C182 := 'Λ' //U+039B    [206 155] Greek Capital Letter Lambda
    C183 := 'Ξ' //U+039E    [206 158] Greek Capital Letter Xi
    C184 := 'Π' //U+03A0    [206 160] Greek Capital Letter Pi
    C185 := 'Σ' //U+03A3    [206 163] Greek Capital Letter Sigma
    C186 := 'Φ' //U+03A6    [206 166] Greek Capital Letter Phi
    C187 := 'Ψ' //U+03A8    [206 168] Greek Capital Letter Psi
    C188 := 'Ω' //U+03A9    [206 169] Greek Capital Letter Omega
    
    //Greek Small Letters 23 runes
    C189 := 'α' //U+03B1    [206 177] Greek Small Letter Alpha
    C190 := 'β' //U+03B2    [206 178] Greek Small Letter Beta
    C191 := 'γ' //U+03B3    [206 179] Greek Small Letter Gamma
    C192 := 'δ' //U+03B4    [206 180] Greek Small Letter Delta
    C193 := 'ε' //U+03B5    [206 181] Greek Small Letter Epsilon
    C194 := 'ζ' //U+03B6    [206 182] Greek Small Letter Zeta
    C195 := 'η' //U+03B7    [206 183] Greek Small Letter Eta
    C196 := 'θ' //U+03B8    [206 184] Greek Small Letter Theta
    C197 := 'ι' //U+03B9    [206 185] Greek Small Letter Iota
    C198 := 'κ' //U+03BA    [206 186] Greek Small Letter Kappa
    C199 := 'λ' //U+03BB    [206 187] Greek Small Letter Lambda
    C200 := 'μ' //U+03BC    [206 188] Greek Small Letter Mu
    C201 := 'ν' //U+03BD    [206 189] Greek Small Letter Nu
    C202 := 'ξ' //U+03BE    [206 190] Greek Small Letter Xi
    C203 := 'π' //U+03C0    [206 192] Greek Small Letter Pi
    C204 := 'ρ' //U+03C1    [206 193] Greek Small Letter Rho
    C205 := 'σ' //U+03C3    [206 195] Greek Small Letter Sigma
    C206 := 'ς' //U+03C2    [206 194] Greek Small Letter Final Sigma
    C207 := 'τ' //U+03C4    [206 196] Greek Small Letter Tau
    C208 := 'φ' //U+03C6    [206 198] Greek Small Letter Phi
    C209 := 'χ' //U+03C7    [206 199] Greek Small Letter Chi
    C210 := 'ψ' //U+03C8    [206 200] Greek Small Letter Psi
    C211 := 'ω' //U+03C9    [206 201] Greek Small Letter Omega
    
    //Cyrillic Capital Letters 19 runes
    C212 := 'Б' //U+0411    [208 145] Cyrillic Capital Letter Be
    C213 := 'Д' //U+0414    [208 148] Cyrillic Capital Letter De
    C214 := 'Ж' //U+0416    [208 150] Cyrillic Capital Letter Zhe
    C215 := 'З' //U+0417    [208 151] Cyrillic Capital Letter Ze
    C216 := 'И' //U+0418    [208 152] Cyrillic Capital Letter I
    C217 := 'Й' //U+0419    [208 153] Cyrillic Capital Letter Short I
    C218 := 'Л' //U+041B    [208 155] Cyrillic Capital Letter El
    C219 := 'П' //U+041F    [208 159] Cyrillic Capital Letter Pe
    C220 := 'У' //U+0423    [208 163] Cyrillic Capital Letter U
    C221 := 'Ц' //U+0426    [208 166] Cyrillic Capital Letter Tse
    C222 := 'Ч' //U+0427    [208 167] Cyrillic Capital Letter Che
    C223 := 'Ш' //U+0428    [208 168] Cyrillic Capital Letter Sha
    C224 := 'Щ' //U+0429    [208 169] Cyrillic Capital Letter Shcha
    C225 := 'Ъ' //U+042A    [208 170] Cyrillic Capital Letter Hard Sign
    C226 := 'Ы' //U+042B    [208 171] Cyrillic Capital Letter Yeru
    C227 := 'Ь' //U+042C    [208 172] Cyrillic Capital Letter Soft Sign
    C228 := 'Э' //U+042D    [208 173] Cyrillic Capital Letter E
    C229 := 'Ю' //U+042E    [208 174] Cyrillic Capital Letter Yu
    C230 := 'Я' //U+042F    [208 175] Cyrillic Capital Letter Ya
    
    //Cyrillic Small Letters 25 runes
    C231 := 'б' //U+0431    [208 177] Cyrillic Small Letter Be
    C232 := 'в' //U+0432    [208 178] Cyrillic Small Letter Ve
    C233 := 'д' //U+0434    [208 180] Cyrillic Small Letter De
    C234 := 'ж' //U+0436    [208 182] Cyrillic Small Letter Zhe
    C235 := 'з' //U+0437    [208 183] Cyrillic Small Letter Ze
    C236 := 'и' //U+0438    [208 184] Cyrillic Small Letter I
    C237 := 'й' //U+0439    [208 185] Cyrillic Small Letter Short I
    C238 := 'к' //U+043A    [208 186] Cyrillic Small Letter Ka
    C239 := 'л' //U+043B    [208 187] Cyrillic Small Letter El
    C240 := 'м' //U+043C    [208 188] Cyrillic Small Letter Em
    C241 := 'н' //U+043D    [208 189] Cyrillic Small Letter En
    C242 := 'п' //U+043F    [208 191] Cyrillic Small Letter Pe
    C243 := 'т' //U+0442    [209 130] Cyrillic Small Letter Te
    C244 := 'у' //U+0443    [209 131] Cyrillic Small Letter U
    C245 := 'ф' //U+0444    [209 132] Cyrillic Small Letter Ef
    C246 := 'ц' //U+0446    [209 134] Cyrillic Small Letter Tse
    C247 := 'ч' //U+0447    [209 135] Cyrillic Small Letter Che
    C248 := 'ш' //U+0448    [209 136] Cyrillic Small Letter Sha
    C249 := 'щ' //U+0449    [209 137] Cyrillic Small Letter Shcha
    C250 := 'ъ' //U+044A    [209 138] Cyrillic Small Letter Hard Sign
    C251 := 'ы' //U+044B    [209 139] Cyrillic Small Letter Yeru
    C252 := 'ь' //U+044C    [209 140] Cyrillic Small Letter Soft Sign
    C253 := 'э' //U+044D    [209 141] Cyrillic Small Letter E
    C254 := 'ю' //U+044E    [209 142] Cyrillic Small Letter Yu
    C255 := 'я' //U+044F    [209 143] Cyrillic Small Letter Ya
    
    Row00 := [...]rune{C000, C001, C002, C003, C004, C005, C006, C007, C008, C009, C010, C011, C012, C013, C014, C015}
    Row01 := [...]rune{C016, C017, C018, C019, C020, C021, C022, C023, C024, C025, C026, C027, C028, C029, C030, C031}
    Row02 := [...]rune{C032, C033, C034, C035, C036, C037, C038, C039, C040, C041, C042, C043, C044, C045, C046, C047}
    Row03 := [...]rune{C048, C049, C050, C051, C052, C053, C054, C055, C056, C057, C058, C059, C060, C061, C062, C063}
    Row04 := [...]rune{C064, C065, C066, C067, C068, C069, C070, C071, C072, C073, C074, C075, C076, C077, C078, C079}
    Row05 := [...]rune{C080, C081, C082, C083, C084, C085, C086, C087, C088, C089, C090, C091, C092, C093, C094, C095}
    Row06 := [...]rune{C096, C097, C098, C099, C100, C101, C102, C103, C104, C105, C106, C107, C108, C109, C110, C111}
    Row07 := [...]rune{C112, C113, C114, C115, C116, C117, C118, C119, C120, C121, C122, C123, C124, C125, C126, C127}
    Row08 := [...]rune{C128, C129, C130, C131, C132, C133, C134, C135, C136, C137, C138, C139, C140, C141, C142, C143}
    Row09 := [...]rune{C144, C145, C146, C147, C148, C149, C150, C151, C152, C153, C154, C155, C156, C157, C158, C159}
    Row10 := [...]rune{C160, C161, C162, C163, C164, C165, C166, C167, C168, C169, C170, C171, C172, C173, C174, C175}
    Row11 := [...]rune{C176, C177, C178, C179, C180, C181, C182, C183, C184, C185, C186, C187, C188, C189, C190, C191}
    Row12 := [...]rune{C192, C193, C194, C195, C196, C197, C198, C199, C200, C201, C202, C203, C204, C205, C206, C207}
    Row13 := [...]rune{C208, C209, C210, C211, C212, C213, C214, C215, C216, C217, C218, C219, C220, C221, C222, C223}
    Row14 := [...]rune{C224, C225, C226, C227, C228, C229, C230, C231, C232, C233, C234, C235, C236, C237, C238, C239}
    Row15 := [...]rune{C240, C241, C242, C243, C244, C245, C246, C247, C248, C249, C250, C251, C252, C253, C254, C255}
    
    Matrix := [16][16]rune{Row00, Row01, Row02, Row03, Row04, Row05, Row06, Row07, Row08, Row09, Row10, Row11, Row12, Row13, Row14, Row15}
    return Matrix
}

//======================================================================================================================
//
// VIc - Complete Key Generation Methods - Part III
//
// VIc.0
// The main Function aggregating all Key Generation Methods together
func (k *FiniteFieldEllipticCurve) MakeCryptoplasmKeys() {
    var (
        FirstAnswer        uint
        Answer11, Answer31 uint
        Answer111          uint
        
        Print1 = `1.Create New Cryptoplasm Keys
2.Open Existing Cryptoplasm Keys`
        Print10 = `How do you want to create the Private Key ?`
        Print11 = `1.1.BitString         - create Cryptoplasm Private Key from a BitString
1.2.Number in Base 10 - create Cryptoplasm Private-Key from a Base 10 Number
1.3.Number in Base 49 - create Cryptoplasm Private-Key from a Base 49 Number
1.4.Monochrome Bitmap - create Cryptoplasm Private-Key from a Monochrome Bitmap. Whites are 0, Blacks are 1
1.5.Custom Seed Words - create Cryptoplasm Private-Key from a custom List of Words`
        Print111 = `1.1.1 Random BitString - create a Cryptoplasm Private-Key from a random BitString
1.1.2 Manual BitString - create a Cryptoplasm Private-Key from a manually inputted BitString`
        Print30 = `How do you have the Private Key saved ?`
        Print31 = `1.Key File          - open Cryptoplasm Private-Key from a Key-File
2.Monochrome Bitmap - the Cryptoplasm Private-Key from a Monochrome-Bitmap`
    )
    fmt.Println("What do you want to do ?")
    fmt.Println(Print1)
    _, _ = fmt.Scanln(&FirstAnswer)
    if FirstAnswer == 1 {
        fmt.Println(Print10)
        fmt.Println(Print11)
        _, _ = fmt.Scanln(&Answer11)
        switch Answer11 {
        case 1:
            fmt.Println(Print111)
            _, _ = fmt.Scanln(&Answer111)
            switch Answer111 {
            case 1:
                Keys, Address := k.CPFromRandomBits()
                PrintKeysWithAddress(Keys, Address)
                //Saving Generated Keys
                BitString := k.PrivKey2BitString(Keys.PrivateKey)
                k.SaveBitString(BitString)
            case 2:
                Keys, Address := k.CPFromManualBits()
                PrintKeysWithAddress(Keys, Address)
                //Saving Generated Keys
                BitString := k.PrivKey2BitString(Keys.PrivateKey)
                k.SaveBitString(BitString)
            }
        case 2:
            Keys, Address := k.CPFromNumber(10)
            PrintKeysWithAddress(Keys, Address)
            //Saving Generated Keys
            BitString := k.PrivKey2BitString(Keys.PrivateKey)
            k.SaveBitString(BitString)
        case 3:
            Keys, Address := k.CPFromNumber(49)
            PrintKeysWithAddress(Keys, Address)
            //Saving Generated Keys
            BitString := k.PrivKey2BitString(Keys.PrivateKey)
            k.SaveBitString(BitString)
        case 4:
            fmt.Println("Reading a Private Key from a Monochrome Bitmap is not implemented yet")
        case 5:
            Keys, Address := k.CPFromSeed()
            PrintKeysWithAddress(Keys, Address)
            //Saving Generated Keys
            BitString := k.PrivKey2BitString(Keys.PrivateKey)
            k.SaveBitString(BitString)
        }
    } else if FirstAnswer == 2 {
        fmt.Println(Print30)
        fmt.Println(Print31)
        _, _ = fmt.Scanln(&Answer31)
        switch Answer31 {
        case 1:
            fmt.Println("Case 1")
            BitString, _ := k.OpenKeys()
            Keys, Address := k.BS2CRYPTOPLASM(BitString)
            PrintKeysWithAddress(Keys, Address)
        case 2:
            fmt.Println("Opening a Private Key from a Monochrome Bitmap is not implemented yet")
        }
    }
}

// VIc.1
//
// Creates a Key-Pair and Address from Random Bits
func (k *FiniteFieldEllipticCurve) CPFromRandomBits() (Keys CPKeyPair, Address string) {
    Keys, Address = k.RBS2CRYPTOPLASM()
    return Keys, Address
}

// VIc.2
//
// Creates a Key-Pair and Address from manually inputted Bits
func (k *FiniteFieldEllipticCurve) CPFromManualBits() (Keys CPKeyPair, Address string) {
    var (
        Answer    string
        Condition bool
    )
    for {
        fmt.Println("Please enter a string of bits(0s and 1s)", k.S, "digits long:")
        _, _ = fmt.Scanln(&Answer)
        TT, T1, T2 := k.ValidateBitString(Answer)
        if TT == true {
            Condition = true
        }
        
        if Condition == true {
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
    Keys, Address = k.BS2CRYPTOPLASM(Answer)
    return Keys, Address
}

// VIc.3
//
// Creates a Key-Pair and Address from number in the given Base.
// It is verified if the inputted number in the given base if of the proper clamped Scalar Form.
func (k *FiniteFieldEllipticCurve) CPFromNumber(Base int) (Keys CPKeyPair, Address string) {
    var (
        Answer    string
        AnswerINT = new(big.Int)
        Condition bool
    )
    for {
        fmt.Println("Please enter a compatible Base", Base, "Number to be used as Private-Key:")
        _, _ = fmt.Scanln(&Answer)
        AnswerINT.SetString(Answer, Base)
        BinaryAnswer := ConvertBase10toBase2(AnswerINT)
        TT, T1, T2 := k.ValidateBigBitString(BinaryAnswer)
        if TT == true {
            Condition = true
        }
        if Condition == true {
            break
        } else {
            if T1 == false && T2 == true {
                fmt.Println("The number entered converted to Binary doesnt have the required length, try again!")
            } else if T1 == true && T2 == false {
                fmt.Println("The number entered doesnt have a proper clamped form according to curve Cofactor, try again!")
            } else if T1 == false && T2 == false {
                fmt.Println("The number entered does neither have the required length, nor is it in a clamped form according to curve Cofactor, try again!")
            }
        }
    }
    Keys, Address = Number2CRYPTOPLASM(AnswerINT)
    return Keys, Address
}

// VIc.4
//
// Creates a Key-Pair and Address from a custom number of Seed Words.
func (k *FiniteFieldEllipticCurve) CPFromSeed() (Keys CPKeyPair, Address string) {
    var (
        SeedNumber       uint32
        Condition1       bool
        ConcatenatedSeed string
    )
    for {
        fmt.Println("How many Seed Words do you have? (Enter a number between 1 and 256)")
        _, _ = fmt.Scanln(&SeedNumber)
        if SeedNumber >= 1 && SeedNumber <= 256 {
            Condition1 = true
        }
        if Condition1 == true {
            break
        }
    }
    
    //Getting Seed Words from input.
    var SeedSlice = make([]string, SeedNumber)
    for i := 0; i < int(SeedNumber); i++ {
        fmt.Println("Enter Seed Word number", i+1, ":")
        _, _ = fmt.Scanln(&SeedSlice[i])
    }
    
    //Concatenating Seed Words
    fmt.Println("Seed Slice is:", SeedSlice)
    for i := 0; i < len(SeedSlice); i++ {
        ConcatenatedSeed = ConcatenatedSeed + SeedSlice[i]
    }
    
    BitString := k.StringToBitString(ConcatenatedSeed)
    fmt.Println("BS is", BitString)
    Keys, Address = k.BS2CRYPTOPLASM(BitString)
    return Keys, Address
}

// VIc.5
//
// Creates a BitString of k.S length from another String using a blake3 Hash
func (k *FiniteFieldEllipticCurve) StringToBitString(Word string) string {
    var (
        kS                                         = new(big.Int)
        SByteArray                                 []byte
        Number                                     = new(big.Int)
        OutputSize                                 int
        Type                                       int
        IntermediaryResult2, Result, MissingString string
    )
    //Converting the String to a slice of bytes
    WordToByteSlice := []byte(Word)
    
    //Hashing the slice of bytes with k.S bits output Size
    //Output size is in bytes, therefore k.S should be greater than 8 and a multiple of 8 for this to work properly.
    //These conditions are in case k.S is smaller than 8 or isn't a multiple of 8.
    //Type 1 is when the k.S is directly divisible by 8.
    kS = big.NewInt(int64(k.S))
    T, Rest := IsDivisibleByEight(kS)
    if T == true {
        OutputSize = int(k.S) / 8
        Type = 1
    } else if Rest.Cmp(Zero) != 0 && k.S > 8 {
        OutputSize = int(k.S)/8 + 1
        Type = 2
    } else {
        OutputSize = 1
        Type = 3
    }
    
    S := Blake3.SumCustom(WordToByteSlice, OutputSize)
    
    //Converting the resulting hash which is a slice of bytes, to hex (byte to hex)
    for i := 0; i < len(S); i++ {
        SByteArray = append(SByteArray, S[i])
    }
    Hex := hex.EncodeToString(SByteArray)
    Number.SetString(Hex, 16)
    IntermediaryResult := ConvertBase10toBase2(Number)
    
    //Check if IntermediaryResult is exactly of length OutputSize * 8, If not add zeroes in front of it.
    IR := []rune(IntermediaryResult)
    if len(IR) == OutputSize*8 {
        IntermediaryResult2 = IntermediaryResult
    } else {
        Missing := OutputSize*8 - len(IR)
        for i := 0; i < Missing; i++ {
            MissingString = MissingString + "0"
        }
        IntermediaryResult2 = MissingString + IntermediaryResult
    }
    
    //This code makes sure the resulting BinaryString is equal to k.S in size.
    //In case Safe-Scalar-Size is not divisible with 8, or is smaller than 8.
    if Type == 1 {
        Result = IntermediaryResult2
    } else if Type == 2 || Type == 3 {
        R := []rune(IntermediaryResult2)
        //Amount := uint64(OutputSize) * 8 - k.S
        //fmt.Println("Amount is",Amount)
        R2 := R[:k.S]
        //Converting the Slice of 0s and 1s to a String
        for i := 0; i < len(R2); i++ {
            Result = Result + string(R2[i])
        }
    }
    
    return Result
}

// VIc.6
//
// Saves a BitString representing a Private-Key to an external Cryptoplasm Key File
func (k *FiniteFieldEllipticCurve) SaveBitString(BitString string) () {
    var (
        Identifier       string
        P1, P2, Password string
        Condition        bool
    )
    fmt.Println("")
    fmt.Println("The BitString representing the Private-Key is being saved !")
    fmt.Println("Enter an identifier for the File to be exported:")
    _, _ = fmt.Scanln(&Identifier)
    for {
        fmt.Println("Enter a password to encrypt the Private-Key:")
        _, _ = fmt.Scanln(&P1)
        fmt.Println("Confirm the entered Password by retyping it:")
        _, _ = fmt.Scanln(&P2)
        if P1 == P2 {
            Condition = true
        }
        if Condition == true {
            Password = P1
            break
        } else {
            fmt.Println("Retyped Password doesn't match the previous entered Password!")
        }
    }
    k.ExportPrivateKey(BitString, Identifier, Password)
}

// VIc.7
//
// Exports the BitString representing a Private-Key to a named files and encrypts it with a Password
func (k *FiniteFieldEllipticCurve) ExportPrivateKey(BitString, KeyName, Password string) () {
    Keys := k.GetKeys(BitString)
    
    EncryptedPK := ConvertBase2toBase49(AES.EncryptBitString(BitString, Password))
    PublicKey := Keys.PublicKey
    Address := PublicKey2CRYPTOPLASMAddress(Keys.PublicKey)
    
    String1 := "Cryptoplasm Private-Key in encrypted form:"
    String2 := "Cryptoplasm Public-Key:"
    String3 := "Cryptoplasm Address:"
    String4 := "=========================================="
    
    FileName := KeyName + "_Keys.Cryptoplasm"
    
    OutputFile, err := os.Create(FileName)
    if err != nil {
        log.Fatal(err)
    }
    defer OutputFile.Close()
    
    //Exporting Data
    _, _ = fmt.Fprintln(OutputFile, String1)
    _, _ = fmt.Fprintln(OutputFile, EncryptedPK)
    _, _ = fmt.Fprintln(OutputFile, String4)
    _, _ = fmt.Fprintln(OutputFile, String2)
    _, _ = fmt.Fprintln(OutputFile, PublicKey)
    _, _ = fmt.Fprintln(OutputFile, String4)
    _, _ = fmt.Fprintln(OutputFile, String3)
    _, _ = fmt.Fprintln(OutputFile, Address)
}

// VIc.7
//
// Opens a saved Key File and decrypts it.
func (k *FiniteFieldEllipticCurve) OpenKeys() (string, error) {
    var (
        FileName, Password, MissingString, BitString, BitStringIntermediary string
        E1                                                                  error
    )
    fmt.Println("Cryptoplasm Keys are being opened!")
    
    fmt.Println("Enter the FileName containing the Keys:")
    _, _ = fmt.Scanln(&FileName)
    
    //Reading the Encrypted Private-Key. It is found on Line 2
    EncryptedPK, _ := ImportPrivateKey(FileName)
    
    for {
        fmt.Println("Enter the Password to decrypt the Private-Key:")
        _, _ = fmt.Scanln(&Password)
        
        //Converting the Encrypted-PK to a BitString of 0s and 1s
        EncryptedPKBitString := ConvertBase49toBase2(EncryptedPK)
        
        //Decrypting the BitString using the given Password
        BitStringIntermediary, E1 = AES.DecryptBitString(EncryptedPKBitString, Password)
        
        //Loop Breaking Condition, the Condition is for the password to be correct
        if E1 == nil {
            fmt.Println("Entered Password is correct !")
            break
        } else {
            fmt.Println("Entered Password is incorrect !")
        }
    }
    
    //Checking if the Length of resulting BitString is correct
    BitStringSlice := []rune(BitStringIntermediary)
    if uint64(len(BitStringSlice)) == k.S {
        BitString = BitStringIntermediary
    } else {
        Missing := k.S - uint64(len(BitStringSlice))
        for i := uint64(0); i < Missing; i++ {
            MissingString = MissingString + "0"
        }
        BitString = MissingString + BitStringIntermediary
    }
    
    //fmt.Println("Correct Decrypted BitString is",BitString)
    return BitString, nil
}

func ImportPrivateKey(FileName string) (string, error) {
    PK, err := aux.ReadNthLine(FileName, 2)
    return PK, err
}
