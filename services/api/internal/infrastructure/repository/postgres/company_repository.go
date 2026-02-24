package postgres

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"shukatsu-flow/api/internal/clock"
	"shukatsu-flow/api/internal/domain/model"
)

type CompanyRepository struct {
	db  *sql.DB
	clk clock.Clock
}

func NewCompanyRepository(db *sql.DB, clk clock.Clock) *CompanyRepository {
	return &CompanyRepository{db: db, clk: clk}
}

func (r *CompanyRepository) Create(ctx context.Context, name string) (model.Company, error) {
	const q = `
INSERT INTO companies (name)
VALUES ($1)
RETURNING id, name, created_at;
`
	var c model.Company
	if err := r.db.QueryRowContext(ctx, q, name).Scan(&c.ID, &c.Name, &c.CreatedAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.Company{}, errors.New("failed to insert company")
		}
		return model.Company{}, err
	}

	if c.CreatedAt.IsZero() {
		// DBが返さない設計でも落ちないように保険
		c.CreatedAt = time.Now()
	}
	return c, nil
}
