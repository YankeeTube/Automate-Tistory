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
type Response struct {
	Success bool `json:"success"`
	content struct {
		Id              string `json:"id"`
		Type            string `json:"type"`
		CategoryId      int    `json:"categoryId"`
		Meta            Meta
		ReactionCounter ReactionCounter
		CreatedDate     string `json:"createdDate"`
		UpdatedDate     string `json:"updatedDate"`
	}
}

func EntryFetch(id string, entryId string) interface{} {
	result := Response{}

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
	json.Unmarshal(body, &result)
	fmt.Printf("%+v\n", result)
	fmt.Println("debug")
	return result
}
