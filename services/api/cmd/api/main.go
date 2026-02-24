package main

import (
	"log"
	"net/http"
	"time"

	"shukatsu-flow/api/internal/clock"
	"shukatsu-flow/api/internal/config"
	"shukatsu-flow/api/internal/infrastructure/db/postgres"
	postgresrepo "shukatsu-flow/api/internal/infrastructure/repository/postgres"
	"shukatsu-flow/api/internal/interface/http/router"
	"shukatsu-flow/api/internal/usecase/company"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("config load failed: %v", err)
	}

	db, err := postgres.Open(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("db open failed: %v", err)
	}
	defer db.Close()

	clk := clock.Real{}

	// Repository
	companyRepo := postgresrepo.NewCompanyRepository(db, clk)

	// Usecase
	companyUC := company.NewUsecase(companyRepo)

	// Router (DI)
	mux := router.New(router.Dependencies{
		CompanyUsecase: companyUC,
	})

	srv := &http.Server{
		Addr:              ":" + cfg.AppPort,
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
	}

	log.Printf("API listening on http://localhost:%s", cfg.AppPort)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("server stopped: %v", err)
	}
}
