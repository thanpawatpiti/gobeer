package domain

import (
	"context"
	"mime/multipart"

	"github.com/thanpawatpiti/gobeer/pkg/domain/entities"
)

type BeerSerives interface {
	GetBeer(ctx context.Context, name string, page int, per_page int) ([]interface{}, error)
	GetTotalBeer(ctx context.Context, name string) (int, error)
	AddBeer(ctx context.Context, input *entities.Beer, image *multipart.FileHeader) error
	UpdateBeer(ctx context.Context, input *entities.Beer, image *multipart.FileHeader, id int) error
	DeleteBeer(ctx context.Context, id int) error
	GetMenu(ctx context.Context) (interface{}, error)
}
