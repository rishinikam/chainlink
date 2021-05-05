package main

import (
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/smartcontractkit/chainlink/core/services/signatures/secp256k1"
	"github.com/smartcontractkit/chainlink/core/services/vrf"
	"github.com/smartcontractkit/chainlink/core/utils"
)

func pErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	a := "e5e384bd6dbeb342c66e484fc10ca3e1c97882928dd657e5a7a04b49c852e9710b59dbd14bdcad8ff68ceec6686173741ad4a5000d8c2ea79229382fc3c0859155764a5926f9c9517d06cf35f4ac502267b3db13e7983a8ee6af7364d76c9aad71dd6737d5a5db50d6d338051c4a21c86f37f1d5de8158deb48accde72e762fc252d3189544c49777458ce579f56efeb3036065f55f2502f3a1b2aee39b56a8384063c2ea0f05810f39534f6f86930387bf9703ac1d87c9638464ea48563358c5a8d9906ec47cb9a897982027e005fd1f299c0cc17b996ca8506d06b2d10dc7a00000000000000000000000001be75936459151bed16a226707ac42e0a2cde3e408689d1305e40f6f3da54d3fa393d0eb846fd50d31971308a0fd353458448a1ade3bf4f327df3966b80b9ced98dc5c2142e0e16d4a87c9fc30128983f8835176c0bbb8c6d644955a1c68976a5d8ee3744931864e4a55b4ee5f649940b9eced834aa7387074b5d680327fb1f62f131283bf51eb3e4c06f3d1ca3a8a5a5730748cfbd361077ac252490eda8822d75041efdf71b74082a5610236e5a311d20cff70000000000000000000000000000000000000000000000000000000000d6d354"
	b, err := hex.DecodeString(a)
	pErr(err)
	var r [vrf.ProofLength + 32]byte
	copy(r[:], b)
	fmt.Println("proof", r)
	p, err := vrf.UnmarshalProofResponse(r)
	pErr(err)
	//return keccak256(abi.encodePacked(_keyHash, _vRFInputSeed));
	//keyHash = return keccak256(abi.encodePacked(_publicKey));
	Uint256_2, _ := abi.NewType("uint256[2]", "", nil)
	Bytes32, _ := abi.NewType("bytes32", "", nil)
	Uint256, _ := abi.NewType("uint256", "", nil)
	var args = abi.Arguments{{Type: Uint256_2}}
	if !secp256k1.IsSecp256k1Point(p.P.PublicKey) {
		panic("invalid")
	}
	x, y := secp256k1.Coordinates(p.P.PublicKey)
	var coords = [2]*big.Int{x, y}
	keyEncoded, err := args.PackValues([]interface{}{coords})
	pErr(err)
	keyHash, err := utils.Keccak256(keyEncoded)
	pErr(err)
	var keyHash32 [32]byte
	copy(keyHash32[:], keyHash)
	fmt.Println("keyHash", hex.EncodeToString(keyHash))
	var args2 = abi.Arguments{{Type: Bytes32}, {Type: Uint256}}
	reqBytes, err := args2.PackValues([]interface{}{keyHash32, p.PreSeed.Big()})
	pErr(err)
	reqID, err := utils.Keccak256(reqBytes)
	pErr(err)
	fmt.Printf("%+v\n", hex.EncodeToString(reqID))
}
