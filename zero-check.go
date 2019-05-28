package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

func parseURL(url string) (err error) {
	// Request the HTML page.
	res, err := http.Get(url)
	if err != nil {
		return
	}

	defer res.Body.Close()
	if res.StatusCode != 200 {
		err = fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
		return
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return
	}

	selector := ".t-gameTableItemStatsCount"
	err = fmt.Errorf("Not found %s", selector)

	doc.Find(selector).EachWithBreak(func(i int, s *goquery.Selection) bool {
		dailyPlayers, err1 := strconv.Atoi(s.Text())
		if err1 != nil {
			err = err1
			return false
		}

		fmt.Printf("Daily players %d: %d\n", i, dailyPlayers)
		if dailyPlayers != 0 {
			err = nil
			return false
		}
		return true
	})

	return
}

func main() {
	err := parseURL("https://dgaming.com/")
	if err != nil {
		panic(err)
	}
}
