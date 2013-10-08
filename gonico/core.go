package gonico

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

var NicoAPIUrls = map[string]string{
	"getthumbinfo": "http://ext.nicovideo.jp/api/getthumbinfo/",
	"getflv":       "http://flapi.nicovideo.jp/api/getflv/",
}

func GetVideoInfo(movieId string) {
	resp, err := http.Get(NicoAPIUrls["getthumbinfo"] + movieId)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("status:", resp.Status)
	fmt.Println(string(body))
}
