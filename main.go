package main

import (
	"goExplore/common"
	localConfig "goExplore/config"
	"goExplore/controller"
)

func main() {
	// start a new chain explore service
 	common.InitChainExploreService(localConfig.OrgGoConfig, localConfig.OrgGo, localConfig.Admin, localConfig.User)
	//start the web controller
	// common.WalletTest()
	// time.Sleep(100000)
	// common.InvokeInfoByChaincode("dsa312dsd3")
	// common.QueryInfoByChaincode("a9465920-70c0-4317-b636-7f9baece61d2")
	controller.StartIris()

	//sdk, err := fabsdk.New(config.FromFile("connect-config/org2-config.yaml"))
	//if err != nil {
	//	fmt.Println(err)
	//}
	//ctx := sdk.Context()
	//client, err := msp.New(ctx)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//req := &msp.RegistrationRequest{
	//	Name: "User1",
	//	Type: "client",
	//	CAName: "rca-org2",
	//	Secret: "123456",
	//}
	//
	//register, err := client.Register(req)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//fmt.Println(register)


}
