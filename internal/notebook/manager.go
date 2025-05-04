package notebook

import (
	"github.com/cukhoaimon/khoainats/pkg/ses"
)

type Manager struct {
	SesClient ses.Client
}

func NewManager() Manager {
	sesClient := ses.NewSesClient("localhost", "8765", nil)

	return Manager{
		SesClient: sesClient,
	}
}

func (m Manager) ExchangeSesCode(email string) (ses.V1ExchangeResponse, error) {
	return m.SesClient.V1CodeExchange(email)
}
