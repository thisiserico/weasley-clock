package telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"path"
	"time"

	"github.com/thisiserico/weasley-clock/pkg/firebase"
)

const telegramTokenEnvVar = "TELEGRAM_TOKEN"

var endpoint *url.URL

func init() {
	rand.Seed(time.Now().UnixNano())

	token := os.Getenv(telegramTokenEnvVar)
	address := fmt.Sprintf("https://api.telegram.org/bot%s", token)
	endpoint, _ = url.Parse(address)
}

type reply struct {
	ChatID      int                 `json:"chat_id"`
	Text        string              `json:"text"`
	ReplyMarkup map[string]keyboard `json:"reply_markup"`
}

type keyboard [][]status

type status string

func prepareReply(chatID int, statuses []string) reply {
	emojis := []string{"🧙", "🔮", "⚡"}
	emoji := emojis[rand.Intn(len(emojis))]
	text := fmt.Sprintf("Where should your clock hand point at next? %s", emoji)

	kb := make(keyboard, 1)
	row := 0
	for _, st := range statuses {
		if len(kb[row]) == 3 {
			kb = append(kb, make([]status, 0, 3))
			row += 1
		}

		kb[row] = append(kb[row], status(st))
	}

	return reply{
		ChatID:      chatID,
		Text:        text,
		ReplyMarkup: map[string]keyboard{"keyboard": kb},
	}
}

func RequestNextStatus(chatID int, statuses firebase.Statuses) error {
	u, _ := url.Parse(endpoint.String())
	u.Path = path.Join(u.Path, "sendMessage")

	payload := prepareReply(chatID, statuses)
	js, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, u.String(), bytes.NewBuffer(js))
	if err != nil {
		return err
	}
	req.Header.Add("content-type", "application/json")

	client := &http.Client{}
	_, err = client.Do(req)

	return err
}
