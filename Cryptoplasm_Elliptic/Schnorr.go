package Cryptoplasm_Elliptic

import "math/big"

func ByteSlice2HexString (Hash []byte) string {
    var (
    	result string
    	HashElement byte
    	HashElementBig = new(big.Int)
    )

    for i := 0; i < len(Hash); i++ {
        HashElement = Hash[i]
        HashElementInt64 := int64(HashElement)
        HashElementBig.SetInt64(HashElementInt64)
        HashElementHex := HashElementBig.Text(16)

        //Converts "c" to "0c" as per correct Blake3Output in Hex
        switch HashElementHex {
        case "a":HashElementHex = "0" + HashElementHex
        case "b":HashElementHex = "0" + HashElementHex
        case "c":HashElementHex = "0" + HashElementHex
        case "d":HashElementHex = "0" + HashElementHex
        case "e":HashElementHex = "0" + HashElementHex
        case "f":HashElementHex = "0" + HashElementHex
        }
        result = result + HashElementHex
    }
    return result
}

func HexString2BigInt (HexString string) *big.Int {
    var result = new(big.Int)
    result.SetString(HexString,16)
    return result
}