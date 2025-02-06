package backend

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/maprost/codeExample/server/chat/backend/obj"
)

func initChat(router *gin.Engine) {
	router.GET("/rest/chat/history", getChatHistory)
	router.GET("/rest/chat", getChat)
	router.POST("/rest/chat", addChat)
}

type ChatElementDto struct {
	Date string
	Line string
}

func getChatHistory(con *gin.Context) {
	fmt.Println("getChatHistory")
	chatC := data.NewChatClient()
	hist := chatC.GetChatHistory()

	var res []ChatElementDto
	for _, elem := range hist {
		res = append(res, ChatElementDto{
			Date: elem.Date.Format(config.DateTimeFormat),
			Line: elem.Line,
		})
	}
	con.JSON(http.StatusOK, res)
}

func getChat(con *gin.Context) {
	fmt.Println("getChat")
	tStr := con.Query("t")
	t := time.Time{}
	var err error
	if tStr != "" {
		t, err = time.ParseInLocation(config.DateTimeFormat, tStr, time.Local)
		if err != nil {
			http.Error(con.Writer, fmt.Sprintf("time parse error: %s", err.Error()), http.StatusNotFound)
			return
		}
	}

	chatC := data.NewChatClient()
	newChats := chatC.GetChat(t)

	res := make([]ChatElementDto, 0)
	for _, elem := range newChats {
		res = append(res, ChatElementDto{
			Date: elem.Date.Format(config.DateTimeFormat),
			Line: elem.Line,
		})
	}
	con.JSON(http.StatusOK, res)
}

func addChat(con *gin.Context) {
	fmt.Println("addChat")
	bytes, err := io.ReadAll(con.Request.Body)
	if err != nil {
		http.Error(con.Writer, fmt.Sprintf("can't read body: %s", err.Error()), http.StatusNotFound)
		return
	}
	line := string(bytes)
	date := config.Now()

	fmt.Println("addChat: ", line)

	chatC := data.NewChatClient()
	chatC.AddChat(obj.ChatElement{
		Date: date,
		Line: line,
	})

	con.Writer.WriteHeader(http.StatusNoContent)
}
