package main

import (
	"encoding/json"
	"os"
)

type Config struct {
	PollingInterval int64
	ListenTo string
	Commands map[string][]string
}

func ReadConfig(path string) (c *Config) {
	f, e := os.Open(path)
	if e != nil {
		panic(e)
	}
	defer f.Close()

	c = &Config{}
	dec := json.NewDecoder(f)
	e = dec.Decode(c)
	if e != nil {
		panic(e)
	}
	return c
}
