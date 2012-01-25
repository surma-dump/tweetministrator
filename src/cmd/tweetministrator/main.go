package main

import (
	"flag"
	"log"
	"time"
	"strings"
)

var (
	configFlag = flag.String("config", "/etc/tweetministrator/tweetministrator.conf", "Path to configfile")
	helpFlag = flag.Bool("h", false, "Show this help")
	config *Config

	startupTime = time.Now()
)

func main() {
	flag.Parse()

	if *helpFlag {
		flag.PrintDefaults()
		return
	}

	config = ReadConfig(*configFlag)

	c := GetUserStream(config.ListenTo)
	for t := range c {
		date, e := time.Parse(time.RubyDate, t.Date)
		if e != nil {
			log.Printf("Could not parse date: %s", e)
			continue
		}
		if date.After(startupTime) {
			for _, cmd := range strings.Fields(t.Text) {
				ExecuteCommand(strings.ToLower(cmd))
			}
		}
	}
}

func GetUserStream(user string) (<-chan Tweet) {
	c := make(chan Tweet)
	go func() {
		done := make(map[string]bool)
		tick := Ticker(config.PollingInterval)
		for _ = range tick {
			ts, e := LatestTweets(user)
			if e != nil {
				log.Printf("Could not get tweets: %s", e)
				continue
			}
			for _, t := range ts {
				_, ok := done[t.Identifier]
				if !ok {
					c <- t
					done[t.Identifier] = true
				}
			}
		}
	}()
	return c
}

func Ticker(ms_intervall int64)  (<-chan bool) {
	c := make(chan bool)
	go func() {
		for {
			c <- true
			time.Sleep(time.Duration(ms_intervall) * time.Millisecond)
		}
	}()
	return c
}
