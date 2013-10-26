package main

import (
	"flag"
	"fmt"
	"log"
	"runtime"
)

const (
	urlHelp = "The video url or id from which you want to extract video infomation."
)

func main() {
	videoUrl := flag.String("url", "foo", urlHelp)
	flag.Parse()

	runtime.GOMAXPROCS(runtime.NumCPU())

	resp, err := GetVideoThumbResponse(*videoUrl)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp)
}
