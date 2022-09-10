package common

import (
	"log"
	"npm3/cli/generated/packagemanager"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func GetPackageManagerInstance(client *ethclient.Client) *packagemanager.Packagemanager {
	blocklistAddressHex := "0x081c78BDcFa10391862E44215543d2B8F3211c92"
	addr := common.HexToAddress(blocklistAddressHex)
	var err error
	instance, err := packagemanager.NewPackagemanager(addr, client)
	if err != nil {
		log.Fatalf("failed to create packagemanager instance at address %v, error: %v", blocklistAddressHex, err.Error())
	}
	return instance
}
