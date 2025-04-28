package services

type AuthService interface {
	CreateLogin(email string) (string, error)
}

type defaultAuthService struct {
}

type NewAuthServiceConfig struct {
}

func NewAuthService(cfg NewAuthServiceConfig) AuthService {
	return defaultAuthService{}
}

// ------

func (s defaultAuthService) CreateLogin(email string) (string, error) {
	// handle logic
	return email, nil
}
