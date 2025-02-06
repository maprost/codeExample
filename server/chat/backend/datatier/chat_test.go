package datatier_test

import (
	"testing"
	"time"

	"github.com/maprost/codeExample/server/chat/backend/cfg"
	"github.com/maprost/codeExample/server/chat/backend/datatier"
	"github.com/maprost/codeExample/server/chat/backend/obj"
	"github.com/maprost/testbox/should"
)

func TestChatClient(t *testing.T) {
	init := func() datatier.ChatClient {
		creator := datatier.NewCreator()
		chatC := creator.NewChatClient()
		return chatC
	}

	t.Run("check empty GetChatHistory", func(t *testing.T) {
		chatC := init()
		should.HaveLength(t, chatC.GetChatHistory(), 0)
	})

	t.Run("check empty GetChat", func(t *testing.T) {
		chatC := init()
		should.HaveLength(t, chatC.GetChat(time.Time{}), 0)
	})

	t.Run("check simple add - get workflow", func(t *testing.T) {
		chatC := init()
		now := cfg.Now()
		oneMinuteBefore := now.Add(-time.Minute)
		oneMinuteAfter := now.Add(time.Minute)

		elem1 := obj.ChatElement{
			Date: oneMinuteBefore,
			Line: "oneMinuteBefore",
		}
		chatC.AddChat(elem1)
		elem2 := obj.ChatElement{
			Date: now,
			Line: "now",
		}
		chatC.AddChat(elem2)
		elem3 := obj.ChatElement{
			Date: oneMinuteAfter,
			Line: "oneMinuteAfter",
		}
		chatC.AddChat(elem3)

		t.Run("check history", func(t *testing.T) {
			hist := chatC.GetChatHistory()
			should.HaveLength(t, hist, 3)
			should.BeEqual(t, hist[0], elem1)
			should.BeEqual(t, hist[1], elem2)
			should.BeEqual(t, hist[2], elem3)
		})

		t.Run("check getChat", func(t *testing.T) {
			t.Run("check oneMinuteBefore", func(t *testing.T) {
				hist := chatC.GetChat(oneMinuteBefore)
				should.HaveLength(t, hist, 2)
				should.BeEqual(t, hist[0], elem2)
				should.BeEqual(t, hist[1], elem3)
			})
			t.Run("check now", func(t *testing.T) {
				hist := chatC.GetChat(now)
				should.HaveLength(t, hist, 1)
				should.BeEqual(t, hist[0], elem3)
			})
			t.Run("check oneMinuteAfter", func(t *testing.T) {
				hist := chatC.GetChat(oneMinuteAfter)
				should.HaveLength(t, hist, 0)
			})
		})
	})
}
