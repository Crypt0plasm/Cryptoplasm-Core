package Cryptoplasm_Elliptic

import (
    blake3 "github.com/Crypt0plasm/Cryptographic-Hash-Functions/Blake3"
    "math/big"
)
//
//  Functions to convert a Hash to a big.Int
//
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

// Schnorr Signature Functions

type Schnurr struct {
    R AffineCoordinates
    S *big.Int
}

//Schnorr Signature Creation:
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
//Schnorr Signature Verification
//
//          obtain the signature: (R,s)
//          Obtain public key : P
//          Obtain message: m
//
//          Verify: s*G =   R + Hash(r||P||m) *   P
//                  s*G = z*G + Hash(r||P||m) * k*G


// Function that creates the big.Int from the Hashed Schnorr Parameters
func SchnorrHash (r *big.Int, PublicKey string, Message []byte) *big.Int {
    SmallRBinary := r.Text(2)

    PKAffine := PublicKey2Affine(PublicKey)
    PublicKeyXBinary := PKAffine.AX.Text(2)
    PublicKeyYBinary := PKAffine.AY.Text(2)

    MessageBinary := Hash2BigInt(Message).Text(2)

    ConcatenatedBinaryString := SmallRBinary + PublicKeyXBinary + PublicKeyYBinary + MessageBinary
    ConcatenatedBinaryStringToByteSlice := []byte(ConcatenatedBinaryString)

    //OutputSize in Bytes must not necessarily be as big as the bit-size of the Prime Number of the Curve.
    //But it is settled on that size.
    Hash := blake3.SumCustom(ConcatenatedBinaryStringToByteSlice,65)
    SchnorrHashInt := Hash2BigInt(Hash)
    return SchnorrHashInt
}
//
// The Schnorr Sign Function. It takes a Pair of CryptoPlasm Keys (Strings in 49Base Format, and special Construction for the Public.Key)
// and a Hashed Output as the Message to be singed for, and provides a Signature.
//
func (k *FiniteFieldEllipticCurve) SchnorrSign (Keys CPKeyPair, Hash []byte) (Signature Schnurr) {
    var (
    	z = new(big.Int)
    	v1 = new(big.Int)
        s = new(big.Int)
    )
    //Step 1, choosing random number kappa within prime field interval
    z = k.GetRandomOnCurve()

    //Step 2, computing Q (x,y) as z multiplied by Curve Generator Point
    Generator := AffineCoordinates {&k.PBX, &k.PBY}
    R  := k.ScalarMultiplier(z,Generator)

    SchnorredHash := SchnorrHash(R.AX,Keys.PublicKey,Hash)

    //If MulMod and AddMod is used, signature doesnt verify.
    //Numbers must be allowed to grow non modulo P
    // For k*G=P(k private, P public), H*k*G != H*P if H*k wraps around modulo prime.
    v1.Mul(SchnorredHash,ConvertBase49toBase10(Keys.PrivateKey))
    s.Add(z,v1)

    Signature.R = R
    Signature.S = s
    return Signature
}

// The Schnorr Verify Function. It takes a Signature, a CryptoPlasm PublicKey (in CryptoPlasm PublicKey Format),
// and a Hashed Output as the Message, and returns true is the Signature is valid for the Public Key.

func (k *FiniteFieldEllipticCurve) SchnorrVerify (Sigma Schnurr, PublicKey string, Hash []byte) bool {
    var Result bool
    Generator := AffineCoordinates {&k.PBX, &k.PBY}
    sG := k.ScalarMultiplier(Sigma.S,Generator)

    SchnorredHash := SchnorrHash(Sigma.R.AX,PublicKey,Hash)
    PublicKeyGenerator := PublicKey2Affine(PublicKey)

    Q2  := k.ScalarMultiplier(SchnorredHash,PublicKeyGenerator)
    Q2ex:= k.Affine2Extended(Q2)

    Q1ex:= k.Affine2Extended(Sigma.R)

    Q := k.AdditionV2(Q1ex,Q2ex)
    QAff := k.Extended2Affine(Q)

    Result = k.ArePointsEqualAf(sG,QAff)
    return Result
}