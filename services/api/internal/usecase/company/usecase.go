package company

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"

	"shukatsu-flow/api/internal/domain/model"
)

var ErrCompanyNameIsRequired = errors.New("company name is required")
var ErrPreferenceLevelMustBeBetweenOneAndFive = errors.New("preference level must be between 1 and 5")
var ErrUserIdIsRequired = errors.New("user id is required")

type CreateCompanyInput struct {
	UserId          string
	Name            string
	Industry        *string
	JobType         *string
	PreferenceLevel *int
	Memo            *string
}

type Usecase interface {
	ListCompanies(context.Context) ([]model.Company, error)
	CreateCompany(context.Context, CreateCompanyInput) (model.Company, error)
}

type usecase struct {
	companyRepository Repository
}

func NewUsecase(companyRepository Repository) Usecase {
	return &usecase{
		companyRepository: companyRepository,
	}
}

func (companyUsecase *usecase) ListCompanies(
	contextObject context.Context,
) ([]model.Company, error) {
	companies, listCompaniesError := companyUsecase.companyRepository.ListCompanies(contextObject)
	if listCompaniesError != nil {
		return nil, listCompaniesError
	}

	return companies, nil
}

func (companyUsecase *usecase) CreateCompany(
	contextObject context.Context,
	createCompanyInput CreateCompanyInput,
) (model.Company, error) {
	trimmedUserId := strings.TrimSpace(createCompanyInput.UserId)
	if trimmedUserId == "" {
		return model.Company{}, ErrUserIdIsRequired
	}

	trimmedCompanyName := strings.TrimSpace(createCompanyInput.Name)
	if trimmedCompanyName == "" {
		return model.Company{}, ErrCompanyNameIsRequired
	}

	if createCompanyInput.PreferenceLevel != nil {
		if *createCompanyInput.PreferenceLevel < 1 || *createCompanyInput.PreferenceLevel > 5 {
			return model.Company{}, ErrPreferenceLevelMustBeBetweenOneAndFive
		}
	}

	currentTime := time.Now().UTC()

	company := model.Company{
		Id:              uuid.NewString(),
		UserId:          trimmedUserId,
		Name:            trimmedCompanyName,
		Industry:        createCompanyInput.Industry,
		JobType:         createCompanyInput.JobType,
		PreferenceLevel: createCompanyInput.PreferenceLevel,
		Memo:            createCompanyInput.Memo,
		CreatedAt:       currentTime,
		UpdatedAt:       currentTime,
	}

	createdCompany, createCompanyError := companyUsecase.companyRepository.CreateCompany(contextObject, company)
	if createCompanyError != nil {
		return model.Company{}, createCompanyError
	}

	return createdCompany, nil
}
