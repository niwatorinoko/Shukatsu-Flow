package handler

import (
	nethttp "net/http"

	"github.com/labstack/echo/v4"

	gen "shukatsu-flow/api/internal/interface/http/gen"
	httpMapper "shukatsu-flow/api/internal/interface/http/mapper"
)

func (companyHandler *CompanyHandler) ListApplications(
	context echo.Context,
	listApplicationsParams gen.ListApplicationsParams,
) error {
	return context.JSON(
		nethttp.StatusNotImplemented,
		httpMapper.ToErrorResponse("NOT_IMPLEMENTED", "ListApplications is not implemented yet"),
	)
}

func (companyHandler *CompanyHandler) CreateApplication(context echo.Context) error {
	return context.JSON(
		nethttp.StatusNotImplemented,
		httpMapper.ToErrorResponse("NOT_IMPLEMENTED", "CreateApplication is not implemented yet"),
	)
}

func (companyHandler *CompanyHandler) GetApplication(context echo.Context, id string) error {
	return context.JSON(
		nethttp.StatusNotImplemented,
		httpMapper.ToErrorResponse("NOT_IMPLEMENTED", "GetApplication is not implemented yet"),
	)
}

func (companyHandler *CompanyHandler) UpdateNextAction(context echo.Context, id string) error {
	return context.JSON(
		nethttp.StatusNotImplemented,
		httpMapper.ToErrorResponse("NOT_IMPLEMENTED", "UpdateNextAction is not implemented yet"),
	)
}

func (companyHandler *CompanyHandler) UpdateStatus(context echo.Context, id string) error {
	return context.JSON(
		nethttp.StatusNotImplemented,
		httpMapper.ToErrorResponse("NOT_IMPLEMENTED", "UpdateStatus is not implemented yet"),
	)
}

func (companyHandler *CompanyHandler) CreateChecklistItem(context echo.Context) error {
	return context.JSON(
		nethttp.StatusNotImplemented,
		httpMapper.ToErrorResponse("NOT_IMPLEMENTED", "CreateChecklistItem is not implemented yet"),
	)
}

func (companyHandler *CompanyHandler) UpdateChecklistItem(context echo.Context, id string) error {
	return context.JSON(
		nethttp.StatusNotImplemented,
		httpMapper.ToErrorResponse("NOT_IMPLEMENTED", "UpdateChecklistItem is not implemented yet"),
	)
}

func (companyHandler *CompanyHandler) GetDashboard(context echo.Context) error {
	return context.JSON(
		nethttp.StatusNotImplemented,
		httpMapper.ToErrorResponse("NOT_IMPLEMENTED", "GetDashboard is not implemented yet"),
	)
}

func (companyHandler *CompanyHandler) CreateEvent(context echo.Context) error {
	return context.JSON(
		nethttp.StatusNotImplemented,
		httpMapper.ToErrorResponse("NOT_IMPLEMENTED", "CreateEvent is not implemented yet"),
	)
}

func (companyHandler *CompanyHandler) GetUpcomingEvents(context echo.Context) error {
	return context.JSON(
		nethttp.StatusNotImplemented,
		httpMapper.ToErrorResponse("NOT_IMPLEMENTED", "GetUpcomingEvents is not implemented yet"),
	)
}

func (companyHandler *CompanyHandler) GetHealth(context echo.Context) error {
	return context.JSON(
		nethttp.StatusOK,
		map[string]bool{"ok": true},
	)
}

func (companyHandler *CompanyHandler) CreateInterview(context echo.Context) error {
	return context.JSON(
		nethttp.StatusNotImplemented,
		httpMapper.ToErrorResponse("NOT_IMPLEMENTED", "CreateInterview is not implemented yet"),
	)
}
