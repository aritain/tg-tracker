package app

import (
	t "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func CompileFlowKeyboard() t.InlineKeyboardMarkup {
	config := Get().Config()
	var keyboard = t.NewInlineKeyboardMarkup(
		t.NewInlineKeyboardRow(
			t.NewInlineKeyboardButtonData("+", "+"),
			t.NewInlineKeyboardButtonData("-", "-"),
		),
		t.NewInlineKeyboardRow(
			t.NewInlineKeyboardButtonData(config.CancelText, "cancel"),
		),
	)
	return keyboard
}

func CompileQueryKeyboard() t.InlineKeyboardMarkup {
	config := Get().Config()
	var keyboard = t.NewInlineKeyboardMarkup(
		t.NewInlineKeyboardRow(
			t.NewInlineKeyboardButtonData(config.QueryText, config.QueryText),
		),
	)
	return keyboard
}
