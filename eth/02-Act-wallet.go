package eth

import (
	"crypto/ecdsa"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func Wallet() {

	privateKeyECDSA, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("privateKeyECDSA: ", privateKeyECDSA)

	privateKeyBytes := crypto.FromECDSA(privateKeyECDSA)
	privateKey := hexutil.Encode(privateKeyBytes)
	fmt.Println("privateKey(string): ", privateKey)

	// 公钥
	publicKeyCrypto := privateKeyECDSA.Public()
	publicKeyECDSA, ok := publicKeyCrypto.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	publicKey := hexutil.Encode(publicKeyBytes)
	fmt.Println("publicKey(string): ", publicKey)

	// address
	// PubKey ---Keccak-256-Hash---> Address
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	fmt.Println("address: ", address)

	// 手动
	// golang.org/x/crypto/sha3
	// hash := sha3.NewLegacyKeccak256()
	// hash.Write(publicKeyBytes[1:])
	// fmt.Println(hexutil.Encode(hash.Sum(nil)[12:]))
}
