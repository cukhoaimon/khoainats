package notebook

import (
	ses2 "github.com/cukhoaimon/khoainats/third_party/ses"
)

type Manager struct {
	SesClient ses2.Client
}

func NewManager() Manager {
	sesClient := ses2.NewSesClient("localhost", "8765", nil)

	return Manager{
		SesClient: sesClient,
	}
}

func (m Manager) ExchangeSesCode(email string) (ses2.V1ExchangeResponse, error) {
	return m.SesClient.V1CodeExchange(email)
}
