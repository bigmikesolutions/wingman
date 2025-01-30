package iam

type Service struct{}

func New() *Service {
	return &Service{}
}

func (s *Service) SignIn(login, pass string) error {
	return nil
}
