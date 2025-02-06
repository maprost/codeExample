package datatier

import (
	"sync"
	"time"

	"github.com/maprost/codeExample/server/chat/backend/obj"
)

type chatDB struct {
	history      []obj.ChatElement
	historyMutex sync.RWMutex
}

func newChatDB() *chatDB {
	return &chatDB{
		history:      make([]obj.ChatElement, 0),
		historyMutex: sync.RWMutex{},
	}
}

type ChatClient struct {
	chatDB *chatDB
}

func (x *ChatClient) GetChatHistory() []obj.ChatElement {
	x.chatDB.historyMutex.RLock()
	defer x.chatDB.historyMutex.RUnlock()

	var res []obj.ChatElement
	for _, e := range x.chatDB.history {
		res = append(res, e)
	}
	return res
}

func (x *ChatClient) GetChat(t time.Time) []obj.ChatElement {
	x.chatDB.historyMutex.RLock()
	defer x.chatDB.historyMutex.RUnlock()

	var res []obj.ChatElement
	for _, e := range x.chatDB.history {
		if e.Date.After(t) {
			res = append(res, e)
		}
	}
	return res
}

func (x *ChatClient) AddChat(elem obj.ChatElement) {
	x.chatDB.historyMutex.Lock()
	defer x.chatDB.historyMutex.Unlock()

	x.chatDB.history = append(x.chatDB.history, elem)
}
