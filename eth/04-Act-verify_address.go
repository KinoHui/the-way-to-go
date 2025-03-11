package eth

import (
	"context"
	"fmt"
	"log"
	"regexp"

	"github.com/ethereum/go-ethereum/common"
)

func CheckAddress() {
	addresses := []string{"0x897527501d5f191080060F435F6aD6B36B8a66c2", "0xe41d2489571d322189246dafa5ebde1f4699f498"}
	client, err := NewDefaultClient("")
	if err != nil {
		log.Fatal(err)
	}

	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	fmt.Printf("is valid: %v\n", re.MatchString(addresses[0]))
	fmt.Printf("is valid: %v\n", re.MatchString(addresses[1]))

	for i := 0; i < len(addresses); i++ {
		address := common.HexToAddress(addresses[i])
		//获取对应地址字节码
		bytecode, err := client.CodeAt(context.Background(), address, nil)
		if err != nil {
			log.Fatal(err)
		}

		// 地址上无字节码时，该地址为合约地址
		isContract := len(bytecode) > 0
		fmt.Printf("address %s isContract: %v\n", addresses[i], isContract)
	}

}
