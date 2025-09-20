package main

import (
	"fmt"
	"log"
	"slices"
	"strconv"

	a "tracker/app"

	t "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	if err := a.Initialize(); err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	config := a.Get().Config()

	bot, err := t.NewBotAPI(config.TGToken)
	if err != nil {
		log.Panic(err)
	}

	cache := 0

	// Create chan for telegram updates
	var ucfg t.UpdateConfig = t.NewUpdate(0)
	ucfg.Timeout = 60
	updates := bot.GetUpdatesChan(ucfg)
	go a.Notifier()

	for update := range updates {
		var text string
		var userID int64
		var value a.Balance
		var msg a.TGMessage
		if (update.Message == nil) && (update.CallbackQuery == nil) { // ignore any non-Message updates
			continue
		}
		// Treat CallbackQueries the same as a message from user
		if update.CallbackQuery != nil {
			callback := update.CallbackQuery
			userID = callback.Message.Chat.ID
			text = callback.Data
		} else {
			userID = update.Message.Chat.ID
			text = update.Message.Text
		}
		if !slices.Contains(config.BotAdmins, userID) {
			continue
		}
		value = a.GetValue()
		msg.Keyboard = a.CompileQueryKeyboard()
		msg.UserID = userID
		if text == config.QueryText {
			msg.Text = fmt.Sprintf("%s\n", config.QueryResponse)
			msg.Text += strconv.Itoa(value.Value)
		} else if text == "+" && cache != 0 {
			value.Value += cache
			a.WriteValue(value)
			msg.Text = fmt.Sprintf("%s\n", config.QueryResponse)
			msg.Text += strconv.Itoa(value.Value)
			log.Printf("%v added %v to value\n", userID, cache)
			cache = 0
		} else if text == "-" && cache != 0 {
			value.Value -= cache
			a.WriteValue(value)
			msg.Text = fmt.Sprintf("%s\n", config.QueryResponse)
			msg.Text += strconv.Itoa(value.Value)
			log.Printf("%v removed %v from value\n", userID, cache)
			cache = 0
		} else if text == "cancel" {
			cache = 0
			msg.Text = "Ok"
		} else {
			cache, err = strconv.Atoi(text)
			if err != nil {
				log.Printf("Text '%s' from user '%v'\n", text, userID)
				msg.Text = config.ErrorResponse
			} else {
				if cache > 100 || cache < 1 {
					log.Printf("Text '%s' from user '%v'\n", text, userID)
					msg.Text = config.ErrorResponse
				} else {
					msg.Text = config.UserPick
					msg.Keyboard = a.CompileFlowKeyboard()
				}
			}
		}

		a.SendTGMessage(msg, false)
	}
}
