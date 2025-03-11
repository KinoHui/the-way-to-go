package eth

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"

	"golang.org/x/crypto/sha3"
)

func Nounce() (uint64, error) {
	nonce, err := client.PendingNonceAt(context.Background(), address)
	return nonce, err
}

func TransferTokens() {
	signedTx, err := DefaultTxn()
	if err != nil {
		log.Fatal(err)
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s", signedTx.Hash().Hex())
}

func DefaultTxn() (*types.Transaction, error) {
	value := big.NewInt(0)
	nonce, _ := client.PendingNonceAt(context.Background(), address)

	transferFnSig := []byte("transfer(address,uint256)")
	hash := sha3.NewLegacyKeccak256()
	hash.Write(transferFnSig)
	methodID := hash.Sum(nil)[:4]
	fmt.Println("methodID: ", hexutil.Encode(methodID))

	paddedToAddress := common.LeftPadBytes(toAddress.Bytes(), 32)

	amount := new(big.Int)
	amount.SetString("10000000000000000000", 10)
	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)

	var data []byte
	data = append(data, methodID...)
	data = append(data, paddedToAddress...)
	data = append(data, paddedAmount...)

	gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
		To:   &toAddress,
		Data: data,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("gasLimit", gasLimit) // 0x5edfe1c7e923650163cbf8ffe3652ff12e1e915f7db2e254253d253a5da30be6
	gasLimit = uint64(70041)

	tx := types.NewTransaction(nonce, tokenAddress, value, gasLimit, gasPrice, data)

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), PrivateKey())
	if err != nil {
		return nil, err
	}
	return signedTx, nil
}
