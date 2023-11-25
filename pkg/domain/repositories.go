package domain

import (
	"context"
)

type BeerRepository interface {
	GetBeer(ctx context.Context, name string, page int, per_page int) ([]interface{}, error)
}
