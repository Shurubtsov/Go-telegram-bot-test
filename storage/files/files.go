package files

import (
	"encoding/gob"
	"errors"
	"math/rand"
	"os"
	"path/filepath"
	"telegram-bot/lib/ers"
	"telegram-bot/storage"
	"time"
)

type Storage struct {
	basePath string
}

const defalultPerm = 0774

var ErrNoSavedPages = errors.New("[ERROR] no saved pages")

func New(basePath string) Storage {
	return Storage{basePath: basePath}
}

func (s Storage) Save(page *storage.Page) (err error) {
	defer func() { err = ers.WrapIfErr("[ERROR] can't save", err) }()

	fPath := filepath.Join(s.basePath, page.UserName)

	if err := os.MkdirAll(fPath, defalultPerm); err != nil {
		return err
	}

	fName, err := fileName(page)
	if err != nil {
		return err
	}

	fPath = filepath.Join(fPath, fName)

	file, err := os.Create(fPath)
	if err != nil {
		return err
	}
	defer func() { _ = file.Close() }() // ignore errorClose

	if err := gob.NewEncoder(file).Encode(page); err != nil {
		return err
	}

	return nil
}

func (s Storage) PickRandom(userName string) (page *storage.Page, err error) {
	defer func() { err = ers.WrapIfErr("[ERROR] can't pick random page", err) }()

	path := filepath.Join(s.basePath, userName)

	files, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}
	if len(files) == 0 {
		return nil, ErrNoSavedPages
	}

	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(len(files))

	file := files[n]

	return s.decodePage(filepath.Join(path, file.Name()))
}

func (s Storage) decodePage(filePath string) (*storage.Page, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, ers.Wrap("[ERROR] can't decode page", err)
	}
	defer func() { _ = f.Close() }()

	var p storage.Page

	if err := gob.NewDecoder(f).Decode(&p); err != nil {
		return nil, ers.Wrap("[ERROR] decoding is failed", err)
	}

	return &p, nil
}

func fileName(p *storage.Page) (string, error) {
	return p.Hash()
}
