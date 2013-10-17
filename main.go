package main

import (
	"flag"
	"fmt"
	"runtime"

	"gonico"
)

const (
	urlHelp = "The video url or id from which you want to extract video infomation."
)

func main() {
	videoUrl := flag.String("url", "foo", urlHelp)
	flag.Parse()

	runtime.GOMAXPROCS(runtime.NumCPU())

	resp, status := gonico.GetVideoThumbResponse(*videoUrl)
	if status != "ok" {
		fmt.Println(resp.ErrorInfo)
	} else {
		videoInfo := resp.VideoInfo
		fmt.Println(videoInfo)
	}
}
