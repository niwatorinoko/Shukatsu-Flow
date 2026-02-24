package company

import (
	"context"

	"shukatsu-flow/api/internal/domain/model"
)

type Repository interface {
	Create(ctx context.Context, name string) (model.Company, error)
}
