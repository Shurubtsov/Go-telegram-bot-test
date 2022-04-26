package main

import (
	"flag"
	"log"
	"telegram-bot/clients/telegram"
)

func main() {

	tgClient := telegram.New(mustBotHost(), mustToken())

	//todo: fetcher = fetcher.New()

	//todo: processor = processor.New()

	//todo: consumer.Start(fetcher, processor)
}

func mustToken() string {
	token := flag.String("token-bot-token", "", "token for access to tg bot")

	flag.Parse()

	if *token == "" {
		log.Fatal("token is empty")
	}
	return *token
}

func mustBotHost() string {
	host := flag.String("bot-host", "", "get host for tg bot")

	if *host == "" {
		log.Fatal("our host is empty....try again")
	}

	return *host
}
