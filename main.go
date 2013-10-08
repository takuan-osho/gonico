package main

import (
	"flag"
	"fmt"
	"runtime"

	"gonico"
)

const (
	urlHelp = "The movie url or id from which you want to extract movie infomation."
)

func main() {
	movieUrl := flag.String("url", "foo", urlHelp)
	flag.Parse()

	runtime.GOMAXPROCS(runtime.NumCPU())

	movieInfo := gonico.GetVideoInfo(*movieUrl)
	fmt.Println(movieInfo)
}
