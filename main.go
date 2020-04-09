package main

import (
	"LikeTistory/dispatch"
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
	BlogId    int    `json:"blogId`
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

func fetch(url string, entryId int, target interface{}) interface{} {
	data := []byte(fmt.Sprintf(`{"reactionType": "LIKE", "entryId": %d}`, entryId))
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
	return target
	//return json.Unmarshal(body, &target)
}

func main() {
	//url := "https://gmyankee.tistory.com/reaction"
	id := "gmyankee"
	dispatch.PostParser(id)
	//result := Response{}
	//res := fetch(url, 291, result)
	//fmt.Printf("%+v\n", res)
}
