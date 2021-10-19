package common

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/ledger"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/fab"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	localConfig "goExplore/config"
	"goExplore/models"
	"goExplore/util"
	"log"
	"os"
)

var mainSDK *fabsdk.FabricSDK
var ledgerClient *ledger.Client

// InitChainExploreService Init ChainExplore Client (ledger client)
func InitChainExploreService(cfg, org, admin, user string) *models.Client{

	log.Println("Initialize the client")

	client := &models.Client{
		ConfigPath: cfg,
		OrgName:    org,
		OrgAdmin:   admin,
		OrgUser:    user,

		ChaincodeID:   localConfig.ChaincodeID,
		ChaincodePath: localConfig.ChaincodePath,
		GoPath:        os.Getenv("GOPATH"),
		ChannelID:     localConfig.ChannelID,
	}

	var err error
	// create fabsdk SDk
	mainSDK, err = fabsdk.New(config.FromFile(client.ConfigPath))
	if err != nil {
		log.Panicf("Failed to create an new SDK:%s\n",err)
	}

	client.SDK = mainSDK
	log.Println("Success to create an new SDK")
	// get channel context
	userChannelContext := mainSDK.ChannelContext(client.ChannelID, fabsdk.WithUser(client.OrgUser))
	// ledger client
	ledgerClient,err = ledger.New(userChannelContext)
	if err != nil {
		log.Printf("Failed to create an new ledgerClient:%s\n",err)
	}
	log.Println("Success to create an new ledgerClient ")

	return client

}

// QueryLedgerInfo Query ledger info
func QueryLedgerInfo() (*fab.BlockchainInfoResponse, error)  {

	log.Println("Query ledger info")

	ledgerInfo , err := ledgerClient.QueryInfo()
	if err != nil {
		log.Printf("Failed to query blockChain info: %s\n ",err)
		return nil, err
	}

	return ledgerInfo,nil

}

// QueryLastesBlocksINfo Query last 5 Blocks info
func QueryLastesBlocksInfo() ([]*models.Block , error) {

	ledgerInfo , err := ledgerClient.QueryInfo()
	if err != nil {
		log.Printf("Failed to query last 5 Blocks info:%s \n",err)
		return nil, err
	}

	var lastesBlockList []*models.Block
	lastesBlockNum := ledgerInfo.BCI.Height - 1

	for i := lastesBlockNum; i > 0 && i > (lastesBlockNum - 5 ) ; i-- {
		block , err := QueryBlockByBlockNum(int64(i))
		if err != nil {
			log.Printf("Failed to Query last 5 Blocks info:%s \n",err)
			return nil, err
		}
		lastesBlockList = append(lastesBlockList,block)
	}

	return lastesBlockList, nil
}

// QueryBlockByBlockNum Query Block info by block's number
func QueryBlockByBlockNum(num int64) (*models.Block, error) {

	rawBlock , err := ledgerClient.QueryBlock(uint64(num))
	if err != nil {
		log.Printf("Failed to query Block info by block's number : %s \n",err)
		return nil, err
	}

	// parse the block body

	var txList []*models.Transaction

	for i := range rawBlock.Data.Data{
		rawEnvelope , err := util.GetEnvelopeFromBlock(rawBlock.Data.Data[i])
		if err != nil{
			log.Printf("Failed to GetEnvelopeFromBlock: %s \n",err)
			return nil, err
		}

		transaction, err := util.GetTxFromEnvelopeDeep(rawEnvelope)
		if err != nil{
			log.Printf("Failed to GetTxFromEnvelopeDeep: %s \n",err)
			return nil, err
		}

		for i := range transaction.TransactionActionList {
			transaction.TransactionActionList[i].BlockNum = rawBlock.Header.Number
		}

		txList = append(txList, transaction)
	}

	block := models.Block{

		Number: rawBlock.Header.Number,
		PreviousHash: rawBlock.Header.PreviousHash,
		DataHash: rawBlock.Header.DataHash,
		BlockHash: rawBlock.Header.DataHash,
		TxNum: len(rawBlock.Data.Data),
		TransactionList: txList,
		CreateTime: txList[0].TransactionActionList[0].Timestamp,
	}

	return &block, nil
}

func QueryTxByTxId(txId string) (*models.Transaction,error) {
	rawTx,err := ledgerClient.QueryTransaction(fab.TransactionID(txId))
	if err != nil{
		log.Printf("Failed to QueryTxByTxId rawTx: %s \n",err)
		return nil, err
	}

	transaction,err := util.GetTxFromEnvelopeDeep(rawTx.TransactionEnvelope)
	if err != nil{
		log.Printf("Failed to QueryTxByTxId transaction: %s \n",err)
		return nil, err
	}

	block , err := ledgerClient.QueryBlockByTxID(fab.TransactionID(txId))
	if err != nil{
		log.Printf("Failed to QueryTxByTxId block: %s \n",err)
		return nil, err
	}

	for i := range transaction.TransactionActionList{
		transaction.TransactionActionList[i].BlockNum = block.Header.Number
	}

	return transaction,nil
}

func QueryTxByTxIdJsonStr(txId string) (string,error) {
	transaction , err := QueryTxByTxId(txId)
	if err != nil{
		log.Printf("Failed to QueryTxByTxIdJsonStr transaction: %s \n",err)
		return "", err
	}

	jsonStr, err := json.Marshal(transaction)
	return string(jsonStr),err
}



// OperateLedgerTest Test for operate ledger
func OperateLedgerTest() {

	bci, err := ledgerClient.QueryInfo()
	if err != nil {
		fmt.Printf("failed to query for blockchain info: %s\n", err)
	}

	if bci != nil {
		fmt.Println("Retrieved ledger info")
	}
	fmt.Println(bci.BCI.Height)
	fmt.Println(bci.BCI.PreviousBlockHash)

	rawBlock, err := ledgerClient.QueryBlock(uint64(4))
	fmt.Println(rawBlock.GetHeader().GetDataHash())

}



