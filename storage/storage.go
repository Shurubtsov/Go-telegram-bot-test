package storage

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"io"
	"telegram-bot/lib/ers"
)

var ErrNoSavedPages = errors.New("[ERROR] no saved pages")

type Storage interface {
	Save(p *Page) error
	PickRandom(userName string) (*Page, error)
	Remove(p *Page) error
	IsExists(p *Page) (bool, error)
}

type Page struct {
	URL      string
	UserName string
}

func (p Page) Hash() (string, error) {
	h := sha1.New()

	if _, err := io.WriteString(h, p.URL); err != nil {
		return "", ers.Wrap("[ERROR] can't write hash", err)
	}

	if _, err := io.WriteString(h, p.URL); err != nil {
		return "", ers.Wrap("[ERROR] can't write hash", err)
	}
	return fmt.Sprintf("%x", h.Sum(nil)), nil
}
