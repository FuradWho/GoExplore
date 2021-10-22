package controller

import (
	"github.com/iris-contrib/swagger/v12"
	"github.com/iris-contrib/swagger/v12/swaggerFiles"
	"github.com/kataras/iris/v12"
	"goExplore/service"

	_ "goExplore/docs"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /v2
func StartIris() {
	app := iris.New()
	app.Use(Cors)
	config := &swagger.Config{
		URL: "http://localhost:9090/swagger/doc.json", //The url pointing to API definition
	}

	// swagger config
	swaggerApi := app.Party("/swagger")
	{
		swaggerApi.Get("/{any:path}", swagger.CustomWrapHandler(config, swaggerFiles.Handler))
	}

	// blocks API operate
	blocksApi := app.Party("/blocks")
	{
		blocksApi.Use(iris.Compression)

		blocksApi.Get("/QueryLastesBlocksInfo", service.GetLastesBlocksInfo)
		blocksApi.Get("/QueryBlockByBlockNum", service.QueryBlockByBlockNum)
		blocksApi.Get("/QueryAllBlocksInfo", service.QueryAllBlocksInfo)
		blocksApi.Get("/QueryBlockInfoByHash", service.QueryBlockInfoByHash)
		blocksApi.Get("/QueryBlockMainInfo",service.QueryBlockMainInfo)

	}

	// txs API operate
	txsApi := app.Party("/txs")
	{
		txsApi.Use(iris.Compression)

		txsApi.Get("/QueryTxByTxId", service.QueryTxByTxId)
		txsApi.Get("/QueryTxByTxIdJsonStr", service.QueryTxByTxIdJsonStr)
	}



	app.Listen(":9090")
}

// Cors Resolve the CORS
func Cors(ctx iris.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")
	if ctx.Request().Method == "OPTIONS" {
		ctx.Header("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,PATCH,OPTIONS")
		ctx.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Authorization")
		ctx.StatusCode(204)
		return
	}
	ctx.Next()
}
