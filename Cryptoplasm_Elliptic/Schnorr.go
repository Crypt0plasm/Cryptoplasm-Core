package Cryptoplasm_Elliptic

import (
    "fmt"
    blake3 "github.com/Crypt0plasm/Cryptographic-Hash-Functions/Blake3"
    "math/big"
)

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

func Hash2BigInt (Hash []byte) *big.Int {
    HexString := ByteSlice2HexString(Hash)
    result := HexString2BigInt(HexString)
    return result
}

type Schnurr struct {
    R *big.Int
    S *big.Int
}

//This Schnorr Signature implementation is used for CryptoPlasm Transaction signing.
//Therefore the "Keys" used as input are generated with the same CurveParameters "k"
//
//Same Function is used as described in Zilliqa Whitepaper Appendix A
func (k *FiniteFieldEllipticCurve) CPSchnorrSign (Keys CPKeyPair, Hash []byte) (Signature Schnurr) {
    //Keys are the CryptoPlasm Key pairs used to generate the Signature
    //Hash is the Blake3 Hash of the Message for which the Signature is generated.

    var (
    	r = new(big.Int)
        s = new(big.Int)
    	kappa = new(big.Int)
    	v1 = new(big.Int)
        v2 = new(big.Int)
    )
    Cmp1 := r.Cmp(Zero)
    Cmp2 := s.Cmp(Zero)

    for Cmp2 == 0 {
        for Cmp1 == 0 {
            //Step 1, choosing random number kappa within prime field interval
            kappa = k.GetRandomOnCurve()

            //Step 2, computing Q (x,y) as kappa multiplied by Curve Generator Point
            Generator := AffineCoordinates {&k.PBX, &k.PBY}
            Q  := k.ScalarMultiplier(kappa,Generator)

            //Step 3a, Concatenating Q,pk,m
            String1 := Q.AX.String()
            String2 := Q.AY.String()

            PKAffine := PublicKey2Affine(Keys.PublicKey)
            String3 := PKAffine.AX.String()
            String4 := PKAffine.AY.String()

            String5 := Hash2BigInt(Hash).String()
            BigString := String1 + String2 + String3 + String4 + String5

            //Step 3b, Hashing the Concatenated Result with Blake3
            BigStringToByteSlice := []byte(BigString)
            BigStringHash := blake3.SumCustom(BigStringToByteSlice,160)

            //Step 3c, Converting the resulted Hash in big.Int and computing mod n
            r.Mod(Hash2BigInt(BigStringHash),&k.P)

            //Step4, Checking if r is zero, if it is, start again from 1
            Cmp1 = Zero.Cmp(r)
        }
        //Step 5, Computing s
        v1.Mul(r,ConvertBase49toBase10(Keys.PrivateKey))
        v2.Sub(kappa,v1)
        s.Mod(v2,&k.P)

        //Step6, Checking if s is zero, if it is, start again from 1
        Cmp2 = Zero.Cmp(s)
    }

    //Step7, Finalizing signature value
    Signature.R = r
    Signature.S = s

    return Signature
}

func (k *FiniteFieldEllipticCurve) CPSchnorrVerify (Sigma Schnurr, PublicKey string, Hash []byte) bool {
    var (
    	Verification bool
        v = new(big.Int)
    )
    //Step 3a, Computing s*G
    Generator := AffineCoordinates {&k.PBX, &k.PBY}
    Q1  := k.ScalarMultiplier(Sigma.S,Generator)
    Q1ex:= k.Affine2Extended(Q1)

    //Step 3b, Computing r*pk
    PublicKeyGenerator := PublicKey2Affine(PublicKey)
    Q2  := k.ScalarMultiplier(Sigma.R,PublicKeyGenerator)
    Q2ex:= k.Affine2Extended(Q2)

    //Step 3c, Computing Q
    Q   := k.AdditionV2(Q1ex,Q2ex)
    Qext:= k.Extended2Affine(Q)

    //Step 4a, Concatenating Q,pk,m
    String1 := Qext.AX.String()
    String2 := Qext.AY.String()

    PKAffine := PublicKey2Affine(PublicKey)
    String3 := PKAffine.AX.String()
    String4 := PKAffine.AY.String()

    String5 := Hash2BigInt(Hash).String()
    BigString := String1 + String2 + String3 + String4 + String5

    //Step 4b, Hashing the Concatenated Result with Blake3
    BigStringToByteSlice := []byte(BigString)
    BigStringHash := blake3.SumCustom(BigStringToByteSlice,160)

    //Step 4c, Converting the resulted Hash in big.Int and computing mod n
    v.Mod(Hash2BigInt(BigStringHash),&k.P)

    fmt.Println("v mic este", v)

    return Verification
}