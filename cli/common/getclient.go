package common

import (
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

var client *ethclient.Client

func GetClient() *ethclient.Client {
	if client != nil {
		return client
	}

	rpcUrl := "https://matic-mumbai.chainstacklabs.com"
	client, err := ethclient.Dial(rpcUrl)
	if err != nil {
		log.Fatalf("failed to dial rpc url %v", rpcUrl)
	}
	return client
}
