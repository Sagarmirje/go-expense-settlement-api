package services

import (
	"expense-tracker/internal/models"
	"fmt"
	"math"
	"sort"
	"sync"
)

// ExpenseService handles all data storage and business logic in-memory
type ExpenseService struct {
	users    map[string]models.User
	expenses []models.Expense
	mu       sync.RWMutex // Protects maps/slices from concurrent access
}

func NewExpenseService() *ExpenseService {
	return &ExpenseService{
		users:    make(map[string]models.User),
		expenses: []models.Expense{},
	}
}

// AddUser stores a new user in memory
func (s *ExpenseService) AddUser(user models.User) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, exists := s.users[user.ID]; exists {
		return fmt.Errorf("user with ID %s already exists", user.ID)
	}
	s.users[user.ID] = user
	return nil
}

// ListUsers returns all users in the system
func (s *ExpenseService) ListUsers() []models.User {
	s.mu.RLock()
	defer s.mu.RUnlock()
	users := []models.User{}
	for _, u := range s.users {
		users = append(users, u)
	}
	return users
}

// AddExpense records a new shared expense
func (s *ExpenseService) AddExpense(exp models.Expense) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Validate payer exists
	if _, exists := s.users[exp.PaidBy]; !exists {
		return fmt.Errorf("payer %s does not exist", exp.PaidBy)
	}

	// Validate total amount matches splits
	var totalSplits float64
	for _, split := range exp.Splits {
		if _, exists := s.users[split.UserID]; !exists {
			return fmt.Errorf("user %s in split does not exist", split.UserID)
		}
		totalSplits += split.Amount
	}

	// Simple check for floating point precision (allowing 0.01 difference)
	if math.Abs(exp.TotalAmount-totalSplits) > 0.01 {
		return fmt.Errorf("total amount does not match sum of splits")
	}

	s.expenses = append(s.expenses, exp)
	return nil
}

// ListExpenses returns all recorded expenses
func (s *ExpenseService) ListExpenses() []models.Expense {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.expenses
}

// GetBalances calculates the net balance for every user
func (s *ExpenseService) GetBalances() []models.UserBalance {
	s.mu.RLock()
	defer s.mu.RUnlock()

	balanceMap := make(map[string]float64)

	// Initialize all users with 0 balance
	for id := range s.users {
		balanceMap[id] = 0
	}

	// Iterate through expenses to calculate net positions
	for _, exp := range s.expenses {
		// Payer gets back the total amount
		balanceMap[exp.PaidBy] += exp.TotalAmount

		// Each person in splits owes their split amount (negative)
		for _, split := range exp.Splits {
			balanceMap[split.UserID] -= split.Amount
		}
	}

	results := []models.UserBalance{}
	for id, bal := range balanceMap {
		results = append(results, models.UserBalance{
			UserID: id,
			Name:   s.users[id].Name,
			Amount: math.Round(bal*100) / 100, // Round to 2 decimals
		})
	}
	return results
}

// Settle calculates the minimum number of transactions to clear all debts
func (s *ExpenseService) Settle() []models.Settlement {
	balances := s.GetBalances()

	type person struct {
		id      string
		balance float64
	}

	var debtors []person
	var creditors []person

	// Categorize users into debtors (-ive) and creditors (+ive)
	for _, b := range balances {
		if b.Amount < -0.01 {
			debtors = append(debtors, person{b.UserID, b.Amount})
		} else if b.Amount > 0.01 {
			creditors = append(creditors, person{b.UserID, b.Amount})
		}
	}

	settlements := []models.Settlement{}

	// Greedy Settlement Algorithm:
	// Always match the largest debtor with the largest creditor
	for len(debtors) > 0 && len(creditors) > 0 {
		// Sort to find extremes
		sort.Slice(debtors, func(i, j int) bool { return debtors[i].balance < debtors[j].balance })
		sort.Slice(creditors, func(i, j int) bool { return creditors[i].balance > creditors[j].balance })

		d := &debtors[0]
		c := &creditors[0]

		// The amount to transfer is the minimum of what debtor owes and what creditor is owed
		payment := math.Min(math.Abs(d.balance), c.balance)
		payment = math.Round(payment*100) / 100

		settlements = append(settlements, models.Settlement{
			From:   d.id,
			To:     c.id,
			Amount: payment,
		})

		// Update balances
		d.balance += payment
		c.balance -= payment

		// Remove if settled
		if math.Abs(d.balance) < 0.01 {
			debtors = debtors[1:]
		}
		if math.Abs(c.balance) < 0.01 {
			creditors = creditors[1:]
		}
	}

	return settlements
}
