package backend

import (
	"github.com/gin-gonic/gin"
	"github.com/maprost/codeExample/server/chat/backend/cfg"
	"github.com/maprost/codeExample/server/chat/backend/datatier"
)

var (
	config *cfg.Config
	data   *datatier.DataTier
)

func Init(router *gin.Engine, conf *cfg.Config) {
	config = conf
	data = datatier.NewCreator()

	initChat(router)
	initWeb(router)
}
