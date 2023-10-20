package main

import (
	"os"
	"time"
)

var (
	chatId  string = os.Getenv("CHAT_ID")
	apiLink string = os.Getenv("API_LINK")
	apiKey  string = os.Getenv("API_KEY")
)

var (
	links = SplitData(apiLink)
	ids   = SplitData(chatId)
)

const SLEEP_DURATION = time.Second * time.Duration(1) //10 Seconds
