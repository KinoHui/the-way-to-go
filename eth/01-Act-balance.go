package eth

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

var (
	client, _    = NewDefaultClient("sepolia")
	gasLimit     = uint64(21000) // in units
	gasPrice, _  = client.SuggestGasPrice(context.Background())
	toAddress    = common.HexToAddress("0xb473D05ac3c1690c27F21c3eDFb6F764885B26D2")
	tokenAddress = common.HexToAddress("0x5E012F934C386F2d105d019ff245623f0D5c61DA")
	address      = common.HexToAddress("0x897527501d5f191080060F435F6aD6B36B8a66c2")
)

func Balance() {

	fmt.Println("address: ", address)
	fmt.Println("address: ", address.Hex())

	client, err := NewDefaultClient("")
	if err != nil {
		log.Fatal(err)
	}

	balance, err := client.BalanceAt(context.Background(), address, nil)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("balance(wei): ", balance)

	// wei -> ETH
	fbalance := new(big.Float)
	fbalance.SetString(balance.String())
	feth := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
	fmt.Println("feth: ", feth)

	// pending balance
	pendingBalance, err := client.PendingBalanceAt(context.Background(), address)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("pending balance: ", pendingBalance)
}
