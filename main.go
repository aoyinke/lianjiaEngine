package main

import (
	"gengycSrc/gycSearchEngine/common"
	"gengycSrc/gycSearchEngine/handlers"
	"gengycSrc/gycSearchEngine/persist"
	"gengycSrc/gycSearchEngine/service"
)

// https://changzhou.lianjia.com/ershoufang/pg1/
func main()  {


	cfg := InitConfig()
	server := common.NewRouterClient()
	esClient:= common.NewEsClient(&cfg)
	lianjiaEs := persist.NewLianjiaES(esClient)
	lianjiaService := service.NewLianjiaService(lianjiaEs)
	handler := handlers.NewLianjiaHandler(server,lianjiaService)
	handler.Run()

}
