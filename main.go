package main

import (
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strings"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("6152931647:AAGSLOSlRNkF2CxhWXLP6pgQfrr0Ml8_m2c")
	if err != nil {
		log.Panic(err)
	}
	syn := []string{"бучка",
		"бачер",
		"рудге",
		"пэйджер",
		"джадж",
		"hooker",
		"пиджак",
		"паджеро",
		"паджино",
		"чин чопа хук жопа",
		"рыдж",
		"пудге",
		"пудочка",
		"радж",
		"добряк дуф",
		"пудж",
		"модный или даже сексуальный c момом и башером",
		"пидж",
		"батчер",
		"паджильйон",
		"паджеро",
		"пуджислав",
		"pudge",
		"денди",
		"лучший керри",
		"пуджинашвили",
		"педжамин баттон",
		"пионер-паджатый",
		"падждержка",
		"ридж",
		"пудж"}
	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // If we got a message
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			if searchPudge(update.Message.Text, syn) {
				params := make(map[string]string)
				params["sticker"] = "CAACAgIAAxkBAAEHqplj5fP2XbfyaSiwfl-qi343QJrn5QACzBEAAp2EqUhkgIpPikPbkC4E"
				params["chat_id"] = strconv.FormatInt(update.Message.Chat.ID, 10)
				params["reply_to_message_id"] = strconv.Itoa(update.Message.MessageID)
				bot.MakeRequest("sendSticker", params)
			}

		}
	}
}

func searchPudge(s string, synonyms []string) bool {

	for _, v := range synonyms {
		if strings.Count(strings.ToLower(s), v) >= 1 {
			return true
		}
	}
	return false
}
