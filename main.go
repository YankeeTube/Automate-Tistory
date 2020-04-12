package main

import (
	"LikeTistory/dispatch"
	"flag"
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

func main() {
	//var clearIds []string
	//url := "https://gmyankee.tistory.com/reaction"
	const (
		InfoColor    = "\033[1;34m%s\033[0m"
		NoticeColor  = "\033[1;36m%s\033[0m"
		WarningColor = "\033[1;33m%s\033[0m"
		ErrorColor   = "\033[1;31m%s\033[0m"
		DebugColor   = "\033[0;36m%s\033[0m"
	)
	var (
		wait sync.WaitGroup
	)

	blogId := flag.String("blog", "gmyankee", "<BlogId>.tistory.com")
	single := flag.Bool("only", true, "Like only single post => true or false")
	postId := flag.Int("Post Id", 200, "Single Mode is Required Post ID => <BlogId>.tistory.com/<PostId>")
	target := flag.Int("target Like Count", 0, "Settings to Max Like Count")

	if *single == true {
		if *postId == 0 {
			fmt.Printf(WarningColor, "\nSingle Mode is Required Post ID!\n")
			os.Exit(3)
		}
		entryId := strconv.Itoa(*postId)
		start := dispatch.CheckEntryCount(*blogId, entryId, *target)
		if start == *target {
			fmt.Printf(InfoColor, "Already Count!\n")
			os.Exit(3)
		}
		for x := start; x < *target; x++ {
			result := dispatch.Response{}
			dispatch.EntryFetch(*blogId, entryId, &result)
			if result.Content.ReactionCounter.Like > *target {
				break
			}
		}
		fmt.Printf(InfoColor, "Clear!\n")
	} else {
		fmt.Printf(WarningColor, "\nDue to heavy traffic, it can be blocked and interrupted.\n")
		ids := dispatch.PostParser(*blogId)
		wait.Add(len(ids))
		completes := make([]string, len(ids))
		go func() {
			for _, entryId := range ids {
				time.Sleep(time.Duration(2) * time.Second)
				start := dispatch.CheckEntryCount(*blogId, entryId, *target)
				if start == *target {
					wait.Done()
					continue
				}
				for x := start; x < *target; x++ {
					fmt.Println(entryId)
					result := dispatch.Response{}
					dispatch.EntryFetch(*blogId, entryId, &result)
					if result.Content.ReactionCounter.Like > *target {
						completes = append(completes, entryId)
						wait.Done()
						break
					}
				}
				wait.Done()
			}
		}()
	}
	wait.Wait()
	fmt.Println("Complete")

	//res := fetch(url, 291, result)

}
