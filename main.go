package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/AlexCorn999/webHook-bot-telegram/betypes"
	"github.com/AlexCorn999/webHook-bot-telegram/logger"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var (
	NewBot, BotErr = tgbotapi.NewBotAPI(betypes.BOT_TOKEN)
)

func setWebHook(bot *tgbotapi.BotAPI) {
	webHookInfo := tgbotapi.NewWebhookWithCert(
		fmt.Sprintf("https://%s:%s/%s", betypes.BOT_ADDRESS, betypes.BOT_PORT, betypes.BOT_TOKEN),
		betypes.CERT_PATH)
	_, err := bot.SetWebhook(webHookInfo)
	logger.ForError(err)
}

func main() {
	logger.ForError(BotErr)
	setWebHook(NewBot)

	message := func(w http.ResponseWriter, r *http.Request) {
		text, err := ioutil.ReadAll(r.Body)
		logger.ForError(err)

		var botText betypes.BotMessage
		err = json.Unmarshal(text, &botText)
		logger.ForError(err)

		fmt.Println(fmt.Sprintf("%s", text))
		logger.LogFile.Println(fmt.Sprintf("%s", text))

		username := botText.Message.From.Username
		chatUser := botText.Message.From.Id
		chatGroup := botText.Message.Chat.Id
		messageId := botText.Message.Message_Id
		botCommand := strings.Split(botText.Message.Text, "@")[0]
		commandText := strings.Split(botText.Message.Text, " ")

		fmt.Println(username, chatUser, chatGroup, messageId, botCommand, commandText)
	}

	http.HandleFunc("/", message)
	log.Fatal(http.ListenAndServeTLS(
		fmt.Sprintf("%s:%s", betypes.BOT_ADDRESS, betypes.BOT_PORT),
		betypes.CERT_PATH, betypes.KEY_PATH, nil))
}
