package services

import (
	"context"
	"mime/multipart"

	"github.com/thanpawatpiti/gobeer/pkg/domain/entities"
	"github.com/thanpawatpiti/gobeer/pkg/services/helpers"
)

func (s *Service) AddBeer(ctx context.Context, input *entities.Beer, image *multipart.FileHeader) error {
	filePath := ""

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
	err := s.repo.AddBeer(ctx, input)
	if err != nil {
		return err
	}

	return nil
}
