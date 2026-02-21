package handlers

import (
	"encoding/json"
	"expense-tracker/internal/models"
	"expense-tracker/internal/services"
	"expense-tracker/internal/utils"
	"net/http"
)

type APIHandler struct {
	service *services.ExpenseService
}

func NewAPIHandler(service *services.ExpenseService) *APIHandler {
	return &APIHandler{service: service}
}

// CreateUser handles POST /users (create) and GET /users (list)
func (h *APIHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		users := h.service.ListUsers()
		utils.JSONResponse(w, http.StatusOK, users)
		return
	}

	if r.Method != http.MethodPost {
		utils.ErrorResponse(w, http.StatusMethodNotAllowed, "Only POST or GET are allowed")
		return
	}

	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Invalid JSON")
		return
	}

	if user.ID == "" || user.Name == "" {
		utils.ErrorResponse(w, http.StatusBadRequest, "ID and Name are required")
		return
	}

	if err := h.service.AddUser(user); err != nil {
		utils.ErrorResponse(w, http.StatusConflict, err.Error())
		return
	}

	utils.JSONResponse(w, http.StatusCreated, user)
}

// AddExpense handles POST /expenses (create) and GET /expenses (list)
func (h *APIHandler) AddExpense(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		expenses := h.service.ListExpenses()
		utils.JSONResponse(w, http.StatusOK, expenses)
		return
	}

	if r.Method != http.MethodPost {
		utils.ErrorResponse(w, http.StatusMethodNotAllowed, "Only POST or GET are allowed")
		return
	}

	var exp models.Expense
	if err := json.NewDecoder(r.Body).Decode(&exp); err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Invalid JSON")
		return
	}

	if err := h.service.AddExpense(exp); err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.JSONResponse(w, http.StatusCreated, exp)
}

// GetBalances handles GET /balances
func (h *APIHandler) GetBalances(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.ErrorResponse(w, http.StatusMethodNotAllowed, "Only GET is allowed")
		return
	}

	balances := h.service.GetBalances()
	utils.JSONResponse(w, http.StatusOK, balances)
}

// SettleDebt handles GET /settle
func (h *APIHandler) SettleDebt(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.ErrorResponse(w, http.StatusMethodNotAllowed, "Only GET is allowed")
		return
	}

	plan := h.service.Settle()
	utils.JSONResponse(w, http.StatusOK, plan)
}
