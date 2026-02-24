package company

import (
	"context"
	"errors"
	"strings"

	"shukatsu-flow/api/internal/domain/model"
)

type Usecase struct {
	repo Repository
}

func NewUsecase(repo Repository) *Usecase {
	return &Usecase{repo: repo}
}

func (u *Usecase) CreateCompany(ctx context.Context, name string) (model.Company, error) {
	name = strings.TrimSpace(name)
	if name == "" {
		return model.Company{}, errors.New("name is required")
	}
	if len(name) > 100 {
		return model.Company{}, errors.New("name is too long (max 100)")
	}
	return u.repo.Create(ctx, name)
}
