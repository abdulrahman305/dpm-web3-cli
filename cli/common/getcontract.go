package common

import (
	"dpm3/cli/generated/packagemanager"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func GetPackageManagerInstance(client *ethclient.Client) *packagemanager.Packagemanager {
	blocklistAddressHex := "0x8c6C2BdAf20368a794BcB2A7E0953EAa56Cf2a9F"
	addr := common.HexToAddress(blocklistAddressHex)
	var err error
	instance, err := packagemanager.NewPackagemanager(addr, client)
	if err != nil {
		log.Fatalf("failed to create packagemanager instance at address %v, error: %v", blocklistAddressHex, err.Error())
	}
	return instance
}
