package response

import (
	"time"

	"shukatsu-flow/api/internal/domain/model"
)

type CompanyResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

func FromCompany(c model.Company) CompanyResponse {
	return CompanyResponse{
		ID:        c.ID,
		Name:      c.Name,
		CreatedAt: c.CreatedAt,
	}
}
