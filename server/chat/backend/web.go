package backend

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/maprost/codeExample/server/chat/frontend"
)

func initWeb(router *gin.Engine) {
	addWebsite := func(key string, file string) {
		router.GET(key, func(con *gin.Context) {
			if file == "" {
				file = con.Param("web")
			}
			content, err := os.ReadFile(frontend.BasePath() + file)
			if err != nil {
				panic(err)
			}
			_, err = con.Writer.Write(content)
			if err != nil {
				panic(err)
			}
		})
	}

	addWebsite("/", "index.html")
	addWebsite("/p/*web", "") // for more pages
}
