package wire

import (
	"gengycSrc/gycSearchEngine/common"
	"gengycSrc/gycSearchEngine/config"
	"gengycSrc/gycSearchEngine/handlers"
	"gengycSrc/gycSearchEngine/persist"
	"gengycSrc/gycSearchEngine/service"
	"github.com/google/wire"
)

func InitializeHandler(conf *config.ServerConfig) *handlers.LianjiaHandler  {
	wire.Build(
		common.NewEsClient,
		common.NewRouterClient,
		persist.NewLianjiaES,
		service.NewLianjiaService,
		handlers.NewLianjiaHandler)

	return &handlers.LianjiaHandler{}
}


