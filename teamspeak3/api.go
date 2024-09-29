package teamspeak3

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type APIClient struct {
	BaseURL       *url.URL
	VirtualServer string
	apiKey        string
	httpClient    http.Client
}

func NewAPIClient(baseURL string, virtualServer string, apiKey string) *APIClient {
	parsedURL, parseError := url.Parse(baseURL)
	if parseError != nil {
		panic(parseError)
	}

	return &APIClient{
		BaseURL:       parsedURL,
		VirtualServer: virtualServer,
		apiKey:        apiKey,
	}
}

func APIClientFromEnv() *APIClient {
	baseURL, parseError := url.Parse(getEnv("URL", "http://127.0.0.1:10080"))
	if parseError != nil {
		panic(parseError)
	}

	return &APIClient{
		BaseURL:       baseURL,
		VirtualServer: getEnv("VIRTUAL_SERVER", "1"),
		apiKey:        getEnv("API_KEY", ""),
	}
}

func (c *APIClient) Request(method string, path string, body io.Reader) *http.Request {
	req, _ := http.NewRequest(method, fmt.Sprintf("%s/%s%s", c.BaseURL, c.VirtualServer, path), body)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("x-api-key", c.apiKey)
	// debug.DdRequest(req, false)
	return req
}

func (c *APIClient) HandleResponse(res *http.Response, err error) (*http.Response, error) {
	if res.StatusCode == http.StatusOK {
		return res, nil
	}

	err_body, _ := io.ReadAll(res.Body)

	return res, fmt.Errorf(fmt.Sprintf("%v - %s", res.StatusCode, err_body))
}

func (c *APIClient) Get(path string) (*http.Response, error) {
	return c.HandleResponse(c.httpClient.Do(c.Request(http.MethodGet, path, nil)))
}

func (c *APIClient) Post(path string, body map[string]string) (*http.Response, error) {
	jsonBody, _ := json.Marshal(body)

	req := c.Request(
		http.MethodPost,
		path,
		bytes.NewBuffer(jsonBody),
	)

	req.Header.Add("Content-Type", "application/json")

	return c.HandleResponse(c.httpClient.Do(req))
}

func (c *APIClient) ListClients() ([]Client, error) {
	res, err := c.Post("/clientlist", map[string]string{
		"-uid":     "",
		"-away":    "",
		"-voice":   "",
		"-times":   "",
		"-groups":  "",
		"-info":    "",
		"-country": "",
		"-ip":      "",
		"-badges":  "",
	})

	if err != nil {
		return []Client{}, err
	}

	var client_list ClientList

	if err := json.NewDecoder(res.Body).Decode(&client_list); err != nil {
		return []Client{}, err
	}

	return client_list.Clients, nil
}

func (c *APIClient) GetClientById(clid string) (Client, error) {
	res, err := c.Post("/clientinfo", map[string]string{"clid": clid})

	if err != nil {
		return Client{}, err
	}

	var client_list ClientList

	if err := json.NewDecoder(res.Body).Decode(&client_list); err != nil {
		return Client{}, err
	}

	client := client_list.Clients[0]
	client.Id = clid

	return client, nil
}

func (c *APIClient) ListChannels() ([]Channel, error) {
	res, err := c.Post("/channellist", map[string]string{
		"-topic":        "",
		"-flags":        "",
		"-voice":        "",
		"-limits":       "",
		"-icon":         "",
		"-secondsempty": "",
	})

	if err != nil {
		return []Channel{}, err
	}

	var channel_list ChannelList

	if err := json.NewDecoder(res.Body).Decode(&channel_list); err != nil {
		return []Channel{}, err
	}

	return channel_list.Channels, nil
}

func (c *APIClient) GetChannelById(cid string) (Channel, error) {
	res, err := c.Post("/channelinfo", map[string]string{"cid": cid})

	if err != nil {
		return Channel{}, err
	}

	var channel_list ChannelList

	if err := json.NewDecoder(res.Body).Decode(&channel_list); err != nil {
		return Channel{}, err
	}

	channel := channel_list.Channels[0]
	channel.Id = cid

	return channel, nil
}

func (c *APIClient) MoveClient(client Client, cid string) (bool, error) {
	res, err := c.Post("/clientmove", map[string]string{
		"clid": client.Id,
		"cid":  cid,
	})

	if err != nil || res.StatusCode != 200 {
		return false, err
	}

	return true, nil
}
