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

	switch status {
	case "ok":
		fmt.Println(resp.VideoInfo)
	default:
		fmt.Println(resp.ErrorInfo)
	}
}
