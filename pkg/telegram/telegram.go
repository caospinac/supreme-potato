package telegram

import (
	"net/http"

	"github.com/caospinac/supreme-potato/pkg/util"
)

type Telegram interface {
	SendMessage(params *SendMessageRequest) (*TelegramResult, error)
	SendChatAction(chatID, action string) (*TelegramResult, error)
}

type telegram struct {
	token string
}

func NewTelegram(token string) Telegram {
	return &telegram{
		token,
	}
}

func (t *telegram) SendMessage(params *SendMessageRequest) (*TelegramResult, error) {
	url := t.getWebhookURL() + "/sendMessage"
	res := new(TelegramResult)
	request := &util.Request{
		Method: http.MethodPost,
		URL:    url,
		Body:   params,
		Output: res,
	}

	if err := util.MakeRequest(request); err != nil {
		return nil, err
	}

	return res, nil
}

func (t *telegram) SendChatAction(chatID, action string) (*TelegramResult, error) {
	url := t.getWebhookURL() + "/sendChatAction"
	res := new(TelegramResult)
	request := &util.Request{
		Method: http.MethodPost,
		URL:    url,
		Body: &SendChatActionRequest{
			ChatID: chatID,
			Action: action,
		},
		Output: res,
	}

	if err := util.MakeRequest(request); err != nil {
		return nil, err
	}

	return res, nil
}

func (t *telegram) getWebhookURL() string {
	return "https://api.telegram.org/bot" + t.token
}
