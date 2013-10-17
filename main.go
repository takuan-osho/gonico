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

	videoInfo := gonico.GetVideoInfo(*videoUrl)
	fmt.Println(videoInfo)
}
