package handler

import (
	"errors"
	nethttp "net/http"

	"github.com/labstack/echo/v4"

	gen "shukatsu-flow/api/internal/interface/http/gen"
	httpMapper "shukatsu-flow/api/internal/interface/http/mapper"
	companyUsecase "shukatsu-flow/api/internal/usecase/company"
)

type CompanyHandler struct {
	companyUsecase companyUsecase.Usecase
}

var _ gen.ServerInterface = (*CompanyHandler)(nil)

func NewCompanyHandler(companyUsecase companyUsecase.Usecase) *CompanyHandler {
	return &CompanyHandler{
		companyUsecase: companyUsecase,
	}
}

func (companyHandler *CompanyHandler) ListCompanies(context echo.Context) error {
	companies, listCompaniesError := companyHandler.companyUsecase.ListCompanies(
		context.Request().Context(),
	)
	if listCompaniesError != nil {
		return context.JSON(
			nethttp.StatusInternalServerError,
			httpMapper.ToErrorResponse("INTERNAL_SERVER_ERROR", "failed to list companies"),
		)
	}

	companiesListResponse := httpMapper.ToCompaniesListResponse(companies)

	return context.JSON(nethttp.StatusOK, companiesListResponse)
}

func (companyHandler *CompanyHandler) CreateCompany(context echo.Context) error {
	var createCompanyRequest httpMapper.CreateCompanyRequest

	bindRequestError := context.Bind(&createCompanyRequest)
	if bindRequestError != nil {
		return context.JSON(
			nethttp.StatusBadRequest,
			httpMapper.ToErrorResponse("INVALID_REQUEST", "request body is invalid"),
		)
	}

	createCompanyInput := httpMapper.ToCreateCompanyInput(createCompanyRequest)

	createdCompany, createCompanyError := companyHandler.companyUsecase.CreateCompany(
		context.Request().Context(),
		createCompanyInput,
	)
	if createCompanyError != nil {
		if errors.Is(createCompanyError, companyUsecase.ErrCompanyNameIsRequired) {
			return context.JSON(
				nethttp.StatusBadRequest,
				httpMapper.ToErrorResponse("INVALID_REQUEST", createCompanyError.Error()),
			)
		}

		if errors.Is(createCompanyError, companyUsecase.ErrPreferenceLevelMustBeBetweenOneAndFive) {
			return context.JSON(
				nethttp.StatusBadRequest,
				httpMapper.ToErrorResponse("INVALID_REQUEST", createCompanyError.Error()),
			)
		}

		return context.JSON(
			nethttp.StatusInternalServerError,
			httpMapper.ToErrorResponse("INTERNAL_SERVER_ERROR", "failed to create company"),
		)
	}

	companyResponse := httpMapper.ToCompanyResponse(createdCompany)

	return context.JSON(nethttp.StatusCreated, companyResponse)
}
