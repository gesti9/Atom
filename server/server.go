package server

import (
	"atom/pkg"
	"atom/pkg/parser"
	"fmt"
	"log"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Server() {
	bot, err := tgbotapi.NewBotAPI(pkg.BOT_TOKEN)
	if err != nil {
		log.Panic(err)
	}

	// bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	go task()
	for update := range updates {
		if update.Message != nil { // If we got a message
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			reply := Reply(update.Message.Text)
			if reply == "False" {
				reply = update.Message.Text
			}

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, reply)
			msg.ReplyToMessageID = update.Message.MessageID

			bot.Send(msg)
		}
	}
}

func Reply(s string) string {
	if s == pkg.Item {
		return parser.JoomTovar() + "\n" + parser.KaspiTovar()
	}
	return "False"
}

func task() {
	for {
		fmt.Println("1")
		time.Sleep(time.Minute)
	}

}