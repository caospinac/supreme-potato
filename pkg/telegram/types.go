package telegram

type SendMessageRequest struct {
	ChatID    string `json:"chat_id"`
	Text      string `json:"text"`
	ParseMode string `json:"parse_mode"`
}

type SendChatActionRequest struct {
	ChatID string `json:"chat_id"`
	Action string `json:"action"`
}

type TelegramResult struct {
	Ok     bool `json:"ok"`
	Result any  `json:"result"`
}
