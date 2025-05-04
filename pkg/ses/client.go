package ses

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
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
	// Prepare the request payload
	payload := V1ExchangeRequest{email}
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return V1ExchangeResponse{}, err
	}

	// Send the POST request
	resp, err := c.defaultClient.Post(
		c.baseUrl+"/v1/code/exchange",
		"application/json",
		bytes.NewReader(jsonData),
	)

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
