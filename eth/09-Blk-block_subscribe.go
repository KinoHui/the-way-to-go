package eth

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
)

func SubcribeBlock() {
	clientWss, err := NewDefaultClient("sepolia", 1)
	if err != nil {
		log.Fatal("Failed to create client: ", err)
	}

	chan_headers := make(chan *types.Header)

	sub, err := clientWss.SubscribeNewHead(context.Background(), chan_headers)
	if err != nil {
		log.Fatal("Failed to subscribe to new headers: ", err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal("Subscription error: ", err)
		case header := <-chan_headers:
			fmt.Println("New block header hash:", header.Hash().Hex())

			// 增加延迟，等待区块同步
			time.Sleep(15 * time.Second)

			// 重试机制
			var block *types.Block
			var err error
			for retry := 0; retry < 3; retry++ {
				block, err = clientWss.BlockByHash(context.Background(), header.Hash())
				if err == nil {
					break
				}
				log.Printf("BlockByHash() attempt %d failed: %v\n", retry+1, err)
				time.Sleep(5 * time.Second) // 每次重试前等待 5 秒
			}
			if err != nil {
				log.Fatal("BlockByHash() failed after retries: ", err)
			}

			// 打印区块信息
			fmt.Println("Block hash:", block.Hash().Hex())
			fmt.Println("Block number:", block.Number().Uint64())
			fmt.Println("Block time:", block.Time())
			fmt.Println("Block nonce:", block.Nonce())
			fmt.Println("Transactions count:", len(block.Transactions()))
			fmt.Println("========================================================")
		}
	}
}
