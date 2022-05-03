package main

import (
	"flag"
	"log"
	tgClient "telegram-bot/clients/telegram"
	event_consumer "telegram-bot/consumer/event-consumer"
	"telegram-bot/events/telegram"
	"telegram-bot/storage/files"
)

const (
	tgBotHost   = "api.telegram.org"
	storagePath = "files_storage"
	batchSize   = 100
)

func main() {
	eventsProcessor := telegram.New(
		tgClient.New(tgBotHost, mustToken()),
		files.New(storagePath),
	)

	log.Print("service started")
	consumer := event_consumer.New(eventsProcessor, eventsProcessor, batchSize)

	if err := consumer.Start(); err != nil {
		log.Fatal("service is stopped", err)
	}
}

func mustToken() string {
	token := flag.String("tg-bot-token", "", "token for access to tg bot")

	flag.Parse()

	if *token == "" {
		log.Fatal("token is empty")
	}
	return *token
}

// func mustBotHost() string {
// 	host := flag.String("bot-host", "", "get host for tg bot")

// 	if *host == "" {
// 		log.Fatal("our host is empty....try again")
// 	}

// 	return *host
// }
