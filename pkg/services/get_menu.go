package services

import "context"

func (s *Service) GetMenu(ctx context.Context) (interface{}, error) {
	menu, err := s.repo.GetMenu(ctx)
	if err != nil {
		return nil, err
	}
	return menu, nil
}
