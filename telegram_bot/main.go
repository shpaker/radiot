package main

import (
	rt "github.com/shpaker/radiot"
	"github.com/tucnak/telebot"

	"flag"
	"log"
	"strconv"
	"strings"
	"time"
)

var (
	pathToJsonFile string
	token          string
	updateTime     uint
	bot            *telebot.Bot
	rtJson         *rt.JsonFile
)

func init() {
	var err error

	flag.StringVar(&pathToJsonFile, "path", "./", "path to folder with radiot.json")
	flag.StringVar(&token, "token", "", "telegram's bot token string")
	flag.UintVar(&updateTime, "update", 30, "update time")
	flag.Parse()

	bot, err = telebot.NewBot(token)
	if err != nil {
		log.Fatal(err)
	}

	rtJson, err = rt.NewJsonFile(pathToJsonFile)
	if err != nil {
		log.Fatal(err)
	}

}

func updateFromSite() {
	for {
		rtJson.UpdateJsonFile()
		time.Sleep(time.Duration(updateTime) * time.Minute)
	}
}

func main() {

	go updateFromSite()

	messages := make(chan telebot.Message)
	bot.Listen(messages, 1*time.Second)

	for message := range messages {
		var text string
		switch {
		case strings.SplitN(message.Text, " ", 2)[0] == "/show":
			if strings.Count(message.Text, " ") > 0 {
				epId, err := strconv.Atoi(strings.SplitN(message.Text, " ", 2)[1])
				if err != nil || epId < 0 || epId > rtJson.GetLatestEpisodeId() {
					text = "🖐 Указанного выпуска не существует"
				} else {
					if text, err = rtJson.GetEpisodeMessage(epId); err != nil {
						text = "😡 Указанный выпуск не найден"
					}
				}
			} else {
				text, _ = rtJson.GetEpisodeMessage(rtJson.GetLatestEpisodeId())
			}
			bot.SendMessage(message.Chat, text, nil)
		case message.Text == "/latest":
			text, _ = rtJson.GetEpisodeMessage(rtJson.GetLatestEpisodeId())
			bot.SendMessage(message.Chat, text, nil)
		case message.Text == "/about":
			text, _ = rtJson.GetEpisodeMessage(rtJson.GetLatestEpisodeId())
			bot.SendMessage(message.Chat, `🙈 Author: https://telegram.me/shpaker
🙊 Sources: https://github.com/shpaker/radiot`, nil)
		}
	}
}
