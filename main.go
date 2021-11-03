package main

import (
	"goExplore/common"
	localConfig "goExplore/config"
	"goExplore/controller"
)

func main() {
	// start a new chain explore service
 	common.InitChainExploreService(localConfig.OrgGoConfig, localConfig.OrgGo, localConfig.Admin, localConfig.User)
	// start the web controller
	// common.WalletTest()
	// time.Sleep(100000)
	// common.InvokeInfoByChaincode("dsa312dsd3")
	// common.QueryInfoByChaincode("a9465920-70c0-4317-b636-7f9baece61d2")
	controller.StartIris()



}
