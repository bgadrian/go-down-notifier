package main

import (
	"flag"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/bgadrian/go-down-notifier/godown"
)

func main() {

	//chrome 60 default

	method := flag.String("method", "GET", "web request method (GET,POST...)")
	interval := flag.Uint("interval", 300, "seconds interval to make the check")
	webUrls := flag.String("web", "http://httpbin.org/get?a=test", "a CSV of URL's for web request checks")

	guser := flag.String("guser", "", "gmail username")
	gpass := flag.String("gpass", "", "gmail password")
	email := flag.String("mail", "", "email to receive alerts")

	flag.Parse()
	userAgent := "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/60.0.3112.113 Safari/537.36"
	checkEvery := time.Duration(*interval) * time.Second
	urls := strings.Split(*webUrls, ",")

	checkOneWeb := func(url string) {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in checkOneWeb", r)
			}
		}()

		err := godown.HTTPRequest(10, userAgent, *method, url)
		if err == nil {
			log.Println("OK: " + url)
		} else {
			//TODO implement each alert system in it's own channel? https://stackoverflow.com/questions/16930251/go-one-producer-many-consumers

			godown.LogPrintln(url, err)
			if *guser != "" {
				godown.Gmail(url, err, *guser, *email, *gpass)
			}
		}
	}

	for range time.Tick(checkEvery) {
		for _, url := range urls {
			go checkOneWeb(strings.TrimSpace(url))
		}
	}
}
