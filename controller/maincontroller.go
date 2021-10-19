package controller

import (
	"github.com/kataras/iris/v12"
	"goExplore/service"
)

func StartIris()  {
	app := iris.New()

	blocksApi := app.Party("/blocks")
	{
		blocksApi.Use(iris.Compression)

		blocksApi.Get("/QueryLastesBlocksInfo",service.GetLastesBlocksInfo)
		blocksApi.Get("/QueryBlockByBlockNum",service.QueryBlockByBlockNum)

	}

	txsApi := app.Party("/tx")
	{
		txsApi.Use(iris.Compression)
		txsApi.Get("/QueryTxByTxId",service.QueryTxByTxId)
		txsApi.Get("/QueryTxByTxIdJsonStr",service.QueryTxByTxIdJsonStr)
	}

	app.Listen(":9090")

}





