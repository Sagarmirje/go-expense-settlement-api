# Go Expense Tracker & Bill Splitter API

A professional REST API built with Go to track shared group expenses and calculate optimized debt settlements.
This project focuses on backend architecture, algorithmic thinking, and real-world financial workflow modeling.

---

## Problem Explanation

Managing shared expenses in groups (trips, hostels, roommates, events) requires tracking payments and determining who owes whom.
Instead of everyone transferring money to everyone else, an optimized settlement plan should minimize the number of transactions required to balance accounts.

This API solves that problem programmatically using a greedy settlement algorithm.

---

## Key Features

* Create users
* Record shared expenses with custom splits
* Compute real-time balances
* Generate minimal settlement transactions
* Thread-safe in-memory storage for concurrent API calls

---

## Design Decisions

**Clean Architecture**
The project is structured into packages for handlers, services, models, and utilities to ensure modularity and maintainability.

**Standard Library Only**
Implemented using Go's native packages (`net/http`, `encoding/json`) to demonstrate strong fundamentals without framework dependency.

**Thread-Safe Storage**
Data stored in Go maps and slices protected by `sync.RWMutex` to safely handle concurrent API requests.

**Consistent JSON Responses**
Utility helpers standardize error handling and success responses across all endpoints.

---

## Settlement Algorithm (Greedy Optimization)

The algorithm works in the following steps:

1. **Calculate Net Balances**
   Each payer gains the total amount paid, while participants in splits lose their share.

2. **Classify Users**
   Users are divided into:

   * Creditors (positive balance)
   * Debtors (negative balance)

3. **Greedy Matching**
   The largest debtor is matched with the largest creditor.
   Maximum possible amount is transferred between them.
   Balances update and process repeats until all balances reach zero.

This approach minimizes the number of transactions required to settle all debts.

---

## Example Settlement Scenario

If:

* Alice pays ₹300 for Alice, Bob, Charlie
* Bob pays ₹150 for Bob and Charlie

The system computes balances and generates:

```
Bob pays Alice ₹50  
Charlie pays Alice ₹100
```

This demonstrates optimized debt resolution.

---

## How to Run

Ensure Go is installed.

Run:

```
go run cmd/main.go
```

Server runs at:

```
http://localhost:9000
```

---

## Sample API Usage

### Create Users

```
curl.exe -X POST http://localhost:9000/users -d '{ \"id\": \"u1\", \"name\": \"Alice\" }'
```

### Add Expense

```
curl.exe -X POST http://localhost:9000/expenses -H "Content-Type: application/json" -d '{ \"id\": \"e1\", \"description\": \"Group Dinner\", \"total_amount\": 300, \"paid_by\": \"u1\", \"splits\": [ {\"user_id\": \"u1\", \"amount\": 100}, {\"user_id\": \"u2\", \"amount\": 100}, {\"user_id\": \"u3\", \"amount\": 100} ] }'
```

### View Balances

```
curl.exe http://localhost:9000/balances
```

### Settlement Plan

```
curl.exe http://localhost:9000/settle
```

---

## AI Assistance Disclosure

This project was developed with the assistance of AI tools to accelerate boilerplate generation and improve development productivity.

**IDE Used:** Antigravity IDE
**AI Model Used:** Gemini Flash

AI assistance was used for:

* Generating initial project scaffolding
* Suggesting REST endpoint structures
* Drafting the greedy settlement algorithm logic
* Improving documentation clarity

However, the following were performed independently:

* Understanding and verifying all generated code
* Structuring the final project architecture
* Testing API flows and debugging issues
* Validating algorithm correctness
* Writing explanations and documentation

AI served as a productivity assistant, while design decisions and implementation understanding remain my own.

---

## Prompts Used During Development

### Prompt 1 – Project Generation

```
Generate a Go REST API for shared expense tracking that supports users, expenses, balances, and optimized settlement transactions using a greedy algorithm. Use only Go standard library, clean architecture, and in-memory storage.
```

### Prompt 2 – Algorithm Clarification

```
Explain how to implement a greedy settlement algorithm that minimizes number of financial transactions among multiple users with positive and negative balances.
```

---

## Author

**Sagar Mirje**
Computer Science Engineering Student

This project demonstrates backend engineering skills, API design, concurrency awareness, and algorithmic problem solving.
