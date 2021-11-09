package main

import (
	"fmt"
	mspclient "github.com/hyperledger/fabric-sdk-go/pkg/client/msp"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/resmgmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/msp"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	"log"
	"strings"
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

	/*
	// ordererDomain  := "orderer.example.com"
	orgs :=[]string{"org2"}
	channelId := "mychannel"
	connectConfig,_ := ioutil.ReadFile("./connect-config/client-network.yaml")

	// chaincodeId := "mycc"
	// chaincodePath := "/usr/local/hyper/fabric-ca/chaincode/newchaincode"

	fabric := fabric_ca.NewFabricClient(connectConfig, channelId ,orgs)
	defer fabric.Close()
	fabric.Setup()
	//创建channel
	//fabric.CreateChannel(channelTx)
	//加入channel
	//fabric.JoinChannel()
	 */


	sdkClient , err := fabsdk.New(config.FromFile("connect-config/channel-connection.yaml"))
	if err != nil {
		log.Panicf("Failed to create a sdkClient :%s \n",err)

	}
	resourceProvider := sdkClient.Context(fabsdk.WithUser("User1"),fabsdk.WithOrg("org2"))

	resourceClient , err := resmgmt.New(resourceProvider)
	if err != nil {
		log.Panicf("Failed to create a resourceClient : %s \n",err)
	}


	mspClient , err := mspclient.New(sdkClient.Context(),mspclient.WithOrg("org2"))
	if err != nil {
		log.Printf("Failed to new mspClient : %s \n",err)
	}

	adminidentity, err := mspClient.GetSigningIdentity("User2")
	if err != nil {
		log.Printf("Failed to get signIdentity : %s \n",err)
	}

	channelTx := "/usr/local/hyper/test2/configtx/channel-artifacts/mychannel.tx"
	channelId := "mychannel"

	req := resmgmt.SaveChannelRequest{
		ChannelID: channelId,
		ChannelConfigPath: channelTx,
		SigningIdentities: []msp.SigningIdentity{adminidentity},
	}

	txId , err := resourceClient.SaveChannel(req)
	if err != nil {
		log.Printf("Failed to save channel : %s \n",err)
	}

	fmt.Println(txId)

	err = resourceClient.JoinChannel(channelId,resmgmt.WithTargetEndpoints("peer1-org2"))
	if err != nil && !strings.Contains(err.Error(), "LedgerID already exists") {
		log.Printf("Org peers failed to JoinChannel: %s \n", err)
	}




 	//sdk, err := fabsdk.New(config.FromFile("connect-config/channel-connection.yaml"))
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//ctx := sdk.Context()
	//client, err := mspclient.New(ctx)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//info, err := client.GetCAInfo()
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//fmt.Println(info.CAName)
	//fmt.Println(info.Version)
	//
	//affiliations, err := client.GetAllIdentities()
	//if err != nil {
	//	log.Printf("%s \n",err)
	//}
	//
	//for _ , info := range affiliations{
	//	fmt.Println(info.ID)
	//	fmt.Println(info.Type)
	//	fmt.Println(info.Attributes)
	//	fmt.Println("----------------------")
	//}
	//
	//a1 := mspclient.Attribute{
	//	Name: "hf.Registrar.Roles",
	//	Value:"client,orderer,peer,user",
	//}
	//
	//a2 := mspclient.Attribute{
	//	Name: "hf.Registrar.DelegateRoles",
	//	Value:"client,orderer,peer,user",
	//}
	//
	//a3 := mspclient.Attribute{
	//	Name: "hf.Registrar.Attributes",
	//	Value:"*",
	//}
	//
	//a4 := mspclient.Attribute{
	//	Name: "hf.GenCRL",
	//	Value:"true",
	//}
	//
	//a5 := mspclient.Attribute{
	//	Name: "hf.Revoker",
	//	Value:"true",
	//}
	//
	//a6 := mspclient.Attribute{
	//	Name: "hf.AffiliationMgr",
	//	Value:"true",
	//}
	//
	//a7 := mspclient.Attribute{
	//	Name: "hf.IntermediateCA",
	//	Value:"true",
	//}
	//
	//var attributes []mspclient.Attribute
	//attributes = append(attributes,a1,a2,a3,a4,a5,a6,a7)
	//
	//req := &mspclient.RegistrationRequest{
	//	Name: "User1",
	//	Type: "admin",
	//	CAName: "ca-org2",
	//	Secret: "User1",
	//	Attributes: attributes,
	//	Affiliation: "org2",
	//}
	//register, err := client.Register(req)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//fmt.Println(register)
	//
	//err = client.Enroll("User1",mspclient.WithSecret("User1"))
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//signingIdentity, err := client.GetSigningIdentity("User1")
	//if err != nil {
	//	fmt.Printf("GetSigningIdentity : %s \n",err)
	//}
	//fmt.Println(signingIdentity.Identifier().ID)


}

























