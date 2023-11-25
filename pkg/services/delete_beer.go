package services

import "context"

func (s *Service) DeleteBeer(ctx context.Context, id int) error {
	// Delete beer
	err := s.repo.DeleteBeer(ctx, id)
	if err != nil {
		return err
	}

	// create logs
	err = s.repo.CreateLog(ctx, "delete beer", id)
	if err != nil {
		return err
	}

	return nil
}
