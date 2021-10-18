package main

import (
	"fmt"
	"log"
	"os"

	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/ledger"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/resmgmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
)

const (
	orgGoConfig   = "../orggosdk-config.yaml"
	orgCppConfig  = "../orgcppsdk-config.yaml"
	org1          = "OrgGo"
	org2          = "OrgCpp"
	admin         = "Admin"
	user          = "User1"
	chaincodePath = "github.com/hyperledger/fabric-samples/chaincode/chaincode_example02/go"
)

func main() {
	org1Client := NewClient(orgGoConfig, org1, admin, user)
	org2Client := NewClient(orgCppConfig, org2, admin, user)

	// Close: 关闭并释放有SDK维护的缓存和连接
	defer org1Client.SDK.Close()
	defer org2Client.SDK.Close()
}

func NewClient(cfg, org, admin, user string) *Client {
	client := &Client{
		ConfigPath: cfg,
		OrgName:    org,
		OrgAdmin:   admin,
		OrgUser:    user,

		ChaincodeID:   "example",
		ChaincodePath: chaincodePath,
		GoPath:        os.Getenv("GOPATH"),
		ChannelID:     "mychannel",
	}
	// fabsdk.New：创建fabsdk实例
	sdk, err := fabsdk.New(config.FromFile(client.ConfigPath))
	if err != nil {
		log.Panicf("创建SDK失败：%s", err)
	}
	client.SDK = sdk
	log.Println("SDK初始化完成")
	client.rc, client.cc = NewSDKClient(sdk, client.ChannelID, client.OrgName, client.OrgAdmin, client.OrgUser)
	return client
}

func NewSDKClient(sdk *fabsdk.FabricSDK, channelID string, orgName string, orgAdmin string, orgUser string) (rc *resmgmt.Client, cc *channel.Client) {
	var err error

	//WithIdentity：使用预先构造的身份对象作为访问的凭据
	//WithOrg：使用命名的组织
	//WithUser：使用命名的用户来加载对应的身份证书
	//Context：创建并返回上下文客户端
	rcp := sdk.Context(fabsdk.WithUser(orgAdmin), fabsdk.WithOrg(orgName))
	// resmgmt.New: 返回一个client实例
	rc, err = resmgmt.New(rcp)
	if err != nil {
		log.Panicf("创建RC客户端失败:%s", err)
	}
	log.Println("RC初始化完成")
	//ChannelContext：创建并返回通道上下文
	ccp := sdk.ChannelContext(channelID, fabsdk.WithUser(orgUser))
	//channel.New: 返回一个Client实例
	cc, err = channel.New(ccp)
	if err != nil {
		log.Panicf("创建通道客户端失败：%s", err)
	}
	log.Println("通道客户端初始化完成")
	response, err := cc.Query(channel.Request{
		ChaincodeID: "sacc2",
		Fcn:         "invoke",
		Args:        [][]byte{[]byte("a")},
	})

	if err != nil {
		log.Printf("failed to query chaincode : %s\n", err)
	}

	if len(response.Payload) > 0 {
		log.Println("chaincode query success")
		log.Println(response.TransactionID)
	}

	ctx, err := ledger.New(ccp)
	bci, err := ctx.QueryInfo()
	if err != nil {
		fmt.Printf("failed to query for blockchain info: %s\n", err)
	}

	if bci != nil {
		fmt.Println("Retrieved ledger info")
	}
	fmt.Println(bci.BCI.Height)

	return rc, cc
}

type Client struct {
	// Fabric 网络信息
	ConfigPath string
	OrgName    string
	OrgAdmin   string
	OrgUser    string

	//sdk 客户端
	SDK *fabsdk.FabricSDK
	rc  *resmgmt.Client
	cc  *channel.Client

	// 链码信息
	ChannelID     string //通道名
	ChaincodeID   string //链码ID或者名称
	ChaincodePath string //链码路径
	GoPath        string // GOPATH路径
}
