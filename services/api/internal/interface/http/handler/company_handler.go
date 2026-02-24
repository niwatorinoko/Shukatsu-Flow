package handler

import (
	"encoding/json"
	"net/http"

	"shukatsu-flow/api/internal/interface/http/request"
	"shukatsu-flow/api/internal/interface/http/response"
	"shukatsu-flow/api/internal/usecase/company"
)

type CompanyHandler struct {
	uc *company.Usecase
}

func NewCompanyHandler(uc *company.Usecase) *CompanyHandler {
	return &CompanyHandler{uc: uc}
}

func (h *CompanyHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req request.CreateCompanyRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]any{"error": "invalid json"})
		return
	}

	c, err := h.uc.CreateCompany(r.Context(), req.Name)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]any{"error": err.Error()})
		return
	}

	writeJSON(w, http.StatusCreated, response.FromCompany(c))
}

func writeJSON(w http.ResponseWriter, status int, body any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(body)
}
