<<<<<<< HEAD
# Go Expense Tracker & Bill Splitter

A professional REST API built with Go to track shared group expenses and calculate optimal debt settlements.

## Problem Explanation
Managing group expenses (like travel or shared housing) involves tracking who paid for what and how much others owe. At the end, instead of everyone paying everyone else, we need a "minimal transaction" plan where fewer money transfers happen to reach equilibrium.

## Design Decisions
1. **Clean Architecture**: Use separate packages for `handlers`, `services`, and `models` to maintain a separation of concerns.
2. **Standard Library Only**: Built using `net/http` and `encoding/json` to demonstrate proficiency with Go's core capabilities without external bloat.
3. **In-Memory Storage**: Uses Go Map and Slices protected by a `sync.RWMutex` to ensure thread-safety for concurrent API calls.
4. **JSON Helpers**: Utility functions for consistent API error and success reporting.

## Algorithm Explanation (Greedy Settlement)
The settlement algorithm works as follows:
1. **Calculate Net Balances**: Iterate all expenses. A payer gains the `TotalAmount`, while people in the splits lose their shared `Amount`.
2. **Identify Debtors & Creditors**: Split users into two groups: those with negative balances (debtors) and positive balances (creditors).
3. **Recursive/Greedy Match**:
   - Take the user who owes the most (largest negative).
   - Take the user who is owed the most (largest positive).
   - Transfer the maximum possible amount between them.
   - Update their balances and repeat until everyone's balance is zero.
This ensures the minimum number of transactions required to settle up.

## How to Run
1. Ensure you have [Go](https://go.dev/doc/install) installed.
2. Open a terminal in the project root.
3. Run:
   ```bash
   go run cmd/main.go
   ```
4. The API will be available at `http://localhost:9000`.

## Sample API Usage (Testing)

### 1. Create Users
**Bash:**
```bash
curl -X POST http://localhost:9000/users -d '{"id": "u1", "name": "Alice"}'
curl -X POST http://localhost:9000/users -d '{"id": "u2", "name": "Bob"}'
curl -X POST http://localhost:9000/users -d '{"id": "u3", "name": "Charlie"}'
```

**PowerShell (Windows):**
```powershell
# Use curl.exe with escaped quotes for PowerShell compatibility
curl.exe -X POST http://localhost:9000/users -d '{ \"id\": \"u1\", \"name\": \"Alice\" }'
curl.exe -X POST http://localhost:9000/users -d '{ \"id\": \"u2\", \"name\": \"Bob\" }'
curl.exe -X POST http://localhost:9000/users -d '{ \"id\": \"u3\", \"name\": \"Charlie\" }'
```

### 2. Add a Shared Expense
Alice pays $300, split equally between Alice, Bob, and Charlie.

**Bash:**
```bash
curl -X POST http://localhost:9000/expenses -H "Content-Type: application/json" -d '{
    "id": "e1",
    "description": "Group Dinner",
    "total_amount": 300,
    "paid_by": "u1",
    "splits": [
        {"user_id": "u1", "amount": 100},
        {"user_id": "u2", "amount": 100},
        {"user_id": "u3", "amount": 100}
    ]
}'
```

**PowerShell (Windows):**
```powershell
curl.exe -X POST http://localhost:9000/expenses -H "Content-Type: application/json" -d '{ \"id\": \"e1\", \"description\": \"Group Dinner\", \"total_amount\": 300, \"paid_by\": \"u1\", \"splits\": [ {\"user_id\": \"u1\", \"amount\": 100}, {\"user_id\": \"u2\", \"amount\": 100}, {\"user_id\": \"u3\", \"amount\": 100} ] }'
```

### 3. Get Balances
```bash
curl.exe http://localhost:9000/balances
```

### 4. Get Settlement Plan
```bash
curl.exe http://localhost:9000/settle
```

