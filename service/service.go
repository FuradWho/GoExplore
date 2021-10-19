package service

import (
	"github.com/kataras/iris/v12/context"
	"goExplore/common"
	"log"
	"strconv"
)

func GetLastesBlocksInfo(context *context.Context) {
	blocks ,err :=common.QueryLastesBlocksInfo()
	if err != nil {
		log.Println(err)
	}
	context.JSON(blocks)
}

func QueryAllBlocksInfo(context *context.Context) {
	blocks ,err :=common.QueryAllBlocksInfo()
	if err != nil {
		log.Println(err)
	}
	context.JSON(blocks)
}


func QueryTxByTxId(context *context.Context) {

	txId := context.URLParam("txId")
	if txId == "" {
		context.JSON("fail")
	}else{
		transactions, err := common.QueryTxByTxId(txId)
		if err != nil {
			log.Println(err)
		}

		context.JSON(transactions)
	}
}

func QueryTxByTxIdJsonStr(context *context.Context){

	txId := context.URLParam("txId")
	if txId == "" {
		context.JSON("fail")
	}else{
		transactions, err := common.QueryTxByTxId(txId)
		if err != nil {
			log.Println(err)
		}

		context.JSON(transactions)
	}
}

func QueryBlockByBlockNum(context *context.Context) {
	blockNum := context.URLParam("blockNum")
	if blockNum == "" {
		context.JSON("fail")
	}else{

		num, _ := strconv.ParseInt(blockNum,10,64)
		transactions, err := common.QueryBlockByBlockNum(num)
		if err != nil {
			log.Println(err)
		}

		context.JSON(transactions)
	}
}
