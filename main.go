package main

import (
	"awesomeBot/src/chatGPT"
	"awesomeBot/src/dallE"
	"awesomeBot/src/telegramApi"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {

	telegramKey, exists := os.LookupEnv("TELEGRAM_API_KEY")

	if exists {
		fmt.Println(telegramKey)
	}

	openAIkey, exists := os.LookupEnv("OPENAI_API_KEY")

	if exists {
		fmt.Println(openAIkey)
	}

	bot, updates, err := telegramApi.CreateConnection(telegramKey)
	if err != nil {
		log.Panic(err)
	}
	updates.Clear()

	for update := range updates {
		if update.Message != nil { // If we got a message
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
			randomSource := rand.NewSource(time.Now().UnixNano())
			random := rand.New(randomSource)
			randomInt := random.Intn(100)
			if strings.Count(strings.ToLower(update.Message.Text), "@davzitpudgebot нарисуй") >= 1 && randomInt <= 5 {
				log.Printf("Рандомное число: %d", randomInt)
				err := telegramApi.SendSticker(bot, "CAACAgIAAxkBAAEHshZj6GGRZ28Pp4jRDMxBiC_PXtD7LQACHRQAAn9hyEu5TVczRrvG6y4E", strconv.FormatInt(update.Message.Chat.ID, 10), strconv.Itoa(update.Message.MessageID))
				if err != nil {
					log.Panic(err)
				}
			} else if (strings.Count(strings.ToLower(update.Message.Text), "@davzitpudgebot нарисуй") >= 1) || (strings.Count(strings.ToLower(update.Message.Text), "@davzitpudgebot draw") >= 1) {
				log.Printf("==================================")
				image, err := dallE.GenerateImage(strings.Trim(update.Message.Text, "@davzitpudgebot"))
				if err == nil {
					var result dallE.GeneratedImage
					json.Unmarshal([]byte(image), &result)
					log.Printf("A: %s\n", string(image))
					log.Println(result.Data[0])
					err := telegramApi.SendImage(bot, result.Data[0].Url, strconv.FormatInt(update.Message.Chat.ID, 10), strconv.Itoa(update.Message.MessageID))
					if err != nil {
						log.Panic(err)
					}
				} else {
					err := telegramApi.SendMesage(bot, "Неверный запрос, либо запрос содержит непонятные слова", strconv.FormatInt(update.Message.Chat.ID, 10), strconv.Itoa(update.Message.MessageID))
					if err != nil {
						log.Panic(err)
					}
				}

			} else if strings.Count(strings.ToLower(update.Message.Text), "@davzitpudgebot") >= 1 {
				log.Printf("==================================")
				chat, err := chatGPT.SendQuestion(strings.Trim(update.Message.Text, "@davzitPudgeBot"))
				if err == nil {
					var result chatGPT.Responce

					json.Unmarshal([]byte(chat), &result)
					log.Printf("A: %s\n", string(chat))
					log.Println(result.Choices[0].Text)
					err := telegramApi.SendMesage(bot, result.Choices[0].Text, strconv.FormatInt(update.Message.Chat.ID, 10), strconv.Itoa(update.Message.MessageID))
					if err != nil {
						log.Panic(err)
					}
				} else {
					err := telegramApi.SendMesage(bot, "Неверный запрос, либо запрос содержит непонятные слова", strconv.FormatInt(update.Message.Chat.ID, 10), strconv.Itoa(update.Message.MessageID))
					if err != nil {
						log.Panic(err)
					}
				}
			}
		}
		/*if searchPudge(update.Message.Text, syn) {
			params := make(map[string]string)
			params["sticker"] = "CAACAgIAAxkBAAEHqplj5fP2XbfyaSiwfl-qi343QJrn5QACzBEAAp2EqUhkgIpPikPbkC4E"
			params["chat_id"] = strconv.FormatInt(update.Message.Chat.ID, 10)
			params["reply_to_message_id"] = strconv.Itoa(update.Message.MessageID)
			bot.MakeRequest("sendSticker", params)
		}*/

	}
	/*
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
		}*/
}

func searchPudge(s string, synonyms []string) bool {

	for _, v := range synonyms {
		if strings.Count(strings.ToLower(s), v) >= 1 {
			return true
		}
	}
	return false
}
