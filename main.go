package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

//GetData Function to get the html
func GetData(url string) {

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Find the review items
	doc.Find("#Nse_Prc_tick").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		band := s.Find("strong").Text()
		fmt.Printf("stock price %d: %s \n", i, band)
	})
}

func main() {
	fmt.Printf("starting the stock-scraper\n")
	n := map[string]string{
		"Adani gas":       "https://www.moneycontrol.com/india/stockpricequote/miscellaneous/adanigaslimited/ADG01",
		"Rain Industries": "https://www.moneycontrol.com/india/stockpricequote/miscellaneous/adanigaslimited/ADG01",
	}

	// var urls = []string{"https://www.moneycontrol.com/india/stockpricequote/miscellaneous/adanigaslimited/ADG01",
	// 	"https://www.moneycontrol.com/india/stockpricequote/miscellaneous/adanigaslimited/ADG01"}

	for name, url := range n {
		fmt.Printf("Url %s : %s \n", name, url)
		GetData(url)
	}

}
