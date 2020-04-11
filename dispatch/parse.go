package dispatch

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"math"
	"net/http"
	"regexp"
	"strconv"
)

func PostParser(id string) {
	var (
		links    []string
		lastPage int
	)
	links, lastPage = GetInitialData(id)
	fmt.Println(links)
	fmt.Println(lastPage)
	//for page := 1; page < lastPage; page++{
	//	url := fmt.Sprintf("https://search.daum.net/search?nil_suggest=btn&w=blog&DA=PGD&q=site:%s.tistory.com&page=%d", id, page)
	//	resp := Fetch(url)
	//	pageLinks, _ := DocumentsExtract(resp)
	//	links = append(links, pageLinks)
	//}
}

func Fetch(url string) *http.Response {
	resp, err := http.Get(url)
	if err != nil {
		log.Panic("Post Parser Request Error... : ", err)
	}
	return resp
}

func DocumentsExtract(resp *http.Response) ([]string, *goquery.Document) {
	var links []string
	regex, _ := regexp.Compile("/([0-9]+)")
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	doc.Find(".f_url").Each(func(i int, selection *goquery.Selection) {
		link, ok := selection.Attr("href")
		if ok {
			fmt.Println(regex.FindStringSubmatch(link)[1])
			links = append(links, link)
		}
	})
	return links, doc
}

func GetInitialData(id string) ([]string, int) {
	url := fmt.Sprintf("https://search.daum.net/search?nil_suggest=btn&w=blog&DA=PGD&q=site:%s.tistory.com&page=1", id)
	resp := Fetch(url)
	links, doc := DocumentsExtract(resp)
	lastPage := GetLastPage(doc)
	return links, lastPage
}

func GetLastPage(doc *goquery.Document) int {
	var lastPage int
	regex, _ := regexp.Compile("([0-9]+)건")
	doc.Find(".txt_info").Each(func(i int, selection *goquery.Selection) {
		total, err := strconv.ParseFloat(regex.FindStringSubmatch(selection.Text())[1], 32)
		if err != nil {
			log.Fatal(err)
		}
		lastPage = int(math.Ceil(total / 10))
	})
	return lastPage
}
