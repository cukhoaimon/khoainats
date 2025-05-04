package ses

import (
	"testing"
)

func TestClient(t *testing.T) {
	// start server
	cfg := ServerConfig{
		Port:       "8989",
		Host:       "localhost",
		WebhookUrl: "",
	}
	go func() {
		Start(cfg)
	}()

	c := NewSesClient("http://"+cfg.Host, cfg.Port, nil)
	res, err := c.V1CodeExchange("phucdeptrai@gmail.com")
	if err != nil {
		t.Errorf("fail to run V1CodeExchange %v\n", err)
	}

	if res.Code == "" {
		t.Errorf("expected to have response code but found empty string")
	}
}
