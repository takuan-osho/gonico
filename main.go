package main

import (
	"flag"
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

	gonico.GetVideoInfo(*movieUrl)
}
