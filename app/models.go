package app

import t "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type Balance struct {
	Value int `json:"Value"`
}

type Config struct {
	TGToken          string   `mapstructure:"TG_TOKEN"`
	CancelText       string   `mapstructure:"CANCEL_TEXT"`
	QueryText        string   `mapstructure:"QUERY_TEXT"`
	QueryResponse    string   `mapstructure:"QUERY_RESPONSE"`
	ErrorResponse    string   `mapstructure:"ERROR_RESPONSE"`
	UserPick         string   `mapstructure:"USER_PICK"`
	BotAdmins        []int64  `mapstructure:"BOT_ADMINS"`
	NotificationID   int64    `mapstructure:"NOTIFICATION_ID"`
	NotificationText string   `mapstructure:"NOTIFICATION_TEXT"`
	ConfirmationText string   `mapstructure:"CONFIRMATION_TEXT"`
	NTime            string   `mapstructure:"NTIME"`
	CTime            string   `mapstructure:"CTIME"`
	StickerID        []string `mapstructure:"STICKER_ID"`
}

type TGMessage struct {
	TGToken  string
	UserID   int64
	Text     string
	Keyboard t.InlineKeyboardMarkup
}
