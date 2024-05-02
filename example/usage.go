package example

import "context"

type Service struct {
	PersonStore PersonStore
}

func NewService(personStore PersonStore) *Service {
	return &Service{
		PersonStore: personStore,
	}
}

func (s *Service) GetPerson(ctx context.Context, id string) (*Person, error) {
	return s.PersonStore.Get(ctx, id)
}
