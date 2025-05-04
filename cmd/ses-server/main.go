package main

import (
	ses "github.com/cukhoaimon/khoainats/third_party/ses"
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
