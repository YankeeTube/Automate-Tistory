package main

import (
	"LikeTistory/dispatch"
	"flag"
	"fmt"
	"github.com/cheggaaa/pb/v3"
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
	var wait sync.WaitGroup

	blog := flag.String("blog", "gmyankee", "<BlogId>.tistory.com\nex) -blog=gmyankee")
	post := flag.Int("post", 222, "<BlogId>.tistory.com/<PostId>\nex) -post=291")
	target := flag.Int("target", 21, "Settings to Max Like Count\nex) -target=10")
	flag.Parse()

	if *post == 0 {
		fmt.Printf(WarningColor, "\nSingle Mode is Required Post ID!\n")
		os.Exit(3)
	}
	if *target == 0 {
		fmt.Printf(WarningColor, "\nRequired to target count\n")
		os.Exit(3)
	}

	entryId := strconv.Itoa(*post)
	start := dispatch.CheckEntryCount(*blog, entryId, *target)
	if start == *target {
		fmt.Printf(InfoColor, "Already Count!\n")
		os.Exit(3)
	}
	maxRange := *target - start
	bar := pb.StartNew(maxRange)
	wait.Add(maxRange)
	go func() {
		for x := start; x < *target; x++ {
			bar.Increment()
			time.Sleep(time.Duration(1) * time.Second)
			result := dispatch.Response{}
			dispatch.EntryFetch(*blog, entryId, &result)
			if result.Content.ReactionCounter.Like > *target {
				wait.Done()
				break
			}
			wait.Done()
		}
	}()
	bar.Finish()
	wait.Wait()
	fmt.Printf(InfoColor, "Clear!\n")
}
