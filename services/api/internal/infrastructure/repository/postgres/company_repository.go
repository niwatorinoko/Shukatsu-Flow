package postgres

import (
	"context"
	"database/sql"

	"shukatsu-flow/api/internal/domain/model"
	companyUsecase "shukatsu-flow/api/internal/usecase/company"
)

type CompanyRepository struct {
	databaseConnection *sql.DB
}

var _ companyUsecase.Repository = (*CompanyRepository)(nil)

func NewCompanyRepository(databaseConnection *sql.DB) *CompanyRepository {
	return &CompanyRepository{
		databaseConnection: databaseConnection,
	}
}

func (companyRepository *CompanyRepository) ListCompanies(
	contextObject context.Context,
) ([]model.Company, error) {
	query := `
		SELECT
			id,
			user_id,
			name,
			industry,
			job_type,
			preference_level,
			memo,
			created_at,
			updated_at
		FROM companies
		ORDER BY created_at DESC
	`

	rows, queryError := companyRepository.databaseConnection.QueryContext(contextObject, query)
	if queryError != nil {
		return nil, queryError
	}
	defer rows.Close()

	companies := []model.Company{}

	for rows.Next() {
		var company model.Company

		scanError := rows.Scan(
			&company.Id,
			&company.UserId,
			&company.Name,
			&company.Industry,
			&company.JobType,
			&company.PreferenceLevel,
			&company.Memo,
			&company.CreatedAt,
			&company.UpdatedAt,
		)
		if scanError != nil {
			return nil, scanError
		}

		companies = append(companies, company)
	}

	if rowsError := rows.Err(); rowsError != nil {
		return nil, rowsError
	}

	return companies, nil
}

func (companyRepository *CompanyRepository) CreateCompany(
	contextObject context.Context,
	company model.Company,
) (model.Company, error) {
	query := `
		INSERT INTO companies (
			id,
			user_id,
			name,
			industry,
			job_type,
			preference_level,
			memo,
			created_at,
			updated_at
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING
			id,
			user_id,
			name,
			industry,
			job_type,
			preference_level,
			memo,
			created_at,
			updated_at
	`

	var createdCompany model.Company

	queryRow := companyRepository.databaseConnection.QueryRowContext(
		contextObject,
		query,
		company.Id,
		company.UserId,
		company.Name,
		company.Industry,
		company.JobType,
		company.PreferenceLevel,
		company.Memo,
		company.CreatedAt,
		company.UpdatedAt,
	)

	scanError := queryRow.Scan(
		&createdCompany.Id,
		&createdCompany.UserId,
		&createdCompany.Name,
		&createdCompany.Industry,
		&createdCompany.JobType,
		&createdCompany.PreferenceLevel,
		&createdCompany.Memo,
		&createdCompany.CreatedAt,
		&createdCompany.UpdatedAt,
	)
	if scanError != nil {
		return model.Company{}, scanError
	}

	return createdCompany, nil
}
