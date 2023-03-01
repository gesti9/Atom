package server

import (
	"atom/pkg"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Client struct {
	bot *tgbotapi.BotAPI
}

func New(apiKey string) *Client {
	bot, err := tgbotapi.NewBotAPI(apiKey)
	if err != nil {
		log.Panic(err)
	}

	return &Client{
		bot: bot,
	}
}

func (c *Client) SendMessage(text string, chatId int64) error {
	msg := tgbotapi.NewMessage(chatId, text)
	msg.ParseMode = "Markdown"
	_, err := c.bot.Send(msg)
	return err
}
func OlxParser() {

	url := "https://www.olx.kz/d/elektronika/kompyutery-i-komplektuyuschie/monitory/astana/?search%5Border%5D=created_at:desc"

	response, err := http.Get(url)
	Check(err)
	defer response.Body.Close()

	doc, err := goquery.NewDocumentFromReader(response.Body)
	Check(err)

	massUrl := make([]string, 0)
	doc.Find("div.css-oukcj3").Find("div.css-1sw7q4x").Each(func(index int, item *goquery.Selection) {
		// h6 := item.Find("h6")
		// title := strings.TrimSpace(h6.Text())
		url, _ := item.Find("a").Attr("href")
		// price := strings.TrimSpace(item.Find("div.css-u2ayx9").Find("p").Text())

		massUrl = append(massUrl, "https://www.olx.kz"+url)

	})
	// DataBase(title, "https://www.olx.kz"+url)

	for i := 0; i < len(massUrl); i++ {
		DataBase(massUrl[i])
		// fmt.Println(massUrl[i])
	}
	fmt.Println(len(massUrl))

}
func Check(err error) {

	if err != nil {
		fmt.Println(err)
	}
}

func DataBase(s string) {
	c := New(pkg.BOT_TOKEN)

	database, err := sql.Open("sqlite3", "./list.db")
	Check(err)
	statement, err := database.Prepare("CREATE TABLE IF NOT EXISTS olx (id INTEGER PRIMARY KEY, url TEXT)")
	Check(err)
	statement.Exec()
	rows, _ := database.Query("SELECT id, url FROM olx")
	var id int

	var url1 string
	chek := false
	for rows.Next() {
		if s == url1 {
			chek = true
		}
		rows.Scan(&id, &url1)
		// fmt.Printf("%d: %s \n", id, url1)

	}

	rows.Close()
	statement, err = database.Prepare("INSERT INTO olx ( url) VALUES ( ?)")
	Check(err)
	if chek == false {
		c.SendMessage(s, int64(452639799))
		statement.Exec(s)
	}

}
