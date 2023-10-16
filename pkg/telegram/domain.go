package telegram

import "time"

type Update struct {
	UpdateID int64   `json:"update_id"`
	Message  Message `json:"message"`
}

type Message struct {
	MessageID      int64           `json:"message_id"`
	From           MessageFrom     `json:"from"`
	Chat           Chat            `json:"chat"`
	Date           time.Time       `json:"date"`
	ReplyToMessage *ReplyToMessage `json:"reply_to_message"`
	Text           string          `json:"text"`
	Entities       []Entity        `json:"entities"`
}

type Chat struct {
	ID        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	Type      string `json:"type"`
}

type Entity struct {
	Offset int64  `json:"offset"`
	Length int64  `json:"length"`
	Type   string `json:"type"`
}

type MessageFrom struct {
	ID           int64  `json:"id"`
	IsBot        bool   `json:"is_bot"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Username     string `json:"username"`
	LanguageCode string `json:"language_code"`
}

type ReplyToMessage struct {
	MessageID int64              `json:"message_id"`
	From      ReplyToMessageFrom `json:"from"`
	Chat      Chat               `json:"chat"`
	Date      time.Time          `json:"date"`
	Text      string             `json:"text"`
}

type ReplyToMessageFrom struct {
	ID        int64  `json:"id"`
	IsBot     bool   `json:"is_bot"`
	FirstName string `json:"first_name"`
	Username  string `json:"username"`
}
