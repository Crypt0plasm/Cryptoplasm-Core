package Cryptoplasm_Elliptic

import (
    b "Cryptoplasm-Core/Cryptoplasm_Blockchain_Constants"
    "fmt"
    blake3 "github.com/Crypt0plasm/Cryptographic-Hash-Functions/Blake3"
    p "github.com/Crypt0plasm/Firefly-APD"
    "math/big"
    "strings"
)
// The Schnorr Signature Structure
type Schnurr struct {
    R AffineCoordinates
    S *big.Int
}
//
// Part I - Converting a Hash to a big.Int
//
func Hash2BigInt (Hash []byte) *big.Int {
    // Internal Function 1
    ByteSlice2HexString := func (Hash []byte) string {
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
    // Internal Function 2
    HexString2BigInt := func(HexString string) *big.Int {
        var result = new(big.Int)
        result.SetString(HexString,16)
        return result
    }

    //Main Body Function
    HexString := ByteSlice2HexString(Hash)
    result := HexString2BigInt(HexString)
    return result
}
//
// Part II - Converting the PublicKey string format into Affine Coordinates
//
func PublicKey2Affine (PublicKey string) AffineCoordinates {
    var (
        Result AffineCoordinates
        PublicKeyDecimal *p.Decimal
    )
    SplitString := strings.Split(PublicKey,".")
    XLengthStr := SplitString[0]
    XLength := ConvertBase49toBase10(XLengthStr)
    XLengthInt64 := XLength.Int64()

    PublicKeyIntStrBase49 := SplitString[1]
    PublicKeyInt := ConvertBase49toBase10(PublicKeyIntStrBase49)
    PublicKeyIntStrBase10 := PublicKeyInt.Text(10)

    PublicKeyDecimal = p.NFS(PublicKeyIntStrBase10)
    Digits := b.Count4Coma(PublicKeyDecimal)

    YLength := Digits - XLengthInt64
    Divizor := b.POWxc(p.NFI(10),p.NFI(YLength))

    PublicKeyDecimalDivided := b.DIVxc(PublicKeyDecimal,Divizor)
    X := b.RemoveDecimals(PublicKeyDecimalDivided)
    Subtractor := b.MULxc(X,Divizor)
    Y := b.SUBxc(PublicKeyDecimal,Subtractor)

    Result.AX = &X.Coeff
    Result.AY = &Y.Coeff

    return Result
}
// Part III - Schnorr Signature Functions:
// Schnorr Signature Creation:
//
//          Private key :             k            (integer)
//          Public key :              P = k*G      (curve point)
//          Message hash:             m            (integer)
//
//          Generate a random number: z            (integer)
//          Calculate:                R = z*G      (curve point)
//
//          Calculate: s = z + Hash(r||P||m)*k     (integer)
//
//          where: r = X-coordinate of curve point R
//          and || denotes binary concatenation
//          Signature = (R, s)     (curve point, integer)
//==================================================================
// Schnorr Signature Verification
//
//          obtain the signature: (R,s)
//          Obtain public key : P
//          Obtain message: m
//
//          Verify: s*G =   R + Hash(r||P||m) *   P
//                  s*G = z*G + Hash(r||P||m) * k*G
//
// Function that creates the big.Int from the Hashed Schnorr Parameters
//
// VII.1
func (k *FiniteFieldEllipticCurve) SchnorrHash (r *big.Int, PublicKey string, Message []byte) *big.Int {
    SmallRBinary := r.Text(2)

    PKAffine := PublicKey2Affine(PublicKey)
    PublicKeyXBinary := PKAffine.AX.Text(2)
    PublicKeyYBinary := PKAffine.AY.Text(2)

    MessageBinary := Hash2BigInt(Message).Text(2)

    ConcatenatedBinaryString := SmallRBinary + PublicKeyXBinary + PublicKeyYBinary + MessageBinary
    ConcatenatedBinaryStringToByteSlice := []byte(ConcatenatedBinaryString)

    //OutputSize in Bytes must not necessarily be as big as the bit-size of the Prime Number of the Curve.
    //But it is settled on that size.
    OutputSizeInBytes := int(k.S) / 8
    Hash := blake3.SumCustom(ConcatenatedBinaryStringToByteSlice,OutputSizeInBytes)
    SchnorrHashInt := Hash2BigInt(Hash)
    return SchnorrHashInt
}
//
// The Schnorr Sign Function. It takes a Pair of CryptoPlasm Keys (Strings in 49Base Format, and special Construction for the Public.Key)
// and a Hashed Output as the Message to be singed for, and provides a Signature.
//
// VII.2
func (k *FiniteFieldEllipticCurve) SchnorrSign (Keys CPKeyPair, Hash []byte) (Signature Schnurr) {
    var (
        z = new(big.Int)
        v1 = new(big.Int)
        s = new(big.Int)
    )
    //Step 1, choosing random number kappa within prime field interval
    BitString := k.GetRandomBitsOnCurve()
    z = k.ClampBitString(BitString)

    //Step 2, computing R (x,y) as z multiplied by Curve Generator Point
    RExt  := k.ScalarMultiplierG(z)
    RAff := k.Extended2Affine(RExt)

    SchnorredHash := k.SchnorrHash(RAff.AX,Keys.PublicKey,Hash)

    //If MulMod and AddMod is used, signature doesnt verify.
    //Numbers must be allowed to grow freely. Restricting them modulo Q or P breaks the signature.
    //SchnorredHash size in similar to Curve S size in bytes. Multiplied by the private key which is of similar size,
    //The resulting s is roughly two times the size of the private Key is bits
    v1.Mul(SchnorredHash,ConvertBase49toBase10(Keys.PrivateKey))
    s.Add(z,v1)
    fmt.Println("s is",s)
    //v1 = k.MulModQ(SchnorredHash,ConvertBase49toBase10(Keys.PrivateKey))
    //s  = k.AddModQ(z,v1)

    Signature.R = RAff
    Signature.S = s
    return Signature
}
// The Schnorr Verify Function. It takes a Signature, a CryptoPlasm PublicKey (in CryptoPlasm PublicKey Format),
// and a Hashed Output as the Message, and returns true is the Signature is valid for the Public Key.
//
// VII.3
func (k *FiniteFieldEllipticCurve) SchnorrVerify (Sigma Schnurr, PublicKey string, Hash []byte) bool {
    var Result bool
    //GeneratorAff := AffineCoordinates {&k.PBX, &k.PBY}
    //GeneratorExt := k.Affine2Extended(GeneratorAff)
    sGExt := k.ScalarMultiplierG(Sigma.S)
    sGAff := k.Extended2Affine(sGExt)

    SchnorredHash := k.SchnorrHash(Sigma.R.AX,PublicKey,Hash)
    PublicKeyGeneratorAff := PublicKey2Affine(PublicKey)
    PublicKeyGeneratorExt := k.Affine2Extended(PublicKeyGeneratorAff)

    Q2ex  := k.ScalarMultiplierPt(SchnorredHash,PublicKeyGeneratorExt)

    Q1ex:= k.Affine2Extended(Sigma.R)

    Q := k.AdditionV2(Q1ex,Q2ex)
    QAff := k.Extended2Affine(Q)

    Result = k.ArePointsEqualAf(sGAff,QAff)
    return Result
}