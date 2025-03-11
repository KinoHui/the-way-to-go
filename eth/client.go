package eth

import (
	"fmt"

	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	rpcURL = "https://mainnet.infura.io/v3/7dcf2069eb6b4c4f851e0b1a762a94b1"
)

func NewDefaultClient(net string, option ...int) (*ethclient.Client, error) {
	if len(option) == 0 {
		if net == "" {
			net = "mainnet"
		}
		rpcURL := fmt.Sprintf("https://%s.infura.io/v3/7dcf2069eb6b4c4f851e0b1a762a94b1", net)
		client, err := ethclient.Dial(rpcURL)
		if err != nil {
			return nil, err
		}
		return client, nil
	} else {
		if net == "" {
			net = "mainnet"
		}
		rpcURL := fmt.Sprintf("wss://%s.infura.io/ws/v3/7dcf2069eb6b4c4f851e0b1a762a94b1", net)
		client, err := ethclient.Dial(rpcURL)
		if err != nil {
			return nil, err
		}
		return client, nil
	}

}
