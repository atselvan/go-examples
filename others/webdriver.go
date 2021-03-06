package main

import (
	"github.com/fedesog/webdriver"
	"log"
	"time"
)

func main() {
	chromeDriver := webdriver.NewChromeDriver("/usr/local/bin/chromedriver")
	err := chromeDriver.Start()
	if err != nil {
		log.Println(err)
	}
	desired := webdriver.Capabilities{"Platform": "Linux"}
	required := webdriver.Capabilities{}
	session, err := chromeDriver.NewSession(desired, required)
	if err != nil {
		log.Println(err)
	}
	err = session.Url("http://golang.org")
	if err != nil {
		log.Println(err)
	}
	time.Sleep(60 * time.Second)
	session.Delete()
	chromeDriver.Stop()
}
