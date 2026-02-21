package main

import (
	"expense-tracker/internal/handlers"
	"expense-tracker/internal/services"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Initialize the singleton service (In-memory storage)
	expenseService := services.NewExpenseService()

	// Initialize handlers
	apiHandler := handlers.NewAPIHandler(expenseService)

	// Define Routes using Go standard library
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to Splitwise Clone API! \nAvailable endpoints: /users, /expenses, /balances, /settle")
	})
	http.HandleFunc("/users", apiHandler.CreateUser)
	http.HandleFunc("/expenses", apiHandler.AddExpense)
	http.HandleFunc("/balances", apiHandler.GetBalances)
	http.HandleFunc("/settle", apiHandler.SettleDebt)

	// Start Server
	port := 9000
	fmt.Printf("🚀 Expense Tracker API running on http://localhost:%d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
