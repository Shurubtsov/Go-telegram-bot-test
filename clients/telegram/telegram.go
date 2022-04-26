package telegram

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"path"
	"strconv"
	"telegram-bot/lib/ers"
)

type Client struct {
	host     string
	basePath string
	client   http.Client
}

const (
	getUpdatesMethod  = "getUpdates"
	sendMessageMethod = "sendMessage"
)

// main method for initialize type of Client in "main package"
func New(host string, token string) Client {
	return Client{
		host:     host,
		basePath: newBasePath(token),
		client:   http.Client{},
	}
}

func (c *Client) SendMessage(chatID int, text string) error {
	q := url.Values{}
	q.Add("chat_id", strconv.Itoa(chatID))
	q.Add("text", text)

	_, err := c.doRequest(sendMessageMethod, q)
	if err != nil {
		return ers.Wrap("[ERROR] can't send message, went something wrong: %w", err)
	}
	return nil
}

// url path of our bot which requesting to Telegram API
func newBasePath(token string) string {
	return "bot" + token
}

// method for getUpdates from api which return us some json responses with args "offset" & "limit"
func (c *Client) Updates(offset int, limit int) (updates []Update, err error) {
	defer func() {
		err = ers.WrapIfErr("[ERROR] can't get updates", err)
	}()
	q := url.Values{}
	q.Add("offset", strconv.Itoa(offset))
	q.Add("offset", strconv.Itoa(limit))

	data, err := c.doRequest(getUpdatesMethod, q)
	if err != nil {
		return nil, err
	}

	var res UpdatesResponse

	if err := json.Unmarshal(data, &res); err != nil {
		return nil, err
	}

	return res.Result, nil
}

// function execute Request to api
func (c *Client) doRequest(method string, query url.Values) (data []byte, err error) {
	defer func() { err = ers.WrapIfErr("[ERROR] can't do request", err) }()

	u := url.URL{
		Host: c.host,
		Path: path.Join(c.basePath, method), // path -> useful package for create functionally urls for requests
	}

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}
	req.URL.RawQuery = query.Encode()

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
