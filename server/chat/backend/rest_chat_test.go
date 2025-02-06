package backend_test

import (
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/maprost/codeExample/server/chat/backend"
	"github.com/maprost/codeExample/server/chat/backend/cfg"
	"github.com/maprost/restclient"
	"github.com/maprost/testbox/must"
	"github.com/maprost/testbox/should"
)

type testBackend struct {
	config  *cfg.Config
	now     time.Time
	server  *httptest.Server
	baseUrl string
}

func newTestBackend() testBackend {
	config := cfg.NewConfig()
	now := config.Now()
	config.Now = func() time.Time {
		return now
	}
	router := gin.New()
	backend.Init(router, config)
	server := httptest.NewServer(router)
	baseUrl := server.URL

	return testBackend{
		config:  config,
		now:     now,
		server:  server,
		baseUrl: baseUrl,
	}
}

func (xx testBackend) getChatHistory(t *testing.T) []backend.ChatElementDto {
	c := restclient.Get(xx.baseUrl + "/rest/chat/history")
	var dtos []backend.ChatElementDto
	res := c.SendAndGetJsonResponse(&dtos)
	must.BeNoError(t, res.Error())
	return dtos
}

func (xx testBackend) getChat(t *testing.T, d time.Time) []backend.ChatElementDto {
	c := restclient.Get(xx.baseUrl + "/rest/chat?t=" + url.QueryEscape(d.Format(xx.config.DateTimeFormat)))
	var dtos []backend.ChatElementDto
	res := c.SendAndGetJsonResponse(&dtos)
	must.BeNoError(t, res.Error())
	return dtos
}

func (xx testBackend) addChat(t *testing.T, line string) {
	c := restclient.Post(xx.baseUrl + "/rest/chat")
	c.AddBody([]byte(line), "application/txt; charset=utf-8")
	res := c.Send()
	must.BeNoError(t, res.Error())
}

func TestChatClient(t *testing.T) {
	t.Run("check empty history", func(t *testing.T) {
		xx := newTestBackend()

		chat := xx.getChatHistory(t)
		must.HaveLength(t, chat, 0)
	})

	t.Run("check empty getChat", func(t *testing.T) {
		xx := newTestBackend()

		chat := xx.getChat(t, time.Time{})
		must.HaveLength(t, chat, 0)
	})

	t.Run("check addChat once", func(t *testing.T) {
		xx := newTestBackend()
		xx.addChat(t, "hello world")

		t.Run("check history", func(t *testing.T) {
			hist := xx.getChatHistory(t)
			should.HaveLength(t, hist, 1)
			should.BeEqual(t, hist[0], backend.ChatElementDto{Date: xx.now.Format(xx.config.DateTimeFormat), Line: "hello world"})
		})

		t.Run("check getChat", func(t *testing.T) {
			chat := xx.getChat(t, time.Time{})
			should.HaveLength(t, chat, 1)
			should.BeEqual(t, chat[0], backend.ChatElementDto{Date: xx.now.Format(xx.config.DateTimeFormat), Line: "hello world"})
		})
	})

	t.Run("check addChat four times", func(t *testing.T) {
		xx := newTestBackend()
		nowPlusOneMinute := xx.now.Add(time.Minute)
		nowPlusTwoMinute := xx.now.Add(2 * time.Minute)
		nowPlusThreeMinute := xx.now.Add(3 * time.Minute)

		xx.addChat(t, "hello")
		xx.config.Now = func() time.Time {
			return nowPlusOneMinute
		}
		xx.addChat(t, "world")
		xx.config.Now = func() time.Time {
			return nowPlusTwoMinute
		}
		xx.addChat(t, "blob")
		xx.config.Now = func() time.Time {
			return nowPlusThreeMinute
		}
		xx.addChat(t, "drop")

		t.Run("check history", func(t *testing.T) {
			hist := xx.getChatHistory(t)
			should.HaveLength(t, hist, 4)
			should.BeEqual(t, hist[0], backend.ChatElementDto{Date: xx.now.Format(xx.config.DateTimeFormat), Line: "hello"})
			should.BeEqual(t, hist[1], backend.ChatElementDto{Date: nowPlusOneMinute.Format(xx.config.DateTimeFormat), Line: "world"})
			should.BeEqual(t, hist[2], backend.ChatElementDto{Date: nowPlusTwoMinute.Format(xx.config.DateTimeFormat), Line: "blob"})
			should.BeEqual(t, hist[3], backend.ChatElementDto{Date: nowPlusThreeMinute.Format(xx.config.DateTimeFormat), Line: "drop"})
		})

		t.Run("check getChat", func(t *testing.T) {
			t.Run("no time", func(t *testing.T) {
				chat := xx.getChat(t, time.Time{})
				should.HaveLength(t, chat, 4)
				should.BeEqual(t, chat[0], backend.ChatElementDto{Date: xx.now.Format(xx.config.DateTimeFormat), Line: "hello"})
				should.BeEqual(t, chat[1], backend.ChatElementDto{Date: nowPlusOneMinute.Format(xx.config.DateTimeFormat), Line: "world"})
				should.BeEqual(t, chat[2], backend.ChatElementDto{Date: nowPlusTwoMinute.Format(xx.config.DateTimeFormat), Line: "blob"})
				should.BeEqual(t, chat[3], backend.ChatElementDto{Date: nowPlusThreeMinute.Format(xx.config.DateTimeFormat), Line: "drop"})
			})
			t.Run("now", func(t *testing.T) {
				chat := xx.getChat(t, xx.now)
				should.HaveLength(t, chat, 3)
				should.BeEqual(t, chat[0], backend.ChatElementDto{Date: nowPlusOneMinute.Format(xx.config.DateTimeFormat), Line: "world"})
				should.BeEqual(t, chat[1], backend.ChatElementDto{Date: nowPlusTwoMinute.Format(xx.config.DateTimeFormat), Line: "blob"})
				should.BeEqual(t, chat[2], backend.ChatElementDto{Date: nowPlusThreeMinute.Format(xx.config.DateTimeFormat), Line: "drop"})
			})
			t.Run("nowPlusOneMinute", func(t *testing.T) {
				chat := xx.getChat(t, nowPlusOneMinute)
				should.HaveLength(t, chat, 2)
				should.BeEqual(t, chat[0], backend.ChatElementDto{Date: nowPlusTwoMinute.Format(xx.config.DateTimeFormat), Line: "blob"})
				should.BeEqual(t, chat[1], backend.ChatElementDto{Date: nowPlusThreeMinute.Format(xx.config.DateTimeFormat), Line: "drop"})
			})
			t.Run("nowPlusTwoMinute", func(t *testing.T) {
				chat := xx.getChat(t, nowPlusTwoMinute)
				should.HaveLength(t, chat, 1)
				should.BeEqual(t, chat[0], backend.ChatElementDto{Date: nowPlusThreeMinute.Format(xx.config.DateTimeFormat), Line: "drop"})
			})
			t.Run("nowPlusThreeMinute", func(t *testing.T) {
				chat := xx.getChat(t, nowPlusThreeMinute)
				should.HaveLength(t, chat, 0)
			})
		})
	})
}
