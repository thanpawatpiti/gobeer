package services

import (
	"context"
	"mime/multipart"

	"github.com/thanpawatpiti/gobeer/pkg/domain/entities"
	"github.com/thanpawatpiti/gobeer/pkg/services/helpers"
)

func (s *Service) AddBeer(ctx context.Context, input *entities.Beer, image *multipart.FileHeader) error {
	filePath := ""

	// check beer type
	err := s.repo.GetBeerType(ctx, input.BeerTypeID)
	if err != nil {
		return err
	}

	if image != nil {
		// Upload file
		file, err := helpers.UploadFile(ctx, image)
		if err != nil {
			return err
		}

		filePath = file
	}

	input.Image = filePath

	// Add beer
	err = s.repo.AddBeer(ctx, input)
	if err != nil {
		return err
	}

	// create logs
	err = s.repo.CreateLog(ctx, "add beer", input)
	if err != nil {
		return err
	}

	return nil
}
