package main

import (
	"flag"
	"runtime"

	"gonico"
)

const (
	startHelp    = "Set start page"
	endHelp      = "Set end page"
	outHelp      = "Set download path"
	intervalHelp = "Polling interval time (min)"
)

var (
	start    = flag.Int("start", 1, startHelp)
	end      = flag.Int("end", 1, endHelp)
	out      = flag.String("out", "", outHelp)
	interval = flag.Int("interval", 30, intervalHelp)
)

func init() {
	flag.IntVar(start, "s", 1, startHelp)
	flag.IntVar(end, "e", 1, endHelp)
	flag.StringVar(out, "o", "", outHelp)
	flag.IntVar(interval, "i", 30, intervalHelp)
}

func main() {
	flag.Parse()
	runtime.GOMAXPROCS(runtime.NumCPU())

	gonico.GetVideoInfo("sm9")
}
