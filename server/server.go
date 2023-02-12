package server

import (
	"atom/pkg"
	"atom/pkg/checkport"
	"atom/pkg/exel"
	"atom/pkg/parser"
	"fmt"
	"log"
	"strconv"
	"strings"
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

	// go task()
	for update := range updates {
		if update.Message != nil { // If we got a message
			// log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			reply := Reply(update.Message.Text)
			if reply == "False" {
				reply = update.Message.Text
			}
			// mass := make([]string, 0)
			sps := strings.Split(reply, "\n")

			if len(sps) == 5 {
				stringToTrim := strings.TrimSpace(sps[3])
				stringToTrim1 := strings.TrimSpace(sps[4])
				s1, _ := strconv.Atoi(stringToTrim)
				fmt.Println(len(stringToTrim))
				s2, _ := strconv.Atoi(stringToTrim1)
				fmt.Println(len(sps[4]))
				exel.ChangeExel(sps[0], sps[1], sps[2], s1, s2)
				file := tgbotapi.NewDocument(update.Message.From.ID, tgbotapi.FilePath("schet.xlsx"))
				file.ReplyToMessageID = update.Message.MessageID
				reply = "ГОТОВО"
				bot.Send(file)

			}

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, reply)

			bot.Send(msg)
		}
	}
}

func Reply(s string) string {
	switch s {
	case "/s":
		return `
		Отправьте 
		1. КОД НАЗНАЧЕНИЯ ПЛАТЕЖА
		2. БИН И АДРЕC В ФОРМАТЕ-->(БИН/ИИН 230440042123,ТОО Рон-41,г.Астана, пр.Мангилик Ел11)
		3. НАИМЕНОВАНИЕ ТОВАРА
		4. КОЛИЧЕСТВО
		5.СУММА
		`
	case "/2":
		o8000, p8000, err := checkport.Shellout("nc -vnz 188.130.130.88 8000")
		o8001, p8001, err := checkport.Shellout("nc -vnz 188.130.130.88 8001")
		o8002, p8002, err := checkport.Shellout("nc -vnz 188.130.130.88 8002")
		o8003, p8003, err := checkport.Shellout("nc -vnz 188.130.130.88 8003")
		o5656, p5656, err := checkport.Shellout("nc -vnz 188.130.130.88 5656")

		return o8000 + p8000 + o8001 + p8001 + o8002 + p8002 + o8003 + p8003 + o5656 + p5656
	case "/3":
		o8000, p8000, err := checkport.Shellout("nc -vnz 178.89.108.250 8000")
		o8001, p8001, err := checkport.Shellout("nc -vnz 178.89.108.250 8001")
		o8002, p8002, err := checkport.Shellout("nc -vnz 178.89.108.250 8002")
		o8003, p8003, err := checkport.Shellout("nc -vnz 178.89.108.250 8003")
		o8005, p8005, err := checkport.Shellout("nc -vnz 178.89.108.250 8005")
		o5657, p5657, err := checkport.Shellout("nc -vnzu 178.89.108.250 5657")
		return o8000 + p8000 + o8001 + p8001 + o8002 + p8002 + o8003 + p8003 + o8005 + p8005 + o5657 + p5657
	case "/4":
		o8001, p8001, err := checkport.Shellout("nc -vnz 85.159.27.8 8001")
		o8002, p8002, err := checkport.Shellout("nc -vnz 85.159.27.8 8002")
		o8003, p8003, err := checkport.Shellout("nc -vnz 85.159.27.8 8003")
		o5715, p5715, err := checkport.Shellout("nc -vnz 85.159.27.8 5715")
		return o8001 + p8001 + o8002 + p8002 + o8003 + p8003 + p5715 + o5715
	case "/7":
		o8000, p8000, err := checkport.Shellout("nc -vnz 176.110.125.109 8000")
		o8008, p8008, err := checkport.Shellout("nc -vnz 176.110.125.109 8008")
		o5821, p5821, err := checkport.Shellout("nc -vnzu 176.110.125.109 5821")
		if err != nil {
			log.Printf("error: %v\n", err)
		}
		return o5821 + "\n" + p5821 + "\n" + o8000 + "\n" + p8000 + "\n" + o8008 + "\n" + p8008
	}

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
