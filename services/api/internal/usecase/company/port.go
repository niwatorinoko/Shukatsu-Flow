package company

import (
	"context"

	"shukatsu-flow/api/internal/domain/model"
)

type Repository interface {
	ListCompaniesByUserId(context.Context, string) ([]model.Company, error)
	CreateCompany(context.Context, model.Company) (model.Company, error)
}
