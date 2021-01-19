package main

import (
	"github.com/vulppine/cursed-fox-bot/bot"
	"log"
	"os"
)

func main() {
	if os.Getenv("TWITTER_APIKEY") != "" && os.Getenv("TWITTER_APISECRET") != "" {
		if os.Getenv("TWUSER_TOKEN") == "" && os.Getenv("TWUSER_SECRET") == "" {
			cursedfoxbot.GetTwitterAuth()
		} else {
			cursedfoxbot.MakeCursedFox()
		}
	} else {
		log.Fatal("You need to supply a Twitter consumer key and secret to use this! Set them as variables in your environment as TWITTER_APIKEY and TWITTER_APISECRET.")
	}
}
