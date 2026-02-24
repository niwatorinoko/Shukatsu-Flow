package router

import (
	"net/http"

	"shukatsu-flow/api/internal/interface/http/handler"
	"shukatsu-flow/api/internal/usecase/company"
)

type Dependencies struct {
	CompanyUsecase *company.Usecase
}

func New(dep Dependencies) http.Handler {
	mux := http.NewServeMux()

	// health
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"ok":true}`))
	})

	// companies
	companyHandler := handler.NewCompanyHandler(dep.CompanyUsecase)
	mux.HandleFunc("POST /companies", companyHandler.Create)

	return mux
}
