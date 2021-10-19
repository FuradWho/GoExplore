package main

import (
	"fmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/ledger"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/resmgmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	"log"
	"os"
)

const (
	orgGoConfig   = "../cliconfigs/orggo-config.yaml"
	orgCppConfig  = "../cliconfigs/orgcpp-config.yaml"
	org1          = "OrgGo"
	org2          = "OrgCpp"
	admin         = "Admin"
	user          = "User1"
	chaincodePath = "github.com/hyperledger/fabric-samples/chaincode/chaincode_example02/go"
)

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

type Block struct {
	Number          uint64         `json:"number"`          //区块号
	PreviousHash    []byte         `json:"previousHash"`    //前区块Hash
	DataHash        []byte         `json:"dataHash"`        //交易体Hash
	BlockHash       []byte         `json:"blockHash"`       //区块Hash
	TxNum           int            `json:"txNum"`           //区块内交易个数
	TransactionList []*Transaction `json:"transactionList"` //交易列表
	CreateTime      string         `json:"createTime"`      //区块生成时间
}
type Transaction struct {
	TransactionActionList []*TransactionAction `json:"transactionActionList"` //交易列表
}
type TransactionAction struct {
	TxId         string   `json:"txId"`         //交易ID
	BlockNum     uint64   `json:"blockNum"`     //区块号
	Type         string   `json:"type"`         //交易类型
	Timestamp    string   `json:"timestamp"`    //交易创建时间
	ChannelId    string   `json:"channelId"`    //通道ID
	Endorsements []string `json:"endorsements"` //背书组织ID列表
	ChaincodeId  string   `json:"chaincodeId"`  //链代码名称
	ReadSetList  []string `json:"readSetList"`  //读集
	WriteSetList []string `json:"writeSetList"` //写集
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

func NewSDKClient(sdk *fabsdk.FabricSDK, channelID string, orgName string, orgAdmin string, orgUser string) (rc *resmgmt.Client, ledgerClient *channel.Client) {
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
	ledgerClient, err = channel.New(ccp)
	if err != nil {
		log.Panicf("创建通道客户端失败：%s", err)
	}
	log.Println("通道客户端初始化完成")
	response, err := ledgerClient.Query(channel.Request{
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
	fmt.Println(bci.BCI.PreviousBlockHash)

	rawBlock, err := ctx.QueryBlock(uint64(4))
	fmt.Println(rawBlock.GetHeader().GetDataHash())

	return rc, ledgerClient
}

func main() {
	org1Client := NewClient(orgGoConfig, org1, admin, user)
	//org2Client := NewClient(orgCppConfig, org2, admin, user)

	// Close: 关闭并释放有SDK维护的缓存和连接
	defer org1Client.SDK.Close()
	//defer org2Client.SDK.Close()
}
