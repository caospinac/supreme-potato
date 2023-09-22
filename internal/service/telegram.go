package service

import (
	"context"
	"fmt"

	"github.com/caospinac/supreme-potato/pkg/telegram"
	"github.com/caospinac/supreme-potato/pkg/util"
)

type TelegramService interface {
	ResolveMessage(context.Context, *telegram.Update) (util.ResponseBuilder, util.ApiError)
}

type telegramService struct {
	client telegram.Telegram
}

func NewTelegramService(telegramClient telegram.Telegram) TelegramService {

	return &telegramService{
		telegramClient,
	}
}

func (s *telegramService) ResolveMessage(ctx context.Context, update *telegram.Update) (util.ResponseBuilder, util.ApiError) {
	chatID := fmt.Sprint(update.Message.Chat.ID)
	s.client.SendChatAction(chatID, telegram.ActionTyping)

	_, errMessage := s.client.SendMessage(&telegram.SendMessageRequest{
		ChatID:    chatID,
		Text:      "Hello, " + update.Message.Chat.FirstName,
		ParseMode: telegram.ParseModeMarkdown,
	})
	if errMessage != nil {
		return nil, util.ToApiError(errMessage)
	}

	return util.NewResponse(), nil
}
