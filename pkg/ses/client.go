package ses

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type Client interface {
	V1CodeExchange(email string) (V1ExchangeResponse, error)
}

type client struct {
	baseUrl       string
	defaultClient *http.Client
}

func NewSesClient(host, port string, customClient *http.Client) Client {
	if customClient == nil {
		customClient = http.DefaultClient
	}

	return client{
		defaultClient: customClient,
		baseUrl:       fmt.Sprintf("%s:%s", host, port),
	}
}

func (c client) V1CodeExchange(email string) (V1ExchangeResponse, error) {
	path, err := url.JoinPath(c.baseUrl, "/v1/code/exchange")
	if err != nil {
		return V1ExchangeResponse{}, err
	}
	resp, err := c.defaultClient.Get(path)
	if err != nil {
		return V1ExchangeResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return V1ExchangeResponse{}, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var body V1ExchangeResponse
	if err = json.NewDecoder(resp.Body).Decode(&body); err != nil {
		return V1ExchangeResponse{}, err
	}

	return V1ExchangeResponse{body.Code}, nil
}
