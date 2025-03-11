package eth

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

func PrivateKey() *ecdsa.PrivateKey {
	privateKey, _ := crypto.HexToECDSA("be38d65e72a295c7a6b3cf29018494eca485f628d77b6c3dcb6991f0653abbea")
	return privateKey
}

func Transfer() {
	// client, _ := NewDefaultClient()

	privateKey, err := crypto.HexToECDSA("be38d65e72a295c7a6b3cf29018494eca485f628d77b6c3dcb6991f0653abbea")
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	value := big.NewInt(1000000000000000) // in wei (1 eth)

	// toAddress := common.HexToAddress("0xb473D05ac3c1690c27F21c3eDFb6F764885B26D2")
	var data []byte
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal("SignTx", err)
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal("SendTransaction(): ", err)
	}

	fmt.Printf("tx sent: %s", signedTx.Hash().Hex())
}
