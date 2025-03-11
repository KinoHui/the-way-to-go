package eth

import (
	"context"
	"fmt"
	"log"
)

func Transaction() {
	client, _ := NewDefaultClient("")

	block, err := client.BlockByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	transactions := block.Transactions()

	for i := range 10 {
		tx := transactions[i]
		fmt.Println("=== Transaction Details ===")
		fmt.Printf("Hash: %s\n", tx.Hash().Hex())            // 交易哈希
		fmt.Printf("Value: %s\n", tx.Value().String())       // 交易金额
		fmt.Printf("Gas: %d\n", tx.Gas())                    // Gas 限制
		fmt.Printf("GasPrice: %d\n", tx.GasPrice().Uint64()) // Gas 价格
		fmt.Printf("Nonce: %d\n", tx.Nonce())                // Nonce
		fmt.Printf("Data: %x\n", tx.Data())                  // 交易数据
		fmt.Printf("To: %s\n", tx.To().Hex())                // 接收地址

		// 获取交易收据
		receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("=== Receipt Details ===")
		fmt.Printf("Status: %d\n", receipt.Status) // 交易状态
		fmt.Printf("Logs: %v\n", receipt.Logs)     // 交易日志

		fmt.Println("========================================================")
	}

}
