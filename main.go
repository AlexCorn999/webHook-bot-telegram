package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var (
	NewBot, BotErr := tgbotapi.NewBotAPI(betypes.BotToken)
)

func main() {

}
