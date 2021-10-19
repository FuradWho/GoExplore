package main

import (
	"goExplore/common"
	localConfig "goExplore/config"
	"goExplore/controller"
)

func main() {
	orgGoClient := common.InitChainExploreService(localConfig.OrgGoConfig, localConfig.Org1, localConfig.Admin, localConfig.User)

	controller.StartIris()

	// Close: 关闭并释放有SDK维护的缓存和连接
	defer orgGoClient.SDK.Close()
}
