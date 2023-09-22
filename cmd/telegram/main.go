package main

import (
	"context"
	"encoding/json"
	"net/http"
	"os"

	"github.com/caospinac/supreme-potato/internal/service"
	"github.com/caospinac/supreme-potato/pkg/telegram"
	"github.com/caospinac/supreme-potato/pkg/util"
)

var (
	token = os.Getenv("TELEGRAM_TOKEN")
)

func main() {
	repo := telegram.NewTelegram(token)
	svc := service.NewTelegramService(repo)

	util.Start(func(ctx context.Context, event util.EventRequest) (util.ResponseBuilder, util.ApiError) {
		update := new(telegram.Update)

		if err := json.Unmarshal([]byte(event.Body), update); err != nil {
			return nil, util.NewApiError().WithStatus(http.StatusBadRequest)
		}

		result, err := svc.ResolveMessage(ctx, update)
		if err != nil {
			return nil, err
		}

		return result, err
	})
}
