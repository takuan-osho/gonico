package main

import (
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"net/url"
	"os"
)

func Login() (resp *http.Response, err error) {
	err = godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	mail := os.Getenv("NICO_MAIL")
	password := os.Getenv("NICO_PASSWORD")
	v := url.Values{"mail": {mail}, "password": {password}}
	resp, err = http.PostForm(NicoAPIUrls["login"], v)
	return
}
