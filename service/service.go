package service

import (
	"github.com/kataras/iris/v12/context"
	"goExplore/common"
	"log"
	"strconv"
)

// @Summary ping example
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} string "pong"
// @Failure 400 {string} string "ok"
// @Failure 404 {string} string "ok"
// @Failure 500 {string} string "ok"
// @Router /examples/ping [get]
func GetLastesBlocksInfo(context *context.Context) {
	blocks, err := common.QueryLastesBlocksInfo()
	if err != nil {
		log.Println(err)
	}
	context.JSON(blocks)
}

func QueryAllBlocksInfo(context *context.Context) {
	blocks, err := common.QueryAllBlocksInfo()
	if err != nil {
		log.Println(err)
	}
	context.JSON(blocks)
}

func QueryTxByTxId(context *context.Context) {

	txId := context.URLParam("txId")
	if txId == "" {
		context.JSON("fail")
	} else {
		transactions, err := common.QueryTxByTxId(txId)
		if err != nil {
			log.Println(err)
		}

		context.JSON(transactions)
	}
}

func QueryTxByTxIdJsonStr(context *context.Context) {

	txId := context.URLParam("txId")
	if txId == "" {
		context.JSON("fail")
	} else {
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
	} else {

		num, _ := strconv.ParseInt(blockNum, 10, 64)
		transactions, err := common.QueryBlockByBlockNum(num)
		if err != nil {
			log.Println(err)
		}

		context.JSON(transactions)
	}
}
