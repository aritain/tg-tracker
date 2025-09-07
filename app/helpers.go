package app

import (
	"encoding/json"
	"log"
	"os"
	"time"

	t "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const FILEPATH = "/data/value.json"
const TIMEOUT = 45

func GetValue() (value Balance) {
	filepath := FILEPATH
	data, err := os.ReadFile(filepath)
	if err == nil {
		_ = json.Unmarshal(data, &value)
	}
	return
}

func WriteValue(value Balance) {
	filepath := FILEPATH
	file, _ := os.OpenFile(filepath, os.O_WRONLY|os.O_TRUNC, 0600)
	defer file.Close()
	json.NewEncoder(file).Encode(value)
}

func SendTGMessage(tgm TGMessage, sticker bool) {
	var msg t.Chattable
	config := Get().Config()
	if len(tgm.TGToken) == 0 {
		tgm.TGToken = config.TGToken
	}
	bot, _ := t.NewBotAPI(tgm.TGToken)
	if sticker {
		msg = t.NewSticker(tgm.UserID, t.FileID(config.StickerID))
	} else {
		message := t.NewMessage(tgm.UserID, tgm.Text)
		if len(tgm.Keyboard.InlineKeyboard) == 0 {
			tgm.Keyboard = CompileQueryKeyboard()
		}
		message.ReplyMarkup = tgm.Keyboard
		msg = message
	}

	var err error

	for {
		_, err = bot.Send(msg)
		if err == nil {
			break
		}
		log.Print(err)
		time.Sleep(TIMEOUT * time.Second)
	}
}
