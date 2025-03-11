package eth

import (
	"encoding/hex"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/rlp"
)

func RawTransication() {
	signedTx, err := DefaultTxn()
	if err != nil {
		log.Fatal(err)
	}

	// 使用 rlp.Encode 获取交易的 RLP 编码
	rawTxBytes, err := rlp.EncodeToBytes(signedTx)
	if err != nil {
		log.Fatal(err)
	}

	rawTxHex := hex.EncodeToString(rawTxBytes)
	fmt.Println(rawTxHex) // 输出 RLP 编码的十六进制字符串
}
