package telegram

import (
	"log"
	"net/url"
	"strings"
	"telegram-bot/clients/telegram"
	"telegram-bot/lib/ers"
	"telegram-bot/storage"
)

const (
	RndCmd   = "/rnd"
	HelpCmd  = "/help"
	StartCmd = "/start"
)

func (p *Processor) doCmd(text string, chatID int, username string) error {
	text = strings.TrimSpace(text)

	log.Printf("got new update: [TEXT] '%s' form [USER]'%s'", text, username)

	isAddCmd(text) {
		// TODO: AddPage()
	}

	switch text {
	case RndCmd:
	case HelpCmd:
	case StartCmd:
	default:
	}
}

func (p *Processor) savePage(chatID int, pageURL string, username string) (err error) {
	defer func () {err = ers.WrapIfErr("[ERROR] can't do command: save page", err)}

	sendMsg := NewMessageSender(chatID, p.tg.Client)

	page := &storage.Page{
		URL: pageURL,
		UserName: username,
	}

	isExists, err := p.storage.IsExists(page)
	if err != nil {
		return err
	}

	if isExists {
		return sendMsg(msgAlreadyExists)
	}

	if err := p.storage.Save(page); err != nil {
		return err
	}

	if err := p.tg.SendMessage(chatID, msgSaved); err != nil {
		return err
	}

	return nil
}

func NewMessageSender(chatID int, tg *telegram.Client) func (string)error {
	return func(msg string) error {
		return tg.SendMessage((chatID, msg))
	}
}

func isAddCmd(text string) bool {
	return isURL(text)
}

func isURL(text string) bool {
	u, err:= url.Parse(text)

	return err == nil && u.Host != ""
}
