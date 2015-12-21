package main

import (
	"github.com/Syfaro/telegram-bot-api"
	"github.com/shpaker/radiot"

	"log"
	"strconv"
	"strings"
)

var (
	rtb *radiot.JsonFile
)

func main() {
	// подключаемся к боту с помощью токена
	bot, err := tgbotapi.NewBotAPI("TOKEN")
	if err != nil {
		log.Panic(err)
	}

	rtb, err = radiot.NewJsonFile("./")
	if err != nil {
		log.Fatal(err)
	}
	err = rtb.UpdateJsonFile()
	if err != nil {
		log.Fatal()
	}

	// bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	// инициализируем канал, куда будут прилетать обновления от API
	var ucfg tgbotapi.UpdateConfig = tgbotapi.NewUpdate(0)
	ucfg.Timeout = 60
	err = bot.UpdatesChan(ucfg)
	// читаем обновления из канала
	for {
		select {
		case update := <-bot.Updates:

			var cmd, attr string

			chatID := update.Message.Chat.ID

			// Текст сообщения
			cmd = strings.Split(update.Message.Text, " ")[0]
			if len(strings.Split(update.Message.Text, " ")) > 1 {
				attr = strings.Split(update.Message.Text, " ")[1]
			}

			switch cmd {
			case "/latest":
				msg := tgbotapi.NewMessage(chatID, "reply")
				// и отправляем его
				bot.SendMessage(msg)
			case "/show":
				var id int
				var msg tgbotapi.MessageConfig
				if attr == "" {
					attr = strconv.Itoa(rtb.GetLatestEpisodeId())
				}
				id, err := strconv.Atoi(attr)
				if err != nil {
					msg = tgbotapi.NewMessage(chatID, "Неверно переданный номер выпуска")
				} else {
					text, _ := rtb.GetEpisodeMessage(id)
					msg = tgbotapi.NewMessage(chatID, text)
				}
				bot.SendMessage(msg)

			}
		}
	}
}
