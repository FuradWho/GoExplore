package main

import (
	"goExplore/common"
	localConfig "goExplore/config"
	"goExplore/controller"
)

func main() {
	// start a new chain explore service
	orgGoClient := common.InitChainExploreService(localConfig.OrgGoConfig, localConfig.Org1, localConfig.Admin, localConfig.User)

	// start the web controller
	controller.StartIris()

	defer orgGoClient.SDK.Close()
}
