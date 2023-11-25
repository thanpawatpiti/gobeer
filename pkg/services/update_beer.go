package services

import (
	"context"
	"mime/multipart"

	"github.com/thanpawatpiti/gobeer/pkg/domain/entities"
	"github.com/thanpawatpiti/gobeer/pkg/services/helpers"
)

func (s *Service) UpdateBeer(ctx context.Context, entity *entities.Beer, image *multipart.FileHeader, id int) error {
	// check beer type
	if entity.BeerTypeID > 0 {
		err := s.repo.GetBeerType(ctx, entity.BeerTypeID)
		if err != nil {
			return err
		}
	}

	if image != nil {
		// Upload file
		file, err := helpers.UploadFile(ctx, image)
		if err != nil {
			return err
		}

		entity.Image = file
	}

	// Update beer
	err := s.repo.UpdateBeer(ctx, entity, id)
	if err != nil {
		return err
	}

	// create logs
	err = s.repo.CreateLog(ctx, "update beer", entity)
	if err != nil {
		return err
	}

	return nil
}
