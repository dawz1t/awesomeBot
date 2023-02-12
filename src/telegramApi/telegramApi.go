package telegramApi

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func CreateConnection(key string) (tgbotapi.BotAPI, tgbotapi.UpdatesChannel, error) {

	bot, err := tgbotapi.NewBotAPI(key)
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)
	updates.Clear()

	return *bot, updates, err

}
func SendImage(bot tgbotapi.BotAPI, document string, chatId string, replyToMessageId string) error {

	params := make(map[string]string)
	params["document"] = document
	params["chat_id"] = chatId
	params["reply_to_message_id"] = replyToMessageId

	_, err := bot.MakeRequest("sendDocument", params)

	return err
}

func SendMesage(bot tgbotapi.BotAPI, text string, chatId string, replyToMessageId string) error {

	params := make(map[string]string)
	params["text"] = text
	params["chat_id"] = chatId
	params["reply_to_message_id"] = replyToMessageId

	_, err := bot.MakeRequest("sendMessage", params)

	return err
}

func SendSticker(bot tgbotapi.BotAPI, text string, chatId string, replyToMessageId string) error {

	params := make(map[string]string)
	params["sticker"] = text
	params["chat_id"] = chatId
	params["reply_to_message_id"] = replyToMessageId

	_, err := bot.MakeRequest("sendSticker", params)

	return err
}
