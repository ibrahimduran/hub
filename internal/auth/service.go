package auth

type Service interface {
}

type service struct{}

func NewService() Service {
	s := service{}

	return &s
}
