package main

import (
	"github.com/cukhoaimon/khoainats/pkg/ses"
)

func main() {
	ses.Start(
		ses.ServerConfig{
			Port:       "8765",
			Host:       "localhost",
			WebhookUrl: "",
		},
	)
}
