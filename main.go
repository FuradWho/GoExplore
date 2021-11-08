package main

import (
	"goExplore/fabric-ca"
	"io/ioutil"
)

func main() {
	// start a new chain explore service
 	// common.InitChainExploreService(localConfig.OrgGoConfig, localConfig.OrgGo, localConfig.Admin, localConfig.User)
	//start the web controller
	// common.WalletTest()
	// time.Sleep(100000)
	// common.InvokeInfoByChaincode("dsa312dsd3")
	// common.QueryInfoByChaincode("a9465920-70c0-4317-b636-7f9baece61d2")
	// controller.StartIris()

	//fabric_ca.InitCaClient()

	// ordererDomain  := "orderer.example.com"
	orgs :=[]string{"org2"}
	channelId := "mychannel"
	connectConfig,_ := ioutil.ReadFile("./connect-config/client-network.yaml")
	channelTx := "/usr/local/hyper/fabric-ca/configtx/channel-artifacts/mychannel.tx"
	// chaincodeId := "mycc"
	// chaincodePath := "/usr/local/hyper/fabric-ca/chaincode/newchaincode"

	/*操作fabric start*/
	fabric := fabric_ca.NewFabricClient(connectConfig, channelId ,orgs)
	defer fabric.Close()
	fabric.Setup()
	//创建channel
	fabric.CreateChannel(channelTx)
	//加入channel
	fabric.JoinChannel()


}
