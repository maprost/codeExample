package main

import (
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/maprost/codeExample/server/chat/backend"
	"github.com/maprost/codeExample/server/chat/backend/cfg"
)

func main() {
	conf := cfg.NewConfig()

	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(gzip.Gzip(gzip.DefaultCompression))
	backend.Init(router, conf)

	err := router.Run(conf.Server.Host + ":" + conf.Server.Port)
	if err != nil {
		panic(err)
	}
}
