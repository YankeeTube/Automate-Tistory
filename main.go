package main

import (
	"LikeTistory/dispatch"
	"fmt"
)

func main() {
	//var clearIds []string
	//url := "https://gmyankee.tistory.com/reaction"
	id := "gmyankee"
	ids := dispatch.PostParser(id)
	for _, entryId := range ids {
		fmt.Println(entryId)
		res := dispatch.EntryFetch(id, entryId)
		fmt.Printf("%+v\n", res)
		fmt.Println("debug")
	}

	//res := fetch(url, 291, result)

}
