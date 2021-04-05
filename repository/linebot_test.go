package repository_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/stretchr/testify/assert"
)

func TestLineBotTestClient(t *testing.T) {
	// arrange
	serv := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte("{}"))
		}))

	// act
	bot, _ := linebot.New("test-secret", "test-token", linebot.WithEndpointBase(serv.URL))
	_, actualPushMessageErr := bot.PushMessage("testbot").Do()
	_, actualReplayMessageErr := bot.ReplyMessage("testbot").Do()

	// assert
	assert.NoError(t, actualPushMessageErr)
	assert.NoError(t, actualReplayMessageErr)
}
