package dispatch

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
)

func PostParser(id string) {
	links := make([]string, 0)
	pagination := 0
	url := fmt.Sprintf("https://search.daum.net/search?nil_suggest=btn&w=blog&DA=PGD&q=site:%s.tistory.com&page=%d", id, pagination)
	resp, err := http.Get(url)
	if err != nil {
		log.Panic("Post Parser Request Error... : ", err)
	}
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	doc.Find(".f_url").Each(func(i int, selection *goquery.Selection) {
		link, ok := selection.Attr("href")
		if ok {
			links = append(links, link)
		}
	})
	fmt.Println(links)
	//fmt.Printf("%s\n", body)
}
