package eth

import (
	"context"
	"fmt"
	"log"
)

func Blocks() {
	// blockNumber := big.NewInt(5671744)
	client, _ := NewDefaultClient("sepolia")

	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(header.Number.String())

	block, err := client.BlockByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(block.Number().Uint64())
	fmt.Println(block.Time())
	fmt.Println(block.Difficulty().Uint64())
	fmt.Println(block.Hash().Hex())
	fmt.Println(len(block.Transactions()))

	count, err := client.TransactionCount(context.Background(), block.Hash())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(count) // 144

}
