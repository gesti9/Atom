package parser

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
)

func JoomTovar() string {
	url := "https://www.joom.com/ru/products/62f661838ed09b01ebd4e0e2"
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error:  %d %s", res.StatusCode, res.Status)
	}
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	linkAll := doc.Find(".content___QyukV").Find(".label___Z2o2Y")
	productName := doc.Find(".card___XVq8N").Find(".name___uxWcB")
	name, _ := productName.Html()

	price, _ := linkAll.Html()
	return name + " " + price

}
func KaspiTovar() string {
	url := "https://kaspi.kz/shop/nur-sultan/c/smartphones/?q=%3Acategory%3ASmartphones%3AmanufacturerName%3AApple%3ASmartphones*Series%3AApple%20iPhone%2013&sort=price-asc"
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error:  %d %s", res.StatusCode, res.Status)
	}
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	linkAll := doc.Find(".item-card__debet ").Find(".item-card__prices-price")
	productName := doc.Find(".item-card__name")
	name, _ := productName.Html()

	price, _ := linkAll.Html()
	name = strings.TrimSpace(name)
	price = strings.TrimSpace(price)
	fmt.Println(name, price)
	return name + price

}

func LegalAvto() string {
	var result string
	url := "https://informburo.kz/novosti"
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error:  %d %s", res.StatusCode, res.Status)
	}
	doc, err := html.Parse(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < len(visit(nil, doc)); i++ {
		matched, err := regexp.MatchString(`legalizaciya`, visit(nil, doc)[i])
		matched1, err := regexp.MatchString(`legalizacii`, visit(nil, doc)[i])
		if err != nil {
			log.Fatal(err)
		}
		// fmt.Println(matched) // true
		if matched || matched1 {
			if result == visit(nil, doc)[i] {

			} else {
				result = visit(nil, doc)[i]
				fmt.Println(result)

			}

		}

	}
	return result
}

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}

	return links
}
