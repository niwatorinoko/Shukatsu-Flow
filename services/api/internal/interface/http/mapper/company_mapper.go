package mapper

import (
	"time"

	"shukatsu-flow/api/internal/domain/model"
	companyUsecase "shukatsu-flow/api/internal/usecase/company"
)

type CreateCompanyRequest struct {
	Name            string  `json:"name"`
	Industry        *string `json:"industry"`
	JobType         *string `json:"job_type"`
	PreferenceLevel *int    `json:"preference_level"`
	Memo            *string `json:"memo"`
}

type CompanyResponse struct {
	Id              string    `json:"id"`
	UserId          string    `json:"user_id"`
	Name            string    `json:"name"`
	Industry        *string   `json:"industry,omitempty"`
	JobType         *string   `json:"job_type,omitempty"`
	PreferenceLevel *int      `json:"preference_level,omitempty"`
	Memo            *string   `json:"memo,omitempty"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type CompaniesListResponse struct {
	Data []CompanyResponse `json:"data"`
}

type CompanyEnvelope struct {
	Data CompanyResponse `json:"data"`
}

type ErrorDetail struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type ErrorResponse struct {
	Error ErrorDetail `json:"error"`
}

func ToCreateCompanyInput(
	userId string,
	createCompanyRequest CreateCompanyRequest,
) companyUsecase.CreateCompanyInput {
	return companyUsecase.CreateCompanyInput{
		UserId:          userId,
		Name:            createCompanyRequest.Name,
		Industry:        createCompanyRequest.Industry,
		JobType:         createCompanyRequest.JobType,
		PreferenceLevel: createCompanyRequest.PreferenceLevel,
		Memo:            createCompanyRequest.Memo,
	}
}

func ToCompany(company model.Company) CompanyResponse {
	return CompanyResponse{
		Id:              company.Id,
		UserId:          company.UserId,
		Name:            company.Name,
		Industry:        company.Industry,
		JobType:         company.JobType,
		PreferenceLevel: company.PreferenceLevel,
		Memo:            company.Memo,
		CreatedAt:       company.CreatedAt,
		UpdatedAt:       company.UpdatedAt,
	}
}

func ToCompaniesListResponse(companies []model.Company) CompaniesListResponse {
	companyResponses := make([]CompanyResponse, 0, len(companies))

	for _, company := range companies {
		companyResponses = append(companyResponses, ToCompany(company))
	}

	return CompaniesListResponse{
		Data: companyResponses,
	}
}

func ToCompanyResponse(company model.Company) CompanyEnvelope {
	return CompanyEnvelope{
		Data: ToCompany(company),
	}
}

func ToErrorResponse(errorCode string, errorMessage string) ErrorResponse {
	return ErrorResponse{
		Error: ErrorDetail{
			Code:    errorCode,
			Message: errorMessage,
		},
	}
}
