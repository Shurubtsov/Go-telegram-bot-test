package event_consumer

import (
	"log"
	"telegram-bot/events"
	"time"
)

type Consumer struct {
	fetcher    events.Fetcher
	proccessor events.Processor
	batchSize  int
}

func New(fetcher events.Fetcher, proccessor events.Processor, batchSize int) Consumer {
	return Consumer{
		fetcher:    fetcher,
		proccessor: proccessor,
		batchSize:  batchSize,
	}
}

func (c Consumer) Start() error {
	for {
		gotEvents, err := c.fetcher.Fetch(c.batchSize)
		if err != nil {
			log.Printf("[ERROR]consumer: %s", err.Error())

			continue
		}

		if len(gotEvents) == 0 {
			time.Sleep(1 * time.Second)

			continue
		}

		if err := c.handleEvents(gotEvents); err != nil {
			log.Print(err)

			continue
		}

	}
}

func (c *Consumer) handleEvents(events []events.Event) error {
	for _, event := range events {
		log.Printf("[INFO]got new event to consume: %s", event.Text)

		if err := c.proccessor.Process(event); err != nil {
			log.Printf("can't handle event: %s", err.Error())

			continue
		}
	}

	return nil
}
