package bcra

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

type APIConfig struct {
	host string
	path string
}

type APIRequest interface {
	params() url.Values
}

// Client may be used to make requests to the BCRA APIs
type Client struct {
	Debug      bool
	httpClient *http.Client
}

// NewClient constructs a new Client which can make requests to the BCRA
// WebService APIs.
func NewClient() (*Client, error) {
	client := &Client{Debug: false, httpClient: &http.Client{}}

	return client, nil
}

func (api *Client) GetJson(config *APIConfig, apiRequest APIRequest, apiResponse interface{}) error {
	httpResponse, err := api.Request(config, apiRequest)

	if err != nil {
		return err
	}

	defer httpResponse.Body.Close()

	err = json.NewDecoder(httpResponse.Body).Decode(apiResponse)

	return err
}

func (api *Client) Request(config *APIConfig, apiRequest APIRequest) (*http.Response, error) {
	req, err := http.NewRequest("GET", config.GetURL(), nil)
	if err != nil {
		log.Println("Hello")
		return nil, err
	}

	return api.httpClient.Do(req)
}

func (apiconf *APIConfig) AddPathValues(values ...any) {
	apiconf.path = fmt.Sprintf(apiconf.path, values...)
}

func (apiconf *APIConfig) GetURL() string {
	result, err := url.JoinPath(apiconf.host, apiconf.path)
	if err != nil {
		panic(err)
	}
	return result
}
