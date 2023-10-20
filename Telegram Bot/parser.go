package main

import (
	"os"
	"strconv"
	"strings"
)

type BotConfig struct {
	api_link string
	chat_id  int64
}

func CombineData(link string, id int64) BotConfig {
	return BotConfig{
		api_link: link,
		chat_id:  id,
	}
}

func SplitData(link string) []string {
	return strings.Split(link, " ")
}

func GetEnvBool() bool {
	return strings.ToLower(os.Getenv("DEBUG_MODE")) == "true"
}

func ConvertChatId(str string) int64 {
	if n, err := strconv.ParseInt(str, 10, 64); err == nil {
		return n
	}
	return 0
}
