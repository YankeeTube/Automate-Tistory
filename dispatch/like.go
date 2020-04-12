package dispatch

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Meta struct {
	ServiceId string `json:"serviceId"`
	BlogId    int    `json:"blogId"`
	EntryId   int    `json:"entryId"`
}
type ReactionCounter struct {
	Sum  int `json:"sum"`
	Like int `json:"like"`
}
type ResponseEntry struct {
	ReactionCounter   ReactionCounter `json:"reactionCounter"`
	ReactionActivated string          `json:"reactionActivated"`
}
type Response struct {
	Success bool `json:"success"`
	Content struct {
		Id              string `json:"id"`
		Type            string `json:"type"`
		CategoryId      int    `json:"categoryId"`
		Meta            Meta
		ReactionCounter ReactionCounter
		CreatedDate     string `json:"createdDate"`
		UpdatedDate     string `json:"updatedDate"`
	}
}

func EntryFetch(id string, entryId string, target *Response) Response {
	url := fmt.Sprintf("https://%s.tistory.com/reaction", id)
	data := []byte(fmt.Sprintf(`{"reactionType": "LIKE", "entryId": %s}`, entryId))
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		log.Fatal("Error Request... : ", err.Error())
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.163 Safari/537.36")
	client := &http.Client{Timeout: time.Second * 10}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error Response... : ", err)
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &target)
	//fmt.Printf("%+v\n", target)  // Do u want view of `Like Count` ?
	return *target
}

func CheckEntryCount(id string, entryId string, target int) int {
	result := ResponseEntry{}
	blog := fmt.Sprintf("%s.tistory.com", id)
	url := fmt.Sprintf("https://%s/reaction?entryId=%s", blog, entryId)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("Error Request... : ", err.Error())
	}
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.132 Whale/2.7.97.12 Safari/537.36")
	req.Header.Set("Host", blog)

	client := &http.Client{Timeout: time.Second * 10}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error Request... : ", err.Error())
	} else if resp.StatusCode == 429 {
		log.Panic("Too Many Request... Quit System")
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &result)
	if result.ReactionCounter.Like >= target {
		return target
	}
	return result.ReactionCounter.Like
}
